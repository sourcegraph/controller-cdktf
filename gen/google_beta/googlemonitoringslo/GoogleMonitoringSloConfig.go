package googlemonitoringslo

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleMonitoringSloConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The fraction of service that must be good in order for this objective to be met.
	//
	// 0 < goal <= 0.999
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#goal GoogleMonitoringSlo#goal}
	Goal *float64 `field:"required" json:"goal" yaml:"goal"`
	// ID of the service to which this SLO belongs.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#service GoogleMonitoringSlo#service}
	Service *string `field:"required" json:"service" yaml:"service"`
	// basic_sli block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#basic_sli GoogleMonitoringSlo#basic_sli}
	BasicSli *GoogleMonitoringSloBasicSli `field:"optional" json:"basicSli" yaml:"basicSli"`
	// A calendar period, semantically "since the start of the current <calendarPeriod>". Possible values: ["DAY", "WEEK", "FORTNIGHT", "MONTH"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#calendar_period GoogleMonitoringSlo#calendar_period}
	CalendarPeriod *string `field:"optional" json:"calendarPeriod" yaml:"calendarPeriod"`
	// Name used for UI elements listing this SLO.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#display_name GoogleMonitoringSlo#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#id GoogleMonitoringSlo#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#project GoogleMonitoringSlo#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// request_based_sli block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#request_based_sli GoogleMonitoringSlo#request_based_sli}
	RequestBasedSli *GoogleMonitoringSloRequestBasedSli `field:"optional" json:"requestBasedSli" yaml:"requestBasedSli"`
	// A rolling time period, semantically "in the past X days". Must be between 1 to 30 days, inclusive.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#rolling_period_days GoogleMonitoringSlo#rolling_period_days}
	RollingPeriodDays *float64 `field:"optional" json:"rollingPeriodDays" yaml:"rollingPeriodDays"`
	// The id to use for this ServiceLevelObjective. If omitted, an id will be generated instead.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#slo_id GoogleMonitoringSlo#slo_id}
	SloId *string `field:"optional" json:"sloId" yaml:"sloId"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#timeouts GoogleMonitoringSlo#timeouts}
	Timeouts *GoogleMonitoringSloTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// This field is intended to be used for organizing and identifying the AlertPolicy objects.The field can contain up to 64 entries. Each key and value is limited to 63 Unicode characters or 128 bytes, whichever is smaller. Labels and values can contain only lowercase letters, numerals, underscores, and dashes. Keys must begin with a letter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#user_labels GoogleMonitoringSlo#user_labels}
	UserLabels *map[string]*string `field:"optional" json:"userLabels" yaml:"userLabels"`
	// windows_based_sli block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_slo#windows_based_sli GoogleMonitoringSlo#windows_based_sli}
	WindowsBasedSli *GoogleMonitoringSloWindowsBasedSli `field:"optional" json:"windowsBasedSli" yaml:"windowsBasedSli"`
}

