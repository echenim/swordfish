package enginee

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Merge tmp files to single file and delete tmp files
func (d Downloader) mergeAll(sections [][2]int) error {
	f, err := os.OpenFile(d.To, os.O_CREATE|os.O_WRONLY|os.O_APPEND, os.ModePerm)
	if err != nil {
		return err
	}
	defer f.Close()
	for i := range sections {
		tmpFileName := fmt.Sprintf("tmp/piece-%v.tmp", i)
		b, err := ioutil.ReadFile(tmpFileName)
		if err != nil {
			return err
		}
		n, err := f.Write(b)
		if err != nil {
			return err
		}
		err = os.Remove(tmpFileName)
		if err != nil {
			return err
		}
		fmt.Printf("%v bytes merged\n", n)
	}
	return nil
}
