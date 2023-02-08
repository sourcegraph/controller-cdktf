// executorsnetworking
package executorsnetworking

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExecutorsnetworkingOptions struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// The GCP region to create the network in.
	Region *string `field:"required" json:"region" yaml:"region"`
	// When true, the network will contain a NAT router.
	//
	// Use when executors should not get public IPs.
	Nat *bool `field:"optional" json:"nat" yaml:"nat"`
	// The minimum number of ports per VM to use when NAT mode is enabled.
	//
	// Consider increasing this when you see egress packets being dropped.
	NatMinPortsPerVm *float64 `field:"optional" json:"natMinPortsPerVm" yaml:"natMinPortsPerVm"`
}

