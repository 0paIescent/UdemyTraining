package exercises

import "fmt"

func Hello() {
	fmt.Println("Hello World")
}

func HelloName() {
	fmt.Println("Hello Andrew")
}

func HelloScanName() {
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %s\n", name)
}

func Remainder(min, max int) {
	fmt.Println(max % min)
}

func Evens() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func FizzBuzz() {
	for i := 0; i <= 100; i++ {
		fmt.Printf("%d ", i)
		if i%3 == 0 {
			fmt.Print("Fizz")
		}
		if i%5 == 0 {
			fmt.Print("Buzz")
		}
		fmt.Println()
	}
}

func ThreesAndFives() {
	var sum int
	for i := 1; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}

func Half(n int) (int, bool) {
	return n / 2, n/2 == 0
}

func Greatest(numbers ...int) int {
	var largest int
	for _, v := range numbers {
		if v > largest {
			largest = v
		}
	}
	return largest
}
