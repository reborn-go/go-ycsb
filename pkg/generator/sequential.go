// Copyright 2018 PingCAP, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// See the License for the specific language governing permissions and
// limitations under the License.

/**
 * Copyright (c) 2010-2016 Yahoo! Inc., 2017 YCSB contributors. All rights reserved.
 * <p>
 * Licensed under the Apache License, Version 2.0 (the "License"); you
 * may not use this file except in compliance with the License. You
 * may obtain a copy of the License at
 * <p>
 * http://www.apache.org/licenses/LICENSE-2.0
 * <p>
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or
 * implied. See the License for the specific language governing
 * permissions and limitations under the License. See accompanying
 * LICENSE file.
 */

package generator

import (
	"math/rand"
	"sync/atomic"
)

// Sequential generates a sequence of integers 0, 1, ...
type Sequential struct {
	counter  int64
	interval int64
	start    int64
}

// NewSequential creates a counter that starts at countStart.
func NewSequential(countStart int64, countEnd int64) *Sequential {
	return &Sequential{
		counter:  0,
		start:    countStart,
		interval: countEnd - countStart + 1,
	}
}

// Next implements the Generator Next interface.
func (s *Sequential) Next(_ *rand.Rand) int64 {
	n := s.start + atomic.AddInt64(&s.counter, 1)%s.interval
	return n
}

// Last implements the Generator Last interface.
func (s *Sequential) Last() int64 {
	return atomic.LoadInt64(&s.counter) + 1
}
