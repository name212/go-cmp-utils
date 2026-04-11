// Copyright 2026
// license that can be found in the LICENSE file.

package tests

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	utils "github.com/name212/go-cmp-utils"
	"github.com/stretchr/testify/require"
)

func TestMapStringComparator(t *testing.T) {
	type test struct {
		name  string
		paths [][]string
		maps  []mapsTuple
	}

	createPaths := func(tt test) []utils.MapPath {
		pathsComparators := make([]utils.MapPath, 0)
		for _, p := range tt.paths {
			pathsComparators = append(pathsComparators, utils.NewMapPathStringComparator(p...))
		}

		return pathsComparators
	}

	t.Run("flat", func(t *testing.T) {
		tests := []test{
			{
				name:  "empty paths",
				paths: nil,
				maps: createTestTuplesFlat(map[string]bool{
					"both nil":                  false,
					"both empty":                false,
					"same flat":                 false,
					"change flat first":         true,
					"change flat int":           true,
					"change flat slice":         true,
					"add notExists to flat":     true,
					"change flat int and first": true,
					"remove from flat int":      true,
				}),
			},
			{
				name:  "skip first",
				paths: [][]string{{"first"}},
				maps: createTestTuplesFlat(map[string]bool{
					"both nil":                  false,
					"both empty":                false,
					"same flat":                 false,
					"change flat first":         false,
					"change flat int":           true,
					"change flat slice":         true,
					"change flat int and first": true,
					"add notExists to flat":     true,
					"remove from flat int":      true,
				}),
			},
			{
				name:  "skip int",
				paths: [][]string{{"int"}},
				maps: createTestTuplesFlat(map[string]bool{
					"both nil":                  false,
					"both empty":                false,
					"same flat":                 false,
					"change flat first":         true,
					"change flat int":           false,
					"change flat slice":         true,
					"change flat int and first": true,
					"add notExists to flat":     true,
					"remove from flat int":      false,
				}),
			},
			{
				name:  "skip notExists",
				paths: [][]string{{"notExists"}},
				maps: createTestTuplesFlat(map[string]bool{
					"both nil":                  false,
					"both empty":                false,
					"same flat":                 false,
					"change flat first":         true,
					"change flat int":           true,
					"change flat slice":         true,
					"change flat int and first": true,
					"add notExists to flat":     false,
					"remove from flat int":      true,
				}),
			},

			{
				name:  "skip first and int",
				paths: [][]string{{"first"}, {"int"}},
				maps: createTestTuplesFlat(map[string]bool{
					"both nil":                  false,
					"both empty":                false,
					"same flat":                 false,
					"change flat first":         false,
					"change flat int":           false,
					"change flat slice":         true,
					"change flat int and first": false,
					"add notExists to flat":     true,
					"remove from flat int":      false,
				}),
			},

			{
				name: "skip all",
				paths: [][]string{
					{"first"},
					{"int"},
					{"slice"},
					{"notExists"},
				},
				maps: createTestTuplesFlat(map[string]bool{
					"both nil":                  false,
					"both empty":                false,
					"same flat":                 false,
					"change flat first":         false,
					"change flat int":           false,
					"change flat slice":         false,
					"change flat int and first": false,
					"add notExists to flat":     false,
					"remove from flat int":      false,
				}),
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				runOverTuples(t, createPaths(tt), tt.maps)
			})
		}
	})
}

func runOverTuples(t *testing.T, paths []utils.MapPath, tuples []mapsTuple) {
	for _, tuple := range tuples {
		t.Run(tuple.name, func(t *testing.T) {
			opts := utils.MapKeysFilter(paths...)
			diff := cmp.Diff(tuple.first, tuple.second, opts)
			assert := require.Empty
			if tuple.hasDiff {
				assert = require.NotEmpty
			}
			assert(t, diff)
		})
	}
}

type mapsTuple struct {
	name    string
	first   mapAny
	second  mapAny
	hasDiff bool
}

func createTestTuplesFlat(hasDiff map[string]bool) []mapsTuple {
	tuples := []mapsTuple{
		{
			name:   "both nil",
			first:  nil,
			second: nil,
		},

		{
			name:   "both empty",
			first:  make(mapAny),
			second: make(mapAny),
		},

		{
			name:   "same flat",
			first:  prepareOverJSON(flatMap),
			second: prepareOverJSON(flatMap),
		},

		{
			name:  "change flat first",
			first: prepareOverJSON(flatMap),
			second: func() mapAny {
				m := prepareOverJSON(flatMap)
				m["first"] = "another"
				return m
			}(),
		},

		{
			name:  "change flat int",
			first: prepareOverJSON(flatMap),
			second: func() mapAny {
				m := prepareOverJSON(flatMap)
				m["int"] = float64(24)
				return m
			}(),
		},

		{
			name:  "change flat int and first",
			first: prepareOverJSON(flatMap),
			second: func() mapAny {
				m := prepareOverJSON(flatMap)
				m["int"] = float64(24)
				m["first"] = "another"
				return m
			}(),
		},

		{
			name:  "change flat slice",
			first: prepareOverJSON(flatMap),
			second: func() mapAny {
				m := prepareOverJSON(flatMap)
				sl := m["slice"].([]any)
				sl[0] = "another"
				return m
			}(),
		},

		{
			name:  "add notExists to flat",
			first: prepareOverJSON(flatMap),
			second: func() mapAny {
				m := prepareOverJSON(flatMap)
				m["notExists"] = "string"
				return m
			}(),
		},

		{
			name:  "remove from flat int",
			first: prepareOverJSON(flatMap),
			second: func() mapAny {
				m := prepareOverJSON(flatMap)
				delete(m, "int")
				return m
			}(),
		},
	}

	return applyHasDiff(hasDiff, tuples)
}

func applyHasDiff(hasDiff map[string]bool, tuples []mapsTuple) []mapsTuple {
	byName := make(map[string]*mapsTuple)

	for _, tt := range tuples {
		ttt := tt
		byName[tt.name] = &ttt
	}

	hasDiffKeys := make(map[string]struct{})

	for k, v := range hasDiff {
		if _, ok := byName[k]; !ok {
			panic(fmt.Sprintf("hasDiff name '%s' not found in tuples", k))
		}

		byName[k].hasDiff = v
		hasDiffKeys[k] = struct{}{}
	}

	res := make([]mapsTuple, 0, len(byName))
	for k, tt := range byName {
		if _, ok := hasDiffKeys[k]; !ok {
			panic(fmt.Sprintf("hasDiff name '%s' not found in tuples", k))
		}
		res = append(res, *tt)
	}

	return res
}

var (
	mapNil   map[string]any = nil
	mapEmpty                = make(map[string]any)

	flatMap = map[string]any{
		"first": "firstvall",
		"int":   42,
		"slice": []string{"first"},
	}

	complexMap = map[string]any{
		"apiVersion": "v1",
		"kind":       "Kind",
		"metadata": map[string]any{
			"name":      "test",
			"namespace": "test-ns",
			"labels": map[string]string{
				"first":                   "first-val",
				"test.example.com/first":  "val1",
				"test.example.com/second": "val2",
			},
			"annotations": map[string]string{
				"first":                  "first-ann",
				"test.example.com/third": "ann1",
				"test.example.com/four":  "ann2",
			},
			"creationTimestamp": "2021-03-18T13:46:17Z",
			"generation":        6,
			"resourceVersion":   "1552261164",
			"uid":               "7f66a236-6931-478d-b635-69ab8862fa75",
		},
		"spec": map[string]any{
			"int":    42,
			"string": "str",
			"slice":  []int{2},
			"innerFirst": map[string]any{
				"innerInner": map[string]any{
					"innerInnerInner": map[string]any{
						"key":    "val",
						"intKey": 42,
					},
					"innerInnerKey":   "valKey",
					"innerInnerSlice": []string{"one"},
				},
				"sub": "subVal",
			},
			"innerSecond": map[string]any{
				"subInner": map[string]any{
					"innerInnerKey":   "valKey",
					"innerInnerSlice": []string{"one"},
				},
				"subSub": "subSubVal",
			},
		},

		"status": map[string]any{
			"phase": "Bound",
			"conditions": []map[string]any{
				{
					"lastTransitionTime": "2023-05-10T07:47:32Z",
					"status":             "True",
					"type":               "Ready",
				},
				{
					"lastTransitionTime": "2025-06-10T17:22:01Z",
					"status":             "False",
					"type":               "Error",
				},
			},
		},
	}
)

type mapAny = map[string]any

func TestClone(t *testing.T) {
	m := prepareOverJSON(map[string]any{
		"a":     "a",
		"b":     1,
		"slice": []string{"one"},
		"sub": map[string]any{
			"key": "val",
		},
	})

	type test struct {
		name   string
		action func(mapAny) mapAny
	}

	assertClone := func(t *testing.T, tst test) {
		cloned := prepareOverJSON(m)
		require.Equal(t, m, cloned, "should equal")

		afterAction := tst.action(cloned)
		require.NotEqual(t, m, afterAction, "should not equal after %s", tst.name)
	}

	tests := []test{
		{
			name: "change a",
			action: func(ma mapAny) mapAny {
				ma["a"] = "another"
				return ma
			},
		},

		{
			name: "change b",
			action: func(ma mapAny) mapAny {
				ma["b"] = 2
				return ma
			},
		},

		{
			name: "change slice",
			action: func(ma mapAny) mapAny {
				sl := ma["slice"].([]any)
				sl[0] = "another"
				return ma
			},
		},

		{
			name: "remove b",
			action: func(ma mapAny) mapAny {
				delete(ma, "b")
				return ma
			},
		},

		{
			name: "add another",
			action: func(ma mapAny) mapAny {
				ma["another"] = "val"
				return ma
			},
		},

		{
			name: "change sub",
			action: func(ma mapAny) mapAny {
				sub := ma["sub"].(mapAny)
				sub["key"] = "another"
				return ma
			},
		},
	}

	for _, tst := range tests {
		t.Run(tst.name, func(t *testing.T) {
			assertClone(t, tst)
		})
	}
}

func prepareOverJSON(m mapAny) mapAny {
	bt, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}

	var prepared mapAny
	err = json.Unmarshal(bt, &prepared)
	if err != nil {
		panic(err)
	}

	return prepared
}

func clone(m mapAny) mapAny {
	return prepareOverJSON(m)
}
