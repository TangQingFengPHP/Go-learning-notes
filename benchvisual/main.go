package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"regexp"
)

type BenchmarkResult struct {
	Name        string  `json:"name"`
	Runs        int     `json:"runs"`
	NsPerOp     float64 `json:"ns_per_op"`
	BytesPerOp  int     `json:"bytes_per_op"`
	AllocsPerOp int     `json:"allocs_per_op"`
}

var results []BenchmarkResult

func parseBenchmark(input *os.File) ([]BenchmarkResult, error) {
	re := regexp.MustCompile(`^(Benchmark\S*)\s+(\d+)\s+([\d\.]+)\s+ns/op\s+(\d+)\s+B/op\s+(\d+)\s+allocs/op`)
	scanner := bufio.NewScanner(input)
	var parsed []BenchmarkResult

	for scanner.Scan() {
		line := scanner.Text()
		matches := re.FindStringSubmatch(line)
		if len(matches) == 6 {
			var r BenchmarkResult
			r.Name = matches[1]
			fmt.Sscanf(matches[2], "%d", &r.Runs)
			fmt.Sscanf(matches[3], "%f", &r.NsPerOp)
			fmt.Sscanf(matches[4], "%d", &r.BytesPerOp)
			fmt.Sscanf(matches[5], "%d", &r.AllocsPerOp)
			parsed = append(parsed, r)
		}
	}

	return parsed, scanner.Err()
}

func dataHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(results)
}

func main() {
	// 判断是否有标准输入
	stat, _ := os.Stdin.Stat()

	if (stat.Mode() & os.ModeCharDevice) == 0 {
		// 从标准输入读取
		res, err := parseBenchmark(os.Stdin)
		fmt.Println(res)
		if err != nil {
			log.Fatal(err)
		}
		results = res
	} else {
		// 默认读取 benchmark.txt 文件
		file, err := os.Open("benchmark.txt")
		if err != nil {
			fmt.Println("打开 benchmark.txt 失败", err)
			return
		}
		defer file.Close()
		res, err := parseBenchmark(file)
		results = res
		fmt.Println(results)
	}

	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/", fs)
	http.HandleFunc("/data", dataHandler)

	fmt.Println("🚀 服务器已启动: http://localhost:8080/")
	http.ListenAndServe(":8080", nil)

}
