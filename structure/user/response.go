package user

import (
	"sampleapi/repository/db"
)

type AuthorizationInfo struct {
	Token string `json:"token" doc:"トークン"`
	db.MUser
}
