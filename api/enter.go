package api

import (
	"GameManageSystem/api/game_api"
	"GameManageSystem/api/score_api"
	"GameManageSystem/api/user_api"
)

type Group struct {
	UserApi  user_api.UserApi
	GameApi  game_api.GameApi
	ScoreApi score_api.ScoreApi
}

var ApiGroup = new(Group)
