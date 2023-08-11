package locallysignedcert

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-tls.locallySignedCert.LocallySignedCert",
		reflect.TypeOf((*LocallySignedCert)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUses", GoGetter: "AllowedUses"},
			_jsii_.MemberProperty{JsiiProperty: "allowedUsesInput", GoGetter: "AllowedUsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "caCertPem", GoGetter: "CaCertPem"},
			_jsii_.MemberProperty{JsiiProperty: "caCertPemInput", GoGetter: "CaCertPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "caKeyAlgorithm", GoGetter: "CaKeyAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "caPrivateKeyPem", GoGetter: "CaPrivateKeyPem"},
			_jsii_.MemberProperty{JsiiProperty: "caPrivateKeyPemInput", GoGetter: "CaPrivateKeyPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "certPem", GoGetter: "CertPem"},
			_jsii_.MemberProperty{JsiiProperty: "certRequestPem", GoGetter: "CertRequestPem"},
			_jsii_.MemberProperty{JsiiProperty: "certRequestPemInput", GoGetter: "CertRequestPemInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isCaCertificate", GoGetter: "IsCaCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "isCaCertificateInput", GoGetter: "IsCaCertificateInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "readyForRenewal", GoGetter: "ReadyForRenewal"},
			_jsii_.MemberMethod{JsiiMethod: "resetEarlyRenewalHours", GoMethod: "ResetEarlyRenewalHours"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsCaCertificate", GoMethod: "ResetIsCaCertificate"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetSubjectKeyId", GoMethod: "ResetSetSubjectKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "setSubjectKeyId", GoGetter: "SetSubjectKeyId"},
			_jsii_.MemberProperty{JsiiProperty: "setSubjectKeyIdInput", GoGetter: "SetSubjectKeyIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "validityEndTime", GoGetter: "ValidityEndTime"},
			_jsii_.MemberProperty{JsiiProperty: "validityPeriodHours", GoGetter: "ValidityPeriodHours"},
			_jsii_.MemberProperty{JsiiProperty: "validityPeriodHoursInput", GoGetter: "ValidityPeriodHoursInput"},
			_jsii_.MemberProperty{JsiiProperty: "validityStartTime", GoGetter: "ValidityStartTime"},
		},
		func() interface{} {
			j := jsiiProxy_LocallySignedCert{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-tls.locallySignedCert.LocallySignedCertConfig",
		reflect.TypeOf((*LocallySignedCertConfig)(nil)).Elem(),
	)
}
