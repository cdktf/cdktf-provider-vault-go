// Prebuilt vault Provider for Terraform CDK (cdktf)
package vault

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/hashicorp/cdktf-provider-vault-go/vault/v2/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/cdktf-provider-vault-go/vault/v2/internal"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

// Represents a {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert vault_pki_secret_backend_root_cert}.
type PkiSecretBackendRootCert interface {
	cdktf.TerraformResource
	AltNames() *[]*string
	SetAltNames(val *[]*string)
	AltNamesInput() *[]*string
	Backend() *string
	SetBackend(val *string)
	BackendInput() *string
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	Certificate() *string
	CommonName() *string
	SetCommonName(val *string)
	CommonNameInput() *string
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
	Country() *string
	SetCountry(val *string)
	CountryInput() *string
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	ExcludeCnFromSans() interface{}
	SetExcludeCnFromSans(val interface{})
	ExcludeCnFromSansInput() interface{}
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	Format() *string
	SetFormat(val *string)
	FormatInput() *string
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	Id() *string
	SetId(val *string)
	IdInput() *string
	IpSans() *[]*string
	SetIpSans(val *[]*string)
	IpSansInput() *[]*string
	IssuingCa() *string
	KeyBits() *float64
	SetKeyBits(val *float64)
	KeyBitsInput() *float64
	KeyType() *string
	SetKeyType(val *string)
	KeyTypeInput() *string
	// Experimental.
	Lifecycle() *cdktf.TerraformResourceLifecycle
	// Experimental.
	SetLifecycle(val *cdktf.TerraformResourceLifecycle)
	Locality() *string
	SetLocality(val *string)
	LocalityInput() *string
	MaxPathLength() *float64
	SetMaxPathLength(val *float64)
	MaxPathLengthInput() *float64
	Namespace() *string
	SetNamespace(val *string)
	NamespaceInput() *string
	// The tree node.
	Node() constructs.Node
	Organization() *string
	SetOrganization(val *string)
	OrganizationInput() *string
	OtherSans() *[]*string
	SetOtherSans(val *[]*string)
	OtherSansInput() *[]*string
	Ou() *string
	SetOu(val *string)
	OuInput() *string
	PermittedDnsDomains() *[]*string
	SetPermittedDnsDomains(val *[]*string)
	PermittedDnsDomainsInput() *[]*string
	PostalCode() *string
	SetPostalCode(val *string)
	PostalCodeInput() *string
	PrivateKeyFormat() *string
	SetPrivateKeyFormat(val *string)
	PrivateKeyFormatInput() *string
	// Experimental.
	Provider() cdktf.TerraformProvider
	// Experimental.
	SetProvider(val cdktf.TerraformProvider)
	Province() *string
	SetProvince(val *string)
	ProvinceInput() *string
	// Experimental.
	Provisioners() *[]interface{}
	// Experimental.
	SetProvisioners(val *[]interface{})
	// Experimental.
	RawOverrides() interface{}
	Serial() *string
	SerialNumber() *string
	StreetAddress() *string
	SetStreetAddress(val *string)
	StreetAddressInput() *string
	// Experimental.
	TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata
	// Experimental.
	TerraformMetaArguments() *map[string]interface{}
	// Experimental.
	TerraformResourceType() *string
	Ttl() *string
	SetTtl(val *string)
	TtlInput() *string
	Type() *string
	SetType(val *string)
	TypeInput() *string
	UriSans() *[]*string
	SetUriSans(val *[]*string)
	UriSansInput() *[]*string
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
	ResetAltNames()
	ResetCountry()
	ResetExcludeCnFromSans()
	ResetFormat()
	ResetId()
	ResetIpSans()
	ResetKeyBits()
	ResetKeyType()
	ResetLocality()
	ResetMaxPathLength()
	ResetNamespace()
	ResetOrganization()
	ResetOtherSans()
	ResetOu()
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	ResetPermittedDnsDomains()
	ResetPostalCode()
	ResetPrivateKeyFormat()
	ResetProvince()
	ResetStreetAddress()
	ResetTtl()
	ResetUriSans()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Adds this resource to the terraform JSON output.
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for PkiSecretBackendRootCert
type jsiiProxy_PkiSecretBackendRootCert struct {
	internal.Type__cdktfTerraformResource
}

func (j *jsiiProxy_PkiSecretBackendRootCert) AltNames() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"altNames",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) AltNamesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"altNamesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Backend() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backend",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) BackendInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"backendInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Certificate() *string {
	var returns *string
	_jsii_.Get(
		j,
		"certificate",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) CommonName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) CommonNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"commonNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Connection() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"connection",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Count() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Country() *string {
	var returns *string
	_jsii_.Get(
		j,
		"country",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) CountryInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countryInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludeCnFromSans() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCnFromSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ExcludeCnFromSansInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"excludeCnFromSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Format() *string {
	var returns *string
	_jsii_.Get(
		j,
		"format",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) FormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"formatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Id() *string {
	var returns *string
	_jsii_.Get(
		j,
		"id",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"idInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IpSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IpSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"ipSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) IssuingCa() *string {
	var returns *string
	_jsii_.Get(
		j,
		"issuingCa",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyBits() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyBits",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyBitsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"keyBitsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) KeyTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Lifecycle() *cdktf.TerraformResourceLifecycle {
	var returns *cdktf.TerraformResourceLifecycle
	_jsii_.Get(
		j,
		"lifecycle",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Locality() *string {
	var returns *string
	_jsii_.Get(
		j,
		"locality",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) LocalityInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"localityInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) MaxPathLength() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPathLength",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) MaxPathLengthInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxPathLengthInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Namespace() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespace",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) NamespaceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"namespaceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Organization() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organization",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) OrganizationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"organizationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) OtherSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) OtherSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"otherSansInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Ou() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ou",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) OuInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ouInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PermittedDnsDomains() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsDomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PermittedDnsDomainsInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"permittedDnsDomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PostalCode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PostalCodeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"postalCodeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PrivateKeyFormat() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFormat",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) PrivateKeyFormatInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"privateKeyFormatInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Provider() cdktf.TerraformProvider {
	var returns cdktf.TerraformProvider
	_jsii_.Get(
		j,
		"provider",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Province() *string {
	var returns *string
	_jsii_.Get(
		j,
		"province",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) ProvinceInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"provinceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Provisioners() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"provisioners",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Serial() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serial",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SerialNumber() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serialNumber",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) StreetAddress() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddress",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) StreetAddressInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"streetAddressInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) TerraformGeneratorMetadata() *cdktf.TerraformProviderGeneratorMetadata {
	var returns *cdktf.TerraformProviderGeneratorMetadata
	_jsii_.Get(
		j,
		"terraformGeneratorMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) TerraformMetaArguments() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"terraformMetaArguments",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) TerraformResourceType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformResourceType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Ttl() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttl",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) TtlInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ttlInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) Type() *string {
	var returns *string
	_jsii_.Get(
		j,
		"type",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) TypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"typeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) UriSans() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uriSans",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PkiSecretBackendRootCert) UriSansInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"uriSansInput",
		&returns,
	)
	return returns
}


// Create a new {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert vault_pki_secret_backend_root_cert} Resource.
func NewPkiSecretBackendRootCert(scope constructs.Construct, id *string, config *PkiSecretBackendRootCertConfig) PkiSecretBackendRootCert {
	_init_.Initialize()

	j := jsiiProxy_PkiSecretBackendRootCert{}

	_jsii_.Create(
		"@cdktf/provider-vault.PkiSecretBackendRootCert",
		[]interface{}{scope, id, config},
		&j,
	)

	return &j
}

// Create a new {@link https://www.terraform.io/docs/providers/vault/r/pki_secret_backend_root_cert vault_pki_secret_backend_root_cert} Resource.
func NewPkiSecretBackendRootCert_Override(p PkiSecretBackendRootCert, scope constructs.Construct, id *string, config *PkiSecretBackendRootCertConfig) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-vault.PkiSecretBackendRootCert",
		[]interface{}{scope, id, config},
		p,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetAltNames(val *[]*string) {
	_jsii_.Set(
		j,
		"altNames",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetBackend(val *string) {
	_jsii_.Set(
		j,
		"backend",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetCommonName(val *string) {
	_jsii_.Set(
		j,
		"commonName",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetConnection(val interface{}) {
	_jsii_.Set(
		j,
		"connection",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetCount(val *float64) {
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetCountry(val *string) {
	_jsii_.Set(
		j,
		"country",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetExcludeCnFromSans(val interface{}) {
	_jsii_.Set(
		j,
		"excludeCnFromSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetFormat(val *string) {
	_jsii_.Set(
		j,
		"format",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetId(val *string) {
	_jsii_.Set(
		j,
		"id",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetIpSans(val *[]*string) {
	_jsii_.Set(
		j,
		"ipSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetKeyBits(val *float64) {
	_jsii_.Set(
		j,
		"keyBits",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetKeyType(val *string) {
	_jsii_.Set(
		j,
		"keyType",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetLifecycle(val *cdktf.TerraformResourceLifecycle) {
	_jsii_.Set(
		j,
		"lifecycle",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetLocality(val *string) {
	_jsii_.Set(
		j,
		"locality",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetMaxPathLength(val *float64) {
	_jsii_.Set(
		j,
		"maxPathLength",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetNamespace(val *string) {
	_jsii_.Set(
		j,
		"namespace",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetOrganization(val *string) {
	_jsii_.Set(
		j,
		"organization",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetOtherSans(val *[]*string) {
	_jsii_.Set(
		j,
		"otherSans",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetOu(val *string) {
	_jsii_.Set(
		j,
		"ou",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetPermittedDnsDomains(val *[]*string) {
	_jsii_.Set(
		j,
		"permittedDnsDomains",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetPostalCode(val *string) {
	_jsii_.Set(
		j,
		"postalCode",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetPrivateKeyFormat(val *string) {
	_jsii_.Set(
		j,
		"privateKeyFormat",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetProvider(val cdktf.TerraformProvider) {
	_jsii_.Set(
		j,
		"provider",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetProvince(val *string) {
	_jsii_.Set(
		j,
		"province",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetProvisioners(val *[]interface{}) {
	_jsii_.Set(
		j,
		"provisioners",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetStreetAddress(val *string) {
	_jsii_.Set(
		j,
		"streetAddress",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetTtl(val *string) {
	_jsii_.Set(
		j,
		"ttl",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetType(val *string) {
	_jsii_.Set(
		j,
		"type",
		val,
	)
}

func (j *jsiiProxy_PkiSecretBackendRootCert) SetUriSans(val *[]*string) {
	_jsii_.Set(
		j,
		"uriSans",
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
func PkiSecretBackendRootCert_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	var returns *bool

	_jsii_.StaticInvoke(
		"@cdktf/provider-vault.PkiSecretBackendRootCert",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func PkiSecretBackendRootCert_TfResourceType() *string {
	_init_.Initialize()
	var returns *string
	_jsii_.StaticGet(
		"@cdktf/provider-vault.PkiSecretBackendRootCert",
		"tfResourceType",
		&returns,
	)
	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) AddOverride(path *string, value interface{}) {
	_jsii_.InvokeVoid(
		p,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) GetListAttribute(terraformAttribute *string) *[]*string {
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) GetNumberAttribute(terraformAttribute *string) *float64 {
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) GetStringAttribute(terraformAttribute *string) *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) InterpolationForAttribute(terraformAttribute *string) cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) OverrideLogicalId(newLogicalId *string) {
	_jsii_.InvokeVoid(
		p,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetAltNames() {
	_jsii_.InvokeVoid(
		p,
		"resetAltNames",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetCountry() {
	_jsii_.InvokeVoid(
		p,
		"resetCountry",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetExcludeCnFromSans() {
	_jsii_.InvokeVoid(
		p,
		"resetExcludeCnFromSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetId() {
	_jsii_.InvokeVoid(
		p,
		"resetId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetIpSans() {
	_jsii_.InvokeVoid(
		p,
		"resetIpSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetKeyBits() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyBits",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetKeyType() {
	_jsii_.InvokeVoid(
		p,
		"resetKeyType",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetLocality() {
	_jsii_.InvokeVoid(
		p,
		"resetLocality",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetMaxPathLength() {
	_jsii_.InvokeVoid(
		p,
		"resetMaxPathLength",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetNamespace() {
	_jsii_.InvokeVoid(
		p,
		"resetNamespace",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetOrganization() {
	_jsii_.InvokeVoid(
		p,
		"resetOrganization",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetOtherSans() {
	_jsii_.InvokeVoid(
		p,
		"resetOtherSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetOu() {
	_jsii_.InvokeVoid(
		p,
		"resetOu",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		p,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPermittedDnsDomains() {
	_jsii_.InvokeVoid(
		p,
		"resetPermittedDnsDomains",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPostalCode() {
	_jsii_.InvokeVoid(
		p,
		"resetPostalCode",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetPrivateKeyFormat() {
	_jsii_.InvokeVoid(
		p,
		"resetPrivateKeyFormat",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetProvince() {
	_jsii_.InvokeVoid(
		p,
		"resetProvince",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetStreetAddress() {
	_jsii_.InvokeVoid(
		p,
		"resetStreetAddress",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetTtl() {
	_jsii_.InvokeVoid(
		p,
		"resetTtl",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ResetUriSans() {
	_jsii_.InvokeVoid(
		p,
		"resetUriSans",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PkiSecretBackendRootCert) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		p,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PkiSecretBackendRootCert) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		p,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

