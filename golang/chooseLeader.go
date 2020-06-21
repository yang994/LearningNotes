package main

import (
	"crypto/sha256"
	"fmt"
	"sort"
)

var KEEPERS = []string{"8MHS9fZzRaHNj4mP1kYDebwySmLzaw", "8MGRZbvn8caS431icB2P1uT74B3EHh", "8MJCzFbpXCvdfzmJy5L8jiw4w1qPdY", "8MKX58Ko5vBeJUkfgpkig53jZzwqoW", "8MHYzNkm6dF9SWU5u7Py8MJ31vJrzS"}

type data struct {
	index int
	h     []byte
}
type sortlist []data 
func (p sortlist) Swap(i, j int) {
	p[i].index, p[j].index = p[j].index, p[i].index
	for x, _ := range p[i].h {
		p[i].h[x], p[j].h[x] = p[j].h[x], p[i].h[x]
	}
}
func (p sortlist) Len() int           { return len(p) }
func (p sortlist) Less(i, j int) bool { return aBiggerThanB(p[i].h, p[j].h) }
func aBiggerThanB(a, b []byte) bool {
	if len(a) != len(b) {
		fmt.Println("len err")
		return len(a) < len(b)
	}
	for i, value := range a {
		if value < b[i] {
			return false
		}
	}
	return true
}

//哈希函数
func h256(kid string, stEnd int64, round int) []byte {
	h := sha256.New()
	in := fmt.Sprintf("%s_%d_%d", kid, stEnd, round)
	h.Write([]byte(in))
	return h.Sum(nil)
}

//选主函数
func keeperIsLeader(uid string, kid string, stStart int64) bool {
	var tsl sortlist
	for j, kid := range KEEPERS {
		thisData := data{
			index: j,
			h:     h256(kid, stStart, 0)[:1],
		}
		tsl = append(tsl, thisData)
	}
	sort.Sort(tsl)
	return true
}
