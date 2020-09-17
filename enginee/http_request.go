package enginee

import (
	"net/http"
)

//use thius method to http request
func (d Downloader) getHTTPRequest(method string) (*http.Request, error) {
	r, er := http.NewRequest(method, d.From, nil)
	check(er, "error occured while retriving the file:")
	r.Header.Set("User-Agent", "download manager")
	return r, nil
}
