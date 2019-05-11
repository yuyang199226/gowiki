package logic

import (
	"gowiki/db"
		"log"
)

func GetBookList() []string {
	rows, err := db.MysqlDB.Query("SELECT * FROM book;")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	names := make([]string, 0)
	for rows.Next() {
		var name string
		var id int
		var publisher string
		if err := rows.Scan(&id, &name, &publisher); err != nil {
			log.Fatal(err)
		}
		names = append(names, name)
	}
	return names
}