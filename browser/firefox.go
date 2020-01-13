package browser

type Firefox struct {
	conn string
}

func (f *Firefox) Open() error {
	return nil
}

func (f *Firefox) GetVisitedList() (*VisitedList, error) {

	return nil, nil
}

func (f *Firefox) GetDownloadedList() (*DownloadList, error) {

	return nil, nil
}
