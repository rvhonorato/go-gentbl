package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

// Check for errors
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	lines := readFile("examples/input.txt")
	GenTBL(lines)
}

// ----
// 1 A 10,20,30 11 2
// 2 B 100,200,300  1
// ----
// Read the input file and return a slice of strings
func readFile(p string) []string {
	f, err := os.Open(p)
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		// fmt.Println(scanner.Text())
		lines = append(lines, scanner.Text())
	}
	return lines
}

// Generate the TBL file
func GenTBL(s []string) {
	actM := make(map[int][]string)
	pasM := make(map[int][]string)
	cM := make(map[int]string)
	for _, l := range s {
		words := strings.Split(l, " ")
		id, _ := strconv.Atoi(words[0])
		chain := words[1]
		act := strings.Split(words[2], ",")
		pas := strings.Split(words[3], ",")
		actM[id] = act
		pasM[id] = pas
		cM[id] = chain
	}

	for _, l := range s {
		words := strings.Split(l, " ")
		idA, _ := strconv.Atoi(words[0])
		actA := strings.Split(words[2], ",")
		chainA := cM[idA]
		partners := strings.Split(words[4], ",")
		tbl := header(idA)
		for _, a := range actA {
			if a == "" {
				break
			}

			passive := make([]string, 0)

			for _, p := range partners {
				pident, _ := strconv.Atoi(p)

				if pident == idA {
					log.Fatal("Self-restraint not possible")
				}

				chainB := cM[pident]
				actB := actM[pident]

				for _, b := range actB {
					if b != "" {
						passive = append(passive, b+chainB)
					}
				}

				pasB := pasM[pident]
				for _, b := range pasB {
					if b != "" {
						passive = append(passive, b+chainB)
					}
				}
			}

			if len(passive) == 0 {
				break
			}

			resA, _ := strconv.Atoi(a)

			tbl += formatTBL(resA, chainA, passive)
		}
		fmt.Println(tbl)
	}
}

// Put the restraints in the CNS format
func formatTBL(actR int, actC string, pa []string) string {
	str := fmt.Sprintf("assign ( resid %d and segid %s)\n", actR, actC)
	str += "       (\n"
	for i, p := range pa {
		// Last char is the chain
		pasR, _ := strconv.Atoi(p[:len(p)-1])
		pasC := p[len(p)-1:]
		str += fmt.Sprintf("        ( resid %d  and segid %s)\n", pasR, pasC)
		if i != len(pa)-1 {
			// last one
			str += "     or\n"
		} else {
			str += "       )  2.0 2.0 0.0\n\n\n"
		}
	}
	return str
}

// Format the header for a selection
func header(id int) string {
	suffix := getSuffix(id)
	return fmt.Sprintf("! HADDOCK AIR restraints for %d%s selection\n!\n", id, suffix)
}

// Get the suffix (st, nd, rd, th) for an integer
func getSuffix(id int) string {
	digit := lastDigit(id)
	if digit == 1 {
		return "st"
	} else if digit == 2 {
		return "nd"
	} else if digit == 3 {
		return "rd"
	} else {
		return "th"
	}
}

// Get the last digit of a number
func lastDigit(num int) int {
	n := strconv.FormatInt(int64(num), 10)
	place, _ := strconv.Atoi(n[:len(n)-1])
	if place == 0 {
		return num
	} else {
		r := num % int(math.Pow(10, float64(place)))
		return r / int(math.Pow(10, float64(place-1)))
	}
}
