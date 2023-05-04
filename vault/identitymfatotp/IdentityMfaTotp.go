package identitymfatotp

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/cdktf/cdktf-provider-vault-go/vault/v7/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/cdktf/cdktf-provider-vault-go/vault/v7/identitymfatotp/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.1/docs/resources/identity_mfa_totp vault_identity_mfa_totp}.
type IdentityMfaTotp interface {
	cdktf.TerraformResource
	Algorithm() *string
	SetAlgorithm(val *string)
	AlgorithmInput() *string
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
	Digits() *float64
	SetDigits(val *float64)
	DigitsInput() *float64
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
	Issuer() *string
	SetIssuer(val *string)
	IssuerInput() *string
	KeySize() *float64
	SetKeySize(val *float64)
	KeySizeInput() *float64
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	MaxValidationAttempts() *float64
	SetMaxValidationAttempts(val *float64)
	MaxValidationAttemptsInput() *float64
	MethodId() *string
	MountAccessor() *string
	Name() *string
	Namespace() *string
	SetNamespace(val *string)
	NamespaceId() *string
	NamespaceInput() *string
	NamespacePath() *string
	// The tree node.
	Node() constructs.Node
	Period() *float64
	SetPeriod(val *float64)
	PeriodInput() *float64
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	QrSize() *float64
	SetQrSize(val *float64)
	QrSizeInput() *float64
	// Experimental.
	RawOverrides() interface{}
	Skew() *float64
	SetSkew(val *float64)
	SkewInput() *float64
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Type() *string
	Uuid() *string
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
	ResetAlgorithm()
	ResetDigits()
	ResetId()
	ResetKeySize()
	ResetMaxValidationAttempts()
	ResetNamespace()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPeriod()
	ResetQrSize()
	ResetSkew()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for IdentityMfaTotp
type jsiiProxy_IdentityMfaTotp struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_IdentityMfaTotp) Algorithm() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithm",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) AlgorithmInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"algorithmInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Count() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Digits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"digits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) DigitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"digitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Issuer() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuer",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) IssuerInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuerInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) KeySize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keySize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) KeySizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keySizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) MaxValidationAttempts() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxValidationAttempts",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) MaxValidationAttemptsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxValidationAttemptsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) MethodId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"methodId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) MountAccessor() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mountAccessor",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) NamespaceId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) NamespacePath() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespacePath",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Period() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"period",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) PeriodInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"periodInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) QrSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qrSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) QrSizeInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"qrSizeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Skew() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skew",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) SkewInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"skewInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_IdentityMfaTotp) Uuid() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uuid",
		&returns,
	)
	return returns
}


// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.1/docs/resources/identity_mfa_totp vault_identity_mfa_totp} Resource.
func NewIdentityMfaTotp(scope constructs.Construct, id *string, config *IdentityMfaTotpConfig) IdentityMfaTotp {
	_init_.Initialize()

	if err := validateNewIdentityMfaTotpParameters(scope, id, config); err != nil {
		panic(err)
	}
	j := jsiiProxy_IdentityMfaTotp{}

	_jsii_.Create(
		"@cdktf/provider-vault.identityMfaTotp.IdentityMfaTotp",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://registry.terraform.io/providers/hashicorp/vault/3.15.1/docs/resources/identity_mfa_totp vault_identity_mfa_totp} Resource.
func NewIdentityMfaTotp_Override(i IdentityMfaTotp, scope constructs.Construct, id *string, config *IdentityMfaTotpConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.identityMfaTotp.IdentityMfaTotp",
		[]interface{}{scope, id, config},
		i,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetAlgorithm(val *string) {
	if err := j.validateSetAlgorithmParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"algorithm",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetConnection(val interface{}) {
	if err := j.validateSetConnectionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetCount(val interface{}) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetDigits(val *float64) {
	if err := j.validateSetDigitsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"digits",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetId(val *string) {
	if err := j.validateSetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetIssuer(val *string) {
	if err := j.validateSetIssuerParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"issuer",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetKeySize(val *float64) {
	if err := j.validateSetKeySizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keySize",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	if err := j.validateSetLifecycleParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetMaxValidationAttempts(val *float64) {
	if err := j.validateSetMaxValidationAttemptsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxValidationAttempts",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetNamespace(val *string) {
	if err := j.validateSetNamespaceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetPeriod(val *float64) {
	if err := j.validateSetPeriodParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"period",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetProvisioners(val *[]interface{}) {
	if err := j.validateSetProvisionersParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetQrSize(val *float64) {
	if err := j.validateSetQrSizeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"qrSize",
		val,
	)
}

func (j *jsiiProxy_IdentityMfaTotp)SetSkew(val *float64) {
	if err := j.validateSetSkewParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"skew",
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
func IdentityMfaTotp_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityMfaTotp_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.identityMfaTotp.IdentityMfaTotp",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityMfaTotp_IsTerraformElement(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityMfaTotp_IsTerraformElementParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.identityMfaTotp.IdentityMfaTotp",
		"isTerraformElement",
		[]interface{}{x},
		&returns,
	)

	return returns
}

// Experimental.
func IdentityMfaTotp_IsTerraformResource(x interface{}) *bool {
	_init_.Initialize()

	if err := validateIdentityMfaTotp_IsTerraformResourceParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.identityMfaTotp.IdentityMfaTotp",
		"isTerraformResource",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func IdentityMfaTotp_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.identityMfaTotp.IdentityMfaTotp",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) AddOverride(path *string, value interface{}) {
	if err := i.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (i *jsiiProxy_IdentityMfaTotp) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := i.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := i.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		i,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := i.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		i,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := i.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		i,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := i.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		i,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := i.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		i,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) GetStringAttribute(terraformAttribute *string) *string {
	if err := i.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		i,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := i.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		i,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := i.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		i,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) OverrideLogicalId(newLogicalId *string) {
	if err := i.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		i,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (i *jsiiProxy_IdentityMfaTotp) ResetAlgorithm() {
	_jsii_.InvokeVoid(
		i,
		"resetAlgorithm",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaTotp) ResetDigits() {
	_jsii_.InvokeVoid(
		i,
		"resetDigits",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaTotp) ResetId() {
	_jsii_.InvokeVoid(
		i,
		"resetId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaTotp) ResetKeySize() {
	_jsii_.InvokeVoid(
		i,
		"resetKeySize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaTotp) ResetMaxValidationAttempts() {
	_jsii_.InvokeVoid(
		i,
		"resetMaxValidationAttempts",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaTotp) ResetNamespace() {
	_jsii_.InvokeVoid(
		i,
		"resetNamespace",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaTotp) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		i,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaTotp) ResetPeriod() {
	_jsii_.InvokeVoid(
		i,
		"resetPeriod",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaTotp) ResetQrSize() {
	_jsii_.InvokeVoid(
		i,
		"resetQrSize",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaTotp) ResetSkew() {
	_jsii_.InvokeVoid(
		i,
		"resetSkew",
		nil, // no parameters
	)
}

func (i *jsiiProxy_IdentityMfaTotp) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		i,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		i,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (i *jsiiProxy_IdentityMfaTotp) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		i,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

