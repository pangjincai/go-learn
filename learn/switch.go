package main

import "fmt"

func main() {

	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}

	switch  {

	case k < 6:
		fmt.Println("k less than 6")
	default:
		fmt.Println("k is equal 6")
	}

	test(10)

}

func test(i int) {
	switch i {
	case 12,1,2:
		fmt.Println("冬天")
	case 3,4,5:
		fmt.Println("春天")
	case 6,7,8:
		fmt.Println("夏天")
	case 9,10,11:
		fmt.Println("秋天")
	default:
		fmt.Println("未知")
	}
}
