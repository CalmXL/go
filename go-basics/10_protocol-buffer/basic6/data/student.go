package data

import (
	"go-basics/10_protocol-buffer/basic6/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
	"time"
)

var Students = []*proto.StudentResponse{
	{
		Sid:    1,
		Name:   "Mike",
		Age:    18,
		Gender: proto.Gender_MALE,
		Scores: &proto.StudentResponse_Scores{
			Interview: 60,
			Written:   80,
		},
		Course:   proto.Course_GO,
		JoinTime: timestamppb.New(time.Date(2023, 5, 3, 0, 0, 0, 0, time.UTC)),
	},
	{
		Sid:    2,
		Name:   "Rose",
		Age:    20,
		Gender: proto.Gender_FEMALE,
		Scores: &proto.StudentResponse_Scores{
			Interview: 70,
			Written:   90,
		},
		Course:   proto.Course_JAVA,
		JoinTime: timestamppb.New(time.Date(2023, 7, 3, 0, 0, 0, 0, time.UTC)),
		Subject: &proto.StudentResponse_Chemistry{
			Chemistry: 80,
		},
	},
	{
		Sid:    3,
		Name:   "Jack",
		Age:    26,
		Gender: proto.Gender_MALE,
		Scores: &proto.StudentResponse_Scores{
			Interview: 90,
			Written:   90,
		},
		Course:   proto.Course_RUST,
		JoinTime: timestamppb.New(time.Date(2023, 11, 3, 0, 0, 0, 0, time.UTC)),
		Subject: &proto.StudentResponse_Geography{
			Geography: 70,
		},
	},
}
