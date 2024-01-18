// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v12/jsii"

	"github.com/cdktf/cdktf-provider-vault-go/vault/v12/databasesecretbackendconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseSecretBackendConnectionMysqlAuroraOutputReference interface {
	cdktf.ComplexObject
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
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
	// Experimental.
	Fqn() *string
	InternalValue() *DatabaseSecretBackendConnectionMysqlAurora
	SetInternalValue(val *DatabaseSecretBackendConnectionMysqlAurora)
	MaxConnectionLifetime() *float64
	SetMaxConnectionLifetime(val *float64)
	MaxConnectionLifetimeInput() *float64
	MaxIdleConnections() *float64
	SetMaxIdleConnections(val *float64)
	MaxIdleConnectionsInput() *float64
	MaxOpenConnections() *float64
	SetMaxOpenConnections(val *float64)
	MaxOpenConnectionsInput() *float64
	Password() *string
	SetPassword(val *string)
	PasswordInput() *string
	ServiceAccountJson() *string
	SetServiceAccountJson(val *string)
	ServiceAccountJsonInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	TlsCa() *string
	SetTlsCa(val *string)
	TlsCaInput() *string
	TlsCertificateKey() *string
	SetTlsCertificateKey(val *string)
	TlsCertificateKeyInput() *string
	Username() *string
	SetUsername(val *string)
	UsernameInput() *string
	UsernameTemplate() *string
	SetUsernameTemplate(val *string)
	UsernameTemplateInput() *string
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
	ResetAuthType()
	ResetConnectionUrl()
	ResetMaxConnectionLifetime()
	ResetMaxIdleConnections()
	ResetMaxOpenConnections()
	ResetPassword()
	ResetServiceAccountJson()
	ResetTlsCa()
	ResetTlsCertificateKey()
	ResetUsername()
	ResetUsernameTemplate()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DatabaseSecretBackendConnectionMysqlAuroraOutputReference
type jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ConnectionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ConnectionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) InternalValue() *DatabaseSecretBackendConnectionMysqlAurora {
	var returns *DatabaseSecretBackendConnectionMysqlAurora
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) MaxConnectionLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) MaxConnectionLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) MaxIdleConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) MaxIdleConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) MaxOpenConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpenConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) MaxOpenConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpenConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ServiceAccountJson() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountJson",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ServiceAccountJsonInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountJsonInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) TlsCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) TlsCaInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) TlsCertificateKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCertificateKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) TlsCertificateKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCertificateKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}


func NewDatabaseSecretBackendConnectionMysqlAuroraOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseSecretBackendConnectionMysqlAuroraOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseSecretBackendConnectionMysqlAuroraOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnectionMysqlAuroraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseSecretBackendConnectionMysqlAuroraOutputReference_Override(d DatabaseSecretBackendConnectionMysqlAuroraOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnectionMysqlAuroraOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetAuthType(val *string) {
	if err := j.validateSetAuthTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetConnectionUrl(val *string) {
	if err := j.validateSetConnectionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionUrl",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetInternalValue(val *DatabaseSecretBackendConnectionMysqlAurora) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetMaxConnectionLifetime(val *float64) {
	if err := j.validateSetMaxConnectionLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionLifetime",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetMaxIdleConnections(val *float64) {
	if err := j.validateSetMaxIdleConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleConnections",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetMaxOpenConnections(val *float64) {
	if err := j.validateSetMaxOpenConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxOpenConnections",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetServiceAccountJson(val *string) {
	if err := j.validateSetServiceAccountJsonParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountJson",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetTlsCa(val *string) {
	if err := j.validateSetTlsCaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCa",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetTlsCertificateKey(val *string) {
	if err := j.validateSetTlsCertificateKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCertificateKey",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
		val,
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetAuthType() {
	_jsii_.InvokeVoid(
		d,
		"resetAuthType",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetConnectionUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetMaxConnectionLifetime() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxConnectionLifetime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetMaxIdleConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxIdleConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetMaxOpenConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxOpenConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetServiceAccountJson() {
	_jsii_.InvokeVoid(
		d,
		"resetServiceAccountJson",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetTlsCa() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsCa",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetTlsCertificateKey() {
	_jsii_.InvokeVoid(
		d,
		"resetTlsCertificateKey",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionMysqlAuroraOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

