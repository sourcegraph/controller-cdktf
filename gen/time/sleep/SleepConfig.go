package sleep

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type SleepConfig struct {
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
	// [Time duration](https://golang.org/pkg/time/#ParseDuration) to delay resource creation. For example, `30s` for 30 seconds or `5m` for 5 minutes. Updating this value by itself will not trigger a delay.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/sleep#create_duration Sleep#create_duration}
	CreateDuration *string `field:"optional" json:"createDuration" yaml:"createDuration"`
	// [Time duration](https://golang.org/pkg/time/#ParseDuration) to delay resource destroy. For example, `30s` for 30 seconds or `5m` for 5 minutes. Updating this value by itself will not trigger a delay. This value or any updates to it must be successfully applied into the Terraform state before destroying this resource to take effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/sleep#destroy_duration Sleep#destroy_duration}
	DestroyDuration *string `field:"optional" json:"destroyDuration" yaml:"destroyDuration"`
	// (Optional) Arbitrary map of values that, when changed, will run any creation or destroy delays again.
	//
	// See [the main provider documentation](../index.md) for more information.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/time/0.9.1/docs/resources/sleep#triggers Sleep#triggers}
	Triggers *map[string]*string `field:"optional" json:"triggers" yaml:"triggers"`
}

