package draw

import (
	. "../entry"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
	"math/rand"
)

func Lagrange(data *NewData){
	scatterData := point(data)
	lineData := line(data)

	p := plot.New()
	p.Title.Text = "Lagrange"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	p.Add(plotter.NewGrid())

	// Make a scatter plotter and set its style.
	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 7, G: 28, B: 112}

	// Make a line plotter and set its style.
	l, err := plotter.NewLine(lineData)
	if err != nil {
		panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 165, G: 116, B: 0}

	// entry for each
	p.Add(s, l)
	p.Legend.Add("points", s)
	p.Legend.Add("function", l)

	// Save the plot to a PNG file.
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "lagrange.png"); err != nil {
		panic(err)
	}
}

func point(data *NewData) plotter.XYs{
	pts := make(plotter.XYs, data.N)
	for i:=0;i<data.N;i++ {
		pts[i].X = float64(data.ArrayX[i])
		pts[i].Y = float64(data.ArrayY[i])
	}
	return pts
}

func line(data *NewData) plotter.XYs {
	var h float32 = 0.01
	var l int = 1
	for k:=data.ArrayX[0];k < data.ArrayX[data.N-1];k+=h {
		l+=1
	}
	pts := make(plotter.XYs, l)

	var up,down float32 = 1,1
	var res float32 = 0
	for k:=0;k < l;k++ {
		for i := 0; i < data.N; i++ {
			for j := 0; j < data.N; j++ {
				if i != j {
					up *= data.ArrayX[0] - data.ArrayX[j] + float32(k) * h
					down *= data.ArrayX[i] - data.ArrayX[j]
				}
			}
			res += data.ArrayY[i] * up / down
			up,down = 1,1
		}
		pts[k].X = float64(data.ArrayX[0]+float32(k) * h)
		pts[k].Y = float64(res)
		res=0
	}
	return pts
}

func Newton(data *NewData){
	rand.Seed(int64(0))
	scatterData := point(data)
	lineData := line(data)

	// Create a new plot, set its title and
	// axis labels.
	p := plot.New()

	p.Title.Text = "Lagrange"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"
	// Draw a grid behind the data
	p.Add(plotter.NewGrid())

	// Make a scatter plotter and set its style.
	s, err := plotter.NewScatter(scatterData)
	if err != nil {
		panic(err)
	}
	s.GlyphStyle.Color = color.RGBA{R: 7, G: 28, B: 112}

	// Make a line plotter and set its style.
	l, err := plotter.NewLine(lineData)
	if err != nil {
		panic(err)
	}
	l.LineStyle.Width = vg.Points(1)
	l.LineStyle.Color = color.RGBA{R: 165, G: 116, B: 0}

	// entry for each
	p.Add(s, l)
	p.Legend.Add("points", s)
	p.Legend.Add("function", l)

	// Save the plot to a PNG file.
	if err := p.Save(6*vg.Inch, 6*vg.Inch, "lagrange.png"); err != nil {
		panic(err)
	}
}