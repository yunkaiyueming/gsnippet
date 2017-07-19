package main

import (
	"fmt"
	"log"

	"gsnippet/serialize/pb"

	"github.com/golang/protobuf/proto"
)

func main() {
	// 为 AllPerson 填充数据
	p1 := pb.Person{
		Id:   *proto.Int32(1),
		Name: *proto.String("xiexie"),
	}

	p2 := pb.Person{
		Id:   2,
		Name: "gopher",
	}

	all_p := pb.AllPerson{
		Per: []*pb.Person{&p1, &p2},
	}

	// 对数据进行序列化
	data, err := proto.Marshal(&all_p)
	if err != nil {
		log.Fatalln("Mashal data error:", err)
	} else {
		fmt.Println(data)
	}

	// 对已经序列化的数据进行反序列化
	var target pb.AllPerson
	err = proto.Unmarshal(data, &target)
	if err != nil {
		log.Fatalln("UnMashal data error:", err)
	} else {
		for _, person := range target.Per {
			fmt.Printf("%d,%s\n", person.GetId(), person.GetName())
		}
	}
}
