package feature

import (
	"bytes"
	"crypto/md5"
	"fmt"
	"io"
)

// @author  wzz_714105382@icloud.com
// @date  2020/6/27 16:59
// @description
// @version

func LearnCrypto() {
	ex := "I am from Hebei Prov."
	exChanged := "I am from Hebei Prov." // 少了个句号
	hs := md5.New()
	_, _ = io.WriteString(hs, ex) // 这样写整洁一些,自动转[]byte
	bSend := hs.Sum(nil)

	hs.Reset() // 可以直接重置,接着用
	nhs := md5.New()
	_, _ = io.WriteString(hs, exChanged)
	_, _ = io.WriteString(nhs, exChanged)
	bReceived1 := hs.Sum(nil)
	bReceived2 := nhs.Sum(nil)
	fmt.Printf("md5 send:%x\n", bSend)
	fmt.Println("bReceived1 == bReceived2:", bytes.Equal(bReceived1, bReceived2))
	fmt.Println("data check : unchanged? ", bytes.Equal(bSend, bReceived2))
}
