package go_challenges

// efficient lookup using map
func CountVowels(word string) int {
	vowelMap := map[rune]struct{}{
		'a': {}, 'e': {}, 'i': {}, 'o': {}, 'u': {},
	}
	count := 0
	for _, val:= range word {
		if _, ok:= vowelMap[val]; ok{
			count++
		}
	}
	return count
}

func CountVowels2(word string) int {
	vowelMap := [5]string{"a", "e", "i", "o", "u"}
	count := 0
	for _, val:= range word {
		for _, ch:= range vowelMap{
			if ch == string(val){
				count++
			}
		}
	}
	return count
}
