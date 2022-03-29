package utils

import "fmt"

func PrintImage(image [][]int) {
	for i := 0; i < len(image); i++ {
		for j := 0; j < len(image[i]); j++ {
			fmt.Print(image[i][j])
			fmt.Print(" ")
		}
		fmt.Println()
	}
}

func PrintStringArray(strs []string) {
	for _, s := range strs {
		fmt.Println(s)

	}
}
