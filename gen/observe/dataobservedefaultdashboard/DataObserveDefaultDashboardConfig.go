package dataobservedefaultdashboard

import (
	"github.com/open-constructs/cdk-terrain-go/cdktn"
)

type DataObserveDefaultDashboardConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count interface{} `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktn.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktn.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktn.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktn.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]interface{} `field:"optional" json:"provisioners" yaml:"provisioners"`
	// The OID of the dataset that the default dashboard is being set for.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/default_dashboard#dataset DataObserveDefaultDashboard#dataset}
	Dataset *string `field:"required" json:"dataset" yaml:"dataset"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/default_dashboard#id DataObserveDefaultDashboard#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

