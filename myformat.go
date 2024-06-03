package myformat

import (
	"encoding/json"
	"encoding/xml"
	"os"
	"sort"
)

type patient struct {
	Name  string `xml:"name"`
	Age   int    `xml:"age"`
	Email string `xml:"email"`
}

type patients struct {
	List []patient `xml:"Patient"`
}

func fileDecode(fileIn string) (*patients, error) {

	f, err := os.Open(fileIn)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	res := patients{}

	dec := json.NewDecoder(f)
	for dec.More() {
		var a patient
		err = dec.Decode(&a)
		if err != nil {
			return nil, err
		}
		res.List = append(res.List, a)
	}

	sort.Slice(res.List, func(i, j int) bool {
		return res.List[i].Age < res.List[j].Age
	})

	return &res, nil
}

func fileEncode(fileOut string, a *patients) error {

	f, err := os.Create(fileOut)

	if err != nil {
		return err
	}
	defer f.Close()

	_, err = f.WriteString(xml.Header)
	if err != nil {
		return err
	}

	encoder := xml.NewEncoder(f)
	encoder.Indent("", " ")

	err = encoder.Encode(a)
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
