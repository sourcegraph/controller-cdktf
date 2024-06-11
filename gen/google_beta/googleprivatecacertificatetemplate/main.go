package googleprivatecacertificatetemplate

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplate",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createTime", GoGetter: "CreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "effectiveLabels", GoGetter: "EffectiveLabels"},
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
			_jsii_.MemberProperty{JsiiProperty: "identityConstraints", GoGetter: "IdentityConstraints"},
			_jsii_.MemberProperty{JsiiProperty: "identityConstraintsInput", GoGetter: "IdentityConstraintsInput"},
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "labelsInput", GoGetter: "LabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "maximumLifetime", GoGetter: "MaximumLifetime"},
			_jsii_.MemberProperty{JsiiProperty: "maximumLifetimeInput", GoGetter: "MaximumLifetimeInput"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughExtensions", GoGetter: "PassthroughExtensions"},
			_jsii_.MemberProperty{JsiiProperty: "passthroughExtensionsInput", GoGetter: "PassthroughExtensionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "predefinedValues", GoGetter: "PredefinedValues"},
			_jsii_.MemberProperty{JsiiProperty: "predefinedValuesInput", GoGetter: "PredefinedValuesInput"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putIdentityConstraints", GoMethod: "PutIdentityConstraints"},
			_jsii_.MemberMethod{JsiiMethod: "putPassthroughExtensions", GoMethod: "PutPassthroughExtensions"},
			_jsii_.MemberMethod{JsiiMethod: "putPredefinedValues", GoMethod: "PutPredefinedValues"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdentityConstraints", GoMethod: "ResetIdentityConstraints"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabels", GoMethod: "ResetLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaximumLifetime", GoMethod: "ResetMaximumLifetime"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetPassthroughExtensions", GoMethod: "ResetPassthroughExtensions"},
			_jsii_.MemberMethod{JsiiMethod: "resetPredefinedValues", GoMethod: "ResetPredefinedValues"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformLabels", GoGetter: "TerraformLabels"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "updateTime", GoGetter: "UpdateTime"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplateConfig",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplateConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplateIdentityConstraints",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplateIdentityConstraints)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplateIdentityConstraintsCelExpression",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplateIdentityConstraintsCelExpression)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplateIdentityConstraintsCelExpressionOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplateIdentityConstraintsCelExpressionOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "expression", GoGetter: "Expression"},
			_jsii_.MemberProperty{JsiiProperty: "expressionInput", GoGetter: "ExpressionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetExpression", GoMethod: "ResetExpression"},
			_jsii_.MemberMethod{JsiiMethod: "resetLocation", GoMethod: "ResetLocation"},
			_jsii_.MemberMethod{JsiiMethod: "resetTitle", GoMethod: "ResetTitle"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "title", GoGetter: "Title"},
			_jsii_.MemberProperty{JsiiProperty: "titleInput", GoGetter: "TitleInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplateIdentityConstraintsCelExpressionOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplateIdentityConstraintsOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplateIdentityConstraintsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "allowSubjectAltNamesPassthrough", GoGetter: "AllowSubjectAltNamesPassthrough"},
			_jsii_.MemberProperty{JsiiProperty: "allowSubjectAltNamesPassthroughInput", GoGetter: "AllowSubjectAltNamesPassthroughInput"},
			_jsii_.MemberProperty{JsiiProperty: "allowSubjectPassthrough", GoGetter: "AllowSubjectPassthrough"},
			_jsii_.MemberProperty{JsiiProperty: "allowSubjectPassthroughInput", GoGetter: "AllowSubjectPassthroughInput"},
			_jsii_.MemberProperty{JsiiProperty: "celExpression", GoGetter: "CelExpression"},
			_jsii_.MemberProperty{JsiiProperty: "celExpressionInput", GoGetter: "CelExpressionInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putCelExpression", GoMethod: "PutCelExpression"},
			_jsii_.MemberMethod{JsiiMethod: "resetCelExpression", GoMethod: "ResetCelExpression"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplateIdentityConstraintsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePassthroughExtensions",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePassthroughExtensions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsList",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsList)(nil)).Elem(),
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
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "objectIdPath", GoGetter: "ObjectIdPath"},
			_jsii_.MemberProperty{JsiiProperty: "objectIdPathInput", GoGetter: "ObjectIdPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePassthroughExtensionsAdditionalExtensionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePassthroughExtensionsOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePassthroughExtensionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalExtensions", GoGetter: "AdditionalExtensions"},
			_jsii_.MemberProperty{JsiiProperty: "additionalExtensionsInput", GoGetter: "AdditionalExtensionsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "knownExtensions", GoGetter: "KnownExtensions"},
			_jsii_.MemberProperty{JsiiProperty: "knownExtensionsInput", GoGetter: "KnownExtensionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAdditionalExtensions", GoMethod: "PutAdditionalExtensions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalExtensions", GoMethod: "ResetAdditionalExtensions"},
			_jsii_.MemberMethod{JsiiMethod: "resetKnownExtensions", GoMethod: "ResetKnownExtensions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePassthroughExtensionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValues",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValues)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsList",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsList)(nil)).Elem(),
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
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectId)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "objectIdPath", GoGetter: "ObjectIdPath"},
			_jsii_.MemberProperty{JsiiProperty: "objectIdPathInput", GoGetter: "ObjectIdPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsObjectIdOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "critical", GoGetter: "Critical"},
			_jsii_.MemberProperty{JsiiProperty: "criticalInput", GoGetter: "CriticalInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "objectId", GoGetter: "ObjectId"},
			_jsii_.MemberProperty{JsiiProperty: "objectIdInput", GoGetter: "ObjectIdInput"},
			_jsii_.MemberMethod{JsiiMethod: "putObjectId", GoMethod: "PutObjectId"},
			_jsii_.MemberMethod{JsiiMethod: "resetCritical", GoMethod: "ResetCritical"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "value", GoGetter: "Value"},
			_jsii_.MemberProperty{JsiiProperty: "valueInput", GoGetter: "ValueInput"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesAdditionalExtensionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesCaOptions",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesCaOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "isCa", GoGetter: "IsCa"},
			_jsii_.MemberProperty{JsiiProperty: "isCaInput", GoGetter: "IsCaInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxIssuerPathLength", GoGetter: "MaxIssuerPathLength"},
			_jsii_.MemberProperty{JsiiProperty: "maxIssuerPathLengthInput", GoGetter: "MaxIssuerPathLengthInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsCa", GoMethod: "ResetIsCa"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxIssuerPathLength", GoMethod: "ResetMaxIssuerPathLength"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesCaOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsage",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsage)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsage)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "certSign", GoGetter: "CertSign"},
			_jsii_.MemberProperty{JsiiProperty: "certSignInput", GoGetter: "CertSignInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "contentCommitment", GoGetter: "ContentCommitment"},
			_jsii_.MemberProperty{JsiiProperty: "contentCommitmentInput", GoGetter: "ContentCommitmentInput"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "crlSign", GoGetter: "CrlSign"},
			_jsii_.MemberProperty{JsiiProperty: "crlSignInput", GoGetter: "CrlSignInput"},
			_jsii_.MemberProperty{JsiiProperty: "dataEncipherment", GoGetter: "DataEncipherment"},
			_jsii_.MemberProperty{JsiiProperty: "dataEnciphermentInput", GoGetter: "DataEnciphermentInput"},
			_jsii_.MemberProperty{JsiiProperty: "decipherOnly", GoGetter: "DecipherOnly"},
			_jsii_.MemberProperty{JsiiProperty: "decipherOnlyInput", GoGetter: "DecipherOnlyInput"},
			_jsii_.MemberProperty{JsiiProperty: "digitalSignature", GoGetter: "DigitalSignature"},
			_jsii_.MemberProperty{JsiiProperty: "digitalSignatureInput", GoGetter: "DigitalSignatureInput"},
			_jsii_.MemberProperty{JsiiProperty: "encipherOnly", GoGetter: "EncipherOnly"},
			_jsii_.MemberProperty{JsiiProperty: "encipherOnlyInput", GoGetter: "EncipherOnlyInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "keyAgreement", GoGetter: "KeyAgreement"},
			_jsii_.MemberProperty{JsiiProperty: "keyAgreementInput", GoGetter: "KeyAgreementInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyEncipherment", GoGetter: "KeyEncipherment"},
			_jsii_.MemberProperty{JsiiProperty: "keyEnciphermentInput", GoGetter: "KeyEnciphermentInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCertSign", GoMethod: "ResetCertSign"},
			_jsii_.MemberMethod{JsiiMethod: "resetContentCommitment", GoMethod: "ResetContentCommitment"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrlSign", GoMethod: "ResetCrlSign"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataEncipherment", GoMethod: "ResetDataEncipherment"},
			_jsii_.MemberMethod{JsiiMethod: "resetDecipherOnly", GoMethod: "ResetDecipherOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetDigitalSignature", GoMethod: "ResetDigitalSignature"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncipherOnly", GoMethod: "ResetEncipherOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyAgreement", GoMethod: "ResetKeyAgreement"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyEncipherment", GoMethod: "ResetKeyEncipherment"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageBaseKeyUsageOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsage)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientAuth", GoGetter: "ClientAuth"},
			_jsii_.MemberProperty{JsiiProperty: "clientAuthInput", GoGetter: "ClientAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "codeSigning", GoGetter: "CodeSigning"},
			_jsii_.MemberProperty{JsiiProperty: "codeSigningInput", GoGetter: "CodeSigningInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "emailProtection", GoGetter: "EmailProtection"},
			_jsii_.MemberProperty{JsiiProperty: "emailProtectionInput", GoGetter: "EmailProtectionInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "ocspSigning", GoGetter: "OcspSigning"},
			_jsii_.MemberProperty{JsiiProperty: "ocspSigningInput", GoGetter: "OcspSigningInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientAuth", GoMethod: "ResetClientAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetCodeSigning", GoMethod: "ResetCodeSigning"},
			_jsii_.MemberMethod{JsiiMethod: "resetEmailProtection", GoMethod: "ResetEmailProtection"},
			_jsii_.MemberMethod{JsiiMethod: "resetOcspSigning", GoMethod: "ResetOcspSigning"},
			_jsii_.MemberMethod{JsiiMethod: "resetServerAuth", GoMethod: "ResetServerAuth"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeStamping", GoMethod: "ResetTimeStamping"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serverAuth", GoGetter: "ServerAuth"},
			_jsii_.MemberProperty{JsiiProperty: "serverAuthInput", GoGetter: "ServerAuthInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "timeStamping", GoGetter: "TimeStamping"},
			_jsii_.MemberProperty{JsiiProperty: "timeStampingInput", GoGetter: "TimeStampingInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageExtendedKeyUsageOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "baseKeyUsage", GoGetter: "BaseKeyUsage"},
			_jsii_.MemberProperty{JsiiProperty: "baseKeyUsageInput", GoGetter: "BaseKeyUsageInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "extendedKeyUsage", GoGetter: "ExtendedKeyUsage"},
			_jsii_.MemberProperty{JsiiProperty: "extendedKeyUsageInput", GoGetter: "ExtendedKeyUsageInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putBaseKeyUsage", GoMethod: "PutBaseKeyUsage"},
			_jsii_.MemberMethod{JsiiMethod: "putExtendedKeyUsage", GoMethod: "PutExtendedKeyUsage"},
			_jsii_.MemberMethod{JsiiMethod: "putUnknownExtendedKeyUsages", GoMethod: "PutUnknownExtendedKeyUsages"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaseKeyUsage", GoMethod: "ResetBaseKeyUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetExtendedKeyUsage", GoMethod: "ResetExtendedKeyUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetUnknownExtendedKeyUsages", GoMethod: "ResetUnknownExtendedKeyUsages"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "unknownExtendedKeyUsages", GoGetter: "UnknownExtendedKeyUsages"},
			_jsii_.MemberProperty{JsiiProperty: "unknownExtendedKeyUsagesInput", GoGetter: "UnknownExtendedKeyUsagesInput"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsages)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesList",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesList)(nil)).Elem(),
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
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "objectIdPath", GoGetter: "ObjectIdPath"},
			_jsii_.MemberProperty{JsiiProperty: "objectIdPathInput", GoGetter: "ObjectIdPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesKeyUsageUnknownExtendedKeyUsagesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "additionalExtensions", GoGetter: "AdditionalExtensions"},
			_jsii_.MemberProperty{JsiiProperty: "additionalExtensionsInput", GoGetter: "AdditionalExtensionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "aiaOcspServers", GoGetter: "AiaOcspServers"},
			_jsii_.MemberProperty{JsiiProperty: "aiaOcspServersInput", GoGetter: "AiaOcspServersInput"},
			_jsii_.MemberProperty{JsiiProperty: "caOptions", GoGetter: "CaOptions"},
			_jsii_.MemberProperty{JsiiProperty: "caOptionsInput", GoGetter: "CaOptionsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "keyUsage", GoGetter: "KeyUsage"},
			_jsii_.MemberProperty{JsiiProperty: "keyUsageInput", GoGetter: "KeyUsageInput"},
			_jsii_.MemberProperty{JsiiProperty: "policyIds", GoGetter: "PolicyIds"},
			_jsii_.MemberProperty{JsiiProperty: "policyIdsInput", GoGetter: "PolicyIdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAdditionalExtensions", GoMethod: "PutAdditionalExtensions"},
			_jsii_.MemberMethod{JsiiMethod: "putCaOptions", GoMethod: "PutCaOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putKeyUsage", GoMethod: "PutKeyUsage"},
			_jsii_.MemberMethod{JsiiMethod: "putPolicyIds", GoMethod: "PutPolicyIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetAdditionalExtensions", GoMethod: "ResetAdditionalExtensions"},
			_jsii_.MemberMethod{JsiiMethod: "resetAiaOcspServers", GoMethod: "ResetAiaOcspServers"},
			_jsii_.MemberMethod{JsiiMethod: "resetCaOptions", GoMethod: "ResetCaOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyUsage", GoMethod: "ResetKeyUsage"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicyIds", GoMethod: "ResetPolicyIds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIds",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIds)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIdsList",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIdsList)(nil)).Elem(),
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
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIdsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIdsOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIdsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "objectIdPath", GoGetter: "ObjectIdPath"},
			_jsii_.MemberProperty{JsiiProperty: "objectIdPathInput", GoGetter: "ObjectIdPathInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GooglePrivatecaCertificateTemplatePredefinedValuesPolicyIdsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplateTimeouts",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplateTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googlePrivatecaCertificateTemplate.GooglePrivatecaCertificateTemplateTimeoutsOutputReference",
		reflect.TypeOf((*GooglePrivatecaCertificateTemplateTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_GooglePrivatecaCertificateTemplateTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
