// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tokenauthbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v13/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v13/tokenauthbackendrole/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/token_auth_backend_role vault_token_auth_backend_role}.
type TokenAuthBackendRole interface {
	cdktf.TerraformResource
	AllowedEntityAliases() *[]*string
	SetAllowedEntityAliases(val *[]*string)
	AllowedEntityAliasesInput() *[]*string
	AllowedPolicies() *[]*string
	SetAllowedPolicies(val *[]*string)
	AllowedPoliciesGlob() *[]*string
	SetAllowedPoliciesGlob(val *[]*string)
	AllowedPoliciesGlobInput() *[]*string
	AllowedPoliciesInput() *[]*string
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
	DisallowedPolicies() *[]*string
	SetDisallowedPolicies(val *[]*string)
	DisallowedPoliciesGlob() *[]*string
	SetDisallowedPoliciesGlob(val *[]*string)
	DisallowedPoliciesGlobInput() *[]*string
	DisallowedPoliciesInput() *[]*string
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
	Orphan() interface{}
	SetOrphan(val interface{})
	OrphanInput() interface{}
	PathSuffix() *string
	SetPathSuffix(val *string)
	PathSuffixInput() *string
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
	Renewable() interface{}
	SetRenewable(val interface{})
	RenewableInput() interface{}
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenBoundCidrs() *[]*string
	SetTokenBoundCidrs(val *[]*string)
	TokenBoundCidrsInput() *[]*string
	TokenExplicitMaxTtl() *float64
	SetTokenExplicitMaxTtl(val *float64)
	TokenExplicitMaxTtlInput() *float64
	TokenMaxTtl() *float64
	SetTokenMaxTtl(val *float64)
	TokenMaxTtlInput() *float64
	TokenNoDefaultPolicy() interface{}
	SetTokenNoDefaultPolicy(val interface{})
	TokenNoDefaultPolicyInput() interface{}
	TokenNumUses() *float64
	SetTokenNumUses(val *float64)
	TokenNumUsesInput() *float64
	TokenPeriod() *float64
	SetTokenPeriod(val *float64)
	TokenPeriodInput() *float64
	TokenPolicies() *[]*string
	SetTokenPolicies(val *[]*string)
	TokenPoliciesInput() *[]*string
	TokenTtl() *float64
	SetTokenTtl(val *float64)
	TokenTtlInput() *float64
	TokenType() *string
	SetTokenType(val *string)
	TokenTypeInput() *string
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
	ResetAllowedEntityAliases()
	ResetAllowedPolicies()
	ResetAllowedPoliciesGlob()
	ResetDisallowedPolicies()
	ResetDisallowedPoliciesGlob()
	ResetId()
	ResetNamespace()
	ResetOrphan()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPathSuffix()
	ResetRenewable()
	ResetTokenBoundCidrs()
	ResetTokenExplicitMaxTtl()
	ResetTokenMaxTtl()
	ResetTokenNoDefaultPolicy()
	ResetTokenNumUses()
	ResetTokenPeriod()
	ResetTokenPolicies()
	ResetTokenTtl()
	ResetTokenType()
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

// The jsii proxy struct for TokenAuthBackendRole
type jsiiProxy_TokenAuthBackendRole struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_TokenAuthBackendRole) AllowedEntityAliases() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEntityAliases",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) AllowedEntityAliasesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEntityAliasesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) AllowedPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) AllowedPoliciesGlob() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPoliciesGlob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) AllowedPoliciesGlobInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPoliciesGlobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) AllowedPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) DisallowedPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disallowedPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) DisallowedPoliciesGlob() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disallowedPoliciesGlob",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) DisallowedPoliciesGlobInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disallowedPoliciesGlobInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) DisallowedPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"disallowedPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Orphan() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orphan",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) OrphanInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"orphanInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) PathSuffix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathSuffix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) PathSuffixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathSuffixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) Renewable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renewable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) RenewableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"renewableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenBoundCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenBoundCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenExplicitMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenExplicitMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenNoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenNoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenNumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenNumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_TokenAuthBackendRole) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/token_auth_backend_role vault_token_auth_backend_role} Resource.
func NewTokenAuthBackendRole(scope constructs.Construct, id *string, config *TokenAuthBackendRoleConfig) TokenAuthBackendRole {
	_init_.Initialize()

	if err := validateNewTokenAuthBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_TokenAuthBackendRole{}

	_jsii_.Create(
		"@cdktf/provider-vault.tokenAuthBackendRole.TokenAuthBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/4.4.0/docs/resources/token_auth_backend_role vault_token_auth_backend_role} Resource.
func NewTokenAuthBackendRole_Override(t TokenAuthBackendRole, scope constructs.Construct, id *string, config *TokenAuthBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.tokenAuthBackendRole.TokenAuthBackendRole",
		[]interface{}{scope, id, config},
		t,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetAllowedEntityAliases(val *[]*string) {
	if err := j.validateSetAllowedEntityAliasesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedEntityAliases",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetAllowedPolicies(val *[]*string) {
	if err := j.validateSetAllowedPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedPolicies",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetAllowedPoliciesGlob(val *[]*string) {
	if err := j.validateSetAllowedPoliciesGlobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedPoliciesGlob",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetDisallowedPolicies(val *[]*string) {
	if err := j.validateSetDisallowedPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disallowedPolicies",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetDisallowedPoliciesGlob(val *[]*string) {
	if err := j.validateSetDisallowedPoliciesGlobParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disallowedPoliciesGlob",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetOrphan(val interface{}) {
	if err := j.validateSetOrphanParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"orphan",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetPathSuffix(val *string) {
	if err := j.validateSetPathSuffixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"pathSuffix",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetRenewable(val interface{}) {
	if err := j.validateSetRenewableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"renewable",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetRoleName(val *string) {
	if err := j.validateSetRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetTokenBoundCidrs(val *[]*string) {
	if err := j.validateSetTokenBoundCidrsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenBoundCidrs",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetTokenExplicitMaxTtl(val *float64) {
	if err := j.validateSetTokenExplicitMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenExplicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetTokenMaxTtl(val *float64) {
	if err := j.validateSetTokenMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenMaxTtl",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetTokenNoDefaultPolicy(val interface{}) {
	if err := j.validateSetTokenNoDefaultPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNoDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetTokenNumUses(val *float64) {
	if err := j.validateSetTokenNumUsesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenNumUses",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetTokenPeriod(val *float64) {
	if err := j.validateSetTokenPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPeriod",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetTokenPolicies(val *[]*string) {
	if err := j.validateSetTokenPoliciesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenPolicies",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetTokenTtl(val *float64) {
	if err := j.validateSetTokenTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenTtl",
		val,
	)
}

func (j *jsiiProxy_TokenAuthBackendRole)SetTokenType(val *string) {
	if err := j.validateSetTokenTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenType",
		val,
	)
}

// Generates CDKTF code for importing a TokenAuthBackendRole resource upon running "cdktf plan <stack-name>".
func TokenAuthBackendRole_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validateTokenAuthBackendRole_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.tokenAuthBackendRole.TokenAuthBackendRole",
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
func TokenAuthBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTokenAuthBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.tokenAuthBackendRole.TokenAuthBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TokenAuthBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTokenAuthBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.tokenAuthBackendRole.TokenAuthBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func TokenAuthBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateTokenAuthBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.tokenAuthBackendRole.TokenAuthBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func TokenAuthBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.tokenAuthBackendRole.TokenAuthBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) AddMoveTarget(moveTarget *string) {
	if err := t.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) AddOverride(path *string, value interface{}) {
	if err := t.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := t.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := t.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		t,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := t.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		t,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := t.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		t,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := t.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		t,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := t.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		t,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) GetStringAttribute(terraformAttribute *string) *string {
	if err := t.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		t,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := t.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		t,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := t.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := t.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		t,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) MoveFromId(id *string) {
	if err := t.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveFromId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) MoveTo(moveTarget *string, index interface{}) {
	if err := t.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) MoveToId(id *string) {
	if err := t.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"moveToId",
		[]interface{}{id},
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := t.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		t,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetAllowedEntityAliases() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedEntityAliases",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetAllowedPolicies() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedPolicies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetAllowedPoliciesGlob() {
	_jsii_.InvokeVoid(
		t,
		"resetAllowedPoliciesGlob",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetDisallowedPolicies() {
	_jsii_.InvokeVoid(
		t,
		"resetDisallowedPolicies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetDisallowedPoliciesGlob() {
	_jsii_.InvokeVoid(
		t,
		"resetDisallowedPoliciesGlob",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		t,
		"resetId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		t,
		"resetNamespace",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetOrphan() {
	_jsii_.InvokeVoid(
		t,
		"resetOrphan",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		t,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetPathSuffix() {
	_jsii_.InvokeVoid(
		t,
		"resetPathSuffix",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetRenewable() {
	_jsii_.InvokeVoid(
		t,
		"resetRenewable",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetTokenBoundCidrs() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenBoundCidrs",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetTokenExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenExplicitMaxTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetTokenNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenNoDefaultPolicy",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetTokenNumUses() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenNumUses",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetTokenPeriod() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenPeriod",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetTokenPolicies() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenPolicies",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetTokenTtl() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenTtl",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) ResetTokenType() {
	_jsii_.InvokeVoid(
		t,
		"resetTokenType",
		nil, // no parameters
	)
}

func (t *jsiiProxy_TokenAuthBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		t,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		t,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (t *jsiiProxy_TokenAuthBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		t,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

