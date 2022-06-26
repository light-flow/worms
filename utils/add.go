package utils

type Point struct {
	X int8
	Y int8
}

func Add(a int8, b int8) int8 {
	return a + b
}

func (point Point) GetDistance(point1 Point, point2 Point) int8 {
	return (point1.X - point2.X) * (point1.Y - point2.Y)
}
