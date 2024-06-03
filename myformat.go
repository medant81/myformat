package myformat

import (
	"encoding/json"
	"os"
)

type animal struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

func fileDecode(fileIn string) ([]animal, error) {

	f, err := os.Open(fileIn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := make([]animal, 0, 4)
	dec := json.NewDecoder(f)
	for dec.More() {
		var a animal
		err = dec.Decode(&a)
		if err != nil {
			return nil, err
		}
		res = append(res, a)
	}

	return res, nil
}

func fileEncode(fileOut string, a []animal) error {

	f, err := os.Create(fileOut)
	if err != nil {
		return err
	}

	err = json.NewEncoder(f).Encode(a)
	if err != nil {
		return err
	}
	err = f.Close()
	if err != nil {
		return err
	}

	return nil
}

func Do(fileIn, fileOut string) error {

	res, err := fileDecode(fileIn)
	if err != nil {
		return err
	}

	err = fileEncode(fileOut, res)
	if err != nil {
		return err
	}

	return nil
}
