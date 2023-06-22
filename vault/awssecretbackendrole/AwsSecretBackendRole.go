package awssecretbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v9/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v9/awssecretbackendrole/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/3.17.0/docs/resources/aws_secret_backend_role vault_aws_secret_backend_role}.
type AwsSecretBackendRole interface {
	cdktf.TerraformResource
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
	CredentialType() *string
	SetCredentialType(val *string)
	CredentialTypeInput() *string
	DefaultStsTtl() *float64
	SetDefaultStsTtl(val *float64)
	DefaultStsTtlInput() *float64
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
	IamGroups() *[]*string
	SetIamGroups(val *[]*string)
	IamGroupsInput() *[]*string
	Id() *string
	SetId(val *string)
	IdInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxStsTtl() *float64
	SetMaxStsTtl(val *float64)
	MaxStsTtlInput() *float64
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	PermissionsBoundaryArn() *string
	SetPermissionsBoundaryArn(val *string)
	PermissionsBoundaryArnInput() *string
	PolicyArns() *[]*string
	SetPolicyArns(val *[]*string)
	PolicyArnsInput() *[]*string
	PolicyDocument() *string
	SetPolicyDocument(val *string)
	PolicyDocumentInput() *string
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
	RoleArns() *[]*string
	SetRoleArns(val *[]*string)
	RoleArnsInput() *[]*string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	UserPath() *string
	SetUserPath(val *string)
	UserPathInput() *string
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
	ResetDefaultStsTtl()
	ResetIamGroups()
	ResetId()
	ResetMaxStsTtl()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermissionsBoundaryArn()
	ResetPolicyArns()
	ResetPolicyDocument()
	ResetRoleArns()
	ResetUserPath()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AwsSecretBackendRole
type jsiiProxy_AwsSecretBackendRole struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AwsSecretBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) CredentialType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) CredentialTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"credentialTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) DefaultStsTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultStsTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) DefaultStsTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"defaultStsTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) IamGroups() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamGroups",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) IamGroupsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"iamGroupsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) MaxStsTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStsTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) MaxStsTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxStsTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) PermissionsBoundaryArn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionsBoundaryArn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) PermissionsBoundaryArnInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"permissionsBoundaryArnInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) PolicyArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policyArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) PolicyArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"policyArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) PolicyDocument() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocument",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) PolicyDocumentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"policyDocumentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) RoleArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roleArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) RoleArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"roleArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) UserPath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsSecretBackendRole) UserPathInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"userPathInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.17.0/docs/resources/aws_secret_backend_role vault_aws_secret_backend_role} Resource.
func NewAwsSecretBackendRole(scope constructs.Construct, id *string, config *AwsSecretBackendRoleConfig) AwsSecretBackendRole {
	_init_.Initialize()

	if err := validateNewAwsSecretBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_AwsSecretBackendRole{}

	_jsii_.Create(
		"@cdktf/provider-vault.awsSecretBackendRole.AwsSecretBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.17.0/docs/resources/aws_secret_backend_role vault_aws_secret_backend_role} Resource.
func NewAwsSecretBackendRole_Override(a AwsSecretBackendRole, scope constructs.Construct, id *string, config *AwsSecretBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.awsSecretBackendRole.AwsSecretBackendRole",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetCredentialType(val *string) {
	if err := j.validateSetCredentialTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"credentialType",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetDefaultStsTtl(val *float64) {
	if err := j.validateSetDefaultStsTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"defaultStsTtl",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetIamGroups(val *[]*string) {
	if err := j.validateSetIamGroupsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"iamGroups",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetMaxStsTtl(val *float64) {
	if err := j.validateSetMaxStsTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxStsTtl",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetPermissionsBoundaryArn(val *string) {
	if err := j.validateSetPermissionsBoundaryArnParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"permissionsBoundaryArn",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetPolicyArns(val *[]*string) {
	if err := j.validateSetPolicyArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyArns",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetPolicyDocument(val *string) {
	if err := j.validateSetPolicyDocumentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"policyDocument",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetRoleArns(val *[]*string) {
	if err := j.validateSetRoleArnsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"roleArns",
		val,
	)
}

func (j *jsiiProxy_AwsSecretBackendRole)SetUserPath(val *string) {
	if err := j.validateSetUserPathParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"userPath",
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
func AwsSecretBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSecretBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.awsSecretBackendRole.AwsSecretBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSecretBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSecretBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.awsSecretBackendRole.AwsSecretBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func AwsSecretBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateAwsSecretBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.awsSecretBackendRole.AwsSecretBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsSecretBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.awsSecretBackendRole.AwsSecretBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsSecretBackendRole) AddOverride(path *string, value interface{}) {
	if err := a.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (a *jsiiProxy_AwsSecretBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AwsSecretBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (a *jsiiProxy_AwsSecretBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (a *jsiiProxy_AwsSecretBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (a *jsiiProxy_AwsSecretBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (a *jsiiProxy_AwsSecretBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (a *jsiiProxy_AwsSecretBackendRole) GetStringAttribute(terraformAttribute *string) *string {
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

func (a *jsiiProxy_AwsSecretBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (a *jsiiProxy_AwsSecretBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (a *jsiiProxy_AwsSecretBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := a.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetDefaultStsTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetDefaultStsTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetIamGroups() {
	_jsii_.InvokeVoid(
		a,
		"resetIamGroups",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetMaxStsTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetMaxStsTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetPermissionsBoundaryArn() {
	_jsii_.InvokeVoid(
		a,
		"resetPermissionsBoundaryArn",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetPolicyArns() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetPolicyDocument() {
	_jsii_.InvokeVoid(
		a,
		"resetPolicyDocument",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetRoleArns() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) ResetUserPath() {
	_jsii_.InvokeVoid(
		a,
		"resetUserPath",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsSecretBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecretBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecretBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsSecretBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

