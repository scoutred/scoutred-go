# Scoutred Lambda

Scoutred Lambda is used for bulk processing CSV files using AWS Lambda and Scoutred's API.

## Environment Variables

The following environment variables are necessary:

* `SCOUTRED_API_KEY`: Your Scoutred API Key.
* `DESTINATION_S3_BUCKET`: The bucket the files should be place in
* `DESTINATION_S3_PATH`: The basepath inside the `DESTINATION_S3_BUCKET` to place the output. 
* `CSV_LAT_IDX`: The column index in the CSV that contains the latitude field starting at 0.
* `CSV_LON_IDX`: The column index in the CSV that contains the longitude field starting at 0.
* `CSV_HEADER`: Boolean indicating if the CSV has a header. Defaults to false. 
* `CSV_APPEND`: Boolean indicating if the output should be appended to the incoming CSV file. Defaults to true.