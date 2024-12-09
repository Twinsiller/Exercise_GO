package main

import (
	"fmt"
	"math"

	//"regexp"
	"strconv"
	"strings"
)

func sum(n int) int {
	var res int = 0
	for n > 0 {
		res += n % 10
		n /= 10
	}
	return res
}
func temp(s string) string {
	strs := strings.Split(s, " ")
	num, err := strconv.Atoi(strs[0])
	if err != nil {
		return "Error!!!"
	}
	if strs[1] == "(Celsius)" {
		return strconv.Itoa(((num)*9/5 + 32)) + " (Fahrenheit)"
	}
	return strconv.Itoa((num-32)*5/9) + " (Celsius)"
}
func doubleMas(n []int) []int {
	for i, _ := range n {
		n[i] *= 2
	}
	return n
}
func strConcat(s []string) string {
	return strings.Join(s, " ")
}

func linePos(s string) float64 {
	var coords []int
	for i, cH := range s {
		if cH == '=' {
			var l string = ""
			for ; s[i] != ','; i++ {
				l = l + s[i]
			}
			coords = append(coords)
		}
	}

	// Регулярное выражение для поиска чисел
	re := regexp.MustCompile(`\d+`)

	// Найти все вхождения чисел
	matches := re.FindAllString(s, -1)

	// Преобразуем найденные строки в числа

	for _, match := range matches {
		num, err := strconv.Atoi(match)
		if err == nil {
			coords = append(coords, num)
		}
	}
	a := float64(coords[1] - coords[0])
	b := float64(coords[3] - coords[2])
	return math.Sqrt(float64(math.Pow(a, 2) + math.Pow(b, 2)))
}

func isEven(n float64) string {
	if int(math.Abs(n))%2 == 0 {
		return "Чётное"
	}
	return "Нечётное"
}
func isLeapYear(year int) string {
	// Високосный год: делится на 4, но не делится на 100,
	// или делится на 400
	if (year%4 == 0 && year%100 != 0) || (year%400 == 0) {
		return "Високосный"
	}
	return "Невисокосный"
}
func maxInt(s string) int {
	numbers := strings.Split(s, ", ")
	max, _ := strconv.Atoi(numbers[0])
	for i := range numbers {
		check, _ := strconv.Atoi(numbers[i])
		if check > max {
			max = check
		}
	}
	return max
}
func ages(age int) string {
	if age < 0 {
		return "Error"
	} else if age < 14 {
		return "Ребёнок"
	} else if age < 18 {
		return "Подросток"
	}
	return "Взрослый"
}
func d35(n int) string {
	if n%15 == 0 {
		return "Делится"
	}
	return "Не делится"
}

func factorial(n int) int {
	result := 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	return result
}
func fib(n int) string {
	if n < 0 {
		return "Error"
	} else if n == 0 {
		return "Empty"
	} else if n == 1 {
		return "0"
	}
	var res string = "0, 1"
	a := 0
	b := 1
	for i := 2; i < n; i++ {
		c := a + b
		a = b
		b = c
		res += ", " + strconv.Itoa(c)
	}
	return res
}
func revMas(numbers []int) []int {
	var res []int
	for i := len(numbers) - 1; i > -1; i-- {
		res = append(res, numbers[i])
	}
	return res
}
func simpleNum(n int) string {
	if n < 2 {
		return ""
	}
	var res string = "2"
	var u bool = true
	for i := 3; i <= n; i++ {
		u = true
		for j := 2; j < i; j++ {
			if i%j == 0 {
				u = false
				break
			}
		}
		if u {
			res += ", " + strconv.Itoa(i)
		}
	}
	return res
}
func sumMas(numbers []int) int {
	var res int = 0
	for _, i := range numbers {
		res += i
	}
	return res
}

func main() {
	fmt.Println(sum(341284))
	fmt.Println(temp("25 (Celsius)"))
	fmt.Println(doubleMas([]int{1, 2, 3, 4}))
	fmt.Println(strConcat([]string{"Hello", "world"}))
	fmt.Println(linePos("(x1=1, y1=1), (x2=4, y2=5)"))

	fmt.Println(isEven(-33))
	fmt.Println(isLeapYear(400))
	fmt.Println(maxInt("4, 9, 7"))
	fmt.Println(ages(24))
	fmt.Println(d35(45))

	fmt.Printf("Факториал %d равен %d\n", 5, factorial(5))
	fmt.Println(fib(10))
	fmt.Println(revMas([]int{1, 2, 3, 4, 5}))
	fmt.Println(simpleNum(42))
	fmt.Println(sumMas([]int{1, 2, 3, 4, 5}))
}
