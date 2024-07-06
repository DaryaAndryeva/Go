package main

import "fmt"

// Привет, мир!
func task1() {
	fmt.Println("Привет, мир!")
}

// Сложение чисел
func task2(x, y int) int {
	return x + y
}

// Четное или нечетное
func task3(x int) {
	if x%2 == 1 {
		fmt.Printf("%v is odd\n", x)
	} else {
		fmt.Printf("%v is even\n", x)
	}
}

// Максимум из трех чисел
func task4(x, y, z int) int {
	return max(x, y, z)
}

// Факториал числа
func task5(n int) int {
	f := 1
	for i := 1; i <= n; i++ {
		f *= i
	}
	return f
}

// Проверка символа
func task6(s rune) {
	if s == 'a' || s == 'e' || s == 'y' || s == 'u' || s == 'i' || s == 'o' ||
		s == 'A' || s == 'E' || s == 'Y' || s == 'U' || s == 'I' || s == 'O' {
		fmt.Println(string(s), "is a vowel letter")
	} else {
		fmt.Println(string(s), "is not a vowel letter")
	}
}

// Простые числа
func task7(n int) {
	sieve := make([]int, n+1)
	sieve[0] = 1
	sieve[1] = 1
	for i := 2; i <= n; i++ {
		if sieve[i] == 0 {
			for j := i * i; j <= n; j += i {
				sieve[j] = 1
			}
		}
	}
	fmt.Printf("Все простые до %v:\n", n)
	for i := 1; i <= n; i++ {
		if sieve[i] == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()
}

// Строка и ее перевертыш
func task8(s string) string {
	t := ""
	for i := len(s) - 1; i >= 0; i-- {
		t += string(s[i])
	}
	return t
}

// Массив и его сумма
func task9(arr []int) int {
	sum := 0
	for _, val := range arr {
		sum += val
	}
	return sum
}

// Структуры и методы
type Rectangle struct {
	height, width int
}

func (r Rectangle) area() int {
	return r.height * r.width
}

func main() {
	var x, y, z int

	// 1
	task1()
	fmt.Println()

	// 2
	x = 3
	y = 9
	fmt.Printf("Сумма %v и %v -- %v\n\n", x, y, task2(x, y))

	// 3
	task3(17)
	fmt.Println()

	// 4
	x = 5
	y = 7
	z = 4
	fmt.Printf("Максимум из %v, %v и %v -- %v\n\n", x, y, z, task4(x, y, z))

	// 5
	fmt.Printf("Факториал %v -- %v\n\n", 6, task5(6))

	// 6
	task6('k') // Проверка символа
	fmt.Println()

	// 7
	task7(100)
	fmt.Println()

	// 8
	s := "qwerty"
	fmt.Println(s, "--", task8(s))
	fmt.Println()

	// 9
	arr := []int{1, 2, 3, 4, 5}
	fmt.Printf("Сумма %o -- %v\n", arr, task9(arr))
	fmt.Println()

	// 10
	r := Rectangle{4, 9}
	fmt.Printf("Площадь прямоугольника со сторонами %d и %d -- %v\n", r.height, r.width, r.area())
	fmt.Println()
}
