package cache

import (
	"testing"
)

func TestCommonCacheService(t *testing.T) {
	var err error
	err = CommonCacheService().Set("11111", 1)
	if err != nil {
		t.Error(err)
		return
	}
	err = CommonCacheService().Set("22222", 1)
	if err != nil {
		t.Error(err)
		return
	}
	err = CommonCacheService().Set("33333", 1)
	if err != nil {
		t.Error(err)
		return
	}

	CommonCacheService().RemoveAll()

	t.Log(CommonCacheService().Len())
}
