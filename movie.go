package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/GuilhermRodovalho/movies/Movies"

	"github.com/jedib0t/go-pretty/table"
)

func inputNewMovie(handler Movies.Handler) {
	// handler := Movies.NewFileHandler()
	mv := getMovieFromUser()

	handler.Write(mv)
}

func getMovieFromUser() *Movies.Movie {
	mv := Movies.Movie{}
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Inputing a new movie")
	fmt.Print("Name: ")
	tmp, _ := reader.ReadString('\n')
	mv.Name = strings.Trim(tmp, "\n")

	fmt.Print("Year: ")
	fmt.Scan(&mv.Year)
	fmt.Print("My rating: ")
	fmt.Scan(&mv.MyRating)
	for {
		fmt.Print("Genre (enter when finished): ")
		var genre string
		tmp, _ := reader.ReadString('\n')
		genre = strings.Trim(tmp, "\n")
		if genre == "" {
			break
		}
		mv.Genre = append(mv.Genre, genre)
	}

	return &mv
}

func showAllMovies(handler Movies.Handler) {
	mvs, err := handler.Read()
	if err != nil {
		log.Println("Error has occured reading the movies ", err)
	}

	t := table.NewWriter()
	t.SetOutputMirror(os.Stdout)
	t.AppendHeader(table.Row{"Name", "Year", "Rating"})
	var rows []table.Row
	for _, mv := range mvs {
		r := table.Row{mv.Name, mv.Year, fmt.Sprintf("%d/10", mv.MyRating)}
		rows = append(rows, r)
	}
	t.AppendRows(rows)
	t.Render()
}

func editMovie(toEdit string, handler Movies.Handler) {
	mvs, err := handler.Read()
	if err != nil {
		fmt.Println("Problem reading movies ", err)
	}
	var desiredMovie Movies.Movie
	found := -1
	for i, mv := range mvs {
		if mv.Name == toEdit {
			desiredMovie = mv
			found = i
		}
	}

	if found == -1 {
		fmt.Println("Cannot find this movie, are you sure you added it to the list?")
		return
	}

	showUserEditOptions()
	choice := getUserChoice()

	getNewDataFromUser(&desiredMovie, choice)
}

type option rune

const (
	name   = 'n'
	year   = 'y'
	rating = 'r'
	genres = 'g'
)

func getUserChoice() option {
	var choice option
	fmt.Scanf("%c", &choice)

	return choice
}

func showUserEditOptions() {
	fmt.Println("Edit options")
	fmt.Println("(n) Edit name")
	fmt.Println("(y) Edit year")
	fmt.Println("(r) Edit rating")
	fmt.Println("(g) Edit genres")
	fmt.Print("What do you want to do? ")
}

func getNewDataFromUser(desiredMovie *Movies.Movie, choice option) *Movies.Movie {
	switch choice {
	case name:
		return getNewName(desiredMovie)
	case year:
		return getNewYear(desiredMovie)
	case rating:
		return getNewRating(desiredMovie)
	case genres:
		return getNewGenres(desiredMovie)
	}

	return desiredMovie
}

func getNewName(mv *Movies.Movie) *Movies.Movie {
	scn := bufio.NewReader(os.Stdin)
	scn.ReadLine()
	fmt.Print("Type the new name: ")
	name, err := scn.ReadString('\n')
	if err != nil {
		log.Panic("Problem reading name: ", err)
	}
	mv.Name = name
	return mv
}

func getNewYear(mv *Movies.Movie) *Movies.Movie {
	fmt.Print("Type the new year: ")
	fmt.Scan(&mv.Year)

	return mv
}

func getNewRating(mv *Movies.Movie) *Movies.Movie {
	fmt.Print("Type the new rating: ")
	fmt.Scan(&mv.MyRating)

	return mv
}

func getNewGenres(mv *Movies.Movie) *Movies.Movie {
	genres := make([]string, 1)

	reader := bufio.NewReader(os.Stdin)
	for {
		tmp, _ := reader.ReadString('\n')
		genre := strings.Trim(tmp, "\n")
		if genre == "" {
			break
		}
		genres = append(genres, genre)
	}
	mv.Genre = genres

	return mv
}
