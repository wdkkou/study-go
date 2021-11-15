package main

import (
	"fmt"
	"math"
)

type Point struct{ X, Y int }
type FloatPoint struct{ X, Y float64 }

func (p *Point) Render() {
	fmt.Printf("<%d, %d>\n", p.X, p.Y)

}

func callRender() {
	p := &Point{X: 5, Y: 12}
	p.Render()

}

func (p *Point) Distance(dp *Point) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y

	return math.Sqrt(float64(x*x + y*y))

}

func (p *FloatPoint) Distance(dp *FloatPoint) float64 {
	x, y := p.X-dp.X, p.Y-dp.Y

	return math.Sqrt(float64(x*x + y*y))

}

func callDistance() {
	p := &Point{X: 0, Y: 0}
	fmt.Println(p.Distance(&Point{X: 1, Y: 1}))

	fp := &FloatPoint{X: 4.3, Y: 4.3}
	fmt.Println(fp.Distance(&FloatPoint{X: 2.2, Y: 1.1}))
}

type MyInt int

func (m MyInt) Plus(i int) int {
	return int(m) + i
}

func callPlus() {
	ans := MyInt(4).Plus(3)
	fmt.Println(ans)
}

type IntPair [2]int

func (ip IntPair) First() int {
	return ip[0]
}

func (ip IntPair) Last() int {
	return ip[1]
}

type Strings []string

func (s Strings) Join(d string) string {
	sum := ""
	for _, v := range s {
		if sum == "" {
			sum += v
			continue
		}
		sum += d
		sum += v
	}
	return sum

}

func callJoin() {
	ip := IntPair{1, 2}
	fmt.Println(ip.First())
	fmt.Println(ip.Last())

	str := Strings{"A", "B", "C"}.Join(", ")
	fmt.Println(str)

}

func main() {
	// callRender()
	// callDistance()
	// callPlus()
	callJoin()

}
