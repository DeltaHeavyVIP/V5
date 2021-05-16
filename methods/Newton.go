package methods

import (
	"../draw"
	. "../entry"
	"fmt"
	_ "gonum.org/v1/plot"
	"os"
)

type Newton struct{}

func (Newton) Count(data *NewData) {
	answer := data.ArrayY[0]
	zn := float32(1.0)
	for i := 1; i < data.N; i++ {
		for j := 0; j < i; j++ {
			zn *= data.X - data.ArrayX[j]
		}
		answer += getF(data, i+1, i) * zn
		zn = 1.0
	}
	_, _ = fmt.Fprintf(os.Stdout, "Вычесленный Y методом Ньютона для неравноотстоющих узлов: %f \n", answer)
	draw.Graph(data,0)
}

func getF(data *NewData, k int, max int) float32 {
	if k == 2 {
		 return (data.ArrayY[max] - data.ArrayY[max-1]) / (data.ArrayX[max] - data.ArrayX[max-1])
	} else {
		return (getF(data,k-1,max)-getF(data,2,max-1))/(data.ArrayX[max] - data.ArrayX[max - k + 1])
	}
}
