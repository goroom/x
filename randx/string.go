package randx

const (
	RST_LOWER  = 0x1
	RST_UPPER  = 0x2
	RST_NUMBER = 0x4
	RST_SYMBOL = 0x8
)

const (
	_lowerString  = "qwertyuiopasdfghjklzxcvbnm"
	_upperString  = "QWERTYUIOPASDFGHJKLZXCVBNM"
	_numberString = "1234567890"
	_symbolString = "`~!@#$%^&*()-_=+[{]}\\|;:'\",<.>/?"
)

var (
	mString = map[int]string{
		RST_LOWER:  _lowerString,
		RST_UPPER:  _upperString,
		RST_NUMBER: _numberString,
		RST_SYMBOL: _symbolString,

		RST_LOWER | RST_UPPER:   _lowerString + _upperString,
		RST_LOWER | RST_NUMBER:  _lowerString + _numberString,
		RST_LOWER | RST_SYMBOL:  _lowerString + _symbolString,
		RST_UPPER | RST_NUMBER:  _upperString + _numberString,
		RST_UPPER | RST_SYMBOL:  _upperString + _symbolString,
		RST_NUMBER | RST_SYMBOL: _numberString + _symbolString,

		RST_LOWER | RST_UPPER | RST_NUMBER:  _lowerString + _upperString + _numberString,
		RST_LOWER | RST_UPPER | RST_SYMBOL:  _lowerString + _upperString + _symbolString,
		RST_LOWER | RST_NUMBER | RST_SYMBOL: _lowerString + _numberString + _symbolString,
		RST_UPPER | RST_NUMBER | RST_SYMBOL: _upperString + _numberString + _symbolString,

		RST_LOWER | RST_UPPER | RST_NUMBER | RST_SYMBOL: _lowerString + _upperString + _numberString + _symbolString,
	}
)

func (r *Rand) String(length int, flag int) string {
	return r.StringLib(length, mString[flag])
}

func (r *Rand) StringLib(length int, str string) string {
	_s := ""
	for i := 0; i < length; i++ {
		_s += r.RangeString(str)
	}
	return _s
}

func (r *Rand) StringArray(list []string) string {
	return list[r.Intn(len(list))]
}

func (r *Rand) RangeString(s string) string {
	str := []rune(s)
	index := r.Intn(len(str))
	return string(str[index])
}
