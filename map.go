// Copyright 2026
// license that can be found in the LICENSE file.

package gocmputils

import (
	"fmt"
	"regexp"

	"github.com/google/go-cmp/cmp"
)

// MapPathComparator
// interface for compare key parts
type MapPathComparator interface {
	// Compare
	// compare key part with index
	Compare(indx int, keyPathPart string) bool
	// Parts
	// returns number of parts
	Parts() int
}

// MapPath
// alias to MapPathComparator
type MapPath MapPathComparator

// MapKeysFilter
// returns go-cmp option for ignore golang map keys by path
// Every MapPath are paths to keys
// With it you can filter multiple keys by own path
// For create MapPath you can use
// NewMapPathStringComparator
// MapPathReComparator
func MapKeysFilter(paths ...MapPath) cmp.Option {
	if len(paths) == 0 {
		return filterOption(noFilter)
	}

	filters := make([]filter, 0, len(paths))
	for _, p := range paths {
		filters = append(filters, constructFilter(p))
	}

	composeFilter := func(p cmp.Path) bool {
		for _, f := range filters {
			if f(p) {
				return true
			}
		}

		return false
	}

	return filterOption(composeFilter)
}

// MapPathStringComparator
// implementation of MapPathComparator
// This comparator compare every path part
// as full string equals
type MapPathStringComparator []string

// NewMapPathStringComparator
// creates MapPathStringComparator
// parts is full path to key for filter
// for example, for filter "second" key in map:
//
//	map[string]any {
//	  "root": map[string]any{
//	    "first": 42,
//	    "second": "val",
//	  }
//	}
//
// you can use next call
// NewMapPathStringComparator("root", "second")
func NewMapPathStringComparator(parts ...string) MapPathStringComparator {
	res := make(MapPathStringComparator, 0, len(parts))
	res = append(res, parts...)

	return res
}

func (c MapPathStringComparator) Parts() int {
	return len(c)
}

func (c MapPathStringComparator) Compare(indx int, keyPathPart string) bool {
	if indx >= c.Parts() {
		return false
	}

	return c[indx] == keyPathPart
}

var (
	// SkipNotEmptyRe
	// regexp for filter any not empty key
	SkipNotEmptyRe = regexp.MustCompile(`.+`)
	// SkipWithEmptyRe
	// regexp for filter any key (include empty)
	SkipWithEmptyRe = regexp.MustCompile(`.*`)
)

// RepeatCompileRe
// compile regexp and create list of regexp with count len
func RepeatCompileRe(r string, count int) ([]*regexp.Regexp, error) {
	re, err := regexp.Compile(r)
	if err != nil {
		return nil, err
	}

	res := make([]*regexp.Regexp, count)
	for i := range count {
		res[i] = re
	}

	return res, nil
}

// RepeatRe
// create list of regexp with count len
func RepeatRe(r *regexp.Regexp, count int) []*regexp.Regexp {
	res := make([]*regexp.Regexp, count)
	for i := range count {
		res[i] = r
	}

	return res
}

// MapPathReComparator
// implementation of MapPathComparator
// This comparator compare every path part as regexp match
type MapPathReComparator []*regexp.Regexp

// NewMapPathReComparator
// creates MapPathReComparator
// parts is regexp compatibility strings
// all parts will compile to regexp
// for example, for filter all keys "second" and "first" key
// as single regexp in map:
//
//	map[string]any {
//	  "root": map[string]any{
//	    "first": 42,
//	    "second": "val",
//	    "another": []string{"hello"}
//	  }
//	}
//
// you can use next call
// NewMapPathReComparator("root", "^(second|first)$")
// function returns all compile errors as single error
func NewMapPathReComparator(parts ...string) (MapPathReComparator, error) {
	res := make(MapPathReComparator, 0, len(parts))

	var errors []error
	for i, pp := range parts {
		re, err := regexp.Compile(pp)
		if err != nil {
			errors = append(errors, fmt.Errorf("failed to compile part %d: %w", i, err))
			continue
		}
		res = append(res, re)
	}

	if len(errors) == 0 {
		return res, nil
	}

	resError := errors[0]
	for i := 1; i < len(errors); i++ {
		resError = fmt.Errorf("%w\n%w", resError, errors[i])
	}

	return nil, resError
}

// NewMapPathReComparator
// creates MapPathReComparator
// parts is regexp instances
func NewMapPathReComparatorFromRe(parts ...*regexp.Regexp) MapPathReComparator {
	res := make(MapPathReComparator, 0, len(parts))
	res = append(res, parts...)

	return res
}

func (c MapPathReComparator) Parts() int {
	return len(c)
}

func (c MapPathReComparator) Compare(indx int, keyPathPart string) bool {
	if indx >= c.Parts() {
		return false
	}

	return c[indx].MatchString(keyPathPart)
}

type filter = func(p cmp.Path) bool

func noFilter(p cmp.Path) bool {
	return false
}

func filterOption(f filter) cmp.Option {
	return cmp.FilterPath(f, cmp.Ignore())
}

func constructFilter(path MapPath) filter {
	l := path.Parts()

	if l == 0 {
		return noFilter
	}

	return func(p cmp.Path) bool {
		if len(p) != 2*l {
			return false
		}

		result := true

		for i := range l {
			num := (2 * i) + 1
			index, ok := p.Index(num).(cmp.MapIndex)
			if !ok {
				result = false
				break
			}

			if !path.Compare(i, index.Key().String()) {
				result = false
				break
			}
		}

		return result
	}
}
