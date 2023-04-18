package databasesecretsmount

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v6/jsii"

	"github.com/cdktf/cdktf-provider-vault-go/vault/v6/databasesecretsmount/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseSecretsMountOracleOutputReference interface {
	cdktf.ComplexObject
	AllowedRoles() *[]*string
	SetAllowedRoles(val *[]*string)
	AllowedRolesInput() *[]*string
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	ConnectionUrl() *string
	SetConnectionUrl(val *string)
	ConnectionUrlInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	Data() *map[string]*string
	SetData(val *map[string]*string)
	DataInput() *map[string]*string
	// Experimental.
	Fqn() *string
	InternalValue() interface{}
	SetInternalValue(val interface{})
	MaxConnectionLifetime() *float64
	SetMaxConnectionLifetime(val *float64)
	MaxConnectionLifetimeInput() *float64
	MaxIdleConnections() *float64
	SetMaxIdleConnections(val *float64)
	MaxIdleConnectionsInput() *float64
	MaxOpenConnections() *float64
	SetMaxOpenConnections(val *float64)
	MaxOpenConnectionsInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	PluginName() *string
	SetPluginName(val *string)
	PluginNameInput() *string
	RootRotationStatements() *[]*string
	SetRootRotationStatements(val *[]*string)
	RootRotationStatementsInput() *[]*string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	UsernameTemplate() *string
	SetUsernameTemplate(val *string)
	UsernameTemplateInput() *string
	VerifyConnection() interface{}
	SetVerifyConnection(val interface{})
	VerifyConnectionInput() interface{}
	// Experimental.
	ComputeFqn() *string
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
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetAllowedRoles()
	ResetConnectionUrl()
	ResetData()
	ResetMaxConnectionLifetime()
	ResetMaxIdleConnections()
	ResetMaxOpenConnections()
	ResetPassword()
	ResetPluginName()
	ResetRootRotationStatements()
	ResetUsername()
	ResetUsernameTemplate()
	ResetVerifyConnection()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseSecretsMountOracleOutputReference
type jsiiProxy_DatabaseSecretsMountOracleOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) AllowedRoles() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) AllowedRolesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ConnectionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ConnectionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) Data() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"data",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) DataInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"dataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) MaxConnectionLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) MaxConnectionLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) MaxIdleConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) MaxIdleConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) MaxOpenConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpenConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) MaxOpenConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpenConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) PluginName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) PluginNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pluginNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) RootRotationStatements() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootRotationStatements",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) RootRotationStatementsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"rootRotationStatementsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) VerifyConnection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference) VerifyConnectionInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyConnectionInput",
		&returns,
	)
	return returns
}


func NewDatabaseSecretsMountOracleOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DatabaseSecretsMountOracleOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseSecretsMountOracleOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSecretsMountOracleOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vault.databaseSecretsMount.DatabaseSecretsMountOracleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDatabaseSecretsMountOracleOutputReference_Override(d DatabaseSecretsMountOracleOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.databaseSecretsMount.DatabaseSecretsMountOracleOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetAllowedRoles(val *[]*string) {
	if err := j.validateSetAllowedRolesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedRoles",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetConnectionUrl(val *string) {
	if err := j.validateSetConnectionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionUrl",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetData(val *map[string]*string) {
	if err := j.validateSetDataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"data",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetMaxConnectionLifetime(val *float64) {
	if err := j.validateSetMaxConnectionLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionLifetime",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetMaxIdleConnections(val *float64) {
	if err := j.validateSetMaxIdleConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleConnections",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetMaxOpenConnections(val *float64) {
	if err := j.validateSetMaxOpenConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxOpenConnections",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetPluginName(val *string) {
	if err := j.validateSetPluginNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pluginName",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetRootRotationStatements(val *[]*string) {
	if err := j.validateSetRootRotationStatementsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rootRotationStatements",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretsMountOracleOutputReference)SetVerifyConnection(val interface{}) {
	if err := j.validateSetVerifyConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verifyConnection",
		val,
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetAllowedRoles() {
	_jsii_.InvokeVoid(
		d,
		"resetAllowedRoles",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetConnectionUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetData() {
	_jsii_.InvokeVoid(
		d,
		"resetData",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetMaxConnectionLifetime() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxConnectionLifetime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetMaxIdleConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxIdleConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetMaxOpenConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxOpenConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetPluginName() {
	_jsii_.InvokeVoid(
		d,
		"resetPluginName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetRootRotationStatements() {
	_jsii_.InvokeVoid(
		d,
		"resetRootRotationStatements",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ResetVerifyConnection() {
	_jsii_.InvokeVoid(
		d,
		"resetVerifyConnection",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		d,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretsMountOracleOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

