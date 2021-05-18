package methods

import (
	"../draw"
	. "../entry"
	"fmt"
	"os"
)

type Lagrange struct{
	Data *NewData
}

func (l *Lagrange) Count() {
	data := l.Data
	var up float32 = 1
	var down float32 = 1
	var res float32 = 0
	for i := 0; i < data.N; i++ {
		for j := 0; j < data.N; j++ {
			if i != j {
				up *= data.X - data.ArrayX[j]
				down *= data.ArrayX[i] - data.ArrayX[j]
			}
		}
		res +=data.ArrayY[i] * up / down
		up,down = 1,1
	}
	_, _ = fmt.Fprintf(os.Stdout, "Вычесленный Y методом Лагранжа: %f \n", res)
	draw.Graph(data,1)
}
