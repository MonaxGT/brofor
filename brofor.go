package brofor

import (
	"errors"

	"github.com/MonaxGT/brofor/browser"
	"github.com/MonaxGT/brofor/output"
)

// Config main struct
type Config struct {
	br  browser.Browser
	out output.File
}

func (c *Config) forensicProc() error {
	visited, err := c.br.GetVisitedList()
	if err != nil {
		return err
	}
	err = c.out.SaveVisitedLinks(&output.VisitedLinks{})

	downloaded, err := c.br.GetDownloadedList()
	if err != nil {
		return err
	}
	err = c.out.SaveDownloadedLinks(&output.DownloadedLinks{})

	return nil
}

func (c *Config) threatHuntProc() error {

	return nil
}

// New initiate interface and open files
func New(dbPath string, broType string) (*Config, error) {
	var br browser.Browser
	switch broType {
	case "chrome":
		c := browser.Chrome{}
		err := c.Open()
		if err != nil {
			return nil, err
		}
	}

	var out output.File
	switch outType {
	case "csv":
		c := output.CSV{}
		err := c.Open()
		if err != nil {
			return nil, err
		}
	}
	return &Config{
		br: br,
	}, nil
}

// Run function where mode chosen
func (c *Config) Run(mode string) error {
	if mode == "df" {
		err := c.forensicProc()
		if err != nil {
			return err
		}
	} else if mode == "th" {
		err := c.threatHuntProc()
		if err != nil {
			return err
		}
	} else {
		return errors.New("unknown mode type")
	}
	return nil
}
