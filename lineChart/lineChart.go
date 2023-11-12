package lineChart

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"fmt"
)

type LineValues struct {
	Id string
	Value []int
}

func transformToLineData(v []int) ([]opts.LineData) {
	items := make([]opts.LineData, 0)
	for a := 0; a < len(v); a++ {
		value := v[a]
		items = append(items, opts.LineData{Value: value})
	}
	return items
} 

func CreateLine(xaxis []string, lv []LineValues, title, xaxisname, yaxisname string, smooth, symbol bool) *charts.Line {
	line := charts.NewLine()
	line.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{
			Title:    title,
			Subtitle: "",
			Link:     "",
		}),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:  "inside",
			Start: 10,
			End:   50,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: xaxisname,
			Data: xaxis, // Set x-axis data to your xaxis values
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: yaxisname,
		}),
	)

	if len(xaxis) != len(lv[0].Value) {
		fmt.Println("Error: Length of xaxis and lv.Value must be the same")
		return nil
	}

	line.SetXAxis(xaxis)
	
	for _, l := range lv {
		line.AddSeries(l.Id, transformToLineData(l.Value)).SetSeriesOptions(charts.WithLineChartOpts(
			opts.LineChart{Smooth: smooth, ShowSymbol: symbol, SymbolSize: 15, Symbol: "diamond"},
		),
			charts.WithLabelOpts(opts.Label{
				Show: true,
			}),
		)
	}

	return line
}