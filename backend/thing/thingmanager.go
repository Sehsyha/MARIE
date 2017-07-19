package thing

import (
	"log"

	"gopkg.in/mgo.v2/bson"

	"github.com/Zenika/MARIE/backend/config"

	"github.com/Zenika/MARIE/backend/utils"
)

// Create a new thing and add it to the database
func Create(t Thing) {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	c := s.DB(cfg.DbName).C(ThingCollectionName)
	err := c.Insert(t)
	if err != nil {
		log.Fatal(err)
	}

}

// ReadAll things in database
func ReadAll() []Thing {
	cfg := config.Load()
	s := utils.GetSession()
	defer s.Close()

	var things []Thing

	c := s.DB(cfg.DbName).C(ThingCollectionName)
	err := c.Find(bson.M{}).All(&things)
	if err != nil {
		log.Fatal(err)
	}

	return things
}

// Read a thing in the database with its id
func Read(id bson.ObjectId) (Thing, error) {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	c := s.DB(cfg.DbName).C(ThingCollectionName)

	res := Thing{}
	err := c.FindId(id).One(&res)

	if err != nil {
		return res, err
	}

	return res, nil
}

// Update a thing in database
func Update(t Thing) error {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	c := s.DB(cfg.DbName).C(ThingCollectionName)

	return c.Update(bson.M{"_id": t.ID}, bson.M{"getters": t.Getters,
		"actions":  t.Actions,
		"location": t.Location,
		"protocol": t.Protocol,
		"name":     t.Name,
		"type":     t.Type})
}

// ReadGetterName return things that have a getter with the given name
func ReadGetterName(name string) []Thing {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	// Select all things with this parameter
	c := s.DB(cfg.DbName).C(ThingCollectionName)
	things := []Thing{}

	err := c.Pipe([]bson.M{{"$match": bson.M{"getters.name": name}}}).All(&things)
	if err != nil {
		log.Fatal(err)
	}
	return things
}

// ReadMacAddress return thing with mac address
func ReadMacAddress(mac string) (Thing, error) {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	c := s.DB(cfg.DbName).C(ThingCollectionName)
	t := Thing{}

	err := c.Find(bson.M{"macaddress": mac}).One(t)

	if err != nil {
		return t, err
	}

	return t, nil
}

// ReadActionName return things that have an action with the given name
func ReadActionName(name string) []Thing {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	c := s.DB(cfg.DbName).C(ThingCollectionName)
	things := []Thing{}

	err := c.Pipe([]bson.M{{"$match": bson.M{"actions.name": name}}}).All(&things)
	if err != nil {
		log.Fatal(err)
	}
	return things
}

// Delete the thing from the database
func Delete(id bson.ObjectId) error {
	cfg := config.Load()

	s := utils.GetSession()
	defer s.Close()

	c := s.DB(cfg.DbName).C(ThingCollectionName)

	return c.RemoveId(id)
}
