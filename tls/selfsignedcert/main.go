// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package selfsignedcert

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-tls.selfSignedCert.SelfSignedCert",
		reflect.TypeOf((*SelfSignedCert)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUses", GoGetter: "AllowedUses"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUsesInput", GoGetter: "AllowedUsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certPem", GoGetter: "CertPem"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "dnsNames", GoGetter: "DnsNames"},
			_jsii_.MemberProperty{JsiiProperty: "dnsNamesInput", GoGetter: "DnsNamesInput"},
			_jsii_.MemberProperty{JsiiProperty: "earlyRenewalHours", GoGetter: "EarlyRenewalHours"},
			_jsii_.MemberProperty{JsiiProperty: "earlyRenewalHoursInput", GoGetter: "EarlyRenewalHoursInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddresses", GoGetter: "IpAddresses"},
			_jsii_.MemberProperty{JsiiProperty: "ipAddressesInput", GoGetter: "IpAddressesInput"},
			_jsii_.MemberProperty{JsiiProperty: "isCaCertificate", GoGetter: "IsCaCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "isCaCertificateInput", GoGetter: "IsCaCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyAlgorithm", GoGetter: "KeyAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyPem", GoGetter: "PrivateKeyPem"},
			_jsii_.MemberProperty{JsiiProperty: "privateKeyPemInput", GoGetter: "PrivateKeyPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putSubject", GoMethod: "PutSubject"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "readyForRenewal", GoGetter: "ReadyForRenewal"},
			_jsii_.MemberMethod{JsiiMethod: "resetDnsNames", GoMethod: "ResetDnsNames"},
			_jsii_.MemberMethod{JsiiMethod: "resetEarlyRenewalHours", GoMethod: "ResetEarlyRenewalHours"},
			_jsii_.MemberMethod{JsiiMethod: "resetIpAddresses", GoMethod: "ResetIpAddresses"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsCaCertificate", GoMethod: "ResetIsCaCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetAuthorityKeyId", GoMethod: "ResetSetAuthorityKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetSubjectKeyId", GoMethod: "ResetSetSubjectKeyId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubject", GoMethod: "ResetSubject"},
			_jsii_.MemberMethod{JsiiMethod: "resetUris", GoMethod: "ResetUris"},
			_jsii_.MemberProperty{JsiiProperty: "setAuthorityKeyId", GoGetter: "SetAuthorityKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "setAuthorityKeyIdInput", GoGetter: "SetAuthorityKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "setSubjectKeyId", GoGetter: "SetSubjectKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "setSubjectKeyIdInput", GoGetter: "SetSubjectKeyIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "subject", GoGetter: "Subject"},
			_jsii_.MemberProperty{JsiiProperty: "subjectInput", GoGetter: "SubjectInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "uris", GoGetter: "Uris"},
			_jsii_.MemberProperty{JsiiProperty: "urisInput", GoGetter: "UrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "validityEndTime", GoGetter: "ValidityEndTime"},
			_jsii_.MemberProperty{JsiiProperty: "validityPeriodHours", GoGetter: "ValidityPeriodHours"},
			_jsii_.MemberProperty{JsiiProperty: "validityPeriodHoursInput", GoGetter: "ValidityPeriodHoursInput"},
			_jsii_.MemberProperty{JsiiProperty: "validityStartTime", GoGetter: "ValidityStartTime"},
		},
		func() interface{} {
			j := jsiiProxy_SelfSignedCert{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-tls.selfSignedCert.SelfSignedCertConfig",
		reflect.TypeOf((*SelfSignedCertConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-tls.selfSignedCert.SelfSignedCertSubject",
		reflect.TypeOf((*SelfSignedCertSubject)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-tls.selfSignedCert.SelfSignedCertSubjectList",
		reflect.TypeOf((*SelfSignedCertSubjectList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "allWithMapKey", GoMethod: "AllWithMapKey"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_SelfSignedCertSubjectList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-tls.selfSignedCert.SelfSignedCertSubjectOutputReference",
		reflect.TypeOf((*SelfSignedCertSubjectOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "commonName", GoGetter: "CommonName"},
			_jsii_.MemberProperty{JsiiProperty: "commonNameInput", GoGetter: "CommonNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "country", GoGetter: "Country"},
			_jsii_.MemberProperty{JsiiProperty: "countryInput", GoGetter: "CountryInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "emailAddress", GoGetter: "EmailAddress"},
			_jsii_.MemberProperty{JsiiProperty: "emailAddressInput", GoGetter: "EmailAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "locality", GoGetter: "Locality"},
			_jsii_.MemberProperty{JsiiProperty: "localityInput", GoGetter: "LocalityInput"},
			_jsii_.MemberProperty{JsiiProperty: "organization", GoGetter: "Organization"},
			_jsii_.MemberProperty{JsiiProperty: "organizationalUnit", GoGetter: "OrganizationalUnit"},
			_jsii_.MemberProperty{JsiiProperty: "organizationalUnitInput", GoGetter: "OrganizationalUnitInput"},
			_jsii_.MemberProperty{JsiiProperty: "organizationInput", GoGetter: "OrganizationInput"},
			_jsii_.MemberProperty{JsiiProperty: "postalCode", GoGetter: "PostalCode"},
			_jsii_.MemberProperty{JsiiProperty: "postalCodeInput", GoGetter: "PostalCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "province", GoGetter: "Province"},
			_jsii_.MemberProperty{JsiiProperty: "provinceInput", GoGetter: "ProvinceInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCommonName", GoMethod: "ResetCommonName"},
			_jsii_.MemberMethod{JsiiMethod: "resetCountry", GoMethod: "ResetCountry"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailAddress", GoMethod: "ResetEmailAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocality", GoMethod: "ResetLocality"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrganization", GoMethod: "ResetOrganization"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrganizationalUnit", GoMethod: "ResetOrganizationalUnit"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostalCode", GoMethod: "ResetPostalCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetProvince", GoMethod: "ResetProvince"},
			_jsii_.MemberMethod{JsiiMethod: "resetSerialNumber", GoMethod: "ResetSerialNumber"},
			_jsii_.MemberMethod{JsiiMethod: "resetStreetAddress", GoMethod: "ResetStreetAddress"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serialNumber", GoGetter: "SerialNumber"},
			_jsii_.MemberProperty{JsiiProperty: "serialNumberInput", GoGetter: "SerialNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "streetAddress", GoGetter: "StreetAddress"},
			_jsii_.MemberProperty{JsiiProperty: "streetAddressInput", GoGetter: "StreetAddressInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_SelfSignedCertSubjectOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
