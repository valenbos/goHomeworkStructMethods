package main

import "fmt"


type Rectangle struct {
	Width int64
	Lenght int64
}


func (r Rectangle) surfaceArea() int64 {
	return r.Width * r.Lenght
} 


func (r Rectangle) perimeter() int64 {
	return 2 * (r.Width + r.Lenght)
} 


func main() {
	rectangle := Rectangle{
		Width: 10,
		Lenght: 30,
	}

	rectSurface := rectangle.surfaceArea()
	rectPerimeter := rectangle.perimeter()

	fmt.Printf("Surface area of the rectangle: %vcm.\n", rectSurface)
	fmt.Printf("Perimeter of the rectangle: %vcm.", rectPerimeter)

}
