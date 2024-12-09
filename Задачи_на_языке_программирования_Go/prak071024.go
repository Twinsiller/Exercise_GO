package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ConvertInt(val string, base, toBase int) (string, error) {
	i, err := strconv.ParseInt(val, base, 64)
	f := 1
	if err != nil {
		return "", err
	}
	return strconv.FormatInt(i, toBase), nil
}
func u1() string {

	var (
		s string
		a,
		b int
	)
	fmt.Scanf("%s %d %d", &s, &a, &b)
	res, err := ConvertInt(s, a, b)
	if err != nil {
		return "Error"
	}
	return res
}

func u2() string {
	var (
		a,
		b,
		c float64
	)
	fmt.Scanf("%f %f %f", &a, &b, &c)
	d := b*b - 4*a*c
	if d > 0 {
		r1 := (-b + math.Sqrt(d)) / (2 * a)
		r2 := (-b - math.Sqrt(d)) / (2 * a)
		return fmt.Sprintf("%f", r1) + " " + fmt.Sprintf("%f", r2)
	} else if d == 0 {
		res := (-b) / (2 * a)
		return fmt.Sprintf("%f", res)
	}
	r1 := complex((-b)/(2*a), (math.Sqrt(-d))/(2*a))
	r2 := complex((-b)/(2*a), -((math.Sqrt(-d)) / (2 * a)))
	return fmt.Sprintf("%f", r1) + " " + fmt.Sprintf("%f", r2)
}
func u3(numbers []int) []int {
	var temp int
	for i := 0; i < len(numbers)-1; i++ {
		for j := i + 1; j < len(numbers); j++ {
			if math.Abs(float64(numbers[i])) > math.Abs(float64(numbers[j])) {
				temp = numbers[i]
				numbers[i] = numbers[j]
				numbers[j] = temp
			}
		}
	}
	return numbers
}
func u4(ns1 []int, ns2 []int) []int {
	var (
		a       = 0
		b   int = 0
		res []int
	)
	for a < len(ns1) && b < len(ns2) {
		if ns1[a] < ns2[b] {
			res = append(res, ns1[a])
			a++
		} else {
			res = append(res, ns2[b])
			b++
		}
	}
	if a == len(ns1) {
		for ; b < len(ns2); b++ {
			res = append(res, ns2[b])
		}
	} else {
		for ; a < len(ns1); a++ {
			res = append(res, ns1[a])
		}
	}
	return res
}
func u5(s1, s2 string) int {
	var res int = -1
	var check bool
	if len(s2) > len(s1) {
		s1, s2 = s2, s1
	}
	for i, ch := range s1 {
		if string(ch) != string(s2[0]) {
			continue
		}
		check = true
		for j, val := range s2 {
			if string(val) != string(s1[i+j]) {
				check = false

			}
		}
		if check {
			res = i
			break
		}
	}
	return res
}

func us1(a, b float64, op rune) (float64, string) {
	if op == '+' {
		return a + b, ""
	} else if op == '-' {
		return a - b, ""
	} else if op == '*' {
		return a * b, ""
	} else if op == '/' {
		if b == 0 {
			return 0, "шибка"
		}
		return a / b, ""
	} else if op == '^' {
		return math.Pow(a, b), ""
	}
	if a-math.Round(a) == 0 && b-math.Round(b) == 0 {
		return float64(int(a) % int(b)), ""
	}
	return 0, "шибка"
}

func us2(s string) bool {
	//fmt.Println(int('a'), int('z'))
	s = strings.ToLower(s)
	res := ""
	for _, ch := range s {
		if 96 < int(ch) && int(ch) < 123 {
			res += string(ch)
		}
	}
	s = res
	res = ""
	for _, ch := range s {
		res = string(ch) + res
	}
	fmt.Println(s, res)
	return res == s
}
func us3ALT(ns []int) bool {
	for i := 2; i < len(ns); i++ {
		if i%2 == 0 {
			if ns[i] > ns[1] {
				return false
			}
		} else {
			if ns[i] < ns[0] {
				return false
			}
		}
	}
	return true
}
func us3(ns []int) bool{
	if ns == nil{
		return true
	}
	if ns[0] < ns[2] && ns[1] < ns[3]
}

func main() {
	//fmt.Println(u1())
	//fmt.Println(u2())
	//fmt.Println(u3([]int{23, -24, 21, 5, -6, 90, -36}))
	//fmt.Println(u4([]int{1, 5, 7, 23, 56, 73, 91}, []int{2, 4, 15, 35, 57, 87, 94, 101}))
	//fmt.Println(u5("qweryurtiop", "rt"))

	//fmt.Println(us1(5, 2.5, '%'))
	//fmt.Println(us2("q WE, ew,q"))
	fmt.Println(us3([]int{5, 8, 8, 15, 2, 7}))
}
