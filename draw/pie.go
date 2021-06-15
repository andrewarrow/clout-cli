package draw

import (
	"fmt"
	"os"

	"github.com/fogleman/gg"
	"github.com/wcharczuk/go-chart"
)

type CreatorCoin struct {
	Username string
	Supply   int64
	Price    int64
	CapTable map[string]int64
	History  []int64
	Rewards  []int64
}

func (cc *CreatorCoin) ToChartMap() map[string]int {
	chartMap := map[string]int{}
	sum := int64(0)
	for k, v := range cc.CapTable {
		val := int((float64(v) / float64(cc.Supply)) * 100)
		chartMap[fmt.Sprintf("%s_%d", k, val)] = val
		sum += v
	}
	chartMap["other"] = int((float64(cc.Supply-sum) / float64(cc.Supply)) * 100)
	return chartMap
}
func (cc *CreatorCoin) Buy(who string, amount int64, rate float64) int64 {
	fr := int64(float64(amount) * rate)
	fixed := amount - fr
	cc.History = append(cc.History, fixed)
	cc.Rewards = append(cc.Rewards, fr)
	cc.Supply += fixed
	cc.CapTable[who] += fixed
	return fr
}
func (cc *CreatorCoin) BuyNoFR(who string, amount int64) {
	cc.Supply += amount
	cc.CapTable[who] += amount
}
func (cc *CreatorCoin) DrawChartWithBuys(filename string) {
	dc := gg.NewContext(600, 600)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	DrawChart(cc.ToChartMap(), "chart.png")
	im, _ := gg.LoadPNG("chart.png")
	dc.DrawImage(im, 0, 0)
	font := "arial.ttf"
	dc.LoadFontFace(font, 48)
	dc.DrawStringAnchored("price", 25, 50, 0.5, 0.5)
	dc.SavePNG(filename)
}

func DrawPieImage() {
	cc := CreatorCoin{}
	cc.Username = "andrewarrow"
	cc.Supply = 15950600000
	cc.CapTable = map[string]int64{}
	cc.History = []int64{}
	cc.Rewards = []int64{}
	cc.CapTable["andrewarrow"] = 8370500000
	cc.CapTable["donhardman"] = 1940800000
	cc.CapTable["Clout_Cast"] = 1711100000
	cc.CapTable["JasonDevlin"] = 1523500000
	cc.CapTable["clayoglesby"] = 663300000
	cc.CapTable["Salvo"] = 500700000

	fr := 0.0
	cc.DrawChartWithBuys("001.png")
	cc.Buy("Clout_Cast", 1711100000, fr)
	cc.DrawChartWithBuys("002.png")
	//DrawChart(cc.ToChartMap(), "002.png")

	/*
			c.BuyNoFR("andrewarrow", val)
			DrawChart(cc.ToChartMap(), "002.png")
			val = cc.Buy("Clout_Cast", 1711100000*4)
			cc.BuyNoFR("andrewarrow", val)
			DrawChart(cc.ToChartMap(), "003.png")
			val = cc.Buy("Clout_Cast", 1711100000*4)
			cc.BuyNoFR("andrewarrow", val)
			DrawChart(cc.ToChartMap(), "004.png")
			val = cc.Buy("Clout_Cast", 1711100000*8)
			cc.BuyNoFR("andrewarrow", val)
		DrawChart(cc.ToChartMap(), "005.png")
	*/
}

func DrawChart(m map[string]int, filename string) {
	items := []chart.Value{}
	for k, v := range m {
		items = append(items, chart.Value{Value: float64(v), Label: k})
	}

	pie := chart.PieChart{
		Width:  400,
		Height: 400,
		Values: items,
	}

	f, _ := os.Create(filename)
	defer f.Close()
	pie.Render(chart.PNG, f)
}
