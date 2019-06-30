package vdf

import (
	"fmt"
	"math/big"
	"testing"
)

func TestCodec(t *testing.T) {
	vdf := NewVDF(3201)
	fmt.Println("初始化编解码器")

	sourceData := big.NewInt(334324)
	key := big.NewInt(2433141)
	round := 5
	en := vdf.encode(sourceData, key, round)
	fmt.Println("")
	fmt.Println("原数据：", sourceData)
	fmt.Println("key:", key)
	fmt.Println("加密次数：", round)
	fmt.Println("加密后数据：", en)

	de := vdf.decode(en, key, round)
	fmt.Println("解密后数据", de)

}
