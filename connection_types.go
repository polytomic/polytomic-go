package polytomic

const (
	AthenaConnectionType             = "awsathena"
	PostgresqlConnectionType         = "postgresql"
	SQLServerConnectionType          = "azuresql"
	S3ConnectionType                 = "s3"
	AzureBlobConnectionType          = "azureblob"
	GoogleCloudStorageConnectionType = "gcs"
	BigQueryConnectionType           = "bigquery"
	GoogleSheetsConnectionType       = "gsheets"
	GoogleAdsConnectionType          = "googleads"
	SalesforceConnectionType         = "salesforce"
	HubspotConnectionType            = "hubspot"
	SnowflakeConnectionType          = "snowflake"
)

type SnowflakeConfiguration struct {
	Account          string `json:"account"`
	Username         string `json:"username"`
	Password         string `json:"password"`
	Database         string `json:"dbname"`
	Warehouse        string `json:"warehouse"`
	AdditionalParams string `json:"params"`
}

type HubspotConnfiguration struct {
	HubDomain string `json:"hub_domain"`
	HubId     int    `json:"hub_id"`
}

type GoogleAdsConfiguration struct {
	ResourceId struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"resource_id"`
}

type GoogleSheetsConfiguration struct {
	UserEmail     string `json:"user_email"`
	SpreadsheetID struct {
		Value string `json:"value"`
		Label string `json:"label"`
	} `json:"spreadsheet_id"`
}

type SalesforceConfiguration struct {
	Username    string `json:"username"`
	Domain      string `json:"domain"`
	InstanceURL string `json:"instance_url"`
	APIVersion  int    `json:"api_version"`
}

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

	// Client Certificates
	ClientCerts       bool   `json:"client_certs"`
	ClientCertificate string `json:"client_certificate"`
	ClientKey         string `json:"client_key"`
	CACert            string `json:"ca_cert"`

	// Change Detection
	ChangeDetection bool   `json:"change_detection"`
	Publication     string `json:"publication"`

	SSH        bool   `json:"ssh"`
	SSHUser    string `json:"ssh_user,omitempty"`
	SSHHost    string `json:"ssh_host,omitempty"`
	SSHPort    int    `json:"ssh_port,omitempty"`
	PrivateKey string `json:"ssh_private_key,omitempty"`
}
