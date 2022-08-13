// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-vault-go/vault/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-vault-go/vault/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/vault/d/azure_access_credentials vault_azure_access_credentials}.
type DataVaultAzureAccessCredentials interface {
	cdktf.TerraformDataSource
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	ClientId() *string
	ClientSecret() *string
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
	Environment() *string
	SetEnvironment(val *string)
	EnvironmentInput() *string
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
	LeaseDuration() *float64
	LeaseId() *string
	LeaseRenewable() cdktf.IResolvable
	LeaseStartTime() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxCredValidationSeconds() *float64
	SetMaxCredValidationSeconds(val *float64)
	MaxCredValidationSecondsInput() *float64
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	NumSecondsBetweenTests() *float64
	SetNumSecondsBetweenTests(val *float64)
	NumSecondsBetweenTestsInput() *float64
	NumSequentialSuccesses() *float64
	SetNumSequentialSuccesses(val *float64)
	NumSequentialSuccessesInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	RawOverrides() interface{}
	Role() *string
	SetRole(val *string)
	RoleInput() *string
	SubscriptionId() *string
	SetSubscriptionId(val *string)
	SubscriptionIdInput() *string
	TenantId() *string
	SetTenantId(val *string)
	TenantIdInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	ValidateCreds() interface{}
	SetValidateCreds(val interface{})
	ValidateCredsInput() interface{}
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
	ResetEnvironment()
	ResetId()
	ResetMaxCredValidationSeconds()
	ResetNamespace()
	ResetNumSecondsBetweenTests()
	ResetNumSequentialSuccesses()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetSubscriptionId()
	ResetTenantId()
	ResetValidateCreds()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for DataVaultAzureAccessCredentials
type jsiiProxy_DataVaultAzureAccessCredentials struct {
	internal.Type__cdktfTerraformDataSource
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) ClientId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) ClientSecret() *string {
	var returns *string
	_jsii_.Get(
		j,
		"clientSecret",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) Environment() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environment",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) EnvironmentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"environmentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) LeaseDuration() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"leaseDuration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) LeaseId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leaseId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) LeaseRenewable() cdktf.IResolvable {
	var returns cdktf.IResolvable
	_jsii_.Get(
		j,
		"leaseRenewable",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) LeaseStartTime() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leaseStartTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) MaxCredValidationSeconds() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCredValidationSeconds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) MaxCredValidationSecondsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxCredValidationSecondsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) NumSecondsBetweenTests() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numSecondsBetweenTests",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) NumSecondsBetweenTestsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numSecondsBetweenTestsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) NumSequentialSuccesses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numSequentialSuccesses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) NumSequentialSuccessesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"numSequentialSuccessesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SubscriptionId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SubscriptionIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subscriptionIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) TenantId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) TenantIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tenantIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) ValidateCreds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateCreds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) ValidateCredsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"validateCredsInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/vault/d/azure_access_credentials vault_azure_access_credentials} Data Source.
func NewDataVaultAzureAccessCredentials(scope constructs.Construct, id *string, config *DataVaultAzureAccessCredentialsConfig) DataVaultAzureAccessCredentials {
	_init_.Initialize()

	j := jsiiProxy_DataVaultAzureAccessCredentials{}

	_jsii_.Create(
		"@cdktf/provider-vault.DataVaultAzureAccessCredentials",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/vault/d/azure_access_credentials vault_azure_access_credentials} Data Source.
func NewDataVaultAzureAccessCredentials_Override(d DataVaultAzureAccessCredentials, scope constructs.Construct, id *string, config *DataVaultAzureAccessCredentialsConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.DataVaultAzureAccessCredentials",
		[]interface{}{scope, id, config},
		d,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetBackend(val *string) {
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetEnvironment(val *string) {
	_jsii_.Set(
		j,
		"environment",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetMaxCredValidationSeconds(val *float64) {
	_jsii_.Set(
		j,
		"maxCredValidationSeconds",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetNumSecondsBetweenTests(val *float64) {
	_jsii_.Set(
		j,
		"numSecondsBetweenTests",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetNumSequentialSuccesses(val *float64) {
	_jsii_.Set(
		j,
		"numSequentialSuccesses",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetSubscriptionId(val *string) {
	_jsii_.Set(
		j,
		"subscriptionId",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetTenantId(val *string) {
	_jsii_.Set(
		j,
		"tenantId",
		val,
	)
}

func (j *jsiiProxy_DataVaultAzureAccessCredentials) SetValidateCreds(val interface{}) {
	_jsii_.Set(
		j,
		"validateCreds",
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
func DataVaultAzureAccessCredentials_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.DataVaultAzureAccessCredentials",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func DataVaultAzureAccessCredentials_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.DataVaultAzureAccessCredentials",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		d,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		d,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ResetEnvironment() {
	_jsii_.InvokeVoid(
		d,
		"resetEnvironment",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ResetId() {
	_jsii_.InvokeVoid(
		d,
		"resetId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ResetMaxCredValidationSeconds() {
	_jsii_.InvokeVoid(
		d,
		"resetMaxCredValidationSeconds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ResetNamespace() {
	_jsii_.InvokeVoid(
		d,
		"resetNamespace",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ResetNumSecondsBetweenTests() {
	_jsii_.InvokeVoid(
		d,
		"resetNumSecondsBetweenTests",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ResetNumSequentialSuccesses() {
	_jsii_.InvokeVoid(
		d,
		"resetNumSequentialSuccesses",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		d,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ResetSubscriptionId() {
	_jsii_.InvokeVoid(
		d,
		"resetSubscriptionId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ResetTenantId() {
	_jsii_.InvokeVoid(
		d,
		"resetTenantId",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ResetValidateCreds() {
	_jsii_.InvokeVoid(
		d,
		"resetValidateCreds",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		d,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DataVaultAzureAccessCredentials) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		d,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
