package main

import "fmt"

func main() {
	demoArr := []int{2, 1, 4, 5, 8, 8, 10, 19}
	println("排序前：")
	fmt.Printf("%v", demoArr)
	//bubbleSort(demoArr)
	//standardBubbleSort(demoArr)
	simpleSelectionSort(demoArr)
	println("")
	println("排序后：")
	fmt.Printf("%v", demoArr)
}

//通过n-i次关键字的比较，从n-i+1个记录中选出最小的记录，并和第i个记录交换
func simpleSelectionSort(intArr []int) {
	for i := 0; i < len(intArr); i++ {
		minIndex := 0
		for j := i + 1; j <= len(intArr)-1; j++ {
			if intArr[i] < intArr[j] {
				minIndex = j //将此值下标给minIndex
			}
			if minIndex != i { //若min不等于i，说明找到最小值，交换
				intArr[i], intArr[j] = intArr[j], intArr[i]
			}
		}
	}
}
