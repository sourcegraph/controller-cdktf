package googledocumentaiwarehousedocumentschema

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchema",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchema)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "documentIsFolder", GoGetter: "DocumentIsFolder"},
			_jsii_.MemberProperty{JsiiProperty: "documentIsFolderInput", GoGetter: "DocumentIsFolderInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "projectNumber", GoGetter: "ProjectNumber"},
			_jsii_.MemberProperty{JsiiProperty: "projectNumberInput", GoGetter: "ProjectNumberInput"},
			_jsii_.MemberProperty{JsiiProperty: "propertyDefinitions", GoGetter: "PropertyDefinitions"},
			_jsii_.MemberProperty{JsiiProperty: "propertyDefinitionsInput", GoGetter: "PropertyDefinitionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putPropertyDefinitions", GoMethod: "PutPropertyDefinitions"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDocumentIsFolder", GoMethod: "ResetDocumentIsFolder"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
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
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchema{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaConfig",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsDateTimeTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "possibleValues", GoGetter: "PossibleValues"},
			_jsii_.MemberProperty{JsiiProperty: "possibleValuesInput", GoGetter: "PossibleValuesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetValidationCheckDisabled", GoMethod: "ResetValidationCheckDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "validationCheckDisabled", GoGetter: "ValidationCheckDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "validationCheckDisabledInput", GoGetter: "ValidationCheckDisabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsEnumTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsFloatTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsIntegerTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsList",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsMapTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dateTimeTypeOptions", GoGetter: "DateTimeTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "dateTimeTypeOptionsInput", GoGetter: "DateTimeTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "enumTypeOptions", GoGetter: "EnumTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "enumTypeOptionsInput", GoGetter: "EnumTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "floatTypeOptions", GoGetter: "FloatTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "floatTypeOptionsInput", GoGetter: "FloatTypeOptionsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "integerTypeOptions", GoGetter: "IntegerTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "integerTypeOptionsInput", GoGetter: "IntegerTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isFilterable", GoGetter: "IsFilterable"},
			_jsii_.MemberProperty{JsiiProperty: "isFilterableInput", GoGetter: "IsFilterableInput"},
			_jsii_.MemberProperty{JsiiProperty: "isMetadata", GoGetter: "IsMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "isMetadataInput", GoGetter: "IsMetadataInput"},
			_jsii_.MemberProperty{JsiiProperty: "isRepeatable", GoGetter: "IsRepeatable"},
			_jsii_.MemberProperty{JsiiProperty: "isRepeatableInput", GoGetter: "IsRepeatableInput"},
			_jsii_.MemberProperty{JsiiProperty: "isRequired", GoGetter: "IsRequired"},
			_jsii_.MemberProperty{JsiiProperty: "isRequiredInput", GoGetter: "IsRequiredInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSearchable", GoGetter: "IsSearchable"},
			_jsii_.MemberProperty{JsiiProperty: "isSearchableInput", GoGetter: "IsSearchableInput"},
			_jsii_.MemberProperty{JsiiProperty: "mapTypeOptions", GoGetter: "MapTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "mapTypeOptionsInput", GoGetter: "MapTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "propertyTypeOptions", GoGetter: "PropertyTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "propertyTypeOptionsInput", GoGetter: "PropertyTypeOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDateTimeTypeOptions", GoMethod: "PutDateTimeTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEnumTypeOptions", GoMethod: "PutEnumTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putFloatTypeOptions", GoMethod: "PutFloatTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putIntegerTypeOptions", GoMethod: "PutIntegerTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putMapTypeOptions", GoMethod: "PutMapTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putPropertyTypeOptions", GoMethod: "PutPropertyTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putSchemaSources", GoMethod: "PutSchemaSources"},
			_jsii_.MemberMethod{JsiiMethod: "putTextTypeOptions", GoMethod: "PutTextTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putTimestampTypeOptions", GoMethod: "PutTimestampTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDateTimeTypeOptions", GoMethod: "ResetDateTimeTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnumTypeOptions", GoMethod: "ResetEnumTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetFloatTypeOptions", GoMethod: "ResetFloatTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegerTypeOptions", GoMethod: "ResetIntegerTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsFilterable", GoMethod: "ResetIsFilterable"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsMetadata", GoMethod: "ResetIsMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsRepeatable", GoMethod: "ResetIsRepeatable"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsRequired", GoMethod: "ResetIsRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsSearchable", GoMethod: "ResetIsSearchable"},
			_jsii_.MemberMethod{JsiiMethod: "resetMapTypeOptions", GoMethod: "ResetMapTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetPropertyTypeOptions", GoMethod: "ResetPropertyTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetrievalImportance", GoMethod: "ResetRetrievalImportance"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchemaSources", GoMethod: "ResetSchemaSources"},
			_jsii_.MemberMethod{JsiiMethod: "resetTextTypeOptions", GoMethod: "ResetTextTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimestampTypeOptions", GoMethod: "ResetTimestampTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retrievalImportance", GoGetter: "RetrievalImportance"},
			_jsii_.MemberProperty{JsiiProperty: "retrievalImportanceInput", GoGetter: "RetrievalImportanceInput"},
			_jsii_.MemberProperty{JsiiProperty: "schemaSources", GoGetter: "SchemaSources"},
			_jsii_.MemberProperty{JsiiProperty: "schemaSourcesInput", GoGetter: "SchemaSourcesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "textTypeOptions", GoGetter: "TextTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "textTypeOptionsInput", GoGetter: "TextTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "timestampTypeOptions", GoGetter: "TimestampTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "timestampTypeOptionsInput", GoGetter: "TimestampTypeOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "propertyDefinitions", GoGetter: "PropertyDefinitions"},
			_jsii_.MemberProperty{JsiiProperty: "propertyDefinitionsInput", GoGetter: "PropertyDefinitionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPropertyDefinitions", GoMethod: "PutPropertyDefinitions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitions)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsDateTimeTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsDateTimeTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsDateTimeTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsDateTimeTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsDateTimeTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsEnumTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsEnumTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsEnumTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsEnumTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "possibleValues", GoGetter: "PossibleValues"},
			_jsii_.MemberProperty{JsiiProperty: "possibleValuesInput", GoGetter: "PossibleValuesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetValidationCheckDisabled", GoMethod: "ResetValidationCheckDisabled"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "validationCheckDisabled", GoGetter: "ValidationCheckDisabled"},
			_jsii_.MemberProperty{JsiiProperty: "validationCheckDisabledInput", GoGetter: "ValidationCheckDisabledInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsEnumTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsFloatTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsFloatTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsFloatTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsFloatTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsFloatTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsIntegerTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsIntegerTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsIntegerTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsIntegerTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsIntegerTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsList",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsMapTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsMapTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsMapTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsMapTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsMapTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dateTimeTypeOptions", GoGetter: "DateTimeTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "dateTimeTypeOptionsInput", GoGetter: "DateTimeTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "enumTypeOptions", GoGetter: "EnumTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "enumTypeOptionsInput", GoGetter: "EnumTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "floatTypeOptions", GoGetter: "FloatTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "floatTypeOptionsInput", GoGetter: "FloatTypeOptionsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "integerTypeOptions", GoGetter: "IntegerTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "integerTypeOptionsInput", GoGetter: "IntegerTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isFilterable", GoGetter: "IsFilterable"},
			_jsii_.MemberProperty{JsiiProperty: "isFilterableInput", GoGetter: "IsFilterableInput"},
			_jsii_.MemberProperty{JsiiProperty: "isMetadata", GoGetter: "IsMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "isMetadataInput", GoGetter: "IsMetadataInput"},
			_jsii_.MemberProperty{JsiiProperty: "isRepeatable", GoGetter: "IsRepeatable"},
			_jsii_.MemberProperty{JsiiProperty: "isRepeatableInput", GoGetter: "IsRepeatableInput"},
			_jsii_.MemberProperty{JsiiProperty: "isRequired", GoGetter: "IsRequired"},
			_jsii_.MemberProperty{JsiiProperty: "isRequiredInput", GoGetter: "IsRequiredInput"},
			_jsii_.MemberProperty{JsiiProperty: "isSearchable", GoGetter: "IsSearchable"},
			_jsii_.MemberProperty{JsiiProperty: "isSearchableInput", GoGetter: "IsSearchableInput"},
			_jsii_.MemberProperty{JsiiProperty: "mapTypeOptions", GoGetter: "MapTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "mapTypeOptionsInput", GoGetter: "MapTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDateTimeTypeOptions", GoMethod: "PutDateTimeTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putEnumTypeOptions", GoMethod: "PutEnumTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putFloatTypeOptions", GoMethod: "PutFloatTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putIntegerTypeOptions", GoMethod: "PutIntegerTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putMapTypeOptions", GoMethod: "PutMapTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putSchemaSources", GoMethod: "PutSchemaSources"},
			_jsii_.MemberMethod{JsiiMethod: "putTextTypeOptions", GoMethod: "PutTextTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "putTimestampTypeOptions", GoMethod: "PutTimestampTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDateTimeTypeOptions", GoMethod: "ResetDateTimeTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetDisplayName", GoMethod: "ResetDisplayName"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnumTypeOptions", GoMethod: "ResetEnumTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetFloatTypeOptions", GoMethod: "ResetFloatTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetIntegerTypeOptions", GoMethod: "ResetIntegerTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsFilterable", GoMethod: "ResetIsFilterable"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsMetadata", GoMethod: "ResetIsMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsRepeatable", GoMethod: "ResetIsRepeatable"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsRequired", GoMethod: "ResetIsRequired"},
			_jsii_.MemberMethod{JsiiMethod: "resetIsSearchable", GoMethod: "ResetIsSearchable"},
			_jsii_.MemberMethod{JsiiMethod: "resetMapTypeOptions", GoMethod: "ResetMapTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetRetrievalImportance", GoMethod: "ResetRetrievalImportance"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchemaSources", GoMethod: "ResetSchemaSources"},
			_jsii_.MemberMethod{JsiiMethod: "resetTextTypeOptions", GoMethod: "ResetTextTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimestampTypeOptions", GoMethod: "ResetTimestampTypeOptions"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "retrievalImportance", GoGetter: "RetrievalImportance"},
			_jsii_.MemberProperty{JsiiProperty: "retrievalImportanceInput", GoGetter: "RetrievalImportanceInput"},
			_jsii_.MemberProperty{JsiiProperty: "schemaSources", GoGetter: "SchemaSources"},
			_jsii_.MemberProperty{JsiiProperty: "schemaSourcesInput", GoGetter: "SchemaSourcesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "textTypeOptions", GoGetter: "TextTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "textTypeOptionsInput", GoGetter: "TextTypeOptionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "timestampTypeOptions", GoGetter: "TimestampTypeOptions"},
			_jsii_.MemberProperty{JsiiProperty: "timestampTypeOptionsInput", GoGetter: "TimestampTypeOptionsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsSchemaSources",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsSchemaSources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsSchemaSourcesList",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsSchemaSourcesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsSchemaSourcesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsSchemaSourcesOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsSchemaSourcesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "processorType", GoGetter: "ProcessorType"},
			_jsii_.MemberProperty{JsiiProperty: "processorTypeInput", GoGetter: "ProcessorTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessorType", GoMethod: "ResetProcessorType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsSchemaSourcesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsTextTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsTextTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsTextTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsTextTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsTextTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsTimestampTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsTimestampTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsTimestampTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsTimestampTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsPropertyTypeOptionsPropertyDefinitionsTimestampTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSources",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSources)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesList",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "processorType", GoGetter: "ProcessorType"},
			_jsii_.MemberProperty{JsiiProperty: "processorTypeInput", GoGetter: "ProcessorTypeInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetName", GoMethod: "ResetName"},
			_jsii_.MemberMethod{JsiiMethod: "resetProcessorType", GoMethod: "ResetProcessorType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsSchemaSourcesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTextTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptionsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptionsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaPropertyDefinitionsTimestampTypeOptionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaTimeouts",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDocumentAiWarehouseDocumentSchema.GoogleDocumentAiWarehouseDocumentSchemaTimeoutsOutputReference",
		reflect.TypeOf((*GoogleDocumentAiWarehouseDocumentSchemaTimeoutsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDocumentAiWarehouseDocumentSchemaTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
