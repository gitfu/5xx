package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var stime float64
var etime float64
var m map[string]*Stats

type Stats struct {
	total, errors int
}

func (stats *Stats) percent() float64 {
	e := float64(stats.errors)
	t := float64(stats.total)
	p := (e / t) * 100.0
	return p
}

func report(m map[string]*Stats) {
	fmt.Printf("Between time %.2f and time  %.2f \n", stime, etime)
	for k, v := range m {
		fmt.Printf("%s returned %.2f%% 500 errors\n", k, v.percent())
	}
}

func parseLog(f string) {
	fmt.Println(f)
	file, err := os.Open(f)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		parseLine(scanner.Text())
	}
}

func checkHost(hostname string) {
	if _, ok := m[hostname]; !ok {
		m[hostname] = &Stats{
			0, 0,
		}
	}
}

func checkTime(timestamp float64) int {
	if etime > timestamp {
		if timestamp >= stime {
			return 1
		}
	}
	return 0
}

func checkHttpCode(httpcode string) int {
	if strings.HasPrefix(httpcode, " 5") {
		return 1
	}
	return 0
}

func parseLine(line string) {
	values := strings.Split(line, "|")
	hostname := strings.TrimSpace(values[2])
	checkHost(hostname)
	timestamp, _ := strconv.ParseFloat((strings.TrimSpace(values[0])), 64)
	m[hostname].total += checkTime(timestamp)
	httpcode := values[4]
	m[hostname].errors += checkHttpCode(httpcode)
}

func parseArgs() []string {
	sptr := flag.Float64("s", 0.0, "start time")
	eptr := flag.Float64("e", 9999999999.0, "end time")
	flag.Parse()
	files := flag.Args()
	stime = *sptr
	etime = *eptr
	return files
}

func main() {
	files := parseArgs()
	m = make(map[string]*Stats)
	for _, f := range files {
		parseLog(f)
	}
	report(m)
}
