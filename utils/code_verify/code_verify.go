package code_verify

import (
	"fmt"

	"github.com/pangxianfei/framework/helpers/cache"
	"github.com/pangxianfei/framework/helpers/zone"

	"github.com/pangxianfei/framework/helpers/str"
)

const VALIDATION_CACHE_KEY = "TMAIC_VALIDATION_%s"

func validationCacheKey(index string) string {
	return fmt.Sprintf(VALIDATION_CACHE_KEY, index)
}

func Generate(index string, codeLen uint, expiredMinute uint) string {
	code := str.RandNumberString(codeLen)
	cache.Put(validationCacheKey(index), code, zone.Now().Add(zone.Duration(expiredMinute)*zone.Minute))
	return code
}

func Verify(index string, code string) bool {
	cacheCode := cache.Pull(validationCacheKey(index))
	if cacheCode == nil {
		return false
	}

	if cacheCode.(string) == code {
		return true
	}

	return false
}
