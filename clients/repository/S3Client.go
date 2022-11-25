package repository

import (
	json "encoding/json"
	fmt "fmt"
	types "mongomonitor/types"

	"github.com/aws/aws-sdk-go/aws/session"
	s3 "github.com/aws/aws-sdk-go/service/s3"
)

type S3Client struct {
	ConnectionString string
	Options          map[string]string
	client           *s3.S3
	database         string
	collection       string
}

func (S3Client *S3Client) Insert(docs []interface{}) types.WriteResponse {
	s3Client := S3Client.getDatabaseInstance()
	//coll := s3Client.database(S3Client.database).Collection(S3Client.collection)

	//_, err := coll.InsertMany(context.TODO(), docs)
	//if err != nil {
	//	panic(err)
	//}

	fmt.Println(s3Client)

	return types.WriteResponse{Inserted: 0, Success: true}
}

func (s3Client *S3Client) FindById() types.WriteResponse {
	return types.WriteResponse{Inserted: 0, Success: true}
}

func (s3Client *S3Client) FindAll() types.WriteResponse {
	return types.WriteResponse{Inserted: 0, Success: true}
}
func (s3Client *S3Client) SetDatabase(database string) {
	s3Client.database = database
}

func (s3Client *S3Client) SetCollection(collection string) {
	s3Client.collection = collection
}

func (s3Client *S3Client) getDatabaseInstance() *s3.S3 {

	//clientOpts := options.Client().ApplyURI(s3Client.ConnectionString)
	client := s3.New(session.New())

	response, _ := json.Marshal(client)

	fmt.Println(string(response))

	s3Client.client = client

	return client
}
