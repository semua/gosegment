package main

import (
	"fmt"
	"strconv"

	"github.com/semua/gosegment/segment"
	"github.com/semua/gosegment/segment/dict"
)

func main() {
	seg := segment.NewSegment()
	err := seg.Init("./dicts")
	if err != nil {
		fmt.Println("%v", err)
	}
	ret := seg.DoSegment(`盘古分词 简介: 盘古分词 是由eaglet 开发的一款基于字典的中英文分词组件
	主要功能: 中英文分词，未登录词识别,多元歧义自动识别,全角字符识别能力
	主要性能指标:
	分词准确度:90%以上
	处理速度: 300-600KBytes/s Core Duo 1.8GHz
	用于测试的句子:
	长春市长春节致词
	长春市长春药店
	IＢM的技术和服务都不错
	张三在一月份工作会议上说的确实在理
	于北京时间5月10日举行运动会
	我的和服务必在明天做好`)

	for cur := ret.Front(); cur != nil; cur = cur.Next() {
		w := cur.Value.(*dict.WordInfo)
		fmt.Print(w.Word, "(", w.Rank, ")/")
	}

	s := `7月5日至6日，中共中央政治局常委、国务院总理李克强连续到安徽阜阳、湖南岳阳、湖北武汉考察长江、淮河流域防汛抗洪和抢险救灾工作。
　　今年入汛早，大范围强降雨集中且持续时间长，保障大江大河大湖安全度汛、确保人民群众生命安全至关重要。党中央、国务院高度重视，习近平总书记等中央领导同志多次作出重要指示和批示，各地各有关部门扎实有效推进防汛抗洪工作。淮河历来是灾害多发区域。李克强来到“千里淮河第一闸”安徽阜阳王家坝闸，察看上游来水，了解汛情变化，并嘱咐工作人员当好“耳目尖兵”，确保监测预报精准及时。他还听取淮河水利委员会负责人汇报。李克强说，七八月份是防汛关键期，防汛抗洪的攻坚还在后面，要始终紧绷安全这根弦，上下游统筹协作，做好各种应急准备，把握抗洪抢险主动权。
　　在蒙洼蓄洪区的郑台子庄台，李克强向群众了解生产生活和抗洪的食品药品准备情况，询问庄台坚固性是否有保障。他对当地负责人说，你们处在淮河防汛的关键地带，一定要以人民群众生命财产安全为重，工作前移、保障到位，让群众对战胜洪水更有信心、挺起脊梁。党和政府会继续加大扶持力度，绝不让蓄洪洼地变成民生洼地。`
	ret = seg.DoSegment(s)
	var scoreWords []string
	for cur := ret.Front(); cur != nil; cur = cur.Next() {
		w := cur.Value.(*dict.WordInfo)
		if len(w.Word) > 3 {
			//单个汉字忽略
			scoreWords = append(scoreWords, w.Word)
		}
	}
	term := segment.TextRank(scoreWords, 20)
	fmt.Println()
	for _, n := range term {
		fmt.Println(n.Word + ":" + strconv.FormatFloat(n.Score, 'f', -1, 64))
	}
}
