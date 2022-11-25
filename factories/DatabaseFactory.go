package factories

import (
	repository "mongomonitor/clients/repository"
)

type DatabaseFactory struct{}

func (d *DatabaseFactory) CreateDatabase(db string, connectionString string, options map[string]string) repository.IRepositoryClient {

	var client repository.IRepositoryClient

	if db == "database" {
		client = &repository.DatabaseClient{ConnectionString: connectionString, Options: options}
	} else if db == "s3" {
		client = &repository.S3Client{ConnectionString: connectionString, Options: options}
	}
	return client
}
