package main

import "database/sql"

type Product struct {
	Id       int
	Name     string
	Category int
	Price    float64
}

func queryDatabase(db *sql.DB) []Product {
	rows, err := db.Query("SELECT * from Products")

	products := []Product{}

	if err == nil {
		for rows.Next() {
			// var id, category int
			// var name string
			// var price float64

			// var id, category, name, price interface{} // ну или псевдоним any

			// rows.Scan(&id, &name, &category, &price)
			// Printfln("Row: %v %v %v %v", id, name, category, price)
			p := Product{}
			scanErr := rows.Scan(&p.Id, &p.Name, &p.Category, &p.Price)

			if scanErr == nil {
				products = append(products, p)
			} else {
				Printfln("Scan error: %v", scanErr)
				break
			}
		}
	} else {
		Printfln("Error: %v", err)
	}
	return products
}

func insertAndUseCategory(name string, productIDs ...int) {
	result, err := insertNewCategory.Exec(name)
	if err == nil {
		newID, _ := result.LastInsertId()
		for _, id := range productIDs {
			changeProductCategory.Exec(int(newID), id)
		}
	} else {
		Printfln("Prepared statement error: %v", err)
	}
}

func main() {
	listDrivers()
	db, err := openDatabase()
	if err == nil {
		products := queryDatabase(db)
		for _, v := range products {
			Printfln("id = %v, name = %v, category = %v, price = %v", v.Id, v.Name, v.Category, v.Price)
		}

		insertAndUseCategory("Misc Products", 2)
        p := queryDatabase(db)
        Printfln("Product: %v", p)

		db.Close()
	} else {
		panic(err)
	}

}
