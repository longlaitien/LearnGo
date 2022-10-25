package main

import (
	"fmt"
	"math"
)

func main() {
	run := 0
	fmt.Println("Case chay: ")
	_, _ = fmt.Scanln(&run)
	switch run {
	case 1:
		fmt.Print("Chiu Chiu Chiu!!!")
	case 2:
		a := 5
		b := 10.5
		c := " da xong"
		fmt.Printf("a: %v,%T; b:%v,%T %v", a, a, b, b, c)
	case 3:
		var a float64
		var b float64
		fmt.Println("So nguyen:")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println("So thuc:")
		_, err = fmt.Scanln(&b)
		if err != nil || b == 0 {
			fmt.Println("B error")
		}
		fmt.Println("2 so vua nhap: ", a, " va ", b)
		fmt.Println("Tong: ", a+b)
		fmt.Println("Tich: ", a*b)
		fmt.Println("Du: ", math.Mod(a, b))
	case 4:
		var d, r float64
		for {
			fmt.Println("dai: ")
			_, _ = fmt.Scanln(&d)
			fmt.Println("rong: ")
			_, _ = fmt.Scanln(&r)
			if d > r {
				break
			}
		}
		c := (d + r) * 2
		s := d * r
		fmt.Println("Chu vi va dien tich: ", c, " va ", s)
	case 5:
		var c1, c2, c3 float64
		for {
			fmt.Println("canh 1: ")
			_, _ = fmt.Scanln(&c1)
			fmt.Println("canh 2: ")
			_, _ = fmt.Scanln(&c2)
			fmt.Println("canh 3: ")
			_, _ = fmt.Scanln(&c3)
			if c1 > 0 && c2 > 0 && c3 > 0 {
				break
			} else {
				fmt.Println("Nhap lai!")
			}
		}
		c := c1 + c2 + c3
		p := c / 2
		s := math.Sqrt(p * (p - c1) * (p - c2) * (p - c3))

		fmt.Println("Chu vi va dien tich: ", c, " va ", s)
	case 6:
		h := 7
		w := 11
		s := "*"
		//L
		for i := 1; i <= h; i++ {
			for j := 1; j <= w; j++ {
				if i == j || i == h {
					fmt.Print(s)
				}
			}
			fmt.Println(" ")
		}
		fmt.Println(" ")
		//m
		n := 1
		m := w
		for i := 1; i <= h; i++ {
			for j := 1; j <= w; j++ {
				if j == 1 || j == w || (n == j && n < m) || (m == j && n < m) || (n == m && n == w/2) {
					fmt.Print(s)
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println(" ")
			n += w / h
			m -= w / h
		}

		fmt.Println(" ")
		//o
		for i := 1; i <= h; i++ {
			for j := 1; j <= w; j++ {
				if i == 1 || j == 1 || j == w || i == h {
					fmt.Print(s)
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println(" ")
		}

		fmt.Println(" ")
		//H
		for i := 1; i <= h; i++ {
			for j := 1; j <= w; j++ {
				if j == 1 || j == w || i == h/2 {
					fmt.Print(s)
				} else {
					fmt.Print(" ")
				}
			}
			fmt.Println(" ")
		}
	case 7:
		var s1, s2, s3, stg int
		for {
			fmt.Println("So1:")
			_, err := fmt.Scanln(&s1)
			fmt.Println("So 2:")
			_, err = fmt.Scanln(&s2)
			fmt.Println("So 3:")
			_, err = fmt.Scanln(&s3)
			if err == nil {
				break
			}
		}
		if s1 < s2 {
			stg = s1
			s1 = s2
			s2 = stg
		}
		if s3 >= s1 {
			fmt.Println("Cac so vua nhap theo thu tu tang dan: ", s3, s1, s2)
			fmt.Println("Cac so vua nhap theo thu tu giam dan: ", s2, s1, s3)
		} else {
			if s3 >= s2 {
				fmt.Println("Cac so vua nhap theo thu tu tang dan: ", s1, s3, s2)
				fmt.Println("Cac so vua nhap theo thu tu giam dan: ", s2, s3, s1)
			} else {
				fmt.Println("Cac so vua nhap theo thu tu tang dan: ", s1, s2, s3)
				fmt.Println("Cac so vua nhap theo thu tu giam dan: ", s3, s2, s1)
			}
		}
		break
	case 8:
		var n int
		s := 0
		fmt.Println("Nhap so n: ")
		_, _ = fmt.Scanln(&n)
		for i := 1; i <= n; i++ {
			s += i
		}
		fmt.Println("S(n) = 1 + 2 + 3 + … + n = ", s)
	case 9:
		var n int
		s := 0.0
		fmt.Println("Nhap so n: ")
		_, _ = fmt.Scanln(&n)
		for i := 1; i <= n; i++ {
			s += math.Pow(float64(i), 2.0)
			fmt.Println(s)
		}
		fmt.Println("S(n) = 1^2 + 2^2 + … + n^2 = ", s)
	}

}
