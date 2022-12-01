package internal

type vertex string
type edge struct {
	start vertex
	end   vertex
}

func (e *edge) Start() vertex {
	return e.start
}

func (e *edge) End() vertex {
	return e.end
}

func NewEdge(start, end vertex) *edge {
	e := &edge{start: start, end: end}
	return e
}
