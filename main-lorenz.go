package main

// Reference: https://duckduckgo.com/?q=DuckDuckGo+AI+Chat&ia=chat&duckai=1&atb=v268-1
//            Previous Duck.ai qurery "Write the shortest possible Lorenz Attractor program in Python. Do not output graphics, only output the x,y,z co-ordinates required to plot."
//            https://github.com/gonum/plot/issues/494
//            https://deepwiki.com/gonum/plot/1.1-installation-and-getting-started
//            https://pkg.go.dev/gonum.org/v1/plot/plotter#XYRange
// as of 2025-08-22
// V1.0 2025-08-24: First Customer Ship (FCS)

import (
	"fmt"
	"image/color"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

type GraphingCoordinates struct {
	X float64
	Y float64
	Z float64
}

func main() {

	/********/
	/* Init */
	/********/

	const sigma = 10.0
	const beta = 8 / 3
	const rho = 28.0
	const dt = 0.01
	//
	const steps = 1000

	var dx, dy, dz float64 = 0.0, 0.0, 0.0 // Initial conditions
	var i int = 0                          // loop itterator only
	var k = GraphingCoordinates{}          // Loop itteration only

	current_coordinates := GraphingCoordinates{
		// Really wish I could inialize these in the struct itself!  It may not be necessary (like start plotting somewhere else rather than 1.0),  but I'd like to have that option!
		X: 1.0,
		Y: 1.0,
		Z: 1.0,
	}
	var CoordinateList []GraphingCoordinates

	/**************/
	/* Main Block */
	/**************/

	for i = 0; i < steps; i++ {
		dx = sigma * (current_coordinates.Y - current_coordinates.X) * dt
		dy = (current_coordinates.X*(rho-current_coordinates.Z) - current_coordinates.Y) * dt
		dz = (current_coordinates.X*current_coordinates.Y - beta*current_coordinates.Z) * dt
		current_coordinates.X += dx
		current_coordinates.Y += dy
		current_coordinates.Z += dz

		//fmt.Printf("%v , %v , %v\n", current_coordinates.X, current_coordinates.Y, current_coordinates.Z)
		CoordinateList = append(CoordinateList, current_coordinates)
	}

	p := plot.New()
	p.Title.Text = "My First Plot"
	p.X.Label.Text = "X axis"
	p.Y.Label.Text = "Y axis"
	points := plotter.XYZs{}
	//
	for _, k = range CoordinateList {
		fmt.Printf("%v , %v , %v\n", k.X, k.Y, k.Z)                       // Print to STDOUT as well.
		points = append(points, struct{ X, Y, Z float64 }{k.X, k.Y, k.Z}) // https://github.com/gonum/plot/issues/494
	}
	//
	line, err := plotter.NewLine(points)
	if err != nil {
		// Handle error
	}
	//
	// Customize line appearance
	line.LineStyle.Width = vg.Points(1.5)
	line.LineStyle.Color = color.RGBA{R: 0, G: 0, B: 255, A: 255}
	//
	// Add the line to the plot
	p.Add(line)
	//
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "my_plot.png"); err != nil {
		fmt.Print("Could not write png file: ", err)
	}

	os.Exit(0) // Because I always prefer to return properly with a code of my choosing...
}
