package main

import (
	"bytes"
	"fmt"
)

func main() {
	a := 5;
	var buffer bytes.Buffer
	for i := 3; i < 10; i++ {
		buffer.WriteString("hihihihihi\n");
	}
	fmt.Println(buffer.String());
	fmt.Printf(":",a);
}