package database

import (
	"encoding/csv"
	"log"
	"os"
)

func writeLine(records [][]string) {
	w := csv.NewWriter(os.Stdout)
	w.WriteAll(records) // calls Flush internally

	if err := w.Error(); err != nil {
		log.Fatalln("error writing csv:", err)
	}
}

func WriteToDatabase(line string) {
	records := [][]string{{line}}

	writeLine(records)
}
