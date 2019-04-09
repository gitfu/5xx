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
var m map[string]*Site

type Site struct {
	total, errors int
}

func (site *Site) percent() float64 {
	e := float64(site.errors)
	t := float64(site.total)
	p := (e / t) * 100.0
	return p
}

func report(m map[string]*Site) {
	for k, v := range m {
		fmt.Printf(k)
		fmt.Println(" ", v.percent())
	}
}

func do(line string) {
	fu := strings.Split(line, "|")
	hostname := strings.TrimSpace(fu[2])
	if _, ok := m[hostname]; !ok {
		m[hostname] = &Site{
			0, 0,
		}
	}
	floated, _ := strconv.ParseFloat((strings.TrimSpace(fu[0])), 64)
	if etime > floated {
		if floated >= stime {
			m[hostname].total += 1
			if strings.HasPrefix(fu[4], " 5") {
				m[hostname].errors += 1
			}

		}
	}
}

func main() {
	sptr := flag.Float64("s", 0.0, "start time")
	eptr := flag.Float64("e", 9999999999.0, "end time")
	flag.Parse()
	files := flag.Args()
	stime = *sptr
	etime = *eptr
	m = make(map[string]*Site)
	for _, f := range files {
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
			do(scanner.Text())
		}
	}
	report(m)
}
