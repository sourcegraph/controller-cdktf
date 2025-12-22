package dropfilter

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DropFilterConfig struct {
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
	// The percentage of matched observations to drop specified as a floating point number between 0.0 and 1.0.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/drop_filter#drop_rate DropFilter#drop_rate}
	DropRate *float64 `field:"required" json:"dropRate" yaml:"dropRate"`
	// The name of the drop filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/drop_filter#name DropFilter#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// The opal that defines the drop filter.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/drop_filter#pipeline DropFilter#pipeline}
	Pipeline *string `field:"required" json:"pipeline" yaml:"pipeline"`
	// The source dataset that the drop filter should apply to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/drop_filter#source_dataset DropFilter#source_dataset}
	SourceDataset *string `field:"required" json:"sourceDataset" yaml:"sourceDataset"`
	// OID of the workspace this object is contained in.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/drop_filter#workspace DropFilter#workspace}
	Workspace *string `field:"required" json:"workspace" yaml:"workspace"`
	// When true the drop filter will drop data, and when false the drop filter will not drop data.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/drop_filter#enabled DropFilter#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/drop_filter#id DropFilter#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
}

