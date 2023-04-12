package authenticator

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.authenticator.Authenticator",
		reflect.TypeOf((*Authenticator)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "providerAuthPort", GoGetter: "ProviderAuthPort"},
			_jsii_.MemberProperty{JsiiProperty: "providerAuthPortInput", GoGetter: "ProviderAuthPortInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerHost", GoGetter: "ProviderHost"},
			_jsii_.MemberProperty{JsiiProperty: "providerHostInput", GoGetter: "ProviderHostInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerHostname", GoGetter: "ProviderHostname"},
			_jsii_.MemberProperty{JsiiProperty: "providerHostnameInput", GoGetter: "ProviderHostnameInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerInstanceId", GoGetter: "ProviderInstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "providerIntegrationKey", GoGetter: "ProviderIntegrationKey"},
			_jsii_.MemberProperty{JsiiProperty: "providerIntegrationKeyInput", GoGetter: "ProviderIntegrationKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerJson", GoGetter: "ProviderJson"},
			_jsii_.MemberProperty{JsiiProperty: "providerJsonInput", GoGetter: "ProviderJsonInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerSecretKey", GoGetter: "ProviderSecretKey"},
			_jsii_.MemberProperty{JsiiProperty: "providerSecretKeyInput", GoGetter: "ProviderSecretKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerSharedSecret", GoGetter: "ProviderSharedSecret"},
			_jsii_.MemberProperty{JsiiProperty: "providerSharedSecretInput", GoGetter: "ProviderSharedSecretInput"},
			_jsii_.MemberProperty{JsiiProperty: "providerType", GoGetter: "ProviderType"},
			_jsii_.MemberProperty{JsiiProperty: "providerUserNameTemplate", GoGetter: "ProviderUserNameTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "providerUserNameTemplateInput", GoGetter: "ProviderUserNameTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderAuthPort", GoMethod: "ResetProviderAuthPort"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderHost", GoMethod: "ResetProviderHost"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderHostname", GoMethod: "ResetProviderHostname"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderIntegrationKey", GoMethod: "ResetProviderIntegrationKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderJson", GoMethod: "ResetProviderJson"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderSecretKey", GoMethod: "ResetProviderSecretKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderSharedSecret", GoMethod: "ResetProviderSharedSecret"},
			_jsii_.MemberMethod{JsiiMethod: "resetProviderUserNameTemplate", GoMethod: "ResetProviderUserNameTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resetSettings", GoMethod: "ResetSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetStatus", GoMethod: "ResetStatus"},
			_jsii_.MemberProperty{JsiiProperty: "settings", GoGetter: "Settings"},
			_jsii_.MemberProperty{JsiiProperty: "settingsInput", GoGetter: "SettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "status", GoGetter: "Status"},
			_jsii_.MemberProperty{JsiiProperty: "statusInput", GoGetter: "StatusInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "type", GoGetter: "Type"},
		},
		func() interface{} {
			j := jsiiProxy_Authenticator{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.authenticator.AuthenticatorConfig",
		reflect.TypeOf((*AuthenticatorConfig)(nil)).Elem(),
	)
}
