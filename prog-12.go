package main

import (
	"errors"
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimetre() float64

	// check if true, return 1
	// otherwise, return 0
	myfunc(bool) (int, error)
}

type circle struct {
	r float64
}

func (c circle) area() float64 {
	return math.Pi * c.r * c.r
}

func (c circle) perimetre() float64 {
	return 2 * math.Pi * c.r
}

type rect struct {
	w, h float64
}

func (r rect) area() float64 {
	return r.w * r.h
}

func (r rect) perimetre() float64 {
	return (2 * r.w) + (2 * r.h)
}

func (r rect) myfunc(flag bool) (int, error) {
	if flag {
		return 1, nil
	}
	return 0, errors.New("Random error")
}

func getPerimetreWierdly(g geometry) float64 {
	_, err := g.myfunc(true)
	if err == nil {
		return g.perimetre()
	}
	return 0	
}

func main() {
	c1 := circle{4}
	r1 := rect{4, 3}

	fmt.Println("Area:", c1.area())
	fmt.Println("Perimitre:", c1.perimetre())
	fmt.Println("getPerimitre:", getPerimetreWierdly(r1))

}
