// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package identitymfaduo

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v10/identitymfaduo/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.0/docs/resources/identity_mfa_duo vault_identity_mfa_duo}.
type IdentityMfaDuo interface {
	cdktf.TerraformResource
	ApiHostname() *string
	SetApiHostname(val *string)
	ApiHostnameInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() interface{}
	// Experimental.
	SetCount(val interface{})
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IntegrationKey() *string
	SetIntegrationKey(val *string)
	IntegrationKeyInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MethodId() *string
	MountAccessor() *string
	Name() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceId() *string
	NamespaceInput() *string
	NamespacePath() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	PushInfo() *string
	SetPushInfo(val *string)
	PushInfoInput() *string
	// Experimental.
	RawOverrides() interface{}
	SecretKey() *string
	SetSecretKey(val *string)
	SecretKeyInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	UsePasscode() interface{}
	SetUsePasscode(val interface{})
	UsePasscodeInput() interface{}
	UsernameFormat() *string
	SetUsernameFormat(val *string)
	UsernameFormatInput() *string
	Uuid() *string
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
	ResetId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPushInfo()
	ResetUsePasscode()
	ResetUsernameFormat()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for IdentityMfaDuo
type jsiiProxy_IdentityMfaDuo struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IdentityMfaDuo) ApiHostname() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiHostname",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) ApiHostnameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiHostnameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) IntegrationKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) IntegrationKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"integrationKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) MethodId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) MountAccessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) NamespacePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespacePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) PushInfo() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushInfo",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) PushInfoInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pushInfoInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) SecretKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) SecretKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) UsePasscode() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePasscode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) UsePasscodeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePasscodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) UsernameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) UsernameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaDuo) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.0/docs/resources/identity_mfa_duo vault_identity_mfa_duo} Resource.
func NewIdentityMfaDuo(scope constructs.Construct, id *string, config *IdentityMfaDuoConfig) IdentityMfaDuo {
	_init_.Initialize()

	if err := validateNewIdentityMfaDuoParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityMfaDuo{}

	_jsii_.Create(
		"@cdktf/provider-vault.identityMfaDuo.IdentityMfaDuo",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.0/docs/resources/identity_mfa_duo vault_identity_mfa_duo} Resource.
func NewIdentityMfaDuo_Override(i IdentityMfaDuo, scope constructs.Construct, id *string, config *IdentityMfaDuoConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.identityMfaDuo.IdentityMfaDuo",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetApiHostname(val *string) {
	if err := j.validateSetApiHostnameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"apiHostname",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetIntegrationKey(val *string) {
	if err := j.validateSetIntegrationKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"integrationKey",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetPushInfo(val *string) {
	if err := j.validateSetPushInfoParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pushInfo",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetSecretKey(val *string) {
	if err := j.validateSetSecretKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"secretKey",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetUsePasscode(val interface{}) {
	if err := j.validateSetUsePasscodeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePasscode",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaDuo)SetUsernameFormat(val *string) {
	if err := j.validateSetUsernameFormatParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usernameFormat",
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
func IdentityMfaDuo_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityMfaDuo_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.identityMfaDuo.IdentityMfaDuo",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityMfaDuo_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityMfaDuo_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.identityMfaDuo.IdentityMfaDuo",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityMfaDuo_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityMfaDuo_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.identityMfaDuo.IdentityMfaDuo",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IdentityMfaDuo_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.identityMfaDuo.IdentityMfaDuo",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IdentityMfaDuo) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IdentityMfaDuo) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaDuo) ResetNamespace() {
	_jsii_.InvokeVoid(
		i,
		"resetNamespace",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaDuo) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaDuo) ResetPushInfo() {
	_jsii_.InvokeVoid(
		i,
		"resetPushInfo",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaDuo) ResetUsePasscode() {
	_jsii_.InvokeVoid(
		i,
		"resetUsePasscode",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaDuo) ResetUsernameFormat() {
	_jsii_.InvokeVoid(
		i,
		"resetUsernameFormat",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaDuo) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaDuo) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

