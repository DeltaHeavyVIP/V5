package draw

import (
	. "../entry"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
	"image/color"
)

func Graph(data *NewData, name int) {
	scatterData := point(data)
	var lineData plotter.XYs
	if name == 1 {
		lineData = lineLag(data)
	} else {
		lineData = lineNewt(data)
	}

	p := plot.New()
	if name == 1 {
		p.Title.Text = "Lagrange"
	} else {
		p.Title.Text = "Newton"
	}
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
	if name == 1 {
		if err := p.Save(6*vg.Inch, 6*vg.Inch, "lagrange.png"); err != nil {
			panic(err)
		}
	} else {
		if err := p.Save(6*vg.Inch, 6*vg.Inch, "newton.png"); err != nil {
			panic(err)
		}
	}
}

func point(data *NewData) plotter.XYs {
	pts := make(plotter.XYs, data.N)
	for i := 0; i < data.N; i++ {
		pts[i].X = float64(data.ArrayX[i])
		pts[i].Y = float64(data.ArrayY[i])
	}
	return pts
}

func lineLag(data *NewData) plotter.XYs {
	var h float32 = 0.01
	var l int = 1
	for k := data.ArrayX[0]; k < data.ArrayX[data.N-1]; k += h {
		l += 1
	}
	pts := make(plotter.XYs, l)

	var up, down float32 = 1, 1
	var res float32 = 0
	for k := 0; k < l; k++ {
		for i := 0; i < data.N; i++ {
			for j := 0; j < data.N; j++ {
				if i != j {
					up *= data.ArrayX[0] - data.ArrayX[j] + float32(k)*h
					down *= data.ArrayX[i] - data.ArrayX[j]
				}
			}
			res += data.ArrayY[i] * up / down
			up, down = 1, 1
		}
		pts[k].X = float64(data.ArrayX[0] + float32(k)*h)
		pts[k].Y = float64(res)
		res = 0
	}
	return pts
}

/*func lineNewt(data *NewData) plotter.XYs {
	var h float32 = 0.01
	var l = 1
	for k := data.ArrayX[0]; k < data.ArrayX[data.N-1]; k += h {
		l += 1
	}
	pts := make(plotter.XYs, l)

	var res float32 = data.ArrayY[0]
	for k := 0; k < l; k++ {
		res = bitch(data,data.ArrayX[0] + float32(k)*h)
		pts[k].X = float64(data.ArrayX[0] + float32(k)*h)
		pts[k].Y = float64(res)
		//fmt.Println("----",data.ArrayX[0] + float32(k)*h,"\t\t",res)
	}
	return pts
}

func bitch(data *NewData, x float32) float32 {
	answer := data.ArrayY[0]
	zn := float32(1.0)
	for i := 1; i < data.N; i++ {
		for j := 0; j < i; j++ {
			zn *= x - data.ArrayX[j]
		}
		answer += getF(data, i+1, i) * zn
		zn = 1.0
	}
	return answer
}
func getF(data *NewData, k int, max int) float32 {
	if k == 2 {
		return (data.ArrayY[max] - data.ArrayY[max-1]) / (data.ArrayX[max] - data.ArrayX[max-1])
	} else {
		return (getF(data,k-1,max)-getF(data,2,max-1))/(data.ArrayX[max] - data.ArrayX[max - k + 1])
	}
}*/

//Я писал код ниже и разбирался как работает метод Ньютона для равноотстоящих узлов часов 6, а мог просто вызывать функцию из класса Newton, господи, какой я тупой, просто ебнешься
func lineNewt(data *NewData) plotter.XYs {
	var h float32 = 0.01
	var l = 1
	for k := data.ArrayX[0]; k < data.ArrayX[data.N-1]; k += h {
		l += 1
	}
	pts := make(plotter.XYs, l)

	//fmt.Println(getAns(3,data.ArrayY))
	var up float32 = 1
	var res float32 = data.ArrayY[0]
	for k := 0; k < l; k++ {
		for i := 1; i < data.N; i++ {
			up *= data.ArrayX[0] + float32(k)*h - data.ArrayX[i-1]
			res += getAns(i, data.ArrayY) * up / float32(fuctorial(i))
		}
		pts[k].X = float64(data.ArrayX[0] + float32(k)*h)
		pts[k].Y = float64(res)
		res = data.ArrayY[0]
		up = 1
	}
	return pts
}
func getAns(up int, array []float32) float32 {
	res := float32(array[up])
	k := 1
	if up%2 == 1 {
		for i := up - 1; i >= 0; i-- {
			k = k * (i + 1) / (up - i)
			if i%2 == 1 {
				res += array[i] * float32(k)
			} else {
				res -= array[i] * float32(k)
			}
		}
	} else {
		for i := up - 1; i >= 0; i-- {
			k = k * (i + 1) / (up - i)
			if i%2 == 0 {
				res += array[i] * float32(k)
			} else {
				res -= array[i] * float32(k)
			}
		}
	}
	return res
}

func fuctorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
	}
	return res
}