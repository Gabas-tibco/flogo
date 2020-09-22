package csv

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/project-flogo/core/data"
	"github.com/project-flogo/core/data/expression/function"
)

type CsvToJson struct {
}

type Row struct {
	Row []string `json:"row"`
}

type ParsedCsv struct {
	Columns []string `json:"columns"`
	Rows    []Row `json:"rows"`
}

func init() {
	function.Register(&CsvToJson{})
}

func (s *CsvToJson) Name() string {
	return "csvToJson"
}

func (s *CsvToJson) Sig() (paramTypes []data.Type, isVariadic bool) {
	return []data.Type{data.TypeString, data.TypeString}, false
}

func (s *CsvToJson) Eval(params ...interface{}) (interface{}, error) {
	csv := params[0].(string)
	separator := params[1].(string)

	scanner := bufio.NewScanner(strings.NewReader(csv))
	var columns []string
	var rows []Row = []Row{}
	i := 0
	for scanner.Scan() {
		if i == 0 {
			columns = strings.Split(strings.Trim(scanner.Text(), " "), separator)
		} else {
			rows = append(rows, Row{strings.Split(strings.Trim(scanner.Text(), " "), separator)})
		}
		i++
	}

	parsedCsv := ParsedCsv{
		columns,
		rows,
	}

	fmt.Println("CsvToJson -> Total cols " + strconv.Itoa(len(parsedCsv.Columns)) + " Total rows " + strconv.Itoa(len(parsedCsv.Rows)))

	return parsedCsv, nil
}
