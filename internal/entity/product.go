package entity

type Product struct {
	ID    string `bson:"_id"`
	Name  string `bson:"name"`
	Votes int    `bson:"votes"`
}
