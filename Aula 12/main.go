package main

import "fmt"

func main() {

	sum := 0
	for i := 0; i < 10; i++ {
		sum++
	}

	pessoas := map[string]int{
		"Pedro": 18,
		"Maria": 19,
	}
	fmt.Println(sum)

	count := 0

	for i := 0; i < len(pessoas); i++ {
		count++
	}

	nums := []int{1, 2, 3, 4, 5}

	for i := 0; i < len(nums); i++ {
		fmt.Println(nums[i])
	}

	fmt.Println("A quantidade de pessoas na lista é de : ", count)
}
