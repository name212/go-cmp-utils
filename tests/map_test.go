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
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
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

	t.Run("complex", func(t *testing.T) {
		tests := []test{
			{
				name:  "empty paths",
				paths: nil,
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              true,
					"change complex name":                                              true,
					"change complex labels":                                            true,
					"change complex annotations":                                       true,
					"change complex annotations and labels":                            true,
					"add complex notExists to labels":                                  true,
					"add complex notExists to annotations":                             true,
					"delete complex first from labels":                                 true,
					"delete complex test.example.com/third from annotations":           true,
					"delete complex first from annotations and labels":                 true,
					"delete complex all labels":                                        true,
					"delete complex all annotations":                                   true,
					"delete complex all annotations and labels":                        true,
					"add complex all labels":                                           true,
					"add complex all annotations":                                      true,
					"add complex all annotations and labels":                           true,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip kind",
				paths: [][]string{
					{"kind"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              false,
					"change complex name":                                              true,
					"change complex labels":                                            true,
					"change complex annotations":                                       true,
					"change complex annotations and labels":                            true,
					"add complex notExists to labels":                                  true,
					"add complex notExists to annotations":                             true,
					"delete complex first from labels":                                 true,
					"delete complex test.example.com/third from annotations":           true,
					"delete complex first from annotations and labels":                 true,
					"delete complex all labels":                                        true,
					"delete complex all annotations":                                   true,
					"delete complex all annotations and labels":                        true,
					"add complex all labels":                                           true,
					"add complex all annotations":                                      true,
					"add complex all annotations and labels":                           true,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip name",
				paths: [][]string{
					{"metadata", "name"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              true,
					"change complex name":                                              false,
					"change complex labels":                                            true,
					"change complex annotations":                                       true,
					"change complex annotations and labels":                            true,
					"add complex notExists to labels":                                  true,
					"add complex notExists to annotations":                             true,
					"delete complex first from labels":                                 true,
					"delete complex test.example.com/third from annotations":           true,
					"delete complex first from annotations and labels":                 true,
					"delete complex all labels":                                        true,
					"delete complex all annotations":                                   true,
					"delete complex all annotations and labels":                        true,
					"add complex all labels":                                           true,
					"add complex all annotations":                                      true,
					"add complex all annotations and labels":                           true,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip name and kind",
				paths: [][]string{
					{"metadata", "name"},
					{"kind"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              false,
					"change complex name":                                              false,
					"change complex labels":                                            true,
					"change complex annotations":                                       true,
					"change complex annotations and labels":                            true,
					"add complex notExists to labels":                                  true,
					"add complex notExists to annotations":                             true,
					"delete complex first from labels":                                 true,
					"delete complex test.example.com/third from annotations":           true,
					"delete complex first from annotations and labels":                 true,
					"delete complex all labels":                                        true,
					"delete complex all annotations":                                   true,
					"delete complex all annotations and labels":                        true,
					"add complex all labels":                                           true,
					"add complex all annotations":                                      true,
					"add complex all annotations and labels":                           true,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip labels",
				paths: [][]string{
					{"metadata", "labels"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              true,
					"change complex name":                                              true,
					"change complex labels":                                            false,
					"change complex annotations":                                       true,
					"change complex annotations and labels":                            true,
					"add complex notExists to labels":                                  false,
					"add complex notExists to annotations":                             true,
					"delete complex first from labels":                                 false,
					"delete complex test.example.com/third from annotations":           true,
					"delete complex first from annotations and labels":                 true,
					"delete complex all labels":                                        false,
					"delete complex all annotations":                                   true,
					"delete complex all annotations and labels":                        true,
					"add complex all labels":                                           false,
					"add complex all annotations":                                      true,
					"add complex all annotations and labels":                           true,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip annotations",
				paths: [][]string{
					{"metadata", "annotations"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              true,
					"change complex name":                                              true,
					"change complex labels":                                            true,
					"change complex annotations":                                       false,
					"change complex annotations and labels":                            true,
					"add complex notExists to labels":                                  true,
					"add complex notExists to annotations":                             false,
					"delete complex first from labels":                                 true,
					"delete complex test.example.com/third from annotations":           false,
					"delete complex first from annotations and labels":                 true,
					"delete complex all labels":                                        true,
					"delete complex all annotations":                                   false,
					"delete complex all annotations and labels":                        true,
					"add complex all labels":                                           true,
					"add complex all annotations":                                      false,
					"add complex all annotations and labels":                           true,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip one annotations",
				paths: [][]string{
					{"metadata", "annotations", "test.example.com/four"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              true,
					"change complex name":                                              true,
					"change complex labels":                                            true,
					"change complex annotations":                                       false,
					"change complex annotations and labels":                            true,
					"add complex notExists to labels":                                  true,
					"add complex notExists to annotations":                             true,
					"delete complex first from labels":                                 true,
					"delete complex test.example.com/third from annotations":           true,
					"delete complex first from annotations and labels":                 true,
					"delete complex all labels":                                        true,
					"delete complex all annotations":                                   true,
					"delete complex all annotations and labels":                        true,
					"add complex all labels":                                           true,
					"add complex all annotations":                                      true,
					"add complex all annotations and labels":                           true,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip multiple labels",
				paths: [][]string{
					{"metadata", "labels", "test.example.com/first"},
					{"metadata", "labels", "notExists"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              true,
					"change complex name":                                              true,
					"change complex labels":                                            false,
					"change complex annotations":                                       true,
					"change complex annotations and labels":                            true,
					"add complex notExists to labels":                                  false,
					"add complex notExists to annotations":                             true,
					"delete complex first from labels":                                 true,
					"delete complex test.example.com/third from annotations":           true,
					"delete complex first from annotations and labels":                 true,
					"delete complex all labels":                                        true,
					"delete complex all annotations":                                   true,
					"delete complex all annotations and labels":                        true,
					"add complex all labels":                                           true,
					"add complex all annotations":                                      true,
					"add complex all annotations and labels":                           true,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip annotations and labels",
				paths: [][]string{
					{"metadata", "labels"},
					{"metadata", "annotations"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              true,
					"change complex name":                                              true,
					"change complex labels":                                            false,
					"change complex annotations":                                       false,
					"change complex annotations and labels":                            false,
					"add complex notExists to labels":                                  false,
					"add complex notExists to annotations":                             false,
					"delete complex first from labels":                                 false,
					"delete complex test.example.com/third from annotations":           false,
					"delete complex first from annotations and labels":                 false,
					"delete complex all labels":                                        false,
					"delete complex all annotations":                                   false,
					"delete complex all annotations and labels":                        false,
					"add complex all labels":                                           false,
					"add complex all annotations":                                      false,
					"add complex all annotations and labels":                           false,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip annotations and labels change path order",
				paths: [][]string{
					{"metadata", "annotations"},
					{"metadata", "labels"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              true,
					"change complex name":                                              true,
					"change complex labels":                                            false,
					"change complex annotations":                                       false,
					"change complex annotations and labels":                            false,
					"add complex notExists to labels":                                  false,
					"add complex notExists to annotations":                             false,
					"delete complex first from labels":                                 false,
					"delete complex test.example.com/third from annotations":           false,
					"delete complex first from annotations and labels":                 false,
					"delete complex all labels":                                        false,
					"delete complex all annotations":                                   false,
					"delete complex all annotations and labels":                        false,
					"add complex all labels":                                           false,
					"add complex all annotations":                                      false,
					"add complex all annotations and labels":                           false,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip annotations and labels and name",
				paths: [][]string{
					{"metadata", "annotations"},
					{"metadata", "name"},
					{"metadata", "labels"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              true,
					"change complex name":                                              false,
					"change complex labels":                                            false,
					"change complex annotations":                                       false,
					"change complex annotations and labels":                            false,
					"add complex notExists to labels":                                  false,
					"add complex notExists to annotations":                             false,
					"delete complex first from labels":                                 false,
					"delete complex test.example.com/third from annotations":           false,
					"delete complex first from annotations and labels":                 false,
					"delete complex all labels":                                        false,
					"delete complex all annotations":                                   false,
					"delete complex all annotations and labels":                        false,
					"add complex all labels":                                           false,
					"add complex all annotations":                                      false,
					"add complex all annotations and labels":                           false,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
				}),
			},

			{
				name: "skip metadata",
				paths: [][]string{
					{"metadata"},
				},
				maps: createTestTuplesComplex(map[string]bool{
					"both nil":                                                         false,
					"both empty":                                                       false,
					"same complex":                                                     false,
					"change complex kind":                                              true,
					"change complex name":                                              false,
					"change complex labels":                                            false,
					"change complex annotations":                                       false,
					"change complex annotations and labels":                            false,
					"add complex notExists to labels":                                  false,
					"add complex notExists to annotations":                             false,
					"delete complex first from labels":                                 false,
					"delete complex test.example.com/third from annotations":           false,
					"delete complex first from annotations and labels":                 false,
					"delete complex all labels":                                        false,
					"delete complex all annotations":                                   false,
					"delete complex all annotations and labels":                        false,
					"add complex all labels":                                           false,
					"add complex all annotations":                                      false,
					"add complex all annotations and labels":                           false,
					"complex change spec string":                                       true,
					"complex change full deep spec intKey":                             true,
					"complex change spec innerInner.innerInnerKey":                     true,
					"complex change spec innerFirst.sub":                               true,
					"complex change multiple spec keys":                                true,
					"complex delete spec string":                                       true,
					"complex delete spec multiple keys":                                true,
					"complex add spec notExists key":                                   true,
					"complex add notExists key to spec innerInnerInner and innerFirst": true,
					"complex add change remove in spec":                                true,
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

func createTestTuplesComplex(hasDiff map[string]bool) []mapsTuple {
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
			name:   "same complex",
			first:  prepareOverJSON(complexMap),
			second: prepareOverJSON(complexMap),
		},

		{
			name:  "change complex kind",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				m["kind"] = "Another"
				return m
			}(),
		},

		{
			name:  "change complex name",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return setField(m, "another-name", "metadata", "name")
			}(),
		},

		{
			name:  "change complex labels",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return setField(m, "changed", "metadata", "labels", "test.example.com/first")
			}(),
		},

		{
			name:  "change complex annotations",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return setField(m, "changed", "metadata", "annotations", "test.example.com/four")
			}(),
		},

		{
			name:  "change complex annotations and labels",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				m = setField(m, "changed", "metadata", "labels", "test.example.com/second")
				return setField(m, "changed-ann", "metadata", "annotations", "first")
			}(),
		},

		{
			name:  "add complex notExists to labels",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return setField(m, "added", "metadata", "labels", "notExists")
			}(),
		},

		{
			name:  "add complex notExists to annotations",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return setField(m, "added", "metadata", "annotations", "notExists")
			}(),
		},

		{
			name:  "delete complex first from labels",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return deleteField(m, "metadata", "labels", "first")
			}(),
		},

		{
			name:  "delete complex test.example.com/third from annotations",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return deleteField(m, "metadata", "annotations", "test.example.com/third")
			}(),
		},

		{
			name:  "delete complex first from annotations and labels",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				m = deleteField(m, "metadata", "labels", "first")
				return deleteField(m, "metadata", "annotations", "first")
			}(),
		},

		{
			name:  "delete complex all labels",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return deleteField(m, "metadata", "labels")
			}(),
		},

		{
			name:  "delete complex all annotations",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return deleteField(m, "metadata", "annotations")
			}(),
		},

		{
			name:  "delete complex all annotations and labels",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				m = deleteField(m, "metadata", "annotations")
				return deleteField(m, "metadata", "labels")
			}(),
		},

		{
			name: "add complex all labels",
			first: func() mapAny {
				m := prepareOverJSON(complexMap)
				return deleteField(m, "metadata", "labels")
			}(),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				toSet := map[string]any{
					"first":                        "first-label",
					"test.example.com/first-label": "label",
				}
				return setFieldMap(m, toSet, "metadata", "labels")
			}(),
		},

		{
			name: "add complex all annotations",
			first: func() mapAny {
				m := prepareOverJSON(complexMap)
				return deleteField(m, "metadata", "annotations")
			}(),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				toSet := map[string]any{
					"first":                       "first-ann",
					"test.example.com/second-ann": "ann",
				}
				return setFieldMap(m, toSet, "metadata", "annotations")
			}(),
		},

		{
			name: "add complex all annotations and labels",
			first: func() mapAny {
				m := prepareOverJSON(complexMap)
				m = deleteField(m, "metadata", "labels")
				return deleteField(m, "metadata", "annotations")
			}(),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				toAnn := map[string]any{
					"first":                       "first-ann",
					"test.example.com/second-ann": "ann",
				}
				toLabels := map[string]any{
					"first":                        "first-label",
					"test.example.com/first-label": "label",
				}
				m = setFieldMap(m, toLabels, "metadata", "labels")
				return setFieldMap(m, toAnn, "metadata", "annotations")
			}(),
		},

		{
			name:  "complex change spec string",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return setField(m, "changed",
					"spec",
					"string",
				)
			}(),
		},

		{
			name:  "complex change full deep spec intKey",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return setField(m, float64(22),
					"spec",
					"innerFirst",
					"innerInner",
					"innerInnerInner",
					"intKey",
				)
			}(),
		},

		{
			name:  "complex change spec innerInner.innerInnerKey",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return setField(m, "changed",
					"spec",
					"innerFirst",
					"innerInner",
					"innerInnerKey",
				)
			}(),
		},

		{
			name:  "complex change spec innerFirst.sub",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return setField(m, "changed",
					"spec",
					"innerFirst",
					"sub",
				)
			}(),
		},

		{
			name:  "complex change multiple spec keys",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				m = setField(m, "changed",
					"spec",
					"string",
				)
				m = setField(m, "changed",
					"spec",
					"innerFirst",
					"sub",
				)
				return setField(m, float64(22),
					"spec",
					"innerFirst",
					"innerInner",
					"innerInnerInner",
					"intKey",
				)
			}(),
		},

		{
			name:  "complex delete spec string",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return deleteField(m, "spec", "string")
			}(),
		},

		{
			name:  "complex delete spec multiple keys",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				m = deleteField(m, "spec", "string")
				m = deleteField(m, "spec", "innerFirst", "sub")
				return deleteField(m,
					"spec",
					"innerFirst",
					"innerInner",
					"innerInnerInner",
					"intKey",
				)
			}(),
		},

		{
			name:  "complex add spec notExists key",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				return setField(m, "added",
					"spec",
					"notExists",
				)
			}(),
		},

		{
			name:  "complex add notExists key to spec innerInnerInner and innerFirst",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				m = setField(m, "added",
					"spec",
					"notExists",
				)
				m = setField(m, "added",
					"spec",
					"innerFirst",
					"notExists",
				)
				return setField(m, "added",
					"spec",
					"innerFirst",
					"innerInner",
					"innerInnerInner",
					"notExists",
				)
			}(),
		},

		{
			name:  "complex add change remove in spec",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				m = setField(m, "added",
					"spec",
					"notExists",
				)
				m = deleteField(m,
					"spec",
					"innerFirst",
					"innerInner",
					"innerInnerInner",
					"intKey",
				)
				return setField(m, "changed",
					"spec",
					"innerFirst",
					"sub",
				)
			}(),
		},
	}

	return applyHasDiff(hasDiff, tuples)
}

func applyHasDiff(hasDiff map[string]bool, tuples []mapsTuple) []mapsTuple {
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

func deleteField(m mapAny, path ...string) mapAny {
	unstructured.RemoveNestedField(m, path...)
	return m
}
