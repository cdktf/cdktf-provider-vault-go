// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-vault-go/vault/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-vault-go/vault/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/vault/r/aws_auth_backend_role vault_aws_auth_backend_role}.
type AwsAuthBackendRole interface {
	cdktf.TerraformResource
	AllowInstanceMigration() interface{}
	SetAllowInstanceMigration(val interface{})
	AllowInstanceMigrationInput() interface{}
	AuthType() *string
	SetAuthType(val *string)
	AuthTypeInput() *string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	BoundAccountIds() *[]*string
	SetBoundAccountIds(val *[]*string)
	BoundAccountIdsInput() *[]*string
	BoundAmiIds() *[]*string
	SetBoundAmiIds(val *[]*string)
	BoundAmiIdsInput() *[]*string
	BoundEc2InstanceIds() *[]*string
	SetBoundEc2InstanceIds(val *[]*string)
	BoundEc2InstanceIdsInput() *[]*string
	BoundIamInstanceProfileArns() *[]*string
	SetBoundIamInstanceProfileArns(val *[]*string)
	BoundIamInstanceProfileArnsInput() *[]*string
	BoundIamPrincipalArns() *[]*string
	SetBoundIamPrincipalArns(val *[]*string)
	BoundIamPrincipalArnsInput() *[]*string
	BoundIamRoleArns() *[]*string
	SetBoundIamRoleArns(val *[]*string)
	BoundIamRoleArnsInput() *[]*string
	BoundRegions() *[]*string
	SetBoundRegions(val *[]*string)
	BoundRegionsInput() *[]*string
	BoundSubnetIds() *[]*string
	SetBoundSubnetIds(val *[]*string)
	BoundSubnetIdsInput() *[]*string
	BoundVpcIds() *[]*string
	SetBoundVpcIds(val *[]*string)
	BoundVpcIdsInput() *[]*string
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
	DisallowReauthentication() interface{}
	SetDisallowReauthentication(val interface{})
	DisallowReauthenticationInput() interface{}
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
	InferredAwsRegion() *string
	SetInferredAwsRegion(val *string)
	InferredAwsRegionInput() *string
	InferredEntityType() *string
	SetInferredEntityType(val *string)
	InferredEntityTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
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
	ResolveAwsUniqueIds() interface{}
	SetResolveAwsUniqueIds(val interface{})
	ResolveAwsUniqueIdsInput() interface{}
	Role() *string
	SetRole(val *string)
	RoleId() *string
	RoleInput() *string
	RoleTag() *string
	SetRoleTag(val *string)
	RoleTagInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenBoundCidrs() *[]*string
	SetTokenBoundCidrs(val *[]*string)
	TokenBoundCidrsInput() *[]*string
	TokenExplicitMaxTtl() *float64
	SetTokenExplicitMaxTtl(val *float64)
	TokenExplicitMaxTtlInput() *float64
	TokenMaxTtl() *float64
	SetTokenMaxTtl(val *float64)
	TokenMaxTtlInput() *float64
	TokenNoDefaultPolicy() interface{}
	SetTokenNoDefaultPolicy(val interface{})
	TokenNoDefaultPolicyInput() interface{}
	TokenNumUses() *float64
	SetTokenNumUses(val *float64)
	TokenNumUsesInput() *float64
	TokenPeriod() *float64
	SetTokenPeriod(val *float64)
	TokenPeriodInput() *float64
	TokenPolicies() *[]*string
	SetTokenPolicies(val *[]*string)
	TokenPoliciesInput() *[]*string
	TokenTtl() *float64
	SetTokenTtl(val *float64)
	TokenTtlInput() *float64
	TokenType() *string
	SetTokenType(val *string)
	TokenTypeInput() *string
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
	ResetAllowInstanceMigration()
	ResetAuthType()
	ResetBackend()
	ResetBoundAccountIds()
	ResetBoundAmiIds()
	ResetBoundEc2InstanceIds()
	ResetBoundIamInstanceProfileArns()
	ResetBoundIamPrincipalArns()
	ResetBoundIamRoleArns()
	ResetBoundRegions()
	ResetBoundSubnetIds()
	ResetBoundVpcIds()
	ResetDisallowReauthentication()
	ResetId()
	ResetInferredAwsRegion()
	ResetInferredEntityType()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetResolveAwsUniqueIds()
	ResetRoleTag()
	ResetTokenBoundCidrs()
	ResetTokenExplicitMaxTtl()
	ResetTokenMaxTtl()
	ResetTokenNoDefaultPolicy()
	ResetTokenNumUses()
	ResetTokenPeriod()
	ResetTokenPolicies()
	ResetTokenTtl()
	ResetTokenType()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for AwsAuthBackendRole
type jsiiProxy_AwsAuthBackendRole struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_AwsAuthBackendRole) AllowInstanceMigration() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInstanceMigration",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) AllowInstanceMigrationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowInstanceMigrationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) AuthType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) AuthTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"authTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundAccountIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAccountIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundAccountIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAccountIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundAmiIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAmiIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundAmiIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundAmiIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundEc2InstanceIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundEc2InstanceIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundEc2InstanceIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundEc2InstanceIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamInstanceProfileArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamInstanceProfileArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamInstanceProfileArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamInstanceProfileArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamPrincipalArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamPrincipalArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamPrincipalArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamPrincipalArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamRoleArns() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamRoleArns",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundIamRoleArnsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundIamRoleArnsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundRegions() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundRegions",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundRegionsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundRegionsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundSubnetIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundSubnetIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundSubnetIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundSubnetIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundVpcIds() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundVpcIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) BoundVpcIdsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"boundVpcIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) DisallowReauthentication() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disallowReauthentication",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) DisallowReauthenticationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"disallowReauthenticationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) InferredAwsRegion() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredAwsRegion",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) InferredAwsRegionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredAwsRegionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) InferredEntityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredEntityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) InferredEntityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"inferredEntityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) ResolveAwsUniqueIds() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolveAwsUniqueIds",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) ResolveAwsUniqueIdsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"resolveAwsUniqueIdsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) Role() *string {
	var returns *string
	_jsii_.Get(
		j,
		"role",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) RoleId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) RoleInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) RoleTag() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleTag",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) RoleTagInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"roleTagInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenBoundCidrs() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenBoundCidrsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenBoundCidrsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenExplicitMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenExplicitMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenExplicitMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenNoDefaultPolicy() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenNoDefaultPolicyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"tokenNoDefaultPolicyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenNumUses() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUses",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenNumUsesInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenNumUsesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenPeriod() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriod",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenPeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenPeriodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenPolicies() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPolicies",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenPoliciesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"tokenPoliciesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_AwsAuthBackendRole) TokenTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tokenTypeInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/vault/r/aws_auth_backend_role vault_aws_auth_backend_role} Resource.
func NewAwsAuthBackendRole(scope constructs.Construct, id *string, config *AwsAuthBackendRoleConfig) AwsAuthBackendRole {
	_init_.Initialize()

	j := jsiiProxy_AwsAuthBackendRole{}

	_jsii_.Create(
		"@cdktf/provider-vault.AwsAuthBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/vault/r/aws_auth_backend_role vault_aws_auth_backend_role} Resource.
func NewAwsAuthBackendRole_Override(a AwsAuthBackendRole, scope constructs.Construct, id *string, config *AwsAuthBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.AwsAuthBackendRole",
		[]interface{}{scope, id, config},
		a,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetAllowInstanceMigration(val interface{}) {
	_jsii_.Set(
		j,
		"allowInstanceMigration",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetAuthType(val *string) {
	_jsii_.Set(
		j,
		"authType",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetBackend(val *string) {
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetBoundAccountIds(val *[]*string) {
	_jsii_.Set(
		j,
		"boundAccountIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetBoundAmiIds(val *[]*string) {
	_jsii_.Set(
		j,
		"boundAmiIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetBoundEc2InstanceIds(val *[]*string) {
	_jsii_.Set(
		j,
		"boundEc2InstanceIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetBoundIamInstanceProfileArns(val *[]*string) {
	_jsii_.Set(
		j,
		"boundIamInstanceProfileArns",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetBoundIamPrincipalArns(val *[]*string) {
	_jsii_.Set(
		j,
		"boundIamPrincipalArns",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetBoundIamRoleArns(val *[]*string) {
	_jsii_.Set(
		j,
		"boundIamRoleArns",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetBoundRegions(val *[]*string) {
	_jsii_.Set(
		j,
		"boundRegions",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetBoundSubnetIds(val *[]*string) {
	_jsii_.Set(
		j,
		"boundSubnetIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetBoundVpcIds(val *[]*string) {
	_jsii_.Set(
		j,
		"boundVpcIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetDisallowReauthentication(val interface{}) {
	_jsii_.Set(
		j,
		"disallowReauthentication",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetInferredAwsRegion(val *string) {
	_jsii_.Set(
		j,
		"inferredAwsRegion",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetInferredEntityType(val *string) {
	_jsii_.Set(
		j,
		"inferredEntityType",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetResolveAwsUniqueIds(val interface{}) {
	_jsii_.Set(
		j,
		"resolveAwsUniqueIds",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetRole(val *string) {
	_jsii_.Set(
		j,
		"role",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetRoleTag(val *string) {
	_jsii_.Set(
		j,
		"roleTag",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetTokenBoundCidrs(val *[]*string) {
	_jsii_.Set(
		j,
		"tokenBoundCidrs",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetTokenExplicitMaxTtl(val *float64) {
	_jsii_.Set(
		j,
		"tokenExplicitMaxTtl",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetTokenMaxTtl(val *float64) {
	_jsii_.Set(
		j,
		"tokenMaxTtl",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetTokenNoDefaultPolicy(val interface{}) {
	_jsii_.Set(
		j,
		"tokenNoDefaultPolicy",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetTokenNumUses(val *float64) {
	_jsii_.Set(
		j,
		"tokenNumUses",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetTokenPeriod(val *float64) {
	_jsii_.Set(
		j,
		"tokenPeriod",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetTokenPolicies(val *[]*string) {
	_jsii_.Set(
		j,
		"tokenPolicies",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetTokenTtl(val *float64) {
	_jsii_.Set(
		j,
		"tokenTtl",
		val,
	)
}

func (j *jsiiProxy_AwsAuthBackendRole) SetTokenType(val *string) {
	_jsii_.Set(
		j,
		"tokenType",
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
func AwsAuthBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.AwsAuthBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func AwsAuthBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.AwsAuthBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		a,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		a,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		a,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		a,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		a,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		a,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		a,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		a,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		a,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetAllowInstanceMigration() {
	_jsii_.InvokeVoid(
		a,
		"resetAllowInstanceMigration",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetAuthType() {
	_jsii_.InvokeVoid(
		a,
		"resetAuthType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBackend() {
	_jsii_.InvokeVoid(
		a,
		"resetBackend",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundAccountIds() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundAccountIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundAmiIds() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundAmiIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundEc2InstanceIds() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundEc2InstanceIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundIamInstanceProfileArns() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundIamInstanceProfileArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundIamPrincipalArns() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundIamPrincipalArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundIamRoleArns() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundIamRoleArns",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundRegions() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundRegions",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundSubnetIds() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundSubnetIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetBoundVpcIds() {
	_jsii_.InvokeVoid(
		a,
		"resetBoundVpcIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetDisallowReauthentication() {
	_jsii_.InvokeVoid(
		a,
		"resetDisallowReauthentication",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		a,
		"resetId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetInferredAwsRegion() {
	_jsii_.InvokeVoid(
		a,
		"resetInferredAwsRegion",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetInferredEntityType() {
	_jsii_.InvokeVoid(
		a,
		"resetInferredEntityType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		a,
		"resetNamespace",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		a,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetResolveAwsUniqueIds() {
	_jsii_.InvokeVoid(
		a,
		"resetResolveAwsUniqueIds",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetRoleTag() {
	_jsii_.InvokeVoid(
		a,
		"resetRoleTag",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenBoundCidrs() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenBoundCidrs",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenExplicitMaxTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenExplicitMaxTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenNoDefaultPolicy() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenNoDefaultPolicy",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenNumUses() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenNumUses",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenPeriod() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenPeriod",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenPolicies() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenPolicies",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenTtl() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenTtl",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) ResetTokenType() {
	_jsii_.InvokeVoid(
		a,
		"resetTokenType",
		nil, // no parameters
	)
}

func (a *jsiiProxy_AwsAuthBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		a,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		a,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (a *jsiiProxy_AwsAuthBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		a,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}
