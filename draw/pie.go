package draw

import (
	"os"

	"github.com/fogleman/gg"
	"github.com/wcharczuk/go-chart"
)

type CreatorCoin struct {
	Username string
	Supply   int64
	Price    int64
}

func (cc *CreatorCoin) ToChartMap(capTable map[string]int64) map[string]int {
	chartMap := map[string]int{}
	sum := int64(0)
	for k, v := range capTable {
		chartMap[k] = int((float64(v) / float64(cc.Supply)) * 100)
		sum += v
	}
	chartMap["other"] = int((float64(cc.Supply-sum) / float64(cc.Supply)) * 100)
	return chartMap
}

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

	cc := CreatorCoin{}
	cc.Username = "andrewarrow"
	cc.Supply = 15950600000
	capTable := map[string]int64{}
	capTable["andrew_52"] = 8370500000
	capTable["donhardman_12"] = 1940800000
	capTable["Clout_Cast_10"] = 1711100000
	capTable["JasonDevlin"] = 1523500000
	capTable["clayoglesby"] = 663300000
	capTable["Salvo"] = 500700000

	DrawChart(cc.ToChartMap(capTable))
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
