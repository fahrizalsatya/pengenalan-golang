package package_2

import (
	"fmt"

	"github.com/FahrizalSatya/pengenalan-golang/Basic_Package_Variable_Functions/package_1"
)

func PrintPackage2() {
	fmt.Println("Print dari package_2")
}

func ImportDariPackage1() {
	package_1.PrintPackage1()
}
