package names

import (
	"math/rand"
	"strconv"
)

var nameList = []string{
	"귀여운",
	"예쁜",
	"가냘픈",
	"게으른",
	"고마운",
	"그리운",
	"기쁜",
	"깨끗한",
	"슬픈",
	"날랜",
	"너그러운",
	"아쉬운",
	"외로운",
	"편안한",
	"기분좋은",
	"지혜로운",
	"멋진",
	"한결같은",
	"반가운",
	"힘찬",
	"시원한",
	"추운",
	"더운",
	"느긋한",
	"사려깊은",
	"관대한",
	"친근한",
	"정직한",
	"솔직한",
	"순수한",
	"공손한",
	"정중한",
	"용감한",
	"근면한",
	"진지한",
	"성실한",
	"용맹한",
	"밝은",
	"긍정적인",
	"행복한",
	"즐거운",
	"명랑한",
	"유쾌한",
	"재미있는",
	"신나는",
}

func MakeNewName() string {
	idx := rand.Int() % len(nameList)
	randNum := strconv.FormatInt(int64(rand.Int()%49995+5), 10)

	return nameList[idx] + " 깨코 " + randNum + "호"
}
