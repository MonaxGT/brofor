package output

import (
	"encoding/json"
	"os"
)

type JSON struct {
	FileS     string
	FileD     string
	fileFiles *os.File
	fileSites *os.File
}

func (j *JSON) Open() error {
	var err error
	j.fileSites, err = os.OpenFile(j.FileS, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	j.fileFiles, err = os.OpenFile(j.FileD, os.O_CREATE, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

func (j *JSON) SaveVisitedLinks(links *[]VisitedLinks) error {
	encoder := json.NewEncoder(j.fileSites)
	err := encoder.Encode(links)
	if err != nil {
		return err
	}
	return nil
}

func (j *JSON) SaveDownloadedLinks(links *[]DownloadedLinks) error {
	encoder := json.NewEncoder(j.fileFiles)
	err := encoder.Encode(links)
	if err != nil {
		return err
	}
	return nil
}
