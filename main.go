package main

import "fmt"

func stringNilai(variable string, value int) (result string) {
	result = fmt.Sprintf("Nilai %v = %v", variable, value)
	return
}

func main() {
	const iLimit, jLimit int = 4, 10

	for i := 0; i <= iLimit; i++ {
		fmt.Println(stringNilai("i", i))

		if i == iLimit {
			for j := 0; j <= jLimit; j++ {
				if j == 5 {
					const russiaString = "САШАРВО"
					for index, runeValue := range russiaString {
						fmt.Printf("character %#U starts at byte position %d\n", runeValue, index)
					}
				} else {
					fmt.Println(stringNilai("j", j))
				}
			}
		}
	}
}
