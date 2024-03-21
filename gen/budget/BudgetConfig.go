package budget

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BudgetConfig struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The amount to use as the budget.
	Amount *float64 `field:"required" json:"amount" yaml:"amount"`
	// ID of the billing account to set a budget on.
	BillingAccount *string `field:"required" json:"billingAccount" yaml:"billingAccount"`
	// The project ids to include in this budget.
	//
	// If empty budget will include all projects.
	Projects *[]*string `field:"required" json:"projects" yaml:"projects"`
	// The name of the Cloud Pub/Sub topic where budget related messages will be published, in the form of `projects/{project_id}/topics/{topic_id}`.
	AlertPubsubTopic *string `field:"optional" json:"alertPubsubTopic" yaml:"alertPubsubTopic"`
	// The type of basis used to determine if spend has passed the threshold CURRENT_SPEND.
	AlertSpendBasis *string `field:"optional" json:"alertSpendBasis" yaml:"alertSpendBasis"`
	// A list of percentages of the budget to alert on when threshold is exceeded 0.5 0.7 1.
	AlertSpentPercents *[]*float64 `field:"optional" json:"alertSpentPercents" yaml:"alertSpentPercents"`
	// Specifies the calendar period for the budget.
	//
	// Possible values are MONTH, QUARTER, YEAR, CALENDAR_PERIOD_UNSPECIFIED, CUSTOM. custom_period_start_date and custom_period_end_date must be set if CUSTOM
	CalendarPeriod *string `field:"optional" json:"calendarPeriod" yaml:"calendarPeriod"`
	// If the budget should be created true.
	CreateBudget *bool `field:"optional" json:"createBudget" yaml:"createBudget"`
	// Specifies how credits should be treated when determining spend for threshold calculations INCLUDE_ALL_CREDITS.
	CreditTypesTreatment *string `field:"optional" json:"creditTypesTreatment" yaml:"creditTypesTreatment"`
	// Specifies the end date (DD-MM-YYYY) for the calendar_period CUSTOM.
	CustomPeriodEndDate *string `field:"optional" json:"customPeriodEndDate" yaml:"customPeriodEndDate"`
	// Specifies the start date (DD-MM-YYYY) for the calendar_period CUSTOM.
	CustomPeriodStartDate *string `field:"optional" json:"customPeriodStartDate" yaml:"customPeriodStartDate"`
	// The display name of the budget.
	//
	// If not set defaults to `Budget For <projects[0]|All Projects>`.
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// A single label and value pair specifying that usage from only this set of labeled resources should be included in the budget.
	//
	// The property type contains a map, they have special handling, please see {@link cdk.tf /module-map-inputs the docs}
	Labels *map[string]*string `field:"optional" json:"labels" yaml:"labels"`
	// A list of monitoring notification channels in the form `[projects/{project_id}/notificationChannels/{channel_id}]`.
	//
	// A maximum of 5 channels are allowed.
	MonitoringNotificationChannels *[]*string `field:"optional" json:"monitoringNotificationChannels" yaml:"monitoringNotificationChannels"`
	// A list of services ids to be included in the budget.
	//
	// If omitted, all services will be included in the budget. Service ids can be found at https://cloud.google.com/skus/
	Services *[]*string `field:"optional" json:"services" yaml:"services"`
}

