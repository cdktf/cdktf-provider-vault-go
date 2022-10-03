package provider

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-vault-go/vault/v3/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-vault-go/vault/v3/provider/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/vault vault}.
type VaultProvider interface {
	cdktf.TerraformProvider
	AddAddressToEnv() *string
	SetAddAddressToEnv(val *string)
	AddAddressToEnvInput() *string
	Address() *string
	SetAddress(val *string)
	AddressInput() *string
	Alias() *string
	SetAlias(val *string)
	AliasInput() *string
	AuthLogin() interface{}
	SetAuthLogin(val interface{})
	AuthLoginInput() interface{}
	CaCertDir() *string
	SetCaCertDir(val *string)
	CaCertDirInput() *string
	CaCertFile() *string
	SetCaCertFile(val *string)
	CaCertFileInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientAuth() interface{}
	SetClientAuth(val interface{})
	ClientAuthInput() interface{}
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Headers() interface{}
	SetHeaders(val interface{})
	HeadersInput() interface{}
	MaxLeaseTtlSeconds() *float64
	SetMaxLeaseTtlSeconds(val *float64)
	MaxLeaseTtlSecondsInput() *float64
	MaxRetries() *float64
	SetMaxRetries(val *float64)
	MaxRetriesCcc() *float64
	SetMaxRetriesCcc(val *float64)
	MaxRetriesCccInput() *float64
	MaxRetriesInput() *float64
	// Experimental.
	MetaAttributes() *map[string]interface{}
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	// Experimental.
	RawOverrides() interface{}
	SkipChildToken() interface{}
	SetSkipChildToken(val interface{})
	SkipChildTokenInput() interface{}
	SkipTlsVerify() interface{}
	SetSkipTlsVerify(val interface{})
	SkipTlsVerifyInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformProviderSource() *string
	// Experimental.
	TerraformResourceType() *string
	TlsServerName() *string
	SetTlsServerName(val *string)
	TlsServerNameInput() *string
	Token() *string
	SetToken(val *string)
	TokenInput() *string
	TokenName() *string
	SetTokenName(val *string)
	TokenNameInput() *string
	// Experimental.
	AddOverride(path *string, value interface{})
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	ResetAddAddressToEnv()
	ResetAlias()
	ResetAuthLogin()
	ResetCaCertDir()
	ResetCaCertFile()
	ResetClientAuth()
	ResetHeaders()
	ResetMaxLeaseTtlSeconds()
	ResetMaxRetries()
	ResetMaxRetriesCcc()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSkipChildToken()
	ResetSkipTlsVerify()
	ResetTlsServerName()
	ResetToken()
	ResetTokenName()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for VaultProvider
type jsiiProxy_VaultProvider struct {
	internal.Type__cdktfTerraformProvider
}

func (j *jsiiProxy_VaultProvider) AddAddressToEnv() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addAddressToEnv",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AddAddressToEnvInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addAddressToEnvInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Address() *string {
	var returns *string
	_jsii_.Get(
		j,
		"address",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"addressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Alias() *string {
	var returns *string
	_jsii_.Get(
		j,
		"alias",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AliasInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"aliasInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) AuthLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"authLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) CaCertDir() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertDir",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) CaCertDirInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertDirInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) CaCertFile() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertFile",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) CaCertFileInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"caCertFileInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) ClientAuth() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuth",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) ClientAuthInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"clientAuthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Headers() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) HeadersInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"headersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxRetries() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetries",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxRetriesCcc() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesCcc",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxRetriesCccInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesCccInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MaxRetriesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetriesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) MetaAttributes() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"metaAttributes",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SkipChildToken() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipChildToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SkipChildTokenInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipChildTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SkipTlsVerify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipTlsVerify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) SkipTlsVerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"skipTlsVerifyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TerraformProviderSource() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformProviderSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TlsServerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsServerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TlsServerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsServerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) Token() *string {
	var returns *string
	_jsii_.Get(
		j,
		"token",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TokenName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_VaultProvider) TokenNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenNameInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/vault vault} Resource.
func NewVaultProvider(scope constructs.Construct, id *string, config *VaultProviderConfig) VaultProvider {
	_init_.Initialize()

	if err := validateNewVaultProviderParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_VaultProvider{}

	_jsii_.Create(
		"@cdktf/provider-vault.provider.VaultProvider",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/vault vault} Resource.
func NewVaultProvider_Override(v VaultProvider, scope constructs.Construct, id *string, config *VaultProviderConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.provider.VaultProvider",
		[]interface{}{scope, id, config},
		v,
	)
}

func (j *jsiiProxy_VaultProvider)SetAddAddressToEnv(val *string) {
	_jsii_.Set(
		j,
		"addAddressToEnv",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAddress(val *string) {
	_jsii_.Set(
		j,
		"address",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAlias(val *string) {
	_jsii_.Set(
		j,
		"alias",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetAuthLogin(val interface{}) {
	if err := j.validateSetAuthLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"authLogin",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetCaCertDir(val *string) {
	_jsii_.Set(
		j,
		"caCertDir",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetCaCertFile(val *string) {
	_jsii_.Set(
		j,
		"caCertFile",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetClientAuth(val interface{}) {
	if err := j.validateSetClientAuthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientAuth",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetHeaders(val interface{}) {
	if err := j.validateSetHeadersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"headers",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetMaxLeaseTtlSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetMaxRetries(val *float64) {
	_jsii_.Set(
		j,
		"maxRetries",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetMaxRetriesCcc(val *float64) {
	_jsii_.Set(
		j,
		"maxRetriesCcc",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetSkipChildToken(val interface{}) {
	if err := j.validateSetSkipChildTokenParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipChildToken",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetSkipTlsVerify(val interface{}) {
	if err := j.validateSetSkipTlsVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skipTlsVerify",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetTlsServerName(val *string) {
	_jsii_.Set(
		j,
		"tlsServerName",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetToken(val *string) {
	_jsii_.Set(
		j,
		"token",
		val,
	)
}

func (j *jsiiProxy_VaultProvider)SetTokenName(val *string) {
	_jsii_.Set(
		j,
		"tokenName",
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
func VaultProvider_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateVaultProvider_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.provider.VaultProvider",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func VaultProvider_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.provider.VaultProvider",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (v *jsiiProxy_VaultProvider) AddOverride(path *string, value interface{}) {
	if err := v.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (v *jsiiProxy_VaultProvider) OverrideLogicalId(newLogicalId *string) {
	if err := v.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		v,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (v *jsiiProxy_VaultProvider) ResetAddAddressToEnv() {
	_jsii_.InvokeVoid(
		v,
		"resetAddAddressToEnv",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAlias() {
	_jsii_.InvokeVoid(
		v,
		"resetAlias",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetAuthLogin() {
	_jsii_.InvokeVoid(
		v,
		"resetAuthLogin",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetCaCertDir() {
	_jsii_.InvokeVoid(
		v,
		"resetCaCertDir",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetCaCertFile() {
	_jsii_.InvokeVoid(
		v,
		"resetCaCertFile",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetClientAuth() {
	_jsii_.InvokeVoid(
		v,
		"resetClientAuth",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetHeaders() {
	_jsii_.InvokeVoid(
		v,
		"resetHeaders",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetMaxLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetMaxRetries() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxRetries",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetMaxRetriesCcc() {
	_jsii_.InvokeVoid(
		v,
		"resetMaxRetriesCcc",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetNamespace() {
	_jsii_.InvokeVoid(
		v,
		"resetNamespace",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		v,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetSkipChildToken() {
	_jsii_.InvokeVoid(
		v,
		"resetSkipChildToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetSkipTlsVerify() {
	_jsii_.InvokeVoid(
		v,
		"resetSkipTlsVerify",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetTlsServerName() {
	_jsii_.InvokeVoid(
		v,
		"resetTlsServerName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetToken() {
	_jsii_.InvokeVoid(
		v,
		"resetToken",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) ResetTokenName() {
	_jsii_.InvokeVoid(
		v,
		"resetTokenName",
		nil, // no parameters
	)
}

func (v *jsiiProxy_VaultProvider) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		v,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultProvider) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultProvider) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		v,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (v *jsiiProxy_VaultProvider) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		v,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

