package pizza

import "slices"

type Option func(*Pizza)

func WithTopping(topping Topping) Option {
	return func(p *Pizza) {
		if slices.Contains(p.Toppings, topping) {
			return
		}

		p.Toppings = append(p.Toppings, topping)
	}
}

func WithSize(size Size) Option {
	return func(p *Pizza) {
		p.Size = size
	}
}

func WithBase(base Base) Option {
	return func(p *Pizza) {
		p.Base = base
	}
}
