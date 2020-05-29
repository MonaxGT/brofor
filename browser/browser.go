package browser

import (
	"database/sql"
	"time"
)

type VisitedList struct {
	ID              sql.NullInt64
	URL             sql.NullString
	Title           sql.NullString
	RevHost         sql.NullString
	VisitCount      sql.NullInt64
	TypedCount      sql.NullInt64
	Frecency        sql.NullInt64
	LastVisitedTime sql.NullInt64
	Hidden          sql.NullInt32
	GUID            sql.NullString
	ForeignCount    sql.NullInt64
	URLHash         sql.NullString
	Description     sql.NullString
	PreviewImageURL sql.NullString
	OriginID        sql.NullInt32
	LastTime        time.Time
}

type DownloadList struct {
	ID               sql.NullInt64
	PlaceID          sql.NullInt32
	AnnoAttributeId  sql.NullInt32
	Flags            sql.NullInt32
	Expiration       sql.NullInt32
	Type             sql.NullInt32
	GUID             sql.NullString
	CurrentPath      sql.NullString
	TargetPath       sql.NullString
	StartTime        sql.NullInt64
	ReceivedBytes    sql.NullInt64
	TotalBytes       sql.NullInt64
	State            sql.NullInt32
	DangerType       sql.NullInt32
	InterruptReason  sql.NullInt32
	Hash             sql.NullString
	EndTime          sql.NullInt64
	Opened           sql.NullInt32
	LastAccessTime   sql.NullInt64
	Transient        sql.NullInt32
	Referrer         sql.NullString
	SiteURL          sql.NullString
	TabURL           sql.NullString
	TabReferrerURL   sql.NullString
	HTTPMethod       sql.NullString
	ByEXTId          sql.NullString
	ByEXTName        sql.NullString
	ETag             sql.NullString
	LastModified     sql.NullString
	MimeType         sql.NullString
	OriginalMimeType sql.NullString
	Time             time.Time
}

type Browser interface {
	Open(dbPath string, User string) error
	GetVisitedList() (*[]VisitedList, error)
	GetDownloadedList() (*[]DownloadList, error)
}
