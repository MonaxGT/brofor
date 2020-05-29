package browser

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Chrome struct {
	conn *sql.DB
}

func (c *Chrome) Open(dbPath string, User string) error {
	var err error
	c.conn, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	return nil
}

func (c *Chrome) GetVisitedList() (*[]VisitedList, error) {
	rows, err := c.conn.Query("SELECT id,url,title,visit_count,typed_count,last_visit_time,hidden FROM urls")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	domains := []VisitedList{}

	for rows.Next() {
		p := VisitedList{}
		err := rows.Scan(&p.ID, &p.URL, &p.Title, &p.VisitCount, &p.TypedCount, &p.LastVisitedTime, &p.Hidden)
		if err != nil {
			return nil, err
		}
		p.LastTime = time.Unix(p.LastVisitedTime.Int64/1000000-11644473600, 0)
		domains = append(domains, p)
	}

	return &domains, nil
}

func (c *Chrome) GetDownloadedList() (*[]DownloadList, error) {
	rows, err := c.conn.Query("SELECT * FROM downloads")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dowloads := []DownloadList{}

	for rows.Next() {
		p := DownloadList{}
		err := rows.Scan(&p.ID, &p.GUID, &p.CurrentPath, &p.TargetPath, &p.StartTime, &p.ReceivedBytes, &p.TotalBytes, &p.State, &p.DangerType, &p.InterruptReason, &p.Hash, &p.EndTime, &p.Opened, &p.LastAccessTime, &p.Transient, &p.Referrer, &p.SiteURL, &p.TabURL, &p.TabReferrerURL, &p.HTTPMethod, &p.ByEXTId, &p.ByEXTName, &p.ETag, &p.LastModified, &p.MimeType, &p.OriginalMimeType)
		if err != nil {
			return nil, err
		}
		p.Time = time.Unix(p.StartTime.Int64/1000000-11644473600, 0)
		dowloads = append(dowloads, p)
	}

	return &dowloads, nil
}
