package mongodb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

// CrawlTarget struct
type CrawlTarget struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	Platform  string	     `bson:"platform,omitempty"`
	Channel   string	     `bson:"channel,omitempty"`
	ChannelID string	     `bson:"channelID,omitempty"`
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

	filter := bson.D{primitive.E{Key: "creatorDataName", Value: "soso"}}
	update := bson.D{
		{
			"$set", bson.D{
				primitive.E{Key: "newdate", Value: 199412345333999996},
			},
		},
	}
	UpdateData(filter, update)

	delfilter := bson.D{primitive.E{Key: "creatorDataName", Value: "soso"}}
	DeleteData(delfilter)
	ListData()
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

////////////////////////ADMIN FUNCTION////////////////////////////

// CrawlList func
func CrawlList() string {
	// DB 연결하기
	client, ctx, cancel := connectDB()
	// func 종료 후 mongodb 연결 끊기
	defer client.Disconnect(ctx)
	defer cancel()

	// 특정 collection 가져오기
	moaData := client.Database("meerkatonair").Collection("crawl_target")

	res, err := moaData.Find(ctx, bson.M{})
	fmt.Println(res)
	checkErr(err)

	var datas []bson.M
	if err = res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	jsonBytes, err := json.Marshal(datas)
	checkErr(err)

	jsonString := string(jsonBytes)
	//fmt.Println(jsonString)

	return jsonString
}

// SearchDBbyID func
func SearchDBbyID(id string) string {
	// DB 연결하기
	client, ctx, cancel := connectDB()
	// func 종료 후 mongodb 연결 끊기
	defer client.Disconnect(ctx)
	defer cancel()
	fmt.Println(id)
	// 특정 collection 가져오기
	moaData := client.Database("meerkatonair").Collection("crawl_target")

	// _id type 변환작업
	docID, err := primitive.ObjectIDFromHex(id)
	checkErr(err)
	res, err := moaData.Find(ctx, bson.M{"_id": docID})
	fmt.Println(res)
	checkErr(err)

	var datas []bson.M
	if err = res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}

	jsonBytes, err := json.Marshal(datas)
	checkErr(err)

	jsonString := string(jsonBytes)
	//fmt.Println(jsonString)

	return jsonString
}

// DeleteDBbyID func
func DeleteDBbyID(id string) string {
	// DB 연결하기
	client, ctx, cancel := connectDB()
	// func 종료 후 mongodb 연결 끊기
	defer client.Disconnect(ctx)
	defer cancel()
	fmt.Println(id)
	// 특정 collection 가져오기
	moaData := client.Database("meerkatonair").Collection("crawl_target")

	// _id type 변환작업
	docID, err := primitive.ObjectIDFromHex(id)
	checkErr(err)

	//delfilter := bson.D{primitive.E{Key: "_id", Value: docID}}
	delfilter := bson.M{"_id": docID}
	res, err := moaData.DeleteOne(ctx, delfilter)
	checkErr(err)
	fmt.Println(res)

	return "delete!"
}

//UpdateDBbyID func
func UpdateDBbyID(id, platform, channel, channelID string) string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	moaData := client.Database("meerkatonair").Collection("crawl_target")
	fmt.Println(moaData)
	docID, err := primitive.ObjectIDFromHex(id)
	checkErr(err)
	fmt.Println(docID)

	filter := bson.M{"_id": docID}

	update := bson.D{
		{"$set", bson.D{
				{"platform", platform},
				{"channel", channel},
				{"channelID", channelID},
			},
		},
	}
	res, err := moaData.UpdateOne(ctx, filter, update)
	fmt.Println(res)
	checkErr(err)

	return "updated!"
}

// CreateDB func
func CreateDB(platform, channel, channelID string) string {
	client, ctx, cancel := connectDB()
	defer client.Disconnect(ctx)
	defer cancel()

	moaData := client.Database("meerkatonair").Collection("crawl_target")

	newData := CrawlTarget{
		Platform:  platform,
		Channel:   channel,
		ChannelID: channelID,
	}
	// data insert 처리
	res, err := moaData.InsertOne(ctx, newData)
	fmt.Println(res)
	checkErr(err)

	return "create!"
}

////////////////////////TEST FUNCTION////////////////////////////

// ListData func
func ListData() string {
	// DB 연결하기
	client, ctx, cancel := connectDB()
	// func 종료 후 mongodb 연결 끊기
	defer client.Disconnect(ctx)
	defer cancel()

	// 특정 collection 가져오기
	moaData := client.Database("meerkatonair").Collection("live_list")

	res, err := moaData.Find(ctx, bson.M{"onLive": true})
	fmt.Println(res)
	checkErr(err)

	var datas []bson.M
	if err = res.All(ctx, &datas); err != nil {
		fmt.Println(err)
	}
	//fmt.Println(datas)

	jsonBytes, err := json.Marshal(datas)
	checkErr(err)

	jsonString := string(jsonBytes)
	//fmt.Println(jsonString)

	return jsonString
}

// DeleteData func
func DeleteData(filter bson.D) {
	// DB 연결하기
	client, ctx, cancel := connectDB()
	// func 종료 후 mongodb 연결 끊기
	defer client.Disconnect(ctx)
	defer cancel()

	// 특정 collection 가져오기
	moaData := client.Database("meerkatonair").Collection("live_list")

	res, err := moaData.DeleteOne(ctx, filter)
	checkErr(err)
	fmt.Println(res)
}

//UpdateData func
func UpdateData(filter bson.D, update bson.D) {
	// DB 연결하기
	client, ctx, cancel := connectDB()
	// func 종료 후 mongodb 연결 끊기
	defer client.Disconnect(ctx)
	defer cancel()

	// 특정 collection 가져오기
	moaData := client.Database("moadata").Collection("moadata")

	// 해당 필드가 존재하면 업데이트, 없을경우 필드와 값 추가
	res, err := moaData.UpdateOne(ctx, filter, update)

	checkErr(err)
	fmt.Println(res)
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
