package theme

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-okta.theme.Theme",
		reflect.TypeOf((*Theme)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "backgroundImage", GoGetter: "BackgroundImage"},
			_jsii_.MemberProperty{JsiiProperty: "backgroundImageInput", GoGetter: "BackgroundImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "backgroundImageUrl", GoGetter: "BackgroundImageUrl"},
			_jsii_.MemberProperty{JsiiProperty: "brandId", GoGetter: "BrandId"},
			_jsii_.MemberProperty{JsiiProperty: "brandIdInput", GoGetter: "BrandIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "emailTemplateTouchPointVariant", GoGetter: "EmailTemplateTouchPointVariant"},
			_jsii_.MemberProperty{JsiiProperty: "emailTemplateTouchPointVariantInput", GoGetter: "EmailTemplateTouchPointVariantInput"},
			_jsii_.MemberProperty{JsiiProperty: "endUserDashboardTouchPointVariant", GoGetter: "EndUserDashboardTouchPointVariant"},
			_jsii_.MemberProperty{JsiiProperty: "endUserDashboardTouchPointVariantInput", GoGetter: "EndUserDashboardTouchPointVariantInput"},
			_jsii_.MemberProperty{JsiiProperty: "errorPageTouchPointVariant", GoGetter: "ErrorPageTouchPointVariant"},
			_jsii_.MemberProperty{JsiiProperty: "errorPageTouchPointVariantInput", GoGetter: "ErrorPageTouchPointVariantInput"},
			_jsii_.MemberProperty{JsiiProperty: "favicon", GoGetter: "Favicon"},
			_jsii_.MemberProperty{JsiiProperty: "faviconInput", GoGetter: "FaviconInput"},
			_jsii_.MemberProperty{JsiiProperty: "faviconUrl", GoGetter: "FaviconUrl"},
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
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "links", GoGetter: "Links"},
			_jsii_.MemberProperty{JsiiProperty: "logo", GoGetter: "Logo"},
			_jsii_.MemberProperty{JsiiProperty: "logoInput", GoGetter: "LogoInput"},
			_jsii_.MemberProperty{JsiiProperty: "logoUrl", GoGetter: "LogoUrl"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "primaryColorContrastHex", GoGetter: "PrimaryColorContrastHex"},
			_jsii_.MemberProperty{JsiiProperty: "primaryColorContrastHexInput", GoGetter: "PrimaryColorContrastHexInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryColorHex", GoGetter: "PrimaryColorHex"},
			_jsii_.MemberProperty{JsiiProperty: "primaryColorHexInput", GoGetter: "PrimaryColorHexInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackgroundImage", GoMethod: "ResetBackgroundImage"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailTemplateTouchPointVariant", GoMethod: "ResetEmailTemplateTouchPointVariant"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndUserDashboardTouchPointVariant", GoMethod: "ResetEndUserDashboardTouchPointVariant"},
			_jsii_.MemberMethod{JsiiMethod: "resetErrorPageTouchPointVariant", GoMethod: "ResetErrorPageTouchPointVariant"},
			_jsii_.MemberMethod{JsiiMethod: "resetFavicon", GoMethod: "ResetFavicon"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogo", GoMethod: "ResetLogo"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryColorContrastHex", GoMethod: "ResetPrimaryColorContrastHex"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryColorHex", GoMethod: "ResetPrimaryColorHex"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecondaryColorContrastHex", GoMethod: "ResetSecondaryColorContrastHex"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecondaryColorHex", GoMethod: "ResetSecondaryColorHex"},
			_jsii_.MemberMethod{JsiiMethod: "resetSignInPageTouchPointVariant", GoMethod: "ResetSignInPageTouchPointVariant"},
			_jsii_.MemberMethod{JsiiMethod: "resetThemeId", GoMethod: "ResetThemeId"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryColorContrastHex", GoGetter: "SecondaryColorContrastHex"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryColorContrastHexInput", GoGetter: "SecondaryColorContrastHexInput"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryColorHex", GoGetter: "SecondaryColorHex"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryColorHexInput", GoGetter: "SecondaryColorHexInput"},
			_jsii_.MemberProperty{JsiiProperty: "signInPageTouchPointVariant", GoGetter: "SignInPageTouchPointVariant"},
			_jsii_.MemberProperty{JsiiProperty: "signInPageTouchPointVariantInput", GoGetter: "SignInPageTouchPointVariantInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "themeId", GoGetter: "ThemeId"},
			_jsii_.MemberProperty{JsiiProperty: "themeIdInput", GoGetter: "ThemeIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
		},
		func() interface{} {
			j := jsiiProxy_Theme{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-okta.theme.ThemeConfig",
		reflect.TypeOf((*ThemeConfig)(nil)).Elem(),
	)
}
