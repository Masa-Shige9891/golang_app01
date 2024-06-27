package main

import (
	"gorm.io/gorm"
	"time"
)

type Post struct {
	ID        int64     `json:"ID" gorm:"primaryKey"`
	Title     string    `json:"title" gorm:"unique"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

type CreatePostPayload struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

type PostRepo interface {
	createPost(*Post) (*Post, error)
}

type Storage struct {
	db *gorm.DB
}
