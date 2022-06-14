package main

import (
	"fmt"
	"mongo_test/crud"
)

func main() {
	// fmt.Println("##### connections #####")
	// crud.MainConnect()

	fmt.Println("##### inserts #####")
	crud.MainInsert()

	// fmt.Println("##### finds #####")
	// crud.MainFind()

	// fmt.Println("##### updates #####")
	// crud.MainUpdate()

	// fmt.Println("##### deletes #####")
	// crud.DeleteMain()
}
