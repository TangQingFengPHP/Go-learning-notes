package main

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"
	"regexp"
)

type BenchmarkResult struct {
	Name        string
	Runs        string
	NsPerOp     string
	BytesPerOp  string
	AllocsPerOp string
}

func main() {
	inpuFile := "benchmark.txt"
	outputFile := "benchmark.csv"

	file, err := os.Open(inpuFile)
	if err != nil {
		fmt.Println("读取文件失败", err)
		return
	}
	defer file.Close()

	// 正则匹配 benchmark 输出行
	re := regexp.MustCompile(`^Benchmark(\w+)-\d+\s+(\d+)\s+(\d+)\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op`)

	var results []BenchmarkResult

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 6 {
			results = append(results, BenchmarkResult{
				Name:        matches[1],
				Runs:        matches[2],
				NsPerOp:     matches[3],
				BytesPerOp:  matches[4],
				AllocsPerOp: matches[5],
			})
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("读取行失败", err)
	}

	// 写入 CSV 文件
	outFile, err := os.Create(outputFile)
	if err != nil {
		fmt.Println("创建 CSV 文件失败", err)
		return
	}
	defer outFile.Close()

	writer := csv.NewWriter(outFile)
	defer writer.Flush()

	header := []string{"name", "runs", "ns_per_op", "bytes_per_op", "allocs_per_op"}
	if err := writer.Write(header); err != nil {
		fmt.Println("写入 CSV 头失败", err)
		return
	}

	for _, r := range results {
		row := []string{r.Name, r.Runs, r.NsPerOp, r.BytesPerOp, r.AllocsPerOp}
		if err := writer.Write(row); err != nil {
			fmt.Println("写入 CSV 行失败", err)
			return
		}
	}

	fmt.Println("写入 CSV 文件成功", outputFile)
}
