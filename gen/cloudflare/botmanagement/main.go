package botmanagement

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-cloudflare.botManagement.BotManagement",
		reflect.TypeOf((*BotManagement)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "autoUpdateModel", GoGetter: "AutoUpdateModel"},
			_jsii_.MemberProperty{JsiiProperty: "autoUpdateModelInput", GoGetter: "AutoUpdateModelInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "enableJs", GoGetter: "EnableJs"},
			_jsii_.MemberProperty{JsiiProperty: "enableJsInput", GoGetter: "EnableJsInput"},
			_jsii_.MemberProperty{JsiiProperty: "fightMode", GoGetter: "FightMode"},
			_jsii_.MemberProperty{JsiiProperty: "fightModeInput", GoGetter: "FightModeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "optimizeWordpress", GoGetter: "OptimizeWordpress"},
			_jsii_.MemberProperty{JsiiProperty: "optimizeWordpressInput", GoGetter: "OptimizeWordpressInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoUpdateModel", GoMethod: "ResetAutoUpdateModel"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableJs", GoMethod: "ResetEnableJs"},
			_jsii_.MemberMethod{JsiiMethod: "resetFightMode", GoMethod: "ResetFightMode"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptimizeWordpress", GoMethod: "ResetOptimizeWordpress"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetSbfmDefinitelyAutomated", GoMethod: "ResetSbfmDefinitelyAutomated"},
			_jsii_.MemberMethod{JsiiMethod: "resetSbfmLikelyAutomated", GoMethod: "ResetSbfmLikelyAutomated"},
			_jsii_.MemberMethod{JsiiMethod: "resetSbfmStaticResourceProtection", GoMethod: "ResetSbfmStaticResourceProtection"},
			_jsii_.MemberMethod{JsiiMethod: "resetSbfmVerifiedBots", GoMethod: "ResetSbfmVerifiedBots"},
			_jsii_.MemberMethod{JsiiMethod: "resetSuppressSessionScore", GoMethod: "ResetSuppressSessionScore"},
			_jsii_.MemberProperty{JsiiProperty: "sbfmDefinitelyAutomated", GoGetter: "SbfmDefinitelyAutomated"},
			_jsii_.MemberProperty{JsiiProperty: "sbfmDefinitelyAutomatedInput", GoGetter: "SbfmDefinitelyAutomatedInput"},
			_jsii_.MemberProperty{JsiiProperty: "sbfmLikelyAutomated", GoGetter: "SbfmLikelyAutomated"},
			_jsii_.MemberProperty{JsiiProperty: "sbfmLikelyAutomatedInput", GoGetter: "SbfmLikelyAutomatedInput"},
			_jsii_.MemberProperty{JsiiProperty: "sbfmStaticResourceProtection", GoGetter: "SbfmStaticResourceProtection"},
			_jsii_.MemberProperty{JsiiProperty: "sbfmStaticResourceProtectionInput", GoGetter: "SbfmStaticResourceProtectionInput"},
			_jsii_.MemberProperty{JsiiProperty: "sbfmVerifiedBots", GoGetter: "SbfmVerifiedBots"},
			_jsii_.MemberProperty{JsiiProperty: "sbfmVerifiedBotsInput", GoGetter: "SbfmVerifiedBotsInput"},
			_jsii_.MemberProperty{JsiiProperty: "suppressSessionScore", GoGetter: "SuppressSessionScore"},
			_jsii_.MemberProperty{JsiiProperty: "suppressSessionScoreInput", GoGetter: "SuppressSessionScoreInput"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "usingLatestModel", GoGetter: "UsingLatestModel"},
			_jsii_.MemberProperty{JsiiProperty: "zoneId", GoGetter: "ZoneId"},
			_jsii_.MemberProperty{JsiiProperty: "zoneIdInput", GoGetter: "ZoneIdInput"},
		},
		func() interface{} {
			j := jsiiProxy_BotManagement{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-cloudflare.botManagement.BotManagementConfig",
		reflect.TypeOf((*BotManagementConfig)(nil)).Elem(),
	)
}
