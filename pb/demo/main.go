package main

import (
	"fmt"

	"github.com/x627234313/mmorpg-zinx/pb"
	"google.golang.org/protobuf/proto"
)

func main() {
	persion := &pb.Person{
		Name:   "wuyazi",
		Age:    23,
		Emails: []string{"danbing.at@gmail.com", "danbing@163.com"},
		Phones: []*pb.PhoneNumber{
			{
				Number: "13109890980",
				Type:   pb.PhoneType_MOBILE,
			},
			{
				Number: "18902902020",
				Type:   pb.PhoneType_WORK,
			},
			{
				Number: "152857093899",
				Type:   pb.PhoneType_HOME,
			},
		},
	}

	data, err := proto.Marshal(persion)
	if err != nil {
		fmt.Println("marshal err:", err)
		return
	}

	newPersion := &pb.Person{}

	err = proto.Unmarshal(data, newPersion)
	if err != nil {
		fmt.Println("unmarshal error:", err)
		return
	}

	fmt.Println("源数据：", persion)
	fmt.Println("解码数据：", newPersion)
}
