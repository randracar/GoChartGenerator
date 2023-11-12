package line3d

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

type Line3dData struct {
	Id    string
	Value [][3]int
}

var line3DColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

func transformTo3dLineData(Value [][3]int) []opts.Chart3DData {
	items := make([]opts.Chart3DData, 0, len(Value))
	for _, a := range Value {
		items = append(items, opts.Chart3DData{Value: []interface{}{a[0], a[1], a[2]}})
	}
	return items
}

func CreateLine3d(linedataraw []Line3dData, title string) *charts.Line3D {
	line3d := charts.NewLine3D()
	line3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			Max:        30,
			InRange:    &opts.VisualMapInRange{Color: line3DColor},
		}),
	)

	for _, a := range linedataraw {
		line3d.AddSeries(a.Id, transformTo3dLineData(a.Value))
	}

	return line3d
}