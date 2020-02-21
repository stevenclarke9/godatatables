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


func TestSelectDataTableWithNoHeader(t *testing.T) {

// there must be a carriage return after the last data line.
// this is what is printed as the last character output from the
// godatatables.String() function.
var stringSelectedColumnsTable string = `a|c|value
1|bb|10
2|cc|15
3|ee|30
4|dd|60
5|cc|120
6|aa|240
7|ee|480
8|dd|960
`
	stringTableReader := strings.NewReader(stringTable)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

    stringSelectedColumnsTableReader := strings.NewReader(stringSelectedColumnsTable)
	dataTableFromStringSelectedColumnsTableReader, err := godatatables.ReadTable(stringSelectedColumnsTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

	sprintDataTableFromStringSelectedColumnsTableReader := fmt.Sprint(dataTableFromStringSelectedColumnsTableReader)
	printOrLog(t,"sprintDataTableFromStringSelectedColumnsTableReader:")
	printOrLog(t,sprintDataTableFromStringSelectedColumnsTableReader)

    dataTableSelectedColumns := dataTableFromStringTableReader.Select([]int{0,2,3})
    fmt.Println("dataTableSelectedColunms row count =",dataTableSelectedColumns.Count())
	sprintDataTableSelectedColumns := fmt.Sprint(dataTableSelectedColumns)
	printOrLog(t,"sprintDataTableSelectedColumns:")
	printOrLog(t,sprintDataTableSelectedColumns)

	if ok := dataTableSelectedColumns.Cmp(&dataTableFromStringSelectedColumnsTableReader); !ok {
		t.Errorf("tables are not equal")
	}	
}

func TestSelectDataTableWithInvalidColumnIndexValues(t *testing.T) {

// there must be a carriage return after the last data line.
// this is what is printed as the last character output from the
// godatatables.String() function.
var stringSelectedColumnsTable string = `a|c|value
1|bb|10
2|cc|15
3|ee|30
4|dd|60
5|cc|120
6|aa|240
7|ee|480
8|dd|960
`
	stringTableReader := strings.NewReader(stringTable)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

    stringSelectedColumnsTableReader := strings.NewReader(stringSelectedColumnsTable)
	dataTableFromStringSelectedColumnsTableReader, err := godatatables.ReadTable(stringSelectedColumnsTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

	sprintDataTableFromStringSelectedColumnsTableReader := fmt.Sprint(dataTableFromStringSelectedColumnsTableReader)
	printOrLog(t,"sprintDataTableFromStringSelectedColumnsTableReader:")
	printOrLog(t,sprintDataTableFromStringSelectedColumnsTableReader)

    dataTableSelectedColumns := dataTableFromStringTableReader.Select([]int{0,-2,2,3})
    fmt.Println("dataTableSelectedColunms row count =",dataTableSelectedColumns.Count())

	sprintDataTableSelectedColumns := fmt.Sprint(dataTableSelectedColumns)
	printOrLog(t,"sprintDataTableSelectedColumns:")
	printOrLog(t,sprintDataTableSelectedColumns)

	if ok := dataTableSelectedColumns.Cmp(&dataTableFromStringSelectedColumnsTableReader); !ok {
		t.Errorf("tables are not equal")
	}	
}

func TestJoinDataTablesWithValidColumnIndexValues(t *testing.T) {

var tableA string = `a|c|value
1|bb|10
2|cc|15
3|ee|30
4|dd|60
5|cc|120
6|aa|240
7|ee|480
8|dd|960
`

var tableB string = `a|b
1|aa
2|aa
3|aa
4|aa
5|bb
6|bb
7|bb
8|bb
`

	stringTableReader := strings.NewReader(stringTable)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	tableAReader := strings.NewReader(tableA)
	dataTableFromTableAReader, err := godatatables.ReadTable(tableAReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	tableBReader := strings.NewReader(tableB)
	dataTableFromTableBReader, err := godatatables.ReadTable(tableBReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

	sprintDataTableFromTableAReader := fmt.Sprint(dataTableFromTableAReader)
	printOrLog(t,"sprintDataTableFromTableAReader:")
	printOrLog(t,sprintDataTableFromTableAReader)

	sprintDataTableFromTableBReader := fmt.Sprint(dataTableFromTableBReader)
	printOrLog(t,"sprintDataTableFromTableBReader:")
	printOrLog(t,sprintDataTableFromTableBReader)

    joinedTable := dataTableFromTableBReader.InnerJoin(true, []int{0}, []int{0}, dataTableFromTableAReader)

	sprintJoinedTable := fmt.Sprint(joinedTable)
	printOrLog(t,"sprintJoinedTable:")
	printOrLog(t,sprintJoinedTable)

	if ok := joinedTable.Cmp(&dataTableFromStringTableReader); !ok {
		t.Errorf("tables are not equal")
	}	

}

func TestJoinDataTablesWithInvalidColumnIndexValues(t *testing.T) {
var tableA string = `a|c|value
1|bb|10
2|cc|15
3|ee|30
4|dd|60
5|cc|120
6|aa|240
7|ee|480
8|dd|960
`

var tableB string = `a|b
1|aa
2|aa
3|aa
4|aa
5|bb
6|bb
7|bb
8|bb
`

	stringTableReader := strings.NewReader(stringTable)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	tableAReader := strings.NewReader(tableA)
	dataTableFromTableAReader, err := godatatables.ReadTable(tableAReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	tableBReader := strings.NewReader(tableB)
	dataTableFromTableBReader, err := godatatables.ReadTable(tableBReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

	sprintDataTableFromTableAReader := fmt.Sprint(dataTableFromTableAReader)
	printOrLog(t,"sprintDataTableFromTableAReader:")
	printOrLog(t,sprintDataTableFromTableAReader)

	sprintDataTableFromTableBReader := fmt.Sprint(dataTableFromTableBReader)
	printOrLog(t,"sprintDataTableFromTableBReader:")
	printOrLog(t,sprintDataTableFromTableBReader)

    joinedTable := dataTableFromTableBReader.InnerJoin(true, []int{0,-1}, []int{-2,0}, dataTableFromTableAReader)

	sprintJoinedTable := fmt.Sprint(joinedTable)
	printOrLog(t,"sprintJoinedTable:")
	printOrLog(t,sprintJoinedTable)

	if ok := joinedTable.Cmp(&dataTableFromStringTableReader); !ok {
		t.Errorf("tables are not equal")
	}	

}

func TestJoinDataTablesWithAllOfTableA_InvalidJoinedColumnIndexValues(t *testing.T) {
var tableA string = `a|c|value
1|bb|10
2|cc|15
3|ee|30
4|dd|60
5|cc|120
6|aa|240
7|ee|480
8|dd|960
`

var tableB string = `a|b
1|aa
2|aa
3|aa
4|aa
5|bb
6|bb
7|bb
8|bb
`

	stringTableReader := strings.NewReader(stringTable)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	tableAReader := strings.NewReader(tableA)
	dataTableFromTableAReader, err := godatatables.ReadTable(tableAReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	tableBReader := strings.NewReader(tableB)
	dataTableFromTableBReader, err := godatatables.ReadTable(tableBReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

    tableAjoinIndex := []int{-1}
    fmt.Println("TableA join index =", tableAjoinIndex)
    tableBjoinIndex := []int{0}
    fmt.Println("TableB join index =", tableBjoinIndex)

	sprintDataTableFromTableAReader := fmt.Sprint(dataTableFromTableAReader)
	printOrLog(t,"sprintDataTableFromTableAReader:")
	printOrLog(t,sprintDataTableFromTableAReader)

	sprintDataTableFromTableBReader := fmt.Sprint(dataTableFromTableBReader)
	printOrLog(t,"sprintDataTableFromTableBReader:")
	printOrLog(t,sprintDataTableFromTableBReader)

    joinedTable := dataTableFromTableBReader.InnerJoin(true, tableBjoinIndex, tableAjoinIndex, dataTableFromTableAReader)

    if joinedTable != nil {
    	sprintJoinedTable := fmt.Sprint(joinedTable)
	    printOrLog(t,"sprintJoinedTable:")
    	printOrLog(t,sprintJoinedTable)
        t.Errorf("we shouldn't get here, thereis a problem with tableA joinIndex.")
    }

}

func TestJoinDataTablesWithAllOfTableB_InvalidJoinedColumnIndexValues(t *testing.T) {
var tableA string = `a|c|value
1|bb|10
2|cc|15
3|ee|30
4|dd|60
5|cc|120
6|aa|240
7|ee|480
8|dd|960
`

var tableB string = `a|b
1|aa
2|aa
3|aa
4|aa
5|bb
6|bb
7|bb
8|bb
`

	stringTableReader := strings.NewReader(stringTable)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	tableAReader := strings.NewReader(tableA)
	dataTableFromTableAReader, err := godatatables.ReadTable(tableAReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	tableBReader := strings.NewReader(tableB)
	dataTableFromTableBReader, err := godatatables.ReadTable(tableBReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

    tableAjoinIndex := []int{0}
    fmt.Println("TableA join index =", tableAjoinIndex)
    tableBjoinIndex := []int{-2}
    fmt.Println("TableB join index =", tableBjoinIndex)

	sprintDataTableFromTableAReader := fmt.Sprint(dataTableFromTableAReader)
	printOrLog(t,"sprintDataTableFromTableAReader:")
	printOrLog(t,sprintDataTableFromTableAReader)

	sprintDataTableFromTableBReader := fmt.Sprint(dataTableFromTableBReader)
	printOrLog(t,"sprintDataTableFromTableBReader:")
	printOrLog(t,sprintDataTableFromTableBReader)

    joinedTable := dataTableFromTableBReader.InnerJoin(true, tableBjoinIndex, tableAjoinIndex, dataTableFromTableAReader)

    if joinedTable != nil {
    	sprintJoinedTable := fmt.Sprint(joinedTable)
	    printOrLog(t,"sprintJoinedTable:")
    	printOrLog(t,sprintJoinedTable)
        t.Errorf("we shouldn't get here, thereis a problem with tableA joinIndex.")
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
