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

func InspectBundle(zipPath string) (*AudioRecord, error) {
	log.Println("opening zip file")
	zf, err := zip.OpenReader(zipPath)
	if err != nil {
		return nil, err
	}

	defer zf.Close()

	// just two files nothing more
	log.Println("expecting two files in zip")
	if len(zf.File) < 1 {
		return nil, errors.New("security risk bundle has more than two files")
	}

	record := &AudioRecord{}

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

	_, validType := VALIDTYPES[record.Tipo]

	if !validType {
		return record, errors.New("FATAL ERROR NOT SUPPORTED " + record.Tipo)
	}

	if record.Tipo == "CO2" {
		// TODO validated required fields for CO2
		if record.CO2 <= 0 {
			return record, errors.New("invalid CO2 value , please provide a positive number")
		}
		return record, nil
	}

	rawData, ok := index[record.Name]

	if !ok {
		return nil, errors.New("filename pointed by meta is no found")
	}
	data, err = loadFieldData(rawData)

	if err != nil {
		return nil, err
	}

	record.Content = data
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
