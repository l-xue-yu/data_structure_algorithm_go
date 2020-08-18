package main

import "fmt"

func main() {
	demoArr := []int{20, 1, 14, 5, 18, 8, 6, 19}
	println("排序前：")
	fmt.Printf("%v", demoArr)
	println("")
	//bubbleSort(demoArr)
	//standardBubbleSort(demoArr)
	//simpleSelectionSort(demoArr)
	straightInsertionSort(demoArr)
	println("排序后：")
	fmt.Printf("%v", demoArr)
}

//插入排序：一个数组，前i-1个元素是有序的，后i个是无序的，先将待插入元素i给临时变量temp，当待插入变量i小于i-1时，继续向数组第i-2个比较，

//当待插入元素小于元素i-1,且未到数组第一个元素，将前i个元素后移，当待插入元素大于元素i-1或达到第一个元素，则停止比较，交换待插入元素
func straightInsertionSort(intArr []int) {
	var j, temp int
	for i := 1; i < len(intArr); i++ {
		j = i
		//待排序元素赋值给临时变量
		temp = intArr[i]
		//循环条件，待插入元素小于当前元素或未到数组的第一个元素
		for j > 0 && temp < intArr[j-1] { //当未达到数组的第一个元素或者待插入元素小于当前元素继续循环
			intArr[j] = intArr[j-1] //将该元素后移
			j--                     //下标减一，继续比较
			fmt.Printf("排序index j：%d,%v", j, intArr)
		}
		intArr[j] = temp //插入位置已经找到，立即插入

		fmt.Printf("排序index：%d,%v", i, intArr)
		println("")
	}
}
