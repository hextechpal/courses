package mergetables

import "fmt"

type Table struct {
	ID   int
	Rows int
}

type info struct {
	totalRows int
	rank      int
	parentId  int
}

type TableMerger struct {
	parents []*info
	maxRows int
}

func NewTableMerger(tables []Table) *TableMerger {
	tm := &TableMerger{parents: make([]*info, len(tables))}
	for _, tb := range tables {
		tm.MakeSet(tb)
	}
	return tm
}

func (t *TableMerger) MakeSet(table Table) {
	t.parents[table.ID] = &info{
		totalRows: table.Rows,
		rank:      0,
		parentId:  table.ID,
	}
	if t.maxRows < table.Rows {
		t.maxRows = table.Rows
	}
}

func (t *TableMerger) Find(tableId int) int {
	for t.parents[tableId].parentId != tableId {
		tableId = t.parents[tableId].parentId
	}
	return tableId
}

func (t *TableMerger) Union(dest, src int) {
	srcId := t.Find(src)
	destId := t.Find(dest)

	if srcId == destId {
		return
	}

	if t.parents[srcId].rank < t.parents[destId].rank {
		t.parents[srcId].parentId = destId
		t.parents[destId].totalRows += t.parents[srcId].totalRows
		if t.maxRows < t.parents[destId].totalRows {
			t.maxRows = t.parents[destId].totalRows
		}
	} else {
		t.parents[destId].parentId = srcId
		t.parents[srcId].totalRows += t.parents[destId].totalRows
		if t.maxRows < t.parents[srcId].totalRows {
			t.maxRows = t.parents[srcId].totalRows
		}
		if t.parents[destId].rank == t.parents[srcId].rank {
			t.parents[srcId].rank++
		}
	}
}

func (t *TableMerger) PrintMaxRows() {
	fmt.Println(t.maxRows)
}
