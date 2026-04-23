package dv360

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"strings"
)

// convertCsvToArray Convert a csv content as rows of struct where csv headers are map keys
// Struct fields json tag values must match csv header (lower case & spaces replaced with _ )
// Ex csv with headers as 'Campaign ID', 'Impressions', 'Total Conversions'
// Structure json tag must be campaign_id, impressions, total_conversions
func convertCsvToArray[T any](r io.Reader) ([]*T, error) {
	csvReader := csv.NewReader(r)
	csvReader.Comma = ','
	csvReader.FieldsPerRecord = -1

	headers := []string{}
	results := make([]*T, 0)

	for {
		records, err := csvReader.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			return nil, err
		}

		if len(records) == 1 && records[0] == "No data returned by the reporting service." {
			return results, nil
		}

		if len(records) > 0 && len(records[0]) == 0 {
			// stop when first cell has empty value
			break
		}

		if len(headers) == 0 {
			for _, record := range records {
				headers = append(headers, slugify(record))
			}
		} else {
			if len(records) != len(headers) {
				break
			}

			// convert as map to marshal/unmarshall in structure
			dataAsMap := map[string]string{}
			for i, record := range records {
				dataAsMap[headers[i]] = record
			}

			buf, _ := json.Marshal(dataAsMap)
			var row T
			if err = json.Unmarshal(buf, &row); err != nil {
				return nil, err
			}
			results = append(results, &row)
		}
	}

	return results, nil
}

func slugify(columnName string) string {
	columnName = strings.ToLower(columnName)
	columnName = strings.ReplaceAll(columnName, "(", "")
	columnName = strings.ReplaceAll(columnName, ")", "")
	columnName = strings.ReplaceAll(columnName, " ", "_")

	return columnName
}
