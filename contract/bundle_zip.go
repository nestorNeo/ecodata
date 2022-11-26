package contract

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"io/ioutil"
	"log"
)

const (
	INFOFILE = "meta.json"
)

func InspectBundle(zipPath string) (*Record, error) {
	log.Println("opening zip file")
	zf, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}

	defer zf.Close()

	// just two files nothing more
	log.Println("Starting file checks")
	if len(zf.File) < 1 {
		return nil, errors.New("security risk bundle has more than two files")
	}

	record := &Record{}

	index := make(map[string]*zip.File)

	for _, file := range zf.File {
		index[file.Name] = file
		log.Println(file.Name)
	}

	meta, ok := index[INFOFILE]

	if !ok {
		return nil, errors.New("missing meta.json")
	}

	data, err := loadFieldData(meta)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(data, record)
	if err != nil {
		return nil, err
	}

	switch record.Tipo {
	case "CO2_MEASURE":
		log.Println("found co2 type doing processing")
		if isValidError := record.IsValidCO2(); isValidError != nil {
			return record, isValidError
		}
		break
	case "NOISE_MEASURE":
		log.Println("found noise record doingi processing")

		if isValidError := record.IsValidNoise(); isValidError != nil {
			return record, isValidError
		}
		break
	default:
		return record, errors.New("FATAL ERROR NOT SUPPORTED " + record.Tipo)
	}

	return record, nil
}

// individual zip file record
func loadFieldData(file *zip.File) ([]byte, error) {
	fc, err := file.Open()
	if err != nil {
		return nil, err
	}
	defer fc.Close()
	return ioutil.ReadAll(fc)
}
