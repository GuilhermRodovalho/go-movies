package main

import (
	"flag"
	"fmt"
)

type Options struct {
	isNew  bool
	isAll  *bool
	toEdit string
}

func main() {
	// userOptions := getFlags()
	uo := Options{}
	// isNew := flag.Bool("new", false, "user wants to input a new movie")
	uo.isAll = flag.Bool("all", false, "get all the movies loaded")
	// toEdit := flag.String("edit", "", "edit the movie with the name. Put it on quotes if is more than one name")
	flag.Parse()

	fmt.Printf("%T, %v", *uo.isAll, *uo.isAll)

	// handler := Movies.NewFileHandler()

	switch {
	// case userOptions.isNew:
	// 	inputNewMovie(handler)
	// case userOptions.isAll:
	// 	showAllMovies(handler)
	// case userOptions.toEdit != "":
	// 	editMovie(userOptions.toEdit, handler)
	}
}

func getFlags() Options {
	opt := Options{}
	opt.isNew = *flag.Bool("new", false, "user wants to input a new movie")
	// opt.isAll = *flag.Bool("all", false, "get all the movies loaded")
	opt.toEdit = *flag.String("edit", "", "edit the movie with the name. Put it on quotes if is more than one name")
	// test := flag.Bool("test", true, "testing")
	flag.Parse()

	return opt
}
