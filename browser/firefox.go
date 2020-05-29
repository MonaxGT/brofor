package browser

import (
	"database/sql"
	"time"

	_ "github.com/mattn/go-sqlite3"
)

type Firefox struct {
	conn *sql.DB
}

func (f *Firefox) Open(dbPath string, User string) error {
	var err error
	f.conn, err = sql.Open("sqlite3", dbPath)
	if err != nil {
		return err
	}
	return nil
}

func (f *Firefox) GetVisitedList() (*[]VisitedList, error) {
	rows, err := f.conn.Query("SELECT * FROM moz_places")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	domains := []VisitedList{}

	for rows.Next() {
		p := VisitedList{}
		err := rows.Scan(&p.ID, &p.URL, &p.Title, &p.RevHost, &p.VisitCount, &p.Hidden, &p.TypedCount, &p.Frecency, &p.LastVisitedTime, &p.GUID, &p.ForeignCount, &p.URLHash, &p.Description, &p.PreviewImageURL, &p.OriginID)
		if err != nil {
			return nil, err
		}
		p.LastTime = time.Unix(p.LastVisitedTime.Int64/1000000, 0)
		domains = append(domains, p)
	}

	return &domains, nil
}

func (f *Firefox) GetDownloadedList() (*[]DownloadList, error) {
	rows, err := f.conn.Query("SELECT * FROM moz_annos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	dowloads := []DownloadList{}

	for rows.Next() {
		p := DownloadList{}
		err := rows.Scan(&p.ID, &p.PlaceID, &p.AnnoAttributeId, &p.CurrentPath, &p.Flags, &p.Expiration, &p.Type, &p.StartTime, &p.LastAccessTime)
		if err != nil {
			return nil, err
		}
		p.Time = time.Unix(p.StartTime.Int64/1000000, 0)
		dowloads = append(dowloads, p)
	}

	return &dowloads, nil
}
