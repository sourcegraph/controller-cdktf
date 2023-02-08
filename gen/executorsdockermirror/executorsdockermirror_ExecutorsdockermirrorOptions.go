// executorsdockermirror
package executorsdockermirror

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type ExecutorsdockermirrorOptions struct {
	// Experimental.
	DependsOn *[]cdktf.ITerraformDependable `field:"optional" json:"dependsOn" yaml:"dependsOn"`
	// Experimental.
	ForEach cdktf.ITerraformIterator `field:"optional" json:"forEach" yaml:"forEach"`
	// Experimental.
	Providers *[]interface{} `field:"optional" json:"providers" yaml:"providers"`
	// Experimental.
	SkipAssetCreationFromLocalModules *bool `field:"optional" json:"skipAssetCreationFromLocalModules" yaml:"skipAssetCreationFromLocalModules"`
	// A label tag to add to all the machines;
	//
	// can be used for filtering out the right instances in stackdriver monitoring and in Prometheus instance discovery.
	InstanceTagPrefix *string `field:"required" json:"instanceTagPrefix" yaml:"instanceTagPrefix"`
	// The network to run the VM in.
	NetworkId *string `field:"required" json:"networkId" yaml:"networkId"`
	// The subnet to run the VM in.
	SubnetId *string `field:"required" json:"subnetId" yaml:"subnetId"`
	// The Google zone to create the docker mirror resources in.
	Zone *string `field:"required" json:"zone" yaml:"zone"`
	// If false, no public IP will be associated with the executors.
	AssignPublicIp *bool `field:"optional" json:"assignPublicIp" yaml:"assignPublicIp"`
	// Docker registry mirror node disk size in GB.
	BootDiskSize *float64 `field:"optional" json:"bootDiskSize" yaml:"bootDiskSize"`
	// Persistent Docker registry mirror disk size in GB.
	DiskSize *float64 `field:"optional" json:"diskSize" yaml:"diskSize"`
	// CIDR range from where HTTP access to the Docker registry is acceptable.
	HttpAccessCidrRanges *[]*string `field:"optional" json:"httpAccessCidrRanges" yaml:"httpAccessCidrRanges"`
	// DEPRECATED.
	//
	// This is not used anymore.
	HttpMetricsAccessCidrRanges *[]*string `field:"optional" json:"httpMetricsAccessCidrRanges" yaml:"httpMetricsAccessCidrRanges"`
	// Docker registry mirror node machine disk image to use for creating the boot volume.
	//
	// Leave empty to use latest compatible with the Sourcegraph version.
	MachineImage *string `field:"optional" json:"machineImage" yaml:"machineImage"`
	// Docker registry mirror node machine type.
	MachineType *string `field:"optional" json:"machineType" yaml:"machineType"`
}

