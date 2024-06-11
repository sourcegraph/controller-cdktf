package offset

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-time.offset.Offset",
		reflect.TypeOf((*Offset)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addMoveTarget", GoMethod: "AddMoveTarget"},
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberProperty{JsiiProperty: "baseRfc3339", GoGetter: "BaseRfc3339"},
			_jsii_.MemberProperty{JsiiProperty: "baseRfc3339Input", GoGetter: "BaseRfc3339Input"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "connection", GoGetter: "Connection"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "count", GoGetter: "Count"},
			_jsii_.MemberProperty{JsiiProperty: "day", GoGetter: "Day"},
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
			_jsii_.MemberMethod{JsiiMethod: "hasResourceMove", GoMethod: "HasResourceMove"},
			_jsii_.MemberProperty{JsiiProperty: "hour", GoGetter: "Hour"},
			_jsii_.MemberProperty{JsiiProperty: "id", GoGetter: "Id"},
			_jsii_.MemberMethod{JsiiMethod: "importFrom", GoMethod: "ImportFrom"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForAttribute", GoMethod: "InterpolationForAttribute"},
			_jsii_.MemberProperty{JsiiProperty: "lifecycle", GoGetter: "Lifecycle"},
			_jsii_.MemberProperty{JsiiProperty: "minute", GoGetter: "Minute"},
			_jsii_.MemberProperty{JsiiProperty: "month", GoGetter: "Month"},
			_jsii_.MemberMethod{JsiiMethod: "moveFromId", GoMethod: "MoveFromId"},
			_jsii_.MemberMethod{JsiiMethod: "moveTo", GoMethod: "MoveTo"},
			_jsii_.MemberMethod{JsiiMethod: "moveToId", GoMethod: "MoveToId"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberProperty{JsiiProperty: "offsetDays", GoGetter: "OffsetDays"},
			_jsii_.MemberProperty{JsiiProperty: "offsetDaysInput", GoGetter: "OffsetDaysInput"},
			_jsii_.MemberProperty{JsiiProperty: "offsetHours", GoGetter: "OffsetHours"},
			_jsii_.MemberProperty{JsiiProperty: "offsetHoursInput", GoGetter: "OffsetHoursInput"},
			_jsii_.MemberProperty{JsiiProperty: "offsetMinutes", GoGetter: "OffsetMinutes"},
			_jsii_.MemberProperty{JsiiProperty: "offsetMinutesInput", GoGetter: "OffsetMinutesInput"},
			_jsii_.MemberProperty{JsiiProperty: "offsetMonths", GoGetter: "OffsetMonths"},
			_jsii_.MemberProperty{JsiiProperty: "offsetMonthsInput", GoGetter: "OffsetMonthsInput"},
			_jsii_.MemberProperty{JsiiProperty: "offsetSeconds", GoGetter: "OffsetSeconds"},
			_jsii_.MemberProperty{JsiiProperty: "offsetSecondsInput", GoGetter: "OffsetSecondsInput"},
			_jsii_.MemberProperty{JsiiProperty: "offsetYears", GoGetter: "OffsetYears"},
			_jsii_.MemberProperty{JsiiProperty: "offsetYearsInput", GoGetter: "OffsetYearsInput"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "provider", GoGetter: "Provider"},
			_jsii_.MemberProperty{JsiiProperty: "provisioners", GoGetter: "Provisioners"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetBaseRfc3339", GoMethod: "ResetBaseRfc3339"},
			_jsii_.MemberMethod{JsiiMethod: "resetOffsetDays", GoMethod: "ResetOffsetDays"},
			_jsii_.MemberMethod{JsiiMethod: "resetOffsetHours", GoMethod: "ResetOffsetHours"},
			_jsii_.MemberMethod{JsiiMethod: "resetOffsetMinutes", GoMethod: "ResetOffsetMinutes"},
			_jsii_.MemberMethod{JsiiMethod: "resetOffsetMonths", GoMethod: "ResetOffsetMonths"},
			_jsii_.MemberMethod{JsiiMethod: "resetOffsetSeconds", GoMethod: "ResetOffsetSeconds"},
			_jsii_.MemberMethod{JsiiMethod: "resetOffsetYears", GoMethod: "ResetOffsetYears"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberMethod{JsiiMethod: "resetTriggers", GoMethod: "ResetTriggers"},
			_jsii_.MemberProperty{JsiiProperty: "rfc3339", GoGetter: "Rfc3339"},
			_jsii_.MemberProperty{JsiiProperty: "second", GoGetter: "Second"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberProperty{JsiiProperty: "terraformGeneratorMetadata", GoGetter: "TerraformGeneratorMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "terraformMetaArguments", GoGetter: "TerraformMetaArguments"},
			_jsii_.MemberProperty{JsiiProperty: "terraformResourceType", GoGetter: "TerraformResourceType"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "triggers", GoGetter: "Triggers"},
			_jsii_.MemberProperty{JsiiProperty: "triggersInput", GoGetter: "TriggersInput"},
			_jsii_.MemberProperty{JsiiProperty: "unix", GoGetter: "Unix"},
			_jsii_.MemberProperty{JsiiProperty: "year", GoGetter: "Year"},
		},
		func() interface{} {
			j := jsiiProxy_Offset{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformResource)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-time.offset.OffsetConfig",
		reflect.TypeOf((*OffsetConfig)(nil)).Elem(),
	)
}
