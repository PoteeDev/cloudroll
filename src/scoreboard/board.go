package scoreboard

import (
	"github.com/PoteeDev/cloudroll/src/scoreboard/models"
	"gorm.io/gorm/clause"
)

func (s *ScoreBoard) GetBoard() (*models.Board, error) {
	var team models.Team
	res := s.DB.Preload("Board.Columns.Cards.Tasks.Task").Preload(clause.Associations).First(&team)
	if res.Error != nil {
		return nil, res.Error
	}
	return &team.Board, nil
}

func (s *ScoreBoard) UpdateBoard(board *models.Board) error {
	var team models.Team
	s.DB.First(&team)
	team.Board = *board
	return s.DB.Save(&team).Error
}

func (s *ScoreBoard) GenerateBoard() {
	task := models.Task{Name: "Test Task", Description: "Test Task Description", Points: 100}

	s.DB.Create(&task)
	team := models.Team{Name: "testTeam"}
	s.DB.Create(&team)
	board := models.Board{
		TeamID: 1,
		Columns: []*models.Column{
			{Name: "TODO", Cards: []*models.Card{
				{Title: "Test Card", Tasks: []*models.TeamTask{
					{Completed: false, Task: &task},
				}},
			}},
			{Name: "In progress"},
			{Name: "Review"},
			{Name: "Done"},
		},
	}
	team.Board = board
	s.DB.Save(&team)
}
