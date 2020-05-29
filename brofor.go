package brofor

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/url"
	"os"
	"regexp"
	"runtime"
	"strings"
	"sync"

	"github.com/MonaxGT/brofor/browser"
	"github.com/MonaxGT/brofor/output"
)

// Config main struct
type Config struct {
	hash bool
	br   browser.Browser
	out  output.File
}

type infoDB struct {
	pathDB string
	user   string
}

type hashResult struct {
	path   string
	sha256 string
	err    error
}

func (c *Config) calcHash(downloaded *[]browser.DownloadList) (map[string]string, error) {

	numcpu := runtime.NumCPU()
	fmt.Println("NumCPU", numcpu)
	runtime.GOMAXPROCS(numcpu)

	done := make(chan struct{})
	defer close(done)
	paths, errc := extPaths(done, downloaded)

	hashChan := make(chan hashResult, 1000)
	var wg sync.WaitGroup
	const numDigesters = 20
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			findHashFiles(done, paths, hashChan)
			wg.Done()
		}()
	}
	go func() {
		wg.Wait()
		close(hashChan)
	}()
	m := make(map[string]string)
	for r := range hashChan {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sha256
	}
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil
}

func normalizeFilePath(path string) string {
	if strings.HasPrefix(path, "file:///") {
		path = strings.TrimPrefix(path, "file:///")
	}
	pathNew, err := url.QueryUnescape(path)
	if err != nil {
		log.Println(err) // should think how remove this
	}
	return pathNew
}

func ByteCountBinary(b int64) string {
	const unit = 1024
	if b < unit {
		return fmt.Sprintf("%d B", b)
	}
	div, exp := int64(unit), 0
	for n := b / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f %ciB", float64(b)/float64(div), "KMGTPE"[exp])
}

func intToBool(i int32) bool {
	if i == 1 {
		return true
	}
	return false
}

func extractDBPaths(browser string) ([]infoDB, error) {
	files, err := ioutil.ReadDir("C:\\Users")
	if err != nil {
		return nil, err
	}
	var paths []infoDB
	for _, f := range files {
		if f.IsDir() {
			switch browser {
			case "chrome":
				dbPath := fmt.Sprintf("C:\\Users\\%s\\AppData\\Local\\Google\\Chrome\\User Data\\Default\\History", f.Name())
				if _, err := os.Stat(dbPath); err != nil {
					continue
				}
				paths = append(paths, infoDB{
					pathDB: dbPath,
					user:   f.Name(),
				})
			case "opera":
				dbPath := fmt.Sprintf("C:\\Users\\%s\\AppData\\Roaming\\Opera Software\\Opera Stable\\History", f.Name())
				if _, err := os.Stat(dbPath); err != nil {
					continue
				}
				paths = append(paths, infoDB{
					pathDB: dbPath,
					user:   f.Name(),
				})
			case "firefox":
				t := fmt.Sprintf("C:\\Users\\%s\\AppData\\Roaming\\Mozilla\\Firefox\\Profiles", f.Name())
				if _, err := os.Stat(t); err != nil {
					continue
				}
				dir, err := ioutil.ReadDir(t)
				if err != nil {
					return nil, err
				}
				for _, e := range dir {
					if e.IsDir() {
						dbPath := fmt.Sprintf("%s\\%s\\places.sqlite", t, e.Name())
						if _, err := os.Stat(dbPath); err != nil {
							continue
						}
						paths = append(paths, infoDB{
							pathDB: dbPath,
							user:   f.Name(),
						})
					}
				}

			}
		}
	}
	return paths, nil
}

func (c *Config) forensicProc(dbPath string, Username string) error {
	err := c.br.Open(dbPath, Username)
	if err != nil {
		return err
	}

	visited, err := c.br.GetVisitedList()
	if err != nil {
		return err
	}

	v := []output.VisitedLinks{}
	for _, i := range *visited {
		u, err := url.Parse(i.URL.String)
		if err != nil {
			return err
		}
		t := output.VisitedLinks{
			Username:      Username,
			URL:           i.URL.String,
			Domain:        u.Hostname(),
			Title:         i.Title.String,
			VisitCount:    i.VisitCount.Int64,
			LastVisitTime: i.LastTime,
			Hidden:        intToBool(i.Hidden.Int32), //test
		}
		v = append(v, t)

	}
	err = c.out.SaveVisitedLinks(&v)
	if err != nil {
		return err
	}

	downloaded, err := c.br.GetDownloadedList()
	if err != nil {
		return err
	}

	h := make(map[string]string)
	if c.hash {
		fmt.Println("Started calculated hashes")
		h, err = c.calcHash(downloaded)
		if err != nil {
			return err
		}

	}

	d := []output.DownloadedLinks{}
	for _, i := range *downloaded {

		t := output.DownloadedLinks{
			Username:        Username,
			CurrentPath:     normalizeFilePath(i.CurrentPath.String),
			StartTime:       i.Time,
			ReceivedBytes:   ByteCountBinary(i.ReceivedBytes.Int64),
			TotalBytes:      ByteCountBinary(i.TotalBytes.Int64),
			DangerType:      i.DangerType.Int32,
			InterruptReason: i.InterruptReason.Int32,
			Opened:          intToBool(i.Opened.Int32),
			Referrer:        i.Referrer.String,
			SiteURL:         i.SiteURL.String,
			LastModified:    i.LastModified.String,
			MimeType:        i.MimeType.String,
			Hash:            h[normalizeFilePath(i.CurrentPath.String)],
		}
		d = append(d, t)
	}
	err = c.out.SaveDownloadedLinks(&d)
	if err != nil {
		return err
	}
	return nil
}

// Open initiate interface and open files
func New(broType string, outType string, address string, hash bool) (*Config, error) {
	var br browser.Browser
	if broType == "" {
		return nil, errors.New("browser type is not choosen")
	}
	switch broType {
	case "chrome":
		c := &browser.Chrome{}
		br = c
	case "firefox":
		c := &browser.Firefox{}
		br = c
	case "opera":
		c := &browser.Chrome{}
		br = c
	}

	var out output.File
	switch outType {
	case "csv":
		c := &output.CSV{
			FileSites: "resultSites.csv",
			FileFiles: "resultFiles.csv",
		}
		err := c.Open()
		if err != nil {
			return nil, err
		}
		out = c
	case "xls":
		c := &output.Excel{
			Report: "result.xlsx",
		}
		err := c.Open()
		if err != nil {
			return nil, err
		}
		out = c
	case "json":
		c := &output.JSON{
			FileS: "resultSites.json",
			FileD: "resultFiles.json",
		}
		err := c.Open()
		if err != nil {
			return nil, err
		}
		out = c
	case "remote":
		re := regexp.MustCompile(`\b(?:\d{1,3}\.){3}\d{1,3}\b\:\d{1,5}`)
		if !re.Match([]byte(address)) {
			return nil, errors.New("wrong format of socket. IP:Port")
		}
		c := &output.Remote{
			Socket: address,
		}
		err := c.Open()
		if err != nil {
			return nil, err
		}
		out = c
	default:
		c := &output.Console{}
		err := c.Open()
		if err != nil {
			return nil, err
		}
		out = c

	}
	return &Config{
		br:   br,
		out:  out,
		hash: hash,
	}, nil
}

// Run function where mode chosen
func (c *Config) Run(mode string, live bool, dbPath string, browser string) error {
	if live {
		paths, err := extractDBPaths(browser)
		for _, i := range paths {
			err := c.forensicProc(i.pathDB, i.user)
			if err != nil {
				return err
			}
		}
		if err != nil {
			return err
		}
	} else {
		if mode == "df" {
			err := c.forensicProc(dbPath, "")
			if err != nil {
				return err
			}
		} else {
			return errors.New("unknown mode type")
		}

	}
	return nil
}
