package types

type Student struct {
	Name  string `bson:"name"`
	Grade int    `bson:"grade"`
}
