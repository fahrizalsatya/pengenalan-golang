package main

import (
	"fmt"

	"github.com/FahrizalSatya/pengenalan-golang/Basic_Package_Variable_Functions/package_1"
)

func main() {

	var x, y int

	nama1, nama2 := "fahrizal", "satya"

	x = 10
	y = 5

	result1 := package_1.Add(x, y)
	fmt.Println(fmt.Sprintf("Hasil penjumlahannya adalah %d", result1))

	fmt.Println(package_1.CustomPrint(nama1))
	fmt.Println(package_1.CustomPrint(nama2))

	result2, flag2 := package_1.AddDoubleReturn(x, y)

	fmt.Println(fmt.Sprintf("Hasil penjumlahannya adalah %d , %t", result2, flag2))
}
