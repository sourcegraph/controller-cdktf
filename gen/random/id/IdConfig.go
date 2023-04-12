package id

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type IdConfig struct {
	// Experimental.
	Connection interface{} `field:"optional" json:"connection" yaml:"connection"`
	// Experimental.
	Count *float64 `field:"optional" json:"count" yaml:"count"`
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
	// The number of random bytes to produce. The minimum value is 1, which produces eight bits of randomness.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/id#byte_length Id#byte_length}
	ByteLength *float64 `field:"required" json:"byteLength" yaml:"byteLength"`
	// Arbitrary map of values that, when changed, will trigger recreation of resource.
	//
	// See [the main provider documentation](../index.html) for more information.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/id#keepers Id#keepers}
	Keepers *map[string]*string `field:"optional" json:"keepers" yaml:"keepers"`
	// Arbitrary string to prefix the output value with.
	//
	// This string is supplied as-is, meaning it is not guaranteed to be URL-safe or base64 encoded.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/random/r/id#prefix Id#prefix}
	Prefix *string `field:"optional" json:"prefix" yaml:"prefix"`
}

