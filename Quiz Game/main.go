package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"flag"
	"strings"
	"time"
	"bufio"
)

func main(){
	filename:= flag.String("csv", "problems.csv", "a csv file in the format of 'question,answer'")
	limit:= flag.Int("limit", 30, "a 30 second limit")

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

	var cor int
	fmt.Printf("yo have %d seond to finish the game\n", *limit)
	fmt.Println("pressn any Enter to start")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	timer := time.NewTimer(time.Duration(*limit) * time.Second)

	for i, record := range records{
		fmt.Printf("Problem #%d: %s = ", i+1, record[0])

		answerch := make(chan string)

		go func ()  {
			var ans string
			fmt.Scanln(&ans)
			answerch <- ans
		}()

		select{
		case <- timer.C:
				fmt.Println("time has ended")
				fmt.Printf("You answered %v of %v\n", cor, len(records))
				return

		case ans := <- answerch :
				if strings.TrimSpace(ans) == strings.TrimSpace(record[1]) {
				cor++
				}
			}
		}
		

	

	fmt.Printf("you have answerd %v of %v", cor, len(records))
}



