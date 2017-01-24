package wscltable

import (
	"fmt"
	"testing"
)

func TestColumn(t *testing.T) {
	// inistanciate Table
	table := NewTable([]Column{
		Column{Name: "ID", Width: 5, Alignment: "left"},
		Column{Name: "Title", Width: 20, Alignment: "left"},
		Column{Name: "Stock", Width: 5, Alignment: "right"},
	})

	for _, c := range table.Columns {
		if c.Name == "Title" {
			if c.Width != 20 {
				t.Fatalf("wrong!!! c.Width: %d", c.Width)
			}
		}
	}

	s := table.getTopLine()
	if s != "┌─────┬────────────────────┬─────┐" {
		t.Fatalf("wrong!!! s: %s", s)
	}

	s = table.getSeparateLine()
	if s != "|─────|────────────────────|─────|" {
		t.Fatalf("wrong!!! s: %s", s)
	}

	// Test AddRow method
	table.AddRow(map[string]interface{}{"ID": 4, "Title": "golanggolanggolanggolanggolanggolanggolang", "Stock": 8})
	for _, row := range table.Rows {
		fmt.Println(len(row["ID"]))
		if len(row["ID"]) != 3 {
			t.Fatalf("wrong!!!: %d", len(row["ID"]))
		}
	}

	/*
		table.AddRow(map[string]interface{}{"ID": 5, "Title": "GOGOGOGOGOGOGOGOGOOGO", "Stock": 16})
		table.Print()
	*/
}
