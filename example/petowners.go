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

func main() {

    dtOwner , errOwner := readFile("owners.txt",true)
    if errOwner != nil {
        fmt.Println(errOwner)
        os.Exit(1)
    }
    dtPets, errPets := readFile("pets.txt",true)
    if errPets != nil {
        fmt.Println(errPets)
        os.Exit(1)
    }

	dtOwnerPets := dtOwner.InnerJoin(true, []int{0}, []int{1}, dtPets)

    fmt.Println("owners.txt")
    fmt.Println(dtOwner)

    fmt.Println("pets.txt")
    fmt.Println(dtPets)

    fmt.rintln("Owners And Pets")
    fmt.Println(dtOwnerPets)

    os.Exit(0)

}
