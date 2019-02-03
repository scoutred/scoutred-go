package main

import (
	"fmt"
	"os"
	"strconv"
)

// Env manages parsing various environment params
type Env struct {
	ScoutredAPIKey string
	DestS3Bucket   string
	DestS3Path     string
	CSVLatIdx      uint64
	CSVLonIdx      uint64
	CSVHeader      bool
	CSVAppend      bool
}

// Read reads various env vars, validates them and fills them into the Env struct
func (env *Env) Read() error {
	var err error

	key := os.Getenv("SCOUTRED_API_KEY")
	if key == "" {
		return fmt.Errorf("missing SCOUTRED_API_KEY")
	}
	env.ScoutredAPIKey = key

	latIdx := os.Getenv("CSV_LAT_IDX")
	if latIdx == "" {
		return fmt.Errorf("missing LAT_IDX")
	}
	env.CSVLatIdx, err = strconv.ParseUint(latIdx, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid value for CSV_LAT_IDX: %d : %v", latIdx, err)
	}

	lonIdx := os.Getenv("CSV_LON_IDX")
	if lonIdx == "" {
		return fmt.Errorf("missing LON_IDX")
	}
	env.CSVLonIdx, err = strconv.ParseUint(lonIdx, 10, 64)
	if err != nil {
		return fmt.Errorf("invalid value for CSV_LON_IDX: %d : %v", lonIdx, err)
	}

	env.DestS3Bucket = os.Getenv("DESTINATION_S3_BUCKET")
	if env.DestS3Bucket == "" {
		return fmt.Errorf("missing DESTINATION_S3_BUCKET")
	}

	env.DestS3Path = os.Getenv("DESTINATION_S3_PATH")
	if env.DestS3Path == "" {
		return fmt.Errorf("missing DESTINATION_S3_PATH")
	}

	csvHeader := os.Getenv("CSV_HEADER")
	if csvHeader != "" {
		env.CSVHeader, err = strconv.ParseBool(csvHeader)
		if err != nil {
			return fmt.Errorf("invalid value for CSV_HEADER: %v", csvHeader)
		}
	}

	csvAppend := os.Getenv("CSV_APPEND")
	if csvAppend != "" {
		env.CSVAppend, err = strconv.ParseBool(csvAppend)
		if err != nil {
			return fmt.Errorf("invalid value for CSV_APPEND: %v", csvAppend)
		}
	} else {
		// default to append
		env.CSVAppend = true
	}

	return nil
}
