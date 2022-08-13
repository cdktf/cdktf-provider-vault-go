// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-vault-go/vault/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-vault-go/vault/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/vault/r/azure_secret_backend_role vault_azure_secret_backend_role}.
type AzureSecretBackendRole interface {
	cdktf.TerraformResource
	ApplicationObjectId() *string
	SetApplicationObjectId(val *string)
	ApplicationObjectIdInput() *string
	AzureGroups() AzureSecretBackendRoleAzureGroupsList
	AzureGroupsInput() interface{}
	AzureRoles() AzureSecretBackendRoleAzureRolesList
	AzureRolesInput() interface{}
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
	Count() *float64
	// Experimental.
	SetCount(val *float64)
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
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
	MaxTtl() *string
	SetMaxTtl(val *string)
	MaxTtlInput() *string
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
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
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
	PutAzureGroups(value interface{})
	PutAzureRoles(value interface{})
	ResetApplicationObjectId()
	ResetAzureGroups()
	ResetAzureRoles()
	ResetBackend()
	ResetDescription()
	ResetId()
	ResetMaxTtl()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetTtl()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AzureSecretBackendRole
type jsiiProxy_AzureSecretBackendRole struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AzureSecretBackendRole) ApplicationObjectId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationObjectId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) ApplicationObjectIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"applicationObjectIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) AzureGroups() AzureSecretBackendRoleAzureGroupsList {
	var returns AzureSecretBackendRoleAzureGroupsList
	_jsii_.Get(
		j,
		"azureGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) AzureGroupsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) AzureRoles() AzureSecretBackendRoleAzureRolesList {
	var returns AzureSecretBackendRoleAzureRolesList
	_jsii_.Get(
		j,
		"azureRoles",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) AzureRolesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"azureRolesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) MaxTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) MaxTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"maxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AzureSecretBackendRole) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/vault/r/azure_secret_backend_role vault_azure_secret_backend_role} Resource.
func NewAzureSecretBackendRole(scope constructs.Construct, id *string, config *AzureSecretBackendRoleConfig) AzureSecretBackendRole {
	_init_.Initialize()

	j := jsiiProxy_AzureSecretBackendRole{}

	_jsii_.Create(
		"@cdktf/provider-vault.AzureSecretBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/vault/r/azure_secret_backend_role vault_azure_secret_backend_role} Resource.
func NewAzureSecretBackendRole_Override(a AzureSecretBackendRole, scope constructs.Construct, id *string, config *AzureSecretBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.AzureSecretBackendRole",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetApplicationObjectId(val *string) {
	_jsii_.Set(
		j,
		"applicationObjectId",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetBackend(val *string) {
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetDescription(val *string) {
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetMaxTtl(val *string) {
	_jsii_.Set(
		j,
		"maxTtl",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_AzureSecretBackendRole) SetTtl(val *string) {
	_jsii_.Set(
		j,
		"ttl",
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
func AzureSecretBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.AzureSecretBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AzureSecretBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.AzureSecretBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) PutAzureGroups(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putAzureGroups",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) PutAzureRoles(value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"putAzureRoles",
		[]interface{}{value},
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) ResetApplicationObjectId() {
	_jsii_.InvokeVoid(
		a,
		"resetApplicationObjectId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) ResetAzureGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) ResetAzureRoles() {
	_jsii_.InvokeVoid(
		a,
		"resetAzureRoles",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) ResetDescription() {
	_jsii_.InvokeVoid(
		a,
		"resetDescription",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) ResetMaxTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) ResetTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AzureSecretBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AzureSecretBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

