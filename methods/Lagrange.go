package methods

import (
	"../draw"
	. "../entry"
	"fmt"
	"os"
)

type Lagrange struct{}

func (Lagrange) Count(data *NewData) {
	var up float32 = 1
	var down float32 = 1
	var res float32 = 0
	for i := 0; i < data.N; i++ {
		for j := 0; j < data.N; j++ {
			if i != j {
				up *= data.X - data.ArrayX[i]
				down *= data.ArrayX[i] - data.ArrayX[j]
			}
		}
		res += data.ArrayY[i] * up / down
	}
	_, _ = fmt.Fprintf(os.Stdout, "Вычесленный Y методом Лагранжа: %d", res)
	draw.DrawLagrange(data)
}
