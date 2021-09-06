package random

type Phone struct {
	Operators string
	Number    int
}

type operators struct {
	Name  string
	Heads []int
}

var dianxin *operators = &operators{"中国电信", []int{133, 149, 153, 173, 177, 180, 181, 189, 191, 199}}
var liantong *operators = &operators{"中国联通", []int{130, 131, 132, 145, 155, 156, 166, 171, 175, 176, 185, 186}}
var yidong *operators = &operators{"中国移动", []int{134, 135, 136, 137, 138, 139, 147, 150, 151, 152, 157, 158, 159, 172, 178, 182, 183, 184, 187, 188, 198}}

func (ext *Extend) Phone() *Phone {
	var op *operators
	switch ext.Intn(3) {
	case 0:
		op = dianxin
	case 1:
		op = liantong
	case 2:
		op = yidong
	}

	h := op.Heads[ext.Intn(len(op.Heads))] * 100000000
	return &Phone{
		Operators: op.Name,
		Number:    h + ext.Intn(100000000),
	}
}
