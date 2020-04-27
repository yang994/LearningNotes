package main

import (
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
	"sort"
	"time"
)

var KEEPERS = []string{"8MHS9fZzRaHNj4mP1kYDebwySmLzaw", "8MGRZbvn8caS431icB2P1uT74B3EHh", "8MJCzFbpXCvdfzmJy5L8jiw4w1qPdY", "8MKX58Ko5vBeJUkfgpkig53jZzwqoW", "8MHYzNkm6dF9SWU5u7Py8MJ31vJrzS"}

func main() {
	count := []int{0, 0, 0, 0, 0}
	now := int64(1587198745)
	if true {
		for i := 0; i < 1000000; i++ {
			var tsl timesortlist
			now += RangeRand(30*60, 60*60)
			for j, kid := range KEEPERS {
				thisData := data{
					index: j,
					h:     h256(kid, now, 0)[:1],
				}
				tsl = append(tsl, thisData)
			}
			sort.Sort(tsl)
			count[tsl[0].index]++
		}
		fmt.Println(count)
	} else {
		test()
	}
}
func h256(kid string, stEnd int64, round int) []byte {
	h := sha256.New()
	in := fmt.Sprintf("%s_%d_%d", kid, stEnd, round)
	h.Write([]byte(in))
	return h.Sum(nil)
}

type data struct {
	index int
	h     []byte
}

type timesortlist []data //该结构用来对挑战结果按时间进行排序，以便计算时空值
func (p timesortlist) Swap(i, j int) {
	p[i].index, p[j].index = p[j].index, p[i].index
	for x, _ := range p[i].h {
		p[i].h[x], p[j].h[x] = p[j].h[x], p[i].h[x]
	}
}
func (p timesortlist) Len() int           { return len(p) }
func (p timesortlist) Less(i, j int) bool { return aBiggerThanB(p[i].h, p[j].h) }
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

func test() {
	fmt.Println(time.Now().Unix())
}

func RangeRand(min, max int64) int64 {
	if min > max {
		panic("the min is greater than max!")
	}

	if min < 0 {
		f64Min := math.Abs(float64(min))
		i64Min := int64(f64Min)
		result, _ := rand.Int(rand.Reader, big.NewInt(max+1+i64Min))

		return result.Int64() - i64Min
	} else {
		result, _ := rand.Int(rand.Reader, big.NewInt(max-min+1))
		return min + result.Int64()
	}
}
