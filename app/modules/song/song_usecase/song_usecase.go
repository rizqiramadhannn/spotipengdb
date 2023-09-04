package song_usecase

import (
	"github.com/labstack/echo/v4"

	"spotipeng/app/domain"
	"spotipeng/app/global"
)

type songUsecase struct {
}

func (u songUsecase) RegisterSong(ctx echo.Context, song domain.Song) (err error) {
	// Create a new song object
	newSong := domain.Song{
		Title:  song.Title,
		Album:  song.Album,
		Singer: song.Singer,
		URL:    song.URL,
		// Set other song properties as needed
	}

	// Store the new song in the database
	err = global.SongRepo.Post(ctx, newSong)
	if err != nil {
		// Handle the error (e.g., log it)
		return err
	}

	// Registration successful
	return nil
}

func (u songUsecase) Get(ctx echo.Context) (songs []domain.Song, err error) {
	return global.SongRepo.Get(ctx)
}

func (u songUsecase) GetById(ctx echo.Context, id int64) (song domain.Song, err error) {
	song, err = global.SongRepo.GetById(ctx, id)
	return
}

func (u songUsecase) Patch(ctx echo.Context, song domain.Song) (err error) {
	return global.SongRepo.Patch(ctx, song)
}

func (u songUsecase) Delete(ctx echo.Context, song domain.Song) (err error) {
	return global.SongRepo.Delete(ctx, song)
}

func New() domain.SongUsecase {
	return &songUsecase{}
}
