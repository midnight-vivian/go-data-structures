package use_case_api

import (
	agraph "github.com/midnight-vivian/go-data-structures/algorithms/graph"

	dgraph "github.com/midnight-vivian/go-data-structures/data-structures/graph"
)

//CC focuses on the connected components in a graph

type CC struct {
	mark  []bool
	id    []int
	count int //count of cc in the graph
}

func NewCC(udg *dgraph.UdGraph) *CC {
	cc := &CC{
		mark: make([]bool, udg.GetV()),
		id:   make([]int, udg.GetV()),
	}

	for s := 0; s < udg.GetV(); s++ {
		if !cc.mark[s] {
			agraph.UndirectedDfsRecursion(udg, s, -1, cc.processed, cc.process)
			cc.count++
		}
	}
	return cc
}

func (cc *CC) Connected(v, w int) bool {
	return cc.id[v] == cc.id[w]
}

func (cc *CC) Count() int {
	return cc.count
}

func (cc *CC) Id(v int) int {
	return cc.id[v]
}

func (cc *CC) processed(vertex int) bool {
	return cc.mark[vertex]
}

func (cc *CC) process(vertex, pvertex int) {
	cc.mark[vertex] = true
	cc.id[vertex] = cc.count
}
