package output

import (
	"encoding/json"
	"errors"
	"fmt"
	"net"
)

type Remote struct {
	Socket string
	conn   net.Conn
	enc    *json.Encoder
}

func (r *Remote) Open() error {
	var err error
	r.conn, err = net.Dial("tcp", r.Socket)
	if err != nil {
		return errors.New(fmt.Sprintf("can't connect to remote tcp log server: %s", r.Socket))
	}
	r.enc = json.NewEncoder(r.conn)
	return nil
}

func (r *Remote) SaveVisitedLinks(links *[]VisitedLinks) error {
	for _, i := range *links {
		err := r.enc.Encode(i)
		if err != nil {
			return err
		}
	}
	return nil
}

func (r *Remote) SaveDownloadedLinks(links *[]DownloadedLinks) error {
	for _, i := range *links {
		err := r.enc.Encode(i)
		if err != nil {
			return err
		}
	}
	return nil
}
