//This script will disable defiant by removing every instance of them
package defiant

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"github.com/xackery/eqcleanup/tools/item"
	"github.com/xackery/eqemuconfig"
)

var focus = "defiant"

func Clean(db *sqlx.DB, config *eqemuconfig.Config) (err error) {
	fmt.Println("Removing every defiant item drop, spawn, merchant entry, etc in game...")
	//Find all item IDS

	ids := []int64{
		17525,
		40598,
		50005,
		50006,
		50007,
		50008,
		50009,
		50010,
		50011,
		50012,
		50013,
		50014,
		50015,
		50016,
		50017,
		50018,
		50019,
		50020,
		50021,
		50022,
		50023,
		50024,
		50025,
		50026,
		50027,
		50028,
		50029,
		50030,
		50031,
		50032,
		50033,
		50034,
		50035,
		50036,
		50037,
		50038,
		50039,
		50040,
		50041,
		50042,
		50043,
		50044,
		50045,
		50046,
		50047,
		50048,
		50049,
		50050,
		50051,
		50052,
		50053,
		50054,
		50055,
		50056,
		50057,
		50058,
		50059,
		50060,
		50061,
		50062,
		50063,
		50064,
		50065,
		50066,
		50067,
		50068,
		50069,
		50070,
		50071,
		50072,
		50073,
		50074,
		50075,
		50076,
		50077,
		50078,
		50079,
		50080,
		50081,
		50082,
		50083,
		50084,
		50085,
		50086,
		50087,
		50088,
		50089,
		50090,
		50091,
		50092,
		50093,
		50094,
		50095,
		50096,
		50097,
		50098,
		50099,
		50100,
		50101,
		50102,
		50103,
		50104,
		50105,
		50106,
		50107,
		50108,
		50109,
		50110,
		50111,
		50112,
		50113,
		50114,
		50115,
		50116,
		50117,
		50118,
		50119,
		50120,
		50121,
		50122,
		50123,
		50124,
		50125,
		50126,
		50127,
		50128,
		50129,
		50130,
		50131,
		50132,
		50133,
		50134,
		50135,
		50136,
		50137,
		50138,
		50139,
		50140,
		50141,
		50142,
		50143,
		50144,
		50145,
		50146,
		50147,
		50148,
		50149,
		50150,
		50151,
		50152,
		50153,
		50154,
		50155,
		50156,
		50157,
		50158,
		50159,
		50160,
		50161,
		50162,
		50163,
		50164,
		50165,
		50166,
		50167,
		50168,
		50169,
		50170,
		50171,
		50172,
		50173,
		50174,
		50175,
		50176,
		50177,
		50178,
		50179,
		50180,
		50181,
		50182,
		50183,
		50184,
		50185,
		50186,
		50187,
		50188,
		50189,
		50190,
		50191,
		50192,
		50193,
		50194,
		50195,
		50196,
		50197,
		50198,
		50199,
		50200,
		50201,
		50202,
		50203,
		50204,
		50205,
		50206,
		50207,
		50208,
		50209,
		50210,
		50211,
		50212,
		50213,
		50214,
		50215,
		50216,
		50217,
		50218,
		50219,
		50220,
		50221,
		50222,
		50223,
		50224,
		50225,
		50226,
		50227,
		50228,
		50229,
		50230,
		50231,
		50232,
		50233,
		50234,
		50235,
		50236,
		50237,
		50238,
		50239,
		50240,
		50241,
		50242,
		50243,
		50244,
		50245,
		50246,

		50282,
		50283,
		50284,
		50285,
		50286,
		50287,
		50288,
		50289,
		50290,
		50291,
		50296,
		50299,
		50300,
		50301,
		50302,
		50303,
		50304,
		50305,
		50306,
		50307,
		50308,
		50309,
		50310,
		50311,
		50312,
		50313,
		50314,
		50315,
		50316,
		50317,
		50318,
		50319,
		50320,
		50321,
		50322,
		50323,
		50324,
		50325,
		50326,
		50327,
		50328,
		50329,
		50330,
		50331,
		50332,
		50333,
		50334,
		50335,
		50336,
		50337,
		50338,
		50339,
		50340,
		50341,
		50342,
		50343,
		50344,
		50345,
		50346,
		50347,
		50348,
		50349,
		50350,
		50351,
		50352,
		50353,
		50354,
		50355,
		50356,
		50357,
		50358,
		50359,
		50360,
		50361,
		50362,
		50363,
		50364,
		50365,
		50366,
		50367,
		50368,
		50369,
		50370,
		50371,
		50372,
		50373,
		50374,
		50375,
		50376,
		50377,
		50378,
		50379,
		50380,
		50381,
		50382,
		50383,
		50384,
		50385,
		50386,
		50387,
		50388,
		50389,
		50390,
		50391,
		50392,
		50393,
		50394,
		50395,
		50396,
		50397,
		50398,
		50399,
		50401,
		50402,
		50403,
		50404,
		50405,
		50406,
		50407,
		50408,
		50409,
		50410,
		50411,
		50412,
		50413,
		50414,
		50415,
		50416,
		50417,
		50418,
		50419,
		50420,
		50421,
		50422,
		50423,
		50424,
		50425,
		50426,
		50427,
		50428,
		50429,
		50430,
		50431,
		50432,
		50433,
		50434,
		50435,
		50436,
		50437,
		50438,
		50439,
		50440,
		50441,
		50442,
		50443,
		50444,
		50445,
		50446,
		50447,
		50448,
		50449,
		50450,
		50451,
		50452,
		50453,
		50454,
		50455,
		50456,
		50457,
		50458,
		50459,
		50460,
		50461,
		50462,
		50463,
		50468,
		50473,
		50485,
		50486,
		50487,
		50488,
		50489,
		50490,
		50491,
		50492,
		50493,
		50494,
		50495,
		50496,
		50497,
		50498,
		50499,
		50500,
		50501,
		50502,
		50503,
		50504,
		50505,
		50506,
		50507,
		50508,
		50509,
		50510,
		50511,
		50512,
		50513,
		50514,
		50515,
		50516,
		50517,
		50518,
		50519,
		50520,
		50521,
		50522,
		50523,
		50524,
		50525,
		50526,
		50527,
		50528,
		50529,
		50530,
		50531,
		50532,
		50533,
		50534,
		50535,
		50536,
		50537,
		50540,
		50541,
		50542,
		50543,
		50544,
		50545,
		50546,
		50547,
		50548,
		50549,
		50550,
		50551,
		50552,
		50553,
		50554,
		50555,
		50556,
		50557,
		50558,
		50559,
		50560,
		50561,
		50562,
		50563,
		50564,
		50565,
		50566,
		50567,
		50568,
		50569,
		50570,
		50571,
		50572,
		50573,
		50574,
		50575,
		50576,
		50577,
		50578,
		50579,
		50580,
		50581,
		50582,
		50583,
		50584,
		50585,
		50586,
		50587,
		50588,
		50589,
		50590,
		50591,
		50592,
		50593,
		50594,
		50595,
		50596,
		50597,
		50598,
		50599,
		50600,
		50601,
		50602,
		50603,
		50604,
		50605,
		50606,
		50607,
		50608,
		50609,
		50610,
		50611,
		50612,
		50613,
		50614,
		50615,
		50616,
		50617,
		50618,
		50619,
		50620,
		50621,
		50622,
		50623,
		50624,
		50625,
		50626,
		50627,
		50628,
		50629,
		50630,
		50631,

		60500,
		60501,
		60502,
		60503,
		60504,
		60505,
		60506,
		60507,
		60508,
		60509,
		60510,
		60511,
		60512,
		60513,
		60514,
		60515,
		60516,
		60517,
		60518,
		60519,
		60520,
		60521,
		60522,
		60523,
		60524,
		60525,
		60526,
		60527,
		60548,
		60549,
		60550,
		60551,
		60552,
		60553,
		60554,
		60555,
		60556,
		60557,
		60558,
		60559,
		60560,
		60561,
	}

	totalChanged, err := item.RemoveAllInstancesOfItems(db, ids)
	if err != nil {
		return
	}
	fmt.Println("Removed", totalChanged, " DB entries related to", focus, "in all player-accessible item locations.")
	return
}
