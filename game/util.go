package game

func mapRuneToIndex(char rune) int {
	for i, k := range keys {
		if k.r == char {
			return i
		}
	}
	return -1 // Return -1 if character is not in slice
}

// func mapIndexToRune(i int) rune {
// 	return keys[i].r
// }
