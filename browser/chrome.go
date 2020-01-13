package browser

type Chrome struct {
	conn string
}

func (c *Chrome) Open() error {
	return nil
}

func (c *Chrome) GetVisitedList() (*VisitedList, error) {

	return nil, nil
}

func (c *Chrome) GetDownloadedList() (*DownloadList, error) {

	return nil, nil
}
