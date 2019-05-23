package main

import (
	"encoding/json"
	"io/ioutil"
)

// DataFileName is the default name of the data file.
var DataFileName = "data.json"

func parseDataFile(file string) (*consignment, error) {
	result := consignment{}

	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}

	json.Unmarshal(data, &result)
	return &result, nil
}
