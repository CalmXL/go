package data

type Phone struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
}

var Phones = []Phone{
	{
		Id:   1,
		Name: "Iphone15 prx max",
	},
	{
		Id:   2,
		Name: "HUAWEI Mate60 pro+",
	},
	{
		Id:   3,
		Name: "xiaomi 14pro",
	},
}
