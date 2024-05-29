// @cdktf/provider-budget
package budget

import (
	"reflect"

	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
)

func init() {
	_jsii_.RegisterClass(
		"@cdktf/provider-budget.Budget",
		reflect.TypeOf((*Budget)(nil)).Elem(),
		[]_jsii_.Member{
			_jsii_.MemberMethod{JsiiMethod: "addOverride", GoMethod: "AddOverride"},
			_jsii_.MemberMethod{JsiiMethod: "addProvider", GoMethod: "AddProvider"},
			_jsii_.MemberProperty{JsiiProperty: "alertPubsubTopic", GoGetter: "AlertPubsubTopic"},
			_jsii_.MemberProperty{JsiiProperty: "alertSpendBasis", GoGetter: "AlertSpendBasis"},
			_jsii_.MemberProperty{JsiiProperty: "alertSpentPercents", GoGetter: "AlertSpentPercents"},
			_jsii_.MemberProperty{JsiiProperty: "amount", GoGetter: "Amount"},
			_jsii_.MemberProperty{JsiiProperty: "billingAccount", GoGetter: "BillingAccount"},
			_jsii_.MemberProperty{JsiiProperty: "calendarPeriod", GoGetter: "CalendarPeriod"},
			_jsii_.MemberProperty{JsiiProperty: "cdktfStack", GoGetter: "CdktfStack"},
			_jsii_.MemberProperty{JsiiProperty: "constructNodeMetadata", GoGetter: "ConstructNodeMetadata"},
			_jsii_.MemberProperty{JsiiProperty: "createBudget", GoGetter: "CreateBudget"},
			_jsii_.MemberProperty{JsiiProperty: "creditTypesTreatment", GoGetter: "CreditTypesTreatment"},
			_jsii_.MemberProperty{JsiiProperty: "customPeriodEndDate", GoGetter: "CustomPeriodEndDate"},
			_jsii_.MemberProperty{JsiiProperty: "customPeriodStartDate", GoGetter: "CustomPeriodStartDate"},
			_jsii_.MemberProperty{JsiiProperty: "dependsOn", GoGetter: "DependsOn"},
			_jsii_.MemberProperty{JsiiProperty: "displayName", GoGetter: "DisplayName"},
			_jsii_.MemberProperty{JsiiProperty: "forEach", GoGetter: "ForEach"},
			_jsii_.MemberProperty{JsiiProperty: "fqn", GoGetter: "Fqn"},
			_jsii_.MemberProperty{JsiiProperty: "friendlyUniqueId", GoGetter: "FriendlyUniqueId"},
			_jsii_.MemberMethod{JsiiMethod: "getString", GoMethod: "GetString"},
			_jsii_.MemberMethod{JsiiMethod: "interpolationForOutput", GoMethod: "InterpolationForOutput"},
			_jsii_.MemberProperty{JsiiProperty: "labels", GoGetter: "Labels"},
			_jsii_.MemberProperty{JsiiProperty: "monitoringNotificationChannels", GoGetter: "MonitoringNotificationChannels"},
			_jsii_.MemberProperty{JsiiProperty: "nameOutput", GoGetter: "NameOutput"},
			_jsii_.MemberProperty{JsiiProperty: "node", GoGetter: "Node"},
			_jsii_.MemberMethod{JsiiMethod: "overrideLogicalId", GoMethod: "OverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "projects", GoGetter: "Projects"},
			_jsii_.MemberProperty{JsiiProperty: "providers", GoGetter: "Providers"},
			_jsii_.MemberProperty{JsiiProperty: "rawOverrides", GoGetter: "RawOverrides"},
			_jsii_.MemberMethod{JsiiMethod: "resetOverrideLogicalId", GoMethod: "ResetOverrideLogicalId"},
			_jsii_.MemberProperty{JsiiProperty: "services", GoGetter: "Services"},
			_jsii_.MemberProperty{JsiiProperty: "skipAssetCreationFromLocalModules", GoGetter: "SkipAssetCreationFromLocalModules"},
			_jsii_.MemberProperty{JsiiProperty: "source", GoGetter: "Source"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeAttributes", GoMethod: "SynthesizeAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "synthesizeHclAttributes", GoMethod: "SynthesizeHclAttributes"},
			_jsii_.MemberMethod{JsiiMethod: "toHclTerraform", GoMethod: "ToHclTerraform"},
			_jsii_.MemberMethod{JsiiMethod: "toMetadata", GoMethod: "ToMetadata"},
			_jsii_.MemberMethod{JsiiMethod: "toString", GoMethod: "ToString"},
			_jsii_.MemberMethod{JsiiMethod: "toTerraform", GoMethod: "ToTerraform"},
			_jsii_.MemberProperty{JsiiProperty: "version", GoGetter: "Version"},
		},
		func() interface{} {
			j := jsiiProxy_Budget{}
			_jsii_.InitJsiiProxy(&j.Type__cdktfTerraformModule)
			return &j
		},
	)
	_jsii_.RegisterStruct(
		"@cdktf/provider-budget.BudgetConfig",
		reflect.TypeOf((*BudgetConfig)(nil)).Elem(),
	)
}
