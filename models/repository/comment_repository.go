package repository

import "time"

type Comment struct {
	Id       int        `json:"-"`
	Uuid     string     `json:"uuid,omitempty"`
	ParentId string     `json:"parentid,omitempty"`
	Comment  string     `json:"comment,omitempty"`
	Author   string     `json:"author,omitempty"`
	Favorite bool       `json:"favorite,omitempty"`
	CreateAt *time.Time `json:"-"`
	UpdateAt *time.Time `json:"update,omitempty"`
}
