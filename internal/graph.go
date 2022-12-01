package internal

import mapset "github.com/deckarep/golang-set/v2"

type VertexSet struct {
	vertices mapset.Set[vertex]
}

func (g *VertexSet) Add(v vertex) {
	g.vertices.Add(v)
}

func (g *VertexSet) IsVertex(v vertex) bool {
	return g.vertices.Contains(v)
}

func (g *VertexSet) Order() int {
	return g.vertices.Cardinality()
}

func (g *VertexSet) Remove(v vertex) {
	if g.IsVertex(v) {
		g.vertices.Remove(v)
	}
}
