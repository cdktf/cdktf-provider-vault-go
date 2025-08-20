// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendconfigscep

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v15/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v15/pkisecretbackendconfigscep/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/5.2.1/docs/resources/pki_secret_backend_config_scep vault_pki_secret_backend_config_scep}.
type PkiSecretBackendConfigScep interface {
	cdktf.TerraformResource
	AllowedDigestAlgorithms() *[]*string
	SetAllowedDigestAlgorithms(val *[]*string)
	AllowedDigestAlgorithmsInput() *[]*string
	AllowedEncryptionAlgorithms() *[]*string
	SetAllowedEncryptionAlgorithms(val *[]*string)
	AllowedEncryptionAlgorithmsInput() *[]*string
	Authenticators() PkiSecretBackendConfigScepAuthenticatorsOutputReference
	AuthenticatorsInput() *PkiSecretBackendConfigScepAuthenticators
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
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
	DefaultPathPolicy() *string
	SetDefaultPathPolicy(val *string)
	DefaultPathPolicyInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Enabled() interface{}
	SetEnabled(val interface{})
	EnabledInput() interface{}
	ExternalValidation() PkiSecretBackendConfigScepExternalValidationList
	ExternalValidationInput() interface{}
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
	LastUpdated() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	LogLevel() *string
	SetLogLevel(val *string)
	LogLevelInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
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
	// Experimental.
	RawOverrides() interface{}
	RestrictCaChainToIssuer() interface{}
	SetRestrictCaChainToIssuer(val interface{})
	RestrictCaChainToIssuerInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
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
	PutAuthenticators(value *PkiSecretBackendConfigScepAuthenticators)
	PutExternalValidation(value interface{})
	ResetAllowedDigestAlgorithms()
	ResetAllowedEncryptionAlgorithms()
	ResetAuthenticators()
	ResetDefaultPathPolicy()
	ResetEnabled()
	ResetExternalValidation()
	ResetId()
	ResetLogLevel()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRestrictCaChainToIssuer()
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

// The jsii proxy struct for PkiSecretBackendConfigScep
type jsiiProxy_PkiSecretBackendConfigScep struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) AllowedDigestAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDigestAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) AllowedDigestAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedDigestAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) AllowedEncryptionAlgorithms() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEncryptionAlgorithms",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) AllowedEncryptionAlgorithmsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedEncryptionAlgorithmsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Authenticators() PkiSecretBackendConfigScepAuthenticatorsOutputReference {
	var returns PkiSecretBackendConfigScepAuthenticatorsOutputReference
	_jsii_.Get(
		j,
		"authenticators",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) AuthenticatorsInput() *PkiSecretBackendConfigScepAuthenticators {
	var returns *PkiSecretBackendConfigScepAuthenticators
	_jsii_.Get(
		j,
		"authenticatorsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) DefaultPathPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPathPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) DefaultPathPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultPathPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Enabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) EnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) ExternalValidation() PkiSecretBackendConfigScepExternalValidationList {
	var returns PkiSecretBackendConfigScepExternalValidationList
	_jsii_.Get(
		j,
		"externalValidation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) ExternalValidationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"externalValidationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) LastUpdated() *string {
	var returns *string
	_jsii_.Get(
		j,
		"lastUpdated",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) LogLevel() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevel",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) LogLevelInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"logLevelInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) RestrictCaChainToIssuer() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictCaChainToIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) RestrictCaChainToIssuerInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"restrictCaChainToIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendConfigScep) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.2.1/docs/resources/pki_secret_backend_config_scep vault_pki_secret_backend_config_scep} Resource.
func NewPkiSecretBackendConfigScep(scope constructs.Construct, id *string, config *PkiSecretBackendConfigScepConfig) PkiSecretBackendConfigScep {
	_init_.Initialize()

	if err := validateNewPkiSecretBackendConfigScepParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiSecretBackendConfigScep{}

	_jsii_.Create(
		"@cdktf/provider-vault.pkiSecretBackendConfigScep.PkiSecretBackendConfigScep",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/5.2.1/docs/resources/pki_secret_backend_config_scep vault_pki_secret_backend_config_scep} Resource.
func NewPkiSecretBackendConfigScep_Override(p PkiSecretBackendConfigScep, scope constructs.Construct, id *string, config *PkiSecretBackendConfigScepConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.pkiSecretBackendConfigScep.PkiSecretBackendConfigScep",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetAllowedDigestAlgorithms(val *[]*string) {
	if err := j.validateSetAllowedDigestAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedDigestAlgorithms",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetAllowedEncryptionAlgorithms(val *[]*string) {
	if err := j.validateSetAllowedEncryptionAlgorithmsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedEncryptionAlgorithms",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetDefaultPathPolicy(val *string) {
	if err := j.validateSetDefaultPathPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultPathPolicy",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetEnabled(val interface{}) {
	if err := j.validateSetEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enabled",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetLogLevel(val *string) {
	if err := j.validateSetLogLevelParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"logLevel",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendConfigScep)SetRestrictCaChainToIssuer(val interface{}) {
	if err := j.validateSetRestrictCaChainToIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"restrictCaChainToIssuer",
		val,
	)
}

// Generates CDKTF code for importing a PkiSecretBackendConfigScep resource upon running "cdktf plan <stack-name>".
func PkiSecretBackendConfigScep_GenerateConfigForImport(scope constructs.Construct, importToId *string, importFromId *string, provider cdktf.TerraformProvider) cdktf.ImportableResource {
	_init_.Initialize()

	if err := validatePkiSecretBackendConfigScep_GenerateConfigForImportParameters(scope, importToId, importFromId); err != nil {
		panic(err)
	}
	var returns cdktf.ImportableResource

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.pkiSecretBackendConfigScep.PkiSecretBackendConfigScep",
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
func PkiSecretBackendConfigScep_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendConfigScep_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.pkiSecretBackendConfigScep.PkiSecretBackendConfigScep",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendConfigScep_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendConfigScep_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.pkiSecretBackendConfigScep.PkiSecretBackendConfigScep",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendConfigScep_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendConfigScep_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.pkiSecretBackendConfigScep.PkiSecretBackendConfigScep",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiSecretBackendConfigScep_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.pkiSecretBackendConfigScep.PkiSecretBackendConfigScep",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) AddMoveTarget(moveTarget *string) {
	if err := p.validateAddMoveTargetParameters(moveTarget); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addMoveTarget",
		[]interface{}{moveTarget},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) HasResourceMove() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"hasResourceMove",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ImportFrom(id *string, provider cdktf.TerraformProvider) {
	if err := p.validateImportFromParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"importFrom",
		[]interface{}{id, provider},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) MoveFromId(id *string) {
	if err := p.validateMoveFromIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveFromId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) MoveTo(moveTarget *string, index interface{}) {
	if err := p.validateMoveToParameters(moveTarget, index); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveTo",
		[]interface{}{moveTarget, index},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) MoveToId(id *string) {
	if err := p.validateMoveToIdParameters(id); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"moveToId",
		[]interface{}{id},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) PutAuthenticators(value *PkiSecretBackendConfigScepAuthenticators) {
	if err := p.validatePutAuthenticatorsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putAuthenticators",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) PutExternalValidation(value interface{}) {
	if err := p.validatePutExternalValidationParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"putExternalValidation",
		[]interface{}{value},
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetAllowedDigestAlgorithms() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedDigestAlgorithms",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetAllowedEncryptionAlgorithms() {
	_jsii_.InvokeVoid(
		p,
		"resetAllowedEncryptionAlgorithms",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetAuthenticators() {
	_jsii_.InvokeVoid(
		p,
		"resetAuthenticators",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetDefaultPathPolicy() {
	_jsii_.InvokeVoid(
		p,
		"resetDefaultPathPolicy",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetEnabled() {
	_jsii_.InvokeVoid(
		p,
		"resetEnabled",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetExternalValidation() {
	_jsii_.InvokeVoid(
		p,
		"resetExternalValidation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetLogLevel() {
	_jsii_.InvokeVoid(
		p,
		"resetLogLevel",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ResetRestrictCaChainToIssuer() {
	_jsii_.InvokeVoid(
		p,
		"resetRestrictCaChainToIssuer",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) SynthesizeHclAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeHclAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ToHclTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toHclTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendConfigScep) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

