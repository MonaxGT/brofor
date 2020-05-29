package output

import (
	"strconv"

	"github.com/360EntSecGroup-Skylar/excelize"
)

type Excel struct {
	f      *excelize.File
	Report string
}

func mapSheet1(m map[string]string, file *excelize.File) error {
	for k, v := range m {
		err := file.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func mapSheet2(m map[string]string, file *excelize.File) error {
	for k, v := range m {
		err := file.SetCellValue("Sheet2", k, v)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Excel) Open() error {
	var err error
	c.f = excelize.NewFile()
	const sheet1 = "Sheet1"
	if err = c.f.SetPageMargins(sheet1,
		excelize.PageMarginBottom(0.5),
		excelize.PageMarginFooter(0.5),
		excelize.PageMarginHeader(0.5),
		excelize.PageMarginLeft(0.5),
		excelize.PageMarginRight(0.5),
		excelize.PageMarginTop(0.5),
	); err != nil {
		return err
	}

	if err := c.f.SetPageLayout(
		"Sheet1",
		excelize.PageLayoutOrientation(excelize.OrientationLandscape),
	); err != nil {
		return err
	}
	if err := c.f.SetPageLayout(
		"Sheet1",
		excelize.PageLayoutPaperSize(10),
	); err != nil {
		return err
	}
	c.f.NewSheet("Sheet2")

	const sheet2 = "Sheet2"
	if err = c.f.SetPageMargins(sheet2,
		excelize.PageMarginBottom(0.5),
		excelize.PageMarginFooter(0.5),
		excelize.PageMarginHeader(0.5),
		excelize.PageMarginLeft(0.5),
		excelize.PageMarginRight(0.5),
		excelize.PageMarginTop(0.5),
	); err != nil {
		return err
	}

	if err := c.f.SetPageLayout(
		"Sheet2",
		excelize.PageLayoutOrientation(excelize.OrientationLandscape),
	); err != nil {
		return err
	}
	if err := c.f.SetPageLayout(
		"Sheet2",
		excelize.PageLayoutPaperSize(10),
	); err != nil {
		return err
	}

	return nil
}

func (c *Excel) SaveVisitedLinks(links *[]VisitedLinks) error {
	valN := make(map[string]int)
	valUrl := make(map[string]string)
	valTitle := make(map[string]string)
	valDomain := make(map[string]string)
	valLastVisitTime := make(map[string]string)
	valVisitCount := make(map[string]int64)
	valReputation := make(map[string]uint64)
	valHidden := make(map[string]bool)

	for i, p := range *links {
		valN["A"+strconv.Itoa(i+4)] = i + 1
		valUrl["B"+strconv.Itoa(i+4)] = p.URL
		valTitle["C"+strconv.Itoa(i+4)] = p.Title
		valDomain["D"+strconv.Itoa(i+4)] = p.Domain
		valLastVisitTime["E"+strconv.Itoa(i+4)] = p.LastVisitTime.String()
		valVisitCount["F"+strconv.Itoa(i+4)] = p.VisitCount
		valReputation["G"+strconv.Itoa(i+4)] = p.Reputation
		valHidden["H"+strconv.Itoa(i+4)] = p.Hidden
	}

	err := c.f.MergeCell("Sheet1", "A1", "H1")
	if err != nil {
		return err
	}
	err = c.f.SetCellValue("Sheet1", "A1", "Report VisitedLinks")
	tableHeader := map[string]string{"A3": "№", "B3": "Username", "C3": "URL",
		"D3": "Domain", "E3": "Title", "F3": "VisitCount", "G3": "LastVisitTime",
		"H3": "Hidden"}

	err = mapSheet1(tableHeader, c.f)
	if err != nil {
		return err
	}
	for k, v := range valN {
		err := c.f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}

	err = mapSheet1(valUrl, c.f)
	if err != nil {
		return err
	}
	err = mapSheet1(valTitle, c.f)
	if err != nil {
		return err
	}
	err = mapSheet1(valDomain, c.f)
	if err != nil {
		return err
	}
	for k, v := range valLastVisitTime {
		err := c.f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}
	for k, v := range valVisitCount {
		err := c.f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}
	for k, v := range valReputation {
		err := c.f.SetCellValue("Sheet1", k, v)
		if err != nil {
			return err
		}
	}
	for k, v := range valHidden {
		var val string
		if v == false {
			val = "hidden"
		} else {
			val = "unhidden"
		}
		err := c.f.SetCellValue("Sheet1", k, val)
		if err != nil {
			return err
		}
	}

	err = c.f.AddTable("Sheet1", "A3", "H"+strconv.Itoa(len(valN)+3), `{"table_name":"table","table_style":"TableStyleLight18",
"show_first_column":true,"show_last_column":true,"show_row_stripes":false,"show_column_stripes":true}`)
	if err != nil {
		return err
	}
	styleTitle, err := c.f.NewStyle(`{"alignment":{"horizontal":"center", "vertical":"center"},
"font":{"bold":true,"family":"Calibri","size":14,"color":"#000000"}}`)
	if err != nil {
		return err
	}
	styleTitleTableCenter, err := c.f.NewStyle(`{"alignment":{"horizontal":"center", "vertical":"center", "wrap_text":true},
"font":{"family":"Calibri","size":11}}`)
	if err != nil {
		return err
	}
	styleTableCenter, err := c.f.NewStyle(`{"alignment":{"horizontal":"center", "vertical":"center"},
	"font":{"family":"Calibri","size":10}}`)
	if err != nil {
		return err
	}
	styleTableLeft, err := c.f.NewStyle(`{"alignment":{"horizontal":"left", "vertical":"center", "wrap_text":true},
"font":{"family":"Calibri","size":10}}`)
	if err != nil {
		return err
	}

	n := strconv.Itoa(len(valN) + 4)
	err = c.f.SetCellStyle("Sheet1", "A1", "H1", styleTitle)
	err = c.f.SetCellStyle("Sheet1", "A3", "H3", styleTitleTableCenter)
	err = c.f.SetCellStyle("Sheet1", "A4", "A"+n, styleTableCenter)
	err = c.f.SetCellStyle("Sheet1", "B4", "D"+n, styleTableLeft)
	err = c.f.SetCellStyle("Sheet1", "E4", "H"+n, styleTableCenter)
	if err != nil {
		return err
	}

	err = c.f.SetColWidth("Sheet1", "A", "A", 5)
	err = c.f.SetColWidth("Sheet1", "B", "B", 55)
	err = c.f.SetColWidth("Sheet1", "C", "C", 60)
	err = c.f.SetColWidth("Sheet1", "D", "D", 30)
	err = c.f.SetColWidth("Sheet1", "E", "E", 29)
	err = c.f.SetColWidth("Sheet1", "F", "F", 11)
	err = c.f.SetColWidth("Sheet1", "G", "G", 13)
	err = c.f.SetColWidth("Sheet1", "H", "H", 10)
	if err != nil {
		return err
	}

	if err := c.f.SaveAs(c.Report); err != nil {
		return err
	}
	return nil
}

func (c *Excel) SaveDownloadedLinks(links *[]DownloadedLinks) error {
	valN := make(map[string]int)
	valCurrentPath := make(map[string]string)
	valTargetPath := make(map[string]string)
	valLastModified := make(map[string]string)
	valReceivedBytes := make(map[string]string)
	valTotalBytes := make(map[string]string)
	valSiteURL := make(map[string]string)
	valReferrer := make(map[string]string)
	valStartTime := make(map[string]string)
	valInterruptReason := make(map[string]int32)
	valMimeType := make(map[string]string)

	for i, p := range *links {
		valN["A"+strconv.Itoa(i+4)] = i + 1
		valCurrentPath["B"+strconv.Itoa(i+4)] = p.CurrentPath
		valTargetPath["C"+strconv.Itoa(i+4)] = p.TargetPath
		valLastModified["D"+strconv.Itoa(i+4)] = p.LastModified
		valReceivedBytes["E"+strconv.Itoa(i+4)] = p.ReceivedBytes
		valTotalBytes["F"+strconv.Itoa(i+4)] = p.TotalBytes
		valSiteURL["G"+strconv.Itoa(i+4)] = p.SiteURL
		valReferrer["H"+strconv.Itoa(i+4)] = p.Referrer
		valStartTime["I"+strconv.Itoa(i+4)] = p.StartTime.String()
		valInterruptReason["J"+strconv.Itoa(i+4)] = p.InterruptReason
		valMimeType["K"+strconv.Itoa(i+4)] = p.MimeType
	}

	err := c.f.MergeCell("Sheet2", "A1", "K1")
	if err != nil {
		return err
	}
	err = c.f.SetCellValue("Sheet2", "A1", "Report DownloadedLinks")

	tableHeader := map[string]string{"A3": "№", "B3": "CurrentPath", "C3": "TargetPath",
		"D3": "LastModified", "E3": "ReceivedBytes", "F3": "TotalBytes", "G3": "SiteURL",
		"H3": "Referrer", "I3": "StartTime", "J3": "InterruptReason", "K3": "MimeType"}

	err = mapSheet2(tableHeader, c.f)
	if err != nil {
		return err
	}
	for k, v := range valN {
		err := c.f.SetCellValue("Sheet2", k, v)
		if err != nil {
			return err
		}
	}

	err = mapSheet2(valCurrentPath, c.f)
	if err != nil {
		return err
	}
	err = mapSheet2(valTargetPath, c.f)
	if err != nil {
		return err
	}
	err = mapSheet2(valLastModified, c.f)
	if err != nil {
		return err
	}
	err = mapSheet2(valReceivedBytes, c.f)
	if err != nil {
		return err
	}
	err = mapSheet2(valTotalBytes, c.f)
	if err != nil {
		return err
	}
	err = mapSheet2(valSiteURL, c.f)
	if err != nil {
		return err
	}
	err = mapSheet2(valReferrer, c.f)
	if err != nil {
		return err
	}
	for k, v := range valStartTime {
		err := c.f.SetCellValue("Sheet2", k, v)
		if err != nil {
			return err
		}
	}
	for k, v := range valInterruptReason {
		err := c.f.SetCellValue("Sheet2", k, v)
		if err != nil {
			return err
		}
	}
	err = mapSheet2(valMimeType, c.f)
	if err != nil {
		return err
	}

	err = c.f.AddTable("Sheet2", "A3", "K"+strconv.Itoa(len(valN)+3), `{"table_name":"table","table_style":"TableStyleLight18",
"show_first_column":true,"show_last_column":true,"show_row_stripes":false,"show_column_stripes":true}`)
	if err != nil {
		return err
	}

	styleTitle, err := c.f.NewStyle(`{"alignment":{"horizontal":"center", "vertical":"center"},
"font":{"bold":true,"family":"Calibri","size":14,"color":"#000000"}}`)
	if err != nil {
		return err
	}
	styleTitleTableCenter, err := c.f.NewStyle(`{"alignment":{"horizontal":"center", "vertical":"center", "wrap_text":true},
"font":{"family":"Calibri","size":11}}`)
	if err != nil {
		return err
	}
	styleTableCenter, err := c.f.NewStyle(`{"alignment":{"horizontal":"center", "vertical":"center"},
	"font":{"family":"Calibri","size":10}}`)
	if err != nil {
		return err
	}
	styleTableLeft, err := c.f.NewStyle(`{"alignment":{"horizontal":"left", "vertical":"center", "wrap_text":true},
"font":{"family":"Calibri","size":10}}`)
	if err != nil {
		return err
	}

	n := strconv.Itoa(len(valN) + 4)
	err = c.f.SetCellStyle("Sheet2", "A1", "H1", styleTitle)
	err = c.f.SetCellStyle("Sheet2", "A3", "K3", styleTitleTableCenter)
	err = c.f.SetCellStyle("Sheet2", "A4", "A"+n, styleTableCenter)
	err = c.f.SetCellStyle("Sheet2", "B4", "C"+n, styleTableLeft)
	err = c.f.SetCellStyle("Sheet2", "D4", "F"+n, styleTableCenter)
	err = c.f.SetCellStyle("Sheet2", "G4", "H"+n, styleTableLeft)
	err = c.f.SetCellStyle("Sheet2", "I4", "J"+n, styleTableCenter)
	err = c.f.SetCellStyle("Sheet2", "K4", "K"+n, styleTableLeft)
	if err != nil {
		return err
	}

	err = c.f.SetColWidth("Sheet2", "A", "A", 5)
	err = c.f.SetColWidth("Sheet2", "B", "B", 55)
	err = c.f.SetColWidth("Sheet2", "C", "C", 20)
	err = c.f.SetColWidth("Sheet2", "D", "D", 28)
	err = c.f.SetColWidth("Sheet2", "E", "E", 13)
	err = c.f.SetColWidth("Sheet2", "F", "F", 11)
	err = c.f.SetColWidth("Sheet2", "G", "H", 40)
	err = c.f.SetColWidth("Sheet2", "I", "I", 27)
	err = c.f.SetColWidth("Sheet2", "J", "J", 15)
	err = c.f.SetColWidth("Sheet2", "K", "K", 27)
	if err != nil {
		return err
	}

	if err := c.f.SaveAs(c.Report); err != nil {
		return err
	}
	return nil
}
