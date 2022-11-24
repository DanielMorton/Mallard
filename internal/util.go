package internal

func reverseCompliment(text string) string {
	T := len(text)
	complement := ""
	for i := T - 1; i >= 0; i-- {
		if text[i] == 'A' {
			complement += "T"
		} else if text[i] == 'C' {
			complement += "G"
		} else if text[i] == 'G' {
			complement += "C"
		} else {
			complement += "A"
		}
	}
	return complement
}
