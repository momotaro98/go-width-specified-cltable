package wscltable

import (
	"fmt"
	"strings"
)

type Table struct {
	Columns []Column
	Rows    []map[string][]string
}

func NewTable(columns []Column) *Table {
	return &Table{
		Columns: columns,
		Rows:    make([]map[string][]string, 0),
	}
}

func (t *Table) Print() {
	// print head
	fmt.Println(t.getTopLine())
	fmt.Println(t.getColumnNameLine())
}

func (t *Table) getTopLine() string {
	csList := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		csList[i] = c.makeFilledInStr("─")
	}
	return "┌" + strings.Join(csList, "┬") + "┐"
}

func (t *Table) getColumnNameLine() string {
	csList := make([]string, len(t.Columns))
	for i, c := range t.Columns {
		csList[i] = c.makeFilledInStr("─")
	}
	return "|" + strings.Join(csList, "|") + "|"
}

type Column struct {
	Name      string
	Width     int
	Alignment string
}

func (c *Column) makeFilledInStr(char string) string {
	return strings.Repeat(char, c.Width)
}
