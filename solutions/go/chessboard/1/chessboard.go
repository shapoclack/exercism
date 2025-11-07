package chessboard

type File []bool

type Chessboard map[string]File

func CountInFile(cb Chessboard, file string) int {
	counter := 0
	if f, ok := cb[file]; ok {
		for _, occupied := range f {
			if occupied {
				counter++
			}
		}
	}
	return counter
}

func CountInRank(cb Chessboard, rank int) int {
	if rank < 1 || rank > 8 {
		return 0
	}
	counter := 0
	for _, file := range cb {

		if file[rank-1] {
			counter++
		}
	}
	return counter
}


func CountAll(cb Chessboard) int {
	return len(cb) * 8
}


func CountOccupied(cb Chessboard) int {
	counter := 0
	for _, file := range cb {
		for _, occupied := range file {
			if occupied {
				counter++
			}
		}
	}
	return counter
}
