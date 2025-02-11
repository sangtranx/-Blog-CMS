package common

import (
	"log"
)

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

const (
	CurrentUser = "user"
)

const (
	AdminRole = "admin"
)

const (
	TopicUserLikePost    = "TopicUserLikePost"
	TopicUserDisLikePost = "TopicUserDisLikePost"
)

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("recover err:", err)
	}
}
