// Package godatatables creates and manipulates data files in CSV format.
// Date: Tue Dec 18, 2:42 PM

package godatatables

import (
	"encoding/csv"
	_ "fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

type DataRow []string

type DataTable struct {
	header   []string
	Table    []DataRow
	rowCount int64
}

type DataTables []DataTable

type JoinedColumn struct {
	leftTableColumn  int
	rightTableColumn int
}

func removeOne(slice DataRow, s int) DataRow {
	return append(slice[:s], slice[s+1:]...)
}

func removeColumns(elements DataRow, columnIndexes []int) DataRow {
	result := DataRow{}

	elementsIndex := 0
	lenElementsIndex := len(elements)
	columnIndex := 0
	lenColumnIndex := len(columnIndexes)

	for elementsIndex < lenElementsIndex {
		if columnIndex < lenColumnIndex {
			if elementsIndex == columnIndexes[columnIndex] {
				// skip.
				columnIndex++
			} else {
				// append the element to the result
				result = append(result, elements[elementsIndex])
			}
			elementsIndex++
		} else {
			// append the element to the result
			result = append(result, elements[elementsIndex])
			elementsIndex++
		}
	}
	// Return the new slice.
	return result
}

// ToDo: DataTableReader
func ReadTable(r io.Reader) (dt DataTable, err error) {
	tableReader := csv.NewReader(r)
	tableReader.Comma = '|'
	tableReader.Comment = '#'
	stringsArray := [][]string{}

	stringsArray, err = tableReader.ReadAll()
	if err == nil {
		dt = NewDataTable(stringsArray)
	}
	return dt, err
}

// ToDo: DataTableWriter
func NewDataTable(records [][]string) DataTable {
	dt := DataTable{header: []string{}, Table: []DataRow{}, rowCount: 0}
	for _, row := range records {
		dr := DataRow{}
		for _, col := range row {
			dr = append(dr, col)
		}
		dt.AppendRow(dr)
		dt.rowCount++
	}
	return dt
}

func (dt *DataTable) Where(f func(dr DataRow) bool) *DataTable {
	newDt := DataTable{
		header:   []string{},
		Table:    []DataRow{},
		rowCount: 0}
	for _, dtRow := range dt.Table {
		if f(dtRow) {
			newDt.AppendRow(dtRow)
		}
	}
	return &newDt
}

func (dt *DataTable) AppendRow(dr DataRow) {
	dt.Table = append(dt.Table, dr)
	dt.rowCount++
}

func (dt *DataTable) InnerJoin(removeDuplicateColumns bool, joinLeftColumnIndexes []int, joinRightColumnIndexes []int, joinTable DataTable) *DataTable {
	dtJoined := DataTable{
		header:   []string{},
		Table:    []DataRow{},
		rowCount: 0}
	lenLeftColumnIndexes := len(joinLeftColumnIndexes)
	lenRightColumnIndexes := len(joinRightColumnIndexes)

	if lenLeftColumnIndexes == lenRightColumnIndexes {
		for _, leftTableRow := range dt.Table {
			colEqual := 0
			for _, rightTableRow := range joinTable.Table {
				for colIndex, leftColNumber := range joinLeftColumnIndexes {
					rightColNumber := joinRightColumnIndexes[colIndex]
					if leftTableRow[leftColNumber] == rightTableRow[rightColNumber] {
						colEqual = colEqual + 1
					}
				}
				// fmt.Println("colEqual = ", colEqual)
				if colEqual == lenLeftColumnIndexes {
					// a = append(a[:i], a[i+1:]...)
					tableRow := leftTableRow
					if removeDuplicateColumns {
						columnsRemovedTableRow := removeColumns(rightTableRow, joinRightColumnIndexes)
						tableRow = append(tableRow, columnsRemovedTableRow[:]...)
					} else {
						tableRow = append(tableRow, rightTableRow[:]...)
					}
					dtJoined.AppendRow(tableRow)
				}
				colEqual = 0
			}
		}
		return &dtJoined
	} else {
		return nil
	}
}

func (dt *DataTable) Order(colIndexes []int) *DataTable {
	//fmt.Println("the order method")
	sort.Sort(dRow(dt.Table))
	return dt
}

// Select returns a new DataTable that contains the selected columns from the "dt" DataTable.
// The column index starts from 0 to number of columns in the "dt" DataTable less 1.
// This func allows call chaining. eg. 	recs.Select([]int{0,2}).Select([]int{1})
func (dt *DataTable) Select(colIndexes []int) *DataTable {
	newDt := NewDataTable([][]string{})
	for _, dtRow := range dt.Table {
		newDtRow := DataRow{}
		for _, colNumber := range colIndexes {
			newDtRow = append(newDtRow, dtRow[colNumber])
		}
		newDt.Table = append(newDt.Table, newDtRow)
	}
	return &newDt
}

func (dt *DataTable) Count() int64 {
	return dt.rowCount
}

type dRow []DataRow

func (t dRow) Len() int {
	count := 0
	for _, i := range t {
		_ = i
		count++
	}
	return count
}

func (t dRow) Less(i, j int) bool {
	//	fmt.Println(i, t[i][0], t[i][1], t[i][2], " : ", j, t[j][0], t[j][1], t[j][2])
	lenIndexI := len(t[i])
	lenIndexJ := len(t[j])
	result := false
	for k := 0; (k < lenIndexI) && (k < lenIndexJ); k++ {
		var float32I float64
		var float32J float64
		float32I = 0.0
		float32J = 0.0
		isFloat64I := false
		isFloat64J := false
		isIntI := false
		isIntJ := false
		// first check t[i][k] is an integer number
		numberI , okI := strconv.Atoi(t[i][k])
		if okI == nil {
			isIntI = true
		}
		// and check that t[j][k] is an integer number
		numberJ , okJ := strconv.Atoi(t[j][k])
		if okI == nil {
			isIntJ = true
		}
		if (okI != nil) {
			// check for a float number
			float32I, okI = strconv.ParseFloat(t[i][k],32)
			if okI == nil {
				isFloat64I = true
			}
		}
		if (okJ != nil) {
			// check for a float number
			float32J, okJ = strconv.ParseFloat(t[j][k],32)
			if okI == nil {
				isFloat64J = true
			}
		}
		if (okI == nil) && (okJ == nil) {
			// then change both strings to a number and compare them as numbers.
			if isIntI && isIntJ {
				if numberI == numberJ {
					continue
				} else {
					if numberI < numberJ {
						result = true
						k = lenIndexI
					} else {
						result = false
						k = lenIndexI
					}
				}
			}
			if isIntI && isFloat64J {
				float32I = float64(numberI)
				if float32I == float32J {
					continue
				} else {
					if float32I < float32J {
						result = true
						k = lenIndexI
					} else {
						result = false
						k = lenIndexI
					}
				}
			}			
			if isFloat64I && isIntJ {
				float32J = float64(numberJ)
				if float32I == float32J {
					continue
				} else {
					if float32I < float32J {
						result = true
						k = lenIndexI
					} else {
						result = false
						k = lenIndexI
					}
				}
			}
			if isFloat64I && isFloat64J {
				if float32I == float32J {
					continue
				} else {
					if float32I < float32J {
						result = true
						k = lenIndexI
					} else {
						result = false
						k = lenIndexI
					}
				}
			}			

		} else {
			// otherwise they can be compared as strings.
			if t[i][k] == t[j][k] {
				continue
			} else {
				if t[i][k] < t[j][k] {
					result = true
					k = lenIndexI - 1
				} else {
					result = false
					k = lenIndexI - 1
				}
			}
		}
	}
	return result
}

func (t dRow) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (dt DataTable) String() string {
    sfmt := ""
	//sfmt = "TABLE START\n"
	/* Ignore the dataRowIndex value */
	for _, dataRow := range dt.Table {
		sfmt = sfmt + strings.Join(dataRow, "|")
		sfmt = sfmt + "\n"
	}
	//sfmt = sfmt + "TABLE END\n"
	return sfmt
}

//Print the DataTables as a string.
func (dts DataTables) String() string {
	sfmt := ""
	for dataTableIndex, dataTable := range dts {
		sfmt = sfmt + "TABLE " + strconv.Itoa(dataTableIndex) + " START\n"
		/* Ignore the dataRowIndex value */
		for _, dataRow := range dataTable.Table {
			sfmt = sfmt + strings.Join(dataRow, "|")
			sfmt = sfmt + "\n"
		}
		sfmt = sfmt + "TABLE " + strconv.Itoa(dataTableIndex) + " END\n"
	}
	return sfmt
}

//Print a DataRow as a string
func (dtr DataRow) String() string {
	return strings.Join(dtr, "|")
}

