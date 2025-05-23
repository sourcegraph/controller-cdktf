package appsecurepasswordstore

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.appSecurePasswordStore.AppSecurePasswordStore",
		reflect.TypeOf((*AppSecurePasswordStore)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessibilityErrorRedirectUrl", GoGetter: "AccessibilityErrorRedirectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilityErrorRedirectUrlInput", GoGetter: "AccessibilityErrorRedirectUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilityLoginRedirectUrl", GoGetter: "AccessibilityLoginRedirectUrl"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilityLoginRedirectUrlInput", GoGetter: "AccessibilityLoginRedirectUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilitySelfService", GoGetter: "AccessibilitySelfService"},
			_jsii_.MemberProperty{JsiiProperty: "accessibilitySelfServiceInput", GoGetter: "AccessibilitySelfServiceInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "adminNote", GoGetter: "AdminNote"},
			_jsii_.MemberProperty{JsiiProperty: "adminNoteInput", GoGetter: "AdminNoteInput"},
			_jsii_.MemberProperty{JsiiProperty: "appLinksJson", GoGetter: "AppLinksJson"},
			_jsii_.MemberProperty{JsiiProperty: "appLinksJsonInput", GoGetter: "AppLinksJsonInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubmitToolbar", GoGetter: "AutoSubmitToolbar"},
			_jsii_.MemberProperty{JsiiProperty: "autoSubmitToolbarInput", GoGetter: "AutoSubmitToolbarInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsScheme", GoGetter: "CredentialsScheme"},
			_jsii_.MemberProperty{JsiiProperty: "credentialsSchemeInput", GoGetter: "CredentialsSchemeInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enduserNote", GoGetter: "EnduserNote"},
			_jsii_.MemberProperty{JsiiProperty: "enduserNoteInput", GoGetter: "EnduserNoteInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "hideIos", GoGetter: "HideIos"},
			_jsii_.MemberProperty{JsiiProperty: "hideIosInput", GoGetter: "HideIosInput"},
			_jsii_.MemberProperty{JsiiProperty: "hideWeb", GoGetter: "HideWeb"},
			_jsii_.MemberProperty{JsiiProperty: "hideWebInput", GoGetter: "HideWebInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "label", GoGetter: "Label"},
			_jsii_.MemberProperty{JsiiProperty: "labelInput", GoGetter: "LabelInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "logo", GoGetter: "Logo"},
			_jsii_.MemberProperty{JsiiProperty: "logoInput", GoGetter: "LogoInput"},
			_jsii_.MemberProperty{JsiiProperty: "logoUrl", GoGetter: "LogoUrl"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField1", GoGetter: "OptionalField1"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField1Input", GoGetter: "OptionalField1Input"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField1Value", GoGetter: "OptionalField1Value"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField1ValueInput", GoGetter: "OptionalField1ValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField2", GoGetter: "OptionalField2"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField2Input", GoGetter: "OptionalField2Input"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField2Value", GoGetter: "OptionalField2Value"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField2ValueInput", GoGetter: "OptionalField2ValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField3", GoGetter: "OptionalField3"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField3Input", GoGetter: "OptionalField3Input"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField3Value", GoGetter: "OptionalField3Value"},
			_jsii_.MemberProperty{JsiiProperty: "optionalField3ValueInput", GoGetter: "OptionalField3ValueInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passwordField", GoGetter: "PasswordField"},
			_jsii_.MemberProperty{JsiiProperty: "passwordFieldInput", GoGetter: "PasswordFieldInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessibilityErrorRedirectUrl", GoMethod: "ResetAccessibilityErrorRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessibilityLoginRedirectUrl", GoMethod: "ResetAccessibilityLoginRedirectUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessibilitySelfService", GoMethod: "ResetAccessibilitySelfService"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdminNote", GoMethod: "ResetAdminNote"},
			_jsii_.MemberMethod{JsiiMethod: "resetAppLinksJson", GoMethod: "ResetAppLinksJson"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoSubmitToolbar", GoMethod: "ResetAutoSubmitToolbar"},
			_jsii_.MemberMethod{JsiiMethod: "resetCredentialsScheme", GoMethod: "ResetCredentialsScheme"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnduserNote", GoMethod: "ResetEnduserNote"},
			_jsii_.MemberMethod{JsiiMethod: "resetHideIos", GoMethod: "ResetHideIos"},
			_jsii_.MemberMethod{JsiiMethod: "resetHideWeb", GoMethod: "ResetHideWeb"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogo", GoMethod: "ResetLogo"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptionalField1", GoMethod: "ResetOptionalField1"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptionalField1Value", GoMethod: "ResetOptionalField1Value"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptionalField2", GoMethod: "ResetOptionalField2"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptionalField2Value", GoMethod: "ResetOptionalField2Value"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptionalField3", GoMethod: "ResetOptionalField3"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptionalField3Value", GoMethod: "ResetOptionalField3Value"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRevealPassword", GoMethod: "ResetRevealPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedPassword", GoMethod: "ResetSharedPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetSharedUsername", GoMethod: "ResetSharedUsername"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplate", GoMethod: "ResetUserNameTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplatePushStatus", GoMethod: "ResetUserNameTemplatePushStatus"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplateSuffix", GoMethod: "ResetUserNameTemplateSuffix"},
			_jsii_.MemberMethod{JsiiMethod: "resetUserNameTemplateType", GoMethod: "ResetUserNameTemplateType"},
			_jsii_.MemberProperty{JsiiProperty: "revealPassword", GoGetter: "RevealPassword"},
			_jsii_.MemberProperty{JsiiProperty: "revealPasswordInput", GoGetter: "RevealPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedPassword", GoGetter: "SharedPassword"},
			_jsii_.MemberProperty{JsiiProperty: "sharedPasswordInput", GoGetter: "SharedPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "sharedUsername", GoGetter: "SharedUsername"},
			_jsii_.MemberProperty{JsiiProperty: "sharedUsernameInput", GoGetter: "SharedUsernameInput"},
			_jsii_.MemberProperty{JsiiProperty: "signOnMode", GoGetter: "SignOnMode"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "usernameField", GoGetter: "UsernameField"},
			_jsii_.MemberProperty{JsiiProperty: "usernameFieldInput", GoGetter: "UsernameFieldInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplate", GoGetter: "UserNameTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateInput", GoGetter: "UserNameTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplatePushStatus", GoGetter: "UserNameTemplatePushStatus"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplatePushStatusInput", GoGetter: "UserNameTemplatePushStatusInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateSuffix", GoGetter: "UserNameTemplateSuffix"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateSuffixInput", GoGetter: "UserNameTemplateSuffixInput"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateType", GoGetter: "UserNameTemplateType"},
			_jsii_.MemberProperty{JsiiProperty: "userNameTemplateTypeInput", GoGetter: "UserNameTemplateTypeInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppSecurePasswordStore{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.appSecurePasswordStore.AppSecurePasswordStoreConfig",
		reflect.TypeOf((*AppSecurePasswordStoreConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.appSecurePasswordStore.AppSecurePasswordStoreTimeouts",
		reflect.TypeOf((*AppSecurePasswordStoreTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.appSecurePasswordStore.AppSecurePasswordStoreTimeoutsOutputReference",
		reflect.TypeOf((*AppSecurePasswordStoreTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
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
			_jsii_.MemberProperty{JsiiProperty: "read", GoGetter: "Read"},
			_jsii_.MemberProperty{JsiiProperty: "readInput", GoGetter: "ReadInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetRead", GoMethod: "ResetRead"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_AppSecurePasswordStoreTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
