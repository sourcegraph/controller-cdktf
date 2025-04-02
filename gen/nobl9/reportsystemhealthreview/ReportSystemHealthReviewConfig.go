package reportsystemhealthreview

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ReportSystemHealthReviewConfig struct {
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
	// column block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#column ReportSystemHealthReview#column}
	Column interface{} `field:"required" json:"column" yaml:"column"`
	// Unique name of the resource, must conform to the naming convention from [DNS RFC1123](https://kubernetes.io/docs/concepts/overview/working-with-objects/names/#names).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#name ReportSystemHealthReview#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Grouping methods of report table rows [project/service].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#row_group_by ReportSystemHealthReview#row_group_by}
	RowGroupBy *string `field:"required" json:"rowGroupBy" yaml:"rowGroupBy"`
	// thresholds block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#thresholds ReportSystemHealthReview#thresholds}
	Thresholds *ReportSystemHealthReviewThresholds `field:"required" json:"thresholds" yaml:"thresholds"`
	// time_frame block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#time_frame ReportSystemHealthReview#time_frame}
	TimeFrame *ReportSystemHealthReviewTimeFrame `field:"required" json:"timeFrame" yaml:"timeFrame"`
	// User-friendly display name of the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#display_name ReportSystemHealthReview#display_name}
	DisplayName *string `field:"optional" json:"displayName" yaml:"displayName"`
	// filters block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#filters ReportSystemHealthReview#filters}
	Filters *ReportSystemHealthReviewFilters `field:"optional" json:"filters" yaml:"filters"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#id ReportSystemHealthReview#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Is report shared for all users with access to included projects.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/nobl9/nobl9/0.37.0/docs/resources/report_system_health_review#shared ReportSystemHealthReview#shared}
	Shared interface{} `field:"optional" json:"shared" yaml:"shared"`
}

