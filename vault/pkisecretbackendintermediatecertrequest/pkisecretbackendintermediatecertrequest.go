package pkisecretbackendintermediatecertrequest

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-vault.pkiSecretBackendIntermediateCertRequest.PkiSecretBackendIntermediateCertRequest",
		reflect.TypeOf((*PkiSecretBackendIntermediateCertRequest)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "altNames", GoGetter: "AltNames"},
			_jsii_.MemberProperty{JsiiProperty: "altNamesInput", GoGetter: "AltNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "backend", GoGetter: "Backend"},
			_jsii_.MemberProperty{JsiiProperty: "backendInput", GoGetter: "BackendInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "commonName", GoGetter: "CommonName"},
			_jsii_.MemberProperty{JsiiProperty: "commonNameInput", GoGetter: "CommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "country", GoGetter: "Country"},
			_jsii_.MemberProperty{JsiiProperty: "countryInput", GoGetter: "CountryInput"},
			_jsii_.MemberProperty{JsiiProperty: "csr", GoGetter: "Csr"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCnFromSans", GoGetter: "ExcludeCnFromSans"},
			_jsii_.MemberProperty{JsiiProperty: "excludeCnFromSansInput", GoGetter: "ExcludeCnFromSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "format", GoGetter: "Format"},
			_jsii_.MemberProperty{JsiiProperty: "formatInput", GoGetter: "FormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipSans", GoGetter: "IpSans"},
			_jsii_.MemberProperty{JsiiProperty: "ipSansInput", GoGetter: "IpSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyBits", GoGetter: "KeyBits"},
			_jsii_.MemberProperty{JsiiProperty: "keyBitsInput", GoGetter: "KeyBitsInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyType", GoGetter: "KeyType"},
			_jsii_.MemberProperty{JsiiProperty: "keyTypeInput", GoGetter: "KeyTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "locality", GoGetter: "Locality"},
			_jsii_.MemberProperty{JsiiProperty: "localityInput", GoGetter: "LocalityInput"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberProperty{JsiiProperty: "namespaceInput", GoGetter: "NamespaceInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "organization", GoGetter: "Organization"},
			_jsii_.MemberProperty{JsiiProperty: "organizationInput", GoGetter: "OrganizationInput"},
			_jsii_.MemberProperty{JsiiProperty: "otherSans", GoGetter: "OtherSans"},
			_jsii_.MemberProperty{JsiiProperty: "otherSansInput", GoGetter: "OtherSansInput"},
			_jsii_.MemberProperty{JsiiProperty: "ou", GoGetter: "Ou"},
			_jsii_.MemberProperty{JsiiProperty: "ouInput", GoGetter: "OuInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "postalCode", GoGetter: "PostalCode"},
			_jsii_.MemberProperty{JsiiProperty: "postalCodeInput", GoGetter: "PostalCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKey", GoGetter: "PrivateKey"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyFormat", GoGetter: "PrivateKeyFormat"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyFormatInput", GoGetter: "PrivateKeyFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyType", GoGetter: "PrivateKeyType"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "province", GoGetter: "Province"},
			_jsii_.MemberProperty{JsiiProperty: "provinceInput", GoGetter: "ProvinceInput"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAltNames", GoMethod: "ResetAltNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetCountry", GoMethod: "ResetCountry"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeCnFromSans", GoMethod: "ResetExcludeCnFromSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetFormat", GoMethod: "ResetFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpSans", GoMethod: "ResetIpSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyBits", GoMethod: "ResetKeyBits"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyType", GoMethod: "ResetKeyType"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocality", GoMethod: "ResetLocality"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespace", GoMethod: "ResetNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrganization", GoMethod: "ResetOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "resetOtherSans", GoMethod: "ResetOtherSans"},
			_jsii_.MemberMethod{JsiiMethod: "resetOu", GoMethod: "ResetOu"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostalCode", GoMethod: "ResetPostalCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateKeyFormat", GoMethod: "ResetPrivateKeyFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvince", GoMethod: "ResetProvince"},
			_jsii_.MemberMethod{JsiiMethod: "resetStreetAddress", GoMethod: "ResetStreetAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetUriSans", GoMethod: "ResetUriSans"},
			_jsii_.MemberProperty{JsiiProperty: "streetAddress", GoGetter: "StreetAddress"},
			_jsii_.MemberProperty{JsiiProperty: "streetAddressInput", GoGetter: "StreetAddressInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "typeInput", GoGetter: "TypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "uriSans", GoGetter: "UriSans"},
			_jsii_.MemberProperty{JsiiProperty: "uriSansInput", GoGetter: "UriSansInput"},
		},
		func() interface{} {
			j := jsiiProxy_PkiSecretBackendIntermediateCertRequest{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-vault.pkiSecretBackendIntermediateCertRequest.PkiSecretBackendIntermediateCertRequestConfig",
		reflect.TypeOf((*PkiSecretBackendIntermediateCertRequestConfig)(nil)).Elem(),
	)
}
