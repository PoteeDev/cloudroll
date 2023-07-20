package scoreboard

import (
	"errors"
	"strings"

	"github.com/PoteeDev/cloudroll/src/scoreboard/models"
)

var (
	ErrDublicate = "duplicate key value violates unique constraint"
)

func (s *ScoreBoard) CreateTeam(name string) (*models.Team, error) {
	t := models.Team{Name: name}
	res := s.DB.Create(&t)
	if res.Error != nil {
		switch {
		case strings.Contains(res.Error.Error(), ErrDublicate):
			return nil, errors.New("team already exists")
		default:
			return nil, res.Error
		}

	}
	return &t, nil
}

func (s *ScoreBoard) JoinTeam(teamID string, userID string) (*models.Team, error) {
	var t models.Team
	if err := s.DB.First(&t, teamID).Error; err != nil {
		return nil, err
	}
	t.Users = append(t.Users, models.User{UserID: userID})
	if err := s.DB.Save(&t).Error; err != nil {
		if err != nil {
			switch {
			case strings.Contains(err.Error(), ErrDublicate):
				return nil, errors.New("user already joined team")
			default:
				return nil, err
			}

		}
		return nil, err
	}
	return &t, nil
}
