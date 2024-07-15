package googledatastreamstream

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStream",
		reflect.TypeOf((*GoogleDatastreamStream)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "backfillAll", GoGetter: "BackfillAll"},
			_jsii_.MemberProperty{JsiiProperty: "backfillAllInput", GoGetter: "BackfillAllInput"},
			_jsii_.MemberProperty{JsiiProperty: "backfillNone", GoGetter: "BackfillNone"},
			_jsii_.MemberProperty{JsiiProperty: "backfillNoneInput", GoGetter: "BackfillNoneInput"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createWithoutValidation", GoGetter: "CreateWithoutValidation"},
			_jsii_.MemberProperty{JsiiProperty: "createWithoutValidationInput", GoGetter: "CreateWithoutValidationInput"},
			_jsii_.MemberProperty{JsiiProperty: "customerManagedEncryptionKey", GoGetter: "CustomerManagedEncryptionKey"},
			_jsii_.MemberProperty{JsiiProperty: "customerManagedEncryptionKeyInput", GoGetter: "CustomerManagedEncryptionKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "desiredState", GoGetter: "DesiredState"},
			_jsii_.MemberProperty{JsiiProperty: "desiredStateInput", GoGetter: "DesiredStateInput"},
			_jsii_.MemberProperty{JsiiProperty: "destinationConfig", GoGetter: "DestinationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "destinationConfigInput", GoGetter: "DestinationConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "displayNameInput", GoGetter: "DisplayNameInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idInput", GoGetter: "IdInput"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "labelsInput", GoGetter: "LabelsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putBackfillAll", GoMethod: "PutBackfillAll"},
			_jsii_.MemberMethod{JsiiMethod: "putBackfillNone", GoMethod: "PutBackfillNone"},
			_jsii_.MemberMethod{JsiiMethod: "putDestinationConfig", GoMethod: "PutDestinationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSourceConfig", GoMethod: "PutSourceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackfillAll", GoMethod: "ResetBackfillAll"},
			_jsii_.MemberMethod{JsiiMethod: "resetBackfillNone", GoMethod: "ResetBackfillNone"},
			_jsii_.MemberMethod{JsiiMethod: "resetCreateWithoutValidation", GoMethod: "ResetCreateWithoutValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetCustomerManagedEncryptionKey", GoMethod: "ResetCustomerManagedEncryptionKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetDesiredState", GoMethod: "ResetDesiredState"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabels", GoMethod: "ResetLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "sourceConfig", GoGetter: "SourceConfig"},
			_jsii_.MemberProperty{JsiiProperty: "sourceConfigInput", GoGetter: "SourceConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "state", GoGetter: "State"},
			_jsii_.MemberProperty{JsiiProperty: "streamId", GoGetter: "StreamId"},
			_jsii_.MemberProperty{JsiiProperty: "streamIdInput", GoGetter: "StreamIdInput"},
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
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStream{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAll",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAll)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjects",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjects)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabases",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabases)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTables",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTables)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesMysqlColumns",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesMysqlColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesMysqlColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesMysqlColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesMysqlColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesMysqlColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesMysqlColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "collation", GoGetter: "Collation"},
			_jsii_.MemberProperty{JsiiProperty: "collationInput", GoGetter: "CollationInput"},
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "nullableInput", GoGetter: "NullableInput"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPositionInput", GoGetter: "OrdinalPositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyInput", GoGetter: "PrimaryKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCollation", GoMethod: "ResetCollation"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullable", GoMethod: "ResetNullable"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrdinalPosition", GoMethod: "ResetOrdinalPosition"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryKey", GoMethod: "ResetPrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesMysqlColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlColumns", GoGetter: "MysqlColumns"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlColumnsInput", GoGetter: "MysqlColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlColumns", GoMethod: "PutMysqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetMysqlColumns", GoMethod: "ResetMysqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesMysqlTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInput", GoGetter: "DatabaseInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlTables", GoGetter: "MysqlTables"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlTablesInput", GoGetter: "MysqlTablesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlTables", GoMethod: "PutMysqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetMysqlTables", GoMethod: "ResetMysqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsMysqlDatabasesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlDatabases", GoGetter: "MysqlDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlDatabasesInput", GoGetter: "MysqlDatabasesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlDatabases", GoMethod: "PutMysqlDatabases"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllMysqlExcludedObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjects",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjects)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemas",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTables",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTables)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOracleColumns",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOracleColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOracleColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOracleColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOracleColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOracleColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOracleColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "encoding", GoGetter: "Encoding"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "precision", GoGetter: "Precision"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scale", GoGetter: "Scale"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOracleColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "oracleColumns", GoGetter: "OracleColumns"},
			_jsii_.MemberProperty{JsiiProperty: "oracleColumnsInput", GoGetter: "OracleColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleColumns", GoMethod: "PutOracleColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetOracleColumns", GoMethod: "ResetOracleColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOracleTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "oracleTables", GoGetter: "OracleTables"},
			_jsii_.MemberProperty{JsiiProperty: "oracleTablesInput", GoGetter: "OracleTablesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleTables", GoMethod: "PutOracleTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetOracleTables", GoMethod: "ResetOracleTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOracleSchemasOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "oracleSchemas", GoGetter: "OracleSchemas"},
			_jsii_.MemberProperty{JsiiProperty: "oracleSchemasInput", GoGetter: "OracleSchemasInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleSchemas", GoMethod: "PutOracleSchemas"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllOracleExcludedObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlExcludedObjects", GoGetter: "MysqlExcludedObjects"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlExcludedObjectsInput", GoGetter: "MysqlExcludedObjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "oracleExcludedObjects", GoGetter: "OracleExcludedObjects"},
			_jsii_.MemberProperty{JsiiProperty: "oracleExcludedObjectsInput", GoGetter: "OracleExcludedObjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlExcludedObjects", GoGetter: "PostgresqlExcludedObjects"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlExcludedObjectsInput", GoGetter: "PostgresqlExcludedObjectsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlExcludedObjects", GoMethod: "PutMysqlExcludedObjects"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleExcludedObjects", GoMethod: "PutOracleExcludedObjects"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlExcludedObjects", GoMethod: "PutPostgresqlExcludedObjects"},
			_jsii_.MemberMethod{JsiiMethod: "putSqlServerExcludedObjects", GoMethod: "PutSqlServerExcludedObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetMysqlExcludedObjects", GoMethod: "ResetMysqlExcludedObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetOracleExcludedObjects", GoMethod: "ResetOracleExcludedObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostgresqlExcludedObjects", GoMethod: "ResetPostgresqlExcludedObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetSqlServerExcludedObjects", GoMethod: "ResetSqlServerExcludedObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sqlServerExcludedObjects", GoGetter: "SqlServerExcludedObjects"},
			_jsii_.MemberProperty{JsiiProperty: "sqlServerExcludedObjectsInput", GoGetter: "SqlServerExcludedObjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjects",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjects)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "postgresqlSchemas", GoGetter: "PostgresqlSchemas"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlSchemasInput", GoGetter: "PostgresqlSchemasInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlSchemas", GoMethod: "PutPostgresqlSchemas"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemas",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "postgresqlTables", GoGetter: "PostgresqlTables"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlTablesInput", GoGetter: "PostgresqlTablesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlTables", GoMethod: "PutPostgresqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostgresqlTables", GoMethod: "ResetPostgresqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTables",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTables)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "postgresqlColumns", GoGetter: "PostgresqlColumns"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlColumnsInput", GoGetter: "PostgresqlColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlColumns", GoMethod: "PutPostgresqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostgresqlColumns", GoMethod: "ResetPostgresqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumns",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "nullableInput", GoGetter: "NullableInput"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPositionInput", GoGetter: "OrdinalPositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "precision", GoGetter: "Precision"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyInput", GoGetter: "PrimaryKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullable", GoMethod: "ResetNullable"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrdinalPosition", GoMethod: "ResetOrdinalPosition"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryKey", GoMethod: "ResetPrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scale", GoGetter: "Scale"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllPostgresqlExcludedObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjects",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjects)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putSchemas", GoMethod: "PutSchemas"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schemas", GoGetter: "Schemas"},
			_jsii_.MemberProperty{JsiiProperty: "schemasInput", GoGetter: "SchemasInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemas",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putTables", GoMethod: "PutTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetTables", GoMethod: "ResetTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "tables", GoGetter: "Tables"},
			_jsii_.MemberProperty{JsiiProperty: "tablesInput", GoGetter: "TablesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTables",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesColumns",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "precision", GoGetter: "Precision"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scale", GoGetter: "Scale"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "columns", GoGetter: "Columns"},
			_jsii_.MemberProperty{JsiiProperty: "columnsInput", GoGetter: "ColumnsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putColumns", GoMethod: "PutColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumns", GoMethod: "ResetColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamBackfillAllSqlServerExcludedObjectsSchemasTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillNone",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillNone)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamBackfillNoneOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamBackfillNoneOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamBackfillNoneOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamConfig",
		reflect.TypeOf((*GoogleDatastreamStreamConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfig",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataFreshness", GoGetter: "DataFreshness"},
			_jsii_.MemberProperty{JsiiProperty: "dataFreshnessInput", GoGetter: "DataFreshnessInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putSingleTargetDataset", GoMethod: "PutSingleTargetDataset"},
			_jsii_.MemberMethod{JsiiMethod: "putSourceHierarchyDatasets", GoMethod: "PutSourceHierarchyDatasets"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataFreshness", GoMethod: "ResetDataFreshness"},
			_jsii_.MemberMethod{JsiiMethod: "resetSingleTargetDataset", GoMethod: "ResetSingleTargetDataset"},
			_jsii_.MemberMethod{JsiiMethod: "resetSourceHierarchyDatasets", GoMethod: "ResetSourceHierarchyDatasets"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "singleTargetDataset", GoGetter: "SingleTargetDataset"},
			_jsii_.MemberProperty{JsiiProperty: "singleTargetDatasetInput", GoGetter: "SingleTargetDatasetInput"},
			_jsii_.MemberProperty{JsiiProperty: "sourceHierarchyDatasets", GoGetter: "SourceHierarchyDatasets"},
			_jsii_.MemberProperty{JsiiProperty: "sourceHierarchyDatasetsInput", GoGetter: "SourceHierarchyDatasetsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDataset)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDatasetOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDatasetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "datasetId", GoGetter: "DatasetId"},
			_jsii_.MemberProperty{JsiiProperty: "datasetIdInput", GoGetter: "DatasetIdInput"},
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
			j := jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSingleTargetDatasetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasets)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplate",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplate)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplateOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplateOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "datasetIdPrefix", GoGetter: "DatasetIdPrefix"},
			_jsii_.MemberProperty{JsiiProperty: "datasetIdPrefixInput", GoGetter: "DatasetIdPrefixInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyName", GoGetter: "KmsKeyName"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyNameInput", GoGetter: "KmsKeyNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetDatasetIdPrefix", GoMethod: "ResetDatasetIdPrefix"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKeyName", GoMethod: "ResetKmsKeyName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsDatasetTemplateOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "datasetTemplate", GoGetter: "DatasetTemplate"},
			_jsii_.MemberProperty{JsiiProperty: "datasetTemplateInput", GoGetter: "DatasetTemplateInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putDatasetTemplate", GoMethod: "PutDatasetTemplate"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamDestinationConfigBigqueryDestinationConfigSourceHierarchyDatasetsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigGcsDestinationConfig",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigGcsDestinationConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormat",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormat)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormatOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormatOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamDestinationConfigGcsDestinationConfigAvroFileFormatOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormat",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormat)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormatOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormatOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberProperty{JsiiProperty: "compression", GoGetter: "Compression"},
			_jsii_.MemberProperty{JsiiProperty: "compressionInput", GoGetter: "CompressionInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetCompression", GoMethod: "ResetCompression"},
			_jsii_.MemberMethod{JsiiMethod: "resetSchemaFileFormat", GoMethod: "ResetSchemaFileFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schemaFileFormat", GoGetter: "SchemaFileFormat"},
			_jsii_.MemberProperty{JsiiProperty: "schemaFileFormatInput", GoGetter: "SchemaFileFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamDestinationConfigGcsDestinationConfigJsonFileFormatOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigGcsDestinationConfigOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigGcsDestinationConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "avroFileFormat", GoGetter: "AvroFileFormat"},
			_jsii_.MemberProperty{JsiiProperty: "avroFileFormatInput", GoGetter: "AvroFileFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileRotationInterval", GoGetter: "FileRotationInterval"},
			_jsii_.MemberProperty{JsiiProperty: "fileRotationIntervalInput", GoGetter: "FileRotationIntervalInput"},
			_jsii_.MemberProperty{JsiiProperty: "fileRotationMb", GoGetter: "FileRotationMb"},
			_jsii_.MemberProperty{JsiiProperty: "fileRotationMbInput", GoGetter: "FileRotationMbInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "jsonFileFormat", GoGetter: "JsonFileFormat"},
			_jsii_.MemberProperty{JsiiProperty: "jsonFileFormatInput", GoGetter: "JsonFileFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "path", GoGetter: "Path"},
			_jsii_.MemberProperty{JsiiProperty: "pathInput", GoGetter: "PathInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAvroFileFormat", GoMethod: "PutAvroFileFormat"},
			_jsii_.MemberMethod{JsiiMethod: "putJsonFileFormat", GoMethod: "PutJsonFileFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetAvroFileFormat", GoMethod: "ResetAvroFileFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileRotationInterval", GoMethod: "ResetFileRotationInterval"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileRotationMb", GoMethod: "ResetFileRotationMb"},
			_jsii_.MemberMethod{JsiiMethod: "resetJsonFileFormat", GoMethod: "ResetJsonFileFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetPath", GoMethod: "ResetPath"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamDestinationConfigGcsDestinationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamDestinationConfigOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamDestinationConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDestinationConfig", GoGetter: "BigqueryDestinationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "bigqueryDestinationConfigInput", GoGetter: "BigqueryDestinationConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "destinationConnectionProfile", GoGetter: "DestinationConnectionProfile"},
			_jsii_.MemberProperty{JsiiProperty: "destinationConnectionProfileInput", GoGetter: "DestinationConnectionProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "gcsDestinationConfig", GoGetter: "GcsDestinationConfig"},
			_jsii_.MemberProperty{JsiiProperty: "gcsDestinationConfigInput", GoGetter: "GcsDestinationConfigInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putBigqueryDestinationConfig", GoMethod: "PutBigqueryDestinationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putGcsDestinationConfig", GoMethod: "PutGcsDestinationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetBigqueryDestinationConfig", GoMethod: "ResetBigqueryDestinationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetGcsDestinationConfig", GoMethod: "ResetGcsDestinationConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamDestinationConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfig",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfig",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjects)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabases",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabases)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTables",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTables)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesMysqlColumns",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesMysqlColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "collation", GoGetter: "Collation"},
			_jsii_.MemberProperty{JsiiProperty: "collationInput", GoGetter: "CollationInput"},
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "nullableInput", GoGetter: "NullableInput"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPositionInput", GoGetter: "OrdinalPositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyInput", GoGetter: "PrimaryKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCollation", GoMethod: "ResetCollation"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullable", GoMethod: "ResetNullable"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrdinalPosition", GoMethod: "ResetOrdinalPosition"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryKey", GoMethod: "ResetPrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlColumns", GoGetter: "MysqlColumns"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlColumnsInput", GoGetter: "MysqlColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlColumns", GoMethod: "PutMysqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetMysqlColumns", GoMethod: "ResetMysqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesMysqlTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInput", GoGetter: "DatabaseInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlTables", GoGetter: "MysqlTables"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlTablesInput", GoGetter: "MysqlTablesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlTables", GoMethod: "PutMysqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetMysqlTables", GoMethod: "ResetMysqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsMysqlDatabasesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlDatabases", GoGetter: "MysqlDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlDatabasesInput", GoGetter: "MysqlDatabasesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlDatabases", GoMethod: "PutMysqlDatabases"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigExcludeObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjects)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabases",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabases)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTables",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTables)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesMysqlColumns",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesMysqlColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "collation", GoGetter: "Collation"},
			_jsii_.MemberProperty{JsiiProperty: "collationInput", GoGetter: "CollationInput"},
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "nullableInput", GoGetter: "NullableInput"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPositionInput", GoGetter: "OrdinalPositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyInput", GoGetter: "PrimaryKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCollation", GoMethod: "ResetCollation"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullable", GoMethod: "ResetNullable"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrdinalPosition", GoMethod: "ResetOrdinalPosition"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryKey", GoMethod: "ResetPrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesMysqlColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlColumns", GoGetter: "MysqlColumns"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlColumnsInput", GoGetter: "MysqlColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlColumns", GoMethod: "PutMysqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetMysqlColumns", GoMethod: "ResetMysqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesMysqlTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "database", GoGetter: "Database"},
			_jsii_.MemberProperty{JsiiProperty: "databaseInput", GoGetter: "DatabaseInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlTables", GoGetter: "MysqlTables"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlTablesInput", GoGetter: "MysqlTablesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlTables", GoMethod: "PutMysqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetMysqlTables", GoMethod: "ResetMysqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsMysqlDatabasesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlDatabases", GoGetter: "MysqlDatabases"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlDatabasesInput", GoGetter: "MysqlDatabasesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlDatabases", GoMethod: "PutMysqlDatabases"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigIncludeObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludeObjects", GoGetter: "ExcludeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "excludeObjectsInput", GoGetter: "ExcludeObjectsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includeObjects", GoGetter: "IncludeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "includeObjectsInput", GoGetter: "IncludeObjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentBackfillTasks", GoGetter: "MaxConcurrentBackfillTasks"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentBackfillTasksInput", GoGetter: "MaxConcurrentBackfillTasksInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentCdcTasks", GoGetter: "MaxConcurrentCdcTasks"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentCdcTasksInput", GoGetter: "MaxConcurrentCdcTasksInput"},
			_jsii_.MemberMethod{JsiiMethod: "putExcludeObjects", GoMethod: "PutExcludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "putIncludeObjects", GoMethod: "PutIncludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeObjects", GoMethod: "ResetExcludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeObjects", GoMethod: "ResetIncludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConcurrentBackfillTasks", GoMethod: "ResetMaxConcurrentBackfillTasks"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConcurrentCdcTasks", GoMethod: "ResetMaxConcurrentCdcTasks"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigMysqlSourceConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfig",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjects",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjects)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjectsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigDropLargeObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjects",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjects)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemas",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTables",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTables)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOracleColumns",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOracleColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOracleColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOracleColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOracleColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "encoding", GoGetter: "Encoding"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "precision", GoGetter: "Precision"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scale", GoGetter: "Scale"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "oracleColumns", GoGetter: "OracleColumns"},
			_jsii_.MemberProperty{JsiiProperty: "oracleColumnsInput", GoGetter: "OracleColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleColumns", GoMethod: "PutOracleColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetOracleColumns", GoMethod: "ResetOracleColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOracleTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "oracleTables", GoGetter: "OracleTables"},
			_jsii_.MemberProperty{JsiiProperty: "oracleTablesInput", GoGetter: "OracleTablesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleTables", GoMethod: "PutOracleTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetOracleTables", GoMethod: "ResetOracleTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOracleSchemasOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "oracleSchemas", GoGetter: "OracleSchemas"},
			_jsii_.MemberProperty{JsiiProperty: "oracleSchemasInput", GoGetter: "OracleSchemasInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleSchemas", GoMethod: "PutOracleSchemas"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigExcludeObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjects",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjects)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemas",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTables",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTables)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumns",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "encoding", GoGetter: "Encoding"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "precision", GoGetter: "Precision"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scale", GoGetter: "Scale"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOracleColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "oracleColumns", GoGetter: "OracleColumns"},
			_jsii_.MemberProperty{JsiiProperty: "oracleColumnsInput", GoGetter: "OracleColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleColumns", GoMethod: "PutOracleColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetOracleColumns", GoMethod: "ResetOracleColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOracleTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "oracleTables", GoGetter: "OracleTables"},
			_jsii_.MemberProperty{JsiiProperty: "oracleTablesInput", GoGetter: "OracleTablesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleTables", GoMethod: "PutOracleTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetOracleTables", GoMethod: "ResetOracleTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOracleSchemasOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "oracleSchemas", GoGetter: "OracleSchemas"},
			_jsii_.MemberProperty{JsiiProperty: "oracleSchemasInput", GoGetter: "OracleSchemasInput"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleSchemas", GoMethod: "PutOracleSchemas"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigIncludeObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dropLargeObjects", GoGetter: "DropLargeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "dropLargeObjectsInput", GoGetter: "DropLargeObjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "excludeObjects", GoGetter: "ExcludeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "excludeObjectsInput", GoGetter: "ExcludeObjectsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includeObjects", GoGetter: "IncludeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "includeObjectsInput", GoGetter: "IncludeObjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentBackfillTasks", GoGetter: "MaxConcurrentBackfillTasks"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentBackfillTasksInput", GoGetter: "MaxConcurrentBackfillTasksInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentCdcTasks", GoGetter: "MaxConcurrentCdcTasks"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentCdcTasksInput", GoGetter: "MaxConcurrentCdcTasksInput"},
			_jsii_.MemberMethod{JsiiMethod: "putDropLargeObjects", GoMethod: "PutDropLargeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "putExcludeObjects", GoMethod: "PutExcludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "putIncludeObjects", GoMethod: "PutIncludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "putStreamLargeObjects", GoMethod: "PutStreamLargeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetDropLargeObjects", GoMethod: "ResetDropLargeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeObjects", GoMethod: "ResetExcludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeObjects", GoMethod: "ResetIncludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConcurrentBackfillTasks", GoMethod: "ResetMaxConcurrentBackfillTasks"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConcurrentCdcTasks", GoMethod: "ResetMaxConcurrentCdcTasks"},
			_jsii_.MemberMethod{JsiiMethod: "resetStreamLargeObjects", GoMethod: "ResetStreamLargeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "streamLargeObjects", GoGetter: "StreamLargeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "streamLargeObjectsInput", GoGetter: "StreamLargeObjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjects",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjects)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjectsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOracleSourceConfigStreamLargeObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "mysqlSourceConfig", GoGetter: "MysqlSourceConfig"},
			_jsii_.MemberProperty{JsiiProperty: "mysqlSourceConfigInput", GoGetter: "MysqlSourceConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "oracleSourceConfig", GoGetter: "OracleSourceConfig"},
			_jsii_.MemberProperty{JsiiProperty: "oracleSourceConfigInput", GoGetter: "OracleSourceConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlSourceConfig", GoGetter: "PostgresqlSourceConfig"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlSourceConfigInput", GoGetter: "PostgresqlSourceConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putMysqlSourceConfig", GoMethod: "PutMysqlSourceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putOracleSourceConfig", GoMethod: "PutOracleSourceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlSourceConfig", GoMethod: "PutPostgresqlSourceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSqlServerSourceConfig", GoMethod: "PutSqlServerSourceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMysqlSourceConfig", GoMethod: "ResetMysqlSourceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOracleSourceConfig", GoMethod: "ResetOracleSourceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostgresqlSourceConfig", GoMethod: "ResetPostgresqlSourceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetSqlServerSourceConfig", GoMethod: "ResetSqlServerSourceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "sourceConnectionProfile", GoGetter: "SourceConnectionProfile"},
			_jsii_.MemberProperty{JsiiProperty: "sourceConnectionProfileInput", GoGetter: "SourceConnectionProfileInput"},
			_jsii_.MemberProperty{JsiiProperty: "sqlServerSourceConfig", GoGetter: "SqlServerSourceConfig"},
			_jsii_.MemberProperty{JsiiProperty: "sqlServerSourceConfigInput", GoGetter: "SqlServerSourceConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfig",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjects",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjects)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "postgresqlSchemas", GoGetter: "PostgresqlSchemas"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlSchemasInput", GoGetter: "PostgresqlSchemasInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlSchemas", GoMethod: "PutPostgresqlSchemas"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemas",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "postgresqlTables", GoGetter: "PostgresqlTables"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlTablesInput", GoGetter: "PostgresqlTablesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlTables", GoMethod: "PutPostgresqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostgresqlTables", GoMethod: "ResetPostgresqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTables",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTables)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "postgresqlColumns", GoGetter: "PostgresqlColumns"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlColumnsInput", GoGetter: "PostgresqlColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlColumns", GoMethod: "PutPostgresqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostgresqlColumns", GoMethod: "ResetPostgresqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumns",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "nullableInput", GoGetter: "NullableInput"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPositionInput", GoGetter: "OrdinalPositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "precision", GoGetter: "Precision"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyInput", GoGetter: "PrimaryKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullable", GoMethod: "ResetNullable"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrdinalPosition", GoMethod: "ResetOrdinalPosition"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryKey", GoMethod: "ResetPrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scale", GoGetter: "Scale"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigExcludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjects",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjects)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "postgresqlSchemas", GoGetter: "PostgresqlSchemas"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlSchemasInput", GoGetter: "PostgresqlSchemasInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlSchemas", GoMethod: "PutPostgresqlSchemas"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemas",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "postgresqlTables", GoGetter: "PostgresqlTables"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlTablesInput", GoGetter: "PostgresqlTablesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlTables", GoMethod: "PutPostgresqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostgresqlTables", GoMethod: "ResetPostgresqlTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTables",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTables)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "postgresqlColumns", GoGetter: "PostgresqlColumns"},
			_jsii_.MemberProperty{JsiiProperty: "postgresqlColumnsInput", GoGetter: "PostgresqlColumnsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putPostgresqlColumns", GoMethod: "PutPostgresqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetPostgresqlColumns", GoMethod: "ResetPostgresqlColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumns",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "nullableInput", GoGetter: "NullableInput"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPositionInput", GoGetter: "OrdinalPositionInput"},
			_jsii_.MemberProperty{JsiiProperty: "precision", GoGetter: "Precision"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKeyInput", GoGetter: "PrimaryKeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNullable", GoMethod: "ResetNullable"},
			_jsii_.MemberMethod{JsiiMethod: "resetOrdinalPosition", GoMethod: "ResetOrdinalPosition"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrimaryKey", GoMethod: "ResetPrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scale", GoGetter: "Scale"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigIncludeObjectsPostgresqlSchemasPostgresqlTablesPostgresqlColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludeObjects", GoGetter: "ExcludeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "excludeObjectsInput", GoGetter: "ExcludeObjectsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includeObjects", GoGetter: "IncludeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "includeObjectsInput", GoGetter: "IncludeObjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentBackfillTasks", GoGetter: "MaxConcurrentBackfillTasks"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentBackfillTasksInput", GoGetter: "MaxConcurrentBackfillTasksInput"},
			_jsii_.MemberProperty{JsiiProperty: "publication", GoGetter: "Publication"},
			_jsii_.MemberProperty{JsiiProperty: "publicationInput", GoGetter: "PublicationInput"},
			_jsii_.MemberMethod{JsiiMethod: "putExcludeObjects", GoMethod: "PutExcludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "putIncludeObjects", GoMethod: "PutIncludeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "replicationSlot", GoGetter: "ReplicationSlot"},
			_jsii_.MemberProperty{JsiiProperty: "replicationSlotInput", GoGetter: "ReplicationSlotInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeObjects", GoMethod: "ResetExcludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeObjects", GoMethod: "ResetIncludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConcurrentBackfillTasks", GoMethod: "ResetMaxConcurrentBackfillTasks"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigPostgresqlSourceConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfig",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjects",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjects)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putSchemas", GoMethod: "PutSchemas"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schemas", GoGetter: "Schemas"},
			_jsii_.MemberProperty{JsiiProperty: "schemasInput", GoGetter: "SchemasInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemas",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putTables", GoMethod: "PutTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetTables", GoMethod: "ResetTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "tables", GoGetter: "Tables"},
			_jsii_.MemberProperty{JsiiProperty: "tablesInput", GoGetter: "TablesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTables",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesColumns",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "precision", GoGetter: "Precision"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scale", GoGetter: "Scale"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "columns", GoGetter: "Columns"},
			_jsii_.MemberProperty{JsiiProperty: "columnsInput", GoGetter: "ColumnsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putColumns", GoMethod: "PutColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumns", GoMethod: "ResetColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigExcludeObjectsSchemasTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjects",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjects)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putSchemas", GoMethod: "PutSchemas"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schemas", GoGetter: "Schemas"},
			_jsii_.MemberProperty{JsiiProperty: "schemasInput", GoGetter: "SchemasInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemas",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemas)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putTables", GoMethod: "PutTables"},
			_jsii_.MemberMethod{JsiiMethod: "resetTables", GoMethod: "ResetTables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "schema", GoGetter: "Schema"},
			_jsii_.MemberProperty{JsiiProperty: "schemaInput", GoGetter: "SchemaInput"},
			_jsii_.MemberProperty{JsiiProperty: "tables", GoGetter: "Tables"},
			_jsii_.MemberProperty{JsiiProperty: "tablesInput", GoGetter: "TablesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTables",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTables)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesColumns",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesColumns)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesColumnsList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesColumnsList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesColumnsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesColumnsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesColumnsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "column", GoGetter: "Column"},
			_jsii_.MemberProperty{JsiiProperty: "columnInput", GoGetter: "ColumnInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataType", GoGetter: "DataType"},
			_jsii_.MemberProperty{JsiiProperty: "dataTypeInput", GoGetter: "DataTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "length", GoGetter: "Length"},
			_jsii_.MemberProperty{JsiiProperty: "nullable", GoGetter: "Nullable"},
			_jsii_.MemberProperty{JsiiProperty: "ordinalPosition", GoGetter: "OrdinalPosition"},
			_jsii_.MemberProperty{JsiiProperty: "precision", GoGetter: "Precision"},
			_jsii_.MemberProperty{JsiiProperty: "primaryKey", GoGetter: "PrimaryKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumn", GoMethod: "ResetColumn"},
			_jsii_.MemberMethod{JsiiMethod: "resetDataType", GoMethod: "ResetDataType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scale", GoGetter: "Scale"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesColumnsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesList",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesList)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "columns", GoGetter: "Columns"},
			_jsii_.MemberProperty{JsiiProperty: "columnsInput", GoGetter: "ColumnsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putColumns", GoMethod: "PutColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resetColumns", GoMethod: "ResetColumns"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "table", GoGetter: "Table"},
			_jsii_.MemberProperty{JsiiProperty: "tableInput", GoGetter: "TableInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigIncludeObjectsSchemasTablesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamSourceConfigSqlServerSourceConfigOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamSourceConfigSqlServerSourceConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "excludeObjects", GoGetter: "ExcludeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "excludeObjectsInput", GoGetter: "ExcludeObjectsInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "includeObjects", GoGetter: "IncludeObjects"},
			_jsii_.MemberProperty{JsiiProperty: "includeObjectsInput", GoGetter: "IncludeObjectsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentBackfillTasks", GoGetter: "MaxConcurrentBackfillTasks"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentBackfillTasksInput", GoGetter: "MaxConcurrentBackfillTasksInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentCdcTasks", GoGetter: "MaxConcurrentCdcTasks"},
			_jsii_.MemberProperty{JsiiProperty: "maxConcurrentCdcTasksInput", GoGetter: "MaxConcurrentCdcTasksInput"},
			_jsii_.MemberMethod{JsiiMethod: "putExcludeObjects", GoMethod: "PutExcludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "putIncludeObjects", GoMethod: "PutIncludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetExcludeObjects", GoMethod: "ResetExcludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetIncludeObjects", GoMethod: "ResetIncludeObjects"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConcurrentBackfillTasks", GoMethod: "ResetMaxConcurrentBackfillTasks"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxConcurrentCdcTasks", GoMethod: "ResetMaxConcurrentCdcTasks"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDatastreamStreamSourceConfigSqlServerSourceConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamTimeouts",
		reflect.TypeOf((*GoogleDatastreamStreamTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"@cdktf/provider-google-beta.googleDatastreamStream.GoogleDatastreamStreamTimeoutsOutputReference",
		reflect.TypeOf((*GoogleDatastreamStreamTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDatastreamStreamTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}
