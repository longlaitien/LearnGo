package ouput

import "math"

func I0019() {
	n := 6
	m := 3
	paper := [6][6]int{
		{5, 1, 8, 2, 4, 3},
		{6, 5, 2, 1, 3, 7},
		{1, 1, 5, 6, 1, 2},
		{2, 7, 1, 5, 7, 6},
		{4, 5, 7, 1, 2, 3},
		{5, 2, 3, 7, 4, 8},
	}
	crease := math.Floor(float64(n / m))
	var rs [3][6]int

	for i:=0; i<6;i++{
		for j:=0;j<6;j++{

		}
	}

	//6-3
	1-4-5 2-3-6

	//12-6 nếp =2
	1-4-5-8-9-12
	2-3-6-7-10-11

	0-3-4-7-8-11
	1-2-5-6-9-10
}
