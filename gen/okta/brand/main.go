package brand

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.brand.Brand",
		reflect.TypeOf((*Brand)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "agreeToCustomPrivacyPolicy", GoGetter: "AgreeToCustomPrivacyPolicy"},
			_jsii_.MemberProperty{JsiiProperty: "agreeToCustomPrivacyPolicyInput", GoGetter: "AgreeToCustomPrivacyPolicyInput"},
			_jsii_.MemberProperty{JsiiProperty: "brandId", GoGetter: "BrandId"},
			_jsii_.MemberProperty{JsiiProperty: "brandIdInput", GoGetter: "BrandIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "customPrivacyPolicyUrl", GoGetter: "CustomPrivacyPolicyUrl"},
			_jsii_.MemberProperty{JsiiProperty: "customPrivacyPolicyUrlInput", GoGetter: "CustomPrivacyPolicyUrlInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAppAppInstanceId", GoGetter: "DefaultAppAppInstanceId"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAppAppInstanceIdInput", GoGetter: "DefaultAppAppInstanceIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAppAppLinkName", GoGetter: "DefaultAppAppLinkName"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAppAppLinkNameInput", GoGetter: "DefaultAppAppLinkNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAppClassicApplicationUri", GoGetter: "DefaultAppClassicApplicationUri"},
			_jsii_.MemberProperty{JsiiProperty: "defaultAppClassicApplicationUriInput", GoGetter: "DefaultAppClassicApplicationUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "emailDomainId", GoGetter: "EmailDomainId"},
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
			_jsii_.MemberProperty{JsiiProperty: "isDefault", GoGetter: "IsDefault"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "links", GoGetter: "Links"},
			_jsii_.MemberProperty{JsiiProperty: "locale", GoGetter: "Locale"},
			_jsii_.MemberProperty{JsiiProperty: "localeInput", GoGetter: "LocaleInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberProperty{JsiiProperty: "removePoweredByOkta", GoGetter: "RemovePoweredByOkta"},
			_jsii_.MemberProperty{JsiiProperty: "removePoweredByOktaInput", GoGetter: "RemovePoweredByOktaInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAgreeToCustomPrivacyPolicy", GoMethod: "ResetAgreeToCustomPrivacyPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resetBrandId", GoMethod: "ResetBrandId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomPrivacyPolicyUrl", GoMethod: "ResetCustomPrivacyPolicyUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultAppAppInstanceId", GoMethod: "ResetDefaultAppAppInstanceId"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultAppAppLinkName", GoMethod: "ResetDefaultAppAppLinkName"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultAppClassicApplicationUri", GoMethod: "ResetDefaultAppClassicApplicationUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocale", GoMethod: "ResetLocale"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetRemovePoweredByOkta", GoMethod: "ResetRemovePoweredByOkta"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_Brand{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.brand.BrandConfig",
		reflect.TypeOf((*BrandConfig)(nil)).Elem(),
	)
}
