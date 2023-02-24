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
	BingAdsConnectionType            = "bingads"
	ChargebeeConnectionType          = "chargebee"
	CloudSQLConnectionType           = "cloudsql"
	CosmosDBConnectionType           = "cosmosdb"
	CustomerIOConnectionType         = "customerio"
	CsvConnectionType                = "csv"
	DialpadConnectionType            = "dialpad"
	FreshdeskConnectionType          = "freshdesk"
	FullstoryConnectionType          = "fullstory"
	HarmonicConnectionType           = "harmonic"
	KlaviyoConnectionType            = "klaviyo"
	KustomerConnectionType           = "kustomer"
	LobConnectionType                = "lob"
	GoogleCloudStorageConnectionType = "gcs"
	GoogleSheetsConnectionType       = "gsheets"
	GoogleAdsConnectionType          = "googleads"
	HubspotConnectionType            = "hubspot"
	IntercomConnectionType           = "intercom"
	IterableConnectionType           = "iterable"
	MarketoConnectionType            = "marketo"
	MongoDBConnectionType            = "mongodb"
	MysqlConnectionType              = "mysql"
	NetsuiteConnectionType           = "netsuite"
	PipedriveConnectionType          = "pipedrive"
	PostgresqlConnectionType         = "postgresql"
	RedshiftConnectionType           = "redshift"
	S3ConnectionType                 = "s3"
	SalesforceConnectionType         = "salesforce"
	SegmentConnectionType            = "segment"
	ShipBobConnectionType            = "shipbob"
	ShopifyConnectionType            = "shopify"
	SmartSheetConnectionType         = "smartsheet"
	SnowflakeConnectionType          = "snowflake"
	StripeConnectionType             = "stripe"
	SurvicateConnectionType          = "survicate"
	SynapseConnectionType            = "synapse"
	UserVoiceConnectionType          = "uservoice"
	VanillaConnectionType            = "vanilla"
	WebhookConnectionType            = "webhook"
	ZendeskConnectionType            = "zendesk"
)

type SnowflakeConfiguration struct {
	Account   string `json:"account" mapstructure:"account,omitempty" tfsdk:"account"`
	Username  string `json:"username" mapstructure:"username,omitempty" tfsdk:"username"`
	Password  string `json:"password" mapstructure:"password,omitempty" tfsdk:"password"`
	Dbname    string `json:"dbname" mapstructure:"dbname,omitempty" tfsdk:"dbname"`
	Warehouse string `json:"warehouse" mapstructure:"warehouse,omitempty" tfsdk:"warehouse"`
	Params    string `json:"params" mapstructure:"params,omitempty" tfsdk:"params"`
}

type HubspotConfiguration struct {
	HubDomain string `json:"hub_domain" mapstructure:"hub_domain" tfsdk:"hub_domain"`
	HubId     int    `json:"hub_id" mapstructure:"hub_id" tfsdk:"hub_id"`
}

type GoogleAdsConfiguration struct {
	ResourceId struct {
		Value string `json:"value" mapstructure:"value" tfsdk:"value"`
		Label string `json:"label" mapstructure:"label" tfsdk:"label"`
	} `json:"resource_id"`
}

type GoogleSheetsConfiguration struct {
	UserEmail     string `json:"user_email" mapstructure:"user_email" tfsdk:"user_email"`
	SpreadsheetID struct {
		Value string `json:"value" mapstructure:"value" tfsdk:"value"`
		Label string `json:"label" mapstructure:"label" tfsdk:"label"`
	} `json:"spreadsheet_id" mapstructure:"spreadsheet_id" tfsdk:"spreadsheet_id"`
}

type SalesforceConfiguration struct {
	Username    string `json:"username" mapstructure:"username" tfsdk:"username"`
	Domain      string `json:"domain" mapstructure:"domain" tfsdk:"domain"`
	InstanceURL string `json:"instance_url" mapstructure:"instance_url" tfsdk:"instance_url"`
	APIVersion  int    `json:"api_version" mapstructure:"api_version" tfsdk:"api_version"`
}

type BigQueryConfiguration struct {
	ProjectID      string `json:"project_id" mapstructure:"project_id" tfsdk:"project_id"`
	ServiceAccount string `json:"service_account" mapstructure:"service_account" tfsdk:"service_account"`
	Location       string `json:"location" mapstructure:"location" tfsdk:"location"`
}

type BingAdsConnectionConfiguration struct{}

type GCSConfiguration struct {
	ProjectId      string `json:"project_id" mapstructure:"project_id" tfsdk:"project_id"`
	ServiceAccount string `json:"service_account" mapstructure:"service_account" tfsdk:"service_account"`
	Bucket         string `json:"bucket" mapstructure:"bucket" tfsdk:"bucket"`
}

type AzureBlobConfiguration struct {
	AccountName   string `json:"account_name" mapstructure:"account_name" tfsdk:"account_name"`
	AccessKey     string `json:"access_key" mapstructure:"access_key" tfsdk:"access_key"`
	ContainerName string `json:"container_name" mapstructure:"container_name" tfsdk:"container_name"`
}

type S3Configuration struct {
	AwsAccessKeyID     string `json:"aws_access_key_id" mapstructure:"aws_access_key_id" tfsdk:"aws_access_key_id"`
	AwsSecretAccessKey string `json:"aws_secret_access_key" mapstructure:"aws_secret_access_key" tfsdk:"aws_secret_access_key"`
	S3BucketRegion     string `json:"s3_bucket_region" mapstructure:"s3_bucket_region" tfsdk:"s3_bucket_region"`
	S3BucketName       string `json:"s3_bucket_name" mapstructure:"s3_bucket_name" tfsdk:"s3_bucket_name"`
}

type AthenaConfiguration struct {
	AccessKeyID     string `json:"access_id" mapstructure:"access_id" tfsdk:"access_id"`
	AccessKeySecret string `json:"secret_access_key" mapstructure:"secret_access_key"` //nolint:gosec
	Region          string `json:"region" mapstructure:"region" tfsdk:"region"`
	OutputBucket    string `json:"outputbucket" mapstructure:"outputbucket" tfsdk:"outputbucket"`
}

type SQLServerConfiguration struct {
	Hostname string `json:"hostname" mapstructure:"hostname" tfsdk:"hostname"`
	Username string `json:"username" mapstructure:"username" tfsdk:"username"`
	Password string `json:"password" mapstructure:"password" tfsdk:"password"`
	Database string `json:"database" mapstructure:"database" tfsdk:"database"`
	Port     int    `json:"port" mapstructure:"port" tfsdk:"port"`
}

type PostgresqlConfiguration struct {
	Hostname string `json:"hostname" mapstructure:"hostname" tfsdk:"hostname"`
	Username string `json:"username" mapstructure:"username" tfsdk:"username"`
	Password string `json:"password" mapstructure:"password" tfsdk:"password"`
	Database string `json:"database" mapstructure:"database" tfsdk:"database"`
	Port     int    `json:"port" mapstructure:"port" tfsdk:"port"`
	SSL      bool   `json:"ssl" mapstructure:"ssl" tfsdk:"ssl"`

	// Client Certificates
	ClientCerts       bool   `json:"client_certs" mapstructure:"client_certs" tfsdk:"client_certs"`
	ClientCertificate string `json:"client_certificate" mapstructure:"client_certificate" tfsdk:"client_certificate"`
	ClientKey         string `json:"client_key" mapstructure:"client_key" tfsdk:"client_key"`
	CACert            string `json:"ca_cert" mapstructure:"ca_cert" tfsdk:"ca_cert"`

	// Change Detection
	ChangeDetection bool   `json:"change_detection" mapstructure:"change_detection" tfsdk:"change_detection"`
	Publication     string `json:"publication" mapstructure:"publication" tfsdk:"publication"`

	SSH           bool   `json:"ssh" mapstructure:"ssh" tfsdk:"ssh"`
	SSHUser       string `json:"ssh_user,omitempty" mapstructure:"ssh_user" tfsdk:"ssh_user"`
	SSHHost       string `json:"ssh_host,omitempty" mapstructure:"ssh_host" tfsdk:"ssh_host"`
	SSHPort       int    `json:"ssh_port,omitempty" mapstructure:"ssh_port" tfsdk:"ssh_port"`
	SSHPrivateKey string `json:"ssh_private_key,omitempty" mapstructure:"ssh_private_key" tfsdk:"ssh_private_key"`
}

type MongoDBConfiguration struct {
	Hosts    string `json:"hosts" mapstructure:"hosts" tfsdk:"hosts"`
	Username string `json:"username" mapstructure:"username" tfsdk:"username"`
	Password string `json:"password" mapstructure:"password" tfsdk:"password"`
	Database string `json:"database" mapstructure:"database" tfsdk:"database"`
	SRV      bool   `json:"srv" mapstructure:"srv" tfsdk:"srv"`
	Params   string `json:"params" mapstructure:"params" tfsdk:"params"`
}

type MarketoConfiguration struct {
	ClientID     string `json:"client_id" mapstructure:"client_id" tfsdk:"client_id"`
	ClientSecret string `json:"client_secret" mapstructure:"client_secret" tfsdk:"client_secret"`
	RESTEndpoint string `json:"rest_endpoint" mapstructure:"rest_endpoint" tfsdk:"rest_endpoint"`

	EnforceAPILimits  bool `json:"enforce_api_limits" mapstructure:"enforce_api_limits" tfsdk:"enforce_api_limits"`
	DailyAPICalls     int  `json:"daily_api_calls" mapstructure:"daily_api_calls" tfsdk:"daily_api_calls"`
	ConcurrentImports int  `json:"concurrent_imports" mapstructure:"concurrent_imports" tfsdk:"concurrent_imports"`
}

type AffinityConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type AirtableConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type AmplitudeConfiguration struct {
	APIKey    string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
	SecretKey string `json:"secret_key" mapstructure:"secret_key" tfsdk:"secret_key"`
}

type APIConnectionConfiguration struct {
	URL         string             `json:"url" mapstructure:"url" tfsdk:"url"`
	Headers     []RequestParameter `json:"headers" mapstructure:"headers" tfsdk:"headers"`
	Body        string             `json:"body" mapstructure:"body" tfsdk:"body"`
	Parameters  []RequestParameter `json:"parameters" mapstructure:"parameters" tfsdk:"parameters"`
	Healthcheck string             `json:"healthcheck" mapstructure:"healthcheck" tfsdk:"healthcheck"`
	Auth        Auth               `json:"auth" mapstructure:",squash" tfsdk:"auth"`
}

type RequestParameter struct {
	Name  string `json:"name" mapstructure:"name" tfsdk:"name"`
	Value string `json:"value" mapstructure:"value" tfsdk:"value"`
}

type Auth struct {
	// BasicAuthConf provides basic authentication credentials
	// e.g.
	// Authorization: Basic <base64 encoded username:password>
	// RFC7617
	Basic *BasicAuthConf `json:"basic,omitempty" mapstructure:"basic" tfsdk:"basic"`
	// HeaderAuthConf provides header authentication credentials
	// e.g.
	// Authorization: Bearer <token>
	Header *RequestParameter `json:"header,omitempty" mapstructure:"header" tfsdk:"header"`
	// OAuthConf provides OAuth authentication using the client credentials flow
	// e.g.
	// Client ID: <client id>
	// Client Secret: <client secret>
	// Callback URL: <callback url>
	// RFC6749 Section-4.4
	OAuth *ClientCredentialConf `json:"oauth,omitempty" mapstructure:"oauth" tfsdk:"oauth"`
}

type BasicAuthConf struct {
	Username string `json:"username,omitempty" mapstructure:"username" tfsdk:"username"`
	Password string `json:"password,omitempty" mapstructure:"password" tfsdk:"password"`
}

type ClientCredentialConf struct {
	ClientID      string             `json:"client_id" mapstructure:"client_id" tfsdk:"client_id"`
	ClientSecret  string             `json:"client_secret" mapstructure:"client_secret" tfsdk:"client_secret"`
	TokenEndpoint string             `json:"token_endpoint" mapstructure:"token_endpoint" tfsdk:"token_endpoint"`
	ExtraFormData []RequestParameter `json:"extra_form_data" mapstructure:"extra_form_data" tfsdk:"extra_form_data"`
}

type CSVConnectionConfiguration struct {
	URL                   string             `json:"url" mapstructure:"url" tfsdk:"url"`
	Headers               []RequestParameter `json:"headers" mapstructure:"headers" tfsdk:"headers"`
	QueryStringParameters []RequestParameter `json:"parameters" mapstructure:"parameters" tfsdk:"parameters"`
	Auth                  Auth               `json:"auth" mapstructure:",squash" tfsdk:"auth"`
}

type ChargeBeeConnectionConfiguration struct {
	Site         string `json:"site" mapstructure:"site" tfsdk:"site"`
	APIKey       string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
	RatelimitRPM int    `json:"ratelimit_rpm" mapstructure:"ratelimit_rpm" tfsdk:"ratelimit_rpm"`
}

type CloudSQLConnectionConfiguration struct {
	ConnectionName string `json:"connection_name" mapstructure:"connection_name" tfsdk:"connection_name"`
	Database       string `json:"database" mapstructure:"database" tfsdk:"database"`
	Username       string `json:"username" mapstructure:"username" tfsdk:"username"`
	Password       string `json:"password" mapstructure:"password" tfsdk:"password"`
	Credentials    string `json:"credentials" mapstructure:"credentials" tfsdk:"credentials"`
}

type CosmosDBConnectionConfiguration struct {
	URI string `json:"uri" mapstructure:"uri" tfsdk:"uri"`
	Key string `json:"key" mapstructure:"key" tfsdk:"key"`
}

type CustomerIOConnectionConfiguration struct {
	SiteID         string `json:"site_id" mapstructure:"site_id" tfsdk:"site_id"`
	TrackingAPIKey string `json:"tracking_api_key" mapstructure:"tracking_api_key" tfsdk:"tracking_api_key"`
	AppAPIKey      string `json:"app_api_key" mapstructure:"app_api_key" tfsdk:"app_api_key"`
}

type DialpadConnectionConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type FreshdeskConnectionConfiguration struct {
	Subdomain string `json:"subdomain" mapstructure:"subdomain" tfsdk:"subdomain"`
	Apikey    string `json:"apikey" mapstructure:"apikey" tfsdk:"apikey"`
}

type FullstoryConnectionConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type HarmonicConnectionConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type IntercomConnectionConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type IterableConnectionConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type KlaviyoConnectionConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type KustomerConnectionConfiguration struct {
	Domain string `json:"domain" mapstructure:"domain" tfsdk:"domain"`
	Apikey string `json:"apikey" mapstructure:"apikey" tfsdk:"apikey"`
}

type LobConnectionConfiguration struct {
	Apikey string `json:"apikey" mapstructure:"apikey" tfsdk:"apikey"`
}

type MysqlConnectionConfiguration struct {
	Hostname string `json:"hostname" mapstructure:"hostname" tfsdk:"hostname"`
	Account  string `json:"account" mapstructure:"account" tfsdk:"account"`
	Passwd   string `json:"passwd" mapstructure:"passwd" tfsdk:"passwd"`
	Dbname   string `json:"dbname" mapstructure:"dbname" tfsdk:"dbname"`
	Port     int    `json:"port" mapstructure:"port" tfsdk:"port"`

	SSH           bool   `json:"ssh" mapstructure:"ssh" tfsdk:"ssh"`
	SSHUser       string `json:"ssh_user" mapstructure:"ssh_user" tfsdk:"ssh_user"`
	SSHHost       string `json:"ssh_host" mapstructure:"ssh_host" tfsdk:"ssh_host"`
	SSHPort       int    `json:"ssh_port" mapstructure:"ssh_port" tfsdk:"ssh_port"`
	SSHPrivateKey string `json:"ssh_private_key" mapstructure:"ssh_private_key" tfsdk:"ssh_private_key"`
}

type NetsuiteConnectionConfiguration struct {
	AccountID      string `json:"account_id" mapstructure:"account_id" tfsdk:"account_id"`
	ConsumerKey    string `json:"consumer_key" mapstructure:"consumer_key" tfsdk:"consumer_key"`
	ConsumerSecret string `json:"consumer_secret" mapstructure:"consumer_secret" tfsdk:"consumer_secret"`
	Token          string `json:"token" mapstructure:"token" tfsdk:"token"`
	TokenSecret    string `json:"token_secret" mapstructure:"token_secret" tfsdk:"token_secret"`
}

type PipedriveConnectionConfiguration struct {
	Domain string `json:"domain" mapstructure:"domain" tfsdk:"domain"`
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type RedshiftConnectionConfiguration struct {
	Hostname string `json:"hostname" mapstructure:"hostname" tfsdk:"hostname"`
	Username string `json:"username" mapstructure:"username" tfsdk:"username"`
	Password string `json:"password" mapstructure:"password" tfsdk:"password"`
	Database string `json:"database" mapstructure:"database" tfsdk:"database"`
	Port     int    `json:"port" mapstructure:"port" tfsdk:"port"`

	SSH           bool   `json:"ssh" mapstructure:"ssh" tfsdk:"ssh"`
	SSHUser       string `json:"ssh_user" mapstructure:"ssh_user" tfsdk:"ssh_user"`
	SSHHost       string `json:"ssh_host" mapstructure:"ssh_host" tfsdk:"ssh_host"`
	SSHPort       int    `json:"ssh_port" mapstructure:"ssh_port" tfsdk:"ssh_port"`
	SSHPrivateKey string `json:"ssh_private_key" mapstructure:"ssh_private_key" tfsdk:"ssh_private_key"`

	AwsAccessKeyID     string `json:"aws_access_key_id" mapstructure:"aws_access_key_id" tfsdk:"aws_access_key_id"`
	AwsSecretAccessKey string `json:"aws_secret_access_key" mapstructure:"aws_secret_access_key" tfsdk:"aws_secret_access_key"`
	S3BucketName       string `json:"s3_bucket_name" mapstructure:"s3_bucket_name" tfsdk:"s3_bucket_name"`
	S3BucketRegion     string `json:"s3_bucket_region" mapstructure:"s3_bucket_region" tfsdk:"s3_bucket_region"`
}

type SegmentConnectionConfiguration struct {
	WriteKey string `json:"write_key" mapstructure:"write_key" tfsdk:"write_key"`
}

type ShipBobConnectionConfiguration struct{}

type ShopifyConnectionConfiguration struct {
	Store string `json:"store" mapstructure:"store" tfsdk:"store"`
}

type SmartSheetConnectionConfiguration struct{}

type StripeConnectionConfiguration struct {
	APIKey  string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
	Version string `json:"version" mapstructure:"version" tfsdk:"version"`
}

type SurvicateConnectionConfiguration struct {
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type SynapseConnectionConfiguration struct {
	Hostname string `json:"hostname" mapstructure:"hostname" tfsdk:"hostname"`
	Username string `json:"username" mapstructure:"username" tfsdk:"username"`
	Password string `json:"password" mapstructure:"password" tfsdk:"password"`
	Database string `json:"database" mapstructure:"database" tfsdk:"database"`
	Port     int    `json:"port" mapstructure:"port" tfsdk:"port"`
}

type UserVoiceConnectionConfiguration struct {
	Domain string `json:"domain" mapstructure:"domain" tfsdk:"domain"`
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type VanillaConnectionConfiguration struct {
	Domain string `json:"domain" mapstructure:"domain" tfsdk:"domain"`
	APIKey string `json:"api_key" mapstructure:"api_key" tfsdk:"api_key"`
}

type WebhookConnectionConfiguration struct {
	URL     string             `json:"url" mapstructure:"url" tfsdk:"url"`
	Secret  string             `json:"secret" mapstructure:"secret" tfsdk:"secret"`
	Headers []RequestParameter `json:"headers" mapstructure:"headers" tfsdk:"headers"`
}

type ZendeskConnectionConfiguration struct {
	Domain string `json:"domain" mapstructure:"domain" tfsdk:"domain"`
}
