package main

import (
    "fmt"
    "os"

    "godatatables"
    // "github.com/StevenClarke9/godatatables"
)

func readFile(fileName string, hasHeader bool) (dt godatatables.DataTable, err error) {

	//	var fileHandle *os.File

	fh, err := os.Open(fileName)
	defer fh.Close()

	if err == nil {
		dt, err = godatatables.ReadTable(fh, hasHeader)
	}

	return dt, err
}

func loadJoinTables(f bool) {
    dtOwner , errOwner := readFile("owners.txt",f)
    if errOwner != nil {
        fmt.Println(errOwner)
        os.Exit(1)
    }
    dtPets, errPets := readFile("pets.txt",f)
    if errPets != nil {
        fmt.Println(errPets)
        os.Exit(1)
    }

	dtOwnerPets := dtOwner.InnerJoin(true, []int{0}, []int{1}, dtPets)

    fmt.Println("owners table, owners.txt")
    fmt.Println(dtOwner)

    fmt.Println("pets table, pets.txt")
    fmt.Println(dtPets)

    fmt.Println("Owners And Pets")
    fmt.Println(dtOwnerPets)

    fmt.Println("pets table after owners table joined with pets table")
    fmt.Println(dtPets)

}

func main() {

    loadJoinTables(true)

    fmt.Println("join the 'owners.txt' and 'pets.txt' tables with no header")
    loadJoinTables(false)

    os.Exit(0)

}
