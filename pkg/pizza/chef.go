package pizza

// makeHawaiPizza creates a Hawaiian pizza.
// Its sole purpose is to demonstrate the options builder usage.
// But man, Hawaiian? Seriously? Of all the pizzas you chose Hawaiian...
func makeHawaiPizza() (pizza *Pizza, err error) {
	return New(
		WithBase(Sugo),
		WithSize(Regular),
		WithTopping(Ham),
		WithTopping(Pineapple),
	)
}
