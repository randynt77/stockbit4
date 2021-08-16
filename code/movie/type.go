package movie

type (
	Movie struct {
		Title  string `json:"Title"`
		Year   string `json:"Year"`
		ImdbID string `json:"imdbID"`
		Type   string `json:"Type"`
		Poster string `json:"Poster"`
	}
	GetMovieResponse struct {
		MovieData []Movie `json:"movie_data,omitempty"`
	}
	OmdbResponse struct {
		Search []Movie `json:"Search"`
	}
	SearchData struct {
		Keyword string
		Page    int
	}
)

type (
	Repository interface {
		InsertMovieData(movieDetail Movie) (err error)
		GetMovieByExactName(MovieName string) (movieData Movie, err error)
		IsAlreadyExist(ImdbID string) (exist bool, err error)

		LogSearchData(searchData SearchData) (err error)
	}
	RestRepository interface {
		GetMovieData(requestParam SearchData) (movieList []Movie, err error)
	}
	Usecase interface {
		GetMovieData(searchData SearchData) (Movies []Movie, err error)
	}
)
