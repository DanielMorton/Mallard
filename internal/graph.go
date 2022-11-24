package internal

import mapset "github.com/deckarep/golang-set/v2"

type Graph struct {
	vertices mapset.Set[vertex]
}

func (g *Graph) Add(v vertex) {
	g.vertices.Add(v)
}

func (g *Graph) IsVertex(v vertex) bool {
	return g.vertices.Contains(v)
}

func (g *Graph) Order() int {
	return g.vertices.Cardinality()
}

func (g *Graph) Remove(v vertex) {
	if g.IsVertex(v) {
		g.vertices.Remove(v)
	}
}
