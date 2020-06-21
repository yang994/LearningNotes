package main

import (
	"crypto/rand"
	"fmt"
	"log"
	"math"
	"math/big"
	"strconv"
	"strings"
	"sort"
)

type chalresult struct {
	Kid           string
	Pid           string
	Uid           string
	ChallengeTime int64
	Length        int64
	Proof         []byte
}

const CHALDELIMITER = "\t"

func (this *chalresult) toString() string {
	res := strings.Join([]string{this.Kid, this.Pid, this.Uid, strconv.Itoa(int(this.ChallengeTime)), strconv.Itoa(int(this.Length)), string(this.Proof)}, CHALDELIMITER)
	return res
}

func (this *chalresult) sign() {
	str := strings.Join([]string{this.Kid, this.Pid, this.Uid, strconv.Itoa(int(this.ChallengeTime)), strconv.Itoa(int(this.Length))}, "")
	s := GetSign([]byte(str), "private.pem")
	this.Proof = s
}

func (this *chalresult) check() bool {
	str := strings.Join([]string{this.Kid, this.Pid, this.Uid, strconv.Itoa(int(this.ChallengeTime)), strconv.Itoa(int(this.Length))}, "")
	return VerifySign([]byte(str), this.Proof, "public.pem")
}
//从字符串获取挑战信息结构体
func getChalresult(str string) *chalresult {
	splitedKey := strings.Split(str, CHALDELIMITER)
	ct, err := strconv.Atoi(splitedKey[3])
	if err != nil {
		fmt.Println("splitedKey[3]=", splitedKey[3])
		return nil
	}
	l, err := strconv.Atoi(splitedKey[4])
	if err != nil {
		fmt.Println("splitedKey[4]=", splitedKey[4])
		return nil
	}
	cr := &chalresult{
		Kid:           splitedKey[0],
		Pid:           splitedKey[1],
		Uid:           splitedKey[2],
		ChallengeTime: int64(ct),
		Length:        int64(l),
		Proof:         []byte(splitedKey[5]),
	}
	return cr
}

type crList []chalresult

func (p crList) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p crList) Len() int           { return len(p) }
func (p crList) Less(i, j int) bool { return p[i].ChallengeTime < p[j].ChallengeTime }

var ChalResult crList

//取挑战结果，挑战结果按照时间排序
func fetchChalresult(uid, pid string, timeStart, timeEnd int64) crList {
	if len(ChalResult) != 0 {
		return ChalResult
	}
	stStart := queryStEnd("uid")
	var chalList crList
	for i := 0; i < 12; i++ {
		temp := &chalresult{
			Kid:           "kid",
			Pid:           "pid",
			Uid:           "uid",
			ChallengeTime: stStart,
		}
		length := RangeRand(5*1024*1024, 10*1024*1024)
		temp.Length = length
		temp.sign()
		chalList = append(chalList, *temp)
		stStart += RangeRand(4*60, 6*60)
	}
	ChalResult = chalList
	return chalList
}



//在要求范围内获得随机数
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

//算时空值，返回算好的时空值和排好序的挑战信息序列
func resultSummary(uid string, pid string, timeStart int64, timeEnd int64) (int64, []chalresult) {
	cl := fetchChalresult(uid, pid, timeStart, timeEnd) //取数据
	chalList := cl[:]
	sort.Sort(chalList)
	spacetime := int64(0)
	if len(chalList) <= 1 {
		fmt.Println("no enough challenge data")
		return 0, nil
	}
	timepre := chalList[0].ChallengeTime
	lengthpre := chalList[0].Length
	//初始化变量
	for _, data := range chalList[1:] { //循环数组进行计算
		length := data.Length
		timeafter := data.ChallengeTime
		spacetime += (timeafter - timepre) * int64(lengthpre+length) / 2
		timepre = data.ChallengeTime
		lengthpre = length
	}
	if spacetime < 0 {
		log.Println("error spacetime<0!")
	}
	return spacetime, chalList
}

//从合约中查询时间
func queryStEnd(account string) int64 {
	return int64(1587198745)
}
