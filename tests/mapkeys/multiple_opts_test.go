// Copyright 2026
// license that can be found in the LICENSE file.

package tests

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	utils "github.com/name212/go-cmp-utils"
	"github.com/stretchr/testify/require"
)

func TestMultipleOptionsWithIgnore(t *testing.T) {
	first := map[string]any{
		"root_first": map[string]any{
			"inner_first":  int(42),
			"inner_second": "val",
		},

		"root_second": map[string]any{
			"inner_first":  int(32),
			"inner_second": "another",
		},
		"slice_int":    []int{22, 44, 66},
		"slice_string": []string{"a", "b", "c"},
	}

	second := map[string]any{
		"root_first": map[string]any{
			"inner_first":  int(12),
			"inner_second": "val",
		},

		"root_second": map[string]any{
			"inner_first":  int(32),
			"inner_second": "another",
		},
		"slice_int":    []int{44, 22, 66},
		"slice_string": []string{"a", "b", "c"},
	}

	runComparatorMultipleOptsTests := func(t *testing.T, tests []testMultipleOpts) {
		for _, tt := range tests {
			run := func(t *testing.T) {
				diff := cmp.Diff(first, second, tt.opts...)
				require.Empty(t, diff, "should not diff")
			}

			runTest(t, &tt, run)
		}

	}

	filterOpt := utils.MapKeysFilter(
		utils.NewMapPathStringComparator("root_first", "inner_first"),
	)

	t.Run("transform", func(t *testing.T) {
		sortIntSlicesOpt := cmpopts.SortSlices(func(f, s any) int {
			fInt, ok := f.(int)
			if !ok {
				return 0
			}

			sInt, ok := s.(int)
			if !ok {
				return 0
			}

			return fInt - sInt
		})

		sortStringSlicesOpt := cmpopts.SortSlices(func(f, s any) int {
			fStr, ok := f.(string)
			if !ok {
				return 0
			}

			sStr, ok := s.(string)
			if !ok {
				return 0
			}

			if fStr == sStr {
				return 0
			}

			if fStr < sStr {
				return -1
			}

			return 1
		})

		tests := []testMultipleOpts{
			{
				name: "first",
				opts: []cmp.Option{filterOpt, sortIntSlicesOpt},
			},
			{
				name: "last",
				opts: []cmp.Option{sortIntSlicesOpt, filterOpt},
			},
			{
				name: "in middle",
				opts: []cmp.Option{sortIntSlicesOpt, filterOpt, sortStringSlicesOpt},
			},
		}

		runComparatorMultipleOptsTests(t, tests)
	})
}

type testMultipleOpts struct {
	name string
	opts []cmp.Option
}

func (tt *testMultipleOpts) Name() string {
	return tt.name
}
