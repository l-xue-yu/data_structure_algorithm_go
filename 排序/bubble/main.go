package main

import "fmt"

func main() {
	//demoArr := []int{7, 13, 4, 5, 8, 1, 11, 9}
	demoArr := []int{2, 1, 4, 5, 8, 8, 10, 19}
	println("排序前：")
	fmt.Printf("%v", demoArr)
	//bubbleSort(demoArr)
	//standardBubbleSort(demoArr)
	superBubbleSort(demoArr)
	println("")
	println("排序后：")
	fmt.Printf("%v", demoArr)

}

//从小到大排序
//最简单排序
func bubbleSort(intArr []int) {

	for i := 0; i < len(intArr); i++ {
		for j := i + 1; j < len(intArr); j++ {
			if intArr[i] > intArr[j] {
				intArr[i], intArr[j] = intArr[j], intArr[i]
			}
			println("排序i：", i, "j:", j)
			fmt.Printf("%v", intArr)
		}
	}

}

//标准冒泡
//与上面最基本排序好在，比较过的都是有序的
func standardBubbleSort(intArr []int) {
	for i := 0; i < len(intArr); i++ {
		//j从后往前循环，初始为倒数第二个和最后一个比较，依次从后往前，从底向上冒泡
		for j := len(intArr) - 2; j >= i; j-- {

			if intArr[j] > intArr[j+1] {
				//如果前者大于后者，交换
				intArr[j], intArr[j+1] = intArr[j+1], intArr[j]
			}
			println("排序i：", i, "j:", j)
			fmt.Printf("%v", intArr)
		}

	}

}

//冒泡优化
//增加标记变量
func superBubbleSort(intArr []int) {
	flag := true
	for i := 0; i < len(intArr) && flag; i++ {
		flag = false
		for j := len(intArr) - 2; j >= i; j-- {
			if intArr[j] > intArr[j+1] {
				intArr[j], intArr[j+1] = intArr[j+1], intArr[j]
				flag = true
			}
			println("排序i：", i, "j:", j)
			fmt.Printf("%v", intArr)
		}
	}
}
