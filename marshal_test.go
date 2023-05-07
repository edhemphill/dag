package dag

import (
	"encoding/json"
	"testing"

	"github.com/go-test/deep"
)

func TestMarshalUnmarshalJSON(t *testing.T) {
	cases := []struct {
		Name     string
		Dag      *DAG[string]
		Expected string
	}{
		{
			Name:     "1",
			Dag:      getTestWalkDAG(),
			Expected: `{"vs":[{"i":"1","v":"v1"},{"i":"2","v":"v2"},{"i":"3","v":"v3"},{"i":"4","v":"v4"},{"i":"5","v":"v5"}],"es":[{"s":"1","d":"2"},{"s":"2","d":"3"},{"s":"2","d":"4"},{"s":"4","d":"5"}]}`,
		},
		{
			Name:     "2",
			Dag:      getTestWalkDAG2(),
			Expected: `{"vs":[{"i":"1","v":"v1"},{"i":"3","v":"v3"},{"i":"5","v":"v5"},{"i":"2","v":"v2"},{"i":"4","v":"v4"}],"es":[{"s":"1","d":"3"},{"s":"3","d":"5"},{"s":"2","d":"3"},{"s":"4","d":"5"}]}`,
		},
		{
			Name:     "3",
			Dag:      getTestWalkDAG3(),
			Expected: `{"vs":[{"i":"1","v":"v1"},{"i":"3","v":"v3"},{"i":"2","v":"v2"},{"i":"4","v":"v4"},{"i":"5","v":"v5"}],"es":[{"s":"1","d":"3"},{"s":"2","d":"3"},{"s":"4","d":"5"}]}`,
		},
		{
			Name:     "4",
			Dag:      getTestWalkDAG4(),
			Expected: `{"vs":[{"i":"1","v":"v1"},{"i":"2","v":"v2"},{"i":"3","v":"v3"},{"i":"5","v":"v5"},{"i":"4","v":"v4"}],"es":[{"s":"1","d":"2"},{"s":"2","d":"3"},{"s":"2","d":"4"},{"s":"3","d":"5"}]}`,
		},
	}

	for _, c := range cases {
		c := c
		t.Run(c.Name, func(t *testing.T) {
			t.Parallel()

			data, err := json.Marshal(c.Dag)
			if err != nil {
				t.Error(err)
			}

			actual := string(data)
			if deep.Equal(c.Expected, actual) != nil {
				t.Errorf("Marshal() = %v, want %v", actual, c.Expected)
			}

			d1 := &DAG[string]{}
			errNotSupported := json.Unmarshal(data, d1)
			if errNotSupported == nil {
				t.Errorf("UnmarshalJSON() = nil, want %v", "This method is not supported")
			}

			var wd testStorableDAG[string]
			dag, err := UnmarshalJSON[string](data, &wd)
			if err != nil {
				t.Fatal(err)
			}
			if deep.Equal(c.Dag, dag) != nil {
				t.Errorf("UnmarshalJSON() = %v, want %v", dag.String(), c.Dag.String())
			}

		})

	}
}
