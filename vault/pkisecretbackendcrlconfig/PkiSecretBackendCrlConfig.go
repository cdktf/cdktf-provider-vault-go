package pkisecretbackendcrlconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v9/pkisecretbackendcrlconfig/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/3.17.0/docs/resources/pki_secret_backend_crl_config vault_pki_secret_backend_crl_config}.
type PkiSecretBackendCrlConfig interface {
	cdktf.TerraformResource
	AutoRebuild() interface{}
	SetAutoRebuild(val interface{})
	AutoRebuildGracePeriod() *string
	SetAutoRebuildGracePeriod(val *string)
	AutoRebuildGracePeriodInput() *string
	AutoRebuildInput() interface{}
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
	CrossClusterRevocation() interface{}
	SetCrossClusterRevocation(val interface{})
	CrossClusterRevocationInput() interface{}
	DeltaRebuildInterval() *string
	SetDeltaRebuildInterval(val *string)
	DeltaRebuildIntervalInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Disable() interface{}
	SetDisable(val interface{})
	DisableInput() interface{}
	EnableDelta() interface{}
	SetEnableDelta(val interface{})
	EnableDeltaInput() interface{}
	Expiry() *string
	SetExpiry(val *string)
	ExpiryInput() *string
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
	OcspDisable() interface{}
	SetOcspDisable(val interface{})
	OcspDisableInput() interface{}
	OcspExpiry() *string
	SetOcspExpiry(val *string)
	OcspExpiryInput() *string
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
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UnifiedCrl() interface{}
	SetUnifiedCrl(val interface{})
	UnifiedCrlInput() interface{}
	UnifiedCrlOnExistingPaths() interface{}
	SetUnifiedCrlOnExistingPaths(val interface{})
	UnifiedCrlOnExistingPathsInput() interface{}
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
	ResetAutoRebuild()
	ResetAutoRebuildGracePeriod()
	ResetCrossClusterRevocation()
	ResetDeltaRebuildInterval()
	ResetDisable()
	ResetEnableDelta()
	ResetExpiry()
	ResetId()
	ResetNamespace()
	ResetOcspDisable()
	ResetOcspExpiry()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetUnifiedCrl()
	ResetUnifiedCrlOnExistingPaths()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PkiSecretBackendCrlConfig
type jsiiProxy_PkiSecretBackendCrlConfig struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) AutoRebuild() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRebuild",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) AutoRebuildGracePeriod() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoRebuildGracePeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) AutoRebuildGracePeriodInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"autoRebuildGracePeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) AutoRebuildInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"autoRebuildInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) CrossClusterRevocation() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossClusterRevocation",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) CrossClusterRevocationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"crossClusterRevocationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) DeltaRebuildInterval() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaRebuildInterval",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) DeltaRebuildIntervalInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"deltaRebuildIntervalInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Disable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) DisableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) EnableDelta() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDelta",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) EnableDeltaInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableDeltaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Expiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) ExpiryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"expiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) OcspDisable() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspDisable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) OcspDisableInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"ocspDisableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) OcspExpiry() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocspExpiry",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) OcspExpiryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ocspExpiryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) UnifiedCrl() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unifiedCrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) UnifiedCrlInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unifiedCrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) UnifiedCrlOnExistingPaths() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unifiedCrlOnExistingPaths",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig) UnifiedCrlOnExistingPathsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"unifiedCrlOnExistingPathsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.17.0/docs/resources/pki_secret_backend_crl_config vault_pki_secret_backend_crl_config} Resource.
func NewPkiSecretBackendCrlConfig(scope constructs.Construct, id *string, config *PkiSecretBackendCrlConfigConfig) PkiSecretBackendCrlConfig {
	_init_.Initialize()

	if err := validateNewPkiSecretBackendCrlConfigParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiSecretBackendCrlConfig{}

	_jsii_.Create(
		"@cdktf/provider-vault.pkiSecretBackendCrlConfig.PkiSecretBackendCrlConfig",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.17.0/docs/resources/pki_secret_backend_crl_config vault_pki_secret_backend_crl_config} Resource.
func NewPkiSecretBackendCrlConfig_Override(p PkiSecretBackendCrlConfig, scope constructs.Construct, id *string, config *PkiSecretBackendCrlConfigConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.pkiSecretBackendCrlConfig.PkiSecretBackendCrlConfig",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetAutoRebuild(val interface{}) {
	if err := j.validateSetAutoRebuildParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRebuild",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetAutoRebuildGracePeriod(val *string) {
	if err := j.validateSetAutoRebuildGracePeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"autoRebuildGracePeriod",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetCrossClusterRevocation(val interface{}) {
	if err := j.validateSetCrossClusterRevocationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crossClusterRevocation",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetDeltaRebuildInterval(val *string) {
	if err := j.validateSetDeltaRebuildIntervalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"deltaRebuildInterval",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetDisable(val interface{}) {
	if err := j.validateSetDisableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disable",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetEnableDelta(val interface{}) {
	if err := j.validateSetEnableDeltaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableDelta",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetExpiry(val *string) {
	if err := j.validateSetExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"expiry",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetOcspDisable(val interface{}) {
	if err := j.validateSetOcspDisableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspDisable",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetOcspExpiry(val *string) {
	if err := j.validateSetOcspExpiryParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspExpiry",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetUnifiedCrl(val interface{}) {
	if err := j.validateSetUnifiedCrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unifiedCrl",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendCrlConfig)SetUnifiedCrlOnExistingPaths(val interface{}) {
	if err := j.validateSetUnifiedCrlOnExistingPathsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"unifiedCrlOnExistingPaths",
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
func PkiSecretBackendCrlConfig_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendCrlConfig_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.pkiSecretBackendCrlConfig.PkiSecretBackendCrlConfig",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendCrlConfig_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendCrlConfig_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.pkiSecretBackendCrlConfig.PkiSecretBackendCrlConfig",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendCrlConfig_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendCrlConfig_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.pkiSecretBackendCrlConfig.PkiSecretBackendCrlConfig",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiSecretBackendCrlConfig_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.pkiSecretBackendCrlConfig.PkiSecretBackendCrlConfig",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (p *jsiiProxy_PkiSecretBackendCrlConfig) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendCrlConfig) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (p *jsiiProxy_PkiSecretBackendCrlConfig) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (p *jsiiProxy_PkiSecretBackendCrlConfig) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (p *jsiiProxy_PkiSecretBackendCrlConfig) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendCrlConfig) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (p *jsiiProxy_PkiSecretBackendCrlConfig) GetStringAttribute(terraformAttribute *string) *string {
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

func (p *jsiiProxy_PkiSecretBackendCrlConfig) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (p *jsiiProxy_PkiSecretBackendCrlConfig) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (p *jsiiProxy_PkiSecretBackendCrlConfig) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetAutoRebuild() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoRebuild",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetAutoRebuildGracePeriod() {
	_jsii_.InvokeVoid(
		p,
		"resetAutoRebuildGracePeriod",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetCrossClusterRevocation() {
	_jsii_.InvokeVoid(
		p,
		"resetCrossClusterRevocation",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetDeltaRebuildInterval() {
	_jsii_.InvokeVoid(
		p,
		"resetDeltaRebuildInterval",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetDisable() {
	_jsii_.InvokeVoid(
		p,
		"resetDisable",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetEnableDelta() {
	_jsii_.InvokeVoid(
		p,
		"resetEnableDelta",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetExpiry() {
	_jsii_.InvokeVoid(
		p,
		"resetExpiry",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetOcspDisable() {
	_jsii_.InvokeVoid(
		p,
		"resetOcspDisable",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetOcspExpiry() {
	_jsii_.InvokeVoid(
		p,
		"resetOcspExpiry",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetUnifiedCrl() {
	_jsii_.InvokeVoid(
		p,
		"resetUnifiedCrl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ResetUnifiedCrlOnExistingPaths() {
	_jsii_.InvokeVoid(
		p,
		"resetUnifiedCrlOnExistingPaths",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendCrlConfig) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

