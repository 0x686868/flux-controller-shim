/*
Copyright 2023 Hidde Beydals <yelling@hhh.computer>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package shim

import (
	"time"

	"github.com/fluxcd/source-controller/internal/cache"
)

// Cache is a shim for an internal cache.Cache.
type Cache *cache.Cache

// CacheRecorder is a shim for an internal cache.CacheRecorder.
type CacheRecorder *cache.CacheRecorder

// NewCache returns a new Cache with the given max items and interval.
func NewCache(maxItems int, interval time.Duration) Cache {
	return cache.New(maxItems, interval)
}

// MustMakeCacheMetrics returns a new CacheRecorder.
func MustMakeCacheMetrics() CacheRecorder {
	return cache.MustMakeMetrics()
}
