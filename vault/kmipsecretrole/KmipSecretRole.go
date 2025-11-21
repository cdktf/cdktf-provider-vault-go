// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kmipsecretrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v15/kmipsecretrole/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.5.0/docs/resources/kmip_secret_role vault_kmip_secret_role}.
type KmipSecretRole interface {
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
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	OperationActivate() interface{}
	SetOperationActivate(val interface{})
	OperationActivateInput() interface{}
	OperationAddAttribute() interface{}
	SetOperationAddAttribute(val interface{})
	OperationAddAttributeInput() interface{}
	OperationAll() interface{}
	SetOperationAll(val interface{})
	OperationAllInput() interface{}
	OperationCreate() interface{}
	SetOperationCreate(val interface{})
	OperationCreateInput() interface{}
	OperationDestroy() interface{}
	SetOperationDestroy(val interface{})
	OperationDestroyInput() interface{}
	OperationDiscoverVersions() interface{}
	SetOperationDiscoverVersions(val interface{})
	OperationDiscoverVersionsInput() interface{}
	OperationGet() interface{}
	SetOperationGet(val interface{})
	OperationGetAttributeList() interface{}
	SetOperationGetAttributeList(val interface{})
	OperationGetAttributeListInput() interface{}
	OperationGetAttributes() interface{}
	SetOperationGetAttributes(val interface{})
	OperationGetAttributesInput() interface{}
	OperationGetInput() interface{}
	OperationLocate() interface{}
	SetOperationLocate(val interface{})
	OperationLocateInput() interface{}
	OperationNone() interface{}
	SetOperationNone(val interface{})
	OperationNoneInput() interface{}
	OperationRegister() interface{}
	SetOperationRegister(val interface{})
	OperationRegisterInput() interface{}
	OperationRekey() interface{}
	SetOperationRekey(val interface{})
	OperationRekeyInput() interface{}
	OperationRevoke() interface{}
	SetOperationRevoke(val interface{})
	OperationRevokeInput() interface{}
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
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	Scope() *string
	SetScope(val *string)
	ScopeInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsClientKeyBits() *float64
	SetTlsClientKeyBits(val *float64)
	TlsClientKeyBitsInput() *float64
	TlsClientKeyType() *string
	SetTlsClientKeyType(val *string)
	TlsClientKeyTypeInput() *string
	TlsClientTtl() *float64
	SetTlsClientTtl(val *float64)
	TlsClientTtlInput() *float64
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
	HasResourceMove() interface{}
	// Experimental.
	ImportFrom(id *string, provider cdktf.TerraformProvider)
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable
	// Move the resource corresponding to "id" to this resource.
	//
	// Note that the resource being moved from must be marked as moved using it's instance function.
	// Experimental.
	MoveFromId(id *string)
	// Moves this resource to the target resource given by moveTarget.
	// Experimental.
	MoveTo(moveTarget *string, index interface{})
	// Moves this resource to the resource corresponding to "id".
	// Experimental.
	MoveToId(id *string)
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetId()
	ResetNamespace()
	ResetOperationActivate()
	ResetOperationAddAttribute()
	ResetOperationAll()
	ResetOperationCreate()
	ResetOperationDestroy()
	ResetOperationDiscoverVersions()
	ResetOperationGet()
	ResetOperationGetAttributeList()
	ResetOperationGetAttributes()
	ResetOperationLocate()
	ResetOperationNone()
	ResetOperationRegister()
	ResetOperationRekey()
	ResetOperationRevoke()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTlsClientKeyBits()
	ResetTlsClientKeyType()
	ResetTlsClientTtl()
	SynthesizeAttributes() *map[string]interface{}
	SynthesizeHclAttributes() *map[string]interface{}
	// Experimental.
	ToHclTerraform() interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KmipSecretRole
type jsiiProxy_KmipSecretRole struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KmipSecretRole) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationActivate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationActivate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationActivateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationActivateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationAddAttribute() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationAddAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationAddAttributeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationAddAttributeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationAll() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationAll",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationAllInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationAllInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationCreate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationCreate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationCreateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationCreateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationDestroy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationDestroy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationDestroyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationDestroyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationDiscoverVersions() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationDiscoverVersions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationDiscoverVersionsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationDiscoverVersionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationGet() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationGet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationGetAttributeList() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationGetAttributeList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationGetAttributeListInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationGetAttributeListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationGetAttributes() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationGetAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationGetAttributesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationGetAttributesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationGetInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationGetInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationLocate() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationLocate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationLocateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationLocateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationNone() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationNone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationNoneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationNoneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationRegister() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationRegister",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationRegisterInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationRegisterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationRekey() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationRekey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationRekeyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationRekeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationRevoke() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationRevoke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) OperationRevokeInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"operationRevokeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) Scope() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scope",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) ScopeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"scopeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) TlsClientKeyBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tlsClientKeyBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) TlsClientKeyBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tlsClientKeyBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) TlsClientKeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsClientKeyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) TlsClientKeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsClientKeyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) TlsClientTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tlsClientTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KmipSecretRole) TlsClientTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tlsClientTtlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.5.0/docs/resources/kmip_secret_role vault_kmip_secret_role} Resource.
func NewKmipSecretRole(scope constructs.Construct, id *string, config *KmipSecretRoleConfig) KmipSecretRole {
	_init_.Initialize()

	if err := validateNewKmipSecretRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KmipSecretRole{}

	_jsii_.Create(
		"@cdktf/provider-vault.kmipSecretRole.KmipSecretRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.5.0/docs/resources/kmip_secret_role vault_kmip_secret_role} Resource.
func NewKmipSecretRole_Override(k KmipSecretRole, scope constructs.Construct, id *string, config *KmipSecretRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.kmipSecretRole.KmipSecretRole",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationActivate(val interface{}) {
	if err := j.validateSetOperationActivateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationActivate",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationAddAttribute(val interface{}) {
	if err := j.validateSetOperationAddAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationAddAttribute",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationAll(val interface{}) {
	if err := j.validateSetOperationAllParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationAll",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationCreate(val interface{}) {
	if err := j.validateSetOperationCreateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationCreate",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationDestroy(val interface{}) {
	if err := j.validateSetOperationDestroyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationDestroy",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationDiscoverVersions(val interface{}) {
	if err := j.validateSetOperationDiscoverVersionsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationDiscoverVersions",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationGet(val interface{}) {
	if err := j.validateSetOperationGetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationGet",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationGetAttributeList(val interface{}) {
	if err := j.validateSetOperationGetAttributeListParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationGetAttributeList",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationGetAttributes(val interface{}) {
	if err := j.validateSetOperationGetAttributesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationGetAttributes",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationLocate(val interface{}) {
	if err := j.validateSetOperationLocateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationLocate",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationNone(val interface{}) {
	if err := j.validateSetOperationNoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationNone",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationRegister(val interface{}) {
	if err := j.validateSetOperationRegisterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationRegister",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationRekey(val interface{}) {
	if err := j.validateSetOperationRekeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationRekey",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetOperationRevoke(val interface{}) {
	if err := j.validateSetOperationRevokeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"operationRevoke",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetRole(val *string) {
	if err := j.validateSetRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetScope(val *string) {
	if err := j.validateSetScopeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"scope",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetTlsClientKeyBits(val *float64) {
	if err := j.validateSetTlsClientKeyBitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsClientKeyBits",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetTlsClientKeyType(val *string) {
	if err := j.validateSetTlsClientKeyTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsClientKeyType",
		val,
	)
}

func (j *jsiiProxy_KmipSecretRole)SetTlsClientTtl(val *float64) {
	if err := j.validateSetTlsClientTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsClientTtl",
		val,
	)
}

// Generates CDKTF code for importing a KmipSecretRole resource upon running "cdktf plan <stack-name>".
func KmipSecretRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateKmipSecretRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kmipSecretRole.KmipSecretRole",
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
func KmipSecretRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmipSecretRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kmipSecretRole.KmipSecretRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KmipSecretRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmipSecretRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kmipSecretRole.KmipSecretRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KmipSecretRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKmipSecretRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kmipSecretRole.KmipSecretRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KmipSecretRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.kmipSecretRole.KmipSecretRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KmipSecretRole) AddMoveTarget(moveTarget *string) {
	if err := k.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (k *jsiiProxy_KmipSecretRole) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KmipSecretRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (k *jsiiProxy_KmipSecretRole) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KmipSecretRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (k *jsiiProxy_KmipSecretRole) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (k *jsiiProxy_KmipSecretRole) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (k *jsiiProxy_KmipSecretRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (k *jsiiProxy_KmipSecretRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (k *jsiiProxy_KmipSecretRole) GetStringAttribute(terraformAttribute *string) *string {
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

func (k *jsiiProxy_KmipSecretRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (k *jsiiProxy_KmipSecretRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretRole) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := k.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (k *jsiiProxy_KmipSecretRole) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (k *jsiiProxy_KmipSecretRole) MoveFromId(id *string) {
	if err := k.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveFromId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KmipSecretRole) MoveTo(moveTarget *string, index interface{}) {
	if err := k.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (k *jsiiProxy_KmipSecretRole) MoveToId(id *string) {
	if err := k.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"moveToId",
		[]interface{}{id},
	)
}

func (k *jsiiProxy_KmipSecretRole) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationActivate() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationActivate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationAddAttribute() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationAddAttribute",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationAll() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationAll",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationCreate() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationCreate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationDestroy() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationDestroy",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationDiscoverVersions() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationDiscoverVersions",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationGet() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationGet",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationGetAttributeList() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationGetAttributeList",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationGetAttributes() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationGetAttributes",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationLocate() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationLocate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationNone() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationNone",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationRegister() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationRegister",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationRekey() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationRekey",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOperationRevoke() {
	_jsii_.InvokeVoid(
		k,
		"resetOperationRevoke",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetTlsClientKeyBits() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsClientKeyBits",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetTlsClientKeyType() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsClientKeyType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) ResetTlsClientTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetTlsClientTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KmipSecretRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KmipSecretRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

