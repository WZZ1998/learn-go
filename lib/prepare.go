package lib

import (
	"math/rand"
	"time"
)

// @author  wzz_714105382@icloud.com
// @date  2020/10/30 18:13
// @description
// @version
var _pRand = rand.New(rand.NewSource(time.Now().UnixNano()))
