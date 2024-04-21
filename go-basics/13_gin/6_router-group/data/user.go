package data

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

var Users = []User{
	{
		Id:   1,
		Name: "Mike",
	},
	{
		Id:   2,
		Name: "Rose",
	},
	{
		Id:   3,
		Name: "Jack",
	},
}
