package data

import "go-basics/10_protocol-buffer/basic5/proto"

var LanguageCourse = []*proto.CourseInfo{
	{
		Cid:     1,
		Cname:   "English",
		Teacher: "Jack",
	},
	{
		Cid:     2,
		Cname:   "French",
		Teacher: "Rose",
	},
	{
		Cid:     3,
		Cname:   "Spanish",
		Teacher: "Crystal",
	},
}

var ComputerCourse = []*proto.CourseInfo{
	{
		Cid:     4,
		Cname:   "Golang",
		Teacher: "Michael",
	},
	{
		Cid:     5,
		Cname:   "Java",
		Teacher: "Curry",
	},
	{
		Cid:     6,
		Cname:   "Rust",
		Teacher: "Mike",
	},
}
