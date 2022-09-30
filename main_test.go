package main

import (
	"strings"
	"testing"
)

func TestReadFile(t *testing.T) {
	// Read the file
	s := readFile("examples/input.txt")
	// Check the number of lines
	if len(s) != 3 {
		t.Errorf("Expected 3 lines, got %d", len(s))
	}
}

func TestCleanLines(t *testing.T) {
	dirtyLines := []string{"# This is a comment", "this is not a comment", ""}
	cleanLines := cleanLines(dirtyLines)
	if len(cleanLines) != 1 {
		t.Errorf("Expected 1 line, got %d", len(cleanLines))
	}
}

func TestParseInput(t *testing.T) {
	l := []string{"A 10,11 20 2", "B 11 12,13 1", "C 10,11,22 20 2"}
	actM, pasM, cM := parseInput(l)
	if len(actM) != 3 {
		t.Errorf("Expected 3 selections, got %d", len(actM))
	}

	if actM[1][0] != "10" {
		t.Errorf("Expected 10, got %s", actM[1][0])
	}
	if actM[1][1] != "11" {
		t.Errorf("Expected 11, got %s", actM[1][1])
	}
	if actM[2][0] != "11" {
		t.Errorf("Expected 11, got %s", actM[2][0])
	}
	if actM[3][0] != "10" {
		t.Errorf("Expected 10, got %s", actM[3][0])
	}
	if actM[3][1] != "11" {
		t.Errorf("Expected 11, got %s", actM[3][1])
	}
	if actM[3][2] != "22" {
		t.Errorf("Expected 22, got %s", actM[3][2])
	}

	if pasM[1][0] != "20" {
		t.Errorf("Expected 20, got %s", pasM[0][0])
	}
	if pasM[2][0] != "12" {
		t.Errorf("Expected 12, got %s", pasM[1][0])
	}
	if pasM[2][1] != "13" {
		t.Errorf("Expected 13, got %s", pasM[1][1])
	}

	if pasM[3][0] != "20" {
		t.Errorf("Expected 20, got %s", pasM[2][0])
	}

	if cM[1] != "A" {
		t.Errorf("Expected A, got %s", cM[1])
	}
	if cM[2] != "B" {
		t.Errorf("Expected B, got %s", cM[2])
	}
	if cM[3] != "C" {
		t.Errorf("Expected C, got %s", cM[3])
	}
}

func TestGenTBL(t *testing.T) {
	l := []string{"A 1  2", "B 11 12 1"}
	result := GenTBL(l)
	lA := strings.Split(result, "\n")

	line1 := "! HADDOCK AIR restraints"
	line2 := "! HADDOCK AIR restraints for 1st selection"
	line3 := "!"
	line4 := "assign ( resid 1 and segid A)"
	line5 := "       ("
	line6 := "        ( resid 11  and segid B)"
	line7 := "     or"
	line8 := "        ( resid 12  and segid B)"
	line9 := "       )  2.0 2.0 0.0"
	line10 := ""
	line11 := "! HADDOCK AIR restraints for 2nd selection"
	line12 := "!"
	line13 := "assign ( resid 11 and segid B)"
	line14 := "       ("
	line15 := "        ( resid 1  and segid A)"
	line16 := "       )  2.0 2.0 0.0"
	line17 := ""
	line18 := ""

	if lA[0] != line1 {
		t.Errorf("Expected %s, got %s", line1, lA[0])
	}
	if lA[1] != line2 {
		t.Errorf("Expected %s, got %s", line2, lA[1])
	}
	if lA[2] != line3 {
		t.Errorf("Expected %s, got %s", line3, lA[2])
	}
	if lA[3] != line4 {
		t.Errorf("Expected %s, got %s", line4, lA[3])
	}
	if lA[4] != line5 {
		t.Errorf("Expected %s, got %s", line5, lA[4])
	}
	if lA[5] != line6 {
		t.Errorf("Expected %s, got %s", line6, lA[5])
	}
	if lA[6] != line7 {
		t.Errorf("Expected %s, got %s", line7, lA[6])
	}
	if lA[7] != line8 {
		t.Errorf("Expected %s, got %s", line8, lA[7])
	}
	if lA[8] != line9 {
		t.Errorf("Expected %s, got %s", line9, lA[8])
	}
	if lA[9] != line10 {
		t.Errorf("Expected %s, got %s", line10, lA[9])
	}
	if lA[10] != line11 {
		t.Errorf("Expected %s, got %s", line11, lA[10])
	}
	if lA[11] != line12 {
		t.Errorf("Expected %s, got %s", line12, lA[11])
	}
	if lA[12] != line13 {
		t.Errorf("Expected %s, got %s", line13, lA[12])
	}
	if lA[13] != line14 {
		t.Errorf("Expected %s, got %s", line14, lA[13])
	}
	if lA[14] != line15 {
		t.Errorf("Expected %s, got %s", line15, lA[14])
	}
	if lA[15] != line16 {
		t.Errorf("Expected %s, got %s", line16, lA[15])
	}
	if lA[16] != line17 {
		t.Errorf("Expected %s, got %s", line17, lA[16])
	}
	if lA[17] != line18 {
		t.Errorf("Expected %s, got %s", line18, lA[17])
	}

}

func TestFormatTBL(t *testing.T) {
	var passive []string
	var result string
	var lA []string

	passive = []string{"111B", "222B", "333B"}
	result = formatTBL(10, "A", passive, false)
	lA = strings.Split(result, "\n")

	lineA := "assign ( resid 10 and segid A)"
	lineB := "        ( resid 111  and segid B)"
	lineC := "        ( resid 222  and segid B)"
	lineD := "        ( resid 333  and segid B)"
	orline := "     or"
	secondline := "       ("
	lastline := "       )  2.0 2.0 0.0"
	if lA[0] != lineA {
		t.Errorf("Expected %s, got %s", lineA, lA[0])
	}
	if lA[1] != secondline {
		t.Errorf("Expected %s, got %s", secondline, lA[1])
	}
	if lA[2] != lineB {
		t.Errorf("Expected %s, got %s", orline, lA[2])
	}
	if lA[3] != orline {
		t.Errorf("Expected %s, got %s", orline, lA[3])
	}
	if lA[4] != lineC {
		t.Errorf("Expected %s, got %s", orline, lA[4])
	}
	if lA[5] != orline {
		t.Errorf("Expected %s, got %s", orline, lA[5])
	}
	if lA[6] != lineD {
		t.Errorf("Expected %s, got %s", orline, lA[6])
	}
	if lA[7] != lastline {
		t.Errorf("Expected %s, got %s", lastline, lA[7])
	}

	passive = []string{"1B"}
	result = formatTBL(10, "A", passive, true)
	lA = strings.Split(result, "\n")
	expected := "! assign ( resid 10 and segid A)"
	if lA[3] != expected {
		t.Errorf("Expected %s, got %s", expected, lA[3])
	}
}

func TestAddComments(t *testing.T) {
	l := addComments("i'm commented")
	lA := strings.Split(l, "\n")
	expected := "! i'm commented"
	if lA[3] != expected {
		t.Errorf("Expected %s, got %s", expected, lA[3])
	}
}

func TestHeader(t *testing.T) {
	l := header(1)
	expected := "! HADDOCK AIR restraints for 1st selection\n!\n"
	if l != expected {
		t.Errorf("Expected %s, got %s", expected, l)
	}
}

func TestGetSuffix(t *testing.T) {
	var val string

	val = getSuffix(1)
	if val != "st" {
		t.Errorf("Expected st, got %s", val)
	}

	val = getSuffix(2)
	if val != "nd" {
		t.Errorf("Expected nd, got %s", val)
	}

	val = getSuffix(3)
	if val != "rd" {
		t.Errorf("Expected rd, got %s", val)
	}

	val = getSuffix(4)
	if val != "th" {
		t.Errorf("Expected th, got %s", val)
	}

}

func TestLastDigit(t *testing.T) {
	var val int

	val = lastDigit(1)
	if val != 1 {
		t.Errorf("Expected 1, got %d", val)
	}

	val = lastDigit(2)
	if val != 2 {
		t.Errorf("Expected 2, got %d", val)
	}

	val = lastDigit(10)
	if val != 0 {
		t.Errorf("Expected 0, got %d", val)
	}

	val = lastDigit(78)
	if val != 8 {
		t.Errorf("Expected 8, got %d", val)
	}
}
