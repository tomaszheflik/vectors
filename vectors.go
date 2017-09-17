package vectors

import (
	"fmt"
)
func GetVector(start Point, stop Point) (Point, error) {
	fmt.Printf("Calulate vector from two points")
	var point Point
	point.X = 10
	point.Y = 10
	return point, nil
}

func Normalize(v Vector) (Vector, error) {
	var vector Vector
	vector.coordinates.X = 10
	vector.coordinates.Y = 10
	return vector, nil
}

func getAngel(v1 Vector, v2 Vector) (int, error) {
	return 20, nil
}
