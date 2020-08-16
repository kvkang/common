package ipdb

import (
	"fmt"
)

// ExampleDistrict for use district example
func ExampleDistrict() {
	got, err := NewDistrict("./city.free.ipdb")
	if err != nil {
		fmt.Printf("open ipdb faild: %v\n", err.Error())
		return
	}
	values, mask, err := got.Find("211.3.11.1", ChineseLanguage)
	fmt.Println(values, mask, err)
}
