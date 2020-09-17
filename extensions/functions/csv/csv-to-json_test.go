package csv

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var cj = &CsvToJson{}

func TestCsvParser(t *testing.T) {
	csv :=
		`col1;col2;col3;col4
row11;row12;row13;row14
row21;row22;row23;row24`

	result, _ := cj.Eval(csv, ";")

	structResult := result.(ParsedCsv)

	assert.Equal(t, len(structResult.Columns), 4)
	assert.Equal(t, len(structResult.Rows), 2)
	assert.Equal(t, len(structResult.Rows[0]), 4)
}
