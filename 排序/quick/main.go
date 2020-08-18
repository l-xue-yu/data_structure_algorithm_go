package main

import "fmt"

func main() {
	intArr := []int{6, 1, 2, 7, 9, 3, 4, 5, 10, 8}
	fmt.Printf("快速排序前：%v", intArr)
	quickSort(0, 9, intArr)
	println("")
	fmt.Printf("快速排序后：%v", intArr)

}

/*
参考博客：https://blog.csdn.net/qq_28584889/article/details/88136498?utm_medium=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-3.channel_param&depth_1-utm_source=distribute.pc_relevant.none-task-blog-BlogCommendFromMachineLearnPai2-3.channel_param
唯一一个一遍就看懂的快排算法
1.先进行第一轮快排，哨兵右从右往左，直到找到比基准数小的停下，哨兵左从左往右，直到找到比基准数大的停下
2.交换左右哨兵
3.左右哨兵继续往下找，直到左右哨兵相遇
4.交换基准数与哨兵（归位），第一轮结束
5.对此时基准数左边的进行上面的操作
6.对此时基准数右边的进行上面的操作
*/
func quickSort(left int, right int, intArr []int) {
	if left >= right {
		return
	}
	i := left
	j := right
	base := intArr[left] //最左边为基数
	temp := 0
	for i < j {
		for intArr[j] >= base && i < j { //从右边开始，'哨兵右'大于基数继续往左走，直到'哨兵右'小于基数
			j--
		}
		for intArr[i] <= base && i < j { //右边的'哨兵右'小于基数，则开始从左边开始，'哨兵左'小于基数则继续往右走，直到'哨兵左'大于基数
			i++
		}
		if i < j { //这时，'哨兵右'小于基数，'哨兵左'大于基数，交换左右哨兵
			temp = intArr[i]
			intArr[i] = intArr[j]
			intArr[j] = temp
		}
		println("")
		fmt.Printf("快速排序后：%v", intArr)
	}

	//此轮哨兵扫描交换结束，i==j,将intArr[i]与基准数交换 基准数归位
	intArr[left] = intArr[i]
	intArr[i] = base
	quickSort(left, i-1, intArr)  //对左边进行快排递归
	quickSort(i+1, right, intArr) //对右边进行快排递归

}
