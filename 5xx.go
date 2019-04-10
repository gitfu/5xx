package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lc int
var startTime float64
var endTime float64
var m map[string]*Stats

type Stats struct {
	total, errors int
}

func (stats *Stats) Percentage() float64 {
	e := float64(stats.errors)
	t := float64(stats.total)
	p := (e / t) * 100.0
	return p
}

func Reporter(m map[string]*Stats) {
	fmt.Printf("Between time %.2f and time  %.2f \n", startTime, endTime)
	for k, v := range m {
		fmt.Printf("%s returned %.2f%% 500 errors\n", k, v.Percentage())
	}
}

func LogParser(log string) {
	fmt.Println(log)
	file, err := os.Open(log)
	if err != nil {
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		LineParser(scanner.Text())
	}

}
func HostChecker(hostname string) {
	if _, ok := m[hostname]; !ok {
		m[hostname] = &Stats{
			0, 0,
		}
	}
}

func TimeChecker(timestamp float64) int {
	if endTime > timestamp {
		if timestamp >= startTime {
			return 1
		}
	}
	return 0
}

func HttpCodeChecker(httpcode string) int {
	if strings.HasPrefix(httpcode, "5") {
		return 1
	}
	return 0
}

func LineParser(line string) {
	lc += 1
	values := strings.Split(line, "|")
	hostname := strings.TrimSpace(values[2])
	HostChecker(hostname)
	timestamp, _ := strconv.ParseFloat((strings.TrimSpace(values[0])), 64)
	m[hostname].total += TimeChecker(timestamp)
	httpcode := strings.TrimSpace(values[4])
	m[hostname].errors += HttpCodeChecker(httpcode)
}

func ArgParser() []string {
	startPtr := flag.Float64("s", 0.0, "start time")
	endPtr := flag.Float64("e", 9999999999.0, "end time")
	flag.Parse()
	files := flag.Args()
	startTime = *startPtr
	endTime = *endPtr
	return files
}

func main() {
	lc = 0
	logFiles := ArgParser()
	m = make(map[string]*Stats)

	for _, log := range logFiles {
		LogParser(log)
		fmt.Println(lc)
	}
	Reporter(m)
}
