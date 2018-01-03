package utils

import mgo "gopkg.in/mgo.v2"

// EnsureIndex ensures the mapping is unique and no collisions would occur
func EnsureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()

	c := session.DB("link-shortener").C("urls")

	index := mgo.Index{
		Key:        []string{"$text:shorturl"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}

	err := c.EnsureIndex(index)

	if err != nil {
		panic(err)
	}
}
