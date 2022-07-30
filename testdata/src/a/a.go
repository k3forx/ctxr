package a

import "context"

func f(ctx context.Context, id int) string {
	return "OK"
}

func g(c context.Context, id int, email string) string { // want "variable name of `context.Context` is invalid"
	return "NG"
}

func h(id int, email string, ctx context.Context) string { // want "variable name of `context.Context` is invalid"
	return "NG" // want
}
