package mongo

type MongoClient interface{
	getLogsWithRange()
}