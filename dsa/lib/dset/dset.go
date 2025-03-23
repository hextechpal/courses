package dset

import (
	"errors"
)

type Node[K comparable] interface {
	Key() K
}

type dSet[K comparable, T Node[K]] struct {
	data   map[K]T
	parent map[K]K
	rank   map[K]int
}

func NewDSet[K comparable]() *dSet[K, Node[K]] {
	return &dSet[K, Node[K]]{data: make(map[K]Node[K]), parent: make(map[K]K), rank: make(map[K]int)}
}

func (d *dSet[K, T]) MakeSet(el T) error {
	k := el.Key()
	_, ok := d.data[k]
	if ok {
		return errors.New("duplicate key")
	}

	d.data[k] = el
	d.parent[k] = k
	d.rank[k] = 1
	return nil
}

func (d *dSet[K, T]) Find(key K) K {
	for d.parent[key] != key {
		return d.Find(d.parent[key])
	}
	return key
}

func (d *dSet[K, T]) Union(src, dest K) {
	srcId := d.Find(src)
	destId := d.Find(dest)

	if srcId == destId {
		return
	}

	srcRank := d.rank[srcId]
	destRank := d.rank[destId]

	if srcRank < destRank {
		d.parent[srcId] = destId
	} else {
		d.parent[destId] = srcId
		if srcRank == destRank {
			d.rank[srcId] = d.rank[srcId] + 1
		}
	}

}
