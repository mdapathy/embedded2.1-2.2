package main

import (
	"github.com/Equanox/gotron"
	"github.com/wcharczuk/go-chart"
	"math"
	"math/rand"
	"os"
	"time"
)

func main() {

	n := 6
	w := 1500
	N := 1024

	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	var y = make([]float64, N, N)
	var x = make([]float64, N, N)

	var Freal = make([]float64, N, N)
	var Fimag = make([]float64, N, N)
	var F = make([]float64, N, N)

	var sum float64
	var dispersion float64

	for i := 0; i < N; i++ {

		var ytemp float64
		for j := 0; j < n; j++ {
			ytemp += r.Float64() *
				(math.Sin(float64((w/n)*(j+1)*i) + r.Float64()))
		}
		sum += ytemp

		y[i] = ytemp
		x[i] = float64(i)

	}

	expect := sum / float64(N)

	for i := 0; i < len(y); i++ {
		dispersion += math.Pow(y[i]-expect, 2)
	}

	dispersion = dispersion / float64(N-1)

	graph := chart.Chart{
		Width:  N * 2,
		Height: 500,
		XAxis: chart.XAxis{
			Name: "Time",
		},
		YAxis: chart.YAxis{
			Name: "X value",
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: y,
			},
		},
	}

	os.Chdir("/home/heirloom/go/src/github.com/mdapathy/embedded2.1/")
	f, _ := os.Create("graph.png")
	graph.Render(chart.PNG, f)
	f.Close()

	for p := 0; p < N; p++ {
		for i := 0; i < N-1; i++ {
			Freal[p] += y[i] * math.Cos(math.Pi*2*float64(p*i)/float64(N))
			Fimag[p] += y[i] * math.Sin(math.Pi*2*float64(p*i)/float64(N))
		}
		F[p] = math.Sqrt(Freal[p]*Freal[p] + Fimag[p]*Fimag[p])
	}

	graph2 := chart.Chart{
		Width:  N * 2,
		Height: 500,
		XAxis: chart.XAxis{
		},
		YAxis: chart.YAxis{
		},
		Series: []chart.Series{
			chart.ContinuousSeries{
				XValues: x,
				YValues: F,
			},
		},
	}

	f, _ = os.Create("graph2.png")
	graph2.Render(chart.PNG, f)
	f.Close()

	// Create a new browser window instance
	window, err := gotron.New("/home/heirloom/go/src/github.com/mdapathy/embedded2.1/html")
	if err != nil {
		panic(err)
	}

	window.WindowOptions.Width = 1200
	window.WindowOptions.Height = 980
	window.WindowOptions.Title = "Lab2.1"

	done, err := window.Start()
	if err != nil {
		panic(err)
	}

	// Wait for the application to close
	<-done

}
