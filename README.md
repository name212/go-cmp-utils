# go-cmp-utils

Some utils for https://github.com/google/go-cmp/ library.

## Filter keys when compare two map[string]any

For filter keys (with deep map keys) by key path you can create cmp.Option with function `MapKeysFilter`.

For compare every key part you can use two comparators:
- `MapPathStringComparator` - match every key part as full equal two string 
- `MapPathReComparator` - match every key part as regexp match

## Examples 

See in [examples dir](./examples/).
