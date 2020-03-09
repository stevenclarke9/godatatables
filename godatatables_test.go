package godatatables_test

import (
	"fmt"
	"godatatables"
	"strings"
	"testing"
)

// there must be a carriage return after the last data line.
// this is what is printed as the last character output from the
// godatatables.String() function.
var stringTableWithoutHeader string = `1|aa|bb|10
2|aa|cc|15
3|aa|ee|30
4|aa|dd|60
5|bb|cc|120
6|bb|aa|240
7|bb|ee|480
8|bb|dd|960
`
var stringTableWithHeader string = `a|b|c|value
1|aa|bb|10
2|aa|cc|15
3|aa|ee|30
4|aa|dd|60
5|bb|cc|120
6|bb|aa|240
7|bb|ee|480
8|bb|dd|960
`

var resultStringTableWithHeader string = `a|b|c|value
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

var	dataTableStringsArrayWithoutHeader [][]string = [][]string{
	{"1","aa","bb","10"},
	{"2","aa","cc","15"},
	{"3","aa","ee","30"},
	{"4","aa","dd","60"},
	{"5","bb","cc","120"},
	{"6","bb","aa","240"},
	{"7","bb","ee","480"},
	{"8","bb","dd","960"}}

var	dataTableStringsArrayWithHeader [][]string = [][]string{
	{"a","b","c","value"},
	{"1","aa","bb","10"},
	{"2","aa","cc","15"},
	{"3","aa","ee","30"},
	{"4","aa","dd","60"},
	{"5","bb","cc","120"},
	{"6","bb","aa","240"},
	{"7","bb","ee","480"},
	{"8","bb","dd","960"}}

// there must be a carriage return after the last data line.
// this is what is printed as the last character output from the
// godatatables.String() function.
var stringSelectedColumnsTableWithoutHeader string = `1|bb|10
2|cc|15
3|ee|30
4|dd|60
5|cc|120
6|aa|240
7|ee|480
8|dd|960
`
// there must be a carriage return after the last data line.
// this is what is printed as the last character output from the
// godatatables.String() function.
var stringSelectedColumnsTableWithHeader string = `a|c|value
1|bb|10
2|cc|15
3|ee|30
4|dd|60
5|cc|120
6|aa|240
7|ee|480
8|dd|960
`

var myVerbose bool = false

func forResultDataTableWithoutHeader(stringTableWithoutHeader string) (godatatables.DataTable, error) {
	stringTableReader := strings.NewReader(stringTableWithoutHeader)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)
    return dataTableFromStringTableReader, err
}

func forResultDataTableWithHeader(stringTableWithHeader string) (godatatables.DataTable, error) {
	stringTableReader := strings.NewReader(stringTableWithHeader)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,true)
    return dataTableFromStringTableReader, err
}

func forTestCreateDataTablesWithNoHeader() (godatatables.DataTable, godatatables.DataTable, error) {
	dataTableFromArray := godatatables.NewDataTable(dataTableStringsArrayWithoutHeader,false)
	stringTableReader := strings.NewReader(stringTableWithoutHeader)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)
    return dataTableFromArray, dataTableFromStringTableReader, err
}

func forTestCreateDataTablesWithHeader() (godatatables.DataTable, godatatables.DataTable, error) {
	dataTableFromArray := godatatables.NewDataTable(dataTableStringsArrayWithHeader,true)
	stringTableReader := strings.NewReader(stringTableWithHeader)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,true)
    return dataTableFromArray, dataTableFromStringTableReader, err
}

func forTestSelectDataTableColumnsWithNoHeader() (godatatables.DataTable, godatatables.DataTable, error) {
    var dataTableFromStringSelectedColumnsTableReader godatatables.DataTable

	stringTableReader := strings.NewReader(stringTableWithoutHeader)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

    if err == nil {
        stringSelectedColumnsTableReader := strings.NewReader(stringSelectedColumnsTableWithoutHeader)
    	dataTableFromStringSelectedColumnsTableReader, err = godatatables.ReadTable(stringSelectedColumnsTableReader,false)
    }
    return dataTableFromStringTableReader, dataTableFromStringSelectedColumnsTableReader, err
}

func forTestSelectDataTableColumnsWithColumnNames() (godatatables.DataTable, godatatables.DataTable, error) {
    var dataTableFromStringSelectedColumnsTableReader godatatables.DataTable

	stringTableReader := strings.NewReader(stringTableWithHeader)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,true)

    if err == nil {
        stringSelectedColumnsTableReader := strings.NewReader(stringSelectedColumnsTableWithHeader)
    	dataTableFromStringSelectedColumnsTableReader, err = godatatables.ReadTable(stringSelectedColumnsTableReader,true)
    }
    return dataTableFromStringTableReader, dataTableFromStringSelectedColumnsTableReader, err
}

/*
Input/Result,DataTable variable, has a Header line, data type
------------------------------------------------
Input,dataTableStringsArrayWithoutHeader, no, array [][]string
Input,stringTableWithoutHeader, no, string
Result,dataTableFromArray, no, DataTable
Result,dataTableFromStringTableReader, no DataTable
Compare,dataTableFromArray,dataTableFromStringTableReader
*/ 

func printOrLog(t *testing.T, s string) {
	if myVerbose {
		fmt.Println(s)
	} else {
		t.Log(s)
	}
}

func TestCreateDataTablesWithNoHeader(t *testing.T) {

    dataTableFromArray, dataTableFromStringTableReader, err := forTestCreateDataTablesWithNoHeader()

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

    dataTableFromArray, dataTableFromStringTableReader, err := forTestCreateDataTablesWithHeader()

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


func TestSelectDataTableColumnsWithNoHeader(t *testing.T) {

    dataTableFromStringTableReader,dataTableFromStringSelectedColumnsTableReader, err := forTestSelectDataTableColumnsWithNoHeader()

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

func TestSelectDataTableColumnsWithValidColumnNames(t *testing.T) {

    dataTableFromStringTableReader,dataTableFromStringSelectedColumnsTableReader, err := forTestSelectDataTableColumnsWithColumnNames()

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

	sprintDataTableFromStringSelectedColumnsTableReader := fmt.Sprint(dataTableFromStringSelectedColumnsTableReader)
	printOrLog(t,"sprintDataTableFromStringSelectedColumnsTableReader:")
	printOrLog(t,sprintDataTableFromStringSelectedColumnsTableReader)

    dataTableSelectedColumns := dataTableFromStringTableReader.SelectByColumnNames([]string{"a","c","value"})
    printOrLog(t,fmt.Sprint("dataTableSelectedColunms row count = ",dataTableSelectedColumns.Count()))

	sprintDataTableSelectedColumns := fmt.Sprint(dataTableSelectedColumns)
	printOrLog(t,"sprintDataTableSelectedColumns:")
	printOrLog(t,sprintDataTableSelectedColumns)

	if ok := dataTableSelectedColumns.Cmp(&dataTableFromStringSelectedColumnsTableReader); !ok {
		t.Errorf("tables are not equal")
	}	
}

func TestSelectDataTableColumnsWithInvalidColumnNames(t *testing.T) {

    dataTableFromStringTableReader,dataTableFromStringSelectedColumnsTableReader, err := forTestSelectDataTableColumnsWithColumnNames()

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

	sprintDataTableFromStringSelectedColumnsTableReader := fmt.Sprint(dataTableFromStringSelectedColumnsTableReader)
	printOrLog(t,"sprintDataTableFromStringSelectedColumnsTableReader:")
	printOrLog(t,sprintDataTableFromStringSelectedColumnsTableReader)

    invalidColumnNames := []string{"a","invalid","c","value"}
    printOrLog(t,fmt.Sprint("invalidColumnNames: ",invalidColumnNames))
    dataTableSelectedColumns := dataTableFromStringTableReader.SelectByColumnNames(invalidColumnNames)

    printOrLog(t,fmt.Sprint("dataTableSelectedColunms row count = ",dataTableSelectedColumns.Count()))

    if dataTableSelectedColumns.Count() != 0 {
    	sprintDataTableSelectedColumns := fmt.Sprint(dataTableSelectedColumns)
    	printOrLog(t,"sprintDataTableSelectedColumns:")
    	printOrLog(t,sprintDataTableSelectedColumns)
        t.Errorf("a non-empty data table is returned!")
    }
}

func TestSelectDataTableWithInvalidColumnIndexValues(t *testing.T) {

    dataTableFromStringTableReader,dataTableFromStringSelectedColumnsTableReader, err := forTestSelectDataTableColumnsWithNoHeader()

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
    printOrLog(t,fmt.Sprint("dataTableSelectedColunms row count =",dataTableSelectedColumns.Count()))

	sprintDataTableSelectedColumns := fmt.Sprint(dataTableSelectedColumns)
	printOrLog(t,"sprintDataTableSelectedColumns:")
	printOrLog(t,sprintDataTableSelectedColumns)

    if dataTableSelectedColumns.IsEmpty() {
	    printOrLog(t,"the dataTableSelectedColumns value IsEmpty")
    }
    if dataTableSelectedColumns.IsEmpty() == false {
    	sprintDataTableSelectedColumns := fmt.Sprint(dataTableSelectedColumns)
    	printOrLog(t,"sprintDataTableSelectedColumns:")
    	printOrLog(t,sprintDataTableSelectedColumns)
        t.Errorf("a non-empty data table is returned!")
    }
}

func TestJoinDataTablesWithValidColumnIndexValues(t *testing.T) {

var sourceJoinTableB string = `a|c|value
1|bb|10
2|cc|15
3|ee|30
4|dd|60
5|cc|120
6|aa|240
7|ee|480
8|dd|960
`

var sourceJoinTableA string = `a|b
1|aa
2|aa
3|aa
4|aa
5|bb
6|bb
7|bb
8|bb
`

	stringTableReader := strings.NewReader(stringTableWithHeader)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	tableAReader := strings.NewReader(sourceJoinTableA)
	dataTableFromTableAReader, err := godatatables.ReadTable(tableAReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	tableBReader := strings.NewReader(sourceJoinTableB)
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

    joinedTable := dataTableFromTableAReader.InnerJoin(true, []int{0}, []int{0}, dataTableFromTableBReader)

	sprintJoinedTable := fmt.Sprint(joinedTable)
	printOrLog(t,"sprintJoinedTable:")
	printOrLog(t,sprintJoinedTable)

    if joinedTable.IsEmpty() == false {
    	if ok := joinedTable.Cmp(&dataTableFromStringTableReader); !ok {
    		t.Errorf("tables are not equal")
    	}
    } else {
   	    t.Errorf("the joinedTable value IsEmpty")
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

	stringTableReader := strings.NewReader(stringTableWithHeader)
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

    if joinedTable.IsEmpty() {
	    printOrLog(t,"the joinedTable value is nil")
    }
    if joinedTable.IsEmpty() == false {
    	sprintJoinedTable := fmt.Sprint(joinedTable)
    	printOrLog(t,"sprintJoinedTable:")
    	printOrLog(t,sprintJoinedTable)
        t.Errorf("a non-empty data table is returned!")
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

	stringTableReader := strings.NewReader(stringTableWithHeader)
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

    if joinedTable.IsEmpty() {
	    printOrLog(t,"the joinedTable value is nil")
    }
    if joinedTable.IsEmpty() == false {
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

    if joinedTable.IsEmpty() {
	    printOrLog(t,"the joinedTable value IsEmpty")
    }
    if joinedTable.IsEmpty() == false {
    	sprintJoinedTable := fmt.Sprint(joinedTable)
	    printOrLog(t,"sprintJoinedTable:")
    	printOrLog(t,sprintJoinedTable)
        t.Errorf("we shouldn't get here, thereis a problem with tableB joinIndex.")
    }
}

func TestDataTableWhereConditionUsingIndexValues(t *testing.T) {
var stringTableWithoutHeader string = `1|aa|bb|10
2|aa|cc|15
3|aa|ee|30
4|aa|dd|60
5|bb|cc|120
6|bb|aa|240
7|bb|ee|480
8|bb|dd|960
`

var stringWhereResultWithoutHeader string = `2|aa|cc|15
5|bb|cc|120
`
	stringTableReader := strings.NewReader(stringTableWithoutHeader)
	dataTableFromStringTableReader, err := godatatables.ReadTable(stringTableReader,false)

	whereResultTableReader := strings.NewReader(stringWhereResultWithoutHeader)
	correctWhereResultDataTable, err := godatatables.ReadTable(whereResultTableReader,false)

	if err != nil {
		t.Errorf("error recorded %s",err)
	}

	sprintDataTableFromStringTableReader := fmt.Sprint(dataTableFromStringTableReader)
	printOrLog(t,"sprintDataTableFromStringTableReader:")
	printOrLog(t,sprintDataTableFromStringTableReader)

	sprintCorrectWhereResultDataTable := fmt.Sprint(correctWhereResultDataTable)
	printOrLog(t,"sprintCorrectWhereResultDataTable:")
	printOrLog(t,sprintCorrectWhereResultDataTable)

    columnValue := "cc"
    columnIndex := 2
    printOrLog(t,fmt.Sprint("column Index: ", columnIndex, " Value: ", columnValue))

    whereFunc := func (dr godatatables.DataRow) bool {
        if dr[columnIndex] == columnValue {
            return true
        } else {
            return false
        }
    }

    whereResultDataTable := dataTableFromStringTableReader.Where(whereFunc)

    if whereResultDataTable.IsEmpty() == false {
    	if ok := whereResultDataTable.Cmp(&correctWhereResultDataTable); !ok {
    		t.Errorf("tables are not equal")
    	}
    } else {
   	    t.Errorf("the joinedTable value IsEmpty")
    }
}

