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

func GetLenght(v Vector) float64 {
	lenght := math.Sqrt(float64(v.Coordinates.X)*float64(v.Coordinates.X)+float64(v.Coordinates.Y)*float64(v.Coordinates.Y))
	return lenght
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
	angRad := math.Acos(product/(lenght1*lenght2))
	angDeg := (180/math.Pi)*angRad
	return angRad, angDeg
}

func GetProd(v1 Vector, v2 Vector) float64 {
	product := (v1.Coordinates.X * v2.Coordinates.X) + (v1.Coordinates.Y * v2.Coordinates.Y)
	return float64(product)
}

func Round(x, unit float64) float64 {
	return float64(int64(x/unit+0.5)) * unit
}