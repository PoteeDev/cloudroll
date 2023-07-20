package scoreboard

import (
	"os"

	"github.com/PoteeDev/cloudroll/src/scoreboard/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type ScoreBoard struct {
	DB *gorm.DB
}

func Init() (*ScoreBoard, error) {
	db, err := gorm.Open(postgres.Open(os.Getenv("DATABASE_URL")+"&application_name=$ docs_simplecrud_gorm"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	err = db.AutoMigrate(&models.Team{}, &models.Score{}, &models.Task{}, &models.User{}, &models.Label{}, &models.TeamTask{}, &models.Card{}, &models.Column{}, &models.Board{})
	if err != nil {
		return nil, err
	}

	s := &ScoreBoard{DB: db}
	return s, nil
}

// Get Full scoreboard
func (s *ScoreBoard) GetScoreboard() []models.Score {
	var scoreboard []models.Score
	s.DB.Find(&scoreboard)
	return scoreboard
}

// Update scoreboard with scoreID and new tasks
func (s *ScoreBoard) UpdateScoreboard(scoreID int, task *models.Task) error {
	var score models.Score
	s.DB.First(&score, scoreID)

	// add task poits
	score.Value += task.Points
	// score.CompletedTasks = append(score.CompletedTasks, task)

	if err := s.DB.Save(&score).Error; err != nil {
		return err
	}
	return nil
}
