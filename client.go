// Package client implements a WebDAV client, as defined in RFC 4918.
package client

import (
	"fmt"
	"net/http"

	"golang.org/x/net/webdav"
)

func ensureHasCode(res *http.Response, expected ...int) error {
	for _, code := range expected {
		if res.StatusCode == code {
			return nil
		}
	}
	return fmt.Errorf("%d %s", res.StatusCode, res.Status)
}

// A WebDAV client.
type Client struct {
	http *http.Client
	root string
}

// Create a new HTTP request.
//
// This function should only be used by packages implementing extensions of
// WebDAV.
func (c *Client) NewRequest(method, name string) (req *http.Request, err error) {
	req, err = http.NewRequest(method, c.root + name, nil)
	// TODO: auth
	return
}

// Perform a HTTP request.
//
// This function should only be used by packages implementing extensions of
// WebDAV.
func (c *Client) Do(req *http.Request) (res *http.Response, err error) {
	return c.http.Do(req)
}

func (c *Client) Mkdir(name string) (err error) {
	req, err := c.NewRequest("MKCOL", name)
	if err != nil {
		return
	}

	res, err := c.Do(req)
	if err != nil {
		return
	}
	res.Body.Close()

	err = ensureHasCode(res, http.StatusCreated)
	return
}

func (c *Client) OpenFile(name string) (f webdav.File, err error) {
	f = &file{
		c: c,
		name: name,
	}
	return
}

// Create a new WebDAV client.
// root is the WebDAV server URL.
func New(root string) *Client {
	return &Client{
		root: root,
		http: &http.Client{},
	}
}
