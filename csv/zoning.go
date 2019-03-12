package csv

import (
	"encoding/csv"
	"fmt"
	"io"
	"net/http"
	"strconv"

	scoutred "github.com/scoutred/scoutred-go"
	"github.com/scoutred/scoutred-go/client"
	"github.com/scoutred/scoutred-go/log"
)

// zoningHeader are the additional zoning headers that will be added
var zoningHeader = []string{
	"designation",
}

// zoningRegulationHeader defines the fields associated with a zoning regulation.
// for each zoning record there could be many regulations. to account for this the
// records are flattened and each designation record will have an index appended to
// it (i.e. lot_size_min_area_1, lot_size_max_area_1, etc.)
var zoningRegulationHeader = []string{
	"lot_size_min_area",
	"lot_size_max_area",
	"lot_size_min_width",
	"lot_size_min_width_corner",
	"lot_size_min_frontage",
	"lot_size_min_depth",
	"density_sf_du",
	"density_sf_du_note",
	"height_limit_max",
	"height_limit_above_enclosed_parking",
	"height_limit_roof_flat",
	"height_limit_roof_pitched",
	"height_limit_note",
	"far_base",
	"far_min",
	"far_max",
	"far_residential",
	"far_commercial",
	"far_mixed",
	"far_note",
	"setbacks_front_min",
	"setbacks_front_min_note",
	"setbacks_front_max",
	"setbacks_front_max_note",
	"setbacks_interior_side",
	"setbacks_interior_side_note",
	"setbacks_street_side",
	"setbacks_street_side_note",
	"setbacks_rear",
	"setbacks_rear_alley",
	"setbacks_rear_note",
	"setbacks_general_note",
	"lot_coverage_min",
	"lot_coverage_max",
	"lot_coverage_note",
}

// ZoningByLonLat process a list of records with lon / lat values
// and decorates them with associated zoning details
func ZoningByLonLat(params LonLatParams) error {
	// zoning records could have 0 to many designations
	// because of this we need to track the total count of
	// designations in our process and then update the CSV
	// header after we're done processing the records
	var designationCount int

	// CSV reader
	r := csv.NewReader(params.Reader)

	// CSV writer
	w := csv.NewWriter(params.Writer)

	// iterate the records
	first := true
	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Error.Println("read error")
			return err
		}

		// skip processing the header
		if params.Header && first {
			first = false

			// decorate the first record with additional headers if we're in append mode
			if params.Append {
				record = append(record, zoningHeader...)
				record = append(record, zoningRegulationHeader...)
			} else {
				record = append(zoningHeader, zoningRegulationHeader...)
			}

			if err := w.Write(record); err != nil {
				return fmt.Errorf("error writing record to csv: %v", err)
			}

			// Write any buffered data to the underlying writer
			w.Flush()
			if err := w.Error(); err != nil {
				return err
			}

			continue
		}

		// error checking
		if len(record) < int(params.LonIdx) {
			return fmt.Errorf("invlaid index provided for longitude")
		}
		if len(record) < int(params.LatIdx) {
			return fmt.Errorf("invlaid index provided for latitude")
		}

		// value parsing
		lon, err := strconv.ParseFloat(record[params.LonIdx], 64)
		if err != nil {
			return fmt.Errorf("invalid float value for longitude: %v", err)
		}

		lat, err := strconv.ParseFloat(record[params.LatIdx], 64)
		if err != nil {
			return fmt.Errorf("invalid float value for latitude: %v", err)
		}

		// make the API request
		zoning, err := params.Client.ZoningByLonLat(lon, lat)
		if err != nil {
			// check if err is a client error
			apiErr, ok := err.(client.Error)
			if !ok {
				return err
			}

			switch apiErr.StatusCode {
			case http.StatusNotFound:
				// log the warning but don't stop processing
				log.Warn.Printf("record not found for LonLat value (%v, %v)", lon, lat)
			default:
				return err
			}
		}

		// we make the slice here incase there is no record. without this the CSV structure
		// could be inconsistent row to row
		zoningVals := make([]string, len(zoningHeader)+len(zoningRegulationHeader))

		if zoning != nil && zoning.Designation != nil {
			// this should map to zoningHeader
			zoningRec := []string{*zoning.Designation}

			zoningRegulationVals := []string{}
			// if we have zoning regulations lets add them to the output
			for i, reg := range zoning.Regulations {
				zoningRegulationVals = append(zoningRegulationVals, zoningRegulationToStringSlice(reg)...)

				// track our designation count max
				if i > designationCount {
					designationCount = i
				}

				// TODO(arolek): this should continue on, but requires us to juggle the header and record attributes
				// since we don't know the total count of attributes until we have processed all the records.
				// for now we will only produce a single regulation record
				break
			}

			// combine the regulation values to the zoning record
			zoningVals = append(zoningRec, zoningRegulationVals...)
		}

		// to append or not to append
		if params.Append {
			record = append(record, zoningVals...)
		} else {
			record = zoningVals
		}

		if err := w.Write(record); err != nil {
			return fmt.Errorf("error writing record to csv: %v", err)
		}

		// write any buffered data to the underlying writer
		w.Flush()

		// check for error
		if err := w.Error(); err != nil {
			return err
		}
	}

	return nil
}

// zoningRegulationToStringSlice maps a zoning regulation to a flat structure
// converting the values to strings for use as a CSV record
func zoningRegulationToStringSlice(reg scoutred.ZoningRegulation) []string {
	regulationVals := []string{}

	if reg.LotSize.MinArea != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.LotSize.MinArea), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.LotSize.MaxArea != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.LotSize.MaxArea), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.LotSize.MinWidth != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.LotSize.MinWidth), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.LotSize.MinWidthCorner != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.LotSize.MinWidthCorner), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.LotSize.MinFrontage != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.LotSize.MinFrontage), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.LotSize.MinDepth != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.LotSize.MinDepth), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Density.SfDu != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.Density.SfDu), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Density.SfDuNote != nil {
		regulationVals = append(regulationVals, *reg.Density.SfDuNote)
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.HeightLimit.Max != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.HeightLimit.Max), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.HeightLimit.AboveEnclosedParking != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.HeightLimit.AboveEnclosedParking), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.HeightLimit.RoofFlat != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.HeightLimit.RoofFlat), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.HeightLimit.RoofPitched != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.HeightLimit.RoofPitched), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.HeightLimit.Note != nil {
		regulationVals = append(regulationVals, *reg.HeightLimit.Note)
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.FAR.Base != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.FAR.Base), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.FAR.Min != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.FAR.Min), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.FAR.Max != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.FAR.Max), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.FAR.Residential != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.FAR.Residential), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.FAR.Commercial != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.FAR.Commercial), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.FAR.Mixed != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.FAR.Mixed), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.FAR.Note != nil {
		regulationVals = append(regulationVals, *reg.FAR.Note)
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.FrontMin != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.Setbacks.FrontMin), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.FrontMinNote != nil {
		regulationVals = append(regulationVals, *reg.Setbacks.FrontMinNote)
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.FrontMax != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.Setbacks.FrontMax), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.FrontMaxNote != nil {
		regulationVals = append(regulationVals, *reg.Setbacks.FrontMaxNote)
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.InteriorSide != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.Setbacks.InteriorSide), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.InteriorSideNote != nil {
		regulationVals = append(regulationVals, *reg.Setbacks.InteriorSideNote)
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.StreetSide != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.Setbacks.StreetSide), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.StreetSideNote != nil {
		regulationVals = append(regulationVals, *reg.Setbacks.StreetSideNote)
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.Rear != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.Setbacks.Rear), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.RearAlley != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.Setbacks.RearAlley), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.RearNote != nil {
		regulationVals = append(regulationVals, *reg.Setbacks.RearNote)
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.Setbacks.GeneralNote != nil {
		regulationVals = append(regulationVals, *reg.Setbacks.GeneralNote)
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.LotCoverage.Min != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.LotCoverage.Min), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.LotCoverage.Max != nil {
		regulationVals = append(regulationVals, strconv.FormatInt(int64(*reg.LotCoverage.Max), 10))
	} else {
		regulationVals = append(regulationVals, "")
	}

	if reg.LotCoverage.Note != nil {
		regulationVals = append(regulationVals, *reg.LotCoverage.Note)
	} else {
		regulationVals = append(regulationVals, "")
	}

	return regulationVals
}
