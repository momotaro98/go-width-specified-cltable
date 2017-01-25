package wscltable

import (
	"fmt"
	"log"
	"strings"

	"github.com/moznion/go-unicode-east-asian-width"
)

type Table struct {
	Columns []Column
	Rows    []map[string][]string
}

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

// Column
type Column struct {
	Name      string
	Width     int
	Alignment string
}

func (c *Column) MakeTurnedLinesAndLen(val string) ([]string, int) {
	lines := c.makeTurnedLine(val)
	length := len(lines)
	return lines, length
}

func (c *Column) makeTurnedLine(str string) (t_lines []string) {
	var isJustMaxLenFlag bool
	var cur_half_len int
	var cur_line []rune

	for _, r := range str { // 'r' means rune
		if eastasianwidth.IsFullwidth(r) {
			cur_half_len += 2
		} else {
			cur_half_len++
		}
		if cur_half_len == c.Width {
			isJustMaxLenFlag = true
		} else if cur_half_len > c.Width {
			// Arrange to Full and Half
			if isJustMaxLenFlag == true {
				t_lines = append(t_lines, string(cur_line))
			} else {
				t_lines = append(t_lines, string(cur_line)+" ")
			}
			// Initalize stat variables
			if eastasianwidth.IsFullwidth(r) {
				cur_half_len = 2
			} else {
				cur_half_len = 1
			}
			isJustMaxLenFlag = false
			cur_line = nil
		}
		cur_line = append(cur_line, r)
	}
	switch c.Alignment {
	case "center":
		t_lines = append(t_lines, CenterAligned(string(cur_line), c.Width))
	case "right":
		t_lines = append(t_lines, RightAligned(string(cur_line), c.Width))
	case "left":
		t_lines = append(t_lines, LeftAligned(string(cur_line), c.Width))
	default:
		t_lines = append(t_lines, LeftAligned(string(cur_line), c.Width))
	}
	return
}

func (c *Column) AddEmptyLine(lines []string, maxLen int) []string {
	if diffLen := maxLen - len(lines); diffLen > 0 {
		for i := 0; i < diffLen; i++ {
			lines = append(lines, strings.Repeat(" ", c.Width))
		}
	}
	return lines
}

// util func
func LeftAligned(str string, max int) (ret string) {
	rest_num := max - len(str)
	ret = str + strings.Repeat(" ", rest_num)
	return
}

// util func
func CenterAligned(str string, max int) (ret string) {
	rest_num := max - len(str)
	ret = strings.Repeat(" ", rest_num/2) + str + strings.Repeat(" ", rest_num-rest_num/2)
	return
}

// util func
func RightAligned(str string, max int) (ret string) {
	rest_num := max - len(str)
	ret = strings.Repeat(" ", rest_num) + str
	return
}
