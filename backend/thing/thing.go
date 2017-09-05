package thing

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

const (
	// CollectionName represents the collection name in the mongo db
	CollectionName = "things"
)

// Thing represents a connected object
type Thing struct {
	ID            bson.ObjectId `json:"id" bson:"_id"`
	MacAddress    string        `json:"macaddress"`
	IPAddress     string        `json:"ipaddress"`
	Type          string        `json:"type"`
	Name          string        `json:"name"`
	Location      string        `json:"location"`
	Protocol      string        `json:"protocol"`
	Actions       []Action      `json:"actions"`
	Getters       []Getter      `json:"getters"`
	State         bool          `json:"state"`
	LastHeartBeat time.Time
}

// Action represents what a thing can do
type Action struct {
	Name       string      `json:"name"`
	Parameters []Parameter `json:"parameters"`
}

// Parameter represents what an action needs to be executed
type Parameter struct {
	Name  string      `json:"name"`
	Type  string      `json:"type"`
	Value interface{} `json:"value"`
}

// Getter represents what information a thing can give
type Getter struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

// DoRequest represents a request to do something
type DoRequest struct {
	Name       string      `json:"name"`
	MacAddress string      `json:"macaddress"`
	Parameters []Parameter `json:"parameters"`
}

// GetRequest represents a request to get something
type GetRequest struct {
	Name       string `json:"name"`
	MacAddress string `json:"macaddress"`
}
