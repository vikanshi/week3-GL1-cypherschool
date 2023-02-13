package db

import "log"

var Key int

func GetKey() {
	Key = 9
}

type Val struct {
	i int
}

// value receiver
// func (v Val) GetVal() int {
// 	v.i = 10
// 	return v.i
// }

// pointer receiver
func (v *Val) GetVal() int {
	return v.i
}

//pointer receiver can be called in case of normal variable as well

func get() {
	vv := Val{}
	vv.GetVal()
	log.Println(vv.i)

	ww := &Val{}
	ww.GetVal()
	log.Println(ww.i)
}

//pointer receiver and value receiver: used in context of structure and method
//pass by reference and pass by value: used in context of function
