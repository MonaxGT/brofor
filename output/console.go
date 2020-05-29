package output

import "fmt"

type Console struct {
}

func (c *Console) Open() error {
	return nil
}

func (c *Console) SaveVisitedLinks(links *[]VisitedLinks) error {
	fmt.Println("Visited links:")
	fmt.Printf("%+v\n", *links)
	return nil
}

func (c *Console) SaveDownloadedLinks(links *[]DownloadedLinks) error {
	fmt.Println("Downloaded links:")
	fmt.Printf("%+v\n", *links)
	return nil
}
