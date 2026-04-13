// Copyright 2026
// license that can be found in the LICENSE file.

package tests

func createTestTuplesComplex(hasDiff hasDiffs) []mapsTuple {
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

		{
			name:  "complex change conditions slice in status",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)
				s := extractSlice[mapAny](m,
					"status",
					"conditions",
				)

				s[0] = mapAny{
					"lastTransitionTime": "2026-03-11T08:11:21Z",
					"status":             "False",
					"type":               "Mount",
				}

				s = append(s, mapAny{
					"lastTransitionTime": "2024-12-01T18:44:02Z",
					"status":             "True",
					"type":               "Ready",
				})

				return setFieldSlice(m, s,
					"status",
					"conditions",
				)
			}(),
		},

		{
			name:  "complex change conditions slice and phase in status",
			first: prepareOverJSON(complexMap),
			second: func() mapAny {
				m := prepareOverJSON(complexMap)

				m = setField(m, "Running",
					"status",
					"phase",
				)

				s := extractSlice[mapAny](m,
					"status",
					"conditions",
				)

				s[0] = mapAny{
					"lastTransitionTime": "2026-03-11T08:11:21Z",
					"status":             "False",
					"type":               "Mount",
				}

				return setFieldSlice(m, s,
					"status",
					"conditions",
				)
			}(),
		},

		{
			name: "complex add status",
			first: func() mapAny {
				m := prepareOverJSON(complexMap)
				return deleteField(m, "status")
			}(),
			second: prepareOverJSON(complexMap),
		},
	}

	return applyHasDiff(hasDiff, tuples)
}

func complexHasDiffsEmpty() hasDiffs {
	return hasDiffs{
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
	}
}

func complexHasDiffsKind() hasDiffs {
	return hasDiffs{
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
		"complex change conditions slice in status":                        true,
		"complex add status":                                               true,
		"complex change conditions slice and phase in status":              true,
	}
}

func complexHasDiffsName() hasDiffs {
	return hasDiffs{
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
		"complex change conditions slice in status":                        true,
		"complex add status":                                               true,
		"complex change conditions slice and phase in status":              true,
	}
}

func complexHasDiffsKindName() hasDiffs {
	return hasDiffs{
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
		"complex change conditions slice in status":                        true,
		"complex add status":                                               true,
		"complex change conditions slice and phase in status":              true,
	}
}

func complexHasDiffsLabels() hasDiffs {
	return hasDiffs{
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
		"complex change conditions slice in status":                        true,
		"complex add status":                                               true,
		"complex change conditions slice and phase in status":              true,
	}
}

func complexHasDiffsAnnotations() hasDiffs {
	return hasDiffs{
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
		"complex change conditions slice in status":                        true,
		"complex add status":                                               true,
		"complex change conditions slice and phase in status":              true,
	}
}

func complexHasDiffsOneAnnotation() hasDiffs {
	return hasDiffs{
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
		"complex change conditions slice in status":                        true,
		"complex add status":                                               true,
		"complex change conditions slice and phase in status":              true,
	}
}

func complexHasDiffsMultipleLabels() hasDiffs {
	return hasDiffs{
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
		"complex change conditions slice in status":                        true,
		"complex add status":                                               true,
		"complex change conditions slice and phase in status":              true,
	}
}

func complexHasDiffsAnnotationsAndLabels() hasDiffs {
	return hasDiffs{
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
		"complex change conditions slice in status":                        true,
		"complex add status":                                               true,
		"complex change conditions slice and phase in status":              true,
	}
}

func complexHasDiffsAnnotationsLabelsName() hasDiffs {
	return hasDiffs{
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
	}
}

func complexHasDiffsSpecString() hasDiffs {
	return hasDiffs{
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
		"complex add spec notExists key":                                   true,
		"complex add notExists key to spec innerInnerInner and innerFirst": true,
		"complex add change remove in spec":                                true,
		"complex change conditions slice in status":                        true,
		"complex add status":                                               true,
		"complex change conditions slice and phase in status":              true,
	}
}

func complexHasDiffsInnerFirstSub() hasDiffs {
	return hasDiffs{
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
		"complex change spec innerFirst.sub":                               false,
		"complex change multiple spec keys":                                true,
		"complex delete spec string":                                       true,
		"complex delete spec multiple keys":                                true,
		"complex add spec notExists key":                                   true,
		"complex add notExists key to spec innerInnerInner and innerFirst": true,
		"complex add change remove in spec":                                true,
		"complex change conditions slice in status":                        true,
		"complex add status":                                               true,
		"complex change conditions slice and phase in status":              true,
	}
}
