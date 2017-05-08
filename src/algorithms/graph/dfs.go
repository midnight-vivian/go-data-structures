package graph

import (
	"data-structures/graph"
	"data-structures/stack"
)

func UndirectedDfsRecursion(udg *graph.UndirectedGraph, vertex, pvertex int, processed func(int) bool, process func(int, int)) {
	process(vertex, pvertex)
	for _, v := range udg.Adj(vertex) {
		if !processed(v) {
			UndirectedDfsRecursion(udg, v, vertex, processed, process)
		}
	}
}

func UndirectedDfsIteration(udg *graph.UndirectedGraph, vertex, pvertex int, processed func(int) bool, process func(int, int)) {
	process(vertex, pvertex)
	s := stack.NewStackLinkedList()
	s.Push(vertex)

	for !s.Empty() {
		v := s.Pop().(int)
		for _, sv := range udg.Adj(v) {
			if !processed(sv) {
				process(sv, v)
				s.Push(sv)
			}
		}
	}

}