package main


import (
	"fmt"
)

type Sisi struct {
	Panjang int
	Lebar int
}

func (s Sisi) Luas() int {
	return s.Panjang * s.Lebar
}

func(s *Sisi) Scale(i int){
	s.Lebar = s.Lebar*i
	s.Panjang = s.Panjang*i
}

func main() {
	s := Sisi{3, 4}
	s.Scale(3)
	fmt.Println(s.Luas())
}
