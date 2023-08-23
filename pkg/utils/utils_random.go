// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

package utils

import (
	"math/rand"
)

type Range struct {
	Begin int
	End   int
}

func RandIntFromRange(r Range) int {
	if r.End-r.Begin <= 0 {
		return r.Begin
	}
	return rand.Intn((r.End-r.Begin)+1) + r.Begin
}

/*
 * amount:金额
 * num:剩余红包个数
 */
func RandomRedEnvelope(amount, num int) int {
	var (
		minAmount = 1
	)
	if amount == 0 || num == 0 {
		return 0
	}
	if num == 1 {
		return amount
	} else {
		return rand.Intn(amount/num) + minAmount
	}
}
