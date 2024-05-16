package pizza

import "errors"

type Topping string

const (
	Prosciutto Topping = "prosciutto"
	Salami     Topping = "salami"
	Ham        Topping = "ham"
	Pineapple  Topping = "pineapple" // Don't even think about it
)

type Size string

const (
	Regular Size = "32cm"
	Large   Size = "40cm"
)

type Base string

const (
	Sugo  Base = "sugo"
	Cream Base = "cream"
)

type Pizza struct {
	Base     Base
	Size     Size
	Toppings []Topping
}

func New(options ...Option) (pizza *Pizza, err error) {
	pizza = &Pizza{}
	for _, option := range options {
		option(pizza)
	}

	if pizza.Size == "" {
		return nil, errors.New("size was not selected")
	}

	return pizza, nil
}
