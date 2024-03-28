package models

import (
	"gorm.io/gorm"
)

type Clipboard struct {
	gorm.Model
	Context string `gorm:"context"`
	Digest  string `gorm:"digest"`
	Short   string `gorm:"short"`
	Size    int    `gorm:"size"`
	Url     string `gorm:"url"`
}
type Clipboard_VIP struct {
	gorm.Model
	Context         string `gorm:"context"`
	Digest          string `gorm:"digest"`
	Short           string `gorm:"short"`
	Size            int    `gorm:"size"`
	Url             string `gorm:"url"`
	Username        string `gorm:"username"`
	Visible_VIP     string `gorm:"visible_VIP"`
	Visible_another string `gorm:"visible_another"`
}

var Db *gorm.DB

type User struct {
	gorm.Model
	Username string `gorm:"username"`
	Password string `gorm:"password"`
}
