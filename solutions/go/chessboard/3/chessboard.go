package chessboard

// Declare a type named File which stores if a square is occupied by a piece - this will be a slice of bools
type File [8]bool

// Declare a type named Chessboard which contains a map of eight Files, accessed with keys from "A" to "H"
type Chessboard map[string]File

// CountInFile returns how many squares are occupied in the chessboard,
// within the given file.
func CountInFile(cb Chessboard, file string) int {
	count := 0
	if len(file) != 1 || file[0] < 'A' || file[0] > 'H' {
		return 0
	}
	for _, square := range cb[file] {
		if square {
			count++
		}
	}
	return count
}

// CountInRank returns how many squares are occupied in the chessboard,
// within the given rank.
func CountInRank(cb Chessboard, rank int) int {
	count := 0
	if rank < 1 || rank > 8 {
		return 0
	}
	for _, file := range cb {
		if file[rank-1] {
			count++
		}
	}
	return count
}

// CountAll should count how many squares are present in the chessboard.
func CountAll(cb Chessboard) int {
	count := 0
	for _, v := range cb {
		count += len(v)
	}
	return count
}

// CountOccupied returns how many squares are occupied in the chessboard.
func CountOccupied(cb Chessboard) int {
	count := 0
	for k := range cb {
		count += CountInFile(cb, k)
	}
	return count
}
