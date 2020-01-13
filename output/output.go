package output

type VisitedLinks struct {
}

type DownloadedLinks struct {
}

type File interface {
	Open() error
	SaveVisitedLinks(links *VisitedLinks) error
	SaveDownloadedLinks(links *DownloadedLinks) error
}
