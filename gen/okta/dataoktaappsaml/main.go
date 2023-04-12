package dataoktaappsaml

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaAppSaml.DataOktaAppSaml",
		reflect.TypeOf((*DataOktaAppSaml)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessibilityErrorRedirectUrl", GoGetter: "AccessibilityErrorRedirectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilityLoginRedirectUrl", GoGetter: "AccessibilityLoginRedirectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilitySelfService", GoGetter: "AccessibilitySelfService"},
			_jsii_.MemberProperty{JsiiProperty: "acsEndpoints", GoGetter: "AcsEndpoints"},
			_jsii_.MemberProperty{JsiiProperty: "activeOnly", GoGetter: "ActiveOnly"},
			_jsii_.MemberProperty{JsiiProperty: "activeOnlyInput", GoGetter: "ActiveOnlyInput"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "appSettingsJson", GoGetter: "AppSettingsJson"},
			_jsii_.MemberProperty{JsiiProperty: "assertionSigned", GoGetter: "AssertionSigned"},
			_jsii_.MemberProperty{JsiiProperty: "attributeStatements", GoGetter: "AttributeStatements"},
			_jsii_.MemberProperty{JsiiProperty: "audience", GoGetter: "Audience"},
			_jsii_.MemberProperty{JsiiProperty: "authnContextClassRef", GoGetter: "AuthnContextClassRef"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubmitToolbar", GoGetter: "AutoSubmitToolbar"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "defaultRelayState", GoGetter: "DefaultRelayState"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "destination", GoGetter: "Destination"},
			_jsii_.MemberProperty{JsiiProperty: "digestAlgorithm", GoGetter: "DigestAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "features", GoGetter: "Features"},
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
			_jsii_.MemberProperty{JsiiProperty: "groups", GoGetter: "Groups"},
			_jsii_.MemberProperty{JsiiProperty: "hideIos", GoGetter: "HideIos"},
			_jsii_.MemberProperty{JsiiProperty: "hideWeb", GoGetter: "HideWeb"},
			_jsii_.MemberProperty{JsiiProperty: "honorForceAuthn", GoGetter: "HonorForceAuthn"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberProperty{JsiiProperty: "idpIssuer", GoGetter: "IdpIssuer"},
			_jsii_.MemberProperty{JsiiProperty: "inlineHookId", GoGetter: "InlineHookId"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "keyId", GoGetter: "KeyId"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "labelPrefix", GoGetter: "LabelPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "labelPrefixInput", GoGetter: "LabelPrefixInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "links", GoGetter: "Links"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "recipient", GoGetter: "Recipient"},
			_jsii_.MemberProperty{JsiiProperty: "requestCompressed", GoGetter: "RequestCompressed"},
			_jsii_.MemberProperty{JsiiProperty: "requestCompressedInput", GoGetter: "RequestCompressedInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetActiveOnly", GoMethod: "ResetActiveOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabel", GoMethod: "ResetLabel"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabelPrefix", GoMethod: "ResetLabelPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequestCompressed", GoMethod: "ResetRequestCompressed"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipGroups", GoMethod: "ResetSkipGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetSkipUsers", GoMethod: "ResetSkipUsers"},
			_jsii_.MemberProperty{JsiiProperty: "responseSigned", GoGetter: "ResponseSigned"},
			_jsii_.MemberProperty{JsiiProperty: "samlSignedRequestEnabled", GoGetter: "SamlSignedRequestEnabled"},
			_jsii_.MemberProperty{JsiiProperty: "signatureAlgorithm", GoGetter: "SignatureAlgorithm"},
			_jsii_.MemberProperty{JsiiProperty: "singleLogoutCertificate", GoGetter: "SingleLogoutCertificate"},
			_jsii_.MemberProperty{JsiiProperty: "singleLogoutIssuer", GoGetter: "SingleLogoutIssuer"},
			_jsii_.MemberProperty{JsiiProperty: "singleLogoutUrl", GoGetter: "SingleLogoutUrl"},
			_jsii_.MemberProperty{JsiiProperty: "skipGroups", GoGetter: "SkipGroups"},
			_jsii_.MemberProperty{JsiiProperty: "skipGroupsInput", GoGetter: "SkipGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "skipUsers", GoGetter: "SkipUsers"},
			_jsii_.MemberProperty{JsiiProperty: "skipUsersInput", GoGetter: "SkipUsersInput"},
			_jsii_.MemberProperty{JsiiProperty: "spIssuer", GoGetter: "SpIssuer"},
			_jsii_.MemberProperty{JsiiProperty: "ssoUrl", GoGetter: "SsoUrl"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "subjectNameIdFormat", GoGetter: "SubjectNameIdFormat"},
			_jsii_.MemberProperty{JsiiProperty: "subjectNameIdTemplate", GoGetter: "SubjectNameIdTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplate", GoGetter: "UserNameTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplatePushStatus", GoGetter: "UserNameTemplatePushStatus"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateSuffix", GoGetter: "UserNameTemplateSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateType", GoGetter: "UserNameTemplateType"},
			_jsii_.MemberProperty{JsiiProperty: "users", GoGetter: "Users"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaAppSaml{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformDataSource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaAppSaml.DataOktaAppSamlAttributeStatements",
		reflect.TypeOf((*DataOktaAppSamlAttributeStatements)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaAppSaml.DataOktaAppSamlAttributeStatementsList",
		reflect.TypeOf((*DataOktaAppSamlAttributeStatementsList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaAppSamlAttributeStatementsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.dataOktaAppSaml.DataOktaAppSamlAttributeStatementsOutputReference",
		reflect.TypeOf((*DataOktaAppSamlAttributeStatementsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "filterType", GoGetter: "FilterType"},
			_jsii_.MemberProperty{JsiiProperty: "filterValue", GoGetter: "FilterValue"},
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "namespace", GoGetter: "Namespace"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
		},
		func() interface{} {
			j := jsiiProxy_DataOktaAppSamlAttributeStatementsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.dataOktaAppSaml.DataOktaAppSamlConfig",
		reflect.TypeOf((*DataOktaAppSamlConfig)(nil)).Elem(),
	)
}
