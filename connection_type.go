package polytomic

const (
	AthenaConnectionType     = "awsathena"
	PostgresqlConnectionType = "postgresql"
	SQLServerConnectionType  = "azuresql"
	S3ConnectionType         = "s3"
)

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
