package jwtauthbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v7/jwtauthbackend/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs/resources/jwt_auth_backend vault_jwt_auth_backend}.
type JwtAuthBackend interface {
	cdktf.TerraformResource
	Accessor() *string
	BoundIssuer() *string
	SetBoundIssuer(val *string)
	BoundIssuerInput() *string
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
	DefaultRole() *string
	SetDefaultRole(val *string)
	DefaultRoleInput() *string
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
	JwksCaPem() *string
	SetJwksCaPem(val *string)
	JwksCaPemInput() *string
	JwksUrl() *string
	SetJwksUrl(val *string)
	JwksUrlInput() *string
	JwtSupportedAlgs() *[]*string
	SetJwtSupportedAlgs(val *[]*string)
	JwtSupportedAlgsInput() *[]*string
	JwtValidationPubkeys() *[]*string
	SetJwtValidationPubkeys(val *[]*string)
	JwtValidationPubkeysInput() *[]*string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Local() interface{}
	SetLocal(val interface{})
	LocalInput() interface{}
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	NamespaceInState() interface{}
	SetNamespaceInState(val interface{})
	NamespaceInStateInput() interface{}
	// The tree node.
	Node() constructs.Node
	OidcClientId() *string
	SetOidcClientId(val *string)
	OidcClientIdInput() *string
	OidcClientSecret() *string
	SetOidcClientSecret(val *string)
	OidcClientSecretInput() *string
	OidcDiscoveryCaPem() *string
	SetOidcDiscoveryCaPem(val *string)
	OidcDiscoveryCaPemInput() *string
	OidcDiscoveryUrl() *string
	SetOidcDiscoveryUrl(val *string)
	OidcDiscoveryUrlInput() *string
	OidcResponseMode() *string
	SetOidcResponseMode(val *string)
	OidcResponseModeInput() *string
	OidcResponseTypes() *[]*string
	SetOidcResponseTypes(val *[]*string)
	OidcResponseTypesInput() *[]*string
	Path() *string
	SetPath(val *string)
	PathInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	ProviderConfig() *map[string]*string
	SetProviderConfig(val *map[string]*string)
	ProviderConfigInput() *map[string]*string
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
	Tune() JwtAuthBackendTuneList
	TuneInput() interface{}
	Type() *string
	SetType(val *string)
	TypeInput() *string
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
	PutTune(value interface{})
	ResetBoundIssuer()
	ResetDefaultRole()
	ResetDescription()
	ResetDisableRemount()
	ResetId()
	ResetJwksCaPem()
	ResetJwksUrl()
	ResetJwtSupportedAlgs()
	ResetJwtValidationPubkeys()
	ResetLocal()
	ResetNamespace()
	ResetNamespaceInState()
	ResetOidcClientId()
	ResetOidcClientSecret()
	ResetOidcDiscoveryCaPem()
	ResetOidcDiscoveryUrl()
	ResetOidcResponseMode()
	ResetOidcResponseTypes()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPath()
	ResetProviderConfig()
	ResetTune()
	ResetType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for JwtAuthBackend
type jsiiProxy_JwtAuthBackend struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_JwtAuthBackend) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) BoundIssuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boundIssuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) BoundIssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"boundIssuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) DefaultRole() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRole",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) DefaultRoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"defaultRoleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) DisableRemount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) DisableRemountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) JwksCaPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksCaPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) JwksCaPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksCaPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) JwksUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) JwksUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"jwksUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) JwtSupportedAlgs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwtSupportedAlgs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) JwtSupportedAlgsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwtSupportedAlgsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) JwtValidationPubkeys() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwtValidationPubkeys",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) JwtValidationPubkeysInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"jwtValidationPubkeysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Local() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) LocalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) NamespaceInState() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namespaceInState",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) NamespaceInStateInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"namespaceInStateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcClientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcClientIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcClientIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcClientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcClientSecretInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcClientSecretInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcDiscoveryCaPem() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcDiscoveryCaPem",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcDiscoveryCaPemInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcDiscoveryCaPemInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcDiscoveryUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcDiscoveryUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcDiscoveryUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcDiscoveryUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcResponseMode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcResponseMode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcResponseModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"oidcResponseModeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcResponseTypes() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oidcResponseTypes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OidcResponseTypesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"oidcResponseTypesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Path() *string {
	var returns *string
	_jsii_.Get(
		j,
		"path",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) PathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"pathInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) ProviderConfig() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"providerConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) ProviderConfigInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"providerConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Tune() JwtAuthBackendTuneList {
	var returns JwtAuthBackendTuneList
	_jsii_.Get(
		j,
		"tune",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) TuneInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tuneInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs/resources/jwt_auth_backend vault_jwt_auth_backend} Resource.
func NewJwtAuthBackend(scope constructs.Construct, id *string, config *JwtAuthBackendConfig) JwtAuthBackend {
	_init_.Initialize()

	if err := validateNewJwtAuthBackendParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_JwtAuthBackend{}

	_jsii_.Create(
		"@cdktf/provider-vault.jwtAuthBackend.JwtAuthBackend",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.0/docs/resources/jwt_auth_backend vault_jwt_auth_backend} Resource.
func NewJwtAuthBackend_Override(j JwtAuthBackend, scope constructs.Construct, id *string, config *JwtAuthBackendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.jwtAuthBackend.JwtAuthBackend",
		[]interface{}{scope, id, config},
		j,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetBoundIssuer(val *string) {
	if err := j.validateSetBoundIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"boundIssuer",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetDefaultRole(val *string) {
	if err := j.validateSetDefaultRoleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultRole",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetDisableRemount(val interface{}) {
	if err := j.validateSetDisableRemountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRemount",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetJwksCaPem(val *string) {
	if err := j.validateSetJwksCaPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksCaPem",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetJwksUrl(val *string) {
	if err := j.validateSetJwksUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwksUrl",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetJwtSupportedAlgs(val *[]*string) {
	if err := j.validateSetJwtSupportedAlgsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtSupportedAlgs",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetJwtValidationPubkeys(val *[]*string) {
	if err := j.validateSetJwtValidationPubkeysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"jwtValidationPubkeys",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetLocal(val interface{}) {
	if err := j.validateSetLocalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"local",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetNamespaceInState(val interface{}) {
	if err := j.validateSetNamespaceInStateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespaceInState",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetOidcClientId(val *string) {
	if err := j.validateSetOidcClientIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcClientId",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetOidcClientSecret(val *string) {
	if err := j.validateSetOidcClientSecretParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcClientSecret",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetOidcDiscoveryCaPem(val *string) {
	if err := j.validateSetOidcDiscoveryCaPemParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcDiscoveryCaPem",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetOidcDiscoveryUrl(val *string) {
	if err := j.validateSetOidcDiscoveryUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcDiscoveryUrl",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetOidcResponseMode(val *string) {
	if err := j.validateSetOidcResponseModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcResponseMode",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetOidcResponseTypes(val *[]*string) {
	if err := j.validateSetOidcResponseTypesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"oidcResponseTypes",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetPath(val *string) {
	if err := j.validateSetPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"path",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetProviderConfig(val *map[string]*string) {
	if err := j.validateSetProviderConfigParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"providerConfig",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_JwtAuthBackend)SetType(val *string) {
	if err := j.validateSetTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"type",
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
func JwtAuthBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJwtAuthBackend_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.jwtAuthBackend.JwtAuthBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func JwtAuthBackend_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJwtAuthBackend_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.jwtAuthBackend.JwtAuthBackend",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func JwtAuthBackend_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateJwtAuthBackend_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.jwtAuthBackend.JwtAuthBackend",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func JwtAuthBackend_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.jwtAuthBackend.JwtAuthBackend",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_JwtAuthBackend) AddOverride(path *string, value interface{}) {
	if err := j.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (j *jsiiProxy_JwtAuthBackend) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := j.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := j.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		j,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := j.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		j,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := j.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		j,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := j.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		j,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := j.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		j,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) GetStringAttribute(terraformAttribute *string) *string {
	if err := j.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		j,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := j.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		j,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := j.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		j,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) OverrideLogicalId(newLogicalId *string) {
	if err := j.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (j *jsiiProxy_JwtAuthBackend) PutTune(value interface{}) {
	if err := j.validatePutTuneParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		j,
		"putTune",
		[]interface{}{value},
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetBoundIssuer() {
	_jsii_.InvokeVoid(
		j,
		"resetBoundIssuer",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetDefaultRole() {
	_jsii_.InvokeVoid(
		j,
		"resetDefaultRole",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetDescription() {
	_jsii_.InvokeVoid(
		j,
		"resetDescription",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetDisableRemount() {
	_jsii_.InvokeVoid(
		j,
		"resetDisableRemount",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetId() {
	_jsii_.InvokeVoid(
		j,
		"resetId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetJwksCaPem() {
	_jsii_.InvokeVoid(
		j,
		"resetJwksCaPem",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetJwksUrl() {
	_jsii_.InvokeVoid(
		j,
		"resetJwksUrl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetJwtSupportedAlgs() {
	_jsii_.InvokeVoid(
		j,
		"resetJwtSupportedAlgs",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetJwtValidationPubkeys() {
	_jsii_.InvokeVoid(
		j,
		"resetJwtValidationPubkeys",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetLocal() {
	_jsii_.InvokeVoid(
		j,
		"resetLocal",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetNamespace() {
	_jsii_.InvokeVoid(
		j,
		"resetNamespace",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetNamespaceInState() {
	_jsii_.InvokeVoid(
		j,
		"resetNamespaceInState",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetOidcClientId() {
	_jsii_.InvokeVoid(
		j,
		"resetOidcClientId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetOidcClientSecret() {
	_jsii_.InvokeVoid(
		j,
		"resetOidcClientSecret",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetOidcDiscoveryCaPem() {
	_jsii_.InvokeVoid(
		j,
		"resetOidcDiscoveryCaPem",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetOidcDiscoveryUrl() {
	_jsii_.InvokeVoid(
		j,
		"resetOidcDiscoveryUrl",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetOidcResponseMode() {
	_jsii_.InvokeVoid(
		j,
		"resetOidcResponseMode",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetOidcResponseTypes() {
	_jsii_.InvokeVoid(
		j,
		"resetOidcResponseTypes",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		j,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetPath() {
	_jsii_.InvokeVoid(
		j,
		"resetPath",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetProviderConfig() {
	_jsii_.InvokeVoid(
		j,
		"resetProviderConfig",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetTune() {
	_jsii_.InvokeVoid(
		j,
		"resetTune",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) ResetType() {
	_jsii_.InvokeVoid(
		j,
		"resetType",
		nil, // no parameters
	)
}

func (j *jsiiProxy_JwtAuthBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		j,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		j,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (j *jsiiProxy_JwtAuthBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		j,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

