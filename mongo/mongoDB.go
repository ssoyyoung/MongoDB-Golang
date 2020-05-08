package mongodb

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func connectDB() (client *mongo.Client, ctx context.Context, cancel context.CancelFunc) {
	// Timeout 설정을 위한 Context생성
	ctx, cancel = context.WithTimeout(context.Background(), 3*time.Second)

	// Auth에러 처리를 위한 client option 구성
	clientOptions := options.Client().ApplyURI("mongodb://49.247.134.77:27017").SetAuth(options.Credential{
		Username: "admin",
		Password: "pwdtlchd50wh",
	})

	// MongoDB 연결
	client, err := mongo.Connect(ctx, clientOptions)
	checkErr(err)

	// MongoDB 연결 검증
	checkErr(client.Ping(ctx, readpref.Primary()))

	return client, ctx, cancel
}

//InsertData func in mongo pkg
func InsertData(dataset Schema) {
	// DB 연결하기
	client, ctx, cancel := connectDB()
	// func 종료 후 mongodb 연결 끊기
	defer client.Disconnect(ctx)
	defer cancel()

	// 특정 database의 collection 연결
	testCollection := client.Database("moadata").Collection("moadata")

	// data insert 처리
	res, err := testCollection.InsertOne(ctx, dataset)
	fmt.Println(res)
	checkErr(err)
}

func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}
