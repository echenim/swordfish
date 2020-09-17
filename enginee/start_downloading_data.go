package enginee

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

//StartDownloading : start the download  with this function
func (d Downloader) StartDownloading() error {
	fmt.Println("\n making connection to ......")
	r, er := d.getHTTPRequest("HEAD")
	Check(er, "fail to open connection : ")

	resp, er := http.DefaultClient.Do(r)
	Check(er, "fail to get header : ")
	log.Printf("\n Response : %v\n ", resp.StatusCode)

	if resp.StatusCode > 299 {
		log.Printf("Can not process, response is %v\n", resp.StatusCode)
	}
	size, er := strconv.Atoi(resp.Header.Get("Content-Length"))
	Check(er, "fail to get content-length")
	log.Printf("Size is %v bytes\n", size)
	pieces := make([][2]int, d.Pieces)
	var eachSize = size / d.Pieces
	d.breakFileInToPieces(pieces, eachSize)

	log.Println(pieces)
	d.concurrentDownload(pieces, er)
	return nil
}
