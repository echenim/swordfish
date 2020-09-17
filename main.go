package main

import (
	"fmt"

	dw "github.com/echenim/swordfish/enginee"
)

func main() {
	d := dw.Downloader{
		From:   "",
		To:     "",
		Pieces: 20,
	}

	er := d.StartDownloading()
	dw.Check(er, "some went wrong while downloading the file:  %s\n")
	fmt.Println("\n download completed")

}
