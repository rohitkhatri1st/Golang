package main

func main() {
	a := (42 == 42)	//'==' , '<=' , '>=' , '!=' , '<' , '>' are few ofthe operators
	b := (42 <= 28)
	c := (42 >= 42)
	d := (42 != 42)	//!= is "not equal to" operator, it checks if 42 is not equal to 42
	e := (42 < 63)
	f := (42 > 63)
	println(a, b, c, d, e, f)
}
