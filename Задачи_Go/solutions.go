package main

import (
	"fmt"
	"math"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func distance() float64 {
	var (
		v,
		a,
		t float64
	)
	fmt.Print("Введите начальную скорсть: ")
	fmt.Scanf("%f\n", &v)
	fmt.Print("Введите ускорение: ")
	fmt.Scanf("%f\n", &a)
	fmt.Print("Введите время движения (>=0): ")
	for {
		fmt.Scanf("%f\n", &t)
		if t >= 0 {
			break
		}
		fmt.Print("Время не может быть отрицательным\nВведите время ещё раз\n")
	}
	return v*t + (a*t*t)/2
}

func systemPol(s []string) float64 {
	//fmt.Println(s)
	for i := 2; i < len(s); i++ {
		if s[i] == "+" || s[i] == "-" || s[i] == "*" || s[i] == "/" {
			a, _ := strconv.ParseFloat(s[i-2], 64)
			b, _ := strconv.ParseFloat(s[i-1], 64)
			var res float64
			switch s[i] {
			case "+":
				res = a + b
			case "-":
				res = a - b
			case "*":
				res = a * b
			case "/":
				res = a / b
			}
			f := s[:i-2]
			//fmt.Println("f - ", f)
			t := fmt.Sprintf("%f", res)
			f = append(f, t)

			for j := i + 1; j < len(s); j++ {
				f = append(f, s[j])
				//fmt.Println("wf - ", f)
			}
			return systemPol(f)
		}
	}
	res, _ := strconv.ParseFloat(s[0], 64)
	return res
}
func bigYear(year int) bool {
	if (year%4 == 0 && year%100 != 0) || year%400 == 0 {
		return true
	}
	return false
}
func solSquare(a, b, c float64) string {
	d := b*b - 4*a*c
	if d > 0 {
		x1 := (-b + math.Sqrt(d)) / (2 * a)
		x2 := (-b - math.Sqrt(d)) / (2 * a)
		return "x1 = " + strconv.FormatFloat(x1, 'g', -1, 64) + "\nx2 = " + strconv.FormatFloat(x2, 'g', -1, 64) + "\n"
	} else if d == 0 {
		x := -b / (2 * a)
		return "x = " + strconv.FormatFloat(x, 'g', -1, 64) + "\n"
	}
	return "Решений нет\n"
}
func calculatorOP(s string) float64 {
	op := strings.Split(s, " ")
	a, _ := strconv.ParseFloat(op[0], 64)
	b, _ := strconv.ParseFloat(op[1], 64)
	if op[2] == "+" {
		return a + b
	} else if op[2] == "-" {
		return a - b
	} else if op[2] == "*" {
		return a - b
	} else if b == 0 {
		return math.NaN()
	}
	return a / b
}
func numberFibIt(n int) {
	if n < 1 {
		return
	}
	if n == 1 {
		println(1)
		return
	}
	if n == 2 {
		println("1 1")
		return
	}
	print("1 1")
	a := 1
	b := 1
	var temp int
	for i := 3; i <= n; i++ {
		temp = a
		a = b
		b = a + temp
		print(" ", b)
	}
	return
}

func numberFibRec(last, n, cur, l int) {
	if l < 1 {
		return
	}
	print(n, " ")
	if cur == l {
		return
	}
	numberFibRec(n, last+n, cur+1, l)
}
func simpleNumber(n int) bool {
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
func numberArmstrong(n int) bool {
	a := n / 100
	b := (n / 10) % 10
	c := n % 10
	return n == a*a*a+b*b*b+c*c*c
}
func tableMultiply(a, b int) {
	if a > b || b > 10 || a < 1 {
		println("Не верно!")
		return
	}
	for i := a; i <= b; i++ {
		for j := a; j <= b; j++ {
			print(i*j, "\t")
		}
		println()
	}
}
func nod(a, b int) int {
	if a == b {
		return a
	}
	if b > a {
		return nod(a, b-a)
	}
	return nod(a-b, b)
}
func simpleRange(a, b int) {
	for i := a; i <= b; i++ {
		if simpleNumber(i) {
			print(i, " ")
		}
	}
}
func tryLucky() {
	// Инициализация источника случайных чисел с текущим временем
	rand.Seed(time.Now().UnixNano())

	n := rand.Intn(101)
	var pol int
	for i := 0; i < 10; i++ {
		fmt.Scan(&pol)
		if n == pol {
			println("You win!!!")
			return
		}
		if pol < n {
			println("More")
		} else {
			println("Less")
		}
	}
	println("You lose!!!\nNumber was ", n)
}
func rhombus(n int) {
	if n < 1 || n%2 == 0 {
		println("Not right")
		return
	}
	for i := 0; i < n/2+1; i++ {
		for j := 0; j < n/2-i; j++ {
			print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			print("*")
		}
		println()
	}
	for i := n/2 - 1; i > -1; i-- {
		for j := 0; j < n/2-i; j++ {
			print(" ")
		}
		for j := 0; j < i*2+1; j++ {
			print("*")
		}
		println()
	}
}
func tenToDouble(n int) int {
	res := 0
	for i := 1; n != 0; i *= 10 {
		res += n % 2 * i
		n /= 2
	}
	return res
}


func countNeighbours(state [][]int, x, y int) (res int){
	res = 0
	for	i:= x-1; i < x+2; i++{
		if y-1 < 0 || y-1 > len
	}
	return
}
func gameOfLife(start [][]int, turns int) {
	next := [8][8]int{}


	for turn := 0; turn < turns; turn++ {
		for i := 0; i < len(start); i++ {
			neighbours := 0
			for y:=i-1;y<i+1;y++{
				if start[y][j]{

				}
			}
		}
	}
}

#---------------------------------------------------

const (
	alive = 1
	dead  = 0
)

func printGrid(grid [][]int) {
	for _, row := range grid {
		for _, cell := range row {
			if cell == alive {
				fmt.Print("O ")
			} else {
				fmt.Print(". ")
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func countNeighbors(grid [][]int, x, y int) int {
	neighbors := 0
	rows := len(grid)
	cols := len(grid[0])

	directions := []struct{ dx, dy int }{
		{-1, -1}, {-1, 0}, {-1, 1},
		{0, -1}, {0, 1},
		{1, -1}, {1, 0}, {1, 1},
	}

	for _, d := range directions {
		nx, ny := x+d.dx, y+d.dy
		if nx >= 0 && nx < rows && ny >= 0 && ny < cols && grid[nx][ny] == alive {
			neighbors++
		}
	}
	return neighbors
}

func nextGeneration(grid [][]int) [][]int {
	rows := len(grid)
	cols := len(grid[0])
	newGrid := make([][]int, rows)
	for i := range newGrid {
		newGrid[i] = make([]int, cols)
	}

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			neighbors := countNeighbors(grid, i, j)

			if grid[i][j] == alive {
				if neighbors == 2 || neighbors == 3 {
					newGrid[i][j] = alive
				} else {
					newGrid[i][j] = dead
				}
			} else {
				if neighbors == 3 {
					newGrid[i][j] = alive
				}
			}
		}
	}

	return newGrid
}

func gameOfLife(){
	var rows, cols, steps int

	fmt.Print("Введите количество строк: ")
	fmt.Scan(&rows)
	fmt.Print("Введите количество столбцов: ")
	fmt.Scan(&cols)

	grid := make([][]int, rows)
	for i := range grid {
		grid[i] = make([]int, cols)
	}

	fmt.Println("Введите начальное состояние поля (0 - мертвая клетка, 1 - живая клетка):")
	scanner := bufio.NewScanner(os.Stdin)
	for i := 0; i < rows; i++ {
		scanner.Scan()
		line := scanner.Text()
		values := strings.Fields(line)
		for j := 0; j < cols; j++ {
			val, _ := strconv.Atoi(values[j])
			grid[i][j] = val
		}
	}

	fmt.Print("Введите количество шагов симуляции: ")
	fmt.Scan(&steps)

	fmt.Println("Начальное состояние:")
	printGrid(grid)

	for step := 0; step < steps; step++ {
		grid = nextGeneration(grid)
		fmt.Printf("Шаг %d:\n", step+1)
		printGrid(grid)
	}
}

func main() {
	fmt.Println(distance())
	fmt.Println(systemPol(strings.Split("8 9 + 1 7 - *", " ")))
	fmt.Println(bigYear(700))
	fmt.Println(solSquare(float64(25)/49, 0, -1))
	fmt.Println(calculatorOP("5 0 /"))

	numberFibIt(15)
	println()
	numberFibRec(0, 1, 1, 15)
	fmt.Println(simpleNumber(1341))
	fmt.Println(numberArmstrong(153))
	tableMultiply(5, 10)
	fmt.Println(nod(108, 300))

	simpleRange(-6, 45)
	tryLucky()
	rhombus(7)
	fmt.Println(tenToDouble(43212))
	gameOfLife()
}
