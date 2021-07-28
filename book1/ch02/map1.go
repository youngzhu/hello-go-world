package main

import "fmt"

// 一个包含个人详细信息的类型
type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {
	var personDB map[string]PersonInfo
	personDB = make(map[string]PersonInfo)

	// 往map里加几条数据
	personDB["1234"] = PersonInfo{"1234", "Tom", "Room 203, ..."}
	personDB["1"] = PersonInfo{"1", "Jack", "Room 2101, ..."}

	// 从map中查找“1234”
	person, ok := personDB["1234"]

	// ok 返回的bool类型，true表示找到了对应数据
	if ok {
		fmt.Println("Found person", person.Name, "whit ID 1234.")
	} else {
		fmt.Println("Did not find person with ID 1234.")
	}
}
