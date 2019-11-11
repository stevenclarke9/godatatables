// Package godatatables creates and manipulates data files in CSV format.
// Date: Monday 11 November 2019
package godatatables

import (
	"encoding/csv"
	// "fmt"
	"io"
	"math"
	"sort"
	"strconv"
	"strings"
)

type DataRow []string

type DataTable struct {
	header   []string
	table    []DataRow
	rowCount int64
}

type DataTables []DataTable

type JoinedColumn struct {
	leftTableColumn  int
	rightTableColumn int
}

type dRow []DataRow

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

func (dr DataRow) GetTableColumn(columnIndex int64) string {
	return dr[columnIndex]
}

func (dt DataTable) GetTableRow(rowIndex int64) *DataRow {
	return &dt.table[rowIndex]
}

// ToDo: DataTableReader
func ReadTable(r io.Reader, header bool) (dt DataTable, err error) {
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
	dt := DataTable{header: []string{}, table: []DataRow{}, rowCount: 0}
	for _, row := range records {
		dr := DataRow{}
		for _, col := range row {
			dr = append(dr, col)
		}
		dt.AppendRow(dr)
		// Disable adding to the rowCount because one is added in the AppendRow function.
		// dt.rowCount++
	}
	return dt
}

func (dt *DataTable) Where(f func(dr DataRow) bool) *DataTable {
	newDt := DataTable{
		header:   []string{},
		table:    []DataRow{},
		rowCount: 0}
	for _, dtRow := range dt.table {
		if f(dtRow) {
			newDt.AppendRow(dtRow)
		}
	}
	return &newDt
}

// Sum
// Return the sum of the supplied index
func (dt DataTable) Sum(index int64) (float64, error) {
	var sum float64 = 0.0

	for _, dtRow := range dt.table {
		colString := dtRow[index]
		colValue, err := strconv.ParseFloat(colString, 64)
		if err == nil {
			sum = sum + colValue
		} else {
			return sum, err
		}
	}
	return sum, nil
}

func (dt *DataTable) AppendRow(dr DataRow) {
	dt.table = append(dt.table, dr)
	dt.rowCount++
}

func (dt *DataTable) InnerJoin(removeDuplicateColumns bool, joinLeftColumnIndexes []int, joinRightColumnIndexes []int, joinTable DataTable) *DataTable {
	dtJoined := DataTable{
		header:   []string{},
		table:    []DataRow{},
		rowCount: 0}
	lenLeftColumnIndexes := len(joinLeftColumnIndexes)
	lenRightColumnIndexes := len(joinRightColumnIndexes)

	if lenLeftColumnIndexes == lenRightColumnIndexes {
		for _, leftTableRow := range dt.table {
			colEqual := 0
			for _, rightTableRow := range joinTable.table {
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
	sortedDt := dt
	sortedDt.sortRow(colIndexes)
	return sortedDt
}

// Select returns a new DataTable that contains the selected columns from the "dt" DataTable.
// The column index starts from 0 to number of columns in the "dt" DataTable less 1.
// This func allows call chaining. eg. 	recs.Select([]int{0,2}).Select([]int{1})
func (dt *DataTable) Select(colIndexes []int) *DataTable {
	newDt := NewDataTable([][]string{})
	for _, dtRow := range dt.table {
		newDtRow := DataRow{}
		for _, colNumber := range colIndexes {
			newDtRow = append(newDtRow, dtRow[colNumber])
		}
		newDt.table = append(newDt.table, newDtRow)
	}
	return &newDt
}

func (dt *DataTable) Count() int64 {
	return dt.rowCount
}


func (a dRow) cmpDataTable(b dRow) bool {
// if we call this function then assume the tables have the same number of rows.
	for rowIndex := 0; rowIndex < len(a); rowIndex++ {
		if (len(a[rowIndex]) != len(b[rowIndex])) {
			return false
		} else {
			// to do: we need to check each element of the row in 'a' is the same as each element in 'b'.
			aLen := len(a[rowIndex])
			for elementIndex := 0; elementIndex < aLen; elementIndex++ {
				if (a[rowIndex][elementIndex] == b[rowIndex][elementIndex]) {
					continue
				}
				return false
			}
		}
	}
	// all the rows have the same number of elements and the values of the elements in each of the rows are the same in both slices of DataRow.
	return true
}

// Cmp
// compare two data tables.
// returns true if both tables are equal.
func (a *DataTable) Cmp(b *DataTable) bool {
	if (a.Count() == b.Count()) {
		if (dRow(a.table).cmpDataTable(dRow(b.table))) {
			return true
		} else {
			return false
		}
	} else {
		return false
	}
}

// compareFloat64 is a helper function for sortRow.
// returns false and -1 if i is equal to j
// returns true and maxIndex if i is less than j
// returns false and maxIndex if i is greater than j
func compareFloat64(i float64, j float64, maxIndex int) (bool, int) {
	if i == j {
		return false, -1
	} else {
		if i < j {
			return true, maxIndex
		} else {
			return false, maxIndex
		}
	}
}

// compareString is a helper function for sortRow.
// returns false and -1 if i ia equal to j
// returns true and maxIndex if i is less than j
// returns false and maxIndex if i is greater than j
func compareString(i string, j string, maxIndex int) (bool, int) {
	if i == j {
		return false, -1
	} else {
		if i < j {
			return true, maxIndex
		} else {
			return false, maxIndex
		}
	}
}

func (dt *DataTable) sortRow(indexes []int) {
	// this is a comment
	sort.Slice(dt.table, func(i, j int) bool {
		// fmt.Println(i, dt.Table[i][0], dt.Table[i][1], dt.Table[i][2], " : ", j, dt.Table[j][0], dt.Table[j][1], dt.Table[j][2])
		var lenIndex int
		if len(indexes) == 0 {
			lenIndexI := len(dt.table[i])
			lenIndexJ := len(dt.table[j])
			lenIndex = int(math.Min(float64(lenIndexI), float64(lenIndexJ)))
		} else {
			lenIndex = len(indexes)
		}
		result := false
		for k := 0; k < lenIndex; k++ {
			var float32I float64 = 0.0
			var float32J float64 = 0.0
			var okI error = nil
			var okJ error = nil
			isFloat64I := false
			isFloat64J := false
			l := k
			reverse := false
			if len(indexes) > 0 {
				l = indexes[k]
				if l < 0 {
					reverse = true
					l = -l
				}
			}
			valueI := dt.table[i][l]
			valueJ := dt.table[j][l]
			if reverse {
				// swap the values
				valueI, valueJ = valueJ, valueI
			}
			// check for a float number
			float32I, okI = strconv.ParseFloat(valueI, 32)
			if okI == nil {
				isFloat64I = true
			}
			// check for a float number
			float32J, okJ = strconv.ParseFloat(valueJ, 32)
			if okJ == nil {
				isFloat64J = true
			}
			if isFloat64I && isFloat64J {
				tmpResult, maxIndex := compareFloat64(float32I, float32J, lenIndex)
				if maxIndex == -1 {
					continue
				} else {
					result = tmpResult
					k = lenIndex
				}
			} else {
				// otherwise they can be compared as strings.
				tmpResult, maxIndex := compareString(valueI, valueJ, lenIndex)
				if maxIndex == -1 {
					continue
				} else {
					result = tmpResult
					k = lenIndex
				}
			}
		}
		return result
	})
}

func (dt DataTable) String() string {
	sfmt := ""
	//sfmt = "TABLE START\n"
	/* Ignore the dataRowIndex value */
	for _, dataRow := range dt.table {
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
		for _, dataRow := range dataTable.table {
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
