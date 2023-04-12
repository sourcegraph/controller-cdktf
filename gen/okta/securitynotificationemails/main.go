package securitynotificationemails

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.securityNotificationEmails.SecurityNotificationEmails",
		reflect.TypeOf((*SecurityNotificationEmails)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
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
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "reportSuspiciousActivityEnabled", GoGetter: "ReportSuspiciousActivityEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "reportSuspiciousActivityEnabledInput", GoGetter: "ReportSuspiciousActivityEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetReportSuspiciousActivityEnabled", GoMethod: "ResetReportSuspiciousActivityEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetSendEmailForFactorEnrollmentEnabled", GoMethod: "ResetSendEmailForFactorEnrollmentEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetSendEmailForFactorResetEnabled", GoMethod: "ResetSendEmailForFactorResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetSendEmailForNewDeviceEnabled", GoMethod: "ResetSendEmailForNewDeviceEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetSendEmailForPasswordChangedEnabled", GoMethod: "ResetSendEmailForPasswordChangedEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "sendEmailForFactorEnrollmentEnabled", GoGetter: "SendEmailForFactorEnrollmentEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "sendEmailForFactorEnrollmentEnabledInput", GoGetter: "SendEmailForFactorEnrollmentEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "sendEmailForFactorResetEnabled", GoGetter: "SendEmailForFactorResetEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "sendEmailForFactorResetEnabledInput", GoGetter: "SendEmailForFactorResetEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "sendEmailForNewDeviceEnabled", GoGetter: "SendEmailForNewDeviceEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "sendEmailForNewDeviceEnabledInput", GoGetter: "SendEmailForNewDeviceEnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "sendEmailForPasswordChangedEnabled", GoGetter: "SendEmailForPasswordChangedEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "sendEmailForPasswordChangedEnabledInput", GoGetter: "SendEmailForPasswordChangedEnabledInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_SecurityNotificationEmails{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.securityNotificationEmails.SecurityNotificationEmailsConfig",
		reflect.TypeOf((*SecurityNotificationEmailsConfig)(nil)).Elem(),
	)
}
