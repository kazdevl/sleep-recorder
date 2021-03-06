package entity_test

import (
	"sleep-manager/entity"
	"testing"
)

func TestAdjustDuration(t *testing.T) {
	input := 9.85
	want := 9.9
	sr := entity.SleepRecord{Duration: input}
	sr.AdjustDuration()
	if sr.Duration != want {
		t.Errorf("error: result = %v, want = %v", sr.Duration, want)
	}
}

func TestConvertToResponse(t *testing.T) {
	tests := []struct {
		name  string
		input entity.Evaluation
		want  string
	}{
		{
			name:  "SuperBad",
			input: entity.SuperBad,
			want:  "๐ฑ ไผธใณไปฃใใใชใ!",
		},
		{
			name:  "Bad",
			input: entity.Bad,
			want:  "๐ฅ ใใใฐใ!",
		},
		{
			name:  "Good",
			input: entity.Good,
			want:  "๐ ่ฏใใญ!",
		},
		{
			name:  "Perfect",
			input: entity.Perfect,
			want:  "๐คฉ ๅฎ็ง!",
		},
		{
			name:  "Error",
			input: entity.Evaluation(100),
			want:  "๐คฉ ใจใฉใผ!",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := test.input.ConvertToResponse()
			if result != test.want {
				t.Errorf("error: result = %v, want = %v", result, test.want)
			}
		})
	}
}

func TestGetLineMessage(t *testing.T) {
	instance := entity.ResponseContents{
		Record: []entity.ResponseContent{
			{
				Date:     "13ๆฅ",
				TimeB:    "1ๆ1ๅ",
				TimeW:    "8ๆ21ๅ",
				Duration: "7.3ๆ้",
				Eval:     "๐ฅ ใใใฐใ!",
			},
			{
				Date:     "14ๆฅ",
				TimeB:    "1ๆ11ๅ",
				TimeW:    "7ๆ30ๅ",
				Duration: "6.3ๆ้",
				Eval:     "๐ฅ ใใใฐใ!",
			},
		},
		Avg: 6.8,
	}
	head := "****็ก็ ่จ้ฒ****\nๅนณๅ็ก็ ๆ้: 6.8ๆ้\n\nๅๆฅใซใกใฎ็ก็ ่จ้ฒ\n"
	bodyOne := "ใ13ๆฅใ: ๐ฅ ใใใฐใ!\n\tๅฐฑๅฏ: 1ๆ1ๅ\n\t่ตทๅบ: 8ๆ21ๅ\n\t็ก็ ๆ้: 7.3ๆ้\n\n"
	bodyTwo := "ใ14ๆฅใ: ๐ฅ ใใใฐใ!\n\tๅฐฑๅฏ: 1ๆ11ๅ\n\t่ตทๅบ: 7ๆ30ๅ\n\t็ก็ ๆ้: 6.3ๆ้\n\n"
	want := head + bodyOne + bodyTwo
	result := instance.GetLineMessage()
	if result != want {
		t.Errorf("unmatched error:\nresult is\n%s\nwant is\n%s", result, want)
	}
}
