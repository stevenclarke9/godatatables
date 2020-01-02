package godatatables_test

import (
	"fmt"
	"github.com/StevenClarke9/godatatables"
	"strings"
	"testing"
)

// there must be a carriage return after the last data line.
// this is what is printed as the last character output from the
// godatatables.String() function.
var stringTable string = `a|b|c|value
1|aa|bb|10
2|aa|cc|15
3|aa|ee|30
4|aa|dd|60
5|bb|cc|120
6|bb|aa|240
7|bb|ee|480
8|bb|dd|960
`
var stringTableWithHeader string = `a|b|c|value
----------
1|aa|bb|10
2|aa|cc|15
3|aa|ee|30
4|aa|dd|60
5|bb|cc|120
6|bb|aa|240
7|bb|ee|480
8|bb|dd|960
`

var	dataTableStringsArray [][]string = [][]string{
	{"a","b","c","value"},
	{"1","aa","bb","10"},
	{"2","aa","cc","15"},
	{"3","aa","ee","30"},
	{"4","aa","dd","60"},
	{"5","bb","cc","120"},
	{"6","bb","aa","240"},
	{"7","bb","ee","480"},
	{"8","bb","dd","960"}}

var myVerbose bool = false

func printOrLog(t *testing.T, s string) {
	if myVerbose {
		fmt.Println(s)
	} else {
		t.Log(s)
	}
}

func TestCreateDatatablesWithNoHeader(t *testing.T) {
	dataTableFromArray := godatatables.NewDataTable(dataTableStringsArray,false)
	stringTableReader := strings.NewReader(stringTable)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromArray := fmt.Sprint(dataTableFromArray)
	printOrLog(t,"sprintDataTableFromArray:")
	printOrLog(t,sprintDataTableFromArray)

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

	if ok := dataTableFromArray.Cmp(&dataTableFromStringTableReader); !ok {
		t.Errorf("tables are not equal")
	}	
}

func TestCreateDatatablesWithHeader(t *testing.T) {
	dataTableFromArray := godatatables.NewDataTable(dataTableStringsArray,true)
	stringTableReader := strings.NewReader(stringTable)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,true)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromArray := fmt.Sprint(dataTableFromArray)
	printOrLog(t,"sprintDataTableFromArray:")
	printOrLog(t,sprintDataTableFromArray)

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

	if ok := dataTableFromArray.Cmp(&dataTableFromStringTableReader); !ok {
		t.Errorf("tables are not equal")
	}	
}

/*
func TestPrintedDatatableWithNoHeader(t *testing.T) {
	stringTableReader := strings.NewReader(stringTable)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	fmt.Println(sprintDataTableFromStringTableReader)

	if stringTable != fmt.Sprint(dataTableFromStringTableReader) {
		t.Errorf("tables are not equal, got '%s', wanted '%s'\n",sprintDataTableFromStringTableReader,stringTable)
	}
}

func TestPrintedDatatableWithHeader(t *testing.T) {
	stringTableReader := strings.NewReader(stringTable)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,true)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	fmt.Println(sprintDataTableFromStringTableReader)

	if stringTable != fmt.Sprint(dataTableFromStringTableReader) {
		t.Errorf("tables are not equal, got '%s', wanted '%s'\n",sprintDataTableFromStringTableReader,stringTableWithHeader)
	}
}
*/
