package discoveryenginewidgetconfig

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfig",
		reflect.TypeOf((*DiscoveryEngineWidgetConfig)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accessSettings", GoGetter: "AccessSettings"},
			_jsii_.MemberProperty{JsiiProperty: "accessSettingsInput", GoGetter: "AccessSettingsInput"},
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "collectionId", GoGetter: "CollectionId"},
			_jsii_.MemberProperty{JsiiProperty: "collectionIdInput", GoGetter: "CollectionIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "engineId", GoGetter: "EngineId"},
			_jsii_.MemberProperty{JsiiProperty: "engineIdInput", GoGetter: "EngineIdInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "homepageSetting", GoGetter: "HomepageSetting"},
			_jsii_.MemberProperty{JsiiProperty: "homepageSettingInput", GoGetter: "HomepageSettingInput"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAccessSettings", GoMethod: "PutAccessSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putHomepageSetting", GoMethod: "PutHomepageSetting"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "putUiBranding", GoMethod: "PutUiBranding"},
			_jsii_.MemberMethod{JsiiMethod: "putUiSettings", GoMethod: "PutUiSettings"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccessSettings", GoMethod: "ResetAccessSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetCollectionId", GoMethod: "ResetCollectionId"},
			_jsii_.MemberMethod{JsiiMethod: "resetHomepageSetting", GoMethod: "ResetHomepageSetting"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetUiBranding", GoMethod: "ResetUiBranding"},
			_jsii_.MemberMethod{JsiiMethod: "resetUiSettings", GoMethod: "ResetUiSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetWidgetConfigId", GoMethod: "ResetWidgetConfigId"},
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
			_jsii_.MemberProperty{JsiiProperty: "uiBranding", GoGetter: "UiBranding"},
			_jsii_.MemberProperty{JsiiProperty: "uiBrandingInput", GoGetter: "UiBrandingInput"},
			_jsii_.MemberProperty{JsiiProperty: "uiSettings", GoGetter: "UiSettings"},
			_jsii_.MemberProperty{JsiiProperty: "uiSettingsInput", GoGetter: "UiSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "widgetConfigId", GoGetter: "WidgetConfigId"},
			_jsii_.MemberProperty{JsiiProperty: "widgetConfigIdInput", GoGetter: "WidgetConfigIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfig{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigAccessSettings",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigAccessSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigAccessSettingsOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigAccessSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowlistedDomains", GoGetter: "AllowlistedDomains"},
			_jsii_.MemberProperty{JsiiProperty: "allowlistedDomainsInput", GoGetter: "AllowlistedDomainsInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowPublicAccess", GoGetter: "AllowPublicAccess"},
			_jsii_.MemberProperty{JsiiProperty: "allowPublicAccessInput", GoGetter: "AllowPublicAccessInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableWebApp", GoGetter: "EnableWebApp"},
			_jsii_.MemberProperty{JsiiProperty: "enableWebAppInput", GoGetter: "EnableWebAppInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "languageCode", GoGetter: "LanguageCode"},
			_jsii_.MemberProperty{JsiiProperty: "languageCodeInput", GoGetter: "LanguageCodeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowlistedDomains", GoMethod: "ResetAllowlistedDomains"},
			_jsii_.MemberMethod{JsiiMethod: "resetAllowPublicAccess", GoMethod: "ResetAllowPublicAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableWebApp", GoMethod: "ResetEnableWebApp"},
			_jsii_.MemberMethod{JsiiMethod: "resetLanguageCode", GoMethod: "ResetLanguageCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkforceIdentityPoolProvider", GoMethod: "ResetWorkforceIdentityPoolProvider"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workforceIdentityPoolProvider", GoGetter: "WorkforceIdentityPoolProvider"},
			_jsii_.MemberProperty{JsiiProperty: "workforceIdentityPoolProviderInput", GoGetter: "WorkforceIdentityPoolProviderInput"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigAccessSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigConfig",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigHomepageSetting",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigHomepageSetting)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigHomepageSettingOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigHomepageSettingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
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
			_jsii_.MemberMethod{JsiiMethod: "putShortcuts", GoMethod: "PutShortcuts"},
			_jsii_.MemberMethod{JsiiMethod: "resetShortcuts", GoMethod: "ResetShortcuts"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "shortcuts", GoGetter: "Shortcuts"},
			_jsii_.MemberProperty{JsiiProperty: "shortcutsInput", GoGetter: "ShortcutsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigHomepageSettingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigHomepageSettingShortcuts",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigHomepageSettingShortcuts)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigHomepageSettingShortcutsIcon",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigHomepageSettingShortcutsIcon)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigHomepageSettingShortcutsIconOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigHomepageSettingShortcutsIconOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetUrl", GoMethod: "ResetUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigHomepageSettingShortcutsIconOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigHomepageSettingShortcutsList",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigHomepageSettingShortcutsList)(nil)).Elem(),
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
			j := jsiiProxy_DiscoveryEngineWidgetConfigHomepageSettingShortcutsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigHomepageSettingShortcutsOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigHomepageSettingShortcutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinationUri", GoGetter: "DestinationUri"},
			_jsii_.MemberProperty{JsiiProperty: "destinationUriInput", GoGetter: "DestinationUriInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "icon", GoGetter: "Icon"},
			_jsii_.MemberProperty{JsiiProperty: "iconInput", GoGetter: "IconInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putIcon", GoMethod: "PutIcon"},
			_jsii_.MemberMethod{JsiiMethod: "resetDestinationUri", GoMethod: "ResetDestinationUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetIcon", GoMethod: "ResetIcon"},
			_jsii_.MemberMethod{JsiiMethod: "resetTitle", GoMethod: "ResetTitle"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberProperty{JsiiProperty: "titleInput", GoGetter: "TitleInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigHomepageSettingShortcutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigTimeouts",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigTimeoutsOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigTimeoutsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "create", GoGetter: "Create"},
			_jsii_.MemberProperty{JsiiProperty: "createInput", GoGetter: "CreateInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "delete", GoGetter: "Delete"},
			_jsii_.MemberProperty{JsiiProperty: "deleteInput", GoGetter: "DeleteInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCreate", GoMethod: "ResetCreate"},
			_jsii_.MemberMethod{JsiiMethod: "resetDelete", GoMethod: "ResetDelete"},
			_jsii_.MemberMethod{JsiiMethod: "resetUpdate", GoMethod: "ResetUpdate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "update", GoGetter: "Update"},
			_jsii_.MemberProperty{JsiiProperty: "updateInput", GoGetter: "UpdateInput"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiBranding",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiBranding)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiBrandingLogo",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiBrandingLogo)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiBrandingLogoOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiBrandingLogoOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetUrl", GoMethod: "ResetUrl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "url", GoGetter: "Url"},
			_jsii_.MemberProperty{JsiiProperty: "urlInput", GoGetter: "UrlInput"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigUiBrandingLogoOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiBrandingOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiBrandingOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
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
			_jsii_.MemberProperty{JsiiProperty: "logo", GoGetter: "Logo"},
			_jsii_.MemberProperty{JsiiProperty: "logoInput", GoGetter: "LogoInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLogo", GoMethod: "PutLogo"},
			_jsii_.MemberMethod{JsiiMethod: "resetLogo", GoMethod: "ResetLogo"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigUiBrandingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettings",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigs",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFacetField",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFacetField)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFacetFieldList",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFacetFieldList)(nil)).Elem(),
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
			j := jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFacetFieldList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFacetFieldOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFacetFieldOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "field", GoGetter: "Field"},
			_jsii_.MemberProperty{JsiiProperty: "fieldInput", GoGetter: "FieldInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFacetFieldOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFieldsUiComponentsMap",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFieldsUiComponentsMap)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFieldsUiComponentsMapList",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFieldsUiComponentsMapList)(nil)).Elem(),
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
			j := jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFieldsUiComponentsMapList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFieldsUiComponentsMapOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFieldsUiComponentsMapOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "deviceVisibility", GoGetter: "DeviceVisibility"},
			_jsii_.MemberProperty{JsiiProperty: "deviceVisibilityInput", GoGetter: "DeviceVisibilityInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayTemplate", GoGetter: "DisplayTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "displayTemplateInput", GoGetter: "DisplayTemplateInput"},
			_jsii_.MemberProperty{JsiiProperty: "field", GoGetter: "Field"},
			_jsii_.MemberProperty{JsiiProperty: "fieldInput", GoGetter: "FieldInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDeviceVisibility", GoMethod: "ResetDeviceVisibility"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayTemplate", GoMethod: "ResetDisplayTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "uiComponent", GoGetter: "UiComponent"},
			_jsii_.MemberProperty{JsiiProperty: "uiComponentInput", GoGetter: "UiComponentInput"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsFieldsUiComponentsMapOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsList",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsList)(nil)).Elem(),
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
			j := jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "facetField", GoGetter: "FacetField"},
			_jsii_.MemberProperty{JsiiProperty: "facetFieldInput", GoGetter: "FacetFieldInput"},
			_jsii_.MemberProperty{JsiiProperty: "fieldsUiComponentsMap", GoGetter: "FieldsUiComponentsMap"},
			_jsii_.MemberProperty{JsiiProperty: "fieldsUiComponentsMapInput", GoGetter: "FieldsUiComponentsMapInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putFacetField", GoMethod: "PutFacetField"},
			_jsii_.MemberMethod{JsiiMethod: "putFieldsUiComponentsMap", GoMethod: "PutFieldsUiComponentsMap"},
			_jsii_.MemberMethod{JsiiMethod: "resetFacetField", GoMethod: "ResetFacetField"},
			_jsii_.MemberMethod{JsiiMethod: "resetFieldsUiComponentsMap", GoMethod: "ResetFieldsUiComponentsMap"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsDataStoreUiConfigsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "disableRelatedQuestions", GoGetter: "DisableRelatedQuestions"},
			_jsii_.MemberProperty{JsiiProperty: "disableRelatedQuestionsInput", GoGetter: "DisableRelatedQuestionsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ignoreAdversarialQuery", GoGetter: "IgnoreAdversarialQuery"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreAdversarialQueryInput", GoGetter: "IgnoreAdversarialQueryInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreLowRelevantContent", GoGetter: "IgnoreLowRelevantContent"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreLowRelevantContentInput", GoGetter: "IgnoreLowRelevantContentInput"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreNonAnswerSeekingQuery", GoGetter: "IgnoreNonAnswerSeekingQuery"},
			_jsii_.MemberProperty{JsiiProperty: "ignoreNonAnswerSeekingQueryInput", GoGetter: "IgnoreNonAnswerSeekingQueryInput"},
			_jsii_.MemberProperty{JsiiProperty: "imageSource", GoGetter: "ImageSource"},
			_jsii_.MemberProperty{JsiiProperty: "imageSourceInput", GoGetter: "ImageSourceInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "languageCode", GoGetter: "LanguageCode"},
			_jsii_.MemberProperty{JsiiProperty: "languageCodeInput", GoGetter: "LanguageCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxRephraseSteps", GoGetter: "MaxRephraseSteps"},
			_jsii_.MemberProperty{JsiiProperty: "maxRephraseStepsInput", GoGetter: "MaxRephraseStepsInput"},
			_jsii_.MemberProperty{JsiiProperty: "modelPromptPreamble", GoGetter: "ModelPromptPreamble"},
			_jsii_.MemberProperty{JsiiProperty: "modelPromptPreambleInput", GoGetter: "ModelPromptPreambleInput"},
			_jsii_.MemberProperty{JsiiProperty: "modelVersion", GoGetter: "ModelVersion"},
			_jsii_.MemberProperty{JsiiProperty: "modelVersionInput", GoGetter: "ModelVersionInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableRelatedQuestions", GoMethod: "ResetDisableRelatedQuestions"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreAdversarialQuery", GoMethod: "ResetIgnoreAdversarialQuery"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreLowRelevantContent", GoMethod: "ResetIgnoreLowRelevantContent"},
			_jsii_.MemberMethod{JsiiMethod: "resetIgnoreNonAnswerSeekingQuery", GoMethod: "ResetIgnoreNonAnswerSeekingQuery"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageSource", GoMethod: "ResetImageSource"},
			_jsii_.MemberMethod{JsiiMethod: "resetLanguageCode", GoMethod: "ResetLanguageCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxRephraseSteps", GoMethod: "ResetMaxRephraseSteps"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelPromptPreamble", GoMethod: "ResetModelPromptPreamble"},
			_jsii_.MemberMethod{JsiiMethod: "resetModelVersion", GoMethod: "ResetModelVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetResultCount", GoMethod: "ResetResultCount"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resultCount", GoGetter: "ResultCount"},
			_jsii_.MemberProperty{JsiiProperty: "resultCountInput", GoGetter: "ResultCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsGenerativeAnswerConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google.discoveryEngineWidgetConfig.DiscoveryEngineWidgetConfigUiSettingsOutputReference",
		reflect.TypeOf((*DiscoveryEngineWidgetConfigUiSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataStoreUiConfigs", GoGetter: "DataStoreUiConfigs"},
			_jsii_.MemberProperty{JsiiProperty: "dataStoreUiConfigsInput", GoGetter: "DataStoreUiConfigsInput"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSearchRequestOrderBy", GoGetter: "DefaultSearchRequestOrderBy"},
			_jsii_.MemberProperty{JsiiProperty: "defaultSearchRequestOrderByInput", GoGetter: "DefaultSearchRequestOrderByInput"},
			_jsii_.MemberProperty{JsiiProperty: "disableUserEventsCollection", GoGetter: "DisableUserEventsCollection"},
			_jsii_.MemberProperty{JsiiProperty: "disableUserEventsCollectionInput", GoGetter: "DisableUserEventsCollectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableAutocomplete", GoGetter: "EnableAutocomplete"},
			_jsii_.MemberProperty{JsiiProperty: "enableAutocompleteInput", GoGetter: "EnableAutocompleteInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableCreateAgentButton", GoGetter: "EnableCreateAgentButton"},
			_jsii_.MemberProperty{JsiiProperty: "enableCreateAgentButtonInput", GoGetter: "EnableCreateAgentButtonInput"},
			_jsii_.MemberProperty{JsiiProperty: "enablePeopleSearch", GoGetter: "EnablePeopleSearch"},
			_jsii_.MemberProperty{JsiiProperty: "enablePeopleSearchInput", GoGetter: "EnablePeopleSearchInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableQualityFeedback", GoGetter: "EnableQualityFeedback"},
			_jsii_.MemberProperty{JsiiProperty: "enableQualityFeedbackInput", GoGetter: "EnableQualityFeedbackInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableSafeSearch", GoGetter: "EnableSafeSearch"},
			_jsii_.MemberProperty{JsiiProperty: "enableSafeSearchInput", GoGetter: "EnableSafeSearchInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableSearchAsYouType", GoGetter: "EnableSearchAsYouType"},
			_jsii_.MemberProperty{JsiiProperty: "enableSearchAsYouTypeInput", GoGetter: "EnableSearchAsYouTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableVisualContentSummary", GoGetter: "EnableVisualContentSummary"},
			_jsii_.MemberProperty{JsiiProperty: "enableVisualContentSummaryInput", GoGetter: "EnableVisualContentSummaryInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "generativeAnswerConfig", GoGetter: "GenerativeAnswerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "generativeAnswerConfigInput", GoGetter: "GenerativeAnswerConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "interactionType", GoGetter: "InteractionType"},
			_jsii_.MemberProperty{JsiiProperty: "interactionTypeInput", GoGetter: "InteractionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putDataStoreUiConfigs", GoMethod: "PutDataStoreUiConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "putGenerativeAnswerConfig", GoMethod: "PutGenerativeAnswerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataStoreUiConfigs", GoMethod: "ResetDataStoreUiConfigs"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultSearchRequestOrderBy", GoMethod: "ResetDefaultSearchRequestOrderBy"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisableUserEventsCollection", GoMethod: "ResetDisableUserEventsCollection"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableAutocomplete", GoMethod: "ResetEnableAutocomplete"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableCreateAgentButton", GoMethod: "ResetEnableCreateAgentButton"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnablePeopleSearch", GoMethod: "ResetEnablePeopleSearch"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableQualityFeedback", GoMethod: "ResetEnableQualityFeedback"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableSafeSearch", GoMethod: "ResetEnableSafeSearch"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableSearchAsYouType", GoMethod: "ResetEnableSearchAsYouType"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableVisualContentSummary", GoMethod: "ResetEnableVisualContentSummary"},
			_jsii_.MemberMethod{JsiiMethod: "resetGenerativeAnswerConfig", GoMethod: "ResetGenerativeAnswerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetInteractionType", GoMethod: "ResetInteractionType"},
			_jsii_.MemberMethod{JsiiMethod: "resetResultDescriptionType", GoMethod: "ResetResultDescriptionType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "resultDescriptionType", GoGetter: "ResultDescriptionType"},
			_jsii_.MemberProperty{JsiiProperty: "resultDescriptionTypeInput", GoGetter: "ResultDescriptionTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_DiscoveryEngineWidgetConfigUiSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
