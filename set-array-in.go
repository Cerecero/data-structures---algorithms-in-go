package main

import "fmt"

type Set map[interface{}]struct{}

func NewSet() Set{
	return make(Set)
}

func (s Set) add(elem interface{}) {
	s[elem] = struct{}{}
}
func (s Set) remove(elem interface{}) {
	delete(s, elem)
}
func (s Set) has(elem interface{}) bool{
	_, ok := s[elem]
	return ok
}

func (s Set) String() string {
	keys := make([]interface{}, 0, len(s))
	for k := range s {
		keys = append(keys, k)
	}

	return fmt.Sprintf("%v", keys)
}

func main(){  
	setArr := NewSet()

	setArr.add("Test")
	setArr.add("Another test")
	setArr.add(23)

	fmt.Printf("This is set: %v\n", setArr)
	setArr.remove("Another test")
	fmt.Printf("This is set: %v\n", setArr)
	fmt.Println(setArr.has(23))
	fmt.Println(setArr.has("Tes"))

}
