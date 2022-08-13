// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-vault-go/vault/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-vault-go/vault/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount vault_database_secrets_mount}.
type DatabaseSecretsMount interface {
	cdktf.TerraformResource
	Accessor() *string
	AuditNonHmacRequestKeys() *[]*string
	SetAuditNonHmacRequestKeys(val *[]*string)
	AuditNonHmacRequestKeysInput() *[]*string
	AuditNonHmacResponseKeys() *[]*string
	SetAuditNonHmacResponseKeys(val *[]*string)
	AuditNonHmacResponseKeysInput() *[]*string
	Cassandra() DatabaseSecretsMountCassandraList
	CassandraInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Couchbase() DatabaseSecretsMountCouchbaseList
	CouchbaseInput() interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	DefaultLeaseTtlSeconds() *float64
	SetDefaultLeaseTtlSeconds(val *float64)
	DefaultLeaseTtlSecondsInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	Elasticsearch() DatabaseSecretsMountElasticsearchList
	ElasticsearchInput() interface{}
	EngineCount() *float64
	ExternalEntropyAccess() interface{}
	SetExternalEntropyAccess(val interface{})
	ExternalEntropyAccessInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hana() DatabaseSecretsMountHanaList
	HanaInput() interface{}
	Id() *string
	SetId(val *string)
	IdInput() *string
	Influxdb() DatabaseSecretsMountInfluxdbList
	InfluxdbInput() interface{}
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Local() interface{}
	SetLocal(val interface{})
	LocalInput() interface{}
	MaxLeaseTtlSeconds() *float64
	SetMaxLeaseTtlSeconds(val *float64)
	MaxLeaseTtlSecondsInput() *float64
	Mongodb() DatabaseSecretsMountMongodbList
	Mongodbatlas() DatabaseSecretsMountMongodbatlasList
	MongodbatlasInput() interface{}
	MongodbInput() interface{}
	Mssql() DatabaseSecretsMountMssqlList
	MssqlInput() interface{}
	Mysql() DatabaseSecretsMountMysqlList
	MysqlAurora() DatabaseSecretsMountMysqlAuroraList
	MysqlAuroraInput() interface{}
	MysqlInput() interface{}
	MysqlLegacy() DatabaseSecretsMountMysqlLegacyList
	MysqlLegacyInput() interface{}
	MysqlRds() DatabaseSecretsMountMysqlRdsList
	MysqlRdsInput() interface{}
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Options() *map[string]*string
	SetOptions(val *map[string]*string)
	OptionsInput() *map[string]*string
	Oracle() DatabaseSecretsMountOracleList
	OracleInput() interface{}
	Path() *string
	SetPath(val *string)
	PathInput() *string
	Postgresql() DatabaseSecretsMountPostgresqlList
	PostgresqlInput() interface{}
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Redshift() DatabaseSecretsMountRedshiftList
	RedshiftInput() interface{}
	SealWrap() interface{}
	SetSealWrap(val interface{})
	SealWrapInput() interface{}
	Snowflake() DatabaseSecretsMountSnowflakeList
	SnowflakeInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	PutCassandra(value interface{})
	PutCouchbase(value interface{})
	PutElasticsearch(value interface{})
	PutHana(value interface{})
	PutInfluxdb(value interface{})
	PutMongodb(value interface{})
	PutMongodbatlas(value interface{})
	PutMssql(value interface{})
	PutMysql(value interface{})
	PutMysqlAurora(value interface{})
	PutMysqlLegacy(value interface{})
	PutMysqlRds(value interface{})
	PutOracle(value interface{})
	PutPostgresql(value interface{})
	PutRedshift(value interface{})
	PutSnowflake(value interface{})
	ResetAuditNonHmacRequestKeys()
	ResetAuditNonHmacResponseKeys()
	ResetCassandra()
	ResetCouchbase()
	ResetDefaultLeaseTtlSeconds()
	ResetDescription()
	ResetElasticsearch()
	ResetExternalEntropyAccess()
	ResetHana()
	ResetId()
	ResetInfluxdb()
	ResetLocal()
	ResetMaxLeaseTtlSeconds()
	ResetMongodb()
	ResetMongodbatlas()
	ResetMssql()
	ResetMysql()
	ResetMysqlAurora()
	ResetMysqlLegacy()
	ResetMysqlRds()
	ResetNamespace()
	ResetOptions()
	ResetOracle()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPostgresql()
	ResetRedshift()
	ResetSealWrap()
	ResetSnowflake()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DatabaseSecretsMount
type jsiiProxy_DatabaseSecretsMount struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabaseSecretsMount) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) AuditNonHmacRequestKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) AuditNonHmacRequestKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacRequestKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) AuditNonHmacResponseKeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) AuditNonHmacResponseKeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"auditNonHmacResponseKeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Cassandra() DatabaseSecretsMountCassandraList {
	var returns DatabaseSecretsMountCassandraList
	_jsii_.Get(
		j,
		"cassandra",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) CassandraInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"cassandraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Couchbase() DatabaseSecretsMountCouchbaseList {
	var returns DatabaseSecretsMountCouchbaseList
	_jsii_.Get(
		j,
		"couchbase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) CouchbaseInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"couchbaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) DefaultLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) DefaultLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Elasticsearch() DatabaseSecretsMountElasticsearchList {
	var returns DatabaseSecretsMountElasticsearchList
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) ElasticsearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) EngineCount() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"engineCount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) ExternalEntropyAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEntropyAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) ExternalEntropyAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalEntropyAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Hana() DatabaseSecretsMountHanaList {
	var returns DatabaseSecretsMountHanaList
	_jsii_.Get(
		j,
		"hana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) HanaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"hanaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Influxdb() DatabaseSecretsMountInfluxdbList {
	var returns DatabaseSecretsMountInfluxdbList
	_jsii_.Get(
		j,
		"influxdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) InfluxdbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"influxdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Local() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) LocalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MaxLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MaxLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Mongodb() DatabaseSecretsMountMongodbList {
	var returns DatabaseSecretsMountMongodbList
	_jsii_.Get(
		j,
		"mongodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Mongodbatlas() DatabaseSecretsMountMongodbatlasList {
	var returns DatabaseSecretsMountMongodbatlasList
	_jsii_.Get(
		j,
		"mongodbatlas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MongodbatlasInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongodbatlasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MongodbInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mongodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Mssql() DatabaseSecretsMountMssqlList {
	var returns DatabaseSecretsMountMssqlList
	_jsii_.Get(
		j,
		"mssql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MssqlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mssqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Mysql() DatabaseSecretsMountMysqlList {
	var returns DatabaseSecretsMountMysqlList
	_jsii_.Get(
		j,
		"mysql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MysqlAurora() DatabaseSecretsMountMysqlAuroraList {
	var returns DatabaseSecretsMountMysqlAuroraList
	_jsii_.Get(
		j,
		"mysqlAurora",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MysqlAuroraInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mysqlAuroraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MysqlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mysqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MysqlLegacy() DatabaseSecretsMountMysqlLegacyList {
	var returns DatabaseSecretsMountMysqlLegacyList
	_jsii_.Get(
		j,
		"mysqlLegacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MysqlLegacyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mysqlLegacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MysqlRds() DatabaseSecretsMountMysqlRdsList {
	var returns DatabaseSecretsMountMysqlRdsList
	_jsii_.Get(
		j,
		"mysqlRds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) MysqlRdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"mysqlRdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Options() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"options",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) OptionsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"optionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Oracle() DatabaseSecretsMountOracleList {
	var returns DatabaseSecretsMountOracleList
	_jsii_.Get(
		j,
		"oracle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) OracleInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"oracleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Postgresql() DatabaseSecretsMountPostgresqlList {
	var returns DatabaseSecretsMountPostgresqlList
	_jsii_.Get(
		j,
		"postgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) PostgresqlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"postgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Redshift() DatabaseSecretsMountRedshiftList {
	var returns DatabaseSecretsMountRedshiftList
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) RedshiftInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) SealWrap() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sealWrap",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) SealWrapInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"sealWrapInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) Snowflake() DatabaseSecretsMountSnowflakeList {
	var returns DatabaseSecretsMountSnowflakeList
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) SnowflakeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMount) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount vault_database_secrets_mount} Resource.
func NewDatabaseSecretsMount(scope constructs.Construct, id *string, config *DatabaseSecretsMountConfig) DatabaseSecretsMount {
	_init_.Initialize()

	j := jsiiProxy_DatabaseSecretsMount{}

	_jsii_.Create(
		"@cdktf/provider-vault.DatabaseSecretsMount",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/vault/r/database_secrets_mount vault_database_secrets_mount} Resource.
func NewDatabaseSecretsMount_Override(d DatabaseSecretsMount, scope constructs.Construct, id *string, config *DatabaseSecretsMountConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.DatabaseSecretsMount",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetAuditNonHmacRequestKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"auditNonHmacRequestKeys",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetAuditNonHmacResponseKeys(val *[]*string) {
	_jsii_.Set(
		j,
		"auditNonHmacResponseKeys",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetDefaultLeaseTtlSeconds(val *float64) {
	_jsii_.Set(
		j,
		"defaultLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetExternalEntropyAccess(val interface{}) {
	_jsii_.Set(
		j,
		"externalEntropyAccess",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetLocal(val interface{}) {
	_jsii_.Set(
		j,
		"local",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetMaxLeaseTtlSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetOptions(val *map[string]*string) {
	_jsii_.Set(
		j,
		"options",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetPath(val *string) {
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMount) SetSealWrap(val interface{}) {
	_jsii_.Set(
		j,
		"sealWrap",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func DatabaseSecretsMount_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.DatabaseSecretsMount",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseSecretsMount_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.DatabaseSecretsMount",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutCassandra(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putCassandra",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutCouchbase(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putCouchbase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutElasticsearch(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutHana(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putHana",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutInfluxdb(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putInfluxdb",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutMongodb(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putMongodb",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutMongodbatlas(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putMongodbatlas",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutMssql(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putMssql",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutMysql(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putMysql",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutMysqlAurora(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putMysqlAurora",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutMysqlLegacy(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putMysqlLegacy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutMysqlRds(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putMysqlRds",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutOracle(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putOracle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutPostgresql(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putPostgresql",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutRedshift(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putRedshift",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) PutSnowflake(value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetAuditNonHmacRequestKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetAuditNonHmacRequestKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetAuditNonHmacResponseKeys() {
	_jsii_.InvokeVoid(
		d,
		"resetAuditNonHmacResponseKeys",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetCassandra() {
	_jsii_.InvokeVoid(
		d,
		"resetCassandra",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetCouchbase() {
	_jsii_.InvokeVoid(
		d,
		"resetCouchbase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetDefaultLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetDefaultLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetDescription() {
	_jsii_.InvokeVoid(
		d,
		"resetDescription",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		d,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetExternalEntropyAccess() {
	_jsii_.InvokeVoid(
		d,
		"resetExternalEntropyAccess",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetHana() {
	_jsii_.InvokeVoid(
		d,
		"resetHana",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetInfluxdb() {
	_jsii_.InvokeVoid(
		d,
		"resetInfluxdb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetLocal() {
	_jsii_.InvokeVoid(
		d,
		"resetLocal",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetMaxLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetMongodb() {
	_jsii_.InvokeVoid(
		d,
		"resetMongodb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetMongodbatlas() {
	_jsii_.InvokeVoid(
		d,
		"resetMongodbatlas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetMssql() {
	_jsii_.InvokeVoid(
		d,
		"resetMssql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetMysql() {
	_jsii_.InvokeVoid(
		d,
		"resetMysql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetMysqlAurora() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlAurora",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetMysqlLegacy() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlLegacy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetMysqlRds() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlRds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetOptions() {
	_jsii_.InvokeVoid(
		d,
		"resetOptions",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetOracle() {
	_jsii_.InvokeVoid(
		d,
		"resetOracle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetPostgresql() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetRedshift() {
	_jsii_.InvokeVoid(
		d,
		"resetRedshift",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetSealWrap() {
	_jsii_.InvokeVoid(
		d,
		"resetSealWrap",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) ResetSnowflake() {
	_jsii_.InvokeVoid(
		d,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMount) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMount) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

