/*
 * Copyright (c) 2014, Yawning Angel <yawning at torproject dot org>
 * All rights reserved.
 *
 * Redistribution and use in source and binary forms, with or without
 * modification, are permitted provided that the following conditions are met:
 *
 *  * Redistributions of source code must retain the above copyright notice,
 *    this list of conditions and the following disclaimer.
 *
 *  * Redistributions in binary form must reproduce the above copyright notice,
 *    this list of conditions and the following disclaimer in the documentation
 *    and/or other materials provided with the distribution.
 *
 * THIS SOFTWARE IS PROVIDED BY THE COPYRIGHT HOLDERS AND CONTRIBUTORS "AS IS"
 * AND ANY EXPRESS OR IMPLIED WARRANTIES, INCLUDING, BUT NOT LIMITED TO, THE
 * IMPLIED WARRANTIES OF MERCHANTABILITY AND FITNESS FOR A PARTICULAR PURPOSE
 * ARE DISCLAIMED. IN NO EVENT SHALL THE COPYRIGHT HOLDER OR CONTRIBUTORS BE
 * LIABLE FOR ANY DIRECT, INDIRECT, INCIDENTAL, SPECIAL, EXEMPLARY, OR
 * CONSEQUENTIAL DAMAGES (INCLUDING, BUT NOT LIMITED TO, PROCUREMENT OF
 * SUBSTITUTE GOODS OR SERVICES; LOSS OF USE, DATA, OR PROFITS; OR BUSINESS
 * INTERRUPTION) HOWEVER CAUSED AND ON ANY THEORY OF LIABILITY, WHETHER IN
 * CONTRACT, STRICT LIABILITY, OR TORT (INCLUDING NEGLIGENCE OR OTHERWISE)
 * ARISING IN ANY WAY OUT OF THE USE OF THIS SOFTWARE, EVEN IF ADVISED OF THE
 * POSSIBILITY OF SUCH DAMAGE.
 */

package obfs4

import (
	csrand "crypto/rand"
	"fmt"
	"math/big"
	"math/rand"
)

var (
	csRandSourceInstance csRandSource
	csRandInstance       = rand.New(csRandSourceInstance)
)

type csRandSource struct {
	// This does not keep any state as it is backed by crypto/rand.
}

func (r csRandSource) Int63() int64 {
	ret, err := csrand.Int(csrand.Reader, big.NewInt(int64((1<<63)-1)))
	if err != nil {
		panic(err)
	}

	return ret.Int64()
}

func (r csRandSource) Seed(seed int64) {
	// No-op.
}

func randRange(min, max int) int {
	if max < min {
		panic(fmt.Sprintf("randRange: min > max (%d, %d)", min, max))
	}

	r := (max + 1) - min
	ret := csRandInstance.Intn(r)
	return ret + min
}

/* vim :set ts=4 sw=4 sts=4 noet : */
