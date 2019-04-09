package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var lc int
var stime float64
var etime float64
var hosts []string
var errors []int
var totals []int

type Site struct {
	total, errors int
}

var m map[string]*Site

func report(m map[string]*Site) {
	for k, _ := range m {
		fmt.Printf(k)
		e := float64(m[k].errors)
		t := float64(m[k].total)
		p := (e / t) * 100.0
		fmt.Println(" ", p)

	}

	fmt.Println(lc)
}

func do(line string) {
	stime := 1493969100.0
	etime := 1493969102.0
	lc += 1
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
	m = make(map[string]*Site)
	lc = 0
	files := []string{"20.data", "20a.data"}
	//files := []string{"10m.data", "10ma.data", "fat.data", "out.data", "out1.data", "out2.data", "out3.data"}
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
		report(m)
	}

}
