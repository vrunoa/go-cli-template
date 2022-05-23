package greeter

import "fmt"

type Greeter struct {
	Lang string
}

func New(lang string) *Greeter {
	return &Greeter{
		Lang: lang,
	}
}

func (g *Greeter) Hello() {
	switch g.Lang {
	case "es":
		fmt.Println("hola")
	case "en":
		fmt.Println("hello")
	default:
		fmt.Println("01001000 01100101 01101100")
	}
}
