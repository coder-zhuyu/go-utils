package maps

import (
    "sort"
    "reflect"
)

type kv struct {
    Key     interface{}
    Value   interface{}
}

func SortMapAscByValue(m map[interface{}]interface{}) ([]kv) {
    kvs := make([]kv, 0, len(m))

    for k, v := range m {
        kvs = append(kvs, kv{k, v})
    }

    sort.Slice(kvs, func(i, j int) bool {
        if reflect.ValueOf(kvs[i].Value).Kind() == reflect.Int {
            return kvs[i].Value.(int) < kvs[j].Value.(int)
        }

        if reflect.ValueOf(kvs[i].Value).Kind() == reflect.Int64 {
            return kvs[i].Value.(int64) < kvs[j].Value.(int64)
        }

        if reflect.ValueOf(kvs[i].Value).Kind() == reflect.String {
            return kvs[i].Value.(string) < kvs[j].Value.(string)
        }

        return true

    })

    return kvs
}

func SortMapDescByValue(m map[interface{}]interface{}) ([]kv) {
    kvs := make([]kv, 0, len(m))

    for k, v := range m {
        kvs = append(kvs, kv{k, v})
    }

    sort.Slice(kvs, func(i, j int) bool {
        if reflect.ValueOf(kvs[i].Value).Kind() == reflect.Int {
            return kvs[i].Value.(int) > kvs[j].Value.(int)
        }

        if reflect.ValueOf(kvs[i].Value).Kind() == reflect.Int64 {
            return kvs[i].Value.(int64) > kvs[j].Value.(int64)
        }

        if reflect.ValueOf(kvs[i].Value).Kind() == reflect.String {
            return kvs[i].Value.(string) > kvs[j].Value.(string)
        }

        return true

    })

    return kvs
}
