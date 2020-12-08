package model

import (
	"github.com/pangxianfei/framework/helpers/zone"
)

type Webroutes struct {
	BaseModel
	ID         uint       `gorm:"column:id;primary_key;auto_increment"`
	RouteName  string     `gorm:"column:routename;type:varchar(255)"`
	Controller string     `gorm:"column:controller;type:varchar(255);unique_index;not null"`
	Function   string     `gorm:"column:Function;type:varchar(255);not null"`
	Path       string     `gorm:"column:path;type:varchar(255);not null"`
	CreatedAt  zone.Time  `gorm:"column:created_at"`
	UpdatedAt  zone.Time  `gorm:"column:updated_at"`
	DeletedAt  *zone.Time `gorm:"column:deleted_at"`
}

func (Webroutes *Webroutes) TableName() string {
	return Webroutes.SetTableName("webroutes")
}

/*
func (webroutes *Webroutes) SetRouteNameAttribute(value interface{}) {
	webroutes.RouteName = webroutes.RouteName
}

func (webroutes *Webroutes) GetRouteNameAttribute(value interface{}) interface{} {
	return value
}*/
