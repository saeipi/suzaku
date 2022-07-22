package main

import (
	"fmt"
	"regexp"
)

func main() {
	// 1 中⽂括号
	str := "华南地区（⼴州）"
	regex := regexp.MustCompile("（(.*?)）")
	list := regex.FindStringSubmatch(str)
	fmt.Println(list)

	// 2 英⽂括号
	str = "华南地区(⼴州)"
	regex = regexp.MustCompile("\\((.*?)\\)")
	list = regex.FindStringSubmatch(str)
	fmt.Println(list)

	// 3 截取取值字段 FindAllStringSubmatch
	str = "$main.countDistinct($city)\n$main.countDistinct($user_id)\n$main.countDistinct($age)"
	regex = regexp.MustCompile("countDistinct\\(\\$(.*?)\\)")
	list1 := regex.FindAllStringSubmatch(str,-1)
	fmt.Println(list1)

	// 3 截取取值字段 FindAllString
	str = "$main.countDistinct($city)\n$main.countDistinct($user_id)\n$main.countDistinct($age)"
	regex = regexp.MustCompile("countDistinct\\(\\$(.*?)\\)")
	list = regex.FindAllString(str,-1)
	fmt.Println(list)
}

// (?<=var hash = ")([\s\S]*?)(?=") || (?<=Begin)([.\S\s]*)(?=End)
// (?<=countDistinct\(\$)([\s\S]*?)(?=\))