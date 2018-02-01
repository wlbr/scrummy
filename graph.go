package scrummy

import (
	"fmt"
	"math"
	"os"

	xlsx "github.com/360EntSecGroup-Skylar/excelize"
	chart "github.com/wcharczuk/go-chart"
	"github.com/wlbr/scrummy/gotils"
)

// A Set represents a set of data on the chart, probably a line.
type Set struct {
	name   string
	data   []float64
	series *chart.Series
}

// NewSet constructs a new object of type Set.
func NewSet(name string, data []float64, series *chart.Series) *Set {
	s := new(Set)
	s.name = name
	s.data = data
	s.series = series
	return s
}

// A Visualizer renders controls the generation of a set of graphs.
type Visualizer struct {
	dataseries map[string]Set
}

// NewVisualizer constructs a new object of type Visualizer.
func NewVisualizer() *Visualizer {
	v := new(Visualizer)
	v.dataseries = make(map[string]Set)
	return v
}

func (v Visualizer) Render() {
	for k, s := range v.dataseries {

	}
}

// DrawGraphs -  well, it draws the graphs
func DrawGraphs(xfile string) {
	xlsx, err := xlsx.OpenFile(xfile)
	if err != nil {
		cwd, _ := os.Getwd()
		gotils.LogError("Cannot open excel file '%s'. Current working directory is '%s'.", xfile, cwd)
		return
	}

	offset := 4
	rows := xlsx.GetRows("Agile Progress")

	sprintno := GetFloatColumn(rows, 1, offset)

	cvelocity := GetFloatColumn(rows, 10, offset)
	cveloSeries := chart.ContinuousSeries{
		Name: "kapazitätsbereinigte Velocity (in SP/PT)",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorBlue,
			FillColor:   chart.ColorBlue.WithAlpha(100),
		},
		XValues: sprintno,
		YValues: cvelocity,
	}

	velocity := GetFloatColumn(rows, 8, offset)
	veloSeries := chart.ContinuousSeries{
		Name: "gleitende Velocity (in SP)",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorAlternateLightGray,
			FillColor:   chart.ColorLightGray.WithAlpha(100),
		},
		XValues: sprintno,
		YValues: velocity,
		YAxis:   chart.YAxisSecondary,
	}

	committment := GetFloatColumn(rows, 6, offset)
	committmentSeries := chart.ContinuousSeries{
		Name: "Commitment  (in SP)",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorOrange,
		},
		XValues: sprintno,
		YValues: committment,
		YAxis:   chart.YAxisSecondary,
	}

	achievement := GetFloatColumn(rows, 7, offset)
	achievementSeries := chart.ContinuousSeries{
		Name: "Achievement  (in SP)",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorGreen,
		},
		XValues: sprintno,
		YValues: achievement,
		YAxis:   chart.YAxisSecondary,
	}

	var result []float64
	for i, c := range committment {
		result = append(result, achievement[i]/c*100)
	}

	GetFloatColumn(rows, 7, offset)
	resultsSeries := chart.ContinuousSeries{
		Name: "Sprint Results (in %)",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorBlue,
			FillColor:   chart.ColorBlue.WithAlpha(100),
		},
		XValues: sprintno,
		YValues: result,
		//YAxis:   chart.YAxisSecondary,
	}

	var percentfloats []float64
	mi := math.Floor(gotils.Minf64(result)/10) * 10
	ma := math.Ceil(gotils.Maxf64(result)/10) * 10
	for i := mi; i <= ma; i = i + 10 {
		percentfloats = append(percentfloats, i)
	}

	var percentticks chart.Ticks
	for _, p := range percentfloats {
		percentticks = append(percentticks, chart.Tick{Value: p, Label: fmt.Sprintf("%.0f", p)})
	}
	var percentlines []chart.GridLine
	for _, s := range percentfloats {
		percentlines = append(percentlines, chart.GridLine{Value: s})
	}

	linreg := &chart.LinearRegressionSeries{
		Name: "Lineare Regression (in SP/PT)",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorAlternateBlue,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: cveloSeries,
	}

	sma := &chart.SMASeries{
		Name: "Einfacher gleitender Duchschnitt (in SP/PT)",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: cveloSeries,
	}

	resultlinreg := &chart.LinearRegressionSeries{
		Name: "Lineare Regression (in %)",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorAlternateBlue,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: resultsSeries,
	}

	resultsma := &chart.SMASeries{
		Name: "Einfacher gleitender Duchschnitt (in %)",
		Style: chart.Style{
			Show:            true,
			StrokeColor:     chart.ColorRed,
			StrokeDashArray: []float64{5.0, 5.0},
		},
		InnerSeries: resultsSeries,
	}

	var gridlines []chart.GridLine
	for _, s := range sprintno {
		gridlines = append(gridlines, chart.GridLine{Value: s})
	}

	var sprintticks chart.Ticks
	for _, s := range sprintno {
		sprintticks = append(sprintticks, chart.Tick{Value: s, Label: fmt.Sprintf("%.0f", s)})
	}

	scope := GetFloatColumn(rows, 12, offset)
	scopeSeries := chart.ContinuousSeries{
		Name: "Gesamtscope  (in SP)",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorAlternateBlue,
		},
		XValues: sprintno,
		YValues: scope,
	}

	openscope := GetFloatColumn(rows, 13, offset)
	openscopeSeries := chart.ContinuousSeries{
		Name: "Offener Scope  (in SP)",
		Style: chart.Style{
			Show:        true,
			StrokeColor: chart.ColorOrange,
		},
		XValues: sprintno,
		YValues: openscope,
	}

	velograph := chart.Chart{
		Width:      800,
		Height:     450,
		Title:      "Sprint Statistiken Done Team 5 (Velocity)",
		TitleStyle: chart.StyleShow(),

		Background: chart.Style{
			Padding: chart.Box{
				Top:  150,
				Left: 40,
			},
		},
		YAxis: chart.YAxis{
			Name:      "SP/PT",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
		},
		YAxisSecondary: chart.YAxis{
			Name: "SP",
			NameStyle: chart.Style{
				Show:                true,
				TextRotationDegrees: 270.0,
			},
			Style: chart.Style{
				Show: true, //enables / displays the secondary y-axis

			},
			ValueFormatter: func(v interface{}) string {
				return fmt.Sprintf("%.0f", v)
			},
		},
		XAxis: chart.XAxis{
			Name:      "Sprint",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: true,
			},
			Ticks: sprintticks,
			GridMajorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorLightGray,
				StrokeWidth: 1.0,
			},
			GridLines: gridlines,
		},
		Series: []chart.Series{
			veloSeries,
			cveloSeries,
			linreg,
			chart.LastValueAnnotation(linreg),
			sma,
			chart.LastValueAnnotation(sma),
		},
	}

	achievementgraph := chart.Chart{
		Width:      800,
		Height:     450,
		Title:      "Sprint Statistiken Done Team 5 (Achievements)",
		TitleStyle: chart.StyleShow(),

		Background: chart.Style{
			Padding: chart.Box{
				Top:  150,
				Left: 40,
			},
		},
		YAxis: chart.YAxis{
			Name:      "%",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			Ticks:     percentticks,
			GridMajorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorLightGray,
				StrokeWidth: 1.0,
			},
			GridLines: percentlines,
		},
		YAxisSecondary: chart.YAxis{
			Name:      "SP",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			ValueFormatter: func(v interface{}) string {
				return fmt.Sprintf("%.0f", v)
			},
		},
		XAxis: chart.XAxis{
			Name:      "Sprint",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: true,
			},
			Ticks: sprintticks,
			GridMajorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorLightGray,
				StrokeWidth: 1.0,
			},
			GridLines: gridlines,
		},
		Series: []chart.Series{
			committmentSeries,
			achievementSeries,
			resultsSeries,
			resultlinreg,
			resultsma,
		},
	}

	projectburndowngraph := chart.Chart{
		Width:      800,
		Height:     450,
		Title:      "Sprint Statistiken Done Team 5 (Project Burndown)",
		TitleStyle: chart.StyleShow(),
		Background: chart.Style{
			Padding: chart.Box{
				Top:  150,
				Left: 40,
			},
		},
		YAxis: chart.YAxis{
			Name:      "SP",
			NameStyle: chart.StyleShow(),
			Style:     chart.StyleShow(),
			GridMajorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorLightGray,
				StrokeWidth: 1.0,
			},
		},
		XAxis: chart.XAxis{
			Name:      "Sprint",
			NameStyle: chart.StyleShow(),
			Style: chart.Style{
				Show: true,
			},
			Ticks: sprintticks,
			GridMajorStyle: chart.Style{
				Show:        true,
				StrokeColor: chart.ColorLightGray,
				StrokeWidth: 1.0,
			},
			GridLines: gridlines,
		},
		Series: []chart.Series{
			scopeSeries,
			openscopeSeries,
		},
	}

	velograph.Elements = []chart.Renderable{chart.LegendThin(&velograph)}
	buffer, err := os.OpenFile("chart1.png", os.O_WRONLY|os.O_CREATE, 0644) // bytes.NewBuffer([]byte{})
	err = velograph.Render(chart.PNG, buffer)

	achievementgraph.Elements = []chart.Renderable{chart.LegendThin(&achievementgraph)}
	buffer, err = os.OpenFile("chart2.png", os.O_WRONLY|os.O_CREATE, 0644) // bytes.NewBuffer([]byte{})
	err = achievementgraph.Render(chart.PNG, buffer)

	projectburndowngraph.Elements = []chart.Renderable{chart.LegendThin(&projectburndowngraph)}
	buffer, err = os.OpenFile("chart3.png", os.O_WRONLY|os.O_CREATE, 0644) // bytes.NewBuffer([]byte{})
	err = projectburndowngraph.Render(chart.PNG, buffer)
}
