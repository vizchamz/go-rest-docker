package models

import "gorm.io/gorm"

type Fact struct {
    gorm.Model
    ID       uint   `json:"ID" gorm:"primaryKey;autoIncrement"`
    Question string `json:"question" gorm:"text;not null;default:null`
    Answer   string `json:"answer" gorm:"text;not null;default:null`
}