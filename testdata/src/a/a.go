package a

import "context"

func f(ctx context.Context, id int, email string) error {
	return nil
}

func g(c context.Context, name string) (string, error) { // want "1st args of func 'g' is context.Context, and its name should be 'ctx'"
	return "", nil
}

func h(id int, ctx context.Context) {} // want "2nd args of func 'h' is context.Context, and it should be first arg"
