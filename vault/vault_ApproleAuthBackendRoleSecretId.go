// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-vault-go/vault/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-vault-go/vault/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/vault/r/approle_auth_backend_role_secret_id vault_approle_auth_backend_role_secret_id}.
type ApproleAuthBackendRoleSecretId interface {
	cdktf.TerraformResource
	Accessor() *string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	CidrList() *[]*string
	SetCidrList(val *[]*string)
	CidrListInput() *[]*string
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
	Metadata() *string
	SetMetadata(val *string)
	MetadataInput() *string
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
	RoleName() *string
	SetRoleName(val *string)
	RoleNameInput() *string
	SecretId() *string
	SetSecretId(val *string)
	SecretIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	WithWrappedAccessor() interface{}
	SetWithWrappedAccessor(val interface{})
	WithWrappedAccessorInput() interface{}
	WrappingAccessor() *string
	WrappingToken() *string
	WrappingTtl() *string
	SetWrappingTtl(val *string)
	WrappingTtlInput() *string
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
	ResetBackend()
	ResetCidrList()
	ResetId()
	ResetMetadata()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSecretId()
	ResetWithWrappedAccessor()
	ResetWrappingTtl()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for ApproleAuthBackendRoleSecretId
type jsiiProxy_ApproleAuthBackendRoleSecretId struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Accessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"accessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) CidrList() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrList",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) CidrListInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"cidrListInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Metadata() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) MetadataInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) RoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) RoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SecretId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SecretIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"secretIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WithWrappedAccessor() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withWrappedAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WithWrappedAccessorInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"withWrappedAccessorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WrappingAccessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WrappingToken() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingToken",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WrappingTtl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) WrappingTtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"wrappingTtlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/vault/r/approle_auth_backend_role_secret_id vault_approle_auth_backend_role_secret_id} Resource.
func NewApproleAuthBackendRoleSecretId(scope constructs.Construct, id *string, config *ApproleAuthBackendRoleSecretIdConfig) ApproleAuthBackendRoleSecretId {
	_init_.Initialize()

	j := jsiiProxy_ApproleAuthBackendRoleSecretId{}

	_jsii_.Create(
		"@cdktf/provider-vault.ApproleAuthBackendRoleSecretId",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/vault/r/approle_auth_backend_role_secret_id vault_approle_auth_backend_role_secret_id} Resource.
func NewApproleAuthBackendRoleSecretId_Override(a ApproleAuthBackendRoleSecretId, scope constructs.Construct, id *string, config *ApproleAuthBackendRoleSecretIdConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.ApproleAuthBackendRoleSecretId",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetBackend(val *string) {
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetCidrList(val *[]*string) {
	_jsii_.Set(
		j,
		"cidrList",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetMetadata(val *string) {
	_jsii_.Set(
		j,
		"metadata",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetRoleName(val *string) {
	_jsii_.Set(
		j,
		"roleName",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetSecretId(val *string) {
	_jsii_.Set(
		j,
		"secretId",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetWithWrappedAccessor(val interface{}) {
	_jsii_.Set(
		j,
		"withWrappedAccessor",
		val,
	)
}

func (j *jsiiProxy_ApproleAuthBackendRoleSecretId) SetWrappingTtl(val *string) {
	_jsii_.Set(
		j,
		"wrappingTtl",
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
func ApproleAuthBackendRoleSecretId_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.ApproleAuthBackendRoleSecretId",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func ApproleAuthBackendRoleSecretId_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.ApproleAuthBackendRoleSecretId",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetCidrList() {
	_jsii_.InvokeVoid(
		a,
		"resetCidrList",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetMetadata() {
	_jsii_.InvokeVoid(
		a,
		"resetMetadata",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetSecretId() {
	_jsii_.InvokeVoid(
		a,
		"resetSecretId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetWithWrappedAccessor() {
	_jsii_.InvokeVoid(
		a,
		"resetWithWrappedAccessor",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ResetWrappingTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetWrappingTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_ApproleAuthBackendRoleSecretId) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
