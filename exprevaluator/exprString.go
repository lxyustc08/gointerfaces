package exprevaluator

import "fmt"

func (v Var) String() string {
	return fmt.Sprintf("%s", string(v))
}

func (l literal) String() string {
	return fmt.Sprintf("%.6g", l)
}

func (u unary) String() string {
	return fmt.Sprintf("%c %s", u.op, u.x)
}

func (c call) String() string {
	s := fmt.Sprintf("%s(", c.fn)
	for index, arg := range c.args {
		s += fmt.Sprintf("%s", arg)
		if index < len(c.args)-1 {
			s += fmt.Sprintf(", ")
		}
	}
	s += fmt.Sprintf(")")
	return s
}
func (b binary) String() string {
	return fmt.Sprintf("%s %c %s", b.x, b.op, b.y)
}
