// Copyright 2026
// license that can be found in the LICENSE file.

package gocmputils

import (
	"fmt"
	"regexp"

	"github.com/google/go-cmp/cmp"
)

type MapPathComparator interface {
	Compare(indx int, keyPathPart string) bool
	Parts() int
}

type MapPath MapPathComparator

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

type MapPathStringComparator []string

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

type MapPathReComparator []*regexp.Regexp

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
