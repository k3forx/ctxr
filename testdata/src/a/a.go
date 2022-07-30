package a

import "context"

func f(ctx context.Context, id int) string {
	return "OK"
}

func g(c context.Context, id int, email string) string {
	return "NG"
}

func h(id int, email string, ctx context.Context) string {
	return "NG"
}
