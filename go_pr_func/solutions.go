package main

import "fmt"

func triangleArea(base float64, height float64) float64 {
	return (base * height / 2)
}
func sortArray(arr []int) []int {
	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				temp := arr[i]
				arr[i] = arr[j]
				arr[j] = temp
			}
		}
	}
	return arr
}
func sumOfSquares(n int) int {
	res := 0
	for i := 2; i <= n; i += 2 {
		res += i * i
	}
	return res
}
func sPalindrome(s string) bool {
	var reverseS string = ""
	for _, ch := range s {
		reverseS = string(ch) + reverseS
	}
	return s == reverseS
}
func isPrime(n int) bool {
	if n < 2 || n%2 == 0 && n != 2 {
		return false
	}
	for i := 2; i <= n/2; i++ {
		if n%i == 0 {
			//println(i, " - число на которое делится")
			return false
		}
	}
	return true
}
func generatePrimes(limit int) []int {
	var res []int
	for i := 2; i <= limit; i++ {
		if isPrime(i) {
			res = append(res, i)
		}
	}
	return res
}

func toBinary(n int64) string {
	if n < 0 {
		return "Not working with negative"
	} else if n == 0 {
		return "0"
	}
	res := ""
	for i := 1; n > 0; i *= 10 {
		if n%2 == 0 {
			res = "0" + res
		} else {
			res = "1" + res
		}
		n /= 2
	}
	return res
}
func findMax(arr []int) int {
	var res int = arr[0]
	for _, i := range arr {
		if res < i {
			res = i
		}
	}
	return res
}
func gcd(a int, b int) int {
	if a == b {
		return a
	}
	if b > a {
		return gcd(a, b-a)
	}
	return gcd(a-b, b)
}
func sumArray(arr []int) int {
	res := 0
	for _, i := range arr {
		res += i
	}
	return res
}

func main() {
	// fmt.Println(triangleArea(4.5, 3.6))
	// fmt.Println(sortArray([]int{5, 2, 8, 3, 0, 9, -3, 568, 34}))
	// fmt.Println(sumOfSquares(8))
	// fmt.Println(sPalindrome("r tt r"))
	// fmt.Println(isPrime(53))
	// fmt.Println(generatePrimes(53))
	// fmt.Println(toBinary(4657))
	// fmt.Println(findMax([]int{5, 3, 8, 6, 1, 0, 23, 1, 4, 35, 6}))
	// fmt.Println(gcd(45, 30))
	// fmt.Println(gcd(45, 30))
	fmt.Println(sumArray([]int{5, 3, 8, 6, 1, 0, 23, 1, 4, 35, 6}))
}
