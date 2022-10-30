package ouput

import "fmt"

// Tô màu:
//cho 1 ma trận NxM (5<=N, M<=50)
//trong mỗi ô sẽ được đánh 1 số, cần tô màu các ô thỏa mãn điều kiện sau:
//đk: bằng cách cộng hoặc trừ các số ở top, down,left, right, bạn có thể tính ra chữ số ở ô hiện tại
//in ra số ô được tô màu

func I0016() {
	n := 5
	m := 6
	arr := [][]int{
		{3, 4, 0, -1, 9, -2},
		{-1, 1, 3, 8, 10, 13},
		{5, 5, 2, 0, 6, -4},
		{-5, 7, 4, -1, 8, -9},
		{3, -7, -3, 5, -3, 12},
	}
	colored := coloring(arr, n, m)
	fmt.Println("colored: ", colored)
}

//func set() [5][6]int32 {
//	//var n, m int
//	//fmt.Println("N: ")
//	//_, err := fmt.Scanln(&n)
//	//if err != nil || n<0{
//	//	return
//	//}
//	//fmt.Println("M: ")
//	//_, err = fmt.Scanln(&m)
//	//if err != nil || m<0{
//	//	return
//	//}
//	//
//	//var array = make([][]float64, int(n),int(m))
//	//for i := 0; i < n; i++ {
//	//	for j := 0; j < m; j++ {
//	//		fmt.Println("arr[",i,"][",j,"]: ")
//	//		_, err := fmt.Scanln(&array[i][j])
//	//		if err != nil {
//	//			return
//	//		}
//	//	}
//	//}
//	//fmt.Println("arr: ")
//	//fmt.Println(array)
//	return [5][6]int32{
//		{3, 4, 0, -1, 9, -2},
//		{-1, 1, 3, 8, 10, 13},
//		{5, 5, 2, 0, 6, -4},
//		{-5, 7, 4, -1, 8, -9},
//		{3, -7, -3, 5, -3, 12},
//	}
//}

func coloring(arr [][]int, n int, m int) (colored int32) {
	colored = 0
	str := ""
	for i := 0; i < n; i++ {
		fmt.Println("i: ", i)
		for j := 0; j < m; j++ {
			top := 0
			down := 0
			left := 0
			right := 0
			num := arr[i][j]
			fmt.Println("num: ", num)
			if isset(i-1, j, n, m) {
				top = arr[i-1][j]
			}
			if isset(i+1, j, n, m) {
				down = arr[i+1][j]
			}
			if isset(i, j-1, n, m) {
				left = arr[i][j-1]
			}
			if isset(i, j+1, n, m) {
				right = arr[i][j+1]
			}
			fmt.Println("Top: ", top, ",Down: ", down, ",Left: ", left, ",Right: ", right)
			if check(num, top, down, left, right) {
				colored++
				str += fmt.Sprintf("%v,", num)
			}
		}
	}
	fmt.Println(str)
	return colored
}
func check(num int, top int, down int, left int, right int) bool {
	if top+down+left+right == num {
		return true
	}
	if top-down+left+right == num {
		return true
	}
	if top+down-left+right == num {
		return true
	}
	if top+down+left-right == num {
		return true
	}
	if top-down-left+right == num {
		return true
	}
	if top+down-left-right == num {
		return true
	}
	if top-down+left-right == num {
		return true
	}
	if top-down-left-right == num {
		return true
	}

	if -top+down+left+right == num {
		return true
	}
	if -top-down+left+right == num {
		return true
	}
	if -top+down-left+right == num {
		return true
	}
	if -top+down+left-right == num {
		return true
	}
	if -top-down-left+right == num {
		return true
	}
	if -top+down-left-right == num {
		return true
	}
	if -top-down+left-right == num {
		return true
	}
	if -top-down-left-right == num {
		return true
	}

	return false
}
func isset(i int, j int, n int, m int) bool {
	if i >= 0 && i < n && j >= 0 && j < m {
		return true
	}
	return false
}
