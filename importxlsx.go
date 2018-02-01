package scrummy

import (
	"strconv"

	"github.com/wlbr/scrummy/gotils"
)

// XlsxImporter is
type XlsxImporter struct {
}

// NewXlsxImporter is a constructor for the struct XlsxImporter.
func NewXlsxImporter() *XlsxImporter {
	c := new(XlsxImporter)
	return c
}

func (x XlsxImporter) Read(s Session) {

}

// GetFloatColumn retirves a column as a float array  out of a string array of rows.
func GetFloatColumn(rows [][]string, columnindex int, offset int) []float64 {
	var column []float64
	for i := offset; i < len(rows) && rows[i][7] != ""; i++ {
		stringcell := rows[i][columnindex]
		floatcell, err := strconv.ParseFloat(stringcell, 64)
		if err != nil {
			gotils.LogError("Found %s being not a number while parsing excel input.", stringcell)
		} else {
			column = append(column, floatcell)
		}
	}
	return column
}
