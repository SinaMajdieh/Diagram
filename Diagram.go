package diagram

import "fmt"

const (
	max_Yaxis = 90
	max_Xaxis = 150
	point     = "*"
	axis      = "O "
	accuracy  = 1
)

type chart struct {
	Yaxis                [max_Yaxis]float64
	Xaxis                [max_Xaxis]float64
	Yaxis_end            float64
	Xaxis_end            float64
	diagram              [][]string
	diagram_Yaxis_length int
	diagram_Xaxis_length int
}

func (c *chart) make( /*accuracy float64 ,*/ max_Yaxis_number float64, max_Xaxis_number float64) {
	c.Yaxis_end = max_Yaxis_number
	c.Xaxis_end = max_Xaxis_number
	c.diagram_Yaxis_length = 1
	c.diagram_Xaxis_length = 1
	for i := 1; i < max_Yaxis && c.Yaxis[i-1] != max_Yaxis_number; i++ {
		//fmt.Println(c.Yaxis[i-1])
		c.Yaxis[i] = c.Yaxis[i-1] + accuracy
		c.diagram_Yaxis_length++
	}
	for i := 1; i < max_Xaxis && c.Xaxis[i-1] != max_Xaxis_number; i++ {
		c.Xaxis[i] = c.Xaxis[i-1] + accuracy
		c.diagram_Xaxis_length++
	}
	c.diagram = make([][]string, c.diagram_Yaxis_length)
	for i := 0; i < c.diagram_Yaxis_length; i++ {
		c.diagram[i] = make([]string, c.diagram_Xaxis_length)
		for j := 0; j < c.diagram_Xaxis_length; j++ {
			c.diagram[i][j] = " "
		}
	}
}

func (c *chart) findpoint(y, x float64) (Y int, X int) {
	for i := 0; i < c.diagram_Yaxis_length; i++ {
		if c.Yaxis[i] == y {
			Y = i
			break
		}
	}
	for i := 0; i < c.diagram_Xaxis_length; i++ {
		if c.Xaxis[i] == x {
			X = i
			break
		}

	}
	return
}

//Adding a point to Diagram
func (c *chart) AddValue(x float64, y float64) {
	var Yindex, Xindex int
	if y <= c.Yaxis_end && x <= c.Xaxis_end {
		Yindex, Xindex = c.findpoint(y, x)
	} else {
		fmt.Println("not available")
		return
	}
	c.diagram[Yindex][Xindex] = point

}

//Printing Diagram
func (c chart) Show() {
	for i := c.diagram_Yaxis_length - 1; i >= 0; i-- {
		fmt.Print(axis)
		for j := 0; j < c.diagram_Xaxis_length; j++ {
			fmt.Print(c.diagram[i][j] + " ")
		}
		fmt.Println()
	}
	fmt.Print("  ")
	for i := 1; i <= c.diagram_Xaxis_length; i++ {
		fmt.Print(axis)
	}
	fmt.Println()
}

//Make an emty Diagram
func MakeDiagram( /*accuracy float64 ,*/ max_Yaxis_number float64, max_Xaxis_number float64) chart {
	c := chart{}
	c.make( /*accuracy,*/ max_Yaxis_number, max_Xaxis_number)
	return c
}
