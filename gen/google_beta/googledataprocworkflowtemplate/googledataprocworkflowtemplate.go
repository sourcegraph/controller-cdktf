package googledataprocworkflowtemplate

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplate",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplate)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "createTime", GoGetter: "CreateTime"},
			_jsii_.MemberProperty{JsiiProperty: "dagTimeout", GoGetter: "DagTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "dagTimeoutInput", GoGetter: "DagTimeoutInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "jobs", GoGetter: "Jobs"},
			_jsii_.MemberProperty{JsiiProperty: "jobsInput", GoGetter: "JobsInput"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "labelsInput", GoGetter: "LabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "location", GoGetter: "Location"},
			_jsii_.MemberProperty{JsiiProperty: "locationInput", GoGetter: "LocationInput"},
			_jsii_.MemberProperty{JsiiProperty: "name", GoGetter: "Name"},
			_jsii_.MemberProperty{JsiiProperty: "nameInput", GoGetter: "NameInput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "parameters", GoGetter: "Parameters"},
			_jsii_.MemberProperty{JsiiProperty: "parametersInput", GoGetter: "ParametersInput"},
			_jsii_.MemberProperty{JsiiProperty: "placement", GoGetter: "Placement"},
			_jsii_.MemberProperty{JsiiProperty: "placementInput", GoGetter: "PlacementInput"},
			_jsii_.MemberProperty{JsiiProperty: "project", GoGetter: "Project"},
			_jsii_.MemberProperty{JsiiProperty: "projectInput", GoGetter: "ProjectInput"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberMethod{JsiiMethod: "putJobs", GoMethod: "PutJobs"},
			_jsii_.MemberMethod{JsiiMethod: "putParameters", GoMethod: "PutParameters"},
			_jsii_.MemberMethod{JsiiMethod: "putPlacement", GoMethod: "PutPlacement"},
			_jsii_.MemberMethod{JsiiMethod: "putTimeouts", GoMethod: "PutTimeouts"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetDagTimeout", GoMethod: "ResetDagTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resetId", GoMethod: "ResetId"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabels", GoMethod: "ResetLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetParameters", GoMethod: "ResetParameters"},
			_jsii_.MemberMethod{JsiiMethod: "resetProject", GoMethod: "ResetProject"},
			_jsii_.MemberMethod{JsiiMethod: "resetTimeouts", GoMethod: "ResetTimeouts"},
			_jsii_.MemberMethod{JsiiMethod: "resetVersion", GoMethod: "ResetVersion"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberProperty{JsiiProperty: "timeouts", GoGetter: "Timeouts"},
			_jsii_.MemberProperty{JsiiProperty: "timeoutsInput", GoGetter: "TimeoutsInput"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "updateTime", GoGetter: "UpdateTime"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
			_jsii_.MemberProperty{JsiiProperty: "versionInput", GoGetter: "VersionInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplate{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobs",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobs)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsHadoopJob",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsHadoopJob)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevels", GoGetter: "DriverLogLevels"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevelsInput", GoGetter: "DriverLogLevelsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDriverLogLevels", GoMethod: "ResetDriverLogLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobLoggingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "archiveUris", GoGetter: "ArchiveUris"},
			_jsii_.MemberProperty{JsiiProperty: "archiveUrisInput", GoGetter: "ArchiveUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "args", GoGetter: "Args"},
			_jsii_.MemberProperty{JsiiProperty: "argsInput", GoGetter: "ArgsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileUris", GoGetter: "FileUris"},
			_jsii_.MemberProperty{JsiiProperty: "fileUrisInput", GoGetter: "FileUrisInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "jarFileUris", GoGetter: "JarFileUris"},
			_jsii_.MemberProperty{JsiiProperty: "jarFileUrisInput", GoGetter: "JarFileUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigInput", GoGetter: "LoggingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "mainClass", GoGetter: "MainClass"},
			_jsii_.MemberProperty{JsiiProperty: "mainClassInput", GoGetter: "MainClassInput"},
			_jsii_.MemberProperty{JsiiProperty: "mainJarFileUri", GoGetter: "MainJarFileUri"},
			_jsii_.MemberProperty{JsiiProperty: "mainJarFileUriInput", GoGetter: "MainJarFileUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfig", GoMethod: "PutLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetArchiveUris", GoMethod: "ResetArchiveUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetArgs", GoMethod: "ResetArgs"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileUris", GoMethod: "ResetFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetJarFileUris", GoMethod: "ResetJarFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfig", GoMethod: "ResetLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMainClass", GoMethod: "ResetMainClass"},
			_jsii_.MemberMethod{JsiiMethod: "resetMainJarFileUri", GoMethod: "ResetMainJarFileUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsHadoopJobOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsHiveJob",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsHiveJob)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsHiveJobOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsHiveJobOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "continueOnFailure", GoGetter: "ContinueOnFailure"},
			_jsii_.MemberProperty{JsiiProperty: "continueOnFailureInput", GoGetter: "ContinueOnFailureInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "jarFileUris", GoGetter: "JarFileUris"},
			_jsii_.MemberProperty{JsiiProperty: "jarFileUrisInput", GoGetter: "JarFileUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryList", GoMethod: "PutQueryList"},
			_jsii_.MemberProperty{JsiiProperty: "queryFileUri", GoGetter: "QueryFileUri"},
			_jsii_.MemberProperty{JsiiProperty: "queryFileUriInput", GoGetter: "QueryFileUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryList", GoGetter: "QueryList"},
			_jsii_.MemberProperty{JsiiProperty: "queryListInput", GoGetter: "QueryListInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetContinueOnFailure", GoMethod: "ResetContinueOnFailure"},
			_jsii_.MemberMethod{JsiiMethod: "resetJarFileUris", GoMethod: "ResetJarFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryFileUri", GoMethod: "ResetQueryFileUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryList", GoMethod: "ResetQueryList"},
			_jsii_.MemberMethod{JsiiMethod: "resetScriptVariables", GoMethod: "ResetScriptVariables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scriptVariables", GoGetter: "ScriptVariables"},
			_jsii_.MemberProperty{JsiiProperty: "scriptVariablesInput", GoGetter: "ScriptVariablesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsHiveJobOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsHiveJobQueryList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsHiveJobQueryList)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsHiveJobQueryListOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsHiveJobQueryListOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "queries", GoGetter: "Queries"},
			_jsii_.MemberProperty{JsiiProperty: "queriesInput", GoGetter: "QueriesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsHiveJobQueryListOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "hadoopJob", GoGetter: "HadoopJob"},
			_jsii_.MemberProperty{JsiiProperty: "hadoopJobInput", GoGetter: "HadoopJobInput"},
			_jsii_.MemberProperty{JsiiProperty: "hiveJob", GoGetter: "HiveJob"},
			_jsii_.MemberProperty{JsiiProperty: "hiveJobInput", GoGetter: "HiveJobInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "labelsInput", GoGetter: "LabelsInput"},
			_jsii_.MemberProperty{JsiiProperty: "pigJob", GoGetter: "PigJob"},
			_jsii_.MemberProperty{JsiiProperty: "pigJobInput", GoGetter: "PigJobInput"},
			_jsii_.MemberProperty{JsiiProperty: "prerequisiteStepIds", GoGetter: "PrerequisiteStepIds"},
			_jsii_.MemberProperty{JsiiProperty: "prerequisiteStepIdsInput", GoGetter: "PrerequisiteStepIdsInput"},
			_jsii_.MemberProperty{JsiiProperty: "prestoJob", GoGetter: "PrestoJob"},
			_jsii_.MemberProperty{JsiiProperty: "prestoJobInput", GoGetter: "PrestoJobInput"},
			_jsii_.MemberMethod{JsiiMethod: "putHadoopJob", GoMethod: "PutHadoopJob"},
			_jsii_.MemberMethod{JsiiMethod: "putHiveJob", GoMethod: "PutHiveJob"},
			_jsii_.MemberMethod{JsiiMethod: "putPigJob", GoMethod: "PutPigJob"},
			_jsii_.MemberMethod{JsiiMethod: "putPrestoJob", GoMethod: "PutPrestoJob"},
			_jsii_.MemberMethod{JsiiMethod: "putPysparkJob", GoMethod: "PutPysparkJob"},
			_jsii_.MemberMethod{JsiiMethod: "putScheduling", GoMethod: "PutScheduling"},
			_jsii_.MemberMethod{JsiiMethod: "putSparkJob", GoMethod: "PutSparkJob"},
			_jsii_.MemberMethod{JsiiMethod: "putSparkRJob", GoMethod: "PutSparkRJob"},
			_jsii_.MemberMethod{JsiiMethod: "putSparkSqlJob", GoMethod: "PutSparkSqlJob"},
			_jsii_.MemberProperty{JsiiProperty: "pysparkJob", GoGetter: "PysparkJob"},
			_jsii_.MemberProperty{JsiiProperty: "pysparkJobInput", GoGetter: "PysparkJobInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetHadoopJob", GoMethod: "ResetHadoopJob"},
			_jsii_.MemberMethod{JsiiMethod: "resetHiveJob", GoMethod: "ResetHiveJob"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabels", GoMethod: "ResetLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resetPigJob", GoMethod: "ResetPigJob"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrerequisiteStepIds", GoMethod: "ResetPrerequisiteStepIds"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrestoJob", GoMethod: "ResetPrestoJob"},
			_jsii_.MemberMethod{JsiiMethod: "resetPysparkJob", GoMethod: "ResetPysparkJob"},
			_jsii_.MemberMethod{JsiiMethod: "resetScheduling", GoMethod: "ResetScheduling"},
			_jsii_.MemberMethod{JsiiMethod: "resetSparkJob", GoMethod: "ResetSparkJob"},
			_jsii_.MemberMethod{JsiiMethod: "resetSparkRJob", GoMethod: "ResetSparkRJob"},
			_jsii_.MemberMethod{JsiiMethod: "resetSparkSqlJob", GoMethod: "ResetSparkSqlJob"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scheduling", GoGetter: "Scheduling"},
			_jsii_.MemberProperty{JsiiProperty: "schedulingInput", GoGetter: "SchedulingInput"},
			_jsii_.MemberProperty{JsiiProperty: "sparkJob", GoGetter: "SparkJob"},
			_jsii_.MemberProperty{JsiiProperty: "sparkJobInput", GoGetter: "SparkJobInput"},
			_jsii_.MemberProperty{JsiiProperty: "sparkRJob", GoGetter: "SparkRJob"},
			_jsii_.MemberProperty{JsiiProperty: "sparkRJobInput", GoGetter: "SparkRJobInput"},
			_jsii_.MemberProperty{JsiiProperty: "sparkSqlJob", GoGetter: "SparkSqlJob"},
			_jsii_.MemberProperty{JsiiProperty: "sparkSqlJobInput", GoGetter: "SparkSqlJobInput"},
			_jsii_.MemberProperty{JsiiProperty: "stepId", GoGetter: "StepId"},
			_jsii_.MemberProperty{JsiiProperty: "stepIdInput", GoGetter: "StepIdInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPigJob",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPigJob)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPigJobLoggingConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPigJobLoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPigJobLoggingConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPigJobLoggingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevels", GoGetter: "DriverLogLevels"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevelsInput", GoGetter: "DriverLogLevelsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDriverLogLevels", GoMethod: "ResetDriverLogLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsPigJobLoggingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPigJobOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPigJobOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "continueOnFailure", GoGetter: "ContinueOnFailure"},
			_jsii_.MemberProperty{JsiiProperty: "continueOnFailureInput", GoGetter: "ContinueOnFailureInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "jarFileUris", GoGetter: "JarFileUris"},
			_jsii_.MemberProperty{JsiiProperty: "jarFileUrisInput", GoGetter: "JarFileUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigInput", GoGetter: "LoggingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfig", GoMethod: "PutLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryList", GoMethod: "PutQueryList"},
			_jsii_.MemberProperty{JsiiProperty: "queryFileUri", GoGetter: "QueryFileUri"},
			_jsii_.MemberProperty{JsiiProperty: "queryFileUriInput", GoGetter: "QueryFileUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryList", GoGetter: "QueryList"},
			_jsii_.MemberProperty{JsiiProperty: "queryListInput", GoGetter: "QueryListInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetContinueOnFailure", GoMethod: "ResetContinueOnFailure"},
			_jsii_.MemberMethod{JsiiMethod: "resetJarFileUris", GoMethod: "ResetJarFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfig", GoMethod: "ResetLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryFileUri", GoMethod: "ResetQueryFileUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryList", GoMethod: "ResetQueryList"},
			_jsii_.MemberMethod{JsiiMethod: "resetScriptVariables", GoMethod: "ResetScriptVariables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scriptVariables", GoGetter: "ScriptVariables"},
			_jsii_.MemberProperty{JsiiProperty: "scriptVariablesInput", GoGetter: "ScriptVariablesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsPigJobOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPigJobQueryList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPigJobQueryList)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPigJobQueryListOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPigJobQueryListOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "queries", GoGetter: "Queries"},
			_jsii_.MemberProperty{JsiiProperty: "queriesInput", GoGetter: "QueriesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsPigJobQueryListOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPrestoJob",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPrestoJob)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPrestoJobLoggingConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPrestoJobLoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPrestoJobLoggingConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPrestoJobLoggingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevels", GoGetter: "DriverLogLevels"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevelsInput", GoGetter: "DriverLogLevelsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDriverLogLevels", GoMethod: "ResetDriverLogLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsPrestoJobLoggingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPrestoJobOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPrestoJobOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clientTags", GoGetter: "ClientTags"},
			_jsii_.MemberProperty{JsiiProperty: "clientTagsInput", GoGetter: "ClientTagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "continueOnFailure", GoGetter: "ContinueOnFailure"},
			_jsii_.MemberProperty{JsiiProperty: "continueOnFailureInput", GoGetter: "ContinueOnFailureInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigInput", GoGetter: "LoggingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "outputFormat", GoGetter: "OutputFormat"},
			_jsii_.MemberProperty{JsiiProperty: "outputFormatInput", GoGetter: "OutputFormatInput"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfig", GoMethod: "PutLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryList", GoMethod: "PutQueryList"},
			_jsii_.MemberProperty{JsiiProperty: "queryFileUri", GoGetter: "QueryFileUri"},
			_jsii_.MemberProperty{JsiiProperty: "queryFileUriInput", GoGetter: "QueryFileUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryList", GoGetter: "QueryList"},
			_jsii_.MemberProperty{JsiiProperty: "queryListInput", GoGetter: "QueryListInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetClientTags", GoMethod: "ResetClientTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetContinueOnFailure", GoMethod: "ResetContinueOnFailure"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfig", GoMethod: "ResetLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetOutputFormat", GoMethod: "ResetOutputFormat"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryFileUri", GoMethod: "ResetQueryFileUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryList", GoMethod: "ResetQueryList"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsPrestoJobOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPrestoJobQueryList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPrestoJobQueryList)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPrestoJobQueryListOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPrestoJobQueryListOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "queries", GoGetter: "Queries"},
			_jsii_.MemberProperty{JsiiProperty: "queriesInput", GoGetter: "QueriesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsPrestoJobQueryListOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPysparkJob",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPysparkJob)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPysparkJobLoggingConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPysparkJobLoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPysparkJobLoggingConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPysparkJobLoggingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevels", GoGetter: "DriverLogLevels"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevelsInput", GoGetter: "DriverLogLevelsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDriverLogLevels", GoMethod: "ResetDriverLogLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsPysparkJobLoggingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsPysparkJobOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsPysparkJobOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "archiveUris", GoGetter: "ArchiveUris"},
			_jsii_.MemberProperty{JsiiProperty: "archiveUrisInput", GoGetter: "ArchiveUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "args", GoGetter: "Args"},
			_jsii_.MemberProperty{JsiiProperty: "argsInput", GoGetter: "ArgsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileUris", GoGetter: "FileUris"},
			_jsii_.MemberProperty{JsiiProperty: "fileUrisInput", GoGetter: "FileUrisInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "jarFileUris", GoGetter: "JarFileUris"},
			_jsii_.MemberProperty{JsiiProperty: "jarFileUrisInput", GoGetter: "JarFileUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigInput", GoGetter: "LoggingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "mainPythonFileUri", GoGetter: "MainPythonFileUri"},
			_jsii_.MemberProperty{JsiiProperty: "mainPythonFileUriInput", GoGetter: "MainPythonFileUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfig", GoMethod: "PutLoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "pythonFileUris", GoGetter: "PythonFileUris"},
			_jsii_.MemberProperty{JsiiProperty: "pythonFileUrisInput", GoGetter: "PythonFileUrisInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetArchiveUris", GoMethod: "ResetArchiveUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetArgs", GoMethod: "ResetArgs"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileUris", GoMethod: "ResetFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetJarFileUris", GoMethod: "ResetJarFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfig", GoMethod: "ResetLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetPythonFileUris", GoMethod: "ResetPythonFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsPysparkJobOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsScheduling",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsScheduling)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSchedulingOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSchedulingOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "maxFailuresPerHour", GoGetter: "MaxFailuresPerHour"},
			_jsii_.MemberProperty{JsiiProperty: "maxFailuresPerHourInput", GoGetter: "MaxFailuresPerHourInput"},
			_jsii_.MemberProperty{JsiiProperty: "maxFailuresTotal", GoGetter: "MaxFailuresTotal"},
			_jsii_.MemberProperty{JsiiProperty: "maxFailuresTotalInput", GoGetter: "MaxFailuresTotalInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxFailuresPerHour", GoMethod: "ResetMaxFailuresPerHour"},
			_jsii_.MemberMethod{JsiiMethod: "resetMaxFailuresTotal", GoMethod: "ResetMaxFailuresTotal"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsSchedulingOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkJob",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkJob)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkJobLoggingConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkJobLoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkJobLoggingConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkJobLoggingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevels", GoGetter: "DriverLogLevels"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevelsInput", GoGetter: "DriverLogLevelsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDriverLogLevels", GoMethod: "ResetDriverLogLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkJobLoggingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkJobOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkJobOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "archiveUris", GoGetter: "ArchiveUris"},
			_jsii_.MemberProperty{JsiiProperty: "archiveUrisInput", GoGetter: "ArchiveUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "args", GoGetter: "Args"},
			_jsii_.MemberProperty{JsiiProperty: "argsInput", GoGetter: "ArgsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileUris", GoGetter: "FileUris"},
			_jsii_.MemberProperty{JsiiProperty: "fileUrisInput", GoGetter: "FileUrisInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "jarFileUris", GoGetter: "JarFileUris"},
			_jsii_.MemberProperty{JsiiProperty: "jarFileUrisInput", GoGetter: "JarFileUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigInput", GoGetter: "LoggingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "mainClass", GoGetter: "MainClass"},
			_jsii_.MemberProperty{JsiiProperty: "mainClassInput", GoGetter: "MainClassInput"},
			_jsii_.MemberProperty{JsiiProperty: "mainJarFileUri", GoGetter: "MainJarFileUri"},
			_jsii_.MemberProperty{JsiiProperty: "mainJarFileUriInput", GoGetter: "MainJarFileUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfig", GoMethod: "PutLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetArchiveUris", GoMethod: "ResetArchiveUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetArgs", GoMethod: "ResetArgs"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileUris", GoMethod: "ResetFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetJarFileUris", GoMethod: "ResetJarFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfig", GoMethod: "ResetLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMainClass", GoMethod: "ResetMainClass"},
			_jsii_.MemberMethod{JsiiMethod: "resetMainJarFileUri", GoMethod: "ResetMainJarFileUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkJobOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkRJob",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkRJob)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkRJobLoggingConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkRJobLoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkRJobLoggingConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkRJobLoggingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevels", GoGetter: "DriverLogLevels"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevelsInput", GoGetter: "DriverLogLevelsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDriverLogLevels", GoMethod: "ResetDriverLogLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkRJobLoggingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkRJobOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkRJobOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "archiveUris", GoGetter: "ArchiveUris"},
			_jsii_.MemberProperty{JsiiProperty: "archiveUrisInput", GoGetter: "ArchiveUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "args", GoGetter: "Args"},
			_jsii_.MemberProperty{JsiiProperty: "argsInput", GoGetter: "ArgsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fileUris", GoGetter: "FileUris"},
			_jsii_.MemberProperty{JsiiProperty: "fileUrisInput", GoGetter: "FileUrisInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigInput", GoGetter: "LoggingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "mainRFileUri", GoGetter: "MainRFileUri"},
			_jsii_.MemberProperty{JsiiProperty: "mainRFileUriInput", GoGetter: "MainRFileUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfig", GoMethod: "PutLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetArchiveUris", GoMethod: "ResetArchiveUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetArgs", GoMethod: "ResetArgs"},
			_jsii_.MemberMethod{JsiiMethod: "resetFileUris", GoMethod: "ResetFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfig", GoMethod: "ResetLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkRJobOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkSqlJob",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkSqlJob)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevels", GoGetter: "DriverLogLevels"},
			_jsii_.MemberProperty{JsiiProperty: "driverLogLevelsInput", GoGetter: "DriverLogLevelsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetDriverLogLevels", GoMethod: "ResetDriverLogLevels"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobLoggingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "jarFileUris", GoGetter: "JarFileUris"},
			_jsii_.MemberProperty{JsiiProperty: "jarFileUrisInput", GoGetter: "JarFileUrisInput"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfig", GoGetter: "LoggingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "loggingConfigInput", GoGetter: "LoggingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "putLoggingConfig", GoMethod: "PutLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putQueryList", GoMethod: "PutQueryList"},
			_jsii_.MemberProperty{JsiiProperty: "queryFileUri", GoGetter: "QueryFileUri"},
			_jsii_.MemberProperty{JsiiProperty: "queryFileUriInput", GoGetter: "QueryFileUriInput"},
			_jsii_.MemberProperty{JsiiProperty: "queryList", GoGetter: "QueryList"},
			_jsii_.MemberProperty{JsiiProperty: "queryListInput", GoGetter: "QueryListInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetJarFileUris", GoMethod: "ResetJarFileUris"},
			_jsii_.MemberMethod{JsiiMethod: "resetLoggingConfig", GoMethod: "ResetLoggingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryFileUri", GoMethod: "ResetQueryFileUri"},
			_jsii_.MemberMethod{JsiiMethod: "resetQueryList", GoMethod: "ResetQueryList"},
			_jsii_.MemberMethod{JsiiMethod: "resetScriptVariables", GoMethod: "ResetScriptVariables"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "scriptVariables", GoGetter: "ScriptVariables"},
			_jsii_.MemberProperty{JsiiProperty: "scriptVariablesInput", GoGetter: "ScriptVariablesInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryList)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryListOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryListOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "queries", GoGetter: "Queries"},
			_jsii_.MemberProperty{JsiiProperty: "queriesInput", GoGetter: "QueriesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateJobsSparkSqlJobQueryListOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateParameters",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateParameters)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateParametersList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateParametersList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_GoogleDataprocWorkflowTemplateParametersList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateParametersOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateParametersOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "description", GoGetter: "Description"},
			_jsii_.MemberProperty{JsiiProperty: "descriptionInput", GoGetter: "DescriptionInput"},
			_jsii_.MemberProperty{JsiiProperty: "fields", GoGetter: "Fields"},
			_jsii_.MemberProperty{JsiiProperty: "fieldsInput", GoGetter: "FieldsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "putValidation", GoMethod: "PutValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resetDescription", GoMethod: "ResetDescription"},
			_jsii_.MemberMethod{JsiiMethod: "resetValidation", GoMethod: "ResetValidation"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "validation", GoGetter: "Validation"},
			_jsii_.MemberProperty{JsiiProperty: "validationInput", GoGetter: "ValidationInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateParametersOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateParametersValidation",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateParametersValidation)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateParametersValidationOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateParametersValidationOutputReference)(nil)).Elem(),
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
			_jsii_.MemberMethod{JsiiMethod: "putRegex", GoMethod: "PutRegex"},
			_jsii_.MemberMethod{JsiiMethod: "putValues", GoMethod: "PutValues"},
			_jsii_.MemberProperty{JsiiProperty: "regex", GoGetter: "Regex"},
			_jsii_.MemberProperty{JsiiProperty: "regexInput", GoGetter: "RegexInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetRegex", GoMethod: "ResetRegex"},
			_jsii_.MemberMethod{JsiiMethod: "resetValues", GoMethod: "ResetValues"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
			_jsii_.MemberProperty{JsiiProperty: "valuesInput", GoGetter: "ValuesInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateParametersValidationOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateParametersValidationRegex",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateParametersValidationRegex)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateParametersValidationRegexOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateParametersValidationRegexOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "regexes", GoGetter: "Regexes"},
			_jsii_.MemberProperty{JsiiProperty: "regexesInput", GoGetter: "RegexesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateParametersValidationRegexOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateParametersValidationValues",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateParametersValidationValues)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateParametersValidationValuesOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateParametersValidationValuesOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
			_jsii_.MemberProperty{JsiiProperty: "valuesInput", GoGetter: "ValuesInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplateParametersValidationValuesOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacement",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacement)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementClusterSelector",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementClusterSelector)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementClusterSelectorOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementClusterSelectorOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clusterLabels", GoGetter: "ClusterLabels"},
			_jsii_.MemberProperty{JsiiProperty: "clusterLabelsInput", GoGetter: "ClusterLabelsInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetZone", GoMethod: "ResetZone"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "zone", GoGetter: "Zone"},
			_jsii_.MemberProperty{JsiiProperty: "zoneInput", GoGetter: "ZoneInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementClusterSelectorOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedCluster",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedCluster)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "policy", GoGetter: "Policy"},
			_jsii_.MemberProperty{JsiiProperty: "policyInput", GoGetter: "PolicyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetPolicy", GoMethod: "ResetPolicy"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigAutoscalingConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "gcePdKmsKeyName", GoGetter: "GcePdKmsKeyName"},
			_jsii_.MemberProperty{JsiiProperty: "gcePdKmsKeyNameInput", GoGetter: "GcePdKmsKeyNameInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetGcePdKmsKeyName", GoMethod: "ResetGcePdKmsKeyName"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEncryptionConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableHttpPortAccess", GoGetter: "EnableHttpPortAccess"},
			_jsii_.MemberProperty{JsiiProperty: "enableHttpPortAccessInput", GoGetter: "EnableHttpPortAccessInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "httpPorts", GoGetter: "HttpPorts"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableHttpPortAccess", GoMethod: "ResetEnableHttpPortAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "nodeGroup", GoGetter: "NodeGroup"},
			_jsii_.MemberProperty{JsiiProperty: "nodeGroupInput", GoGetter: "NodeGroupInput"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigNodeGroupAffinityOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "internalIpOnly", GoGetter: "InternalIpOnly"},
			_jsii_.MemberProperty{JsiiProperty: "internalIpOnlyInput", GoGetter: "InternalIpOnlyInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "metadata", GoGetter: "Metadata"},
			_jsii_.MemberProperty{JsiiProperty: "metadataInput", GoGetter: "MetadataInput"},
			_jsii_.MemberProperty{JsiiProperty: "network", GoGetter: "Network"},
			_jsii_.MemberProperty{JsiiProperty: "networkInput", GoGetter: "NetworkInput"},
			_jsii_.MemberProperty{JsiiProperty: "nodeGroupAffinity", GoGetter: "NodeGroupAffinity"},
			_jsii_.MemberProperty{JsiiProperty: "nodeGroupAffinityInput", GoGetter: "NodeGroupAffinityInput"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpv6GoogleAccess", GoGetter: "PrivateIpv6GoogleAccess"},
			_jsii_.MemberProperty{JsiiProperty: "privateIpv6GoogleAccessInput", GoGetter: "PrivateIpv6GoogleAccessInput"},
			_jsii_.MemberMethod{JsiiMethod: "putNodeGroupAffinity", GoMethod: "PutNodeGroupAffinity"},
			_jsii_.MemberMethod{JsiiMethod: "putReservationAffinity", GoMethod: "PutReservationAffinity"},
			_jsii_.MemberMethod{JsiiMethod: "putShieldedInstanceConfig", GoMethod: "PutShieldedInstanceConfig"},
			_jsii_.MemberProperty{JsiiProperty: "reservationAffinity", GoGetter: "ReservationAffinity"},
			_jsii_.MemberProperty{JsiiProperty: "reservationAffinityInput", GoGetter: "ReservationAffinityInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetInternalIpOnly", GoMethod: "ResetInternalIpOnly"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetadata", GoMethod: "ResetMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "resetNetwork", GoMethod: "ResetNetwork"},
			_jsii_.MemberMethod{JsiiMethod: "resetNodeGroupAffinity", GoMethod: "ResetNodeGroupAffinity"},
			_jsii_.MemberMethod{JsiiMethod: "resetPrivateIpv6GoogleAccess", GoMethod: "ResetPrivateIpv6GoogleAccess"},
			_jsii_.MemberMethod{JsiiMethod: "resetReservationAffinity", GoMethod: "ResetReservationAffinity"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAccount", GoMethod: "ResetServiceAccount"},
			_jsii_.MemberMethod{JsiiMethod: "resetServiceAccountScopes", GoMethod: "ResetServiceAccountScopes"},
			_jsii_.MemberMethod{JsiiMethod: "resetShieldedInstanceConfig", GoMethod: "ResetShieldedInstanceConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetSubnetwork", GoMethod: "ResetSubnetwork"},
			_jsii_.MemberMethod{JsiiMethod: "resetTags", GoMethod: "ResetTags"},
			_jsii_.MemberMethod{JsiiMethod: "resetZone", GoMethod: "ResetZone"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccount", GoGetter: "ServiceAccount"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountInput", GoGetter: "ServiceAccountInput"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountScopes", GoGetter: "ServiceAccountScopes"},
			_jsii_.MemberProperty{JsiiProperty: "serviceAccountScopesInput", GoGetter: "ServiceAccountScopesInput"},
			_jsii_.MemberProperty{JsiiProperty: "shieldedInstanceConfig", GoGetter: "ShieldedInstanceConfig"},
			_jsii_.MemberProperty{JsiiProperty: "shieldedInstanceConfigInput", GoGetter: "ShieldedInstanceConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "subnetwork", GoGetter: "Subnetwork"},
			_jsii_.MemberProperty{JsiiProperty: "subnetworkInput", GoGetter: "SubnetworkInput"},
			_jsii_.MemberProperty{JsiiProperty: "tags", GoGetter: "Tags"},
			_jsii_.MemberProperty{JsiiProperty: "tagsInput", GoGetter: "TagsInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "zone", GoGetter: "Zone"},
			_jsii_.MemberProperty{JsiiProperty: "zoneInput", GoGetter: "ZoneInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinity)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "consumeReservationType", GoGetter: "ConsumeReservationType"},
			_jsii_.MemberProperty{JsiiProperty: "consumeReservationTypeInput", GoGetter: "ConsumeReservationTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "key", GoGetter: "Key"},
			_jsii_.MemberProperty{JsiiProperty: "keyInput", GoGetter: "KeyInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetConsumeReservationType", GoMethod: "ResetConsumeReservationType"},
			_jsii_.MemberMethod{JsiiMethod: "resetKey", GoMethod: "ResetKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetValues", GoMethod: "ResetValues"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "values", GoGetter: "Values"},
			_jsii_.MemberProperty{JsiiProperty: "valuesInput", GoGetter: "ValuesInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigReservationAffinityOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "enableIntegrityMonitoring", GoGetter: "EnableIntegrityMonitoring"},
			_jsii_.MemberProperty{JsiiProperty: "enableIntegrityMonitoringInput", GoGetter: "EnableIntegrityMonitoringInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableSecureBoot", GoGetter: "EnableSecureBoot"},
			_jsii_.MemberProperty{JsiiProperty: "enableSecureBootInput", GoGetter: "EnableSecureBootInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableVtpm", GoGetter: "EnableVtpm"},
			_jsii_.MemberProperty{JsiiProperty: "enableVtpmInput", GoGetter: "EnableVtpmInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetEnableIntegrityMonitoring", GoMethod: "ResetEnableIntegrityMonitoring"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableSecureBoot", GoMethod: "ResetEnableSecureBoot"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableVtpm", GoMethod: "ResetEnableVtpm"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGceClusterConfigShieldedInstanceConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTarget)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clusterNamespace", GoGetter: "ClusterNamespace"},
			_jsii_.MemberProperty{JsiiProperty: "clusterNamespaceInput", GoGetter: "ClusterNamespaceInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetClusterNamespace", GoMethod: "ResetClusterNamespace"},
			_jsii_.MemberMethod{JsiiMethod: "resetTargetGkeCluster", GoMethod: "ResetTargetGkeCluster"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "targetGkeCluster", GoGetter: "TargetGkeCluster"},
			_jsii_.MemberProperty{JsiiProperty: "targetGkeClusterInput", GoGetter: "TargetGkeClusterInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigNamespacedGkeDeploymentTargetOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "namespacedGkeDeploymentTarget", GoGetter: "NamespacedGkeDeploymentTarget"},
			_jsii_.MemberProperty{JsiiProperty: "namespacedGkeDeploymentTargetInput", GoGetter: "NamespacedGkeDeploymentTargetInput"},
			_jsii_.MemberMethod{JsiiMethod: "putNamespacedGkeDeploymentTarget", GoMethod: "PutNamespacedGkeDeploymentTarget"},
			_jsii_.MemberMethod{JsiiMethod: "resetNamespacedGkeDeploymentTarget", GoMethod: "ResetNamespacedGkeDeploymentTarget"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigGkeClusterConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActions",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActions)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "executableFile", GoGetter: "ExecutableFile"},
			_jsii_.MemberProperty{JsiiProperty: "executableFileInput", GoGetter: "ExecutableFileInput"},
			_jsii_.MemberProperty{JsiiProperty: "executionTimeout", GoGetter: "ExecutionTimeout"},
			_jsii_.MemberProperty{JsiiProperty: "executionTimeoutInput", GoGetter: "ExecutionTimeoutInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetExecutableFile", GoMethod: "ResetExecutableFile"},
			_jsii_.MemberMethod{JsiiMethod: "resetExecutionTimeout", GoMethod: "ResetExecutionTimeout"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigInitializationActionsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoDeleteTime", GoGetter: "AutoDeleteTime"},
			_jsii_.MemberProperty{JsiiProperty: "autoDeleteTimeInput", GoGetter: "AutoDeleteTimeInput"},
			_jsii_.MemberProperty{JsiiProperty: "autoDeleteTtl", GoGetter: "AutoDeleteTtl"},
			_jsii_.MemberProperty{JsiiProperty: "autoDeleteTtlInput", GoGetter: "AutoDeleteTtlInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "idleDeleteTtl", GoGetter: "IdleDeleteTtl"},
			_jsii_.MemberProperty{JsiiProperty: "idleDeleteTtlInput", GoGetter: "IdleDeleteTtlInput"},
			_jsii_.MemberProperty{JsiiProperty: "idleStartTime", GoGetter: "IdleStartTime"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoDeleteTime", GoMethod: "ResetAutoDeleteTime"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoDeleteTtl", GoMethod: "ResetAutoDeleteTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resetIdleDeleteTtl", GoMethod: "ResetIdleDeleteTtl"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigLifecycleConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAccelerators)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceleratorCount", GoGetter: "AcceleratorCount"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorCountInput", GoGetter: "AcceleratorCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorType", GoGetter: "AcceleratorType"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorTypeInput", GoGetter: "AcceleratorTypeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorCount", GoMethod: "ResetAcceleratorCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorType", GoMethod: "ResetAcceleratorType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigAcceleratorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bootDiskSizeGb", GoGetter: "BootDiskSizeGb"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiskSizeGbInput", GoGetter: "BootDiskSizeGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiskType", GoGetter: "BootDiskType"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiskTypeInput", GoGetter: "BootDiskTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "numLocalSsds", GoGetter: "NumLocalSsds"},
			_jsii_.MemberProperty{JsiiProperty: "numLocalSsdsInput", GoGetter: "NumLocalSsdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootDiskSizeGb", GoMethod: "ResetBootDiskSizeGb"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootDiskType", GoMethod: "ResetBootDiskType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumLocalSsds", GoMethod: "ResetNumLocalSsds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigDiskConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "instanceGroupManagerName", GoGetter: "InstanceGroupManagerName"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTemplateName", GoGetter: "InstanceTemplateName"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigManagedGroupConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accelerators", GoGetter: "Accelerators"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorsInput", GoGetter: "AcceleratorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskConfig", GoGetter: "DiskConfig"},
			_jsii_.MemberProperty{JsiiProperty: "diskConfigInput", GoGetter: "DiskConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imageInput", GoGetter: "ImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceNames", GoGetter: "InstanceNames"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isPreemptible", GoGetter: "IsPreemptible"},
			_jsii_.MemberProperty{JsiiProperty: "machineType", GoGetter: "MachineType"},
			_jsii_.MemberProperty{JsiiProperty: "machineTypeInput", GoGetter: "MachineTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "managedGroupConfig", GoGetter: "ManagedGroupConfig"},
			_jsii_.MemberProperty{JsiiProperty: "minCpuPlatform", GoGetter: "MinCpuPlatform"},
			_jsii_.MemberProperty{JsiiProperty: "minCpuPlatformInput", GoGetter: "MinCpuPlatformInput"},
			_jsii_.MemberProperty{JsiiProperty: "numInstances", GoGetter: "NumInstances"},
			_jsii_.MemberProperty{JsiiProperty: "numInstancesInput", GoGetter: "NumInstancesInput"},
			_jsii_.MemberProperty{JsiiProperty: "preemptibility", GoGetter: "Preemptibility"},
			_jsii_.MemberProperty{JsiiProperty: "preemptibilityInput", GoGetter: "PreemptibilityInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccelerators", GoMethod: "PutAccelerators"},
			_jsii_.MemberMethod{JsiiMethod: "putDiskConfig", GoMethod: "PutDiskConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccelerators", GoMethod: "ResetAccelerators"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskConfig", GoMethod: "ResetDiskConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetImage", GoMethod: "ResetImage"},
			_jsii_.MemberMethod{JsiiMethod: "resetMachineType", GoMethod: "ResetMachineType"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinCpuPlatform", GoMethod: "ResetMinCpuPlatform"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumInstances", GoMethod: "ResetNumInstances"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreemptibility", GoMethod: "ResetPreemptibility"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMasterConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "dataprocMetastoreService", GoGetter: "DataprocMetastoreService"},
			_jsii_.MemberProperty{JsiiProperty: "dataprocMetastoreServiceInput", GoGetter: "DataprocMetastoreServiceInput"},
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
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "autoscalingConfig", GoGetter: "AutoscalingConfig"},
			_jsii_.MemberProperty{JsiiProperty: "autoscalingConfigInput", GoGetter: "AutoscalingConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfig", GoGetter: "EncryptionConfig"},
			_jsii_.MemberProperty{JsiiProperty: "encryptionConfigInput", GoGetter: "EncryptionConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfig", GoGetter: "EndpointConfig"},
			_jsii_.MemberProperty{JsiiProperty: "endpointConfigInput", GoGetter: "EndpointConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "gceClusterConfig", GoGetter: "GceClusterConfig"},
			_jsii_.MemberProperty{JsiiProperty: "gceClusterConfigInput", GoGetter: "GceClusterConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "getAnyMapAttribute", GoMethod: "GetAnyMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanAttribute", GoMethod: "GetBooleanAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getBooleanMapAttribute", GoMethod: "GetBooleanMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getListAttribute", GoMethod: "GetListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberAttribute", GoMethod: "GetNumberAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberListAttribute", GoMethod: "GetNumberListAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getNumberMapAttribute", GoMethod: "GetNumberMapAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringAttribute", GoMethod: "GetStringAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "getStringMapAttribute", GoMethod: "GetStringMapAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "gkeClusterConfig", GoGetter: "GkeClusterConfig"},
			_jsii_.MemberProperty{JsiiProperty: "gkeClusterConfigInput", GoGetter: "GkeClusterConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "initializationActions", GoGetter: "InitializationActions"},
			_jsii_.MemberProperty{JsiiProperty: "initializationActionsInput", GoGetter: "InitializationActionsInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfig", GoGetter: "LifecycleConfig"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycleConfigInput", GoGetter: "LifecycleConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "masterConfig", GoGetter: "MasterConfig"},
			_jsii_.MemberProperty{JsiiProperty: "masterConfigInput", GoGetter: "MasterConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "metastoreConfig", GoGetter: "MetastoreConfig"},
			_jsii_.MemberProperty{JsiiProperty: "metastoreConfigInput", GoGetter: "MetastoreConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAutoscalingConfig", GoMethod: "PutAutoscalingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putEncryptionConfig", GoMethod: "PutEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putEndpointConfig", GoMethod: "PutEndpointConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putGceClusterConfig", GoMethod: "PutGceClusterConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putGkeClusterConfig", GoMethod: "PutGkeClusterConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putInitializationActions", GoMethod: "PutInitializationActions"},
			_jsii_.MemberMethod{JsiiMethod: "putLifecycleConfig", GoMethod: "PutLifecycleConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putMasterConfig", GoMethod: "PutMasterConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putMetastoreConfig", GoMethod: "PutMetastoreConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSecondaryWorkerConfig", GoMethod: "PutSecondaryWorkerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSecurityConfig", GoMethod: "PutSecurityConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putSoftwareConfig", GoMethod: "PutSoftwareConfig"},
			_jsii_.MemberMethod{JsiiMethod: "putWorkerConfig", GoMethod: "PutWorkerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAutoscalingConfig", GoMethod: "ResetAutoscalingConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetEncryptionConfig", GoMethod: "ResetEncryptionConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetEndpointConfig", GoMethod: "ResetEndpointConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetGceClusterConfig", GoMethod: "ResetGceClusterConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetGkeClusterConfig", GoMethod: "ResetGkeClusterConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetInitializationActions", GoMethod: "ResetInitializationActions"},
			_jsii_.MemberMethod{JsiiMethod: "resetLifecycleConfig", GoMethod: "ResetLifecycleConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMasterConfig", GoMethod: "ResetMasterConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetMetastoreConfig", GoMethod: "ResetMetastoreConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecondaryWorkerConfig", GoMethod: "ResetSecondaryWorkerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetSecurityConfig", GoMethod: "ResetSecurityConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetSoftwareConfig", GoMethod: "ResetSoftwareConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetStagingBucket", GoMethod: "ResetStagingBucket"},
			_jsii_.MemberMethod{JsiiMethod: "resetTempBucket", GoMethod: "ResetTempBucket"},
			_jsii_.MemberMethod{JsiiMethod: "resetWorkerConfig", GoMethod: "ResetWorkerConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryWorkerConfig", GoGetter: "SecondaryWorkerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "secondaryWorkerConfigInput", GoGetter: "SecondaryWorkerConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfig", GoGetter: "SecurityConfig"},
			_jsii_.MemberProperty{JsiiProperty: "securityConfigInput", GoGetter: "SecurityConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "softwareConfig", GoGetter: "SoftwareConfig"},
			_jsii_.MemberProperty{JsiiProperty: "softwareConfigInput", GoGetter: "SoftwareConfigInput"},
			_jsii_.MemberProperty{JsiiProperty: "stagingBucket", GoGetter: "StagingBucket"},
			_jsii_.MemberProperty{JsiiProperty: "stagingBucketInput", GoGetter: "StagingBucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "tempBucket", GoGetter: "TempBucket"},
			_jsii_.MemberProperty{JsiiProperty: "tempBucketInput", GoGetter: "TempBucketInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "workerConfig", GoGetter: "WorkerConfig"},
			_jsii_.MemberProperty{JsiiProperty: "workerConfigInput", GoGetter: "WorkerConfigInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAccelerators)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceleratorCount", GoGetter: "AcceleratorCount"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorCountInput", GoGetter: "AcceleratorCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorType", GoGetter: "AcceleratorType"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorTypeInput", GoGetter: "AcceleratorTypeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorCount", GoMethod: "ResetAcceleratorCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorType", GoMethod: "ResetAcceleratorType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigAcceleratorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bootDiskSizeGb", GoGetter: "BootDiskSizeGb"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiskSizeGbInput", GoGetter: "BootDiskSizeGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiskType", GoGetter: "BootDiskType"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiskTypeInput", GoGetter: "BootDiskTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "numLocalSsds", GoGetter: "NumLocalSsds"},
			_jsii_.MemberProperty{JsiiProperty: "numLocalSsdsInput", GoGetter: "NumLocalSsdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootDiskSizeGb", GoMethod: "ResetBootDiskSizeGb"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootDiskType", GoMethod: "ResetBootDiskType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumLocalSsds", GoMethod: "ResetNumLocalSsds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigDiskConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "instanceGroupManagerName", GoGetter: "InstanceGroupManagerName"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTemplateName", GoGetter: "InstanceTemplateName"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigManagedGroupConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accelerators", GoGetter: "Accelerators"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorsInput", GoGetter: "AcceleratorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskConfig", GoGetter: "DiskConfig"},
			_jsii_.MemberProperty{JsiiProperty: "diskConfigInput", GoGetter: "DiskConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imageInput", GoGetter: "ImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceNames", GoGetter: "InstanceNames"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isPreemptible", GoGetter: "IsPreemptible"},
			_jsii_.MemberProperty{JsiiProperty: "machineType", GoGetter: "MachineType"},
			_jsii_.MemberProperty{JsiiProperty: "machineTypeInput", GoGetter: "MachineTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "managedGroupConfig", GoGetter: "ManagedGroupConfig"},
			_jsii_.MemberProperty{JsiiProperty: "minCpuPlatform", GoGetter: "MinCpuPlatform"},
			_jsii_.MemberProperty{JsiiProperty: "minCpuPlatformInput", GoGetter: "MinCpuPlatformInput"},
			_jsii_.MemberProperty{JsiiProperty: "numInstances", GoGetter: "NumInstances"},
			_jsii_.MemberProperty{JsiiProperty: "numInstancesInput", GoGetter: "NumInstancesInput"},
			_jsii_.MemberProperty{JsiiProperty: "preemptibility", GoGetter: "Preemptibility"},
			_jsii_.MemberProperty{JsiiProperty: "preemptibilityInput", GoGetter: "PreemptibilityInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccelerators", GoMethod: "PutAccelerators"},
			_jsii_.MemberMethod{JsiiMethod: "putDiskConfig", GoMethod: "PutDiskConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccelerators", GoMethod: "ResetAccelerators"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskConfig", GoMethod: "ResetDiskConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetImage", GoMethod: "ResetImage"},
			_jsii_.MemberMethod{JsiiMethod: "resetMachineType", GoMethod: "ResetMachineType"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinCpuPlatform", GoMethod: "ResetMinCpuPlatform"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumInstances", GoMethod: "ResetNumInstances"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreemptibility", GoMethod: "ResetPreemptibility"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecondaryWorkerConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "crossRealmTrustAdminServer", GoGetter: "CrossRealmTrustAdminServer"},
			_jsii_.MemberProperty{JsiiProperty: "crossRealmTrustAdminServerInput", GoGetter: "CrossRealmTrustAdminServerInput"},
			_jsii_.MemberProperty{JsiiProperty: "crossRealmTrustKdc", GoGetter: "CrossRealmTrustKdc"},
			_jsii_.MemberProperty{JsiiProperty: "crossRealmTrustKdcInput", GoGetter: "CrossRealmTrustKdcInput"},
			_jsii_.MemberProperty{JsiiProperty: "crossRealmTrustRealm", GoGetter: "CrossRealmTrustRealm"},
			_jsii_.MemberProperty{JsiiProperty: "crossRealmTrustRealmInput", GoGetter: "CrossRealmTrustRealmInput"},
			_jsii_.MemberProperty{JsiiProperty: "crossRealmTrustSharedPassword", GoGetter: "CrossRealmTrustSharedPassword"},
			_jsii_.MemberProperty{JsiiProperty: "crossRealmTrustSharedPasswordInput", GoGetter: "CrossRealmTrustSharedPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "enableKerberos", GoGetter: "EnableKerberos"},
			_jsii_.MemberProperty{JsiiProperty: "enableKerberosInput", GoGetter: "EnableKerberosInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "kdcDbKey", GoGetter: "KdcDbKey"},
			_jsii_.MemberProperty{JsiiProperty: "kdcDbKeyInput", GoGetter: "KdcDbKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "keyPassword", GoGetter: "KeyPassword"},
			_jsii_.MemberProperty{JsiiProperty: "keyPasswordInput", GoGetter: "KeyPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "keystore", GoGetter: "Keystore"},
			_jsii_.MemberProperty{JsiiProperty: "keystoreInput", GoGetter: "KeystoreInput"},
			_jsii_.MemberProperty{JsiiProperty: "keystorePassword", GoGetter: "KeystorePassword"},
			_jsii_.MemberProperty{JsiiProperty: "keystorePasswordInput", GoGetter: "KeystorePasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKey", GoGetter: "KmsKey"},
			_jsii_.MemberProperty{JsiiProperty: "kmsKeyInput", GoGetter: "KmsKeyInput"},
			_jsii_.MemberProperty{JsiiProperty: "realm", GoGetter: "Realm"},
			_jsii_.MemberProperty{JsiiProperty: "realmInput", GoGetter: "RealmInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrossRealmTrustAdminServer", GoMethod: "ResetCrossRealmTrustAdminServer"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrossRealmTrustKdc", GoMethod: "ResetCrossRealmTrustKdc"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrossRealmTrustRealm", GoMethod: "ResetCrossRealmTrustRealm"},
			_jsii_.MemberMethod{JsiiMethod: "resetCrossRealmTrustSharedPassword", GoMethod: "ResetCrossRealmTrustSharedPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetEnableKerberos", GoMethod: "ResetEnableKerberos"},
			_jsii_.MemberMethod{JsiiMethod: "resetKdcDbKey", GoMethod: "ResetKdcDbKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeyPassword", GoMethod: "ResetKeyPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeystore", GoMethod: "ResetKeystore"},
			_jsii_.MemberMethod{JsiiMethod: "resetKeystorePassword", GoMethod: "ResetKeystorePassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetKmsKey", GoMethod: "ResetKmsKey"},
			_jsii_.MemberMethod{JsiiMethod: "resetRealm", GoMethod: "ResetRealm"},
			_jsii_.MemberMethod{JsiiMethod: "resetRootPrincipalPassword", GoMethod: "ResetRootPrincipalPassword"},
			_jsii_.MemberMethod{JsiiMethod: "resetTgtLifetimeHours", GoMethod: "ResetTgtLifetimeHours"},
			_jsii_.MemberMethod{JsiiMethod: "resetTruststore", GoMethod: "ResetTruststore"},
			_jsii_.MemberMethod{JsiiMethod: "resetTruststorePassword", GoMethod: "ResetTruststorePassword"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "rootPrincipalPassword", GoGetter: "RootPrincipalPassword"},
			_jsii_.MemberProperty{JsiiProperty: "rootPrincipalPasswordInput", GoGetter: "RootPrincipalPasswordInput"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberProperty{JsiiProperty: "tgtLifetimeHours", GoGetter: "TgtLifetimeHours"},
			_jsii_.MemberProperty{JsiiProperty: "tgtLifetimeHoursInput", GoGetter: "TgtLifetimeHoursInput"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "truststore", GoGetter: "Truststore"},
			_jsii_.MemberProperty{JsiiProperty: "truststoreInput", GoGetter: "TruststoreInput"},
			_jsii_.MemberProperty{JsiiProperty: "truststorePassword", GoGetter: "TruststorePassword"},
			_jsii_.MemberProperty{JsiiProperty: "truststorePasswordInput", GoGetter: "TruststorePasswordInput"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigKerberosConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "kerberosConfig", GoGetter: "KerberosConfig"},
			_jsii_.MemberProperty{JsiiProperty: "kerberosConfigInput", GoGetter: "KerberosConfigInput"},
			_jsii_.MemberMethod{JsiiMethod: "putKerberosConfig", GoMethod: "PutKerberosConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetKerberosConfig", GoMethod: "ResetKerberosConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSecurityConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "imageVersion", GoGetter: "ImageVersion"},
			_jsii_.MemberProperty{JsiiProperty: "imageVersionInput", GoGetter: "ImageVersionInput"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "optionalComponents", GoGetter: "OptionalComponents"},
			_jsii_.MemberProperty{JsiiProperty: "optionalComponentsInput", GoGetter: "OptionalComponentsInput"},
			_jsii_.MemberProperty{JsiiProperty: "properties", GoGetter: "Properties"},
			_jsii_.MemberProperty{JsiiProperty: "propertiesInput", GoGetter: "PropertiesInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetImageVersion", GoMethod: "ResetImageVersion"},
			_jsii_.MemberMethod{JsiiMethod: "resetOptionalComponents", GoMethod: "ResetOptionalComponents"},
			_jsii_.MemberMethod{JsiiMethod: "resetProperties", GoMethod: "ResetProperties"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigSoftwareConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfig)(nil)).Elem(),
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAccelerators)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsList)(nil)).Elem(),
		[]_jsii_.Member{
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
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "acceleratorCount", GoGetter: "AcceleratorCount"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorCountInput", GoGetter: "AcceleratorCountInput"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorType", GoGetter: "AcceleratorType"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorTypeInput", GoGetter: "AcceleratorTypeInput"},
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
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorCount", GoMethod: "ResetAcceleratorCount"},
			_jsii_.MemberMethod{JsiiMethod: "resetAcceleratorType", GoMethod: "ResetAcceleratorType"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigAcceleratorsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "bootDiskSizeGb", GoGetter: "BootDiskSizeGb"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiskSizeGbInput", GoGetter: "BootDiskSizeGbInput"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiskType", GoGetter: "BootDiskType"},
			_jsii_.MemberProperty{JsiiProperty: "bootDiskTypeInput", GoGetter: "BootDiskTypeInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "numLocalSsds", GoGetter: "NumLocalSsds"},
			_jsii_.MemberProperty{JsiiProperty: "numLocalSsdsInput", GoGetter: "NumLocalSsdsInput"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootDiskSizeGb", GoMethod: "ResetBootDiskSizeGb"},
			_jsii_.MemberMethod{JsiiMethod: "resetBootDiskType", GoMethod: "ResetBootDiskType"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumLocalSsds", GoMethod: "ResetNumLocalSsds"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigDiskConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfig)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberMethod{JsiiMethod: "get", GoMethod: "Get"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberProperty{JsiiProperty: "wrapsSet", GoGetter: "WrapsSet"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigList{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexList)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigOutputReference)(nil)).Elem(),
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
			_jsii_.MemberProperty{JsiiProperty: "instanceGroupManagerName", GoGetter: "InstanceGroupManagerName"},
			_jsii_.MemberProperty{JsiiProperty: "instanceTemplateName", GoGetter: "InstanceTemplateName"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigManagedGroupConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "accelerators", GoGetter: "Accelerators"},
			_jsii_.MemberProperty{JsiiProperty: "acceleratorsInput", GoGetter: "AcceleratorsInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "creationStack", GoGetter: "CreationStack"},
			_jsii_.MemberProperty{JsiiProperty: "diskConfig", GoGetter: "DiskConfig"},
			_jsii_.MemberProperty{JsiiProperty: "diskConfigInput", GoGetter: "DiskConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "image", GoGetter: "Image"},
			_jsii_.MemberProperty{JsiiProperty: "imageInput", GoGetter: "ImageInput"},
			_jsii_.MemberProperty{JsiiProperty: "instanceNames", GoGetter: "InstanceNames"},
			_jsii_.MemberProperty{JsiiProperty: "internalValue", GoGetter: "InternalValue"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationAsList", GoMethod: "InterpolationAsList"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "isPreemptible", GoGetter: "IsPreemptible"},
			_jsii_.MemberProperty{JsiiProperty: "machineType", GoGetter: "MachineType"},
			_jsii_.MemberProperty{JsiiProperty: "machineTypeInput", GoGetter: "MachineTypeInput"},
			_jsii_.MemberProperty{JsiiProperty: "managedGroupConfig", GoGetter: "ManagedGroupConfig"},
			_jsii_.MemberProperty{JsiiProperty: "minCpuPlatform", GoGetter: "MinCpuPlatform"},
			_jsii_.MemberProperty{JsiiProperty: "minCpuPlatformInput", GoGetter: "MinCpuPlatformInput"},
			_jsii_.MemberProperty{JsiiProperty: "numInstances", GoGetter: "NumInstances"},
			_jsii_.MemberProperty{JsiiProperty: "numInstancesInput", GoGetter: "NumInstancesInput"},
			_jsii_.MemberProperty{JsiiProperty: "preemptibility", GoGetter: "Preemptibility"},
			_jsii_.MemberProperty{JsiiProperty: "preemptibilityInput", GoGetter: "PreemptibilityInput"},
			_jsii_.MemberMethod{JsiiMethod: "putAccelerators", GoMethod: "PutAccelerators"},
			_jsii_.MemberMethod{JsiiMethod: "putDiskConfig", GoMethod: "PutDiskConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetAccelerators", GoMethod: "ResetAccelerators"},
			_jsii_.MemberMethod{JsiiMethod: "resetDiskConfig", GoMethod: "ResetDiskConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetImage", GoMethod: "ResetImage"},
			_jsii_.MemberMethod{JsiiMethod: "resetMachineType", GoMethod: "ResetMachineType"},
			_jsii_.MemberMethod{JsiiMethod: "resetMinCpuPlatform", GoMethod: "ResetMinCpuPlatform"},
			_jsii_.MemberMethod{JsiiMethod: "resetNumInstances", GoMethod: "ResetNumInstances"},
			_jsii_.MemberMethod{JsiiMethod: "resetPreemptibility", GoMethod: "ResetPreemptibility"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigWorkerConfigOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementManagedClusterOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementManagedClusterOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clusterName", GoGetter: "ClusterName"},
			_jsii_.MemberProperty{JsiiProperty: "clusterNameInput", GoGetter: "ClusterNameInput"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIndex", GoGetter: "ComplexObjectIndex"},
			_jsii_.MemberProperty{JsiiProperty: "complexObjectIsFromSet", GoGetter: "ComplexObjectIsFromSet"},
			_jsii_.MemberMethod{JsiiMethod: "computeFqn", GoMethod: "ComputeFqn"},
			_jsii_.MemberProperty{JsiiProperty: "config", GoGetter: "Config"},
			_jsii_.MemberProperty{JsiiProperty: "configInput", GoGetter: "ConfigInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "labelsInput", GoGetter: "LabelsInput"},
			_jsii_.MemberMethod{JsiiMethod: "putConfig", GoMethod: "PutConfig"},
			_jsii_.MemberMethod{JsiiMethod: "resetLabels", GoMethod: "ResetLabels"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementManagedClusterOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplatePlacementOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplatePlacementOutputReference)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberProperty{JsiiProperty: "clusterSelector", GoGetter: "ClusterSelector"},
			_jsii_.MemberProperty{JsiiProperty: "clusterSelectorInput", GoGetter: "ClusterSelectorInput"},
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
			_jsii_.MemberProperty{JsiiProperty: "managedCluster", GoGetter: "ManagedCluster"},
			_jsii_.MemberProperty{JsiiProperty: "managedClusterInput", GoGetter: "ManagedClusterInput"},
			_jsii_.MemberMethod{JsiiMethod: "putClusterSelector", GoMethod: "PutClusterSelector"},
			_jsii_.MemberMethod{JsiiMethod: "putManagedCluster", GoMethod: "PutManagedCluster"},
			_jsii_.MemberMethod{JsiiMethod: "resetClusterSelector", GoMethod: "ResetClusterSelector"},
			_jsii_.MemberMethod{JsiiMethod: "resetManagedCluster", GoMethod: "ResetManagedCluster"},
			_jsii_.MemberMethod{JsiiMethod: "resolve", GoMethod: "Resolve"},
			_jsii_.MemberProperty{JsiiProperty: "terraformAttribute", GoGetter: "TerraformAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResource", GoGetter: "TerraformResource"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
		},
		func() interface{} {
			j := jsiiProxy_GoogleDataprocWorkflowTemplatePlacementOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateTimeouts",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateTimeouts)(nil)).Elem(),
	)
	_jsii_.RegisterClass(
		"google-beta.googleDataprocWorkflowTemplate.GoogleDataprocWorkflowTemplateTimeoutsOutputReference",
		reflect.TypeOf((*GoogleDataprocWorkflowTemplateTimeoutsOutputReference)(nil)).Elem(),
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
			j := jsiiProxy_GoogleDataprocWorkflowTemplateTimeoutsOutputReference{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfComplexObject)
			return &j
		},
	)
}