package pizza

// makeHawaiPizza creates a hawai pizza.
// Its sole purpose is to demonstrate the options builder usage.
// But man, hawai? Seriously? From all the pizzas you chose hawai...
func makeHawaiPizza() (pizza *Pizza, err error) {
	return New(
		WithBase(Sugo),
		WithSize(Regular),
		WithTopping(Ham),
		WithTopping(Pineapple),
	)
}
