package cache

import (
	"manager/models"
)
// will we need a different data type for sdk keys ? currently no model

// FlagCache is implemented by redisCache struct
type FlagCache interface {
	Set(key string, value *models.Flag) // set an array of 
	Get(key string) *models.Flag // not needed? mgr never reads from cache
	FlushAllAsync()
}