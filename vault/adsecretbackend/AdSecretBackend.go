package adsecretbackend

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v9/adsecretbackend/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/3.17.0/docs/resources/ad_secret_backend vault_ad_secret_backend}.
type AdSecretBackend interface {
	cdktf.TerraformResource
	AnonymousGroupSearch() interface{}
	SetAnonymousGroupSearch(val interface{})
	AnonymousGroupSearchInput() interface{}
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	Binddn() *string
	SetBinddn(val *string)
	BinddnInput() *string
	Bindpass() *string
	SetBindpass(val *string)
	BindpassInput() *string
	CaseSensitiveNames() interface{}
	SetCaseSensitiveNames(val interface{})
	CaseSensitiveNamesInput() interface{}
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	SetCertificate(val *string)
	CertificateInput() *string
	ClientTlsCert() *string
	SetClientTlsCert(val *string)
	ClientTlsCertInput() *string
	ClientTlsKey() *string
	SetClientTlsKey(val *string)
	ClientTlsKeyInput() *string
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
	DefaultLeaseTtlSeconds() *float64
	SetDefaultLeaseTtlSeconds(val *float64)
	DefaultLeaseTtlSecondsInput() *float64
	DenyNullBind() interface{}
	SetDenyNullBind(val interface{})
	DenyNullBindInput() interface{}
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
	Discoverdn() interface{}
	SetDiscoverdn(val interface{})
	DiscoverdnInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Formatter() *string
	SetFormatter(val *string)
	FormatterInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Groupattr() *string
	SetGroupattr(val *string)
	GroupattrInput() *string
	Groupdn() *string
	SetGroupdn(val *string)
	GroupdnInput() *string
	Groupfilter() *string
	SetGroupfilter(val *string)
	GroupfilterInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	InsecureTls() interface{}
	SetInsecureTls(val interface{})
	InsecureTlsInput() interface{}
	LastRotationTolerance() *float64
	SetLastRotationTolerance(val *float64)
	LastRotationToleranceInput() *float64
	Length() *float64
	SetLength(val *float64)
	LengthInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Local() interface{}
	SetLocal(val interface{})
	LocalInput() interface{}
	MaxLeaseTtlSeconds() *float64
	SetMaxLeaseTtlSeconds(val *float64)
	MaxLeaseTtlSecondsInput() *float64
	MaxTtl() *float64
	SetMaxTtl(val *float64)
	MaxTtlInput() *float64
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	PasswordPolicy() *string
	SetPasswordPolicy(val *string)
	PasswordPolicyInput() *string
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
	RequestTimeout() *float64
	SetRequestTimeout(val *float64)
	RequestTimeoutInput() *float64
	Starttls() interface{}
	SetStarttls(val interface{})
	StarttlsInput() interface{}
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TlsMaxVersion() *string
	SetTlsMaxVersion(val *string)
	TlsMaxVersionInput() *string
	TlsMinVersion() *string
	SetTlsMinVersion(val *string)
	TlsMinVersionInput() *string
	Ttl() *float64
	SetTtl(val *float64)
	TtlInput() *float64
	Upndomain() *string
	SetUpndomain(val *string)
	UpndomainInput() *string
	Url() *string
	SetUrl(val *string)
	UrlInput() *string
	UsePre111GroupCnBehavior() interface{}
	SetUsePre111GroupCnBehavior(val interface{})
	UsePre111GroupCnBehaviorInput() interface{}
	Userattr() *string
	SetUserattr(val *string)
	UserattrInput() *string
	Userdn() *string
	SetUserdn(val *string)
	UserdnInput() *string
	UseTokenGroups() interface{}
	SetUseTokenGroups(val interface{})
	UseTokenGroupsInput() interface{}
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
	ResetAnonymousGroupSearch()
	ResetBackend()
	ResetCaseSensitiveNames()
	ResetCertificate()
	ResetClientTlsCert()
	ResetClientTlsKey()
	ResetDefaultLeaseTtlSeconds()
	ResetDenyNullBind()
	ResetDescription()
	ResetDisableRemount()
	ResetDiscoverdn()
	ResetFormatter()
	ResetGroupattr()
	ResetGroupdn()
	ResetGroupfilter()
	ResetId()
	ResetInsecureTls()
	ResetLastRotationTolerance()
	ResetLength()
	ResetLocal()
	ResetMaxLeaseTtlSeconds()
	ResetMaxTtl()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPasswordPolicy()
	ResetRequestTimeout()
	ResetStarttls()
	ResetTlsMaxVersion()
	ResetTlsMinVersion()
	ResetTtl()
	ResetUpndomain()
	ResetUrl()
	ResetUsePre111GroupCnBehavior()
	ResetUserattr()
	ResetUserdn()
	ResetUseTokenGroups()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AdSecretBackend
type jsiiProxy_AdSecretBackend struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AdSecretBackend) AnonymousGroupSearch() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousGroupSearch",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) AnonymousGroupSearchInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"anonymousGroupSearchInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Binddn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binddn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) BinddnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"binddnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Bindpass() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpass",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) BindpassInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"bindpassInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) CaseSensitiveNames() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) CaseSensitiveNamesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"caseSensitiveNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) CertificateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) ClientTlsCert() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsCert",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) ClientTlsCertInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsCertInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) ClientTlsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) ClientTlsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientTlsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) DefaultLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) DefaultLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) DenyNullBind() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyNullBind",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) DenyNullBindInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"denyNullBindInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) DisableRemount() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemount",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) DisableRemountInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disableRemountInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Discoverdn() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discoverdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) DiscoverdnInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"discoverdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Formatter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) FormatterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Groupattr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupattr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) GroupattrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupattrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Groupdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) GroupdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Groupfilter() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupfilter",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) GroupfilterInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"groupfilterInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) InsecureTls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) InsecureTlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"insecureTlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) LastRotationTolerance() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastRotationTolerance",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) LastRotationToleranceInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lastRotationToleranceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Length() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"length",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) LengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"lengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Local() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"local",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) LocalInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"localInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) MaxLeaseTtlSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) MaxLeaseTtlSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxLeaseTtlSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) MaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) MaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) PasswordPolicy() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) PasswordPolicyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"passwordPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) RequestTimeout() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeout",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) RequestTimeoutInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"requestTimeoutInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Starttls() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starttls",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) StarttlsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"starttlsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) TlsMaxVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMaxVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) TlsMaxVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMaxVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) TlsMinVersion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) TlsMinVersionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tlsMinVersionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Ttl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) TtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Upndomain() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upndomain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) UpndomainInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"upndomainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Url() *string {
	var returns *string
	_jsii_.Get(
		j,
		"url",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) UrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"urlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) UsePre111GroupCnBehavior() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePre111GroupCnBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) UsePre111GroupCnBehaviorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"usePre111GroupCnBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Userattr() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userattr",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) UserattrInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userattrInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) Userdn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userdn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) UserdnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userdnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) UseTokenGroups() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTokenGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AdSecretBackend) UseTokenGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"useTokenGroupsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.17.0/docs/resources/ad_secret_backend vault_ad_secret_backend} Resource.
func NewAdSecretBackend(scope constructs.Construct, id *string, config *AdSecretBackendConfig) AdSecretBackend {
	_init_.Initialize()

	if err := validateNewAdSecretBackendParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AdSecretBackend{}

	_jsii_.Create(
		"@cdktf/provider-vault.adSecretBackend.AdSecretBackend",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.17.0/docs/resources/ad_secret_backend vault_ad_secret_backend} Resource.
func NewAdSecretBackend_Override(a AdSecretBackend, scope constructs.Construct, id *string, config *AdSecretBackendConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.adSecretBackend.AdSecretBackend",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetAnonymousGroupSearch(val interface{}) {
	if err := j.validateSetAnonymousGroupSearchParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"anonymousGroupSearch",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetBinddn(val *string) {
	if err := j.validateSetBinddnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"binddn",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetBindpass(val *string) {
	if err := j.validateSetBindpassParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"bindpass",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetCaseSensitiveNames(val interface{}) {
	if err := j.validateSetCaseSensitiveNamesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"caseSensitiveNames",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetCertificate(val *string) {
	if err := j.validateSetCertificateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"certificate",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetClientTlsCert(val *string) {
	if err := j.validateSetClientTlsCertParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsCert",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetClientTlsKey(val *string) {
	if err := j.validateSetClientTlsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"clientTlsKey",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetDefaultLeaseTtlSeconds(val *float64) {
	if err := j.validateSetDefaultLeaseTtlSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetDenyNullBind(val interface{}) {
	if err := j.validateSetDenyNullBindParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"denyNullBind",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetDisableRemount(val interface{}) {
	if err := j.validateSetDisableRemountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"disableRemount",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetDiscoverdn(val interface{}) {
	if err := j.validateSetDiscoverdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"discoverdn",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetFormatter(val *string) {
	if err := j.validateSetFormatterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"formatter",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetGroupattr(val *string) {
	if err := j.validateSetGroupattrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupattr",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetGroupdn(val *string) {
	if err := j.validateSetGroupdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupdn",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetGroupfilter(val *string) {
	if err := j.validateSetGroupfilterParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"groupfilter",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetInsecureTls(val interface{}) {
	if err := j.validateSetInsecureTlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"insecureTls",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetLastRotationTolerance(val *float64) {
	if err := j.validateSetLastRotationToleranceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lastRotationTolerance",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetLength(val *float64) {
	if err := j.validateSetLengthParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"length",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetLocal(val interface{}) {
	if err := j.validateSetLocalParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"local",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetMaxLeaseTtlSeconds(val *float64) {
	if err := j.validateSetMaxLeaseTtlSecondsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxLeaseTtlSeconds",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetMaxTtl(val *float64) {
	if err := j.validateSetMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetPasswordPolicy(val *string) {
	if err := j.validateSetPasswordPolicyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"passwordPolicy",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetRequestTimeout(val *float64) {
	if err := j.validateSetRequestTimeoutParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"requestTimeout",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetStarttls(val interface{}) {
	if err := j.validateSetStarttlsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"starttls",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetTlsMaxVersion(val *string) {
	if err := j.validateSetTlsMaxVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsMaxVersion",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetTlsMinVersion(val *string) {
	if err := j.validateSetTlsMinVersionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tlsMinVersion",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetTtl(val *float64) {
	if err := j.validateSetTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetUpndomain(val *string) {
	if err := j.validateSetUpndomainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"upndomain",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetUrl(val *string) {
	if err := j.validateSetUrlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"url",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetUsePre111GroupCnBehavior(val interface{}) {
	if err := j.validateSetUsePre111GroupCnBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usePre111GroupCnBehavior",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetUserattr(val *string) {
	if err := j.validateSetUserattrParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userattr",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetUserdn(val *string) {
	if err := j.validateSetUserdnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userdn",
		val,
	)
}

func (j *jsiiProxy_AdSecretBackend)SetUseTokenGroups(val interface{}) {
	if err := j.validateSetUseTokenGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTokenGroups",
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
func AdSecretBackend_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdSecretBackend_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.adSecretBackend.AdSecretBackend",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AdSecretBackend_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdSecretBackend_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.adSecretBackend.AdSecretBackend",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AdSecretBackend_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAdSecretBackend_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.adSecretBackend.AdSecretBackend",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AdSecretBackend_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.adSecretBackend.AdSecretBackend",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AdSecretBackend) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AdSecretBackend) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := a.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := a.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := a.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := a.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := a.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := a.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) GetStringAttribute(terraformAttribute *string) *string {
	if err := a.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := a.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := a.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetAnonymousGroupSearch() {
	_jsii_.InvokeVoid(
		a,
		"resetAnonymousGroupSearch",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetCaseSensitiveNames() {
	_jsii_.InvokeVoid(
		a,
		"resetCaseSensitiveNames",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetCertificate() {
	_jsii_.InvokeVoid(
		a,
		"resetCertificate",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetClientTlsCert() {
	_jsii_.InvokeVoid(
		a,
		"resetClientTlsCert",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetClientTlsKey() {
	_jsii_.InvokeVoid(
		a,
		"resetClientTlsKey",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetDefaultLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetDenyNullBind() {
	_jsii_.InvokeVoid(
		a,
		"resetDenyNullBind",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetDisableRemount() {
	_jsii_.InvokeVoid(
		a,
		"resetDisableRemount",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetDiscoverdn() {
	_jsii_.InvokeVoid(
		a,
		"resetDiscoverdn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetFormatter() {
	_jsii_.InvokeVoid(
		a,
		"resetFormatter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetGroupattr() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupattr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetGroupdn() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupdn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetGroupfilter() {
	_jsii_.InvokeVoid(
		a,
		"resetGroupfilter",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetInsecureTls() {
	_jsii_.InvokeVoid(
		a,
		"resetInsecureTls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetLastRotationTolerance() {
	_jsii_.InvokeVoid(
		a,
		"resetLastRotationTolerance",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetLength() {
	_jsii_.InvokeVoid(
		a,
		"resetLength",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetLocal() {
	_jsii_.InvokeVoid(
		a,
		"resetLocal",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetMaxLeaseTtlSeconds() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxLeaseTtlSeconds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetPasswordPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetPasswordPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetRequestTimeout() {
	_jsii_.InvokeVoid(
		a,
		"resetRequestTimeout",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetStarttls() {
	_jsii_.InvokeVoid(
		a,
		"resetStarttls",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetTlsMaxVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsMaxVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetTlsMinVersion() {
	_jsii_.InvokeVoid(
		a,
		"resetTlsMinVersion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetUpndomain() {
	_jsii_.InvokeVoid(
		a,
		"resetUpndomain",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetUrl() {
	_jsii_.InvokeVoid(
		a,
		"resetUrl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetUsePre111GroupCnBehavior() {
	_jsii_.InvokeVoid(
		a,
		"resetUsePre111GroupCnBehavior",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetUserattr() {
	_jsii_.InvokeVoid(
		a,
		"resetUserattr",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetUserdn() {
	_jsii_.InvokeVoid(
		a,
		"resetUserdn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) ResetUseTokenGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetUseTokenGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AdSecretBackend) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AdSecretBackend) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

