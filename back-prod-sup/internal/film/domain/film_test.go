package domain

import (
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

func TestNewFilm(t *testing.T) {
	name := "Avatar"
	description := "A sci-fi adventure"
	duration := uint16(180)
	releaseDate := time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC)

	film := NewFilm(name, description, duration, releaseDate)

	assert.NotNil(t, film)
	assert.Equal(t, name, film.Name)
	assert.Equal(t, description, film.Description)
	assert.Equal(t, duration, film.Duration)
	assert.Equal(t, releaseDate, film.ReleaseDate)
	assert.NotEqual(t, uuid.Nil, film.ID)
	assert.NotEqual(t, time.Time{}, film.CreatedAt)
	assert.NotEqual(t, time.Time{}, film.UpdatedAt)
	assert.True(t, film.CreatedAt.Equal(film.UpdatedAt))
}

func TestFilmValidation(t *testing.T) {
	validFilm := &Film{
		Name:        "Avatar",
		Description: "A sci-fi adventure",
		Duration:    180,
		ReleaseDate: time.Date(2009, time.January, 1, 0, 0, 0, 0, time.UTC),
	}

	tests := []struct {
		name    string
		film    *Film
		wantErr bool
	}{
		{"ValidFilm", validFilm, false},
		{"EmptyName", &Film{Name: ""}, true},
		{"LongName", &Film{Name: "This is a very long name that exceeds the maximum character limit for a film's name."}, true},
		{"EmptyDescription", &Film{Description: ""}, true},
		{"LongDescription", &Film{Description: "This is a very long description that exceeds the maximum character limit for a film's description."}, true},
		{"ZeroDuration", &Film{Duration: 0}, true},
		{"NegativeDuration", &Film{Duration: 0}, true},
		{"LongDuration", &Film{Duration: 60001}, true},
		{"FutureReleaseDate", &Film{ReleaseDate: time.Now().AddDate(0, 0, 1)}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := tt.film.Validate()
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestFilmEquality(t *testing.T) {
	film1 := &Film{ID: uuid.New(), Name: "Avatar"}
	film2 := &Film{ID: uuid.New(), Name: "Avatar"}

	assert.False(t, film1 == film2)
}
