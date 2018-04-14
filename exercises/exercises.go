package exercises

import "fmt"

func hello() {
	fmt.Println("Hello World")
}

func helloName() {
	fmt.Println("Hello Andrew")
}

func helloScanName() {
	var name string
	fmt.Scan(&name)
	fmt.Printf("Hello %s\n", name)
}

func remainder(min, max int) {
	fmt.Println(max % min)
}

func evens() {
	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}
}

func fizzbuzz() {
	for i := 0; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Print("Fizz")
		} else if i%5 == 0 {
			fmt.Print("Buzz")
		} else {
			fmt.Print(i)
		}
		fmt.Println()
	}
}

func threesandfives() {
	var sum int
	for i := 0; i <= 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			sum += i
		}
	}
	fmt.Println(sum)
}
