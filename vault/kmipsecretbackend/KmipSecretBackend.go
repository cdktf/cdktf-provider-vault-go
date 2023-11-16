// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmipsecretbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v11/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v11/kmipsecretbackend/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/kmip_secret_backend vault_kmip_secret_backend}.
type KmipSecretBackend interface {
	cdktf.TerraformResource
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
	DefaultTlsClientKeyBits() *float64
	SetDefaultTlsClientKeyBits(val *float64)
	DefaultTlsClientKeyBitsInput() *float64
	DefaultTlsClientKeyType() *string
	SetDefaultTlsClientKeyType(val *string)
	DefaultTlsClientKeyTypeInput() *string
	DefaultTlsClientTtl() *float64
	SetDefaultTlsClientTtl(val *float64)
	DefaultTlsClientTtlInput() *float64
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	DisableRemount() interface{}
	SetDisableRemount(val interface{})
	DisableRemountInput() interface{}
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ListenAddrs() *[]*string
	SetListenAddrs(val *[]*string)
	ListenAddrsInput() *[]*string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Path() *string
	SetPath(val *string)
	PathInput() *string
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
	ServerHostnames() *[]*string
	SetServerHostnames(val *[]*string)
	ServerHostnamesInput() *[]*string
	ServerIps() *[]*string
	SetServerIps(val *[]*string)
	ServerIpsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsCaKeyBits() *float64
	SetTlsCaKeyBits(val *float64)
	TlsCaKeyBitsInput() *float64
	TlsCaKeyType() *string
	SetTlsCaKeyType(val *string)
	TlsCaKeyTypeInput() *string
	TlsMinVersion() *string
	SetTlsMinVersion(val *string)
	TlsMinVersionInput() *string
	// Adds a user defined moveTarget string to this resource to be later used in .moveTo(moveTarget) to resolve the location of the move.
	// Experimental.
	AddMoveTarget(moveTarget *string)
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
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetDefaultTlsClientKeyBits()
	ResetDefaultTlsClientKeyType()
	ResetDefaultTlsClientTtl()
	ResetDescription()
	ResetDisableRemount()
	ResetId()
	ResetListenAddrs()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServerHostnames()
	ResetServerIps()
	ResetTlsCaKeyBits()
	ResetTlsCaKeyType()
	ResetTlsMinVersion()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KmipSecretBackend
type jsiiProxy_KmipSecretBackend struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KmipSecretBackend) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) DefaultTlsClientKeyBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTlsClientKeyBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) DefaultTlsClientKeyBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTlsClientKeyBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) DefaultTlsClientKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTlsClientKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) DefaultTlsClientKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultTlsClientKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) DefaultTlsClientTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTlsClientTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) DefaultTlsClientTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultTlsClientTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) DisableRemount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) DisableRemountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) ListenAddrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listenAddrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) ListenAddrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"listenAddrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) ServerHostnames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverHostnames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) ServerHostnamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverHostnamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) ServerIps() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) ServerIpsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"serverIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) TlsCaKeyBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tlsCaKeyBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) TlsCaKeyBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tlsCaKeyBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) TlsCaKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCaKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) TlsCaKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsCaKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) TlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretBackend) TlsMinVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersionInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/kmip_secret_backend vault_kmip_secret_backend} Resource.
func NewKmipSecretBackend(scope constructs.Construct, id *string, config *KmipSecretBackendConfig) KmipSecretBackend {
	_init_.Initialize()

	if err := validateNewKmipSecretBackendParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KmipSecretBackend{}

	_jsii_.Create(
		"@cdktf/provider-vault.kmipSecretBackend.KmipSecretBackend",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.23.0/docs/resources/kmip_secret_backend vault_kmip_secret_backend} Resource.
func NewKmipSecretBackend_Override(k KmipSecretBackend, scope constructs.Construct, id *string, config *KmipSecretBackendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.kmipSecretBackend.KmipSecretBackend",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetDefaultTlsClientKeyBits(val *float64) {
	if err := j.validateSetDefaultTlsClientKeyBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTlsClientKeyBits",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetDefaultTlsClientKeyType(val *string) {
	if err := j.validateSetDefaultTlsClientKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTlsClientKeyType",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetDefaultTlsClientTtl(val *float64) {
	if err := j.validateSetDefaultTlsClientTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultTlsClientTtl",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetDisableRemount(val interface{}) {
	if err := j.validateSetDisableRemountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRemount",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetListenAddrs(val *[]*string) {
	if err := j.validateSetListenAddrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"listenAddrs",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetServerHostnames(val *[]*string) {
	if err := j.validateSetServerHostnamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverHostnames",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetServerIps(val *[]*string) {
	if err := j.validateSetServerIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serverIps",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetTlsCaKeyBits(val *float64) {
	if err := j.validateSetTlsCaKeyBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCaKeyBits",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetTlsCaKeyType(val *string) {
	if err := j.validateSetTlsCaKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsCaKeyType",
		val,
	)
}

func (j *jsiiProxy_KmipSecretBackend)SetTlsMinVersion(val *string) {
	if err := j.validateSetTlsMinVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsMinVersion",
		val,
	)
}

// Generates CDKTF code for importing a KmipSecretBackend resource upon running "cdktf plan <stack-name>".
func KmipSecretBackend_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKmipSecretBackend_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kmipSecretBackend.KmipSecretBackend",
		"generateConfigForImport",
		[]interface{}{scope, importToId, importFromId, provider},
		&returns,
	)

	return returns
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
func KmipSecretBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmipSecretBackend_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kmipSecretBackend.KmipSecretBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KmipSecretBackend_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmipSecretBackend_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kmipSecretBackend.KmipSecretBackend",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KmipSecretBackend_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmipSecretBackend_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kmipSecretBackend.KmipSecretBackend",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KmipSecretBackend_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.kmipSecretBackend.KmipSecretBackend",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KmipSecretBackend) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KmipSecretBackend) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KmipSecretBackend) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KmipSecretBackend) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KmipSecretBackend) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetDefaultTlsClientKeyBits() {
	_jsii_.InvokeVoid(
		k,
		"resetDefaultTlsClientKeyBits",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetDefaultTlsClientKeyType() {
	_jsii_.InvokeVoid(
		k,
		"resetDefaultTlsClientKeyType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetDefaultTlsClientTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetDefaultTlsClientTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetDescription() {
	_jsii_.InvokeVoid(
		k,
		"resetDescription",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetDisableRemount() {
	_jsii_.InvokeVoid(
		k,
		"resetDisableRemount",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetListenAddrs() {
	_jsii_.InvokeVoid(
		k,
		"resetListenAddrs",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetServerHostnames() {
	_jsii_.InvokeVoid(
		k,
		"resetServerHostnames",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetServerIps() {
	_jsii_.InvokeVoid(
		k,
		"resetServerIps",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetTlsCaKeyBits() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsCaKeyBits",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetTlsCaKeyType() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsCaKeyType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) ResetTlsMinVersion() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsMinVersion",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

