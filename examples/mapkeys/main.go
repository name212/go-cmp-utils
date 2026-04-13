package main

import (
	"fmt"
	"regexp"

	"github.com/google/go-cmp/cmp"
	utils "github.com/name212/go-cmp-utils"
)

func main() {
	first := map[string]any{
		"root_first": map[string]any{
			"inner_first":  int(42),
			"inner_second": "val",
		},

		"root_second": map[string]any{
			"inner_first":  int(22),
			"inner_second": "another",
		},
	}

	second := map[string]any{
		"root_first": map[string]any{
			"inner_first":  int(12),
			"inner_second": "val",
		},

		"root_second": map[string]any{
			"inner_first":  int(32),
			"inner_second": "example",
		},
	}

	examples := []struct {
		name   string
		filter cmp.Option
	}{
		{
			name: "one key string",
			filter: utils.MapKeysFilter(
				utils.NewMapPathStringComparator("root_first", "inner_first"),
			),
		},
		{
			name: "multiple keys string",
			filter: utils.MapKeysFilter(
				utils.NewMapPathStringComparator("root_first", "inner_first"),
				utils.NewMapPathStringComparator("root_second", "inner_second"),
			),
		},
		{
			name: "multiple keys string no diff",
			filter: utils.MapKeysFilter(
				utils.NewMapPathStringComparator("root_first"),
				utils.NewMapPathStringComparator("root_second"),
			),
		},
		{
			name: "one key regexp",
			filter: utils.MapKeysFilter(
				mustCreateRegexpComparator("root_first", "inner(_|-)first"),
			),
		},
		{
			name: "multiple key regexp",
			filter: utils.MapKeysFilter(
				utils.NewMapPathReComparatorFromRe(
					regexp.MustCompile(".+_first"),
					regexp.MustCompile("inner_.*"),
				),
			),
		},
		{
			name: "multiple key one regexp no diff",
			filter: utils.MapKeysFilter(
				mustCreateRegexpComparator("root_.+"),
			),
		},
	}

	for _, e := range examples {
		printDiff(e.name, cmp.Diff(first, second, e.filter))
	}
}

func printDiff(m, diff string) {
	if diff == "" {
		diff = "no diff"
	}

	fmt.Printf("%s diff:\n%s\n\n", m, diff)
}

func mustCreateRegexpComparator(parts ...string) utils.MapPath {
	c, err := utils.NewMapPathReComparator(parts...)
	if err != nil {
		panic(err)
	}

	return c
}
