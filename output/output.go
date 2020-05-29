package output

import (
	"time"
)

type VisitedLinks struct {
	Username      string    `json:"user,omitempty"`
	URL           string    `json:"url"`
	Domain        string    `json:"domain"`
	Title         string    `json:"title"`
	VisitCount    int64     `json:"visitCount"`
	LastVisitTime time.Time `json:"lastVisitTime"`
	Hidden        bool      `json:"hidden"`
	Reputation    uint64    `json:"reputation"`
}

type DownloadedLinks struct {
	Username        string    `json:"user,omitempty"`
	GUID            string    `json:"giud"`
	CurrentPath     string    `json:"currentPath"`
	TargetPath      string    `json:"targetPath"`
	StartTime       time.Time `json:"startTime"`
	ReceivedBytes   string    `json:"receivedBytes"`
	TotalBytes      string    `json:"totalBytes"`
	State           uint8     `json:"state"`
	DangerType      int32     `json:"dangerType"`
	InterruptReason int32     `json:"interruptReason"`
	Hash            string    `json:"hash"`
	Opened          bool      `json:"opened"`
	Referrer        string    `json:"referrer"`
	SiteURL         string    `json:"siteURL"`
	TabURL          string    `json:"tabURL"`
	TabReferrer     string    `json:"tabReferrer"`
	LastModified    string    `json:"lastModified"`
	MimeType        string    `json:"mimeType"`
}

type File interface {
	Open() error
	SaveVisitedLinks(links *[]VisitedLinks) error
	SaveDownloadedLinks(links *[]DownloadedLinks) error
}
