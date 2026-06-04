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
	tests := []struct {
		pass        string
		maxAttempts int
		words       []string
	}{
		{
			pass:        "hbaczn",
			maxAttempts: 10,
			words: []string{
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
			},
		},
		{
			pass:        "acckzz",
			maxAttempts: 10,
			words:       []string{"acckzz", "ccbazz", "eiowzz", "abcczz"},
		},
		{
			pass:        "hamada",
			maxAttempts: 10,
			words:       []string{"hamada", "khaled"},
		},
		{
			pass:        "ccoyyo",
			maxAttempts: 10,
			words: []string{
				"ccoyyo",
				"wichbx",
				"oahwep",
				"tpulot",
				"eqznzs",
				"vvmplb",
				"eywinm",
				"dqefpt",
				"kmjmxr",
				"ihkovg",
				"trbzyb",
				"xqulhc",
				"bcsbfw",
				"rwzslk",
				"abpjhw",
				"mpubps",
				"viyzbc",
				"kodlta",
				"ckfzjh",
				"phuepp",
				"rokoro",
				"nxcwmo",
				"awvqlr",
				"uooeon",
				"hhfuzz",
				"sajxgr",
				"oxgaix",
				"fnugyu",
				"lkxwru",
				"mhtrvb",
				"xxonmg",
				"tqxlbr",
				"euxtzg",
				"tjwvad",
				"uslult",
				"rtjosi",
				"hsygda",
				"vyuica",
				"mbnagm",
				"uinqur",
				"pikenp",
				"szgupv",
				"qpxmsw",
				"vunxdn",
				"jahhfn",
				"kmbeok",
				"biywow",
				"yvgwho",
				"hwzodo",
				"loffxk",
				"xavzqd",
				"vwzpfe",
				"uairjw",
				"itufkt",
				"kaklud",
				"jjinfa",
				"kqbttl",
				"zocgux",
				"ucwjig",
				"meesxb",
				"uysfyc",
				"kdfvtw",
				"vizxrv",
				"rpbdjh",
				"wynohw",
				"lhqxvx",
				"kaadty",
				"dxxwut",
				"vjtskm",
				"yrdswc",
				"byzjxm",
				"jeomdc",
				"saevda",
				"himevi",
				"ydltnu",
				"wrrpoc",
				"khuopg",
				"ooxarg",
				"vcvfry",
				"thaawc",
				"bssybb",
				"ajcwbj",
				"arwfnl",
				"nafmtm",
				"xoaumd",
				"vbejda",
				"kaefne",
				"swcrkh",
				"reeyhj",
				"vmcwaf",
				"chxitv",
				"qkwjna",
				"vklpkp",
				"xfnayl",
				"ktgmfn",
				"xrmzzm",
				"fgtuki",
				"zcffuv",
				"srxuus",
				"pydgmq",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.pass,
			func(t *testing.T) {
				m := &Master{
					maxAttempts: tt.maxAttempts,
					words:       tt.words,
					pass:        tt.pass,
				}
				p := &PasswordFinder{
					words:        tt.words,
					master:       m,
					checkedWords: make(map[string]bool),
					distances:    make(map[string]int),
				}
				res := p.findSecretWordReturns()
				if res != tt.pass {
					t.Errorf("wrong pass: wanted %s got %s", tt.pass, res)
				}

				if m.k > m.maxAttempts {
					t.Errorf("max attempts exceeded: %d", m.k)
				}
			})
	}
}
