package main

import (
	"fmt"
	"strings"
	"unicode"
)

// SOAL 1: Membalikkan String
func reverseWords(input string) string {
	words := strings.Fields(input)
	for i, word := range words {
		words[i] = reverseString(word)
	}
	return strings.Join(words, " ")
}

func reverseString(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

// SOAL 2: FizzBuzz
func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Printf("%d FizzBuzz\n", i)
		} else if i%3 == 0 {
			fmt.Printf("%d Fizz\n", i)
		} else if i%5 == 0 {
			fmt.Printf("%d Buzz\n", i)
		} else {
			fmt.Println(i)
		}
	}
}

// SOAL 3: Deret Fibonacci
func fibonacci(n int) []int {
	if n <= 0 {
		return []int{}
	}
	result := []int{0, 1}
	for i := 2; i < n; i++ {
		next := result[i-1] + result[i-2]
		result = append(result, next)
	}
	return result
}

// SOAL 4: Keuntungan Terbaik
func bestProfit(prices []int) int {
	if len(prices) < 2 {
		return 0
	}
	minPrice := prices[0]
	maxProfit := 0
	for _, price := range prices {
		if price < minPrice {
			minPrice = price
		}
		profit := price - minPrice
		if profit > maxProfit {
			maxProfit = profit
		}
	}
	return maxProfit
}

// SOAL 5: Menghitung Angka dalam Array
func countNumbers(input []string) int {
	count := 0
	for _, char := range input {
		if len(char) == 1 && unicode.IsDigit(rune(char[0])) {
			count++
		}
	}
	return count
}

// MAIN FUNCTION
func main() {
	// SOAL 1
	fmt.Println("SOAL 1:")
	input1 := `italem irad irigayaj iadab itsap ulalreb nalub kusutret gnalali`
	result1 := reverseWords(input1)
	fmt.Println(result1)

	// SOAL 2
	fmt.Println("\nSOAL 2:")
	fizzBuzz()

	// SOAL 3
	fmt.Println("\nSOAL 3:")
	n := 10
	result3 := fibonacci(n)
	fmt.Println(result3)

	// SOAL 4
	fmt.Println("\nSOAL 4:")
	input4_1 := []int{7, 8, 3, 10, 8}
	input4_2 := []int{5, 12, 11, 12, 10}
	input4_3 := []int{7, 18, 27, 10, 29}
	input4_4 := []int{20, 17, 15, 14, 10}
	fmt.Println(bestProfit(input4_1))
	fmt.Println(bestProfit(input4_2))
	fmt.Println(bestProfit(input4_3))
	fmt.Println(bestProfit(input4_4))

	// SOAL 5
	fmt.Println("\nSOAL 5:")
	input5_1 := []string{"b", "7", "h", "6", "h", "k", "i", "5", "g", "7", "8"}
	input5_2 := []string{"7", "b", "8", "5", "6", "9", "n", "f", "y", "6", "9"}
	input5_3 := []string{"u", "h", "b", "n", "7", "6", "5", "1", "g", "7", "9"}
	fmt.Println(countNumbers(input5_1))
	fmt.Println(countNumbers(input5_2))
	fmt.Println(countNumbers(input5_3))
}
