package polytomic

const (
	AthenaConnectionType             = "awsathena"
	PostgresqlConnectionType         = "postgresql"
	SQLServerConnectionType          = "azuresql"
	S3ConnectionType                 = "s3"
	AzureBlobConnectionType          = "azureblob"
	GoogleCloudStorageConnectionType = "gcs"
	BigQueryConnectionType           = "bigquery"
)

type BigQueryConfiguration struct {
	ServiceAccountCredentials string `json:"service_account"`
	Location                  string `json:"location"`
}

type GCSConfiguration struct {
	ProjectId                 string `json:"project_id"`
	ServiceAccountCredentials string `json:"service_account"`
	Bucket                    string `json:"bucket"`
}

type AzureBlobConfiguration struct {
	AccountName   string `json:"account_name"`
	AccessKey     string `json:"access_key"`
	ContainerName string `json:"container_name"`
}

type S3Configuration struct {
	AccessKeyID     string `json:"aws_access_key_id"`
	AccessKeySecret string `json:"aws_secret_access_key"`
	Region          string `json:"s3_bucket_region"`
	Bucket          string `json:"s3_bucket_name"`
}

type AthenaConfiguration struct {
	AccessKeyID     string `json:"access_id"`
	AccessKeySecret string `json:"secret_access_key"`
	Region          string `json:"region"`
	OutputBucket    string `json:"outputbucket"`
}

type SQLServerConfiguration struct {
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     int    `json:"port"`
}

type PostgresqlConfiguration struct {
	Hostname string `json:"hostname"`
	Username string `json:"username"`
	Password string `json:"password"`
	Database string `json:"database"`
	Port     int    `json:"port"`
	SSL      bool   `json:"ssl"`
}
