// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package kubernetessecretbackendrole

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v10/kubernetessecretbackendrole/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/kubernetes_secret_backend_role vault_kubernetes_secret_backend_role}.
type KubernetesSecretBackendRole interface {
	cdktf.TerraformResource
	AllowedKubernetesNamespaces() *[]*string
	SetAllowedKubernetesNamespaces(val *[]*string)
	AllowedKubernetesNamespacesInput() *[]*string
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
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExtraAnnotations() *map[string]*string
	SetExtraAnnotations(val *map[string]*string)
	ExtraAnnotationsInput() *map[string]*string
	ExtraLabels() *map[string]*string
	SetExtraLabels(val *map[string]*string)
	ExtraLabelsInput() *map[string]*string
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	GeneratedRoleRules() *string
	SetGeneratedRoleRules(val *string)
	GeneratedRoleRulesInput() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	KubernetesRoleName() *string
	SetKubernetesRoleName(val *string)
	KubernetesRoleNameInput() *string
	KubernetesRoleType() *string
	SetKubernetesRoleType(val *string)
	KubernetesRoleTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Name() *string
	SetName(val *string)
	NameInput() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	NameTemplate() *string
	SetNameTemplate(val *string)
	NameTemplateInput() *string
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
	ServiceAccountName() *string
	SetServiceAccountName(val *string)
	ServiceAccountNameInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	TokenDefaultTtl() *float64
	SetTokenDefaultTtl(val *float64)
	TokenDefaultTtlInput() *float64
	TokenMaxTtl() *float64
	SetTokenMaxTtl(val *float64)
	TokenMaxTtlInput() *float64
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
	ResetExtraAnnotations()
	ResetExtraLabels()
	ResetGeneratedRoleRules()
	ResetId()
	ResetKubernetesRoleName()
	ResetKubernetesRoleType()
	ResetNamespace()
	ResetNameTemplate()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetServiceAccountName()
	ResetTokenDefaultTtl()
	ResetTokenMaxTtl()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for KubernetesSecretBackendRole
type jsiiProxy_KubernetesSecretBackendRole struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_KubernetesSecretBackendRole) AllowedKubernetesNamespaces() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedKubernetesNamespaces",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) AllowedKubernetesNamespacesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"allowedKubernetesNamespacesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ExtraAnnotations() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraAnnotations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ExtraAnnotationsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraAnnotationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ExtraLabels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraLabels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ExtraLabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"extraLabelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) GeneratedRoleRules() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedRoleRules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) GeneratedRoleRulesInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"generatedRoleRulesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) KubernetesRoleName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesRoleName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) KubernetesRoleNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesRoleNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) KubernetesRoleType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesRoleType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) KubernetesRoleTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kubernetesRoleTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) NameTemplate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameTemplate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) NameTemplateInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameTemplateInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ServiceAccountName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) ServiceAccountNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TokenDefaultTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenDefaultTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TokenDefaultTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenDefaultTtlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TokenMaxTtl() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_KubernetesSecretBackendRole) TokenMaxTtlInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"tokenMaxTtlInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/kubernetes_secret_backend_role vault_kubernetes_secret_backend_role} Resource.
func NewKubernetesSecretBackendRole(scope constructs.Construct, id *string, config *KubernetesSecretBackendRoleConfig) KubernetesSecretBackendRole {
	_init_.Initialize()

	if err := validateNewKubernetesSecretBackendRoleParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_KubernetesSecretBackendRole{}

	_jsii_.Create(
		"@cdktf/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.21.0/docs/resources/kubernetes_secret_backend_role vault_kubernetes_secret_backend_role} Resource.
func NewKubernetesSecretBackendRole_Override(k KubernetesSecretBackendRole, scope constructs.Construct, id *string, config *KubernetesSecretBackendRoleConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		[]interface{}{scope, id, config},
		k,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetAllowedKubernetesNamespaces(val *[]*string) {
	if err := j.validateSetAllowedKubernetesNamespacesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowedKubernetesNamespaces",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetExtraAnnotations(val *map[string]*string) {
	if err := j.validateSetExtraAnnotationsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraAnnotations",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetExtraLabels(val *map[string]*string) {
	if err := j.validateSetExtraLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"extraLabels",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetGeneratedRoleRules(val *string) {
	if err := j.validateSetGeneratedRoleRulesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"generatedRoleRules",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetKubernetesRoleName(val *string) {
	if err := j.validateSetKubernetesRoleNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesRoleName",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetKubernetesRoleType(val *string) {
	if err := j.validateSetKubernetesRoleTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kubernetesRoleType",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetNameTemplate(val *string) {
	if err := j.validateSetNameTemplateParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"nameTemplate",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetServiceAccountName(val *string) {
	if err := j.validateSetServiceAccountNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountName",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetTokenDefaultTtl(val *float64) {
	if err := j.validateSetTokenDefaultTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenDefaultTtl",
		val,
	)
}

func (j *jsiiProxy_KubernetesSecretBackendRole)SetTokenMaxTtl(val *float64) {
	if err := j.validateSetTokenMaxTtlParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"tokenMaxTtl",
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
func KubernetesSecretBackendRole_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesSecretBackendRole_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesSecretBackendRole_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesSecretBackendRole_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func KubernetesSecretBackendRole_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateKubernetesSecretBackendRole_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func KubernetesSecretBackendRole_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.kubernetesSecretBackendRole.KubernetesSecretBackendRole",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) AddOverride(path *string, value interface{}) {
	if err := k.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := k.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := k.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		k,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := k.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		k,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := k.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		k,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := k.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		k,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := k.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		k,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetStringAttribute(terraformAttribute *string) *string {
	if err := k.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		k,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := k.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		k,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := k.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		k,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) OverrideLogicalId(newLogicalId *string) {
	if err := k.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		k,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetExtraAnnotations() {
	_jsii_.InvokeVoid(
		k,
		"resetExtraAnnotations",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetExtraLabels() {
	_jsii_.InvokeVoid(
		k,
		"resetExtraLabels",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetGeneratedRoleRules() {
	_jsii_.InvokeVoid(
		k,
		"resetGeneratedRoleRules",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetId() {
	_jsii_.InvokeVoid(
		k,
		"resetId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetKubernetesRoleName() {
	_jsii_.InvokeVoid(
		k,
		"resetKubernetesRoleName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetKubernetesRoleType() {
	_jsii_.InvokeVoid(
		k,
		"resetKubernetesRoleType",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetNamespace() {
	_jsii_.InvokeVoid(
		k,
		"resetNamespace",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetNameTemplate() {
	_jsii_.InvokeVoid(
		k,
		"resetNameTemplate",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		k,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetServiceAccountName() {
	_jsii_.InvokeVoid(
		k,
		"resetServiceAccountName",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetTokenDefaultTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenDefaultTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ResetTokenMaxTtl() {
	_jsii_.InvokeVoid(
		k,
		"resetTokenMaxTtl",
		nil, // no parameters
	)
}

func (k *jsiiProxy_KubernetesSecretBackendRole) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		k,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		k,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (k *jsiiProxy_KubernetesSecretBackendRole) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		k,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

