// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-vault-go/vault/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-vault-go/vault/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/vault/r/mfa_okta vault_mfa_okta}.
type MfaOkta interface {
	cdktf.TerraformResource
	ApiToken() *string
	SetApiToken(val *string)
	ApiTokenInput() *string
	BaseUrl() *string
	SetBaseUrl(val *string)
	BaseUrlInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	Connection() interface{}
	// Experimental.
	SetConnection(val interface{})
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	Count() *float64
	// Experimental.
	SetCount(val *float64)
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
	MountAccessor() *string
	SetMountAccessor(val *string)
	MountAccessorInput() *string
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	OrgName() *string
	SetOrgName(val *string)
	OrgNameInput() *string
	PrimaryEmail() interface{}
	SetPrimaryEmail(val interface{})
	PrimaryEmailInput() interface{}
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
	UsernameFormat() *string
	SetUsernameFormat(val *string)
	UsernameFormatInput() *string
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
	ResetBaseUrl()
	ResetId()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPrimaryEmail()
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

// The jsii proxy struct for MfaOkta
type jsiiProxy_MfaOkta struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_MfaOkta) ApiToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) ApiTokenInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"apiTokenInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) BaseUrl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) BaseUrlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"baseUrlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) MountAccessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) MountAccessorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountAccessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) OrgName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) OrgNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"orgNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) PrimaryEmail() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) PrimaryEmailInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"primaryEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) UsernameFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_MfaOkta) UsernameFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usernameFormatInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/vault/r/mfa_okta vault_mfa_okta} Resource.
func NewMfaOkta(scope constructs.Construct, id *string, config *MfaOktaConfig) MfaOkta {
	_init_.Initialize()

	j := jsiiProxy_MfaOkta{}

	_jsii_.Create(
		"@cdktf/provider-vault.MfaOkta",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/vault/r/mfa_okta vault_mfa_okta} Resource.
func NewMfaOkta_Override(m MfaOkta, scope constructs.Construct, id *string, config *MfaOktaConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.MfaOkta",
		[]interface{}{scope, id, config},
		m,
	)
}

func (j *jsiiProxy_MfaOkta) SetApiToken(val *string) {
	_jsii_.Set(
		j,
		"apiToken",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetBaseUrl(val *string) {
	_jsii_.Set(
		j,
		"baseUrl",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetMountAccessor(val *string) {
	_jsii_.Set(
		j,
		"mountAccessor",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetName(val *string) {
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetOrgName(val *string) {
	_jsii_.Set(
		j,
		"orgName",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetPrimaryEmail(val interface{}) {
	_jsii_.Set(
		j,
		"primaryEmail",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_MfaOkta) SetUsernameFormat(val *string) {
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
func MfaOkta_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.MfaOkta",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func MfaOkta_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.MfaOkta",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (m *jsiiProxy_MfaOkta) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		m,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (m *jsiiProxy_MfaOkta) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		m,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		m,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		m,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		m,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		m,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		m,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		m,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		m,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (m *jsiiProxy_MfaOkta) ResetBaseUrl() {
	_jsii_.InvokeVoid(
		m,
		"resetBaseUrl",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaOkta) ResetId() {
	_jsii_.InvokeVoid(
		m,
		"resetId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaOkta) ResetNamespace() {
	_jsii_.InvokeVoid(
		m,
		"resetNamespace",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaOkta) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		m,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaOkta) ResetPrimaryEmail() {
	_jsii_.InvokeVoid(
		m,
		"resetPrimaryEmail",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaOkta) ResetUsernameFormat() {
	_jsii_.InvokeVoid(
		m,
		"resetUsernameFormat",
		nil, // no parameters
	)
}

func (m *jsiiProxy_MfaOkta) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		m,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		m,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (m *jsiiProxy_MfaOkta) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		m,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
