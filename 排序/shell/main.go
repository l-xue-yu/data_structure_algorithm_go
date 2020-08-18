package main

import "fmt"

func main() {
	demoArr := []int{7, 13, 4, 5, 8, 1, 11, 9}
	//ShellSort2(demoArr)

	cutSlice(demoArr, 7)

}

//https://blog.csdn.net/qq_36183935/article/details/80224171
func shellSort(demoArr []int) {
	fmt.Println(demoArr)
	d := len(demoArr) / 2
	for d > 0 {
		println("步长：", d)
		for i := 0; i < d; i++ {
			println("i", i)
			for j := i + d; j < len(demoArr); j += d {
				temp := demoArr[j]
				k := j - d
				for k >= i && demoArr[k] > temp {
					demoArr[k+d] = demoArr[k]
					k -= d
				}
				demoArr[k+d] = temp
			}
		}
		d = d / 2
	}
	fmt.Println(demoArr)
}

//https://www.cnblogs.com/xxzhuang/p/7333753.html
func ShellSort2(num []int) {
	fmt.Println(num)
	//increment相隔数量
	for increment := len(num) / 2; increment > 0; increment /= 2 { //步长increment ：5,2,1,0
		//i序号较大的数组下标，i ,j进行比较
		println("步长：", increment)
		for i := increment; i < len(num); i++ { //每一个分组进行排序 i ， 5 ,6, 7,8,9,10      2,3,4,5,6,7,8,9   1,2,3,4,5,6,7,8,9
			println("i", i)
			//进行交换
			temp := num[i]
			//按照increment，数组从j到0进行交换比较
			for j := i - increment; j >= 0; j -= increment { // 0,
				print("j:", num[j])
				if temp < num[j] {
					num[j+increment] = num[j]
					num[j] = temp
					temp = num[j]
				} else { //由于数组前面按照increment已经排好序，如果temp>num[j],则不必继续比较交换下去
					break
				}
			}

		}

	}

	fmt.Println(num)

}

//数组，切片切割
func cutSlice(oldSlice []int, index int) {
	println("")
	fmt.Printf("index,%d old %v", index, oldSlice)
	oldSlice = append(oldSlice[:index], oldSlice[index+1:]...)
	fmt.Printf("new %v", oldSlice)
}
