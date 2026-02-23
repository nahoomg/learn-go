package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"flag"
	"strings"
)

func main(){
	filename:= flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	flag.Parse()

	file, err := os.Open(*filename)
	if err != nil{
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	records, err := reader.ReadAll()

	if err != nil{
		panic(err)
	}


	var ans string
	var cor int
	for i, record := range records{
		fmt.Printf("Problem #%d: %s = ", i+1, record[0])
		fmt.Scanln(&ans)

		if ans == strings.TrimSpace(record[1]) {
			cor++
		}
	}
	fmt.Printf("you have answerd %v of %v", cor, len(records))

}

