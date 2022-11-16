package databasesecretbackendconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v4/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v4/databasesecretbackendconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_connection vault_database_secret_backend_connection}.
type DatabaseSecretBackendConnection interface {
	cdktf.TerraformResource
	AllowedRoles() *[]*string
	SetAllowedRoles(val *[]*string)
	AllowedRolesInput() *[]*string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	Cassandra() DatabaseSecretBackendConnectionCassandraOutputReference
	CassandraInput() *DatabaseSecretBackendConnectionCassandra
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	Couchbase() DatabaseSecretBackendConnectionCouchbaseOutputReference
	CouchbaseInput() *DatabaseSecretBackendConnectionCouchbase
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	Data() *map[string]*string
	SetData(val *map[string]*string)
	DataInput() *map[string]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Elasticsearch() DatabaseSecretBackendConnectionElasticsearchOutputReference
	ElasticsearchInput() *DatabaseSecretBackendConnectionElasticsearch
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Hana() DatabaseSecretBackendConnectionHanaOutputReference
	HanaInput() *DatabaseSecretBackendConnectionHana
	Id() *string
	SetId(val *string)
	IdInput() *string
	Influxdb() DatabaseSecretBackendConnectionInfluxdbOutputReference
	InfluxdbInput() *DatabaseSecretBackendConnectionInfluxdb
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Mongodb() DatabaseSecretBackendConnectionMongodbOutputReference
	Mongodbatlas() DatabaseSecretBackendConnectionMongodbatlasOutputReference
	MongodbatlasInput() *DatabaseSecretBackendConnectionMongodbatlas
	MongodbInput() *DatabaseSecretBackendConnectionMongodb
	Mssql() DatabaseSecretBackendConnectionMssqlOutputReference
	MssqlInput() *DatabaseSecretBackendConnectionMssql
	Mysql() DatabaseSecretBackendConnectionMysqlOutputReference
	MysqlAurora() DatabaseSecretBackendConnectionMysqlAuroraOutputReference
	MysqlAuroraInput() *DatabaseSecretBackendConnectionMysqlAurora
	MysqlInput() *DatabaseSecretBackendConnectionMysql
	MysqlLegacy() DatabaseSecretBackendConnectionMysqlLegacyOutputReference
	MysqlLegacyInput() *DatabaseSecretBackendConnectionMysqlLegacy
	MysqlRds() DatabaseSecretBackendConnectionMysqlRdsOutputReference
	MysqlRdsInput() *DatabaseSecretBackendConnectionMysqlRds
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Oracle() DatabaseSecretBackendConnectionOracleOutputReference
	OracleInput() *DatabaseSecretBackendConnectionOracle
	PluginName() *string
	SetPluginName(val *string)
	PluginNameInput() *string
	Postgresql() DatabaseSecretBackendConnectionPostgresqlOutputReference
	PostgresqlInput() *DatabaseSecretBackendConnectionPostgresql
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
	RedisElasticache() DatabaseSecretBackendConnectionRedisElasticacheOutputReference
	RedisElasticacheInput() *DatabaseSecretBackendConnectionRedisElasticache
	Redshift() DatabaseSecretBackendConnectionRedshiftOutputReference
	RedshiftInput() *DatabaseSecretBackendConnectionRedshift
	RootRotationStatements() *[]*string
	SetRootRotationStatements(val *[]*string)
	RootRotationStatementsInput() *[]*string
	Snowflake() DatabaseSecretBackendConnectionSnowflakeOutputReference
	SnowflakeInput() *DatabaseSecretBackendConnectionSnowflake
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	VerifyConnection() interface{}
	SetVerifyConnection(val interface{})
	VerifyConnectionInput() interface{}
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
	PutCassandra(value *DatabaseSecretBackendConnectionCassandra)
	PutCouchbase(value *DatabaseSecretBackendConnectionCouchbase)
	PutElasticsearch(value *DatabaseSecretBackendConnectionElasticsearch)
	PutHana(value *DatabaseSecretBackendConnectionHana)
	PutInfluxdb(value *DatabaseSecretBackendConnectionInfluxdb)
	PutMongodb(value *DatabaseSecretBackendConnectionMongodb)
	PutMongodbatlas(value *DatabaseSecretBackendConnectionMongodbatlas)
	PutMssql(value *DatabaseSecretBackendConnectionMssql)
	PutMysql(value *DatabaseSecretBackendConnectionMysql)
	PutMysqlAurora(value *DatabaseSecretBackendConnectionMysqlAurora)
	PutMysqlLegacy(value *DatabaseSecretBackendConnectionMysqlLegacy)
	PutMysqlRds(value *DatabaseSecretBackendConnectionMysqlRds)
	PutOracle(value *DatabaseSecretBackendConnectionOracle)
	PutPostgresql(value *DatabaseSecretBackendConnectionPostgresql)
	PutRedisElasticache(value *DatabaseSecretBackendConnectionRedisElasticache)
	PutRedshift(value *DatabaseSecretBackendConnectionRedshift)
	PutSnowflake(value *DatabaseSecretBackendConnectionSnowflake)
	ResetAllowedRoles()
	ResetCassandra()
	ResetCouchbase()
	ResetData()
	ResetElasticsearch()
	ResetHana()
	ResetId()
	ResetInfluxdb()
	ResetMongodb()
	ResetMongodbatlas()
	ResetMssql()
	ResetMysql()
	ResetMysqlAurora()
	ResetMysqlLegacy()
	ResetMysqlRds()
	ResetNamespace()
	ResetOracle()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPluginName()
	ResetPostgresql()
	ResetRedisElasticache()
	ResetRedshift()
	ResetRootRotationStatements()
	ResetSnowflake()
	ResetVerifyConnection()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DatabaseSecretBackendConnection
type jsiiProxy_DatabaseSecretBackendConnection struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) AllowedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) AllowedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Cassandra() DatabaseSecretBackendConnectionCassandraOutputReference {
	var returns DatabaseSecretBackendConnectionCassandraOutputReference
	_jsii_.Get(
		j,
		"cassandra",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) CassandraInput() *DatabaseSecretBackendConnectionCassandra {
	var returns *DatabaseSecretBackendConnectionCassandra
	_jsii_.Get(
		j,
		"cassandraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Couchbase() DatabaseSecretBackendConnectionCouchbaseOutputReference {
	var returns DatabaseSecretBackendConnectionCouchbaseOutputReference
	_jsii_.Get(
		j,
		"couchbase",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) CouchbaseInput() *DatabaseSecretBackendConnectionCouchbase {
	var returns *DatabaseSecretBackendConnectionCouchbase
	_jsii_.Get(
		j,
		"couchbaseInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Data() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) DataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Elasticsearch() DatabaseSecretBackendConnectionElasticsearchOutputReference {
	var returns DatabaseSecretBackendConnectionElasticsearchOutputReference
	_jsii_.Get(
		j,
		"elasticsearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) ElasticsearchInput() *DatabaseSecretBackendConnectionElasticsearch {
	var returns *DatabaseSecretBackendConnectionElasticsearch
	_jsii_.Get(
		j,
		"elasticsearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Hana() DatabaseSecretBackendConnectionHanaOutputReference {
	var returns DatabaseSecretBackendConnectionHanaOutputReference
	_jsii_.Get(
		j,
		"hana",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) HanaInput() *DatabaseSecretBackendConnectionHana {
	var returns *DatabaseSecretBackendConnectionHana
	_jsii_.Get(
		j,
		"hanaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Influxdb() DatabaseSecretBackendConnectionInfluxdbOutputReference {
	var returns DatabaseSecretBackendConnectionInfluxdbOutputReference
	_jsii_.Get(
		j,
		"influxdb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) InfluxdbInput() *DatabaseSecretBackendConnectionInfluxdb {
	var returns *DatabaseSecretBackendConnectionInfluxdb
	_jsii_.Get(
		j,
		"influxdbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Mongodb() DatabaseSecretBackendConnectionMongodbOutputReference {
	var returns DatabaseSecretBackendConnectionMongodbOutputReference
	_jsii_.Get(
		j,
		"mongodb",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Mongodbatlas() DatabaseSecretBackendConnectionMongodbatlasOutputReference {
	var returns DatabaseSecretBackendConnectionMongodbatlasOutputReference
	_jsii_.Get(
		j,
		"mongodbatlas",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) MongodbatlasInput() *DatabaseSecretBackendConnectionMongodbatlas {
	var returns *DatabaseSecretBackendConnectionMongodbatlas
	_jsii_.Get(
		j,
		"mongodbatlasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) MongodbInput() *DatabaseSecretBackendConnectionMongodb {
	var returns *DatabaseSecretBackendConnectionMongodb
	_jsii_.Get(
		j,
		"mongodbInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Mssql() DatabaseSecretBackendConnectionMssqlOutputReference {
	var returns DatabaseSecretBackendConnectionMssqlOutputReference
	_jsii_.Get(
		j,
		"mssql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) MssqlInput() *DatabaseSecretBackendConnectionMssql {
	var returns *DatabaseSecretBackendConnectionMssql
	_jsii_.Get(
		j,
		"mssqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Mysql() DatabaseSecretBackendConnectionMysqlOutputReference {
	var returns DatabaseSecretBackendConnectionMysqlOutputReference
	_jsii_.Get(
		j,
		"mysql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) MysqlAurora() DatabaseSecretBackendConnectionMysqlAuroraOutputReference {
	var returns DatabaseSecretBackendConnectionMysqlAuroraOutputReference
	_jsii_.Get(
		j,
		"mysqlAurora",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) MysqlAuroraInput() *DatabaseSecretBackendConnectionMysqlAurora {
	var returns *DatabaseSecretBackendConnectionMysqlAurora
	_jsii_.Get(
		j,
		"mysqlAuroraInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) MysqlInput() *DatabaseSecretBackendConnectionMysql {
	var returns *DatabaseSecretBackendConnectionMysql
	_jsii_.Get(
		j,
		"mysqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) MysqlLegacy() DatabaseSecretBackendConnectionMysqlLegacyOutputReference {
	var returns DatabaseSecretBackendConnectionMysqlLegacyOutputReference
	_jsii_.Get(
		j,
		"mysqlLegacy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) MysqlLegacyInput() *DatabaseSecretBackendConnectionMysqlLegacy {
	var returns *DatabaseSecretBackendConnectionMysqlLegacy
	_jsii_.Get(
		j,
		"mysqlLegacyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) MysqlRds() DatabaseSecretBackendConnectionMysqlRdsOutputReference {
	var returns DatabaseSecretBackendConnectionMysqlRdsOutputReference
	_jsii_.Get(
		j,
		"mysqlRds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) MysqlRdsInput() *DatabaseSecretBackendConnectionMysqlRds {
	var returns *DatabaseSecretBackendConnectionMysqlRds
	_jsii_.Get(
		j,
		"mysqlRdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Oracle() DatabaseSecretBackendConnectionOracleOutputReference {
	var returns DatabaseSecretBackendConnectionOracleOutputReference
	_jsii_.Get(
		j,
		"oracle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) OracleInput() *DatabaseSecretBackendConnectionOracle {
	var returns *DatabaseSecretBackendConnectionOracle
	_jsii_.Get(
		j,
		"oracleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) PluginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Postgresql() DatabaseSecretBackendConnectionPostgresqlOutputReference {
	var returns DatabaseSecretBackendConnectionPostgresqlOutputReference
	_jsii_.Get(
		j,
		"postgresql",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) PostgresqlInput() *DatabaseSecretBackendConnectionPostgresql {
	var returns *DatabaseSecretBackendConnectionPostgresql
	_jsii_.Get(
		j,
		"postgresqlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) RedisElasticache() DatabaseSecretBackendConnectionRedisElasticacheOutputReference {
	var returns DatabaseSecretBackendConnectionRedisElasticacheOutputReference
	_jsii_.Get(
		j,
		"redisElasticache",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) RedisElasticacheInput() *DatabaseSecretBackendConnectionRedisElasticache {
	var returns *DatabaseSecretBackendConnectionRedisElasticache
	_jsii_.Get(
		j,
		"redisElasticacheInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Redshift() DatabaseSecretBackendConnectionRedshiftOutputReference {
	var returns DatabaseSecretBackendConnectionRedshiftOutputReference
	_jsii_.Get(
		j,
		"redshift",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) RedshiftInput() *DatabaseSecretBackendConnectionRedshift {
	var returns *DatabaseSecretBackendConnectionRedshift
	_jsii_.Get(
		j,
		"redshiftInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) RootRotationStatements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootRotationStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) RootRotationStatementsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootRotationStatementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) Snowflake() DatabaseSecretBackendConnectionSnowflakeOutputReference {
	var returns DatabaseSecretBackendConnectionSnowflakeOutputReference
	_jsii_.Get(
		j,
		"snowflake",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) SnowflakeInput() *DatabaseSecretBackendConnectionSnowflake {
	var returns *DatabaseSecretBackendConnectionSnowflake
	_jsii_.Get(
		j,
		"snowflakeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) VerifyConnection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnection) VerifyConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnectionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_connection vault_database_secret_backend_connection} Resource.
func NewDatabaseSecretBackendConnection(scope constructs.Construct, id *string, config *DatabaseSecretBackendConnectionConfig) DatabaseSecretBackendConnection {
	_init_.Initialize()

	if err := validateNewDatabaseSecretBackendConnectionParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSecretBackendConnection{}

	_jsii_.Create(
		"@cdktf/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnection",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/vault/r/database_secret_backend_connection vault_database_secret_backend_connection} Resource.
func NewDatabaseSecretBackendConnection_Override(d DatabaseSecretBackendConnection, scope constructs.Construct, id *string, config *DatabaseSecretBackendConnectionConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnection",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetAllowedRoles(val *[]*string) {
	if err := j.validateSetAllowedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRoles",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetData(val *map[string]*string) {
	if err := j.validateSetDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetPluginName(val *string) {
	if err := j.validateSetPluginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginName",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetRootRotationStatements(val *[]*string) {
	if err := j.validateSetRootRotationStatementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootRotationStatements",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnection)SetVerifyConnection(val interface{}) {
	if err := j.validateSetVerifyConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyConnection",
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
func DatabaseSecretBackendConnection_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseSecretBackendConnection_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnection",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseSecretBackendConnection_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseSecretBackendConnection_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnection",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func DatabaseSecretBackendConnection_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateDatabaseSecretBackendConnection_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnection",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DatabaseSecretBackendConnection_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnection",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) AddOverride(path *string, value interface{}) {
	if err := d.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) OverrideLogicalId(newLogicalId *string) {
	if err := d.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutCassandra(value *DatabaseSecretBackendConnectionCassandra) {
	if err := d.validatePutCassandraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCassandra",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutCouchbase(value *DatabaseSecretBackendConnectionCouchbase) {
	if err := d.validatePutCouchbaseParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putCouchbase",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutElasticsearch(value *DatabaseSecretBackendConnectionElasticsearch) {
	if err := d.validatePutElasticsearchParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putElasticsearch",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutHana(value *DatabaseSecretBackendConnectionHana) {
	if err := d.validatePutHanaParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putHana",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutInfluxdb(value *DatabaseSecretBackendConnectionInfluxdb) {
	if err := d.validatePutInfluxdbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putInfluxdb",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutMongodb(value *DatabaseSecretBackendConnectionMongodb) {
	if err := d.validatePutMongodbParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMongodb",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutMongodbatlas(value *DatabaseSecretBackendConnectionMongodbatlas) {
	if err := d.validatePutMongodbatlasParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMongodbatlas",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutMssql(value *DatabaseSecretBackendConnectionMssql) {
	if err := d.validatePutMssqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMssql",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutMysql(value *DatabaseSecretBackendConnectionMysql) {
	if err := d.validatePutMysqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMysql",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutMysqlAurora(value *DatabaseSecretBackendConnectionMysqlAurora) {
	if err := d.validatePutMysqlAuroraParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMysqlAurora",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutMysqlLegacy(value *DatabaseSecretBackendConnectionMysqlLegacy) {
	if err := d.validatePutMysqlLegacyParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMysqlLegacy",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutMysqlRds(value *DatabaseSecretBackendConnectionMysqlRds) {
	if err := d.validatePutMysqlRdsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putMysqlRds",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutOracle(value *DatabaseSecretBackendConnectionOracle) {
	if err := d.validatePutOracleParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putOracle",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutPostgresql(value *DatabaseSecretBackendConnectionPostgresql) {
	if err := d.validatePutPostgresqlParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putPostgresql",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutRedisElasticache(value *DatabaseSecretBackendConnectionRedisElasticache) {
	if err := d.validatePutRedisElasticacheParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedisElasticache",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutRedshift(value *DatabaseSecretBackendConnectionRedshift) {
	if err := d.validatePutRedshiftParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putRedshift",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) PutSnowflake(value *DatabaseSecretBackendConnectionSnowflake) {
	if err := d.validatePutSnowflakeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		d,
		"putSnowflake",
		[]interface{}{value},
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetAllowedRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetCassandra() {
	_jsii_.InvokeVoid(
		d,
		"resetCassandra",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetCouchbase() {
	_jsii_.InvokeVoid(
		d,
		"resetCouchbase",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetData() {
	_jsii_.InvokeVoid(
		d,
		"resetData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetElasticsearch() {
	_jsii_.InvokeVoid(
		d,
		"resetElasticsearch",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetHana() {
	_jsii_.InvokeVoid(
		d,
		"resetHana",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetInfluxdb() {
	_jsii_.InvokeVoid(
		d,
		"resetInfluxdb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetMongodb() {
	_jsii_.InvokeVoid(
		d,
		"resetMongodb",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetMongodbatlas() {
	_jsii_.InvokeVoid(
		d,
		"resetMongodbatlas",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetMssql() {
	_jsii_.InvokeVoid(
		d,
		"resetMssql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetMysql() {
	_jsii_.InvokeVoid(
		d,
		"resetMysql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetMysqlAurora() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlAurora",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetMysqlLegacy() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlLegacy",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetMysqlRds() {
	_jsii_.InvokeVoid(
		d,
		"resetMysqlRds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetOracle() {
	_jsii_.InvokeVoid(
		d,
		"resetOracle",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetPluginName() {
	_jsii_.InvokeVoid(
		d,
		"resetPluginName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetPostgresql() {
	_jsii_.InvokeVoid(
		d,
		"resetPostgresql",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetRedisElasticache() {
	_jsii_.InvokeVoid(
		d,
		"resetRedisElasticache",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetRedshift() {
	_jsii_.InvokeVoid(
		d,
		"resetRedshift",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetRootRotationStatements() {
	_jsii_.InvokeVoid(
		d,
		"resetRootRotationStatements",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetSnowflake() {
	_jsii_.InvokeVoid(
		d,
		"resetSnowflake",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ResetVerifyConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetVerifyConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnection) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

