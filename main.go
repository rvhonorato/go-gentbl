// Generate a TBL file to be used as restraints in HADDOCK
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// Main function
func main() {
	inputfile := os.Args[1]
	lines := readFile(inputfile)
	fmt.Println(GenTBL(lines))
}

// Read the input file and return a slice of strings
func readFile(p string) []string {
	f, err := os.Open(p)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	scanner := bufio.NewScanner(f)

	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return cleanLines(lines)
}

// Skip comments and empty lines
func cleanLines(s []string) []string {
	commentPrefix := "#"
	var cleanLines []string
	for _, l := range s {
		if !strings.HasPrefix(l, commentPrefix) && l != "" {
			cleanLines = append(cleanLines, l)
		}
	}
	return cleanLines
}

// Parse the input array and return a map of restraints
func parseInput(s []string) (map[int][]string, map[int][]string, map[int]string) {
	actM := make(map[int][]string)
	pasM := make(map[int][]string)
	cM := make(map[int]string)
	for id, l := range s {
		words := strings.Split(l, " ")
		chain := words[0]
		act := strings.Split(words[1], ",")
		pas := strings.Split(words[2], ",")
		actM[id+1] = act
		pasM[id+1] = pas
		cM[id+1] = chain
	}
	return actM, pasM, cM

}

// GenTBL generates the TBL file
func GenTBL(s []string) string {
	actM, pasM, cM := parseInput(s)

	var result string
	for id, l := range cleanLines(s) {
		idA := id + 1
		var tbl string

		// Put a nice header at the top
		if id == 0 {
			tbl = "! HADDOCK AIR restraints\n"
			tbl += header(idA)
		} else {
			tbl = header(idA)
		}
		words := strings.Split(l, " ")

		actA := strings.Split(words[1], ",")
		chainA := cM[idA]
		partners := strings.Split(words[3], ",")

		for _, a := range actA {
			if a == "" {
				break
			}

			passive := make([]string, 0)
			appendComment := false
			for _, p := range partners {
				pident, _ := strconv.Atoi(p)

				if pident == idA {
					appendComment = true
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

			tbl += formatTBL(resA, chainA, passive, appendComment)
		}
		// fmt.Println(tbl)
		result += tbl
	}
	return result
}

// Put the restraints in the CNS format
func formatTBL(actR int, actC string, pa []string, comment bool) string {
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
			str += "       )  2.0 2.0 0.0\n\n"
		}
	}
	if comment {
		str = addComments(str)
	}
	return str
}

// Add comments to a string
func addComments(str string) string {
	nStr := strings.Repeat("!", 42) + "\n"
	nStr += "!!! Self-restraint not allowed !!!\n"
	nStr += strings.Repeat("!", 42) + "\n"
	for _, l := range strings.Split(str, "\n") {
		nStr += fmt.Sprintf("! %s\n", l)
	}
	nStr += strings.Repeat("!", 42) + "\n"
	return nStr
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
	place, _ := strconv.Atoi(n[len(n)-1:])
	return place
}
