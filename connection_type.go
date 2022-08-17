package polytomic

const (
	AthenaConnectionType     = "awsathena"
	PostgresqlConnectionType = "postgresql"
	SQLServerConnectionType  = "azuresql"
)

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
