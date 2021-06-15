package draw

import (
	"os"

	"github.com/fogleman/gg"
	"github.com/wcharczuk/go-chart"
)

func DrawPieImage() {
	dc := gg.NewContext(600, 600)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(4)
	sizeX = float64(100)
	sizeY = float64(50)
	DrawDiamond(dc, 150.0, 100.0+sizeY+sizeY+(sizeY/2.0))
	DrawDiamond(dc, 150.0, 100.0)

	capTotal := 15.9506
	capTable := map[string]float64{}
	capTable["andrew_52"] = 8.3705
	capTable["donhardman_12"] = 1.9408
	capTable["Clout_Cast_10"] = 1.7111
	capTable["JasonDevlin"] = 1.5235
	capTable["clayoglesby"] = 0.6633
	capTable["Salvo"] = 0.5007

	chartMap := map[string]int{}
	sum := 0.0
	for k, v := range capTable {
		chartMap[k] = int((v / capTotal) * 100)
		sum += v
	}
	chartMap["other"] = int(((capTotal - sum) / capTotal) * 100)
	DrawChart(chartMap)
	dc.SavePNG("001.png")
}

func DrawChart(m map[string]int) {
	items := []chart.Value{}
	for k, v := range m {
		items = append(items, chart.Value{Value: float64(v), Label: k})
	}

	pie := chart.PieChart{
		Width:  400,
		Height: 400,
		Values: items,
	}

	f, _ := os.Create("chart.png")
	defer f.Close()
	pie.Render(chart.PNG, f)
}
