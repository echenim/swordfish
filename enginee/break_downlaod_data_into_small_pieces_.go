package enginee

//for file size is 100 bytes, our pieces should like:
// [[0 20] [21 41] [42 62] [63 83] .....]
func (d Downloader) breakFileInToPieces(pieces [][2]int, eachSize int) {

	for i := range pieces {
		if i == 0 {
			// starting byte of first piece
			pieces[i][0] = 0
		} else {
			// starting byte of other pieces
			pieces[i][0] = pieces[i-1][1] + 1
		}

		if i < d.Pieces-1 {
			// ending byte of other pieces
			pieces[i][1] = pieces[i][0] + eachSize
		} else {
			// ending byte of other pieces
			pieces[i][1] = eachSize - 1
		}
	}
}
