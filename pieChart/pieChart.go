package pieChart

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

func transformToPieData(n []string, v []int) []opts.PieData {
	items := make([]opts.PieData, 0)
	for i, a := range n {
		items = append(items, opts.PieData{Name: a, Value: v[i]})
	}
	return items
}

func CreatePie(names []string, values []int, title string, radius bool, rose bool) *charts.Pie {
	pie := charts.NewPie()
	pie.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
	)

	pie.AddSeries("pie", transformToPieData(names, values))

	pie.SetSeriesOptions(
		charts.WithLabelOpts(opts.Label{
			Show:	true,
			Formatter:	"{b}: {c}",
		}),
	)

	if radius == true {
		pie.SetSeriesOptions(
			charts.WithPieChartOpts(opts.PieChart{
				Radius:   []string{"40%", "75%"},
			}),
		)
	}

	if rose == true {
		pie.SetSeriesOptions(
			charts.WithPieChartOpts(opts.PieChart{
				Radius: []string{"40%", "75%"},
				RoseType: "area",
			}),
		)
	}

	return pie
}