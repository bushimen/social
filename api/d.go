package models

import (
	"strings"
	"time"

	"github.com/go-pg/pg"
)

// Bilibili the bilibili model
type Bilibili struct {
	tableName   struct{}  `sql:"bilibili"`
	Aid         int       `sql:",pk" binding:"required"`
	Title       string    `sql:",notnull" binding:"required"`
	Description string    `sql:",notnull"`
	Favourites  int       `sql:",notnull"`
	Coins       int       `sql:",notnull"`
	Likes       int       `sql:",notnull"`
	Dislikes    int       `sql:",notnull"`
	Danmaku     int       `sql:",notnull"`
	Comments    int       `sql:",notnull"`
	Shares      int       `sql:",notnull"`
	Views       int       `sql:",notnull"`
	Duration    int       `sql:",notnull" binding:"required"`
	Thumbnail   string    `sql:",notnull" binding:"required"`
	Image       string    `sql:",notnull" binding:"required"`
	Timestamp   time.Time `binding:"required"`
}

// Exists whether the bilibili post exists in the database
func (b Bilibili) Exists() (bool, error) {
	if b.Aid == 0 {
		return false, nil
	}

	var bs []Bilibili
	return db.Model(&bs).Where("aid = ?", b.Aid).Exists()
}

// GetBilibili get a single bilibili post
func GetBilibili(aid int) (Bilibili, error) {
	b := Bilibili{
		Aid: aid,
	}

	e := db.Select(&b)

	if e == pg.ErrNoRows {
		return Bilibili{}, nil
	}

	return b, e
}

// GetBilibilis get a list of bilibili posts
func GetBilibilis(q FetchQuery) ([]Bilibili, error) {
	q.sanitize()

	var bs []Bilibili
	builder := db.Model(&bs).Order(q.Order)

	if q.Limit > 0 {
		builder = builder.Limit(q.Limit).Offset(q.Limit * (q.Page - 1))
	}

	fields := strings.Split(q.Fields, ",")
	builder = builder.Column(fields...)

	err := builder.Select()

	return bs, err
}

// UpsertBilibili upserts an bilibili post
func UpsertBilibili(b *Bilibili) error {
	exists, e := b.Exists()

	if e != nil {
		return e
	}

	if exists {
		return db.Update(b)
	}

	return db.Insert(b)
}
