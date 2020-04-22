package common

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"downsampling"
)

func LoadPointsFromCSV(file string) []downsampling.Point {
	csvFile, err := os.Open(file)
	CheckError("Cannot Open the file.", err)
	reader := csv.NewReader(bufio.NewReader(csvFile))

	var data []downsampling.Point
	for {
		line, error := reader.Read()
		if error == io.EOF {
			break
		}
		CheckError("Read file error", err)
		var d downsampling.Point
		d.X, _ = strconv.ParseFloat(line[0], 64)
		d.Y, _ = strconv.ParseFloat(line[1], 64)
		data = append(data, d)
	}
	return data
}

func SavePointsToCSV(file string, points []downsampling.Point) {
	fp, err := os.Create(file)
	CheckError("Cannot create file", err)
	defer fp.Close()

	writer := csv.NewWriter(fp)
	defer writer.Flush()

	for _, point := range points {
		x := fmt.Sprintf("%f", point.X)
		y := fmt.Sprintf("%f", point.Y)
		err := writer.Write([]string{x, y})
		CheckError("Cannot write to file", err)
	}
}
