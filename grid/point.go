package grid

import (
	"Gman/configs"
	"fmt"
	"math"
	"strconv"
)

type Point struct {
	X int
	Y int
}

func CreatePoint(x string, y string) (Point, error){
	X, err1 := convertStringToInt(x)
	Y, err2 := convertStringToInt(y)

	if err1 != nil || err2 != nil {
		return Point{}, fmt.Errorf("invalid point values, (%s, %s)",x, y)
	}
	// checking whether they come inside grid size or not
	if !IsCoordinatesValid(X, Y, configs.GameConfigInstance.GridSize) {
		return Point{}, fmt.Errorf("the given coordinates (%s, %s) do not come within the grid",x, y)
	}
	return Point{X, Y}, nil
}	


func convertStringToInt(arg string) (int, error) {
	
		num, err := strconv.Atoi(arg)
		if err != nil {
			return 0, fmt.Errorf("invalid argument, could not convert str to int")
		}
	
	return num, nil
}

func IsCoordinatesValid(x int, y int, gridSize int)bool {
	if x > gridSize || x < 0{
		return false
	} else if  y > gridSize || y < 0{
		return false
	}
	return true
}

// gives possible values (0, 1 , -1) used for mapping out the relative direction between two points
func (p *Point)GetNormalizedDifference(other Point) Point {
	dx := p.X - other.X
	dy := p.Y - other.Y

	if dx != 0 {
		dx = dx / int(math.Abs(float64(dx)))
	} else {
		dx = 0
	}
	if dy != 0 {
		dy = dy / int(math.Abs(float64(dy)))
	} else {
		dy = 0
	}
	return Point{dx, dy}
}