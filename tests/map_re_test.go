// Copyright 2026
// license that can be found in the LICENSE file.

package tests

import (
	"fmt"
	"regexp"
	"strings"
	"testing"

	utils "github.com/name212/go-cmp-utils"
	"github.com/stretchr/testify/require"
)

func TestMapReComparatorFlat(t *testing.T) {
	tests := []testReComparator{
		{
			name:  "empty paths",
			paths: nil,
			maps:  createTestTuplesFlat(flatHasDiffsEmpty()),
		},
		{
			name: "skip first as string",
			paths: [][]string{
				{"first"},
			},
			maps: createTestTuplesFlat(flatHasDiffsFirst()),
		},
		{
			name: "skip first as re",
			paths: [][]string{
				{
					"fi[a-z]{2}t",
				},
			},
			maps: createTestTuplesFlat(flatHasDiffsFirst()),
		},
		{
			name: "skip int as string",
			paths: [][]string{
				{"int"},
			},
			maps: createTestTuplesFlat(flatHasDiffsInt()),
		},
		{
			name: "skip int as re",
			paths: [][]string{
				{
					"in.+",
				},
			},
			maps: createTestTuplesFlat(flatHasDiffsInt()),
		},
		{
			name: "skip notExists as string",
			paths: [][]string{
				{"notExists"},
			},
			maps: createTestTuplesFlat(flatHasDiffsNotExists()),
		},
		{
			name: "skip notExists as re",
			paths: [][]string{
				{
					"not(E|e)xists",
				},
			},
			maps: createTestTuplesFlat(flatHasDiffsNotExists()),
		},
		{
			name: "skip first and int as re",
			paths: [][]string{
				{
					"f.rst",
				},
				{
					"^[tni]{3}$",
				},
			},
			maps: createTestTuplesFlat(flatHasDiffsFirstInt()),
		},
		{
			name: "skip first and int as one re",
			paths: [][]string{
				{
					"^(first|i[nt]{2})$",
				},
			},
			maps: createTestTuplesFlat(flatHasDiffsFirstInt()),
		},
		{
			name: "skip first and int as string",
			paths: [][]string{
				{"first"},
				{"int"},
			},
			maps: createTestTuplesFlat(flatHasDiffsFirstInt()),
		},
		{
			name: "skip all as strings",
			paths: [][]string{
				{"first"},
				{"int"},
				{"slice"},
				{"notExists"},
			},
			maps: createTestTuplesFlat(flatHasDiffsAll()),
		},
		{
			name: "skip all as re not empty",
			pathsRe: [][]*regexp.Regexp{
				{utils.SkipNotEmptyRe},
			},
			maps: createTestTuplesFlat(flatHasDiffsAll()),
		},
		{
			name: "skip all as re with empty",
			pathsRe: [][]*regexp.Regexp{
				{utils.SkipWithEmptyRe},
			},
			maps: createTestTuplesFlat(flatHasDiffsAll()),
		},
		{
			name: "skip all as one re string",
			paths: [][]string{
				{
					".+",
				},
			},
			maps: createTestTuplesFlat(flatHasDiffsAll()),
		},
	}

	runMapReComparatorTests(t, tests)
}

func TestMapReComparatorComplex(t *testing.T) {
	tests := []testReComparator{
		{
			name:  "empty paths",
			paths: nil,
			maps:  createTestTuplesComplex(complexHasDiffsEmpty()),
		},

		{
			name: "skip kind as re",
			paths: [][]string{
				{"k..d"},
			},
			maps: createTestTuplesComplex(complexHasDiffsKind()),
		},

		{
			name: "skip name as re's",
			paths: [][]string{
				{"meta[a-z]{2}ta", "^name$"},
			},
			maps: createTestTuplesComplex(complexHasDiffsName()),
		},

		{
			name: "skip name and kind as string and re",
			paths: [][]string{
				{"metadata", "n[abc]{1}me"},
				{"kind"},
			},
			maps: createTestTuplesComplex(complexHasDiffsKindName()),
		},

		{
			name: "skip labels as re",
			paths: [][]string{
				{"metadat.+", "lab(e|L)ls"},
			},
			maps: createTestTuplesComplex(complexHasDiffsLabels()),
		},

		{
			name: "skip annotations as re and string",
			paths: [][]string{
				{"metadata", "a[n]{2}otations"},
			},
			maps: createTestTuplesComplex(complexHasDiffsAnnotations()),
		},

		{
			name: "skip one annotations as re's",
			paths: [][]string{
				{"met(a|A)data", "annotation.+", "test.example.com/fo.+"},
			},
			maps: createTestTuplesComplex(complexHasDiffsOneAnnotation()),
		},

		{
			name: "skip multiple labels as re's and strings",
			paths: [][]string{
				{"meta[a-z]+", "labels", "test.example.com/first"},
				{"metadata", "la..ls", "notExists"},
			},
			maps: createTestTuplesComplex(complexHasDiffsMultipleLabels()),
		},

		{
			name: "skip annotations and labels as one re",
			paths: [][]string{
				{"metadata", "^(labels|annotations)$"},
			},
			maps: createTestTuplesComplex(complexHasDiffsAnnotationsAndLabels()),
		},

		{
			name: "skip annotations and labels and name with two re's",
			paths: [][]string{
				{"metadata", "(annotations|name)"},
				{"metadat.", "labels"},
			},
			maps: createTestTuplesComplex(complexHasDiffsAnnotationsLabelsName()),
		},

		{
			name: "skip spec string as re",
			paths: [][]string{
				{"spe.+", "(s|A)tring"},
			},
			maps: createTestTuplesComplex(complexHasDiffsSpecString()),
		},

		{
			name: "skip spec innerInnerKey with re repeat compile",
			pathsRe: func() [][]*regexp.Regexp {
				innerRepeat, err := utils.RepeatCompileRe(`.+Inner.*`, 2)
				if err != nil {
					panic("Cannot compile repeat for innerInnerKey")
				}

				res := []*regexp.Regexp{
					regexp.MustCompile("spec"),
					regexp.MustCompile("inner.nner"),
				}
				res = append(res, innerRepeat...)

				return [][]*regexp.Regexp{res}
			}(),
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
				"complex change conditions slice in status":                        true,
				"complex add status":                                               true,
				"complex change conditions slice and phase in status":              true,
			}),
		},
		{
			name: "skip spec innerFirst.sub as re",
			paths: [][]string{
				{
					"s.ec",
					"inn..First",
					"[sub]{3}",
				},
			},
			maps: createTestTuplesComplex(complexHasDiffsInnerFirstSub()),
		},

		{
			name: "skip spec multiple changes full as re",
			pathsRe: [][]*regexp.Regexp{
				{
					regexp.MustCompile("spec"),
					utils.SkipNotEmptyRe,
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
				"complex change spec innerInner.innerInnerKey":                     false,
				"complex change spec innerFirst.sub":                               false,
				"complex change multiple spec keys":                                false,
				"complex delete spec string":                                       false,
				"complex delete spec multiple keys":                                false,
				"complex add spec notExists key":                                   false,
				"complex add notExists key to spec innerInnerInner and innerFirst": false,
				"complex add change remove in spec":                                false,
				"complex change conditions slice in status":                        true,
				"complex add status":                                               true,
				"complex change conditions slice and phase in status":              true,
			}),
		},

		{
			name: "skip status all as re and string",
			paths: [][]string{
				{
					"status",
					`.+`,
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
				"complex change conditions slice and phase in status":              false,
			}),
		},
		{
			name: "skip full as re",
			pathsRe: [][]*regexp.Regexp{
				{
					utils.SkipWithEmptyRe,
				},
			},

			maps: createTestTuplesComplex(hasDiffs{
				"both nil":                                                         false,
				"both empty":                                                       false,
				"same complex":                                                     false,
				"change complex kind":                                              false,
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
				"complex change full deep spec intKey":                             false,
				"complex change spec innerInner.innerInnerKey":                     false,
				"complex change spec innerFirst.sub":                               false,
				"complex change multiple spec keys":                                false,
				"complex delete spec string":                                       false,
				"complex delete spec multiple keys":                                false,
				"complex add spec notExists key":                                   false,
				"complex add notExists key to spec innerInnerInner and innerFirst": false,
				"complex add change remove in spec":                                false,
				"complex change conditions slice in status":                        false,
				"complex add status":                                               false,
				"complex change conditions slice and phase in status":              false,
			}),
		},
	}

	runMapReComparatorTests(t, tests)
}

func TestMapReComparatorErrors(t *testing.T) {
	t.Run("one error", func(t *testing.T) {
		tests := []testReComparatorError{
			{
				name:  "one",
				paths: []string{"*[aaasw"},
				errors: []string{
					"failed to compile part 0: error parsing regexp: missing argument to repetition operator: `*`",
				},
			},
			{
				name: "middle",
				paths: []string{
					".+",
					"*[aaasw",
					"aaa\\-.+",
				},
				errors: []string{
					"failed to compile part 1: error parsing regexp: missing argument to repetition operator: `*`",
				},
			},
			{
				name: "last",
				paths: []string{
					".+",
					"aaa\\-.+",
					"*[aaasw",
				},
				errors: []string{
					"failed to compile part 2: error parsing regexp: missing argument to repetition operator: `*`",
				},
			},
		}

		runMapReComparatorTestsErrors(t, tests)
	})

	t.Run("multiple errors", func(t *testing.T) {
		tests := []testReComparatorError{
			{
				name: "in first",
				paths: []string{
					".+-*[",
					"aaa\\-.+",
					"*[aaasw",
					".*",
				},
				errors: []string{
					"failed to compile part 0: error parsing regexp: missing closing ]: `[`",
					"failed to compile part 2: error parsing regexp: missing argument to repetition operator: `*`",
				},
			},
			{
				name: "in middle",
				paths: []string{
					"aaa\\-.+",
					".+-*[",
					"*[aaasw",
					".*",
				},
				errors: []string{
					"failed to compile part 1: error parsing regexp: missing closing ]: `[`",
					"failed to compile part 2: error parsing regexp: missing argument to repetition operator: `*`",
				},
			},
			{
				name: "with last",
				paths: []string{
					".+-*[",
					"aaa\\-.+",
					"*[aaasw",
				},
				errors: []string{
					"failed to compile part 0: error parsing regexp: missing closing ]: `[`",
					"failed to compile part 2: error parsing regexp: missing argument to repetition operator: `*`",
				},
			},
		}

		runMapReComparatorTestsErrors(t, tests)
	})
}

type testReComparatorError struct {
	name   string
	paths  []string
	errors []string
}

func (tt *testReComparatorError) Name() string {
	return tt.name
}

func runMapReComparatorTestsErrors(t *testing.T, tests []testReComparatorError) {
	for _, tt := range tests {
		run := func(t *testing.T) {
			comparator, err := utils.NewMapPathReComparator(tt.paths...)
			require.Error(t, err, "should return errors")
			require.Nil(t, comparator, "comparator should be nil")
			errors := strings.Split(err.Error(), "\n")
			require.Len(t, errors, len(tt.errors), "not all errors present")
			require.Equal(t, tt.errors, errors, "all errors should equal")
		}
		
		runTest(t, &tt, run)
	}
}

type testReComparator struct {
	name    string
	paths   [][]string
	pathsRe [][]*regexp.Regexp
	maps    []mapsTuple
}

func (tt *testReComparator) Name() string {
	return tt.name
}

func runMapReComparatorTests(t *testing.T, tests []testReComparator) {
	for _, tt := range tests {
		run := func(t *testing.T) {
			t.Run("from strings", func(t *testing.T) {
				runOverTuples(t, createTestReStringsPaths(tt), tt.maps)
			})

			t.Run("from regexp", func(t *testing.T) {
				runOverTuples(t, createTestRePaths(tt), tt.maps)
			})
		}

		runTest(t, &tt, run)
	}
}

func createTestReComparators(tt testReComparator) []utils.MapPath {
	if len(tt.pathsRe) == 0 {
		return nil
	}

	pathsComparatorsRe := make([]utils.MapPath, 0, len(tt.pathsRe))
	for _, r := range tt.pathsRe {
		pathsComparatorsRe = append(pathsComparatorsRe, utils.NewMapPathReComparatorFromRe(r...))
	}

	return pathsComparatorsRe
}

func createTestReStringsPaths(tt testReComparator) []utils.MapPath {
	if pathsComparatorsRe := createTestReComparators(tt); len(pathsComparatorsRe) > 0 {
		return pathsComparatorsRe
	}

	pathsComparators := make([]utils.MapPath, 0)
	for i, p := range tt.paths {
		comparator, err := utils.NewMapPathReComparator(p...)
		if err != nil {
			panic(fmt.Errorf("cannot create re comparator %d: %v", i, err))
		}
		pathsComparators = append(pathsComparators, comparator)
	}

	return pathsComparators
}

func createTestRePaths(tt testReComparator) []utils.MapPath {
	if pathsComparatorsRe := createTestReComparators(tt); len(pathsComparatorsRe) > 0 {
		return pathsComparatorsRe
	}

	pathsComparators := make([]utils.MapPath, 0)
	for i, pp := range tt.paths {
		rees := make([]*regexp.Regexp, 0, len(pp))
		for j, p := range pp {
			r, err := regexp.Compile(p)
			if err != nil {
				panic(fmt.Errorf("cannot compile re %d/%d: %s: %v", i, j, p, err))
			}
			rees = append(rees, r)
		}

		comparator := utils.NewMapPathReComparatorFromRe(rees...)
		pathsComparators = append(pathsComparators, comparator)
	}

	return pathsComparators
}
