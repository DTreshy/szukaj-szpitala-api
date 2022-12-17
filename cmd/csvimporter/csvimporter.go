package main

import (
	"context"
	"fmt"
	"io"
	"time"
)

func main() {
	start := time.Now()
	c := newConnection()

	err := c.connect("localhost:55010")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	defer c.conn.Close()

	p, err := openParser("./data.csv")
	if err != nil {
		fmt.Printf("%v\n", err.Error())
		return
	}

	defer p.file.Close()

	var (
		counter    uint32
		readHeader bool
	)

	for {
		rec, err := p.readNextRecord()
		if err != nil {
			printInfo(counter, start)

			if err != io.EOF {
				fmt.Printf("%v\n", err.Error())
			}

			break
		}

		if !readHeader {
			p.header = rec.csvData
			readHeader = true

			continue
		}

		if err := c.sendInsertRequest(context.Background(), rec); err != nil {
			printInfo(counter, start)
			fmt.Printf("%v\n", err.Error())

			return
		}

		counter += 1
	}

	printInfo(counter, start)
}

func printInfo(counter uint32, start time.Time) {
	if counter > 0 {
		fmt.Printf("Inserted %d data items in %v\n", counter, time.Since(start))
	}
}
