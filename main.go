package main

import (
	"./srcfile"
)

func main() {
	list, _ := srcfile.TextList(`C:\git\transtool\abcd.txt`)
	list2, _ := srcfile.TransList(`C:\git\transtool\init.txt`)
	map1, _ := srcfile.MakeMap(list, list2)
	for _, m := range map1 {
		println(m.Org, m.Trans)
	}
}
