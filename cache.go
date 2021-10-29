// Package domz1 implementation of data caching
package domz1

type Cache struct {
	cacheMap map[string]interface{}
}

func NewCache() Cache {
	return Cache{cacheMap: map[string]interface{}{}}
}

// Set function implementation add data in cache
func (m Cache) Set(key string, value interface{}) {
	m.cacheMap[key] = value
}

// Get function implementation get data from cache
func (m Cache) Get(key string) interface{} {
	return m.cacheMap[key]
}

// Delete function implementation delete data from cache
func (m Cache) Delete(key string) {
	delete(m.cacheMap, key)
}
