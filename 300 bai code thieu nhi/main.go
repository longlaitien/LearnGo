package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	run := 2
	switch run {
	case 1:
		fmt.Print("Chiu Chiu Chiu!!!")
	case 2:
		a := 5
		b := 10.5
		c := " da xong"
		fmt.Printf("a: %v,%T; b:%v,%T %v", a, a, b, b,c)
	case 3:
		scanner := bufio.NewScanner(os.Stdin);
		fmt.Println("Nhap: ");
		for scanner.Scan(){
			a := scanner.Text();
			if a == "0"{
				break;
			}
			fmt.Println("Chuoi ban vua nhap",a);
		}
		fmt.Println("Exit code 0!");
	}
}
