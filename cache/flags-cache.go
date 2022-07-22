package cache

// will we need a different data type for sdk keys ? currently no model

// FlagCache is implemented by redisCache struct
type FlagCache interface {
	Set(key string, value interface{}) // set an array of
	// Get(key string) interface{}        // not needed? mgr never reads from cache
	FlushAllAsync()
}
