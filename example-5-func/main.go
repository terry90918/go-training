package main

import (
	"fmt"
	"strings"
)

type searchOpts struct {
	username string
	email    string
	sexy     int
}

func getUserListOptsSQL(opts searchOpts) string {
	sql := "select * from user"
	where := []string{}
	if opts.username != "" {
		where = append(where, fmt.Sprintf("username = '%s'", opts.username))
	}
	if opts.email != "" {
		where = append(where, fmt.Sprintf("email = '%s'", opts.email))
	}
	if opts.sexy != 0 {
		where = append(where, fmt.Sprintf("sexy = '%d'", opts.sexy))
	}
	return sql + " where " + strings.Join(where, " or ")
}

func main() {
	fmt.Println(getUserListOptsSQL(searchOpts{
		username: "terry",
		email:    "test@mail.com",
		sexy:     1,
	}))
}
