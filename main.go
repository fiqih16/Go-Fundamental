package main

import (
	"fmt"
	"fundamental/array"
	"fundamental/calculation"
	"fundamental/function"
	"fundamental/ifcontrol"
	"fundamental/map_"
	"fundamental/mapslice"
	"fundamental/method"
	"fundamental/perulanganfor"
	"fundamental/pointer"
	"fundamental/slice"
	"fundamental/struct_"
	"fundamental/switchcontrol"
)

func main() {
	fmt.Println("Hello World, saya sedang belajar golang")
	// Jalankan calculation
	result := calculation.Add(1, 2)
	fmt.Println(result)
	resultKali := calculation.Perkalian(5, 5)
	fmt.Println(resultKali)

	// Jalankan ifcontrol
	ifcontrol.Age()
	ifcontrol.Score()

	// Jalankan switchcontrol
	switchcontrol.Switch()

	// Jalankan perulangan for
	perulanganfor.Forone()
	perulanganfor.Fortwo()
	perulanganfor.Forthree()
	perulanganfor.Findeven()

	// Jalankan array
	array.Arrayone()
	array.Arraytwo()

	// Jalankan slice
	slice.Sliceone()

	// Jalankan map
	map_.Mymap()
	map_.Mymaptwo()
	map_.Mymapthree()
	map_.Deletemap()
	map_.CheckValueMap()

	// Jalankan mapslice
	mapslice.Student()
	mapslice.Score()
	mapslice.Scoretwo()

	// Jalankan function
	function.Printresult()
	function.Printresultkuis()

	// Jalankan structbasic
	struct_.StructUser()
	struct_.PrintUserParam()
	struct_.StructUserGroup()

	// Jalankan method
	method.PrintMethod()

	// Jalankan pointer
	pointer.PrintPointer()
}
// 3:26:36