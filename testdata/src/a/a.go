package a

import "context"

func f(ctx context.Context, id int, email string) error {
	return nil
}

func g(c context.Context, name string) (string, error) { // want "1st args of func 'g' is context.Context, and its name should be 'ctx'"
	return "", nil
}

func h(id int, ctx context.Context) {} // want "2nd args of func 'h' is context.Context, and it should be first arg"

func i(id1, id2, id3, id4, id5, id6, id7, id8 int, str1 string, ctx context.Context) // want "10th args of func 'i' is context.Context, and it should be first arg"

func j(id1, id2, id3, id4, id5, id6, id7, id8 int, str1, str2 string, ctx context.Context) // want "11th args of func 'j' is context.Context, and it should be first arg"
