package progress

import (
	"fight-landlords/util"
)

type PlayCards struct {
	sequence string
	typ      any
}

func loop() {
	s1 := "56789X"
	procedure(s1)
}

func procedure(seq string) {
	seq = AscendStr(seq)
	SwitchToType(seq)
}

func AscendStr(s string) string {
	return util.SortedStr(s)
}

// seq已经升序排列
func SwitchToType(seq string) any {
	if seq != AscendStr(seq) {
		panic("seq is not ascend")
	}
	switch len(seq) {
	case 1:
		return Handle1(seq)
	case 2:
		return Handle2(seq)
	case 3:
		return Handle3(seq)
	case 4:
		return Handle4(seq)
	case 5:
		return Handle5(seq)
	case 6:
		return Handle6(seq)
	case 8:
		return Handle8(seq)
	default:
		return Handle0(seq)
	}
}
