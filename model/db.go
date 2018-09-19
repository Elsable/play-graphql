package model

import "github.com/globalsign/mgo"
// todo still need to find a good pattern for organize dao layer
type baseDao struct {
	coll *mgo.Collection
}

func (dao *baseDao)ready() (*mgo.Collection, func()) {
	db := dao.coll.Database
	s := db.Session.Copy()
	return s.DB(db.Name).C(dao.coll.Name), func() {
		s.Close()
	}
}
