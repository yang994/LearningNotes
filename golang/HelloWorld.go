package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"

	"github.com/memoio/go-mefs/utils/metainfo"
)

type payInfo struct {
	pid       string
	uid       string
	spaceTime int64
	stStart   int64
	stEnd     int64
	sign      []byte
}

func main() {
	count := []int{0, 0, 0, 0, 0}
	now := int64(1587198745)
	if false {
		for i := 0; i < 1000000; i++ {
			var tsl sortlist
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

func test() {
	a:=time.Now()
	for i:=0;i<1000;i++{
		handleSTpayReply()
	}
	fmt.Println(time.Since(a))
}

const (
	TMAX = 3600
	TMIN = 1800
)

//Leader时空支付流程
//选主 -> 算时空值 -> 广播
func spaceTimepay() {
	stStart := queryStEnd("uid")
	if keeperIsLeader("uid", "kid", stStart) {
		spaceTime, chalList := resultSummary("uid", "pid", stStart, stStart+60*3600)
		strChalList := []string{}
		for _, cr := range chalList {
			strChalList = append(strChalList, cr.toString())
		}
		strings.Join(strChalList, "\n")
		metainfo.NewKeyMeta("uid", metainfo.Test, "stStart", "stEnd", strconv.Itoa(int(spaceTime)), "pid")
	}
}

//验证节点验证工作
func handleSpaceTimePay(km *metainfo.KeyMeta, metaValue string) {
	stStart := queryStEnd("uid") //查stStart
	if keeperIsLeader("uid", "kid", stStart) {
		strings.Split(metaValue, "\n")
		spaceTime, chalResult := resultSummary("uid", "pid", stStart, stStart+60*3600)
		for _, data := range chalResult {
			data.check()
		}
		msg:=strings.Join([]string{"stStart", "stEnd", strconv.Itoa(int(spaceTime)), "pid", "uid"}, "")
		GetSign([]byte(msg), "private.pem")
	}
}

func handleSTpayReply(){
	stStart := queryStEnd("uid")
	_, chalResult := resultSummary("uid", "pid", stStart, stStart+60*3600)
	for _, data := range chalResult {
		data.check()
	}
}