// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package pkisecretbackendissuer

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v10/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v10/pkisecretbackendissuer/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/pki_secret_backend_issuer vault_pki_secret_backend_issuer}.
type PkiSecretBackendIssuer interface {
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
	CrlDistributionPoints() *[]*string
	SetCrlDistributionPoints(val *[]*string)
	CrlDistributionPointsInput() *[]*string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	EnableAiaUrlTemplating() interface{}
	SetEnableAiaUrlTemplating(val interface{})
	EnableAiaUrlTemplatingInput() interface{}
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
	IssuerId() *string
	IssuerName() *string
	SetIssuerName(val *string)
	IssuerNameInput() *string
	IssuerRef() *string
	SetIssuerRef(val *string)
	IssuerRefInput() *string
	IssuingCertificates() *[]*string
	SetIssuingCertificates(val *[]*string)
	IssuingCertificatesInput() *[]*string
	LeafNotAfterBehavior() *string
	SetLeafNotAfterBehavior(val *string)
	LeafNotAfterBehaviorInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	ManualChain() *[]*string
	SetManualChain(val *[]*string)
	ManualChainInput() *[]*string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	OcspServers() *[]*string
	SetOcspServers(val *[]*string)
	OcspServersInput() *[]*string
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
	RevocationSignatureAlgorithm() *string
	SetRevocationSignatureAlgorithm(val *string)
	RevocationSignatureAlgorithmInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Usage() *string
	SetUsage(val *string)
	UsageInput() *string
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
	ResetCrlDistributionPoints()
	ResetEnableAiaUrlTemplating()
	ResetId()
	ResetIssuerName()
	ResetIssuingCertificates()
	ResetLeafNotAfterBehavior()
	ResetManualChain()
	ResetNamespace()
	ResetOcspServers()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetRevocationSignatureAlgorithm()
	ResetUsage()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PkiSecretBackendIssuer
type jsiiProxy_PkiSecretBackendIssuer struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) CrlDistributionPoints() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crlDistributionPoints",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) CrlDistributionPointsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"crlDistributionPointsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) EnableAiaUrlTemplating() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAiaUrlTemplating",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) EnableAiaUrlTemplatingInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableAiaUrlTemplatingInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) IssuerId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) IssuerName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) IssuerNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) IssuerRef() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerRef",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) IssuerRefInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerRefInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) IssuingCertificates() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"issuingCertificates",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) IssuingCertificatesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"issuingCertificatesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) LeafNotAfterBehavior() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leafNotAfterBehavior",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) LeafNotAfterBehaviorInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"leafNotAfterBehaviorInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) ManualChain() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"manualChain",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) ManualChainInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"manualChainInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) OcspServers() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ocspServers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) OcspServersInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ocspServersInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) RevocationSignatureAlgorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revocationSignatureAlgorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) RevocationSignatureAlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"revocationSignatureAlgorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) Usage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendIssuer) UsageInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"usageInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/pki_secret_backend_issuer vault_pki_secret_backend_issuer} Resource.
func NewPkiSecretBackendIssuer(scope constructs.Construct, id *string, config *PkiSecretBackendIssuerConfig) PkiSecretBackendIssuer {
	_init_.Initialize()

	if err := validateNewPkiSecretBackendIssuerParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_PkiSecretBackendIssuer{}

	_jsii_.Create(
		"@cdktf/provider-vault.pkiSecretBackendIssuer.PkiSecretBackendIssuer",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.20.1/docs/resources/pki_secret_backend_issuer vault_pki_secret_backend_issuer} Resource.
func NewPkiSecretBackendIssuer_Override(p PkiSecretBackendIssuer, scope constructs.Construct, id *string, config *PkiSecretBackendIssuerConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.pkiSecretBackendIssuer.PkiSecretBackendIssuer",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetBackend(val *string) {
	if err := j.validateSetBackendParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetCrlDistributionPoints(val *[]*string) {
	if err := j.validateSetCrlDistributionPointsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"crlDistributionPoints",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetEnableAiaUrlTemplating(val interface{}) {
	if err := j.validateSetEnableAiaUrlTemplatingParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableAiaUrlTemplating",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetIssuerName(val *string) {
	if err := j.validateSetIssuerNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerName",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetIssuerRef(val *string) {
	if err := j.validateSetIssuerRefParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuerRef",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetIssuingCertificates(val *[]*string) {
	if err := j.validateSetIssuingCertificatesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuingCertificates",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetLeafNotAfterBehavior(val *string) {
	if err := j.validateSetLeafNotAfterBehaviorParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"leafNotAfterBehavior",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetManualChain(val *[]*string) {
	if err := j.validateSetManualChainParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"manualChain",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetOcspServers(val *[]*string) {
	if err := j.validateSetOcspServersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"ocspServers",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetRevocationSignatureAlgorithm(val *string) {
	if err := j.validateSetRevocationSignatureAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"revocationSignatureAlgorithm",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendIssuer)SetUsage(val *string) {
	if err := j.validateSetUsageParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"usage",
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
func PkiSecretBackendIssuer_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendIssuer_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.pkiSecretBackendIssuer.PkiSecretBackendIssuer",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendIssuer_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendIssuer_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.pkiSecretBackendIssuer.PkiSecretBackendIssuer",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func PkiSecretBackendIssuer_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validatePkiSecretBackendIssuer_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.pkiSecretBackendIssuer.PkiSecretBackendIssuer",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiSecretBackendIssuer_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.pkiSecretBackendIssuer.PkiSecretBackendIssuer",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) AddOverride(path *string, value interface{}) {
	if err := p.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) OverrideLogicalId(newLogicalId *string) {
	if err := p.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetCrlDistributionPoints() {
	_jsii_.InvokeVoid(
		p,
		"resetCrlDistributionPoints",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetEnableAiaUrlTemplating() {
	_jsii_.InvokeVoid(
		p,
		"resetEnableAiaUrlTemplating",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetIssuerName() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuerName",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetIssuingCertificates() {
	_jsii_.InvokeVoid(
		p,
		"resetIssuingCertificates",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetLeafNotAfterBehavior() {
	_jsii_.InvokeVoid(
		p,
		"resetLeafNotAfterBehavior",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetManualChain() {
	_jsii_.InvokeVoid(
		p,
		"resetManualChain",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetOcspServers() {
	_jsii_.InvokeVoid(
		p,
		"resetOcspServers",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetRevocationSignatureAlgorithm() {
	_jsii_.InvokeVoid(
		p,
		"resetRevocationSignatureAlgorithm",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ResetUsage() {
	_jsii_.InvokeVoid(
		p,
		"resetUsage",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendIssuer) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendIssuer) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

