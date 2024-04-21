package data

type InfoBody struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Age  uint64 `json:"age"`
}

var Infos = []InfoBody{
	{
		Id:   1,
		Name: "Mike",
		Age:  25,
	},
	{
		Id:   2,
		Name: "Nike",
		Age:  17,
	},
	{
		Id:   3,
		Name: "Mark",
		Age:  29,
	},
}
