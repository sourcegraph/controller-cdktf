package budgetadjustment

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type BudgetAdjustmentConfig struct {
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
	// The duration of the budget adjustment event.
	//
	// The expected value for this field is a string formatted as a time duration. The duration must be defined with a precision of 1 minute. Example: `1h10m`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/budget_adjustment#duration BudgetAdjustment#duration}
	Duration *string `field:"required" json:"duration" yaml:"duration"`
	// The time at which the first event is scheduled to start.
	//
	// The expected value must be a string representing the date and time in RFC3339 format. Example: `2022-12-31T00:00:00Z`
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/budget_adjustment#first_event_start BudgetAdjustment#first_event_start}
	FirstEventStart *string `field:"required" json:"firstEventStart" yaml:"firstEventStart"`
	// Unique name of the resource, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/budget_adjustment#name BudgetAdjustment#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Optional description of the resource.
	//
	// Here, you can add details about who is responsible for the integration (team/owner) or the purpose of creating it.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/budget_adjustment#description BudgetAdjustment#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// User-friendly display name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/budget_adjustment#display_name BudgetAdjustment#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/budget_adjustment#filters BudgetAdjustment#filters}
	Filters interface{} `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/budget_adjustment#id BudgetAdjustment#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// The recurrence rule for the budget adjustment event. The expected value is a string in RRULE format. Example: `FREQ=MONTHLY;BYMONTHDAY=1`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/budget_adjustment#rrule BudgetAdjustment#rrule}
	Rrule *string `field:"optional" json:"rrule" yaml:"rrule"`
}

