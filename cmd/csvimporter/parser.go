package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type parser struct {
	file      *os.File
	csvReader *csv.Reader
	header    []string
}

type record struct {
	csvData  []string
	hospital Hospital
}

type Hospital struct {
	Name    string
	City    string
	Address string
	Phone   string
	Email   string
}

func openParser(path string) (p *parser, err error) {
	p = &parser{}

	p.file, err = os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("error opening file %s", path)
	}

	p.csvReader = csv.NewReader(p.file)

	return p, nil
}

func (p *parser) readNextRecord() (rec *record, err error) {
	rec = &record{}

	rec.csvData, err = p.csvReader.Read()
	if err != nil {
		return nil, err
	}

	rec.hospital = p.parseCSVtoStruct(rec)

	return rec, nil
}

func (p *parser) parseCSVtoStruct(rec *record) Hospital {
	hospital := Hospital{}

	for idx, header := range p.header {
		switch header {
		case "Świadczeniodawca - Nazwa":
			hospital.Name = rec.csvData[idx]
		case "Świadczeniodawca - Miejscowość":
			hospital.City = rec.csvData[idx]
		case "Świadczeniodawca - Adres":
			hospital.Address = rec.csvData[idx]
		case "Świadczeniodawca - Telefon":
			hospital.Phone = rec.csvData[idx]
		case "Świadczeniodawca - adres mailowy":
			hospital.Email = rec.csvData[idx]
		}
	}

	return hospital
}
