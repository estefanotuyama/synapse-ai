package models

import (
	"time"
)

/*
This file defines the structure of our mongoDB database.
*/

type User struct {
	Email      string              `bson:"email"`
	Provider   string              `bson:"provider"`
	ProviderId int                 `bson:"providerId"`
	Name       string              `bson:"name"`
	IsPremium  bool                `bson:"isPremium"`
	Agents     []map[string]string `bson:"agents"`
}

type Agent struct {
	AgentId        string              `bson:"agentId"`
	Owner          string              `bson:"owner"`
	Name           string              `bson:"name"`
	Purpose        string              `bson:"purpose"`
	Instructions   string              `bson:"purpose"`
	CollectionName string              `bson:"collectionName"`
	CreatedAt      time.Time           `bson:"createdAt"`
	LastUsed       time.Time           `bson:"lastUsed"`
	DocumentNames  string              `bson:"documentNames"`
	Chats          []map[string]string `bson:"chats"`
}

type Chat struct {
	ChatId         string              `bson:"chatId"`
	AgentId        string              `bson:"agentId"`
	Name           string              `bson:"name"`
	CreatedAt      time.Time           `bson:"createdAt"`
	LastUpdated    time.Time           `bson:"lastUpdated"`
	MessageHistory []map[string]string `bson:"messageHistory"`
}
