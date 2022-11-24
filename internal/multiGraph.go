package internal

type MuliGraph struct {
	edges map[vertex][]vertex
}

func (g *MuliGraph) Order() int {
	return len(g.edges)
}
