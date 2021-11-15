package main

import "fmt"

type Point struct{ X, Y int }
type Points []*Point

func (ps Points) ToString() string {
	str := ""

	for _, p := range ps {
		if p == nil {
			str += "<nil>"
			continue
		}

		if str == "" {
			str += fmt.Sprintf("[%d, %d]", p.X, p.Y)
			continue
		}

		str += ", "
		str += fmt.Sprintf("[%d, %d]", p.X, p.Y)

	}

	return str

}

func createPoints() {
	ps := Points{}
	fmt.Printf("ps = %v\n", ps)
	ps = append(ps, &Point{1, 2})
	ps = append(ps, nil)
	ps = append(ps, &Point{3, 5})

	fmt.Println(ps.ToString())

}

func main() {
	createPoints()

}
