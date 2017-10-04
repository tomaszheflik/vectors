package vectors_test

import (
	"testing"
	v "github.com/tomaszheflik/vectors"
)

func TestGetVectorFirst(t *testing.T) {
	start := v.Point{2, 2}
	stop := v.Point{3, 1}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if !(testVector.Coordinates.X > 0) || !(testVector.Coordinates.Y > 0) {
		t.Error(
			"For TOP expected X>0 and Y>0 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}

func TestGetVectorSecond(t *testing.T) {
	start := v.Point{2, 2}
	stop := v.Point{3, 3}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if !(testVector.Coordinates.X > 0) || !(testVector.Coordinates.Y < 0) {
		t.Error(
			"For TOP expected X>0 and Y>0 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}

func TestGetVectorThird(t *testing.T) {
	start := v.Point{2, 2}
	stop := v.Point{1, 3}
	testVector,err := v.GetVector(start, stop)
	if err != nil {
		t.Error("Zero length Vector passed")
	}
	if !(testVector.Coordinates.X > 0) || !(testVector.Coordinates.Y > 0) {
		t.Error(
			"For TOP expected X>0 and Y>0 got X:", testVector.Coordinates.X,"Y:", testVector.Coordinates.Y,"\n",
		)
	}
}


func TestVectorLenght(t *testing.T) {
	expected := 7.0710678118654755
	vector := v.Vector{}
	vector.Coordinates.X = 5
	vector.Coordinates.Y = 5
	lenght := v.GetLenght(vector)

	if lenght != expected {
		t.Error(
			"Expected lenght = ",expected," got lenhght =",lenght,"\n",
		)
	}
}