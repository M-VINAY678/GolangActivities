package main

import "fmt"

func main() {
	m := make(map[int][]string)
	m[1] = []string{"A", "E", "I", "O", "U", "L", "N", "R", "S", "T"}
	m[2] = []string{"D", "G"}
	m[3] = []string{"B", "C", "M", "P"}
	m[4] = []string{"F", "H", "V", "W", "Y"}
	m[5] = []string{"K"}
	m[8] = []string{"J", "X"}
	m[10] = []string{"Q", "Z"}
	letterMap := make(map[string]int)
	for key, value := range m {
		fmt.Println(key, value)
		for _, i := range value {
			letterMap[i] = key
		}
		// for i := 0; i < len(value); i++ {
		// 	letterMap[value[i]] = key
		// }
	}
	fmt.Print(letterMap)
}
