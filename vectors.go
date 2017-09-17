package vectors

import (
	"fmt"
	"errors"
)
func GetVector(start Point, stop Point) (Vector, error) {
	fmt.Printf("Calulate vector from two points\n")
	point :=  Vector{}
	// Top orientation
	if stop.X == start.X && stop.Y < start.Y {
		point.coordinates.X = stop.X - start.X
		point.coordinates.Y = start.Y - stop.Y
	}
	// Top-Right orientation
	if stop.X > start.X && stop.Y < start.Y {
		point.coordinates.X = stop.X - start.X
		point.coordinates.Y = (stop.Y - start.Y)*-1
	}
	// Right orientation
	if stop.X > start.X && stop.Y == start.Y {
		point.coordinates.X = stop.X - start.X
		point.coordinates.Y = stop.Y - start.Y
	}
	// Bottom Right
	if stop.X > start.X && stop.Y > start.Y {
		point.coordinates.X = stop.X - start.X
		point.coordinates.Y = (stop.Y - start.Y)*-1
	}
	// Bottom orientation
	if stop.X == start.X && stop.Y > start.Y {
		point.coordinates.X = stop.X - start.X
		point.coordinates.Y = (stop.Y - start.Y)*-1
	}
	// Bottom Left
	if stop.X < start.X && stop.Y > start.Y {
		point.coordinates.X = stop.X - start.X
		point.coordinates.Y = (stop.Y - start.Y)*-1
	}
	// Left orientation
	if stop.X < start.X && stop.Y == start.Y {
		point.coordinates.X = (stop.X - start.X)
		point.coordinates.Y = stop.Y - start.Y
	}
	// Left top orientation
	if stop.X < start.X && stop.Y < start.Y {
		point.coordinates.X = stop.X - start.X
		point.coordinates.Y = (stop.Y - start.Y)*-1
	}
	if stop.X == 0 && start.Y ==0 && stop.X == 0 && stop.Y == 0 {
		return point, errors.New("Zero lenght vectors")
	}
	return point, nil
}

func VectorLenght(v Vector) (int, error) {
	return 10,nil
}

func Normalize(v Vector) (Vector, error) {
	var vector Vector
	vector.coordinates.X = 10
	vector.coordinates.Y = 10
	return vector, nil
}

func GetAngel(v1 Vector, v2 Vector) (int, error) {
	return 20, nil
}
