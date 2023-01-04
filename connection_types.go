package polytomic

const (
	AffinityConnectionType           = "affinity"
	AirtableConnectionType           = "airtable"
	AmplitudeConnectionType          = "amplitude"
	APIConnectionType                = "api"
	AthenaConnectionType             = "awsathena"
	AzureBlobConnectionType          = "azureblob"
	SQLServerConnectionType          = "azuresql"
	BigQueryConnectionType           = "bigquery"
	CsvConnectionType                = "csv"
	GoogleCloudStorageConnectionType = "gcs"
	GoogleSheetsConnectionType       = "gsheets"
	GoogleAdsConnectionType          = "googleads"
	HubspotConnectionType            = "hubspot"
	MarketoConnectionType            = "marketo"
	MongoDBConnectionType            = "mongodb"
	PostgresqlConnectionType         = "postgresql"
	S3ConnectionType                 = "s3"
	SalesforceConnectionType         = "salesforce"
	SnowflakeConnectionType          = "snowflake"
)

type SnowflakeConfiguration struct {
	Account          string `json:"account" mapstructure:"account"`
	Username         string `json:"username" mapstructure:"username"`
	Password         string `json:"password" mapstructure:"password"`
	Dbname           string `json:"dbname" mapstructure:"dbname"`
	Warehouse        string `json:"warehouse" mapstructure:"warehouse"`
	AdditionalParams string `json:"params" mapstructure:"params"`
}

type HubspotConfiguration struct {
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
	ProjectID      string `json:"project_id" mapstructure:"project_id"`
	ServiceAccount string `json:"service_account" mapstructure:"service_account"`
	Location       string `json:"location" mapstructure:"location"`
}

type GCSConfiguration struct {
	ProjectId      string `json:"project_id" mapstructure:"project_id"`
	ServiceAccount string `json:"service_account" mapstructure:"service_account"`
	Bucket         string `json:"bucket" mapstructure:"bucket"`
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

type MongoDBConfiguration struct {
	Hosts            string `json:"hosts" mapstructure:"hosts"`
	Username         string `json:"username" mapstructure:"username"`
	Password         string `json:"password" mapstructure:"password"`
	Database         string `json:"database" mapstructure:"database"`
	SRV              bool   `json:"srv" mapstructure:"srv"`
	AdditionalParams string `json:"params" mapstructure:"params"`
}

type MarketoConfiguration struct {
	ClientID     string `json:"client_id" mapstructure:"client_id"`
	ClientSecret string `json:"client_secret" mapstructure:"client_secret"`
	RESTEndpoint string `json:"rest_endpoint" mapstructure:"rest_endpoint"`

	EnforceAPILimits  bool `json:"enforce_api_limits" mapstructure:"enforce_api_limits"`
	DailyAPICalls     int  `json:"daily_api_calls" mapstructure:"daily_api_calls"`
	ConcurrentImports int  `json:"concurrent_imports" mapstructure:"concurrent_imports"`
}

type AffinityConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key"`
}

type AirtableConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key"`
}

type AmplitudeConfiguration struct {
	APIKey    string `json:"api_key" mapstructure:"api_key"`
	SecretKey string `json:"secret_key" mapstructure:"secret_key"`
}

type APIConnectionConfiguration struct {
	URL                   string             `json:"url" mapstructure:"url"`
	Headers               []RequestParameter `json:"headers" mapstructure:"headers"`
	Body                  string             `json:"body" mapstructure:"body"`
	QueryStringParameters []RequestParameter `json:"parameters" mapstructure:"parameters"`
	Healthcheck           string             `json:"healthcheck" mapstructure:"healthcheck"`
	Auth                  Auth               `json:"auth" mapstructure:",squash"`
}

type RequestParameter struct {
	Name  string      `json:"name" mapstructure:"name"`
	Value interface{} `json:"value" mapstructure:"value"`
}

type Auth struct {
	// BasicAuthConf provides basic authentication credentials
	// e.g.
	// Authorization: Basic <base64 encoded username:password>
	// RFC7617
	Basic *BasicAuthConf `json:"basic,omitempty" mapstructure:"basic"`
	// HeaderAuthConf provides header authentication credentials
	// e.g.
	// Authorization: Bearer <token>
	Header *RequestParameter `json:"header,omitempty" mapstructure:"header"`
	// OAuthConf provides OAuth authentication using the client credentials flow
	// e.g.
	// Client ID: <client id>
	// Client Secret: <client secret>
	// Callback URL: <callback url>
	// RFC6749 Section-4.4
	OAuth *ClientCredentialConf `json:"oauth,omitempty" mapstructure:"oauth"`
}

type BasicAuthConf struct {
	Username string `json:"username,omitempty" mapstructure:"username"`
	Password string `json:"password,omitempty" mapstructure:"password"`
}

type ClientCredentialConf struct {
	ClientID      string             `json:"client_id" mapstructure:"client_id"`
	ClientSecret  string             `json:"client_secret" mapstructure:"client_secret"`
	TokenEndpoint string             `json:"token_endpoint" mapstructure:"token_endpoint"`
	ExtraFormData []RequestParameter `json:"extra_form_data" mapstructure:"extra_form_data"`
}

type CSVConnectionConfiguration struct {
	URL                   string             `json:"url" mapstructure:"url"`
	Headers               []RequestParameter `json:"headers" mapstructure:"headers"`
	QueryStringParameters []RequestParameter `json:"parameters" mapstructure:"parameters"`
	Auth                  Auth               `json:"auth" mapstructure:",squash"`
}
