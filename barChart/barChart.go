package barChart

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type BarValues struct {
	Id string
	Value []int
}

func transformToBarData(v []int) ([]opts.BarData) {
	items := make([]opts.BarData, 0)
	for a := 0; a < len(v); a++ {
		value := v[a]
		items = append(items, opts.BarData{Value: value})
	}
	return items
}

func CreateBar(xaxis []string, bv []BarValues, title string, xaxisname string, yaxisname string) *charts.Bar {
	bar := charts.NewBar()
	// set some global options like Title/Legend/ToolTip or anything else
	bar.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title:    title}),
		charts.WithTooltipOpts(opts.Tooltip{Show: true}),
		charts.WithLegendOpts(opts.Legend{Show: true, Right: "80px"}),
		charts.WithToolboxOpts(opts.Toolbox{
			Show: true,
			Right: "50%", 
			Feature: &opts.ToolBoxFeature{
				SaveAsImage: &opts.ToolBoxFeatureSaveAsImage{
					Show:  true,
					Type:  "png",
					Title: "Download image",
				},
				DataView: &opts.ToolBoxFeatureDataView{
					Show:  true,
					Title: "DataView",
					Lang: []string{"data view", "turn off", "refresh"},
				},
			}},
		),
		charts.WithDataZoomOpts(opts.DataZoom{
			Type:  "inside",
			Start: 10,
			End:   50,
		}),
		charts.WithXAxisOpts(opts.XAxis{
			Name: xaxisname,
		}),
		charts.WithYAxisOpts(opts.YAxis{
			Name: yaxisname,
		}),
	)

	bar.SetXAxis(xaxis)

	for _, b := range bv {
		bar.AddSeries(b.Id, transformToBarData(b.Value))
	}

	return bar
}