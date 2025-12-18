package datasetqueryfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DatasetQueryFilterConfig struct {
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
	// The OID of the dataset this query filter applies to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_query_filter#dataset DatasetQueryFilter#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// OPAL boolean expression string for the filter.
	//
	// For example, `body ~ "phoneNumber"`.
	// This is an exclusion filter, so all observations matching the filter will be filtered out.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_query_filter#filter DatasetQueryFilter#filter}
	Filter *string `field:"required" json:"filter" yaml:"filter"`
	// Human-readable name for the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_query_filter#label DatasetQueryFilter#label}
	Label *string `field:"required" json:"label" yaml:"label"`
	// Long-form description of the filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_query_filter#description DatasetQueryFilter#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Whether the filter is disabled by the user.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_query_filter#disabled DatasetQueryFilter#disabled}
	Disabled interface{} `field:"optional" json:"disabled" yaml:"disabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/dataset_query_filter#id DatasetQueryFilter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

