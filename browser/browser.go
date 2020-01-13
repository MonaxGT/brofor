package browser

type VisitedList struct {
}

type DownloadList struct {
}

type Browser interface {
	Open() error
	GetVisitedList() (*VisitedList, error)
	GetDownloadedList() (*DownloadList, error)
}
