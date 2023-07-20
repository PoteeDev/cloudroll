package server

import (
	"context"
	"fmt"
	"log"

	proto "github.com/PoteeDev/cloudroll/proto"
	"github.com/PoteeDev/cloudroll/src/scoreboard"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Implements of EchoServiceServer

type echoServer struct {
	proto.UnimplementedCloudrollServiceServer
	sb *scoreboard.ScoreBoard
}

func newCloudrollServer() proto.CloudrollServiceServer {
	sb, err := scoreboard.Init()
	if err != nil {
		log.Fatalln(err)
	}

	sb.GenerateBoard()
	return &echoServer{
		sb: sb,
	}
}

func (s *echoServer) Ping(ctx context.Context, req *proto.Empty) (*proto.EchoMessage, error) {

	return &proto.EchoMessage{
		Value: "pong",
	}, nil
}

func (s *echoServer) CreateTeam(ctx context.Context, req *proto.TeamCreateReq) (*proto.TeamInfo, error) {
	team, err := s.sb.CreateTeam(req.GetName())
	if err != nil {
		return nil, status.Error(400, err.Error())
	}
	return &proto.TeamInfo{
		Name:   team.Name,
		Invite: fmt.Sprintf("/team/join/%d", team.ID),
	}, nil
}

func (s *echoServer) JoinTeam(ctx context.Context, req *proto.JoinTeamReq) (*proto.TeamInfo, error) {
	userId, err := GetUserInfoID(ctx)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}

	team, err := s.sb.JoinTeam(req.GetId(), userId)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}
	return &proto.TeamInfo{
		Name:   team.Name,
		Invite: fmt.Sprintf("/team/join/%d", team.ID),
	}, nil
}

func (s *echoServer) AddTask(ctx context.Context, req *proto.Task) (*proto.Task, error) {
	// validata roles
	t, err := s.sb.AddTask(req.Name, req.Description, req.Points)
	if err != nil {
		return nil, status.Error(400, err.Error())
	}
	return &proto.Task{
		Id:          fmt.Sprintf("%d", t.ID),
		Name:        t.Name,
		Description: t.Description,
		Points:      t.Points,
	}, nil
}
func (s *echoServer) ShowTasks(ctx context.Context, req *proto.Empty) (*proto.Tasks, error) {
	tasks, err := s.sb.ShowTasks()
	if err != nil {
		return nil, status.Error(400, err.Error())
	}
	var outTasks proto.Tasks
	for _, t := range tasks {
		outTasks.Tasks = append(outTasks.Tasks, &proto.Task{
			Id:          fmt.Sprintf("%d", t.ID),
			Name:        t.Name,
			Description: t.Description,
			Points:      t.Points,
		})
	}
	return &outTasks, nil

}

// func (s *echoServer) DeleteTask(ctx context.Context, req *proto.Task) (*proto.Empty, error) {
// }

func (s *echoServer) GetScoreboard(ctx context.Context, req *proto.Empty) (*proto.ScoreboardResponse, error) {
	var scores []*proto.Score
	for _, score := range s.sb.GetScoreboard() {
		scores = append(scores, &proto.Score{Score: score.Value})
	}

	return &proto.ScoreboardResponse{
		Time:  timestamppb.Now(),
		Score: scores,
	}, nil
}

func (s *echoServer) SubmitTask(ctx context.Context, req *proto.SubmitRequest) (*proto.EchoMessage, error) {
	log.Println(req.Task.GetName())
	return &proto.EchoMessage{
		Value: "ok",
	}, nil
}

// Board

func (s *echoServer) GetBoard(ctx context.Context, req *proto.Empty) (*proto.Board, error) {

	board, err := s.sb.GetBoard()
	if err != nil {
		return nil, status.Error(400, err.Error())
	}
	return ConvertBoard(board), nil
}

func (s *echoServer) UpdateBoard(ctx context.Context, req *proto.Board) (*proto.Board, error) {

	err := s.sb.UpdateBoard(ConvertProtoToBoard(req))
	if err != nil {
		return nil, status.Error(400, err.Error())
	}
	return req, nil
}
