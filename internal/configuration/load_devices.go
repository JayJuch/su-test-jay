package configuration

import (
	"encoding/csv"
	"os"
)

func LoadDevices(path string) (map[string]struct{}, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	rows, err := csv.NewReader(f).ReadAll()
	if err != nil {
		return nil, err
	}

	devices := make(map[string]struct{})
	for i, row := range rows {
		if i == 0 {
			continue // skip header
		}
		if len(row) > 0 && row[0] != "" {
			devices[row[0]] = struct{}{}
		}
	}
	return devices, nil
}
