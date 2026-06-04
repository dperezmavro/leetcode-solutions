package solution

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaster(t *testing.T) {
	tests := []struct {
		pass        string
		maxAttempts int
		words       []string
		want        int
	}{
		{
			pass:        "AAAAAA",
			words:       []string{"ABCDEF"},
			maxAttempts: 10,
			want:        1,
		},
		{
			pass:        "AAAAAA",
			words:       []string{"AAAAAA"},
			maxAttempts: 10,
			want:        6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.pass,
			func(t *testing.T) {
				m := Master{
					maxAttempts: tt.maxAttempts,
					words:       tt.words,
					pass:        tt.pass,
				}

				assert.Equal(t, tt.want, m.Guess(tt.words[0]))
			})
	}
}

func TestCheckDistances(t *testing.T) {
	var pass string = "hbaczn"
	var guesses []string = []string{
		"gaxckt",
		"trlccr",
		"jxwhkz",
		"ycbfps",
		"peayuf",
		"yiejjw",
		"ldzccp",
		"nqsjoa",
		"qrjasy",
		"pcldos",
		"acrtag",
		"buyeia",
		"ubmtpj",
		"drtclz",
		"zqderp",
		"snywek",
		"caoztp",
		"ibpghw",
		"evtkhl",
		"bhpfla",
		"ymqhxk",
		"qkvipb",
		"tvmued",
		"rvbass",
		"axeasm",
		"qolsjg",
		"roswcb",
		"vdjgxx",
		"bugbyv",
		"zipjpc",
		"tamszl",
		"osdifo",
		"dvxlxm",
		"iwmyfb",
		"wmnwhe",
		"hslnop",
		"nkrfwn",
		"puvgve",
		"rqsqpq",
		"jwoswl",
		"tittgf",
		"evqsqe",
		"aishiv",
		"pmwovj",
		"sorbte",
		"hbaczn",
		"coifed",
		"hrctvp",
		"vkytbw",
		"dizcxz",
		"arabol",
		"uywurk",
		"ppywdo",
		"resfls",
		"tmoliy",
		"etriev",
		"oanvlx",
		"wcsnzy",
		"loufkw",
		"onnwcy",
		"novblw",
		"mtxgwe",
		"rgrdbt",
		"ckolob",
		"kxnflb",
		"phonmg",
		"egcdab",
		"cykndr",
		"lkzobv",
		"ifwmwp",
		"jqmbib",
		"mypnvf",
		"lnrgnj",
		"clijwa",
		"kiioqr",
		"syzebr",
		"rqsmhg",
		"sczjmz",
		"hsdjfp",
		"mjcgvm",
		"ajotcx",
		"olgnfv",
		"mjyjxj",
		"wzgbmg",
		"lpcnbj",
		"yjjlwn",
		"blrogv",
		"bdplzs",
		"oxblph",
		"twejel",
		"rupapy",
		"euwrrz",
		"apiqzu",
		"ydcroj",
		"ldvzgq",
		"zailgu",
		"xgqpsr",
		"wxdyho",
		"alrplq",
		"brklfk",
	}

	m := &Master{
		maxAttempts: 10,
		words:       guesses,
		pass:        pass,
	}

	findSecretWord(guesses, m)

}
