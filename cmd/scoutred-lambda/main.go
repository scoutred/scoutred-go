package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/scoutred/scoutred-go/client"
	"github.com/scoutred/scoutred-go/csv"
)

func main() {
	lambda.Start(HandleRequest)
}

func HandleRequest(ctx context.Context, s3Event events.S3Event) {
	fmt.Printf("%+v", s3Event)

	// new AWS session for interacting with services
	sess := session.Must(session.NewSession())

	// parse env parms
	env := Env{}
	if err := env.Read(); err != nil {
		log.Fatal(err)
	}

	// setup Scoutred client
	srClient := client.New(env.ScoutredAPIKey)

	// read request records
	for _, rec := range s3Event.Records {
		fmt.Printf("[%s - %s] Bucket = %s, Key = %s \n", rec.EventSource, rec.EventTime, rec.S3.Bucket.Name, rec.S3.Object.Key)

		// download the file from S3 to ephemeral storage
		inputFile, err := fetchS3Data(sess, rec.S3)
		if err != nil {
			log.Fatal(err)
		}

		// create a new temp file to write output
		outputFile, err := ioutil.TempFile("", "scoutred")
		if err != nil {
			log.Fatal(err)
		}

		switch {
		case env.CSVLonIdx != 0 && env.CSVLatIdx != 0:
			// setup our params
			csvParams := csv.LonLatParams{
				LonIdx: env.CSVLonIdx,
				LatIdx: env.CSVLatIdx,
				Params: csv.Params{
					Append: env.CSVAppend,
					Header: env.CSVHeader,
					Reader: inputFile,
					Writer: outputFile,
					Client: srClient,
				},
			}

			if err := csv.ZoningByLonLat(csvParams); err != nil {
				log.Fatal(err)
			}

		default:
			log.Fatal("missing CSV indexes to work with")
		}

		// pop off last entry for the file name
		parts := strings.Split(rec.S3.Object.Key, "/")

		// replace the folder name "input" with "output"
		moveToDest(sess, outputFile, env.DestS3Bucket, path.Join(env.DestS3Path, parts[len(parts)-1]))
	}
}

// fetch the file from s3 and slurp it into memory
func fetchS3Data(sess *session.Session, rec events.S3Entity) (*os.File, error) {
	// Create a downloader with the session and default options
	downloader := s3manager.NewDownloader(sess)

	tmpFile, err := ioutil.TempFile("", "scoutred")
	if err != nil {
		return nil, err
	}

	escapedKey, err := url.QueryUnescape(rec.Object.Key)
	if err != nil {
		return nil, err
	}

	// Write the contents of S3 Object to the file
	_, err = downloader.Download(tmpFile, &s3.GetObjectInput{
		Bucket: aws.String(rec.Bucket.Name),
		Key:    aws.String(escapedKey),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to download file, %v", err)
	}

	_, err = tmpFile.Seek(0, 0)
	if err != nil {
		return nil, err
	}

	return tmpFile, nil
}

// moveToDest takes a file and moves it to the destination location
func moveToDest(sess *session.Session, f *os.File, bucket, key string) error {
	_, err := f.Seek(0, 0)
	if err != nil {
		return nil
	}

	uploader := s3manager.NewUploader(sess)

	escapedKey, err := url.QueryUnescape(key)
	if err != nil {
		return err
	}

	result, err := uploader.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(escapedKey),
		Body:   f,
	})
	if err != nil {
		return fmt.Errorf("failed to upload file, %v", err)
	}

	fmt.Printf("file uploaded to, %s\n", result.Location)

	return nil
}
