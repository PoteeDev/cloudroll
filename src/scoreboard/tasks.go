package scoreboard

import (
	"errors"
	"strings"

	"github.com/PoteeDev/cloudroll/src/scoreboard/models"
)

func (s *ScoreBoard) AddTask(name, description string, points int64) (*models.Task, error) {
	t := models.Task{
		Name:        name,
		Description: description,
		Points:      points,
	}
	res := s.DB.Create(&t)
	if res.Error != nil {
		switch {
		case strings.Contains(res.Error.Error(), ErrDublicate):
			return nil, errors.New("task already exists")
		default:
			return nil, res.Error
		}
	}
	return &t, nil
}
func (s *ScoreBoard) ShowTasks() ([]models.Task, error) {
	var tasks []models.Task
	if err := s.DB.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil

}

func (s *ScoreBoard) DeleteTask(taskId string) error {
	return s.DB.Delete(&models.Task{}, taskId).Error
}

func (s *ScoreBoard) AddTasksFromFile(filepath string) ([]*models.Task, error) {
	return nil, nil
}
