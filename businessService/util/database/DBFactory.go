package database

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db, tx *gorm.DB

type Persistences interface {
	Create(model interface{}) error
	Save(model interface{})
	Find(i int, model interface{})
	FindAll(model interface{})
	FindWhere(model []interface{}, where ...interface{})
	FirstWhere(model interface{}, where ...interface{})
}

type Transactions interface {
	Execute(body func())
}

type persistence struct {}

type transaction struct {
	isActive bool
}

func Init(models ...interface{}) (Persistences, Transactions) {
	t := transaction{}
	t.setUnactive()
	var err error
	db, err = gorm.Open("postgres", "user=bms dbname=incogo password=admin sslmode=disable")
	db.LogMode(true)
	if err != nil {
		panic(err)
	}
	db.DropTableIfExists(models...)
	db.AutoMigrate(models...)
	t.setActive()
	return persistence{}, t
}

func (t transaction) Execute(body func()) {
	tx = db.Begin()
	body()
	tx.Commit()
	tx = nil
}

func (p persistence) Save(model interface{}) {
	tx.Save(model)
}

func (p persistence) Create(model interface{}) error {
	if err := tx.Create(model).Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}

func (p persistence) Find(i int, model interface{}) {
	tx.First(model, i)
}

func (p persistence) FindAll(model interface{}) {
	tx.Find(model)
}

func (p persistence) FindWhere(model []interface{}, where ...interface{}) {
	tx.Find(model, where...)
}

func (p persistence) FirstWhere(model interface{}, where ...interface{}) {
	tx.First(model, where...)
}

func (t *transaction) setActive() {
	t.isActive = true
}

func (t *transaction) setUnactive() {
	t.isActive = false
}