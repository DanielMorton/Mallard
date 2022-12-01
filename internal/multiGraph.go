package internal

type MuliGraph struct {
	VertexSet
	edges map[vertex]map[vertex]int
}

func (g *MuliGraph) AddEdge(e edge) {
	if g.IsEdge(e) {
		g.edges[e.start][e.end] += 1
	} else {
		g.AddVertex(e.start)
		g.AddVertex(e.end)
		g.edges[e.start][e.end] = 1
	}
}

func (g *MuliGraph) AddVertex(v vertex) {
	if !g.IsVertex(v) {
		g.Add(v)
		g.edges[v] = make(map[vertex]int)
	}
}

func (g *MuliGraph) Degree(v vertex) int {
	if g.IsVertex(v) {
		return len(g.edges[v])
	}
	panic("Vertex " + v + " not in graph.")
}

func (g *MuliGraph) DropEdge(e edge) {
	if g.IsEdge(e) {
		if g.EdgeMult(e) > 1 {
			g.edges[e.start][e.end] -= 1
		} else {
			delete(g.edges[e.start], e.end)
		}
	}
}

func (g *MuliGraph) DropVertex(v vertex) {
	if g.IsVertex(v) {
		g.Remove(v)
		delete(g.edges, v)
		for w := range g.vertices.Iterator().C {
			if _, ok := g.edges[w][v]; ok {
				delete(g.edges[w], v)
			}
		}
	}
}

func (g *MuliGraph) EdgeMult(e edge) int {
	if g.IsEdge(e) {
		return g.edges[e.start][e.end]
	} else {
		return 0
	}
}

func (g *MuliGraph) IsEdge(e edge) bool {
	if !g.IsVertex(e.start) || !g.IsVertex(e.end) {
		return false
	}
	_, ok := g.edges[e.start][e.end]
	return ok
}
