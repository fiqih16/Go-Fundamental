package map_

import "fmt"

func Mymap() {
	var myMap map[string]int
	myMap = make(map[string]int)

	myMap["key1"] = 1
	myMap["key2"] = 20
	myMap["key3"] = 30

	fmt.Println(myMap["key2"])
}

func Mymaptwo(){
	myMap := map[string]string{"key1": "value1", "key2": "value2"}

	fmt.Println(myMap)
}

func Mymapthree(){
	myMap := map[string]string{"kotlin": "mobileApp", "go": "webApp"}

	for key, value := range myMap {
		fmt.Println("key :", key, ", value :", value)
	}
}

func Deletemap(){
	myMap := map[string]string{"kotlin": "mobileApp", "go": "webApp", "php": "api"}

	for key, value := range myMap {
		fmt.Println("key :", key, ", value :", value)
	}

	fmt.Println("#################")
	delete(myMap, "php")

	for key, value := range myMap {
		fmt.Println("key :", key, ", value :", value)
	}
}

func CheckValueMap(){
	myMap := map[string]string{"kotlin": "mobileApp", "go": "webApp", "php": "api"}

	value, ok := myMap["php"]
	if ok {
		fmt.Println("value :", value)
	} else {
		fmt.Println("value tidak ditemukan")
	}
}