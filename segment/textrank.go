package segment

import (
	"math"
	"reflect"
	"sort"
	"strings"
)

func TextRank(words []string, topN int) []*Term {
	var topWords []*Term
	wordGraph := make(map[string][]string)
	var walkQueue []string
	for _, w := range words {
		walkQueue = append(walkQueue, w)
		if len(walkQueue) > 5 {
			walkQueue = walkQueue[1:]
		}
		for _, q1 := range walkQueue {
			for _, q2 := range walkQueue {
				if strings.EqualFold(q1, q2) {
					continue
				}
				if _, ok := wordGraph[q1]; !ok {
					wordGraph[q1] = make([]string, 0)
				}
				if _, ok := wordGraph[q2]; !ok {
					wordGraph[q2] = make([]string, 0)
				}
				if !Contain(q2, wordGraph[q1]) {
					wordGraph[q1] = append(wordGraph[q1], q2)
				}
				if !Contain(q1, wordGraph[q2]) {
					wordGraph[q2] = append(wordGraph[q2], q1)
				}
			}
		}
	}
	scoreMap := make(map[string]float64)
	for i := 0; i < 200; {
		m := make(map[string]float64)
		var max_diff float64
		for k, v := range wordGraph {
			m[k] = 1 - 0.85
			for _, e := range v {
				size := len(wordGraph[e])
				if strings.EqualFold(k, e) || size == 0 {
					continue
				}
				if _, ok := scoreMap[e]; ok {
					m[k] = m[k] + 0.85*scoreMap[e]/float64(size)
				}
			}
			if _, ok := scoreMap[k]; ok {
				max_diff = math.Max(max_diff, math.Abs(m[k]-scoreMap[k]))
			} else {
				max_diff = math.Max(max_diff, math.Abs(m[k]))
			}
		}
		scoreMap = m
		if max_diff <= 0.001 {
			break
		}
		i++
	}
	for k, v := range scoreMap {
		topWords = append(topWords, &Term{k, v})
	}
	sort.Sort(TermWrapper{topWords, func(p, q *Term) bool {
		return p.Score > q.Score //递减排序
	}})
	if len(topWords) > topN {
		return topWords[:topN]
	} else {
		return topWords
	}
}

type Term struct {
	Word  string  // 分词
	Score float64 // 得分
}

type TermWrapper struct {
	term []*Term
	by   func(p, q *Term) bool
}

func (a TermWrapper) Len() int { // 重写 Len() 方法
	return len(a.term)
}
func (a TermWrapper) Swap(i, j int) { // 重写 Swap() 方法
	a.term[i], a.term[j] = a.term[j], a.term[i]
}
func (a TermWrapper) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a.by(a.term[i], a.term[j])
}
func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}
