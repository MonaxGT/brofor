package output

import (
	"encoding/csv"
	"os"
	"strconv"
)

type CSV struct {
	connS       *os.File
	connD       *os.File
	writerSites *csv.Writer
	writerFiles *csv.Writer
	FileSites   string
	FileFiles   string
}

func (c *CSV) Open() error {
	var err error
	c.connS, err = os.Create(c.FileSites)
	if err != nil {
		return err
	}
	c.connD, err = os.Create(c.FileFiles)
	if err != nil {
		return err
	}
	c.writerSites = csv.NewWriter(c.connS)
	c.writerFiles = csv.NewWriter(c.connD)
	return nil
}

func (c *CSV) SaveVisitedLinks(links *[]VisitedLinks) error {
	for _, i := range *links {
		err := c.writerSites.Write([]string{i.URL, i.Title, i.Domain, i.LastVisitTime.String(), string(i.VisitCount), string(i.Reputation), strconv.FormatBool(i.Hidden)})
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *CSV) SaveDownloadedLinks(links *[]DownloadedLinks) error {
	for _, i := range *links {
		err := c.writerFiles.Write([]string{i.CurrentPath, i.TargetPath, i.LastModified, i.ReceivedBytes, i.TotalBytes, i.SiteURL, i.Referrer, i.StartTime.String(), string(i.InterruptReason), i.MimeType})
		if err != nil {
			return err
		}
	}
	return nil
}
