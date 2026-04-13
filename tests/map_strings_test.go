// Copyright 2026
// license that can be found in the LICENSE file.

package tests

import (
	"testing"

	utils "github.com/name212/go-cmp-utils"
)

func TestMapStringComparatorFlat(t *testing.T) {
	tests := []testStringsComparator{
		{
			name:  "empty paths",
			paths: nil,
			maps:  createTestTuplesFlat(flatHasDiffsEmpty()),
		},
		{
			name:  "skip first",
			paths: [][]string{{"first"}},
			maps:  createTestTuplesFlat(flatHasDiffsFirst()),
		},
		{
			name:  "skip int",
			paths: [][]string{{"int"}},
			maps:  createTestTuplesFlat(flatHasDiffsInt()),
		},
		{
			name:  "skip notExists",
			paths: [][]string{{"notExists"}},
			maps:  createTestTuplesFlat(flatHasDiffsNotExists()),
		},

		{
			name:  "skip first and int",
			paths: [][]string{{"first"}, {"int"}},
			maps:  createTestTuplesFlat(flatHasDiffsFirstInt()),
		},

		{
			name: "skip all",
			paths: [][]string{
				{"first"},
				{"int"},
				{"slice"},
				{"notExists"},
			},
			maps: createTestTuplesFlat(flatHasDiffsAll()),
		},
	}

	runMapStringComparatorTests(t, tests)
}

func TestMapStringComparatorComplex(t *testing.T) {
	tests := []testStringsComparator{
		{
			name:  "empty paths",
			paths: nil,
			maps:  createTestTuplesComplex(complexHasDiffsEmpty()),
		},

		{
			name: "skip kind",
			paths: [][]string{
				{"kind"},
			},
			maps: createTestTuplesComplex(complexHasDiffsKind()),
		},

		{
			name: "skip name",
			paths: [][]string{
				{"metadata", "name"},
			},
			maps: createTestTuplesComplex(complexHasDiffsName()),
		},

		{
			name: "skip name and kind",
			paths: [][]string{
				{"metadata", "name"},
				{"kind"},
			},
			maps: createTestTuplesComplex(complexHasDiffsKindName()),
		},

		{
			name: "skip labels",
			paths: [][]string{
				{"metadata", "labels"},
			},
			maps: createTestTuplesComplex(complexHasDiffsLabels()),
		},

		{
			name: "skip annotations",
			paths: [][]string{
				{"metadata", "annotations"},
			},
			maps: createTestTuplesComplex(complexHasDiffsAnnotations()),
		},

		{
			name: "skip one annotations",
			paths: [][]string{
				{"metadata", "annotations", "test.example.com/four"},
			},
			maps: createTestTuplesComplex(complexHasDiffsOneAnnotation()),
		},

		{
			name: "skip multiple labels",
			paths: [][]string{
				{"metadata", "labels", "test.example.com/first"},
				{"metadata", "labels", "notExists"},
			},
			maps: createTestTuplesComplex(complexHasDiffsMultipleLabels()),
		},

		{
			name: "skip annotations and labels",
			paths: [][]string{
				{"metadata", "labels"},
				{"metadata", "annotations"},
			},
			maps: createTestTuplesComplex(complexHasDiffsAnnotationsAndLabels()),
		},

		{
			name: "skip annotations and labels change path order",
			paths: [][]string{
				{"metadata", "annotations"},
				{"metadata", "labels"},
			},
			maps: createTestTuplesComplex(complexHasDiffsAnnotationsAndLabels()),
		},

		{
			name: "skip annotations and labels and name",
			paths: [][]string{
				{"metadata", "annotations"},
				{"metadata", "name"},
				{"metadata", "labels"},
			},
			maps: createTestTuplesComplex(complexHasDiffsAnnotationsLabelsName()),
		},

		{
			name: "skip metadata",
			paths: [][]string{
				{"metadata"},
			},
			maps: createTestTuplesComplex(hasDiffs{
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
				"complex change conditions slice in status":                        true,
				"complex add status":                                               true,
				"complex change conditions slice and phase in status":              true,
			}),
		},

		{
			name: "skip spec string",
			paths: [][]string{
				{"spec", "string"},
			},
			maps: createTestTuplesComplex(complexHasDiffsSpecString()),
		},

		{
			name: "skip spec innerInnerKey",
			paths: [][]string{
				{
					"spec",
					"innerFirst",
					"innerInner",
					"innerInnerKey",
				},
			},
			maps: createTestTuplesComplex(hasDiffs{
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
				"complex change spec innerInner.innerInnerKey":                     false,
				"complex change spec innerFirst.sub":                               true,
				"complex change multiple spec keys":                                true,
				"complex delete spec string":                                       true,
				"complex delete spec multiple keys":                                true,
				"complex add spec notExists key":                                   true,
				"complex add notExists key to spec innerInnerInner and innerFirst": true,
				"complex add change remove in spec":                                true,
				"complex change conditions slice in status":                        true,
				"complex add status":                                               true,
				"complex change conditions slice and phase in status":              true,
			}),
		},
		{
			name: "skip spec innerFirst.sub",
			paths: [][]string{
				{
					"spec",
					"innerFirst",
					"sub",
				},
			},
			maps: createTestTuplesComplex(complexHasDiffsInnerFirstSub()),
		},

		{
			name: "skip spec multiple changes full",
			paths: [][]string{
				{
					"spec",
					"string",
				},
				{
					"spec",
					"notExists",
				},
				{
					"spec",
					"innerFirst",
					"sub",
				},
				{
					"spec",
					"innerFirst",
					"innerInner",
					"innerInnerInner",
					"intKey",
				},
			},
			maps: createTestTuplesComplex(hasDiffs{
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
				"complex change spec string":                                       false,
				"complex change full deep spec intKey":                             false,
				"complex change spec innerInner.innerInnerKey":                     true,
				"complex change spec innerFirst.sub":                               false,
				"complex change multiple spec keys":                                false,
				"complex delete spec string":                                       false,
				"complex delete spec multiple keys":                                false,
				"complex add spec notExists key":                                   false,
				"complex add notExists key to spec innerInnerInner and innerFirst": true,
				"complex add change remove in spec":                                false,
				"complex change conditions slice in status":                        true,
				"complex add status":                                               true,
				"complex change conditions slice and phase in status":              true,
			}),
		},
		{
			name: "skip spec notExists",
			paths: [][]string{
				{
					"spec",
					"string",
				},
				{
					"spec",
					"notExists",
				},
				{
					"spec",
					"innerFirst",
					"notExists",
				},

				{
					"spec",
					"innerFirst",
					"innerInner",
					"innerInnerInner",
					"notExists",
				},
			},
			maps: createTestTuplesComplex(hasDiffs{
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
				"complex change spec string":                                       false,
				"complex change full deep spec intKey":                             true,
				"complex change spec innerInner.innerInnerKey":                     true,
				"complex change spec innerFirst.sub":                               true,
				"complex change multiple spec keys":                                true,
				"complex delete spec string":                                       false,
				"complex delete spec multiple keys":                                true,
				"complex add spec notExists key":                                   false,
				"complex add notExists key to spec innerInnerInner and innerFirst": false,
				"complex add change remove in spec":                                true,
				"complex change conditions slice in status":                        true,
				"complex add status":                                               true,
				"complex change conditions slice and phase in status":              true,
			}),
		},
		{
			name: "skip status conditions",
			paths: [][]string{
				{
					"status",
					"conditions",
				},
			},
			maps: createTestTuplesComplex(hasDiffs{
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
				"complex change conditions slice in status":                        false,
				"complex add status":                                               true,
				"complex change conditions slice and phase in status":              true,
			}),
		},
		{
			name: "skip status",
			paths: [][]string{
				{
					"status",
				},
			},
			maps: createTestTuplesComplex(hasDiffs{
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
				"complex change conditions slice in status":                        false,
				"complex add status":                                               false,
				"complex change conditions slice and phase in status":              false,
			}),
		},
		{
			name: "skip status metadata and spec string",
			paths: [][]string{
				{
					"status",
				},
				{
					"spec",
					"string",
				},
				{
					"metadata",
				},
			},
			maps: createTestTuplesComplex(hasDiffs{
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
				"complex change spec string":                                       false,
				"complex change full deep spec intKey":                             true,
				"complex change spec innerInner.innerInnerKey":                     true,
				"complex change spec innerFirst.sub":                               true,
				"complex change multiple spec keys":                                true,
				"complex delete spec string":                                       false,
				"complex delete spec multiple keys":                                true,
				"complex add spec notExists key":                                   true,
				"complex add notExists key to spec innerInnerInner and innerFirst": true,
				"complex add change remove in spec":                                true,
				"complex change conditions slice in status":                        false,
				"complex add status":                                               false,
				"complex change conditions slice and phase in status":              false,
			}),
		},
		{
			name: "skip status and labels",
			paths: [][]string{
				{
					"status",
				},
				{
					"metadata",
					"labels",
				},
			},
			maps: createTestTuplesComplex(hasDiffs{
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
				"complex change conditions slice in status":                        false,
				"complex add status":                                               false,
				"complex change conditions slice and phase in status":              false,
			}),
		},
		{
			name: "skip spec string and name",
			paths: [][]string{
				{
					"spec",
					"string",
				},
				{
					"metadata",
					"name",
				},
			},
			maps: createTestTuplesComplex(hasDiffs{
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
				"complex change spec string":                                       false,
				"complex change full deep spec intKey":                             true,
				"complex change spec innerInner.innerInnerKey":                     true,
				"complex change spec innerFirst.sub":                               true,
				"complex change multiple spec keys":                                true,
				"complex delete spec string":                                       false,
				"complex delete spec multiple keys":                                true,
				"complex add spec notExists key":                                   true,
				"complex add notExists key to spec innerInnerInner and innerFirst": true,
				"complex add change remove in spec":                                true,
				"complex change conditions slice in status":                        true,
				"complex add status":                                               true,
				"complex change conditions slice and phase in status":              true,
			}),
		},
	}

	runMapStringComparatorTests(t, tests)
}

type testStringsComparator struct {
	name  string
	paths [][]string
	maps  []mapsTuple
}

func (tt *testStringsComparator) Name() string {
	return tt.name
}

func runMapStringComparatorTests(t *testing.T, tests []testStringsComparator) {
	for _, tt := range tests {
		run := func(t *testing.T) {
			runOverTuples(t, createTestStringsPaths(tt), tt.maps)
		}

		runTest(t, &tt, run)
	}
}

func createTestStringsPaths(tt testStringsComparator) []utils.MapPath {
	pathsComparators := make([]utils.MapPath, 0)
	for _, p := range tt.paths {
		pathsComparators = append(pathsComparators, utils.NewMapPathStringComparator(p...))
	}

	return pathsComparators
}
