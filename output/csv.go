package output

type CSV struct {
	conn string
}

func (c *CSV) Open() error {

	return nil
}

func (c *CSV) SaveVisitedLinks(links *VisitedLinks) error {
	return nil
}

func (c *CSV) SaveDownloadedLinks(links *DownloadedLinks) error {
	return nil
}
