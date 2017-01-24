package wscltable

import (
	// "fmt"
	"testing"
)

func TestColumn(t *testing.T) {
	// inistanciate Table
	table := NewTable([]Column{
		Column{Name: "Title", Width: 5, Alignment: "left"},
		Column{Name: "Article", Width: 20, Alignment: "left"},
		Column{Name: "Stock", Width: 5, Alignment: "right"},
	})

	for _, c := range table.Columns {
		if c.Name == "Article" {
			if c.Width != 20 {
				t.Fatalf("wrong!!! c.Width: %d", c.Width)
			}
		}
	}

	s := table.getTopLine()
	if s != "┌─────┬────────────────────┬─────┐" {
		t.Fatalf("wrong!!! s: %s", s)
	}

	s = table.getColumnNameLine()
	if s != "|─────|────────────────────|─────|" {
		t.Fatalf("wrong!!! s: %s", s)
	}
}
