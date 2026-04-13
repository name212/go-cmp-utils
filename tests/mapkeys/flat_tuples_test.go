// Copyright 2026
// license that can be found in the LICENSE file.

package tests

func createTestTuplesFlat(hasDiff hasDiffs) []mapsTuple {
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
			name:   "nil and empty",
			first:  nil,
			second: make(mapAny),
		},

		{
			name:   "empty and not",
			first:  make(mapAny),
			second: prepareOverJSON(flatMap),
		},

		{
			name:   "nil and not empty",
			first:  prepareOverJSON(flatMap),
			second: nil,
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

func flatHasDiffsEmpty() hasDiffs {
	return hasDiffs{
		"both nil":                  false,
		"both empty":                false,
		"nil and empty":             true,
		"empty and not":             true,
		"nil and not empty":         true,
		"same flat":                 false,
		"change flat first":         true,
		"change flat int":           true,
		"change flat slice":         true,
		"add notExists to flat":     true,
		"change flat int and first": true,
		"remove from flat int":      true,
	}
}

func flatHasDiffsFirst() hasDiffs {
	return hasDiffs{
		"both nil":                  false,
		"both empty":                false,
		"nil and empty":             true,
		"empty and not":             true,
		"nil and not empty":         true,
		"same flat":                 false,
		"change flat first":         false,
		"change flat int":           true,
		"change flat slice":         true,
		"change flat int and first": true,
		"add notExists to flat":     true,
		"remove from flat int":      true,
	}
}

func flatHasDiffsInt() hasDiffs {
	return hasDiffs{
		"both nil":                  false,
		"both empty":                false,
		"nil and empty":             true,
		"empty and not":             true,
		"nil and not empty":         true,
		"same flat":                 false,
		"change flat first":         true,
		"change flat int":           false,
		"change flat slice":         true,
		"change flat int and first": true,
		"add notExists to flat":     true,
		"remove from flat int":      false,
	}
}

func flatHasDiffsNotExists() hasDiffs {
	return hasDiffs{
		"both nil":                  false,
		"both empty":                false,
		"nil and empty":             true,
		"empty and not":             true,
		"nil and not empty":         true,
		"same flat":                 false,
		"change flat first":         true,
		"change flat int":           true,
		"change flat slice":         true,
		"change flat int and first": true,
		"add notExists to flat":     false,
		"remove from flat int":      true,
	}
}

func flatHasDiffsFirstInt() hasDiffs {
	return hasDiffs{
		"both nil":                  false,
		"both empty":                false,
		"nil and empty":             true,
		"empty and not":             true,
		"nil and not empty":         true,
		"same flat":                 false,
		"change flat first":         false,
		"change flat int":           false,
		"change flat slice":         true,
		"change flat int and first": false,
		"add notExists to flat":     true,
		"remove from flat int":      false,
	}
}

func flatHasDiffsAll() hasDiffs {
	return hasDiffs{
		"both nil":                  false,
		"both empty":                false,
		"nil and empty":             true,
		"empty and not":             false,
		"nil and not empty":         true,
		"same flat":                 false,
		"change flat first":         false,
		"change flat int":           false,
		"change flat slice":         false,
		"change flat int and first": false,
		"add notExists to flat":     false,
		"remove from flat int":      false,
	}
}
