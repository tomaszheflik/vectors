package vectors

import (
	"errors"
	"math"
)
func GetVector(start Point, stop Point) (Vector, error) {
	point :=  Vector{}
	// Top orientation
	if stop.X == start.X && stop.Y < start.Y {
		point.Coordinates.X = stop.X - start.X
		point.Coordinates.Y = start.Y - stop.Y
	}
	// Top-Right orientation
	if stop.X > start.X && stop.Y < start.Y {
		point.Coordinates.X = stop.X - start.X
		point.Coordinates.Y = (stop.Y - start.Y)*-1
	}
	// Right orientation
	if stop.X > start.X && stop.Y == start.Y {
		point.Coordinates.X = stop.X - start.X
		point.Coordinates.Y = stop.Y - start.Y
	}
	// Bottom Right
	if stop.X > start.X && stop.Y > start.Y {
		point.Coordinates.X = stop.X - start.X
		point.Coordinates.Y = (stop.Y - start.Y)*-1
	}
	// Bottom orientation
	if stop.X == start.X && stop.Y > start.Y {
		point.Coordinates.X = stop.X - start.X
		point.Coordinates.Y = (stop.Y - start.Y)*-1
	}
	// Bottom Left
	if stop.X < start.X && stop.Y > start.Y {
		point.Coordinates.X = stop.X - start.X
		point.Coordinates.Y = (stop.Y - start.Y)*-1
	}
	// Left orientation
	if stop.X < start.X && stop.Y == start.Y {
		point.Coordinates.X = (stop.X - start.X)
		point.Coordinates.Y = stop.Y - start.Y
	}
	// Left top orientation
	if stop.X < start.X && stop.Y < start.Y {
		point.Coordinates.X = stop.X - start.X
		point.Coordinates.Y = (stop.Y - start.Y)*-1
	}
	if stop.X == 0 && start.Y ==0 && stop.X == 0 && stop.Y == 0 {
		return point, errors.New("Zero lenght vectors")
	}
	return point, nil
}

func VectorLenght(v Vector) (int64, error) {
	lenght := math.Sqrt(float64(v.Coordinates.X*v.Coordinates.X)+float64(v.Coordinates.Y*v.Coordinates.Y))
	return int64(lenght),nil
}

func Normalize(v Vector) (Vector, error) {
	var vector Vector
	vector.Coordinates.X = 10
	vector.Coordinates.Y = 10
	return vector, nil
}

func GetAngel(v1 Vector, v2 Vector) (int, error) {
	return 20, nil
}
