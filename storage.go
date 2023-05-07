package dag

// Vertexer is the interface that wraps the basic Vertex method.
// Vertex returns an id that identifies this vertex and the value of this vertex.
//
// The reason for defining this new structure is that the vertex id may be
// automatically generated when the caller adds a vertex. At this time, the
// vertex structure added by the user does not contain id information.
type Vertexer[T comparable] interface {
	Vertex() (id string, value T)
}

// Edger is the interface that wraps the basic Edge method.
// Edge returns the ids of two vertices that connect an edge.
type Edger interface {
	Edge() (srcID, dstID string)
}

// StorableDAG is the interface that defines a DAG that can be stored.
// It provides methods to get all vertices and all edges of a DAG.
type StorableDAG[T comparable] interface {
	Vertices() []Vertexer[T]
	Edges() []Edger
}

// storableVertex implements the Vertexer interface.
// It is implemented as a storable structure.
// And it uses short json tag to reduce the number of bytes after serialization.
type storableVertex[T comparable] struct {
	WrappedID string `json:"i"`
	Value     T      `json:"v"`
}

func (v storableVertex[T]) Vertex() (id string, value T) {
	return v.WrappedID, v.Value
}

func (v storableVertex[T]) ID() string {
	return v.WrappedID
}

// storableEdge implements the Edger interface.
// It is implemented as a storable structure.
// And it uses short json tag to reduce the number of bytes after serialization.
type storableEdge struct {
	SrcID string `json:"s"`
	DstID string `json:"d"`
}

func (e storableEdge) Edge() (srcID, dstID string) {
	return e.SrcID, e.DstID
}

// storableDAG implements the StorableDAG interface.
// It acts as a serializable operable structure.
// And it uses short json tag to reduce the number of bytes after serialization.
type storableDAG[T comparable] struct {
	StorableVertices []Vertexer[T] `json:"vs"`
	StorableEdges    []Edger       `json:"es"`
}

func (g storableDAG[T]) Vertices() []Vertexer[T] {
	return g.StorableVertices
}

func (g storableDAG[T]) Edges() []Edger {
	return g.StorableEdges
}
