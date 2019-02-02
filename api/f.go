package models

import (
	"strings"
	"time"

	"github.com/go-pg/pg"
)

// Instagram the instagram model
type Instagram struct {
	tableName struct{}  `sql:"instagram"`
	Shortcode string    `sql:",pk" binding:"required" json:"shortcode,omitempty"`
	Caption   string    `sql:",notnull" json:"caption,omitempty"`
	Tags      []string  `sql:",array,notnull" json:"tags,omitempty"`
	Likes     int       `sql:",notnull" json:"likes,omitempty"`
	Comments  int       `sql:",notnull" json:"comments,omitempty"`
	Type      string    `sql:",notnull" binding:"required" json:"type,omitempty"`
	Thumbnail string    `sql:",notnull" binding:"required" json:"thumbnail,omitempty"`
	Image     string    `sql:",notnull" binding:"required" json:"image,omitempty"`
	Timestamp time.Time `binding:"required" json:"timestamp"`
}

// Exists whether the instagram post exists in the database
func (ins Instagram) Exists() (bool, error) {
	if ins.Shortcode == "" {
		return false, nil
	}

	var inses []Instagram
	return db.Model(&inses).Where("shortcode = ?", ins.Shortcode).Exists()
}

// GetInstagram get a single instagram post
func GetInstagram(shortcode string) (Instagram, error) {
	ins := Instagram{
		Shortcode: shortcode,
	}

	err := db.Select(&ins)

	if err == pg.ErrNoRows {
		return Instagram{}, nil
	}

	return ins, err
}

// GetInstagrams get a list of instagram posts
func GetInstagrams(q FetchQuery) ([]Instagram, error) {
	q.sanitize()

	var ins []Instagram

	builder := db.Model(&ins).Order(q.Order).Where("type = ?", "image")

	if q.Limit > 0 {
		builder = builder.Limit(q.Limit).Offset(q.Limit * (q.Page - 1))
	}

	fields := strings.Split(q.Fields, ",")
	builder = builder.Column(fields...)

	err := builder.Select()

	return ins, err
}

// UpsertInstagram upserts an instagram post
func UpsertInstagram(ins *Instagram) error {
	exists, err := ins.Exists()

	if err != nil {
		return err
	}

	if exists {
		return db.Update(ins)
	}

	return db.Insert(ins)
}
