package main

import (
	"fmt"
)

func main() {
	fmt.Println(weighted_strings("abbcccd", []int{1, 3, 9, 8}))
}

func weighted_strings(str string, weights []int) []string {
	if len(weights) == 0 {
		return []string{}
	}

	weights_map := make(map[int]bool)
	for i := 0; i < len(str); i++ {
		weight := get_sequence(str[i])
		weights_map[weight] = true

		for j := i + 1; j < len(str); j++ {
			if str[i] != str[j] {
				break
			}

			weight += get_sequence(str[j])
			weights_map[weight] = true
			i = j
		}
	}

	result := make([]string, 0)
	for _, weight := range weights {
		if weights_map[weight] {
			result = append(result, "Yes")
		} else {
			result = append(result, "No")
		}
	}

	return result
}

func get_sequence(x byte) int {
	return int(x - 'a' + 1)
}
