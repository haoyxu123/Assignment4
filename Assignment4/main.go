package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"math"
	"os"
	"sort"
	"strconv"
	"time"
)

// Statistic holds statistical data
type Statistic struct {
	Count  int
	Mean   float64
	StdDev float64
	Min    float64
	Q1     float64
	Median float64
	Q3     float64
	Max    float64
}

// calculateStatistics calculates descriptive statistics for a slice of floats
func calculateStatistics(data []float64) Statistic {
	sort.Float64s(data) // used for  for calculating the median and quartiles correctly
	count := len(data)
	min := data[0]
	max := data[count-1]
	sum := 0.0   // 0.0 is define the sum as float and initiliza the sum starts at 0
	sqSum := 0.0 // For standard deviation

	for _, value := range data { //"_ " ignore the index, only needs the elements of the dataset
		sum += value
		sqSum += value * value
	}

	mean := sum / float64(count)
	variance := (sqSum / float64(count)) - (mean * mean)
	stdDev := math.Sqrt(variance)

	median := percentile(data, 0.5)
	q1 := percentile(data, 0.25)
	q3 := percentile(data, 0.75)

	return Statistic{
		Count:  count,
		Mean:   mean,
		StdDev: stdDev,
		Min:    min,
		Q1:     q1,
		Median: median,
		Q3:     q3,
		Max:    max,
	}
}

// percentile calculates the p-th percentile of a sorted data slice
func percentile(data []float64, p float64) float64 {
	if p <= 0 {
		return data[0]
	}
	if p >= 1 {
		return data[len(data)-1]
	}

	position := (float64(len(data)) - 1) * p
	lower := math.Floor(position)
	upper := math.Ceil(position)
	if lower == upper {
		return data[int(position)]
	}
	lowerValue := data[int(lower)]
	upperValue := data[int(upper)]
	return lowerValue + (upperValue-lowerValue)*(position-lower)
}

// describeColumn calculates descriptive statistics for a given column in the CSV
func describeColumn(column []string) (Statistic, error) {
	var data []float64
	for _, value := range column {
		floatValue, err := strconv.ParseFloat(value, 64)
		if err != nil {
			return Statistic{}, err
		}
		data = append(data, floatValue)
	}
	return calculateStatistics(data), nil
}

func main() {
	startTime := time.Now()
	// Open the CSV file
	csvFile, err := os.Open("C:/Assignment4/housesInput.csv")
	if err != nil {
		fmt.Println("Error opening CSV file:", err)
		return
	}
	defer csvFile.Close()

	// Create a CSV reader
	reader := csv.NewReader(csvFile)

	// Read the header row
	headers, err := reader.Read()
	if err != nil {
		fmt.Println("Error reading CSV header:", err)
		return
	}

	// Read the rest of the data
	var columns [][]string
	for i := 0; i < len(headers); i++ {
		columns = append(columns, []string{})
	}
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error reading CSV record:", err)
			return
		}
		for i, value := range record {
			columns[i] = append(columns[i], value)
		}
	}

	// Open or create the output file in append mode
	outputFile, err := os.OpenFile("C:/Assignment4/example3.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Println("Error opening output file:", err)
		return
	}
	defer outputFile.Close()

	// Write the descriptive statistics to the output file
	for i, header := range headers {
		stat, err := describeColumn(columns[i])
		if err != nil {
			fmt.Printf("Error describing column %s: %v\n", header, err)
			continue
		}

		output := fmt.Sprintf("Column: %s\nCount: %d\nMean: %.2f\nStdDev: %.2f\nMin: %.2f\n25th percentile (Q1): %.2f\nMedian: %.2f\n75th percentile (Q3): %.2f\nMax: %.2f\n\n",
			header, stat.Count, stat.Mean, stat.StdDev, stat.Min, stat.Q1, stat.Median, stat.Q3, stat.Max)
		if _, err := outputFile.WriteString(output); err != nil {
			fmt.Println("Error writing to output file:", err)
		}
	}
	elapsed := time.Now().Sub(startTime)
	fmt.Printf("Execution took %d nanoseconds\n", elapsed)
}
