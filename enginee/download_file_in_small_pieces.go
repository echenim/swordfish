package enginee

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

//DownloadFilePieceByPiece : Download a single piece and save  to a tmp file
func (d Downloader) DownloadFilePieceByPiece(i int, c [2]int) error {
	r, err := d.getHTTPRequest("GET")
	if err != nil {
		return err
	}
	r.Header.Set("Range", fmt.Sprintf("bytes=%v-%v", c[0], c[1]))
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		return err
	}
	if resp.StatusCode > 299 {
		return fmt.Errorf(fmt.Sprintf("Can't process, response is %v", resp.StatusCode))
	}
	fmt.Printf("Downloaded %v bytes for section %v\n", resp.Header.Get("Content-Length"), i)
	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(fmt.Sprintf("section-%v.tmp", i), b, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}
