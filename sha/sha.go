//  Copyright (C) 晓白齐齐,版权所有.

package sha

import (
	"crypto/sha1"
	"fmt"
)

func CryptSSHA1(word []byte, salts string) []byte {
	s := []byte(salts)
	p := append(word, s...)
	h := sha1.Sum(p)
	return h[:]
}

func SHA(word string, salts ...string) string {
	bytes := []byte(word)
	for _, salt := range salts {
		bytes = CryptSSHA1(bytes, salt)
		//bytes = sqlite3.CryptEncoderSSHA1(salt)(bytes, nil)
		// loger.Debugf("salt is %s", salt)
	}
	return fmt.Sprintf("%x", bytes)
}
