package googledialogflowcxpage

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPage",
		reflect.TypeOf((*GoogleDialogflowCxPage)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSettings", GoGetter: "AdvancedSettings"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSettingsInput", GoGetter: "AdvancedSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "entryFulfillment", GoGetter: "EntryFulfillment"},
			_jsii_.MemberProperty{JsiiProperty: "entryFulfillmentInput", GoGetter: "EntryFulfillmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "eventHandlers", GoGetter: "EventHandlers"},
			_jsii_.MemberProperty{JsiiProperty: "eventHandlersInput", GoGetter: "EventHandlersInput"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "form", GoGetter: "Form"},
			_jsii_.MemberProperty{JsiiProperty: "formInput", GoGetter: "FormInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "languageCode", GoGetter: "LanguageCode"},
			_jsii_.MemberProperty{JsiiProperty: "languageCodeInput", GoGetter: "LanguageCodeInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parent", GoGetter: "Parent"},
			_jsii_.MemberProperty{JsiiProperty: "parentInput", GoGetter: "ParentInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putAdvancedSettings", GoMethod: "PutAdvancedSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putEntryFulfillment", GoMethod: "PutEntryFulfillment"},
			_jsii_.MemberMethod{JsiiMethod: "putEventHandlers", GoMethod: "PutEventHandlers"},
			_jsii_.MemberMethod{JsiiMethod: "putForm", GoMethod: "PutForm"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "putTransitionRoutes", GoMethod: "PutTransitionRoutes"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdvancedSettings", GoMethod: "ResetAdvancedSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetEntryFulfillment", GoMethod: "ResetEntryFulfillment"},
			_jsii_.MemberMethod{JsiiMethod: "resetEventHandlers", GoMethod: "ResetEventHandlers"},
			_jsii_.MemberMethod{JsiiMethod: "resetForm", GoMethod: "ResetForm"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLanguageCode", GoMethod: "ResetLanguageCode"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetParent", GoMethod: "ResetParent"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransitionRouteGroups", GoMethod: "ResetTransitionRouteGroups"},
			_jsii_.MemberMethod{JsiiMethod: "resetTransitionRoutes", GoMethod: "ResetTransitionRoutes"},
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
			_jsii_.MemberProperty{JsiiProperty: "transitionRouteGroups", GoGetter: "TransitionRouteGroups"},
			_jsii_.MemberProperty{JsiiProperty: "transitionRouteGroupsInput", GoGetter: "TransitionRouteGroupsInput"},
			_jsii_.MemberProperty{JsiiProperty: "transitionRoutes", GoGetter: "TransitionRoutes"},
			_jsii_.MemberProperty{JsiiProperty: "transitionRoutesInput", GoGetter: "TransitionRoutesInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPage{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageAdvancedSettings",
		reflect.TypeOf((*GoogleDialogflowCxPageAdvancedSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageAdvancedSettingsDtmfSettings",
		reflect.TypeOf((*GoogleDialogflowCxPageAdvancedSettingsDtmfSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageAdvancedSettingsDtmfSettingsOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageAdvancedSettingsDtmfSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "finishDigit", GoGetter: "FinishDigit"},
			_jsii_.MemberProperty{JsiiProperty: "finishDigitInput", GoGetter: "FinishDigitInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maxDigits", GoGetter: "MaxDigits"},
			_jsii_.MemberProperty{JsiiProperty: "maxDigitsInput", GoGetter: "MaxDigitsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFinishDigit", GoMethod: "ResetFinishDigit"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxDigits", GoMethod: "ResetMaxDigits"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageAdvancedSettingsDtmfSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageAdvancedSettingsOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageAdvancedSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dtmfSettings", GoGetter: "DtmfSettings"},
			_jsii_.MemberProperty{JsiiProperty: "dtmfSettingsInput", GoGetter: "DtmfSettingsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putDtmfSettings", GoMethod: "PutDtmfSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDtmfSettings", GoMethod: "ResetDtmfSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageAdvancedSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageConfig",
		reflect.TypeOf((*GoogleDialogflowCxPageConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillment",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentConditionalCases",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentConditionalCases)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentConditionalCasesList",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentConditionalCasesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentConditionalCasesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentConditionalCasesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentConditionalCasesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cases", GoGetter: "Cases"},
			_jsii_.MemberProperty{JsiiProperty: "casesInput", GoGetter: "CasesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCases", GoMethod: "ResetCases"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentConditionalCasesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessages",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessages)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesConversationSuccess",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesConversationSuccess)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesConversationSuccessOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesConversationSuccessOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentMessagesConversationSuccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesList",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentMessagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesLiveAgentHandoff",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesLiveAgentHandoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesLiveAgentHandoffOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesLiveAgentHandoffOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentMessagesLiveAgentHandoffOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesOutputAudioText",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesOutputAudioText)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesOutputAudioTextOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesOutputAudioTextOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetSsml", GoMethod: "ResetSsml"},
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ssml", GoGetter: "Ssml"},
			_jsii_.MemberProperty{JsiiProperty: "ssmlInput", GoGetter: "SsmlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentMessagesOutputAudioTextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "channel", GoGetter: "Channel"},
			_jsii_.MemberProperty{JsiiProperty: "channelInput", GoGetter: "ChannelInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conversationSuccess", GoGetter: "ConversationSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "conversationSuccessInput", GoGetter: "ConversationSuccessInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "liveAgentHandoff", GoGetter: "LiveAgentHandoff"},
			_jsii_.MemberProperty{JsiiProperty: "liveAgentHandoffInput", GoGetter: "LiveAgentHandoffInput"},
			_jsii_.MemberProperty{JsiiProperty: "outputAudioText", GoGetter: "OutputAudioText"},
			_jsii_.MemberProperty{JsiiProperty: "outputAudioTextInput", GoGetter: "OutputAudioTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "payload", GoGetter: "Payload"},
			_jsii_.MemberProperty{JsiiProperty: "payloadInput", GoGetter: "PayloadInput"},
			_jsii_.MemberProperty{JsiiProperty: "playAudio", GoGetter: "PlayAudio"},
			_jsii_.MemberProperty{JsiiProperty: "playAudioInput", GoGetter: "PlayAudioInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConversationSuccess", GoMethod: "PutConversationSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "putLiveAgentHandoff", GoMethod: "PutLiveAgentHandoff"},
			_jsii_.MemberMethod{JsiiMethod: "putOutputAudioText", GoMethod: "PutOutputAudioText"},
			_jsii_.MemberMethod{JsiiMethod: "putPlayAudio", GoMethod: "PutPlayAudio"},
			_jsii_.MemberMethod{JsiiMethod: "putTelephonyTransferCall", GoMethod: "PutTelephonyTransferCall"},
			_jsii_.MemberMethod{JsiiMethod: "putText", GoMethod: "PutText"},
			_jsii_.MemberMethod{JsiiMethod: "resetChannel", GoMethod: "ResetChannel"},
			_jsii_.MemberMethod{JsiiMethod: "resetConversationSuccess", GoMethod: "ResetConversationSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetLiveAgentHandoff", GoMethod: "ResetLiveAgentHandoff"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputAudioText", GoMethod: "ResetOutputAudioText"},
			_jsii_.MemberMethod{JsiiMethod: "resetPayload", GoMethod: "ResetPayload"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlayAudio", GoMethod: "ResetPlayAudio"},
			_jsii_.MemberMethod{JsiiMethod: "resetTelephonyTransferCall", GoMethod: "ResetTelephonyTransferCall"},
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "telephonyTransferCall", GoGetter: "TelephonyTransferCall"},
			_jsii_.MemberProperty{JsiiProperty: "telephonyTransferCallInput", GoGetter: "TelephonyTransferCallInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentMessagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesPlayAudio",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesPlayAudio)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesPlayAudioOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesPlayAudioOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
			_jsii_.MemberProperty{JsiiProperty: "audioUri", GoGetter: "AudioUri"},
			_jsii_.MemberProperty{JsiiProperty: "audioUriInput", GoGetter: "AudioUriInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentMessagesPlayAudioOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesTelephonyTransferCall",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesTelephonyTransferCall)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesTelephonyTransferCallOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesTelephonyTransferCallOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "phoneNumber", GoGetter: "PhoneNumber"},
			_jsii_.MemberProperty{JsiiProperty: "phoneNumberInput", GoGetter: "PhoneNumberInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentMessagesTelephonyTransferCallOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesText",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesText)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentMessagesTextOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentMessagesTextOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentMessagesTextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalCases", GoGetter: "ConditionalCases"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalCasesInput", GoGetter: "ConditionalCasesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "messages", GoGetter: "Messages"},
			_jsii_.MemberProperty{JsiiProperty: "messagesInput", GoGetter: "MessagesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConditionalCases", GoMethod: "PutConditionalCases"},
			_jsii_.MemberMethod{JsiiMethod: "putMessages", GoMethod: "PutMessages"},
			_jsii_.MemberMethod{JsiiMethod: "putSetParameterActions", GoMethod: "PutSetParameterActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetConditionalCases", GoMethod: "ResetConditionalCases"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessages", GoMethod: "ResetMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resetReturnPartialResponses", GoMethod: "ResetReturnPartialResponses"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetParameterActions", GoMethod: "ResetSetParameterActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTag", GoMethod: "ResetTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebhook", GoMethod: "ResetWebhook"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "returnPartialResponses", GoGetter: "ReturnPartialResponses"},
			_jsii_.MemberProperty{JsiiProperty: "returnPartialResponsesInput", GoGetter: "ReturnPartialResponsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "setParameterActions", GoGetter: "SetParameterActions"},
			_jsii_.MemberProperty{JsiiProperty: "setParameterActionsInput", GoGetter: "SetParameterActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tag", GoGetter: "Tag"},
			_jsii_.MemberProperty{JsiiProperty: "tagInput", GoGetter: "TagInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webhook", GoGetter: "Webhook"},
			_jsii_.MemberProperty{JsiiProperty: "webhookInput", GoGetter: "WebhookInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentSetParameterActions",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentSetParameterActions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentSetParameterActionsList",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentSetParameterActionsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentSetParameterActionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEntryFulfillmentSetParameterActionsOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEntryFulfillmentSetParameterActionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameter", GoGetter: "Parameter"},
			_jsii_.MemberProperty{JsiiProperty: "parameterInput", GoGetter: "ParameterInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameter", GoMethod: "ResetParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEntryFulfillmentSetParameterActionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlers",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlers)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersList",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "event", GoGetter: "Event"},
			_jsii_.MemberProperty{JsiiProperty: "eventInput", GoGetter: "EventInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putTriggerFulfillment", GoMethod: "PutTriggerFulfillment"},
			_jsii_.MemberMethod{JsiiMethod: "resetEvent", GoMethod: "ResetEvent"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetFlow", GoMethod: "ResetTargetFlow"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetPage", GoMethod: "ResetTargetPage"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggerFulfillment", GoMethod: "ResetTriggerFulfillment"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "targetFlow", GoGetter: "TargetFlow"},
			_jsii_.MemberProperty{JsiiProperty: "targetFlowInput", GoGetter: "TargetFlowInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetPage", GoGetter: "TargetPage"},
			_jsii_.MemberProperty{JsiiProperty: "targetPageInput", GoGetter: "TargetPageInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggerFulfillment", GoGetter: "TriggerFulfillment"},
			_jsii_.MemberProperty{JsiiProperty: "triggerFulfillmentInput", GoGetter: "TriggerFulfillmentInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillment",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentConditionalCases",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentConditionalCases)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentConditionalCasesList",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentConditionalCasesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentConditionalCasesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentConditionalCasesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentConditionalCasesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cases", GoGetter: "Cases"},
			_jsii_.MemberProperty{JsiiProperty: "casesInput", GoGetter: "CasesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCases", GoMethod: "ResetCases"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentConditionalCasesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessages",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessages)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccess",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccess)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccessOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccessOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesConversationSuccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesList",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoff",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoffOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoffOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesLiveAgentHandoffOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioText)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetSsml", GoMethod: "ResetSsml"},
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ssml", GoGetter: "Ssml"},
			_jsii_.MemberProperty{JsiiProperty: "ssmlInput", GoGetter: "SsmlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "channel", GoGetter: "Channel"},
			_jsii_.MemberProperty{JsiiProperty: "channelInput", GoGetter: "ChannelInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conversationSuccess", GoGetter: "ConversationSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "conversationSuccessInput", GoGetter: "ConversationSuccessInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "liveAgentHandoff", GoGetter: "LiveAgentHandoff"},
			_jsii_.MemberProperty{JsiiProperty: "liveAgentHandoffInput", GoGetter: "LiveAgentHandoffInput"},
			_jsii_.MemberProperty{JsiiProperty: "outputAudioText", GoGetter: "OutputAudioText"},
			_jsii_.MemberProperty{JsiiProperty: "outputAudioTextInput", GoGetter: "OutputAudioTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "payload", GoGetter: "Payload"},
			_jsii_.MemberProperty{JsiiProperty: "payloadInput", GoGetter: "PayloadInput"},
			_jsii_.MemberProperty{JsiiProperty: "playAudio", GoGetter: "PlayAudio"},
			_jsii_.MemberProperty{JsiiProperty: "playAudioInput", GoGetter: "PlayAudioInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConversationSuccess", GoMethod: "PutConversationSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "putLiveAgentHandoff", GoMethod: "PutLiveAgentHandoff"},
			_jsii_.MemberMethod{JsiiMethod: "putOutputAudioText", GoMethod: "PutOutputAudioText"},
			_jsii_.MemberMethod{JsiiMethod: "putPlayAudio", GoMethod: "PutPlayAudio"},
			_jsii_.MemberMethod{JsiiMethod: "putTelephonyTransferCall", GoMethod: "PutTelephonyTransferCall"},
			_jsii_.MemberMethod{JsiiMethod: "putText", GoMethod: "PutText"},
			_jsii_.MemberMethod{JsiiMethod: "resetChannel", GoMethod: "ResetChannel"},
			_jsii_.MemberMethod{JsiiMethod: "resetConversationSuccess", GoMethod: "ResetConversationSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetLiveAgentHandoff", GoMethod: "ResetLiveAgentHandoff"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputAudioText", GoMethod: "ResetOutputAudioText"},
			_jsii_.MemberMethod{JsiiMethod: "resetPayload", GoMethod: "ResetPayload"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlayAudio", GoMethod: "ResetPlayAudio"},
			_jsii_.MemberMethod{JsiiMethod: "resetTelephonyTransferCall", GoMethod: "ResetTelephonyTransferCall"},
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "telephonyTransferCall", GoGetter: "TelephonyTransferCall"},
			_jsii_.MemberProperty{JsiiProperty: "telephonyTransferCallInput", GoGetter: "TelephonyTransferCallInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudio",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudio)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudioOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudioOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
			_jsii_.MemberProperty{JsiiProperty: "audioUri", GoGetter: "AudioUri"},
			_jsii_.MemberProperty{JsiiProperty: "audioUriInput", GoGetter: "AudioUriInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesPlayAudioOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCallOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCallOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "phoneNumber", GoGetter: "PhoneNumber"},
			_jsii_.MemberProperty{JsiiProperty: "phoneNumberInput", GoGetter: "PhoneNumberInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesTelephonyTransferCallOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesText",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesText)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesTextOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesTextOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentMessagesTextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalCases", GoGetter: "ConditionalCases"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalCasesInput", GoGetter: "ConditionalCasesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "messages", GoGetter: "Messages"},
			_jsii_.MemberProperty{JsiiProperty: "messagesInput", GoGetter: "MessagesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConditionalCases", GoMethod: "PutConditionalCases"},
			_jsii_.MemberMethod{JsiiMethod: "putMessages", GoMethod: "PutMessages"},
			_jsii_.MemberMethod{JsiiMethod: "putSetParameterActions", GoMethod: "PutSetParameterActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetConditionalCases", GoMethod: "ResetConditionalCases"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessages", GoMethod: "ResetMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resetReturnPartialResponses", GoMethod: "ResetReturnPartialResponses"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetParameterActions", GoMethod: "ResetSetParameterActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTag", GoMethod: "ResetTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebhook", GoMethod: "ResetWebhook"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "returnPartialResponses", GoGetter: "ReturnPartialResponses"},
			_jsii_.MemberProperty{JsiiProperty: "returnPartialResponsesInput", GoGetter: "ReturnPartialResponsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "setParameterActions", GoGetter: "SetParameterActions"},
			_jsii_.MemberProperty{JsiiProperty: "setParameterActionsInput", GoGetter: "SetParameterActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tag", GoGetter: "Tag"},
			_jsii_.MemberProperty{JsiiProperty: "tagInput", GoGetter: "TagInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webhook", GoGetter: "Webhook"},
			_jsii_.MemberProperty{JsiiProperty: "webhookInput", GoGetter: "WebhookInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentSetParameterActions",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentSetParameterActions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentSetParameterActionsList",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentSetParameterActionsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentSetParameterActionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageEventHandlersTriggerFulfillmentSetParameterActionsOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageEventHandlersTriggerFulfillmentSetParameterActionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameter", GoGetter: "Parameter"},
			_jsii_.MemberProperty{JsiiProperty: "parameterInput", GoGetter: "ParameterInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameter", GoMethod: "ResetParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageEventHandlersTriggerFulfillmentSetParameterActionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageForm",
		reflect.TypeOf((*GoogleDialogflowCxPageForm)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberMethod{JsiiMethod: "putParameters", GoMethod: "PutParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParameters",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParameters)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersAdvancedSettings",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersAdvancedSettings)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersAdvancedSettingsDtmfSettings",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersAdvancedSettingsDtmfSettings)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersAdvancedSettingsDtmfSettingsOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersAdvancedSettingsDtmfSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enabled", GoGetter: "Enabled"},
			_jsii_.MemberProperty{JsiiProperty: "enabledInput", GoGetter: "EnabledInput"},
			_jsii_.MemberProperty{JsiiProperty: "finishDigit", GoGetter: "FinishDigit"},
			_jsii_.MemberProperty{JsiiProperty: "finishDigitInput", GoGetter: "FinishDigitInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "maxDigits", GoGetter: "MaxDigits"},
			_jsii_.MemberProperty{JsiiProperty: "maxDigitsInput", GoGetter: "MaxDigitsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnabled", GoMethod: "ResetEnabled"},
			_jsii_.MemberMethod{JsiiMethod: "resetFinishDigit", GoMethod: "ResetFinishDigit"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxDigits", GoMethod: "ResetMaxDigits"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersAdvancedSettingsDtmfSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersAdvancedSettingsOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersAdvancedSettingsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dtmfSettings", GoGetter: "DtmfSettings"},
			_jsii_.MemberProperty{JsiiProperty: "dtmfSettingsInput", GoGetter: "DtmfSettingsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putDtmfSettings", GoMethod: "PutDtmfSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDtmfSettings", GoMethod: "ResetDtmfSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersAdvancedSettingsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehavior",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehavior)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillment",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCases",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCases)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCasesList",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCasesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCasesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCasesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCasesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cases", GoGetter: "Cases"},
			_jsii_.MemberProperty{JsiiProperty: "casesInput", GoGetter: "CasesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCases", GoMethod: "ResetCases"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentConditionalCasesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessages",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessages)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesConversationSuccess",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesConversationSuccess)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesConversationSuccessOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesConversationSuccessOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesConversationSuccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesList",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesLiveAgentHandoff",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesLiveAgentHandoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesLiveAgentHandoffOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesLiveAgentHandoffOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesLiveAgentHandoffOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesOutputAudioText",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesOutputAudioText)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesOutputAudioTextOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesOutputAudioTextOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetSsml", GoMethod: "ResetSsml"},
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ssml", GoGetter: "Ssml"},
			_jsii_.MemberProperty{JsiiProperty: "ssmlInput", GoGetter: "SsmlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesOutputAudioTextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "channel", GoGetter: "Channel"},
			_jsii_.MemberProperty{JsiiProperty: "channelInput", GoGetter: "ChannelInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conversationSuccess", GoGetter: "ConversationSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "conversationSuccessInput", GoGetter: "ConversationSuccessInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "liveAgentHandoff", GoGetter: "LiveAgentHandoff"},
			_jsii_.MemberProperty{JsiiProperty: "liveAgentHandoffInput", GoGetter: "LiveAgentHandoffInput"},
			_jsii_.MemberProperty{JsiiProperty: "outputAudioText", GoGetter: "OutputAudioText"},
			_jsii_.MemberProperty{JsiiProperty: "outputAudioTextInput", GoGetter: "OutputAudioTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "payload", GoGetter: "Payload"},
			_jsii_.MemberProperty{JsiiProperty: "payloadInput", GoGetter: "PayloadInput"},
			_jsii_.MemberProperty{JsiiProperty: "playAudio", GoGetter: "PlayAudio"},
			_jsii_.MemberProperty{JsiiProperty: "playAudioInput", GoGetter: "PlayAudioInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConversationSuccess", GoMethod: "PutConversationSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "putLiveAgentHandoff", GoMethod: "PutLiveAgentHandoff"},
			_jsii_.MemberMethod{JsiiMethod: "putOutputAudioText", GoMethod: "PutOutputAudioText"},
			_jsii_.MemberMethod{JsiiMethod: "putPlayAudio", GoMethod: "PutPlayAudio"},
			_jsii_.MemberMethod{JsiiMethod: "putTelephonyTransferCall", GoMethod: "PutTelephonyTransferCall"},
			_jsii_.MemberMethod{JsiiMethod: "putText", GoMethod: "PutText"},
			_jsii_.MemberMethod{JsiiMethod: "resetChannel", GoMethod: "ResetChannel"},
			_jsii_.MemberMethod{JsiiMethod: "resetConversationSuccess", GoMethod: "ResetConversationSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetLiveAgentHandoff", GoMethod: "ResetLiveAgentHandoff"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputAudioText", GoMethod: "ResetOutputAudioText"},
			_jsii_.MemberMethod{JsiiMethod: "resetPayload", GoMethod: "ResetPayload"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlayAudio", GoMethod: "ResetPlayAudio"},
			_jsii_.MemberMethod{JsiiMethod: "resetTelephonyTransferCall", GoMethod: "ResetTelephonyTransferCall"},
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "telephonyTransferCall", GoGetter: "TelephonyTransferCall"},
			_jsii_.MemberProperty{JsiiProperty: "telephonyTransferCallInput", GoGetter: "TelephonyTransferCallInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesPlayAudio",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesPlayAudio)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesPlayAudioOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesPlayAudioOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
			_jsii_.MemberProperty{JsiiProperty: "audioUri", GoGetter: "AudioUri"},
			_jsii_.MemberProperty{JsiiProperty: "audioUriInput", GoGetter: "AudioUriInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesPlayAudioOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTelephonyTransferCall",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTelephonyTransferCall)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTelephonyTransferCallOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTelephonyTransferCallOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "phoneNumber", GoGetter: "PhoneNumber"},
			_jsii_.MemberProperty{JsiiProperty: "phoneNumberInput", GoGetter: "PhoneNumberInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTelephonyTransferCallOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesText",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesText)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTextOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTextOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentMessagesTextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalCases", GoGetter: "ConditionalCases"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalCasesInput", GoGetter: "ConditionalCasesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "messages", GoGetter: "Messages"},
			_jsii_.MemberProperty{JsiiProperty: "messagesInput", GoGetter: "MessagesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConditionalCases", GoMethod: "PutConditionalCases"},
			_jsii_.MemberMethod{JsiiMethod: "putMessages", GoMethod: "PutMessages"},
			_jsii_.MemberMethod{JsiiMethod: "putSetParameterActions", GoMethod: "PutSetParameterActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetConditionalCases", GoMethod: "ResetConditionalCases"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessages", GoMethod: "ResetMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resetReturnPartialResponses", GoMethod: "ResetReturnPartialResponses"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetParameterActions", GoMethod: "ResetSetParameterActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTag", GoMethod: "ResetTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebhook", GoMethod: "ResetWebhook"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "returnPartialResponses", GoGetter: "ReturnPartialResponses"},
			_jsii_.MemberProperty{JsiiProperty: "returnPartialResponsesInput", GoGetter: "ReturnPartialResponsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "setParameterActions", GoGetter: "SetParameterActions"},
			_jsii_.MemberProperty{JsiiProperty: "setParameterActionsInput", GoGetter: "SetParameterActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tag", GoGetter: "Tag"},
			_jsii_.MemberProperty{JsiiProperty: "tagInput", GoGetter: "TagInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webhook", GoGetter: "Webhook"},
			_jsii_.MemberProperty{JsiiProperty: "webhookInput", GoGetter: "WebhookInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActions",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActionsList",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActionsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActionsOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameter", GoGetter: "Parameter"},
			_jsii_.MemberProperty{JsiiProperty: "parameterInput", GoGetter: "ParameterInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameter", GoMethod: "ResetParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorInitialPromptFulfillmentSetParameterActionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "initialPromptFulfillment", GoGetter: "InitialPromptFulfillment"},
			_jsii_.MemberProperty{JsiiProperty: "initialPromptFulfillmentInput", GoGetter: "InitialPromptFulfillmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "putInitialPromptFulfillment", GoMethod: "PutInitialPromptFulfillment"},
			_jsii_.MemberMethod{JsiiMethod: "putRepromptEventHandlers", GoMethod: "PutRepromptEventHandlers"},
			_jsii_.MemberProperty{JsiiProperty: "repromptEventHandlers", GoGetter: "RepromptEventHandlers"},
			_jsii_.MemberProperty{JsiiProperty: "repromptEventHandlersInput", GoGetter: "RepromptEventHandlersInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInitialPromptFulfillment", GoMethod: "ResetInitialPromptFulfillment"},
			_jsii_.MemberMethod{JsiiMethod: "resetRepromptEventHandlers", GoMethod: "ResetRepromptEventHandlers"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlers",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlers)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersList",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "event", GoGetter: "Event"},
			_jsii_.MemberProperty{JsiiProperty: "eventInput", GoGetter: "EventInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putTriggerFulfillment", GoMethod: "PutTriggerFulfillment"},
			_jsii_.MemberMethod{JsiiMethod: "resetEvent", GoMethod: "ResetEvent"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetFlow", GoMethod: "ResetTargetFlow"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetPage", GoMethod: "ResetTargetPage"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggerFulfillment", GoMethod: "ResetTriggerFulfillment"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "targetFlow", GoGetter: "TargetFlow"},
			_jsii_.MemberProperty{JsiiProperty: "targetFlowInput", GoGetter: "TargetFlowInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetPage", GoGetter: "TargetPage"},
			_jsii_.MemberProperty{JsiiProperty: "targetPageInput", GoGetter: "TargetPageInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggerFulfillment", GoGetter: "TriggerFulfillment"},
			_jsii_.MemberProperty{JsiiProperty: "triggerFulfillmentInput", GoGetter: "TriggerFulfillmentInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillment",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentConditionalCases",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentConditionalCases)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentConditionalCasesList",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentConditionalCasesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentConditionalCasesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentConditionalCasesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentConditionalCasesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cases", GoGetter: "Cases"},
			_jsii_.MemberProperty{JsiiProperty: "casesInput", GoGetter: "CasesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCases", GoMethod: "ResetCases"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentConditionalCasesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessages",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessages)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesConversationSuccess",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesConversationSuccess)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesConversationSuccessOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesConversationSuccessOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesConversationSuccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesList",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesLiveAgentHandoff",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesLiveAgentHandoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesLiveAgentHandoffOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesLiveAgentHandoffOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesLiveAgentHandoffOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesOutputAudioText",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesOutputAudioText)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetSsml", GoMethod: "ResetSsml"},
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ssml", GoGetter: "Ssml"},
			_jsii_.MemberProperty{JsiiProperty: "ssmlInput", GoGetter: "SsmlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesOutputAudioTextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "channel", GoGetter: "Channel"},
			_jsii_.MemberProperty{JsiiProperty: "channelInput", GoGetter: "ChannelInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conversationSuccess", GoGetter: "ConversationSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "conversationSuccessInput", GoGetter: "ConversationSuccessInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "liveAgentHandoff", GoGetter: "LiveAgentHandoff"},
			_jsii_.MemberProperty{JsiiProperty: "liveAgentHandoffInput", GoGetter: "LiveAgentHandoffInput"},
			_jsii_.MemberProperty{JsiiProperty: "outputAudioText", GoGetter: "OutputAudioText"},
			_jsii_.MemberProperty{JsiiProperty: "outputAudioTextInput", GoGetter: "OutputAudioTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "payload", GoGetter: "Payload"},
			_jsii_.MemberProperty{JsiiProperty: "payloadInput", GoGetter: "PayloadInput"},
			_jsii_.MemberProperty{JsiiProperty: "playAudio", GoGetter: "PlayAudio"},
			_jsii_.MemberProperty{JsiiProperty: "playAudioInput", GoGetter: "PlayAudioInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConversationSuccess", GoMethod: "PutConversationSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "putLiveAgentHandoff", GoMethod: "PutLiveAgentHandoff"},
			_jsii_.MemberMethod{JsiiMethod: "putOutputAudioText", GoMethod: "PutOutputAudioText"},
			_jsii_.MemberMethod{JsiiMethod: "putPlayAudio", GoMethod: "PutPlayAudio"},
			_jsii_.MemberMethod{JsiiMethod: "putTelephonyTransferCall", GoMethod: "PutTelephonyTransferCall"},
			_jsii_.MemberMethod{JsiiMethod: "putText", GoMethod: "PutText"},
			_jsii_.MemberMethod{JsiiMethod: "resetChannel", GoMethod: "ResetChannel"},
			_jsii_.MemberMethod{JsiiMethod: "resetConversationSuccess", GoMethod: "ResetConversationSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetLiveAgentHandoff", GoMethod: "ResetLiveAgentHandoff"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputAudioText", GoMethod: "ResetOutputAudioText"},
			_jsii_.MemberMethod{JsiiMethod: "resetPayload", GoMethod: "ResetPayload"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlayAudio", GoMethod: "ResetPlayAudio"},
			_jsii_.MemberMethod{JsiiMethod: "resetTelephonyTransferCall", GoMethod: "ResetTelephonyTransferCall"},
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "telephonyTransferCall", GoGetter: "TelephonyTransferCall"},
			_jsii_.MemberProperty{JsiiProperty: "telephonyTransferCallInput", GoGetter: "TelephonyTransferCallInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesPlayAudio",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesPlayAudio)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesPlayAudioOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesPlayAudioOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
			_jsii_.MemberProperty{JsiiProperty: "audioUri", GoGetter: "AudioUri"},
			_jsii_.MemberProperty{JsiiProperty: "audioUriInput", GoGetter: "AudioUriInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesPlayAudioOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesTelephonyTransferCall)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesTelephonyTransferCallOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesTelephonyTransferCallOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "phoneNumber", GoGetter: "PhoneNumber"},
			_jsii_.MemberProperty{JsiiProperty: "phoneNumberInput", GoGetter: "PhoneNumberInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesTelephonyTransferCallOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesText",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesText)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesTextOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesTextOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentMessagesTextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalCases", GoGetter: "ConditionalCases"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalCasesInput", GoGetter: "ConditionalCasesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "messages", GoGetter: "Messages"},
			_jsii_.MemberProperty{JsiiProperty: "messagesInput", GoGetter: "MessagesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConditionalCases", GoMethod: "PutConditionalCases"},
			_jsii_.MemberMethod{JsiiMethod: "putMessages", GoMethod: "PutMessages"},
			_jsii_.MemberMethod{JsiiMethod: "putSetParameterActions", GoMethod: "PutSetParameterActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetConditionalCases", GoMethod: "ResetConditionalCases"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessages", GoMethod: "ResetMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resetReturnPartialResponses", GoMethod: "ResetReturnPartialResponses"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetParameterActions", GoMethod: "ResetSetParameterActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTag", GoMethod: "ResetTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebhook", GoMethod: "ResetWebhook"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "returnPartialResponses", GoGetter: "ReturnPartialResponses"},
			_jsii_.MemberProperty{JsiiProperty: "returnPartialResponsesInput", GoGetter: "ReturnPartialResponsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "setParameterActions", GoGetter: "SetParameterActions"},
			_jsii_.MemberProperty{JsiiProperty: "setParameterActionsInput", GoGetter: "SetParameterActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tag", GoGetter: "Tag"},
			_jsii_.MemberProperty{JsiiProperty: "tagInput", GoGetter: "TagInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webhook", GoGetter: "Webhook"},
			_jsii_.MemberProperty{JsiiProperty: "webhookInput", GoGetter: "WebhookInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentSetParameterActions",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentSetParameterActions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentSetParameterActionsList",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentSetParameterActionsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentSetParameterActionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentSetParameterActionsOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentSetParameterActionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameter", GoGetter: "Parameter"},
			_jsii_.MemberProperty{JsiiProperty: "parameterInput", GoGetter: "ParameterInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameter", GoMethod: "ResetParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersFillBehaviorRepromptEventHandlersTriggerFulfillmentSetParameterActionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersList",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageFormParametersOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageFormParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "advancedSettings", GoGetter: "AdvancedSettings"},
			_jsii_.MemberProperty{JsiiProperty: "advancedSettingsInput", GoGetter: "AdvancedSettingsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "defaultValue", GoGetter: "DefaultValue"},
			_jsii_.MemberProperty{JsiiProperty: "defaultValueInput", GoGetter: "DefaultValueInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "entityType", GoGetter: "EntityType"},
			_jsii_.MemberProperty{JsiiProperty: "entityTypeInput", GoGetter: "EntityTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "fillBehavior", GoGetter: "FillBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "fillBehaviorInput", GoGetter: "FillBehaviorInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "isList", GoGetter: "IsList"},
			_jsii_.MemberProperty{JsiiProperty: "isListInput", GoGetter: "IsListInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAdvancedSettings", GoMethod: "PutAdvancedSettings"},
			_jsii_.MemberMethod{JsiiMethod: "putFillBehavior", GoMethod: "PutFillBehavior"},
			_jsii_.MemberProperty{JsiiProperty: "redact", GoGetter: "Redact"},
			_jsii_.MemberProperty{JsiiProperty: "redactInput", GoGetter: "RedactInput"},
			_jsii_.MemberProperty{JsiiProperty: "required", GoGetter: "Required"},
			_jsii_.MemberProperty{JsiiProperty: "requiredInput", GoGetter: "RequiredInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdvancedSettings", GoMethod: "ResetAdvancedSettings"},
			_jsii_.MemberMethod{JsiiMethod: "resetDefaultValue", GoMethod: "ResetDefaultValue"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEntityType", GoMethod: "ResetEntityType"},
			_jsii_.MemberMethod{JsiiMethod: "resetFillBehavior", GoMethod: "ResetFillBehavior"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsList", GoMethod: "ResetIsList"},
			_jsii_.MemberMethod{JsiiMethod: "resetRedact", GoMethod: "ResetRedact"},
			_jsii_.MemberMethod{JsiiMethod: "resetRequired", GoMethod: "ResetRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageFormParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTimeouts",
		reflect.TypeOf((*GoogleDialogflowCxPageTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTimeoutsOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutes",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutes)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesList",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "condition", GoGetter: "Condition"},
			_jsii_.MemberProperty{JsiiProperty: "conditionInput", GoGetter: "ConditionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "intent", GoGetter: "Intent"},
			_jsii_.MemberProperty{JsiiProperty: "intentInput", GoGetter: "IntentInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberMethod{JsiiMethod: "putTriggerFulfillment", GoMethod: "PutTriggerFulfillment"},
			_jsii_.MemberMethod{JsiiMethod: "resetCondition", GoMethod: "ResetCondition"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntent", GoMethod: "ResetIntent"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetFlow", GoMethod: "ResetTargetFlow"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetPage", GoMethod: "ResetTargetPage"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggerFulfillment", GoMethod: "ResetTriggerFulfillment"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "targetFlow", GoGetter: "TargetFlow"},
			_jsii_.MemberProperty{JsiiProperty: "targetFlowInput", GoGetter: "TargetFlowInput"},
			_jsii_.MemberProperty{JsiiProperty: "targetPage", GoGetter: "TargetPage"},
			_jsii_.MemberProperty{JsiiProperty: "targetPageInput", GoGetter: "TargetPageInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "triggerFulfillment", GoGetter: "TriggerFulfillment"},
			_jsii_.MemberProperty{JsiiProperty: "triggerFulfillmentInput", GoGetter: "TriggerFulfillmentInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillment",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillment)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentConditionalCases",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentConditionalCases)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentConditionalCasesList",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentConditionalCasesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentConditionalCasesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentConditionalCasesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentConditionalCasesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "cases", GoGetter: "Cases"},
			_jsii_.MemberProperty{JsiiProperty: "casesInput", GoGetter: "CasesInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCases", GoMethod: "ResetCases"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentConditionalCasesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessages",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessages)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccess",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccess)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccessOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccessOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesConversationSuccessOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesList",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoff)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoffOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoffOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesLiveAgentHandoffOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioText",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioText)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioTextOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioTextOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetSsml", GoMethod: "ResetSsml"},
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "ssml", GoGetter: "Ssml"},
			_jsii_.MemberProperty{JsiiProperty: "ssmlInput", GoGetter: "SsmlInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputAudioTextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "channel", GoGetter: "Channel"},
			_jsii_.MemberProperty{JsiiProperty: "channelInput", GoGetter: "ChannelInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conversationSuccess", GoGetter: "ConversationSuccess"},
			_jsii_.MemberProperty{JsiiProperty: "conversationSuccessInput", GoGetter: "ConversationSuccessInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "liveAgentHandoff", GoGetter: "LiveAgentHandoff"},
			_jsii_.MemberProperty{JsiiProperty: "liveAgentHandoffInput", GoGetter: "LiveAgentHandoffInput"},
			_jsii_.MemberProperty{JsiiProperty: "outputAudioText", GoGetter: "OutputAudioText"},
			_jsii_.MemberProperty{JsiiProperty: "outputAudioTextInput", GoGetter: "OutputAudioTextInput"},
			_jsii_.MemberProperty{JsiiProperty: "payload", GoGetter: "Payload"},
			_jsii_.MemberProperty{JsiiProperty: "payloadInput", GoGetter: "PayloadInput"},
			_jsii_.MemberProperty{JsiiProperty: "playAudio", GoGetter: "PlayAudio"},
			_jsii_.MemberProperty{JsiiProperty: "playAudioInput", GoGetter: "PlayAudioInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConversationSuccess", GoMethod: "PutConversationSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "putLiveAgentHandoff", GoMethod: "PutLiveAgentHandoff"},
			_jsii_.MemberMethod{JsiiMethod: "putOutputAudioText", GoMethod: "PutOutputAudioText"},
			_jsii_.MemberMethod{JsiiMethod: "putPlayAudio", GoMethod: "PutPlayAudio"},
			_jsii_.MemberMethod{JsiiMethod: "putTelephonyTransferCall", GoMethod: "PutTelephonyTransferCall"},
			_jsii_.MemberMethod{JsiiMethod: "putText", GoMethod: "PutText"},
			_jsii_.MemberMethod{JsiiMethod: "resetChannel", GoMethod: "ResetChannel"},
			_jsii_.MemberMethod{JsiiMethod: "resetConversationSuccess", GoMethod: "ResetConversationSuccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetLiveAgentHandoff", GoMethod: "ResetLiveAgentHandoff"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputAudioText", GoMethod: "ResetOutputAudioText"},
			_jsii_.MemberMethod{JsiiMethod: "resetPayload", GoMethod: "ResetPayload"},
			_jsii_.MemberMethod{JsiiMethod: "resetPlayAudio", GoMethod: "ResetPlayAudio"},
			_jsii_.MemberMethod{JsiiMethod: "resetTelephonyTransferCall", GoMethod: "ResetTelephonyTransferCall"},
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "telephonyTransferCall", GoGetter: "TelephonyTransferCall"},
			_jsii_.MemberProperty{JsiiProperty: "telephonyTransferCallInput", GoGetter: "TelephonyTransferCallInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudio",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudio)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudioOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudioOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
			_jsii_.MemberProperty{JsiiProperty: "audioUri", GoGetter: "AudioUri"},
			_jsii_.MemberProperty{JsiiProperty: "audioUriInput", GoGetter: "AudioUriInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesPlayAudioOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCall)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCallOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCallOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "phoneNumber", GoGetter: "PhoneNumber"},
			_jsii_.MemberProperty{JsiiProperty: "phoneNumberInput", GoGetter: "PhoneNumberInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTelephonyTransferCallOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesText",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesText)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTextOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTextOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowPlaybackInterruption", GoGetter: "AllowPlaybackInterruption"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetText", GoMethod: "ResetText"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "text", GoGetter: "Text"},
			_jsii_.MemberProperty{JsiiProperty: "textInput", GoGetter: "TextInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentMessagesTextOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalCases", GoGetter: "ConditionalCases"},
			_jsii_.MemberProperty{JsiiProperty: "conditionalCasesInput", GoGetter: "ConditionalCasesInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "messages", GoGetter: "Messages"},
			_jsii_.MemberProperty{JsiiProperty: "messagesInput", GoGetter: "MessagesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConditionalCases", GoMethod: "PutConditionalCases"},
			_jsii_.MemberMethod{JsiiMethod: "putMessages", GoMethod: "PutMessages"},
			_jsii_.MemberMethod{JsiiMethod: "putSetParameterActions", GoMethod: "PutSetParameterActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetConditionalCases", GoMethod: "ResetConditionalCases"},
			_jsii_.MemberMethod{JsiiMethod: "resetMessages", GoMethod: "ResetMessages"},
			_jsii_.MemberMethod{JsiiMethod: "resetReturnPartialResponses", GoMethod: "ResetReturnPartialResponses"},
			_jsii_.MemberMethod{JsiiMethod: "resetSetParameterActions", GoMethod: "ResetSetParameterActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTag", GoMethod: "ResetTag"},
			_jsii_.MemberMethod{JsiiMethod: "resetWebhook", GoMethod: "ResetWebhook"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "returnPartialResponses", GoGetter: "ReturnPartialResponses"},
			_jsii_.MemberProperty{JsiiProperty: "returnPartialResponsesInput", GoGetter: "ReturnPartialResponsesInput"},
			_jsii_.MemberProperty{JsiiProperty: "setParameterActions", GoGetter: "SetParameterActions"},
			_jsii_.MemberProperty{JsiiProperty: "setParameterActionsInput", GoGetter: "SetParameterActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "tag", GoGetter: "Tag"},
			_jsii_.MemberProperty{JsiiProperty: "tagInput", GoGetter: "TagInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "webhook", GoGetter: "Webhook"},
			_jsii_.MemberProperty{JsiiProperty: "webhookInput", GoGetter: "WebhookInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentSetParameterActions",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentSetParameterActions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentSetParameterActionsList",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentSetParameterActionsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentSetParameterActionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDialogflowCxPage.GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentSetParameterActionsOutputReference",
		reflect.TypeOf((*GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentSetParameterActionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "parameter", GoGetter: "Parameter"},
			_jsii_.MemberProperty{JsiiProperty: "parameterInput", GoGetter: "ParameterInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameter", GoMethod: "ResetParameter"},
			_jsii_.MemberMethod{JsiiMethod: "resetValue", GoMethod: "ResetValue"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDialogflowCxPageTransitionRoutesTriggerFulfillmentSetParameterActionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
