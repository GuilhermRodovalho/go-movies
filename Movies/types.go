package Movies

type Movie struct {
	Name     string
	Year     int
	MyRating int
	Genre    []string
}

type Movies []Movie

type Handler interface {
	Read() (Movies, error)
	Write(mv *Movie) error
}
