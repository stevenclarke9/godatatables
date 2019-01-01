# godatatables
CSV data file processing implemented in Google's "Go" language.

The data file is a csv with the delimiter as the PIPE ("|") character.
The "#" is used as a comment character.
The specification of the csv file format is the specification in the golang csv package.

Columns values are obtained by the column index number, starting from 0 to the number of columns on the data line.

Most of the methods allow method chaining by returning a pointer to a DataTable.

Dealing with headers is not yet implemented.

Private types and methods.

func removeOne(slice DataRow, s int) DataRow
func removeColumns(elements DataRow, columnIndexes []int) DataRow

Exported types and methods.

type DataRow []string
type DataTable struct {
	header   []string
	Table    []DataRow
	rowCount int64
}

type DataTables []DataTable

The JoinedColumn is not implemented at the moment.

type JoinedColumn struct {
	leftTableColumn  int
	rightTableColumn int
}


func ReadTable(r io.Reader) (dt DataTable, err error)
func NewDataTable(records [][]string) DataTable
func (dt *DataTable) Where(f func(dr DataRow) bool) *DataTable
func (dt *DataTable) AppendRow(dr DataRow)
func (dt *DataTable) InnerJoin(removeDuplicateColumns bool, joinLeftColumnIndexes []int, joinRightColumnIndexes []int, joinTable DataTable) *DataTable
func (dt *DataTable) Order(colIndexes []int) *DataTable
func (dt *DataTable) Count() int

Sort interface implementation

type dRow []DataRow

func (t dRow) Len() int
func (t dRow) Less(i, j int) bool
func (t dRow) Swap(i, j int)

Stringer interface implementation

func (dt DataTable) String() string
func (dts DataTables) String() string
func (dtr DataRow) String() string
