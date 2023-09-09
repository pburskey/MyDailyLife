package domain

type Party interface {
	ID() string
}

type Person struct {
	ID    string `json:"id" bson:"id,omitempty"`
	First string `json:"first" bson:"first,omitempty"`
	Last  string `json:"last" bson:"last,omitempty"`
}
