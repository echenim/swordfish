package main

import (
	"fmt"

	dw "github.com/echenim/swordfish/enginee"
)

func main() {
	d := dw.Downloader{
		From:   "http://192.168.1.73/armwealth/css/Mcknight.mp4",
		To:     "download/home.mp4",
		Pieces: 100,
	}

	er := d.StartDownloading()
	dw.Check(er, "some went wrong while downloading the file:  %s\n")
	fmt.Println("\n download completed")

}
