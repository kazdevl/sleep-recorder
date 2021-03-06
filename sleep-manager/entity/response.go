package entity

import (
	"fmt"
	"strings"
)

type Evaluation int

func (e Evaluation) ConvertToResponse() string {
	switch e {
	case SuperBad:
		return "ð± ä¼¸ã³ä»£ãããªã!"
	case Bad:
		return "ð¥ ããã°ã!"
	case Good:
		return "ð è¯ãã­!"
	case Perfect:
		return "ð¤© å®ç§!"
	default:
		return "ð¤© ã¨ã©ã¼!"
	}
}

const (
	SuperBad Evaluation = iota
	Bad
	Good
	Perfect
)

type ResponseContent struct {
	Date     string `json:"date"`
	TimeB    string `json:"time_bedin"`
	TimeW    string `json:"time_wake"`
	Duration string `json:"duration"`
	Eval     string `json:"evaluation"`
}

type ResponseContents struct {
	Record []ResponseContent `json:"record"`
	Avg    float64           `json:"average"`
}

func (r ResponseContents) GetLineMessage() string {
	var head string = fmt.Sprintf("****ç¡ç è¨é²****\nå¹³åç¡ç æé: %gæé\n\nåæ¥ã«ã¡ã®ç¡ç è¨é²\n", r.Avg)
	strList := make([]string, len(r.Record))
	for i, record := range r.Record {
		str := fmt.Sprintf("ã%sã: %s\n\tå°±å¯: %s\n\tèµ·åº: %s\n\tç¡ç æé: %s\n\n",
			record.Date, record.Eval,
			record.TimeB,
			record.TimeW,
			record.Duration,
		)
		strList[i] = str
	}
	body := strings.Join(strList, "")
	return head + body
}
