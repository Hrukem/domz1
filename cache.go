// Package domz1 implementation of data caching
package domz1

type Cache struct {
	CacheMap map[string]interface{}
}

func NewCache() Cache {
	return Cache{CacheMap: map[string]interface{}{}}
}

// Set function implementation add data in cache
func (m Cache) Set(key string, value interface{}) {
	m.CacheMap[key] = value
}

// Get function implementation get data from cache
func (m Cache) Get(key string) interface{} {
	return m.CacheMap[key]
}

// Delete function implementation delete data from cache
func (m Cache) Delete(key string) {
	delete(m.CacheMap, key)
}
