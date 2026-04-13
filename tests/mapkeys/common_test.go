// Copyright 2026
// license that can be found in the LICENSE file.

package tests

import (
	"encoding/json"
	"fmt"
	"os"
	"testing"

	"github.com/google/go-cmp/cmp"
	utils "github.com/name212/go-cmp-utils"
	"github.com/stretchr/testify/require"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

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

var (
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
			"version": 1,
		},
	}
)

type (
	mapAny   = map[string]any
	hasDiffs = map[string]bool
)

type testNamed interface {
	Name() string
}

func runTest(t *testing.T, tst testNamed, run func(t *testing.T), notPrintSkip ...bool) {
	curTest := tst.Name()

	testNameForRun := os.Getenv("RUN_ONE_TEST")

	shouldSkip := testNameForRun != "" && testNameForRun != curTest

	if (len(notPrintSkip) > 0 && notPrintSkip[0]) && shouldSkip {
		return
	}

	t.Run(curTest, func(t *testing.T) {
		if shouldSkip {
			t.Skipf("Test %s skipped because run one test %s", curTest, testNameForRun)
			return
		}

		run(t)
	})
}

type mapsTuple struct {
	name    string
	first   mapAny
	second  mapAny
	hasDiff bool
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

func applyHasDiff(hasDiff hasDiffs, tuples []mapsTuple) []mapsTuple {
	byName := make(map[string]*mapsTuple)

	names := make([]string, 0, len(tuples))

	for _, tt := range tuples {
		ttt := tt
		byName[tt.name] = &ttt
		names = append(names, ttt.name)
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
	for _, name := range names {
		if _, ok := hasDiffKeys[name]; !ok {
			panic(fmt.Sprintf("hasDiff name '%s' not found in tuples", name))
		}

		tt, ok := byName[name]
		if !ok {
			panic(fmt.Sprintf("byNames name '%s' not found in tuples", name))
		}

		res = append(res, *tt)
	}

	return res
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

func setField(m mapAny, v any, path ...string) mapAny {
	err := unstructured.SetNestedField(m, v, path...)
	if err != nil {
		panic(fmt.Sprintf("Cannot set %v to %v", path, v))
	}

	return m
}

func setFieldMap(m mapAny, v mapAny, path ...string) mapAny {
	err := unstructured.SetNestedMap(m, v, path...)
	if err != nil {
		panic(fmt.Sprintf("Cannot set %v to map %v", path, v))
	}

	return m
}

func extractSlice[T any](m mapAny, path ...string) []T {
	s, exists, err := unstructured.NestedSlice(m, path...)
	if err != nil {
		panic(fmt.Sprintf("Cannot get slice for %v: %v", path, err))
	}

	if !exists {
		panic(fmt.Sprintf("slice not exists for %v", path))
	}

	res := make([]T, 0, len(s))

	for i, v := range s {
		typed, ok := v.(T)
		if !ok {
			panic(
				fmt.Sprintf(
					"cannot convert for %v val for index %d. Has type %T",
					path,
					i,
					v,
				),
			)
		}
		res = append(res, typed)
	}

	return res
}

func setFieldSlice[T any](m mapAny, s []T, path ...string) mapAny {
	converted := make([]any, 0, len(s))
	for _, v := range s {
		converted = append(converted, v)
	}

	err := unstructured.SetNestedSlice(m, converted, path...)
	if err != nil {
		panic(fmt.Sprintf("Cannot set %v as slice %v", path, s))
	}

	return m
}

func deleteField(m mapAny, path ...string) mapAny {
	unstructured.RemoveNestedField(m, path...)
	return m
}
