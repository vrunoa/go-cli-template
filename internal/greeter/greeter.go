package greeter

import "fmt"

type greeter struct {
	Lang string
}

func New(lang string) *greeter {
	return &greeter{
		Lang: lang,
	}
}

func (g *greeter) What() string {
	switch g.Lang {
	case "es":
		return "hola"
	case "en":
		return "hello"
	default:
		return "01001000 01100101 01101100"
	}
}

func (g *greeter) Hello() {
	fmt.Println(g.What())
}
