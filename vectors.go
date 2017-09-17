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

func GetLenght(v Vector) int64 {
	lenght := math.Sqrt(float64(v.Coordinates.X*v.Coordinates.X)+float64(v.Coordinates.Y*v.Coordinates.Y))
	return int64(lenght)
}

func Normalize(v Vector) (Vector, error) {
	var vector Vector
	vector.Coordinates.X = 10
	vector.Coordinates.Y = 10
	return vector, nil
}

func GetAngel(v1 Vector, v2 Vector)  (float64, float64) {
	product := GetProd(v1, v2)
	lenght1 := GetLenght(v1)
	lenght2 := GetLenght(v2)

	angRad := math.Acos(float64(product)/float64(lenght1)*float64(lenght2))
	angDeg := (180/math.Pi)*angRad
	return angRad, angDeg
}

func GetProd(v1 Vector, v2 Vector)  int64 {
	prod := int64(v1.Coordinates.X * v2.Coordinates.X) + int64(v1.Coordinates.Y * v2.Coordinates.Y)
	return prod
}