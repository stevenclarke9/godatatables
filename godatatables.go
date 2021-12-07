// Copyright 2020 <stevenclarke2{AT}bigpond.com>. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.
// Package godatatables creates and manipulates data files in CSV format.
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

type tableHeader []string

type DataRow []string

type DataTable struct {
	header   tableHeader
	table    []DataRow
	rowCount int64
}

type DataTables []DataTable

/*
type JoinedColumn struct {
	leftTableColumn  int
	rightTableColumn int
}
*/

type dRow []DataRow

func removeOne(dr DataRow, s int) DataRow {
	lenDr := len(dr)
	newDr := make(DataRow, lenDr-1)
	j, k := 0, 0
	for j < lenDr {
		if j != s {
			if k < (lenDr - 1) {
				newDr[k] = dr[j]
				k = k + 1
			}
		}
		j = j + 1
	}
	return newDr

	// return append(slice[:s], slice[s+1:]...)
}

func removeHeaderIndex(s []string, i int) []string {
	lenS := len(s)
	newS := make([]string, lenS-1)
	j, k := 0, 0
	for j < lenS {
		if j != i {
			if k < (lenS - 1) {
				newS[k] = s[j]
				k = k + 1
			}
		}
		j = j + 1
	}
	return newS
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

func removeHeaderColumns(elements []string, columnIndexes []int) []string {
	header := elements
	for i := 0; i < len(columnIndexes); i++ {
		header = removeHeaderIndex(header, columnIndexes[i])
	}
	// Return the new slice.
	return header
}

func (dr DataRow) getTableColumn(fieldIndex int64) string {
	return dr[fieldIndex]
}

func (dt DataTable) getTableRow(rowIndex int64) *DataRow {
	return &dt.table[rowIndex]
}

// ToDo: DataTableReader
func ReadTable(r io.Reader, hasHeader bool) (dt DataTable, err error) {
	tableReader := csv.NewReader(r)
	tableReader.Comma = '|'
	tableReader.Comment = '#'
	stringsArray := [][]string{}

	stringsArray, err = tableReader.ReadAll()
	if err == nil {
		if hasHeader {
			stringsHeader := stringsArray[0]
			stringsArray = stringsArray[1:]
			dt = NewDataTable(stringsArray, false)
			dt.header = stringsHeader
		} else {
			dt = NewDataTable(stringsArray, false)
		}
	}
	return dt, err
}

// ToDo: DataTableWriter
func NewDataTable(records [][]string, hasHeader bool) DataTable {
	dt := DataTable{header: []string{}, table: []DataRow{}, rowCount: 0}
	if hasHeader {
		dt.header = records[0]
		records = records[1:]
	}
	dt.table = make([]DataRow, len(records))
	for i := 0; i < len(records); i++ {
		//for _, row := range records {
		dr := make(DataRow, len(records[i]))
		for j := 0; j < len(records[i]); j++ {
			//for _, col := range row {
			//	dr = append(dr, col)
			dr[j] = records[i][j]
		}
		dt.table[i] = dr
		dt.rowCount++
		//dt.AppendRow(dr)
	}
	return dt
}

func (dt *DataTable) Where(f func(dr DataRow) bool) *DataTable {
	newDt := DataTable{
		header:   dt.header,
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
	emptyTable := dtJoined

	checkedLeftColumnIndexes, IsValidLeftColumnIndexes := dt.validateColumnIndexes(joinLeftColumnIndexes)
	checkedRightColumnIndexes, IsValidRightColumnIndexes := dt.validateColumnIndexes(joinRightColumnIndexes)

	if removeDuplicateColumns {
		lenDt := len(dt.header)
		if lenDt > 0 {
			leftTableNames := dt.header
			rightTableNames := []string{}
			// fmt.Println("joinRightColumnIndexes",joinRightColumnIndexes)
			rightTableNames = removeHeaderColumns(joinTable.header, joinRightColumnIndexes)
			// fmt.Println("rightTableNames",rightTableNames)
			dtJoined.header = append(dtJoined.header, leftTableNames...)
			dtJoined.header = append(dtJoined.header, rightTableNames...)
		}
	}
	if IsValidLeftColumnIndexes && IsValidRightColumnIndexes {
		lenLeftColumnIndexes := len(checkedLeftColumnIndexes)
		lenRightColumnIndexes := len(checkedRightColumnIndexes)
		if lenLeftColumnIndexes == lenRightColumnIndexes {
			for _, leftTableRow := range dt.table {
				colEqual := 0
				for _, rightTableRow := range joinTable.table {
					for colIndex, leftColNumber := range checkedLeftColumnIndexes {
						rightColNumber := checkedRightColumnIndexes[colIndex]
						if leftTableRow[leftColNumber] == rightTableRow[rightColNumber] {
							colEqual = colEqual + 1
						}
					}
					// fmt.Println("colEqual = ", colEqual)
					if colEqual == lenLeftColumnIndexes {
						// a = append(a[:i], a[i+1:]...)
						tableRow := leftTableRow
						if removeDuplicateColumns {
							columnsRemovedTableRow := removeColumns(rightTableRow, checkedRightColumnIndexes)
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
			return &emptyTable
		}
	} else {
		// if both tables have invalid join indexes, return nil.
		return &emptyTable
	}
}

func (dt *DataTable) Order(colIndexes []int) *DataTable {
	//fmt.Println("the order method")
	sortedDt := dt
	sortedDt.sortRow(colIndexes)
	return sortedDt
}

func (dt *DataTable) DistinctRows() *DataTable {
	distinctTable := &DataTable{
		header:   []string{},
		table:    []DataRow{},
		rowCount: 0}

	// make a slice of row indexes
	u := make([]int64, 0, len(dt.table))

	// make a map of seen row indexes
	m := make(map[string]int)

	for dtRowIndex, dtRow := range dt.table {
		rowString := string(dtRow.String())
		if _, ok := m[rowString]; !ok {
			m[rowString] = dtRowIndex
			u = append(u, int64(dtRowIndex))
		}
	}
	var rowcount int64 = 0

	for _, v := range u {
		rowcount++
		distinctTable.table = append(distinctTable.table, dt.table[v])
	}
	distinctTable.rowCount = rowcount
	return distinctTable

}

func (dt *DataTable) validateColumnIndexes(colIndexes []int) ([]int, bool) {
	dtRowLen := 0
	if len(dt.table) > 0 {
		dtRowLen = len(dt.table[0])
	}

	colIndexesValid := false
	colIndexesChecked := []int{}
	for _, colNumber := range colIndexes {
		if (colNumber >= 0) && (colNumber < dtRowLen) {
			colIndexesChecked = append(colIndexesChecked, colNumber)
			colIndexesValid = true
		} else {
			colIndexesValid = false
			break
		}
	}
	return colIndexesChecked, colIndexesValid
}

func findInStringSlice(name string, sliceNames []string) int {
	for i := 0; i < len(sliceNames); i++ {
		if sliceNames[i] == name {
			return i
		}
	}
	return -1
}

func (dt *DataTable) validateColumnNames(colNames []string) (colIndexesChecked []int, colIndexesValid bool) {
	// fmt.Println("start validateColumnNames")
	// fmt.Println("dt.header ",dt.header, "colNames ",colNames)
	if len(dt.header) == 0 {
		return []int{}, false
	}

	colIndexesValid = true
	for i := 0; i < len(colNames); i++ {
		ih := findInStringSlice(colNames[i], dt.header)
		// fmt.Println("colNames[",i,"] =",colNames[i],"dt.header =",dt.header)
		if ih > -1 {
			// the value of "ih" is the header array index.
			colIndexesChecked = append(colIndexesChecked, ih)
		} else {
			colIndexesValid = false
			break
		}
	}
	// fmt.Println("colIndexesvalid:", colIndexesValid)
	// fmt.Println("colIndexesChecked:", colIndexesChecked)
	// fmt.Println("exit validateColumnNames")

	return colIndexesChecked, colIndexesValid

}

// IndexOfColumnName return -1 if it is a headerless table.
// and return -1 if the column name is not found in the table header.
func (dt *DataTable) IndexOfColumnName(colName string) int {
	if len(dt.header) > 0 {
		ih := findInStringSlice(colName, dt.header)
		return ih
	} else {
		return -1
	}
}

func (dt *DataTable) SelectByColumnNames(colNames []string) *DataTable {

	newDt := NewDataTable([][]string{}, false)

	colIndexesChecked, colIndexesValid := dt.validateColumnNames(colNames)
	if colIndexesValid {
		selectedDt := dt.Select(colIndexesChecked)
		return selectedDt
	}
	return &newDt
}

// Select returns a new DataTable that contains the selected columns from the "dt" DataTable.
// The column index starts from 0 to number of columns in the "dt" DataTable less 1.
// This func allows call chaining. eg. 	recs.Select([]int{0,2}).Select([]int{1})
func (dt *DataTable) Select(colIndexes []int) *DataTable {

	newDt := NewDataTable([][]string{}, false)

	// silently remove invalid column index values.
	// colIndexesValid is false when all values in the colIndexes array is invalid.
	// Also colIndexesChecked becomes an empty array.
	colIndexesChecked, colIndexesValid := dt.validateColumnIndexes(colIndexes)
	if colIndexesValid {
		if len(dt.header) > 0 {
			for _, colNumber := range colIndexesChecked {
				newDt.header = append(newDt.header, dt.header[colNumber])
			}
		}
		for _, dtRow := range dt.table {
			newDtRow := DataRow{}
			for _, colNumber := range colIndexesChecked {
				newDtRow = append(newDtRow, dtRow[colNumber])
			}
			newDt.AppendRow(newDtRow)
		}
	}
	return &newDt
}

func (dt *DataTable) IsEmpty() bool {
	if (len(dt.header) == 0) && (len(dt.table) == 0) {
		return true
	}
	return false
}

func (dt *DataTable) Count() int64 {
	return dt.rowCount
}

func (a dRow) cmpDataTable(b dRow) bool {
	// if we call this function then assume the tables have the same number of rows.
	for rowIndex := 0; rowIndex < len(a); rowIndex++ {
		if len(a[rowIndex]) != len(b[rowIndex]) {
			return false
		} else {
			// to do: we need to check each element of the row in 'a' is the same as each element in 'b'.
			// at the moment all element types are string.
			aLen := len(a[rowIndex])
			for elementIndex := 0; elementIndex < aLen; elementIndex++ {
				if a[rowIndex][elementIndex] == b[rowIndex][elementIndex] {
					continue
				}
				return false
			}
		}
	}
	// all the rows have the same number of elements and the values of the elements in each of the rows are the same in both slices of DataRow.
	return true
}

func (a tableHeader) cmpHeader(b tableHeader) bool {
	// fmt.Println("a", a, "b", b)
	if len(a) == len(b) {
		for index := 0; index < len(a); index++ {
			if a[index] != b[index] {
				// fmt.Println("a[",index,"] != b[",index,"]", a[index],"!=", b[index])
				return false
			}
		}
		return true
	}
	// the lenghs of a and b are not equal, so return false.
	return false
}

// Cmp
// compare two data tables.
// returns true if both tables are equal.
func (a *DataTable) Cmp(b *DataTable) bool {
	// fmt.Println("a.Count",a.Count(),"b.Count",b.Count())
	if a.Count() == b.Count() {
		// fmt.Println("count equal")
		if a.header.cmpHeader(b.header) {
			// fmt.Println("headers are equal")
			if dRow(a.table).cmpDataTable(dRow(b.table)) {
				// fmt.Println("rows are equal")
				return true
			} else {
				return false
			}
		} else {
			return false
		}
	} else {
		return false
	}
}

// compareFloat64 is a helper function for sortRow.
// returns 0 if i is equal to j
// returns -1 if i is less than j
// returns 1 if i is greater than j
func compareFloat64(i float64, j float64) int {
	if i == j {
		return 0
	} else {
		if i < j {
			return -1
		} else {
			return 1
		}
	}
}

// compareString is a helper function for sortRow.
// returns 0 if i ia equal to j
// returns -1 if i is less than j
// returns 1 if i is greater than j
func compareString(i string, j string) int {
	if i == j {
		return 0
	} else {
		if i < j {
			return -1
		} else {
			return 1
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
				// tmpResult, maxIndex := compareFloat64(float32I, float32J, lenIndex)
				intCmp := compareFloat64(float32I, float32J)
				if intCmp == 0 {
					continue
				}
				if intCmp == -1 {
					result = true
					k = lenIndex
				} else {
					// otherwise intCmp equals 1
					result = false
					k = lenIndex
				}
			} else {
				// otherwise they can be compared as strings.
				// tmpResult, maxIndex := compareString(valueI, valueJ, lenIndex)
				intCmp := compareString(valueI, valueJ)
				if intCmp == 0 {
					continue
				}
				if intCmp == -1 {
					result = true
					k = lenIndex
				} else {
					// otherwise intCmp equals 1
					result = false
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
	if len(dt.header) > 0 {
		sfmt = strings.Join(dt.header, "|")
		sfmt = sfmt + "\n" + "----------" + "\n"
	}
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

