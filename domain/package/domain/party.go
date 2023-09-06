package domain

type Party interface {
	ID() string
}

type Person struct {
	ID          string `json:"id" bson:"id,omitempty"`
	First, Last string
}
