package cmd

import (
	"os"

	"github.com/go-echarts/go-echarts/v2/charts"
	"github.com/go-echarts/go-echarts/v2/opts"
	"github.com/go-echarts/go-echarts/v2/types"
)

type YAxis struct {
    Label string
    Values []int
}

// generate random data for line chart
func generateLineItems(items []int) []opts.LineData {
	//transform items into opts.LineData
	lineItems := make([]opts.LineData, len(items))
	for i, v := range items {
		lineItems[i] = opts.LineData{Value: v}
	}
	return lineItems
}

func GenerateChart(title string, xAxis []string, yAxis ...YAxis) {
	// create a new line instance
	line := charts.NewLine()
	// set some global options like Title/Legend/ToolTip or anything else
	line.SetGlobalOptions(
		charts.WithInitializationOpts(opts.Initialization{Theme: types.ThemeWesteros}),
		charts.WithTitleOpts(opts.Title{
			Title:    title,
		}))

	// Put data into instance
	lineWithX := line.SetXAxis(xAxis)

	for v := range yAxis {
		lineWithX.AddSeries(yAxis[v].Label, generateLineItems(yAxis[v].Values))
	}

	lineWithX.SetSeriesOptions(charts.WithLineChartOpts(opts.LineChart{Smooth: true}))
	f, _ := os.Create("report.html")
	line.Render(f)
}