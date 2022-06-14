package controllers

import (
	"strconv"
	"time"

	"github.com/gofiber/fiber"
	"github.com/rootturk/goreactjwt/database"
	"github.com/rootturk/goreactjwt/models"
)

//TOPIC

func CreateNewTopic(c *fiber.Ctx) error {
	var data map[string]string

	if err := c.BodyParser(&data); err != nil {
		c.Status(503).Send(err)
	}

	userId, err := strconv.Atoi(data["created_user_id"])

	if err != nil {
		c.Status(401).Send(err)
	}

	topic := models.Topics{
		Title:         data["title"],
		CreatedUserId: uint(userId),
		CategoryId:    0,
		LastReplyDate: time.Time{},
		CreateDate:    time.Now(),
	}

	database.DB.Create(&topic)

	return c.JSON(topic)
}

func GetAllTopic() {

}

func GetTopicById() {

}

func DeleteTopic() {

}

func UpdateTopic() {

}

// MESSAGE

func GetAllRepliesFromTopicById() {

}

func GetReplyById() {

}

func DeleteReplyById() {

}

func DeleteReplyByTopicId() {

}
