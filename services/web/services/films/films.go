package films

import (
	"time"

	pb "github.com/bredr/go-grpc-example/proto/films"
	"github.com/bredr/go-grpc-example/services/web/graph/model"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var mapModelToServiceGenre = map[model.Genre]pb.FilmGenre{
	model.GenreAction:      pb.FilmGenre_ACTION,
	model.GenreAdventure:   pb.FilmGenre_ADVENTURE,
	model.GenreAnimation:   pb.FilmGenre_ANIMATION,
	model.GenreBiography:   pb.FilmGenre_BIOGRAPHY,
	model.GenreComedy:      pb.FilmGenre_COMEDY,
	model.GenreCrime:       pb.FilmGenre_CRIME,
	model.GenreDocumentary: pb.FilmGenre_DOCUMENTARY,
	model.GenreDrama:       pb.FilmGenre_DRAMA,
	model.GenreFantasy:     pb.FilmGenre_FANTASY,
	model.GenreHorror:      pb.FilmGenre_HORROR,
	model.GenreRomance:     pb.FilmGenre_ROMANCE,
	model.GenreSciFi:       pb.FilmGenre_SCI_FI,
	model.GenreThriller:    pb.FilmGenre_THRILLER,
	model.GenreUnknown:     pb.FilmGenre_NULL_GENRE,
}
var mapServiceToModelGenre = map[pb.FilmGenre]model.Genre{
	pb.FilmGenre_ACTION:      model.GenreAction,
	pb.FilmGenre_ADVENTURE:   model.GenreAdventure,
	pb.FilmGenre_ANIMATION:   model.GenreAnimation,
	pb.FilmGenre_BIOGRAPHY:   model.GenreBiography,
	pb.FilmGenre_COMEDY:      model.GenreComedy,
	pb.FilmGenre_CRIME:       model.GenreCrime,
	pb.FilmGenre_DOCUMENTARY: model.GenreDocumentary,
	pb.FilmGenre_DRAMA:       model.GenreDrama,
	pb.FilmGenre_FANTASY:     model.GenreFantasy,
	pb.FilmGenre_HORROR:      model.GenreHorror,
	pb.FilmGenre_ROMANCE:     model.GenreRomance,
	pb.FilmGenre_SCI_FI:      model.GenreSciFi,
	pb.FilmGenre_THRILLER:    model.GenreThriller,
	pb.FilmGenre_NULL_GENRE:  model.GenreUnknown,
}

func MapFilmInput(input model.FilmInput) *pb.Film {
	f := &pb.Film{}
	f.Name = input.Name
	f.ReleaseDate = timestamppb.New(input.ReleaseDate)
	f.Genre = mapModelToServiceGenre[input.Genre]
	return f
}

func mapGenres(input []model.Genre) (out []pb.FilmGenre) {
	for _, x := range input {
		v, ok := mapModelToServiceGenre[x]
		if ok {
			out = append(out, v)
		}
	}
	return out
}

func MapFilm(x *pb.Film) model.Film {
	if x != nil {
		releaseDate := time.Time{}
		if x.GetReleaseDate() != nil {
			releaseDate = x.GetReleaseDate().AsTime()
		}
		return model.Film{
			Name:        x.Name,
			ID:          x.ID,
			Genre:       mapServiceToModelGenre[x.Genre],
			ReleaseDate: releaseDate,
		}
	}
	return model.Film{}
}

func MapFilms(input []*pb.Film) (out []model.Film) {
	for _, x := range input {
		if x != nil {
			out = append(out, MapFilm(x))
		}
	}
	return out
}

func GenerateSearchRequest(input *model.FilmSearch) *pb.FilmSearchRequest {
	request := &pb.FilmSearchRequest{}
	if input != nil {
		request.AllowedGenres = mapGenres(input.Genres)
		if input.SearchTerm != nil {
			request.NameSearch = *input.SearchTerm
		}
		if input.ReleasedAfter != nil {
			request.ReleasedAfter = timestamppb.New(*input.ReleasedAfter)
		}
	}
	return request
}
