package wscltable

import (
	"fmt"
	"log"
	"strings"

	"github.com/moznion/go-unicode-east-asian-width"
)

// Table is struct for printing table
type Table struct {
	Columns []Column
	Rows    []map[string][]string
}

// NewTable creates a new instance of Table
func NewTable(columns []Column) *Table {
	for _, c := range columns {
		if c.Name == "" {
			log.Fatalf("There is Name empty Column!")
		}
		if c.Width < len(c.Name) {
			log.Fatalf("Width %d is too short for %s", c.Width, c.Name)
		}
	}
	return &Table{
		Columns: columns,
		Rows:    make([]map[string][]string, 0),
	}
}

// AddRow adds a row to Table's Rows
func (t *Table) AddRow(row map[string]interface{}) {
	newRow := make(map[string][]string)
	var maxLen int
	for _, c := range t.Columns {
		v, ok := row[c.Name]
		if !ok {
			log.Fatalf("No %s in your AddRow arguments", c.Name)
		}
		val := fmt.Sprintf("%v", v)
		value, vLen := c.MakeTurnedLinesAndLen(val)
		newRow[c.Name] = value
		if vLen > maxLen {
			maxLen = vLen
		}
	}

	for _, c := range t.Columns {
		newRow[c.Name] = c.AddEmptyLine(newRow[c.Name], maxLen)
	}

	if len(newRow) > 0 {
		t.Rows = append(t.Rows, newRow)
	}
}

// Print outputs table
func (t *Table) Print() {
	// print head
	fmt.Println(t.getTopLine())
	fmt.Println(t.getColumnNameLine())
	fmt.Println(t.getSeparateLine())

	// print rows
	for _, row := range t.Rows {
		// get maxInnerLinesLen
		var maxInnerLinesLen int
		for _, c := range t.Columns {
			maxInnerLinesLen = len(row[c.Name])
		}

		// join the columns factor to line
		artInnerLineList := make([][]string, maxInnerLinesLen)
		for _, c := range t.Columns {
			for j, cl := range row[c.Name] {
				artInnerLineList[j] = append(artInnerLineList[j], cl)
			}
		}
		// make ret lines
		for _, ails := range artInnerLineList {
			printRow := "|" + strings.Join(ails, "|") + "|"
			fmt.Println(printRow)
		}
		// print separate line
		fmt.Println(t.getSeparateLine())
	}
}

func (t *Table) getTopLine() string {
	csList := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		csList[i] = strings.Repeat("─", c.Width)
	}
	return "┌" + strings.Join(csList, "┬") + "┐"
}

func (t *Table) getColumnNameLine() string {
	list := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		list[i] = CenterAligned(c.Name, c.Width)
	}
	return "|" + strings.Join(list, "|") + "|"
}

func (t *Table) getSeparateLine() string {
	csList := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		csList[i] = strings.Repeat("─", c.Width)
	}
	return "|" + strings.Join(csList, "|") + "|"
}

// Column is struct to manage Table's Columns
type Column struct {
	Name      string
	Width     int
	Alignment string
}

// MakeTurnedLinesAndLen return lines and lines num
func (c *Column) MakeTurnedLinesAndLen(val string) ([]string, int) {
	lines := c.makeTurnedLine(val)
	length := len(lines)
	return lines, length
}

func (c *Column) makeTurnedLine(str string) (tLines []string) {
	var nextHalfLen int
	var halfLen int
	var curRuneSlice []rune
	var isJustMaxLenFlag bool

	for _, r := range str { // 'r' means rune
		if eastasianwidth.IsFullwidth(r) {
			nextHalfLen = halfLen + 2
		} else {
			nextHalfLen = halfLen + 1
		}

		if nextHalfLen > c.Width {
			// append curRuneSlice to goal slice(tLines)
			if isJustMaxLenFlag {
				tLines = append(tLines, string(curRuneSlice))
			} else {
				tLines = append(tLines, string(curRuneSlice)+" ")
			}
			// reset stat variables and continue
			if eastasianwidth.IsFullwidth(r) {
				halfLen = 2
			} else {
				halfLen = 1
			}
			curRuneSlice = []rune{r}
			isJustMaxLenFlag = false
			continue
		}

		// change stat variables
		halfLen = nextHalfLen
		curRuneSlice = append(curRuneSlice, r)
		if nextHalfLen == c.Width {
			isJustMaxLenFlag = true
		} else {
			isJustMaxLenFlag = false
		}
	}
	// alignment the last curRuneSlice and append it to tLines
	switch c.Alignment {
	case "center":
		tLines = append(tLines, CenterAligned(string(curRuneSlice), c.Width))
	case "right":
		tLines = append(tLines, RightAligned(string(curRuneSlice), c.Width))
	case "left":
		tLines = append(tLines, LeftAligned(string(curRuneSlice), c.Width))
	default:
		tLines = append(tLines, LeftAligned(string(curRuneSlice), c.Width))
	}

	return // return tLines
}

// AddEmptyLine add empty line to each column line
func (c *Column) AddEmptyLine(lines []string, maxLen int) []string {
	if diffLen := maxLen - len(lines); diffLen > 0 {
		for i := 0; i < diffLen; i++ {
			lines = append(lines, strings.Repeat(" ", c.Width))
		}
	}
	return lines
}

// LeftAligned return left aligned string
func LeftAligned(str string, max int) (ret string) {
	restNum := max - len(str)
	ret = str + strings.Repeat(" ", restNum)
	return
}

// CenterAligned return center aligned string
func CenterAligned(str string, max int) (ret string) {
	restNum := max - len(str)
	ret = strings.Repeat(" ", restNum/2) + str + strings.Repeat(" ", restNum-restNum/2)
	return
}

// RightAligned return rignt aligned string
func RightAligned(str string, max int) (ret string) {
	restNum := max - len(str)
	ret = strings.Repeat(" ", restNum) + str
	return
}
