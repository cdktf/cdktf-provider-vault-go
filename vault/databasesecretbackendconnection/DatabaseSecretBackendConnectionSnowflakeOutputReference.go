// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package databasesecretbackendconnection

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v10/jsii"

	"github.com/cdktf/cdktf-provider-vault-go/vault/v10/databasesecretbackendconnection/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatabaseSecretBackendConnectionSnowflakeOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *DatabaseSecretBackendConnectionSnowflake
	SetInternalValue(val *DatabaseSecretBackendConnectionSnowflake)
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
	ResetConnectionUrl()
	ResetMaxConnectionLifetime()
	ResetMaxIdleConnections()
	ResetMaxOpenConnections()
	ResetPassword()
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

// The jsii proxy struct for DatabaseSecretBackendConnectionSnowflakeOutputReference
type jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ConnectionUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ConnectionUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"connectionUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) InternalValue() *DatabaseSecretBackendConnectionSnowflake {
	var returns *DatabaseSecretBackendConnectionSnowflake
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) MaxConnectionLifetime() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionLifetime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) MaxConnectionLifetimeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxConnectionLifetimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) MaxIdleConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) MaxIdleConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxIdleConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) MaxOpenConnections() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpenConnections",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) MaxOpenConnectionsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxOpenConnectionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) Password() *string {
	var returns *string
	_jsii_.Get(
		j,
		"password",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) PasswordInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) Username() *string {
	var returns *string
	_jsii_.Get(
		j,
		"username",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) UsernameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) UsernameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) UsernameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameTemplateInput",
		&returns,
	)
	return returns
}


func NewDatabaseSecretBackendConnectionSnowflakeOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) DatabaseSecretBackendConnectionSnowflakeOutputReference {
	_init_.Initialize()

	if err := validateNewDatabaseSecretBackendConnectionSnowflakeOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnectionSnowflakeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewDatabaseSecretBackendConnectionSnowflakeOutputReference_Override(d DatabaseSecretBackendConnectionSnowflakeOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.databaseSecretBackendConnection.DatabaseSecretBackendConnectionSnowflakeOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		d,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetConnectionUrl(val *string) {
	if err := j.validateSetConnectionUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connectionUrl",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetInternalValue(val *DatabaseSecretBackendConnectionSnowflake) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetMaxConnectionLifetime(val *float64) {
	if err := j.validateSetMaxConnectionLifetimeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxConnectionLifetime",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetMaxIdleConnections(val *float64) {
	if err := j.validateSetMaxIdleConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxIdleConnections",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetMaxOpenConnections(val *float64) {
	if err := j.validateSetMaxOpenConnectionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxOpenConnections",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetPassword(val *string) {
	if err := j.validateSetPasswordParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"password",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetUsername(val *string) {
	if err := j.validateSetUsernameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"username",
		val,
	)
}

func (j *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference)SetUsernameTemplate(val *string) {
	if err := j.validateSetUsernameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameTemplate",
		val,
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ResetConnectionUrl() {
	_jsii_.InvokeVoid(
		d,
		"resetConnectionUrl",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ResetMaxConnectionLifetime() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxConnectionLifetime",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ResetMaxIdleConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxIdleConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ResetMaxOpenConnections() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxOpenConnections",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ResetPassword() {
	_jsii_.InvokeVoid(
		d,
		"resetPassword",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ResetUsername() {
	_jsii_.InvokeVoid(
		d,
		"resetUsername",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ResetUsernameTemplate() {
	_jsii_.InvokeVoid(
		d,
		"resetUsernameTemplate",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (d *jsiiProxy_DatabaseSecretBackendConnectionSnowflakeOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

