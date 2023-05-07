package dag

type testVertex[T comparable] struct {
	WID string `json:"i"`
	Val T      `json:"v"`
}

func (tv testVertex[T]) ID() string {
	return tv.WID
}

func (tv testVertex[T]) Vertex() (id string, value T) {
	return tv.WID, tv.Val
}

type testStorableDAG[T comparable] struct {
	StorableVertices []testVertex[T] `json:"vs"`
	StorableEdges    []storableEdge  `json:"es"`
}

func (g testStorableDAG[T]) Vertices() []Vertexer[T] {
	l := make([]Vertexer[T], 0, len(g.StorableVertices))
	for _, v := range g.StorableVertices {
		l = append(l, v)
	}
	return l
}

func (g testStorableDAG[T]) Edges() []Edger {
	l := make([]Edger, 0, len(g.StorableEdges))
	for _, v := range g.StorableEdges {
		l = append(l, v)
	}
	return l
}
