package vectors_test

import (
	"testing"
	v "github.com/tomaszheflik/vectors"
)

func TestGetVectorTop(t *testing.T) {
	start := v.Point{100, 200}
	stop := v.Point{100, 100}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if testVector.Coordinates.X != 0 || testVector.Coordinates.Y != 100 {
		t.Error(
			"For TOP expected X=0 and Y=100 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}

func TestGetVectorTopRight(t *testing.T) {
	start := v.Point{100, 100}
	stop := v.Point{150, 50}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if testVector.Coordinates.X != 50 || testVector.Coordinates.Y != 50 {
		t.Error(
			"For TOP expected X=50 and Y=50 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}

func TestGetVectorRight(t *testing.T) {
	start := v.Point{100, 100}
	stop := v.Point{200, 100}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if testVector.Coordinates.X != 100 || testVector.Coordinates.Y != 0 {
		t.Error(
			"For TOP expected X=0 and Y=100 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}

func TestGetVectorBottomRight(t *testing.T) {
	start := v.Point{100, 100}
	stop := v.Point{150, 150}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if testVector.Coordinates.X != 50 || testVector.Coordinates.Y != -50 {
		t.Error(
			"For TOP expected X=0 and Y=100 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}

func TestGetVectorBottom(t *testing.T) {
	start := v.Point{100, 100}
	stop := v.Point{100, 200}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if testVector.Coordinates.X != 0 || testVector.Coordinates.Y != -100 {
		t.Error(
			"For TOP expected X=0 and Y=100 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}
func TestGetVectorBottomLeft(t *testing.T) {
	start := v.Point{100, 100}
	stop := v.Point{50, 150}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if testVector.Coordinates.X != -50 || testVector.Coordinates.Y != -50 {
		t.Error(
			"For TOP expected X=0 and Y=100 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}
func TestGetVectorLeft(t *testing.T) {
	start := v.Point{100, 100}
	stop := v.Point{50, 100}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if testVector.Coordinates.X != -50 || testVector.Coordinates.Y != 0 {
		t.Error(
			"For TOP expected X=0 and Y=100 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}
func TestGetVectorTopLeft(t *testing.T) {
	start := v.Point{100, 100}
	stop := v.Point{50, 50}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if testVector.Coordinates.X != -50 || testVector.Coordinates.Y != 50 {
		t.Error(
			"For TOP expected X=0 and Y=100 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}

func TestVectorLenght(t *testing.T) {
	expected := int64(7)
	vector := v.Vector{}
	vector.Coordinates.X = 5
	vector.Coordinates.Y = 5
	lenght := v.VectorLenght(vector)

	if lenght != expected {
		t.Error(
			"Expected lenght = ",expected," got lenhght =",lenght,"\n",
		)
	}
}