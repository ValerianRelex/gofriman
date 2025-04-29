package main

import "database/sql"

func queryDatabase(db *sql.DB) {
	rows, err := db.Query("SELECT * from Products")

	if err == nil {
		for rows.Next() {
			// var id, category int
			// var name string
			// var price float64
			var id, category, name, price interface{} // ну или псевдоним any
			rows.Scan(&id, &name, &category, &price)
			Printfln("Row: %v %v %v %v", id, name, category, price)
		}
	} else {
		Printfln("Error: %v", err)
	}
}

func main() {
	listDrivers()
	db, err := openDatabase()
	if err == nil {
		queryDatabase(db)
		db.Close()
	} else {
		panic(err)
	}
}
