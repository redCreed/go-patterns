package main

import "fmt"

func main() {
	vBuilder := newVillaBuilder()
	d := newDirector(vBuilder)
	villa := d.buildHouse()

	fmt.Println("window:", villa.window)
	fmt.Println("door:", villa.door)
	fmt.Println("swimmingPool:", villa.swimmingPool)
	fmt.Println("floor:", villa.floor)
}
