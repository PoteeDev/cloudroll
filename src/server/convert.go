package server

import (
	"fmt"
	"log"

	"github.com/PoteeDev/cloudroll/proto"
	"github.com/PoteeDev/cloudroll/src/scoreboard/models"
)

func ConvertBoard(board *models.Board) *proto.Board {
	var res proto.Board
	log.Println(board)
	for _, column := range board.Columns {
		var cards []*proto.Card
		for _, card := range column.Cards {
			var teamTasks []*proto.TeamTasks
			for _, teamTask := range card.Tasks {
				task := proto.Task{
					Name:        teamTask.Task.Name,
					Description: teamTask.Task.Description,
					Points:      teamTask.Task.Points,
				}
				teamTasks = append(teamTasks, &proto.TeamTasks{
					Id:        fmt.Sprintf("%d", teamTask.ID),
					Completed: teamTask.Completed,
					Info:      &task,
				})
			}
			cards = append(cards, &proto.Card{
				Id:    fmt.Sprintf("%d", card.ID),
				Title: card.Title,
				Tasks: teamTasks,
			})

		}
		log.Println("cards", cards)
		res.Columns = append(res.Columns, &proto.Column{
			Id:    fmt.Sprintf("%d", column.ID),
			Name:  column.Name,
			Cards: cards,
		})
	}
	return &res
}

func ConvertProtoToBoard(board *proto.Board) *models.Board {
	var res models.Board
	log.Println(board)
	for _, column := range board.Columns {
		var cards []*models.Card
		for _, card := range column.Cards {
			var teamTasks []*models.TeamTask
			for _, teamTask := range card.Tasks {
				task := models.Task{
					Name:        teamTask.Info.Name,
					Description: teamTask.Info.Description,
					Points:      teamTask.Info.Points,
				}
				teamTasks = append(teamTasks, &models.TeamTask{
					Completed: teamTask.Completed,
					Task:      &task,
				})
			}
			cards = append(cards, &models.Card{
				Title: card.Title,
				Tasks: teamTasks,
			})

		}
		log.Println("cards", cards)
		res.Columns = append(res.Columns, &models.Column{Name: column.Name, Cards: cards})
	}
	return &res
}
