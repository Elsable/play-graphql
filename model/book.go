package model

import (
	"context"
	"github.com/globalsign/mgo"
	"github.com/globalsign/mgo/bson"
	"time"
)

type IBookDAO interface {
	Find(ctx context.Context, id uint) (*Book, error)
	FindAll(ctx context.Context) ([]Book, error)
	Insert(ctx context.Context, books ...Book) error
	Update(ctx context.Context, book *Book) error
	Remove(ctx context.Context, id uint) error
	FilterByStatus(_ context.Context, status string) ([]Book, error)
	FilterByIDs(_ context.Context, ids []uint) ([]Book, error)
}

var BookDAO IBookDAO = bookDAOHelper{} // just to check if interface is implemented correctly


type bookDAOHelper struct {
	baseDao
}

type Book struct {
	ID               uint `bson:"_id,omitempty"`
	Title            string        `bson:"title"`
	ISBN             string        `bson:"isbn"`
	PageCount        int           `bson:"pageCount"`
	PublishedDate    time.Time     `bson:"publishedDate"`
	ThumbnailURL     string        `bson:"thumbnailUrl"`
	ShortDescription string        `bson:"shortDescription"`
	LongDescription  string        `bson:"longDescription"`
	Status           string        `bson:"status"`
	Authors          []string      `bson:"authors"`
	Categories       []string      `json:"categories"`
}


func (b *Book) Save() error {
	return BookDAO.Insert(context.Background(), *b)
}

func (b *Book) Destroy() error {
	return BookDAO.Remove(context.Background(), b.ID)
}

func InitBookDAO(db *mgo.Database) {
	BookDAO = NewBookDAO(db)
}

func NewBookDAO(db *mgo.Database) IBookDAO {
	return bookDAOHelper{baseDao{db.C("book")}}
}


func (dao bookDAOHelper) getCollection() *mgo.Collection {
	return dao.baseDao.coll
}

func (dao bookDAOHelper) Find(_ context.Context, id uint) (*Book, error) {
	c, clean := dao.ready()
	defer clean()
	var book Book
	err := c.Find(bson.M{"_id": id}).One(&book)
	return &book, err
}

func (dao bookDAOHelper) FindAll(_ context.Context) ([]Book, error) {
	c, clean := dao.ready()
	defer clean()

	books := make([]Book, 0)
	err := c.Find(bson.M{}).All(&books)
	return books, err
}

func (dao bookDAOHelper) Insert(_ context.Context, books ...Book) error {
	c, clean := dao.ready()
	defer clean()
	err := c.Insert(books)
	return err
}

func (dao bookDAOHelper) Update(_ context.Context, book *Book) error {
	c, clean := dao.ready()
	defer clean()
	return c.Update(bson.M{"_id": book.ID}, book)
}

func (dao bookDAOHelper) Remove(_ context.Context, id uint) error {
	c, clean := dao.ready()
	defer clean()
	return c.Remove(bson.M{"_id": id})
}

func (dao bookDAOHelper) FilterByStatus(_ context.Context, status string) ([]Book, error) {
	c, clean := dao.ready()
	defer clean()

	books := make([]Book, 0)
	err := c.Find(bson.M{"status": status}).All(&books)
	return books, err
}

func (dao bookDAOHelper) FilterByIDs(_ context.Context, ids []uint) ([]Book, error) {
	c, clean := dao.ready()
	defer clean()

	books := make([]Book, 0)
	err := c.Find(bson.M{"_id": ids}).All(&books)
	return books, err
}


