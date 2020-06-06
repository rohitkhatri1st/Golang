package main

import "fmt"

type vehicle struct {
	doors int
	color string
}
type truck struct {
	is        vehicle
	fourWheel bool
}
type sedan struct {
	vehicle
	luxury bool
}

func main() {
	t1 := truck{
		is: vehicle{
			doors: 2,
			color: "Red",
		},
		fourWheel: true,
	}
	s1 := sedan{
		vehicle: vehicle{
			doors: 4,
			color: "cream",
		},
		luxury: true,
	}
	fmt.Println(t1)
	fmt.Println(s1)
	fmt.Println(t1.is.doors, t1.is.color, t1.fourWheel)
	fmt.Println(s1.doors, s1.color, s1.luxury)
}
