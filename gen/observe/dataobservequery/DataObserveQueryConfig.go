package dataobservequery

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type DataObserveQueryConfig struct {
	// Experimental.
	Connection any `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count any `field:"optional" json:"count" yaml:"count"`
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Lifecycle *cdktf.TerraformResourceLifecycle `field:"optional" json:"lifecycle" yaml:"lifecycle"`
	// Experimental.
	Provider cdktf.TerraformProvider `field:"optional" json:"provider" yaml:"provider"`
	// Experimental.
	Provisioners *[]any `field:"optional" json:"provisioners" yaml:"provisioners"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#inputs DataObserveQuery#inputs}.
	Inputs *map[string]*string `field:"required" json:"inputs" yaml:"inputs"`
	// stage block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#stage DataObserveQuery#stage}
	Stage any `field:"required" json:"stage" yaml:"stage"`
	// assert block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#assert DataObserveQuery#assert}
	Assert *DataObserveQueryAssert `field:"optional" json:"assert" yaml:"assert"`
	// End timestamp. If omitted, query will be periodically re-run until results are returned.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#end DataObserveQuery#end}
	End *string `field:"optional" json:"end" yaml:"end"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#id DataObserveQuery#id}.
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"optional" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#limit DataObserveQuery#limit}.
	Limit *float64 `field:"optional" json:"limit" yaml:"limit"`
	// poll block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#poll DataObserveQuery#poll}
	Poll *DataObserveQueryPoll `field:"optional" json:"poll" yaml:"poll"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/data-sources/query#start DataObserveQuery#start}.
	Start *string `field:"optional" json:"start" yaml:"start"`
}
