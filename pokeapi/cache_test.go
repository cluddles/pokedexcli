package pokeapi

import (
	"testing"
	"time"
)

type testCase struct {
	toAdd    []input
	expected []result
}

type input struct {
	key string
	str string
}

type result struct {
	key    string
	exists bool
	str    string
}

func TestAddAndGet(t *testing.T) {
	cases := []testCase{
		{
			toAdd: []input{
				{key: "key1", str: "12345"},
				{key: "key3", str: "44 55 66"},
				{key: "key3", str: "overwrite"},
			},
			expected: []result{
				{key: "key1", exists: true, str: "12345"},
				{key: "key2", exists: false},
				{key: "key3", exists: true, str: "overwrite"},
			},
		},
	}

	for i, c := range cases {
		cache := NewCache(time.Second*100, time.Second*100)
		for _, v := range c.toAdd {
			cache.Add(v.key, []byte(v.str))
		}

		for _, v := range c.expected {
			data, exists := cache.Get(v.key)
			if exists != v.exists {
				t.Errorf("testcase %d Get('%s'): exists=%v but expected %v", i, v.key, exists, v.exists)
			}

			if exists {
				str := string(data)
				if str != v.str {
					t.Errorf("testcase %d Get('%s'): data '%v' but expected %v", i, v.key, str, v.str)
				}
			}
		}
	}

}

func TestReapLoop(t *testing.T) {
	const ttl = 5 * time.Millisecond
	cache := NewCache(ttl, time.Millisecond*1)
	cache.Add("https://example.com", []byte("testdata"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Errorf("expected to find key")
		return
	}

	time.Sleep(ttl * 2)

	_, ok = cache.Get("https://example.com")
	if ok {
		t.Errorf("expected to not find key")
		return
	}
}
