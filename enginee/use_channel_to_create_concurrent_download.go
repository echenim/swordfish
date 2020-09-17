package enginee

import (
	"fmt"
	"sync"
)

//use channel to download each piece of the file concurrently
func (d Downloader) concurrentDownload(pieces [][2]int, er error) {
	var wg sync.WaitGroup
	for i, p := range pieces {
		wg.Add(1)
		go func(i int, p [2]int) {
			defer wg.Done()
			fmt.Println(i)
			er = d.DownloadFilePieceByPiece(i, p)
			if er != nil {
				panic(er)
			}
		}(i, p)
	}
	wg.Wait()
}
