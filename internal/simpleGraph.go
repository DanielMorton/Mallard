package internal

import mapset "github.com/deckarep/golang-set/v2"

type SimpleGraph struct {
	VertexSet
	edges map[vertex]mapset.Set[vertex]
}

func (g *SimpleGraph) AddEdge(e edge) {
	if !g.IsEdge(e) {
		g.AddVertex(e.start)
		g.AddVertex(e.end)
		g.edges[e.start].Add(e.end)
	}
}

func (g *SimpleGraph) AddVertex(v vertex) {
	if !g.IsVertex(v) {
		g.Add(v)
		g.edges[v] = mapset.NewSet[vertex]()
	}
}

func (g *SimpleGraph) Degree(v vertex) int {
	if g.IsVertex(v) {
		return g.edges[v].Cardinality()
	}
	panic("Vertex " + v + " not in graph.")
}

func (g *SimpleGraph) DropEdge(e edge) {
	if g.IsEdge(e) {
		g.edges[e.start].Remove(e.end)
	}
}

func (g *SimpleGraph) DropVertex(v vertex) {
	if g.IsVertex(v) {
		g.Remove(v)
		delete(g.edges, v)
		for w := range g.edges {
			if g.edges[w].Contains(v) {
				g.edges[w].Remove(v)
			}
		}
	}
}

func (g *SimpleGraph) IsEdge(e edge) bool {
	if !g.IsVertex(e.start) || !g.IsVertex(e.end) {
		return false
	}
	return g.edges[e.start].Contains(e.end)
}

func build(edges []edge) *SimpleGraph {
	graph := &SimpleGraph{edges: make(map[vertex]mapset.Set[vertex])}
	for _, e := range edges {
		graph.AddVertex(e.start)
		graph.AddVertex(e.end)
		graph.edges[e.start].Add(e.end)
	}
	return graph
}
