package scrummy

// Chart represents each graph build of a coodinates system and (possibly multiple) data lines.
type Chart interface {
	Xaxis() []float64
	Yaxis() []float64
	Xlegend() string
	Ylegend() string
	Values() []float64
	Legend() []string
}

// LineChart is a diagram with one or more data lines.
type LineChart struct {
	xaxis    []float64
	yaxisL   []float64
	yaxisR   []float64
	xlegend  string
	ylegendL string
	ylegendR string
	values   [][]float64
	legend   []string
}

// NewLineChart is the constructor for LineChart.
func NewLineChart(xaxis, yaxisL []float64, xlegend, ylegendR string) *LineChart {
	l := new(LineChart)
	l.xaxis = xaxis
	l.yaxisL = yaxisL
	l.xlegend = xlegend
	l.ylegendR = ylegendR
	return l
}

// Xaxis retrieves the values of the x-axis.
func (l *LineChart) Xaxis() []float64 {
	return l.xaxis
}

// YaxisL retrieves the values of the left y-axis.
func (l *LineChart) YaxisL() []float64 {
	return l.yaxisL
}

// YaxisR retrieves the values of the right y-axis.
func (l *LineChart) YaxisR() []float64 {
	return l.yaxisR
}

// SetYAxisR set the values and the title for the optional right y-axis.
func (l *LineChart) SetYAxisR(values []float64, title string) {
	l.yaxisR = values
	l.ylegendR = title
}

// Xlegend retrieves the title of the x-axis.
func (l *LineChart) Xlegend() string {
	return l.xlegend
}

// YlegendL retrieves the title of the left y-axis.
func (l *LineChart) YlegendL() string {
	return l.ylegendL
}

// YlegendR retrieves the title of the right y-axis.
func (l *LineChart) YlegendR() string {
	return l.ylegendR
}

// ValueMatrix retrieves the values of the line plots.
func (l *LineChart) ValueMatrix() [][]float64 {
	return l.values
}

// Values retrieves the values of the line plots.
func (l *LineChart) Values(i int) []float64 {
	return l.values[i]
}

// AddValues adds a new dataset (line) and its legend to the chart.
// It  returns the index of the added dataset.
func (l *LineChart) AddValues(values []float64, legend string) int {
	l.values = append(l.values, values)
	l.legend = append(l.legend, legend)
	return len(l.values) - 1
}

// Legend retrieves the titles of the line plots.
func (l *LineChart) Legend() []string {
	return l.Legend()
}
