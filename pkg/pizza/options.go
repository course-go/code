package pizza

type Option func(*Pizza)

func WithTopping(topping Topping) Option {
	return func(p *Pizza) {
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
