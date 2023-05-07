package dag

import (
	"testing"

	"github.com/go-test/deep"
)

type testVisitor[T comparable] struct {
	Values []T
}

func (pv *testVisitor[T]) Visit(v Vertexer[T]) {
	_, value := v.Vertex()
	pv.Values = append(pv.Values, value)
}

// schematic diagram:
//
//	v5
//	^
//	|
//	v4
//	^
//	|
//	v2 --> v3
//	^
//	|
//	v1
func getTestWalkDAG() *DAG[string] {
	dag := NewDAG[string]()

	v1, v2, v3, v4, v5 := "1", "2", "3", "4", "5"
	_ = dag.AddVertexByID(v1, "v1")
	_ = dag.AddVertexByID(v2, "v2")
	_ = dag.AddVertexByID(v3, "v3")
	_ = dag.AddVertexByID(v4, "v4")
	_ = dag.AddVertexByID(v5, "v5")
	_ = dag.AddEdge(v1, v2)
	_ = dag.AddEdge(v2, v3)
	_ = dag.AddEdge(v2, v4)
	_ = dag.AddEdge(v4, v5)

	return dag
}

// schematic diagram:
//
//	v4 --> v5
//	       ^
//	       |
//	v1 --> v3
//	       ^
//	       |
//	      v2
func getTestWalkDAG2() *DAG[string] {
	dag := NewDAG[string]()

	v1, v2, v3, v4, v5 := "1", "2", "3", "4", "5"
	_ = dag.AddVertexByID(v1, "v1")
	_ = dag.AddVertexByID(v2, "v2")
	_ = dag.AddVertexByID(v3, "v3")
	_ = dag.AddVertexByID(v4, "v4")
	_ = dag.AddVertexByID(v5, "v5")
	_ = dag.AddEdge(v1, v3)
	_ = dag.AddEdge(v2, v3)
	_ = dag.AddEdge(v3, v5)
	_ = dag.AddEdge(v4, v5)

	return dag
}

// schematic diagram:
//
//	v4 --> v5
//
//
//	v1 --> v3
//	       ^
//	       |
//	      v2
func getTestWalkDAG3() *DAG[string] {
	dag := NewDAG[string]()

	v1, v2, v3, v4, v5 := "1", "2", "3", "4", "5"
	_ = dag.AddVertexByID(v1, "v1")
	_ = dag.AddVertexByID(v2, "v2")
	_ = dag.AddVertexByID(v3, "v3")
	_ = dag.AddVertexByID(v4, "v4")
	_ = dag.AddVertexByID(v5, "v5")
	_ = dag.AddEdge(v1, v3)
	_ = dag.AddEdge(v2, v3)
	_ = dag.AddEdge(v4, v5)

	return dag
}

// schematic diagram:
//
//	v4     v5
//	^      ^
//	|      |
//	v2 --> v3
//	^
//	|
//	v1
func getTestWalkDAG4() *DAG[string] {
	dag := NewDAG[string]()

	v1, v2, v3, v4, v5 := "1", "2", "3", "4", "5"
	_ = dag.AddVertexByID(v1, "v1")
	_ = dag.AddVertexByID(v2, "v2")
	_ = dag.AddVertexByID(v3, "v3")
	_ = dag.AddVertexByID(v4, "v4")
	_ = dag.AddVertexByID(v5, "v5")
	_ = dag.AddEdge(v1, v2)
	_ = dag.AddEdge(v2, v3)
	_ = dag.AddEdge(v3, v5)
	_ = dag.AddEdge(v2, v4)

	return dag
}

func TestDFSWalk(t *testing.T) {
	cases := []struct {
		Name     string
		Dag      *DAG[string]
		Expected []string
	}{
		{
			Name:     "1",
			Dag:      getTestWalkDAG(),
			Expected: []string{"v1", "v2", "v3", "v4", "v5"},
		},
		{
			Name:     "2",
			Dag:      getTestWalkDAG2(),
			Expected: []string{"v1", "v3", "v5", "v2", "v4"},
		},
		{
			Name:     "3",
			Dag:      getTestWalkDAG3(),
			Expected: []string{"v1", "v3", "v2", "v4", "v5"},
		},
		{
			Name:     "4",
			Dag:      getTestWalkDAG4(),
			Expected: []string{"v1", "v2", "v3", "v5", "v4"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			pv := &testVisitor[string]{}
			c.Dag.DFSWalk(pv)

			expected := c.Expected
			actual := pv.Values
			if deep.Equal(expected, actual) != nil {
				t.Errorf("DFSWalk() = %v, want %v", actual, expected)
			}

		})
	}
}

func TestBFSWalk(t *testing.T) {
	cases := []struct {
		Name     string
		Dag      *DAG[string]
		Expected []string
	}{
		{
			Name:     "1",
			Dag:      getTestWalkDAG(),
			Expected: []string{"v1", "v2", "v3", "v4", "v5"},
		},
		{
			Name:     "2",
			Dag:      getTestWalkDAG2(),
			Expected: []string{"v1", "v2", "v4", "v3", "v5"},
		},
		{
			Name:     "3",
			Dag:      getTestWalkDAG3(),
			Expected: []string{"v1", "v2", "v4", "v3", "v5"},
		},
		{
			Name:     "4",
			Dag:      getTestWalkDAG4(),
			Expected: []string{"v1", "v2", "v3", "v4", "v5"},
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			pv := &testVisitor[string]{}
			c.Dag.BFSWalk(pv)

			expected := c.Expected
			actual := pv.Values
			if deep.Equal(expected, actual) != nil {
				t.Errorf("BFSWalk() = %v, want %v", actual, expected)
			}
		})
	}
}
