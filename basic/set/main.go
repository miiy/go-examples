package main

import "fmt"

func main()  {
	set := make(map[string]struct{}, 0)
	set["foo"] = struct{}{}
	for k := range set {
		fmt.Println(k)
	}
	delete(set, "foo")
	size := len(set)
	exists := set["foo"]
	fmt.Println(size)
	fmt.Println(exists)
}
