package client

import (
	"net/http"
	"os"

	"encoding/xml"
)

// Implements webdav.File.
type file struct {
	c *Client
	name string
}

type statRes struct {
	Href string `xml:"href"`
	CreationDate string `xml:"propstat>prop>creationdate"`
	DisplayName string `xml:"propstat>prop>displayname"`
	GetContentLength int `xml:"propstat>prop>getcontentlength"`
	GetContentType string `xml:"propstat>prop>getcontenttype"`
	GetETag string `xml:"propstat>prop>getetag"`
	GetLastModified string `xml:"propstat>prop>getlastmodified"`
	ResourceType string `xml:"propstat>prop>resourcetype"`
	SupportedLock string `xml:"propstat>prop>supportedlock"`
}

func (f *file) Stat() (info os.FileInfo, err error) {
	req, err := f.c.NewRequest("PROPFIND", f.name)
	if err != nil {
		return
	}

	req.Header.Add("Depth", "0")
	req.Header.Add("Translate", "f")

	res, err := f.c.Do(req)
	if err != nil {
		return
	}
	defer res.Body.Close()

	if err = ensureHasCode(res, http.StatusOK); err != nil {
		return
	}

	xmlRes := &statRes{}
	if err = xml.NewDecoder(res.Body).Decode(&xmlRes); err != nil {
		return
	}

	// TODO: populate info
	return
}

func (f *file) Read(b []byte) (int, error) {
	return 0, nil
}

func (f *file) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func (f *file) Readdir(count int) ([]os.FileInfo, error) {
	return nil, nil
}

func (f *file) Write(b []byte) (int, error) {
	return 0, nil
}

func (f *file) Close() error {
	return nil
}
