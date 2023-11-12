package surface3d

import (
	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
)

var surfaceRangeColor = []string{
	"#313695", "#4575b4", "#74add1", "#abd9e9", "#e0f3f8",
	"#fee090", "#fdae61", "#f46d43", "#d73027", "#a50026",
}

type Surface3dData struct {
	Id	string
	Value [][3]int
}

func transformTo3dSurfaceData(Value [][3]int) []opts.Chart3DData {
	items := make([]opts.Chart3DData, 0, len(Value))
	for _, a := range Value {
		items = append(items, opts.Chart3DData{Value: []interface{}{a[0], a[1], a[2]}})
	}
	return items
}

func CreateSurface3d(surfacedataraw []Surface3dData, title string, rose bool) *charts.Surface3D {
	surface3d := charts.NewSurface3D()
	surface3d.SetGlobalOptions(
		charts.WithTitleOpts(opts.Title{Title: title}),
		charts.WithVisualMapOpts(opts.VisualMap{
			Calculable: true,
			InRange:    &opts.VisualMapInRange{Color: surfaceRangeColor},
			Max:        3,
			Min:        -3,
		}),
	)
	
	for _, a := range surfacedataraw {
		surface3d.AddSeries(a.Id, transformTo3dSurfaceData(a.Value))
	}
	
	return surface3d
}