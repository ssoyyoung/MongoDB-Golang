package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

//Schema data struct
type Schema struct {
	ID     primitive.ObjectID `bson:"_id,omitempty"`
	Title  string             `bson:"title,omitempty"`
	Author string             `bson:"author,omitempty"`
	Tags   []string           `bson:"tags,omitempty"`
}

//MongoDB :  golang에서 mongoDB CRUD 테스트
func MongoDB() {
	// input으로 넣을 데이터 정의 struct구조
	dataset := Schema{
		Title:  "CRUD Operation in MongoDB using Golang",
		Author: "Soyoung Park",
		Tags:   []string{"book", "reading", "coding", "new"},
	}
	InsertData(dataset)
}

//InsertData func in mongo pkg
func InsertData(dataset Schema) {
	// timeout 설정을 위한 Context 생성
	ctx, _ := context.WithTimeout(context.Background(), 3*time.Second)

	// Authetication 에러 처리를 위한 client option 구성
	clientOptions := options.Client().ApplyURI("mongodb://49.247.134.77:27017").SetAuth(options.Credential{
		Username: "admin",
		Password: "pwdtlchd50wh",
	})

	// mongodb 연결
	client, err := mongo.Connect(ctx, clientOptions)
	checkErr(err)

	// 연결 검증
	err = client.Ping(context.Background(), nil)
	checkErr(err)

	// 함수 종료 후 mongodb 연결 끊기
	defer client.Disconnect(ctx)

	// 특정 database의 collection 연결
	testCollection := client.Database("moadata").Collection("testCollection")

	// data insert 처리
	res, err := testCollection.InsertOne(ctx, dataset)
	fmt.Println(res)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
