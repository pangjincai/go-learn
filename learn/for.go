package main

import "fmt"

func main()  {
	str := "G"
	for i := 1; i <= 25; i++ {
		println(str)
		str += "G"
	}

	for i := 1;i <= 100;i++ {
		if i%3 == 0 || i%5 == 0 {
			fmt.Println("Fizz")
		} else {
			fmt.Println(i)
		}
	}

	for i := 0; i< 3; i++ {
		fmt.Println("Value of i is now:", i)
	}
}
