package randx

var (
	_phonePrefixList = []string{"130", "131", "132", "133", "134", "135", "136", "137", "138",
		"139", "147", "150", "151", "152", "153", "155", "156", "157", "158", "159", "186",
		"187", "188"}
	_phonePrefixListLength = len(_phonePrefixList)
)

func (r *Rand) Phone() string {
	return _phonePrefixList[r.Intn(_phonePrefixListLength)] + r.String(8, RST_NUMBER)
}
