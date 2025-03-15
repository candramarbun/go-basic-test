package main

import "log"

func main() {
	// main function
	log.Println(FindMatch(6, "a", "b", "c", "d", "e"))
}

func FindMatch(n int, strings ...string) (result bool, idx int) {
	result = false
	if n != len(strings) {
		log.Printf("expected n is %v but array length is %v", n, len(strings))
		return
	}
	idx = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if strings[i] == strings[j] {
				result = true
				idx = i
				break
			}
		}
		if result {
			break
		}
	}
	return
}
