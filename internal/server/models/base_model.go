// Package models 存放模型基类
package models

type BaseMode struct {
	ID uint `json:"id,omitempty";gorm:"column:id;primaryKey;autoIncrement"`
}
