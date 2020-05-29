package brofor

import (
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/MonaxGT/brofor/browser"
)

var excludeExtension = []string{".avi", ".mp3", ".m4v", ".m4b", ".mp4", ".pcap", ".iso", ".ova", ".vmdk", ".jpeg", ".jpg", ".mov", ".gif", ".csv", ".json", ".vsdx"}

func findHashFiles(done <-chan struct{}, paths <-chan string, c chan<- hashResult) {
	for path := range paths {
		fmt.Println(path)
		if _, err := os.Stat(path); err == nil {
			f, err := os.Open(path)
			defer f.Close()
			h := sha256.New()
			_, err = io.Copy(h, f)
			hash := fmt.Sprintf("%x", h.Sum(nil))
			fmt.Println(hash)
			select {
			case c <- hashResult{path, hash, err}:
			case <-done:
				return
			}
		}
	}
	fmt.Println("I've done")
}

func extPaths(done <-chan struct{}, downloaded *[]browser.DownloadList) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)

	go func() {
		defer close(paths)
		errc <- func() error {
			for _, i := range *downloaded {
				ext := filepath.Ext(normalizeFilePath(i.CurrentPath.String))
				if stringInSlice(ext) {
					continue
				}
				select {
				case paths <- normalizeFilePath(i.CurrentPath.String):
				case <-done:
					return errors.New("calc hashes cancelled")
				}

			}
			return nil
		}()

	}()
	return paths, errc
}

func stringInSlice(a string) bool {
	for _, b := range excludeExtension {
		if b == a {
			return true
		}
	}
	return false
}
