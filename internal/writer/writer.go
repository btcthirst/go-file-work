package writer

import (
	"encoding/json"
	"fmt"
	"go-file-work/internal/model"
	"log"
	"os"
)

func CreateNewFile(name string, testData []model.MyData) {
	f, err = os.Create(name)
	if err != nil {
		log.Fatal(err)
	}
	writeLine("simple hello text!!!")
	f.WriteString("\n\n")

	writeJSONToFile(testData)
	f.WriteString("\n\n")

	writeTableToFile(testData)

	defer func() {
		if err = f.Close(); err != nil {
			log.Fatal(err)
		}
	}()
}

func writeLine(line string) {
	f.WriteString(line)
}

func writeJSONToFile(data []model.MyData) {
	err = json.NewEncoder(f).Encode(&data)
	if err != nil {
		log.Fatal(err)
	}
}

func writeTableToFile(testData []model.MyData) {

	line := "----------------------------------------------------------"
	border := "|"
	nextLine := "\n"

	f.WriteString(line + nextLine) //---

	for _, name := range testData[0].Names() {
		f.WriteString(border)
		res := fixLen(name, 18)
		f.WriteString(res)
	}
	f.WriteString(border + nextLine)
	for _, td := range testData {
		f.WriteString(line + nextLine) //---
		for _, data := range td.Iterator() {
			f.WriteString(border)
			res := fixLen(fmt.Sprintf("%v", data), 18)
			f.WriteString(res)
		}
		f.WriteString(border + nextLine)
	}
	f.WriteString(line) //---
}
