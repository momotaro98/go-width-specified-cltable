package wscltable

import (
	"testing"
)

func TestTable(t *testing.T) {
	// inistanciate Table
	table := NewTable([]Column{
		Column{Name: "ID", Width: 5, Alignment: "left"},
		Column{Name: "Title", Width: 20},
		Column{Name: "Tag", Width: 10, Alignment: "center"},
		Column{Name: "Stock", Width: 5, Alignment: "right"},
	})

	// test NewTable func
	if len(table.Columns) != 4 {
		t.Fatalf("Expected: 4; got: %d", len(table.Columns))
	}
	for _, c := range table.Columns {
		if c.Name == "Title" {
			if c.Width != 20 {
				t.Fatalf("Expected: 20; got: %d", c.Width)
			}
		}
	}

	// test getTopLine method
	s := table.getTopLine()
	if s != "┌─────┬────────────────────┬──────────┬─────┐" {
		t.Fatalf("Expected: ┌─────┬────────────────────┬──────────┬─────┐; got:  %s", s)
	}
	// test getColumnNameLine method
	s = table.getColumnNameLine()
	if s != "| ID  |       Title        |   Tag    |Stock|" {
		t.Fatalf("Expected: | ID  |       Title        |   Tag    |Stock|; got: %s", s)
	}
	// test getSeparateLine method
	s = table.getSeparateLine()
	if s != "|─────|────────────────────|──────────|─────|" {
		t.Fatalf("Expected: |─────|────────────────────|──────────|─────|; got: %s", s)
	}

	// test AddRow method
	table.AddRow(map[string]interface{}{"ID": 4, "Title": "golanggolanggolanggolanggolanggolanggolang", "Tag": "python", "Stock": 8})
	for _, row := range table.Rows {
		if len(row["ID"]) != 3 {
			t.Fatalf("Expected: 3; got: %d", len(row["ID"]))
		}
	}
	table.AddRow(map[string]interface{}{"Tag": "cli", "ID": 5, "Title": "TestCase", "Stock": 1})
	table.AddRow(map[string]interface{}{"Tag": "ddd", "ID": 6, "Title": "GOGOGOGOGOGOGOGOGOOGO", "Stock": 16})
	if len(table.Rows) != 3 {
		t.Fatalf("Expected: 3; got: %d", len(table.Rows))
	}

	// test Print method by Visualization
	table.Print()
}

func TestColumn(t *testing.T) {
	// Column instance 01
	// instanciate Column
	column := Column{Name: "ID", Width: 5, Alignment: "left"}

	// test MakeTurnedLinesAndLen method
	line := "13"
	turnedLines, length := column.MakeTurnedLinesAndLen(line)
	if len(turnedLines) != 1 {
		t.Fatalf("Expected: 1; got: %d", len(turnedLines))
	}
	if length != 1 {
		t.Fatalf("Expected: 1; got: %d", length)
	}
	if turnedLines[0] != "13   " {
		t.Fatalf("Expected: '13   '; got: %s", turnedLines[0])
	}

	// test AddEmptyLine method
	addedLines := column.AddEmptyLine(turnedLines, 5)
	if len(addedLines) != 5 {
		t.Fatalf("Expected: 5; got: %d", len(addedLines))
	}

	// Column instance 02
	// Instanciation
	column = Column{Name: "Title", Width: 20, Alignment: "right"}

	// test MakeTurnedLinesAndLen method
	line = "Why don't you learn Go language?" // 32 char
	turnedLines, length = column.MakeTurnedLinesAndLen(line)
	if len(turnedLines) != 2 {
		t.Fatalf("Expected: 2; got: %d", len(turnedLines))
	}
	if length != 2 {
		t.Fatalf("Expected: 2; got: %d", length)
	}
	if turnedLines[0] != "Why don't you learn " {
		t.Fatalf("Expected: 'Why don't you learn '; got: %s", turnedLines[0])
	}
	if turnedLines[1] != "        Go language?" {
		t.Fatalf("Expected: '        Go language?'; got: %s", turnedLines[1])
	}
}
