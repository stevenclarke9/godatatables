package godatatables

import (
	"StevenClarke9/godatatables"
	"strings"
	"testing"
)

func TestCreateDatatables(t *testing.T) {

	stringTableReader := strings.NewReader(`a|b|c|value
1|aa|bb|10
2|aa|cc|15
3|aa|ee|30
4|aa|dd|60
5|bb|cc|120
6|bb|aa|240
7|bb|ee|480
8|bb|dd|960`)
// 9|cc|cc|1920

	dataTableStringsArray := [][]string{
	{"a","b","c","value"},
	{"1","aa","bb","10"},
	{"2","aa","cc","15"},
	{"3","aa","ee","30"},
	{"4","aa","dd","60"},
	{"5","bb","cc","120"},
	{"6","bb","aa","240"},
	{"7","bb","ee","480"},
	{"8","bb","dd","960"}}

	dataTableFromArray := godatatables.NewDataTable(dataTableStringsArray)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}
	if ok := dataTableFromArray.Cmp(&dataTableFromStringTableReader); !ok {
		t.Errorf("tables are not equal")
	}
	
}

