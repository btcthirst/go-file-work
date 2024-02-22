package main

import (
	"go-file-work/internal/model"
	"go-file-work/internal/reader"
	"go-file-work/internal/writer"
)

var testData = []model.MyData{
	{
		Name:  "tester1",
		Place: 1,
		Price: 152,
	},
	{
		Name:  "tester12",
		Place: 11,
		Price: 12,
	},
	{
		Name:  "tester13",
		Place: 12,
		Price: 112,
	},
	{
		Name:  "tester145",
		Place: 5,
		Price: 52,
	},
}

func main() {
	name := "test.txt"
	writer.CreateNewFile(name, testData)

	reader.ReadFile(name)
	reader.ReadToConsole(name)
}
