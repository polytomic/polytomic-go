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
	Account          string `json:"account" mapstructure:"account"`
	Username         string `json:"username" mapstructure:"username"`
	Password         string `json:"password" mapstructure:"password"`
	Database         string `json:"dbname" mapstructure:"dbname"`
	Warehouse        string `json:"warehouse" mapstructure:"warehouse"`
	AdditionalParams string `json:"params" mapstructure:"params"`
}

type HubspotConnfiguration struct {
	HubDomain string `json:"hub_domain" mapstructure:"hub_domain"`
	HubId     int    `json:"hub_id" mapstructure:"hub_id"`
}

type GoogleAdsConfiguration struct {
	ResourceId struct {
		Value string `json:"value" mapstructure:"value"`
		Label string `json:"label" mapstructure:"label"`
	} `json:"resource_id"`
}

type GoogleSheetsConfiguration struct {
	UserEmail     string `json:"user_email" mapstructure:"user_email"`
	SpreadsheetID struct {
		Value string `json:"value" mapstructure:"value"`
		Label string `json:"label" mapstructure:"label"`
	} `json:"spreadsheet_id" mapstructure:"spreadsheet_id"`
}

type SalesforceConfiguration struct {
	Username    string `json:"username" mapstructure:"username"`
	Domain      string `json:"domain" mapstructure:"domain"`
	InstanceURL string `json:"instance_url" mapstructure:"instance_url"`
	APIVersion  int    `json:"api_version" mapstructure:"api_version"`
}

type BigQueryConfiguration struct {
	ServiceAccountCredentials string `json:"service_account" mapstructure:"service_account"`
	Location                  string `json:"location" mapstructure:"location"`
}

type GCSConfiguration struct {
	ProjectId                 string `json:"project_id" mapstructure:"project_id"`
	ServiceAccountCredentials string `json:"service_account" mapstructure:"service_account"`
	Bucket                    string `json:"bucket" mapstructure:"bucket"`
}

type AzureBlobConfiguration struct {
	AccountName   string `json:"account_name" mapstructure:"account_name"`
	AccessKey     string `json:"access_key" mapstructure:"access_key"`
	ContainerName string `json:"container_name" mapstructure:"container_name"`
}

type S3Configuration struct {
	AccessKeyID     string `json:"aws_access_key_id" mapstructure:"aws_access_key_id"`
	AccessKeySecret string `json:"aws_secret_access_key" mapstructure:"aws_secret_access_key"`
	Region          string `json:"s3_bucket_region" mapstructure:"s3_bucket_region"`
	Bucket          string `json:"s3_bucket_name" mapstructure:"s3_bucket_name"`
}

type AthenaConfiguration struct {
	AccessKeyID     string `json:"access_id" mapstructure:"access_id"`
	AccessKeySecret string `json:"secret_access_key" mapstructure:"secret_access_key"`
	Region          string `json:"region" mapstructure:"region"`
	OutputBucket    string `json:"outputbucket" mapstructure:"outputbucket"`
}

type SQLServerConfiguration struct {
	Hostname string `json:"hostname" mapstructure:"hostname"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
	Port     int    `json:"port" mapstructure:"port"`
}

type PostgresqlConfiguration struct {
	Hostname string `json:"hostname" mapstructure:"hostname"`
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Database string `json:"database" mapstructure:"database"`
	Port     int    `json:"port" mapstructure:"port"`
	SSL      bool   `json:"ssl" mapstructure:"ssl"`

	// Client Certificates
	ClientCerts       bool   `json:"client_certs" mapstructure:"client_certs"`
	ClientCertificate string `json:"client_certificate" mapstructure:"client_certificate"`
	ClientKey         string `json:"client_key" mapstructure:"client_key"`
	CACert            string `json:"ca_cert" mapstructure:"ca_cert"`

	// Change Detection
	ChangeDetection bool   `json:"change_detection" mapstructure:"change_detection"`
	Publication     string `json:"publication" mapstructure:"publication"`

	SSH        bool   `json:"ssh" mapstructure:"ssh"`
	SSHUser    string `json:"ssh_user,omitempty" mapstructure:"ssh_user"`
	SSHHost    string `json:"ssh_host,omitempty" mapstructure:"ssh_host"`
	SSHPort    int    `json:"ssh_port,omitempty" mapstructure:"ssh_port"`
	PrivateKey string `json:"ssh_private_key,omitempty" mapstructure:"ssh_private_key"`
}
