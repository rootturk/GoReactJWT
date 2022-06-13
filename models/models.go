package models

import "time"

type Topics struct {
	Id            uint
	Title         string
	CreatedUserId uint
	CategoryId    int
	CreateDate    time.Time
}

type Replies struct {
	Id            uint
	TopicId       uint
	CreatedUserId uint
	CreateDate    time.Time
	IpAddress     string
}

type Login struct {
	Id        uint
	UserId    uint
	CreatedOn time.Time
	IpAddress string
	Token     string
}

type BlackList struct {
	Id        uint
	UserId    uint
	Token     string
	CreatedOn time.Time
}
