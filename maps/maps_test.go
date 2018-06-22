package maps

import "testing"

func TestSortMapAscByValue(t *testing.T) {
    m := map[interface{}]interface{} {
        "one": 1,
        "two": 2,
        "four": 4,
        "three": 3,
    }

    kvs := SortMapAscByValue(m)

    for _, kv := range kvs {
        t.Logf("%s: %d", kv.Key, kv.Value)
    }
}

func TestSortMapDescByValue(t *testing.T) {
    m := map[interface{}]interface{} {
        "one": 1,
        "two": 2,
        "four": 4,
        "three": 3,
    }

    kvs := SortMapDescByValue(m)

    for _, kv := range kvs {
        t.Logf("%s: %d", kv.Key, kv.Value)
    }
}