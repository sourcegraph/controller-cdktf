package googlecomputestoragepool

import (
	"github.com/hashicorp/terraform-cdk-go/cdktf"
)

type GoogleComputeStoragePoolConfig struct {
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
	// Name of the resource.
	//
	// Provided by the client when the resource is created.
	// The name must be 1-63 characters long, and comply with RFC1035.
	// Specifically, the name must be 1-63 characters long and match
	// the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?'
	// which means the first character must be a lowercase letter,
	// and all following characters must be a dash, lowercase letter, or digit,
	// except the last character, which cannot be a dash.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#name GoogleComputeStoragePool#name}
	Name *string `field:"required" json:"name" yaml:"name"`
	// Size, in GiB, of the storage pool. For more information about the size limits, see https://cloud.google.com/compute/docs/disks/storage-pools.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#pool_provisioned_capacity_gb GoogleComputeStoragePool#pool_provisioned_capacity_gb}
	PoolProvisionedCapacityGb *string `field:"required" json:"poolProvisionedCapacityGb" yaml:"poolProvisionedCapacityGb"`
	// Provisioned throughput, in MB/s, of the storage pool. Only relevant if the storage pool type is 'hyperdisk-balanced' or 'hyperdisk-throughput'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#pool_provisioned_throughput GoogleComputeStoragePool#pool_provisioned_throughput}
	PoolProvisionedThroughput *string `field:"required" json:"poolProvisionedThroughput" yaml:"poolProvisionedThroughput"`
	// Type of the storage pool. For example, the following are valid values:.
	//
	// * 'https://www.googleapis.com/compute/v1/projects/{project_id}/zones/{zone}/storagePoolTypes/hyperdisk-balanced'
	// * 'hyperdisk-throughput'
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#storage_pool_type GoogleComputeStoragePool#storage_pool_type}
	StoragePoolType *string `field:"required" json:"storagePoolType" yaml:"storagePoolType"`
	// Provisioning type of the byte capacity of the pool. Possible values: ["STANDARD", "ADVANCED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#capacity_provisioning_type GoogleComputeStoragePool#capacity_provisioning_type}
	CapacityProvisioningType *string `field:"optional" json:"capacityProvisioningType" yaml:"capacityProvisioningType"`
	// Whether Terraform will be prevented from destroying the StoragePool.
	//
	// When the field is set to true or unset in Terraform state, a 'terraform apply'
	// or 'terraform destroy' that would delete the StoragePool will fail.
	// When the field is set to false, deleting the StoragePool is allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#deletion_protection GoogleComputeStoragePool#deletion_protection}
	DeletionProtection interface{} `field:"optional" json:"deletionProtection" yaml:"deletionProtection"`
	// A description of this resource. Provide this property when you create the resource.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#description GoogleComputeStoragePool#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
	// Provisioning type of the performance-related parameters of the pool, such as throughput and IOPS. Possible values: ["STANDARD", "ADVANCED"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#performance_provisioning_type GoogleComputeStoragePool#performance_provisioning_type}
	PerformanceProvisioningType *string `field:"optional" json:"performanceProvisioningType" yaml:"performanceProvisioningType"`
	// Provisioned IOPS of the storage pool. Only relevant if the storage pool type is 'hyperdisk-balanced'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#pool_provisioned_iops GoogleComputeStoragePool#pool_provisioned_iops}
	PoolProvisionedIops *string `field:"optional" json:"poolProvisionedIops" yaml:"poolProvisionedIops"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#project GoogleComputeStoragePool#project}.
	Project *string `field:"optional" json:"project" yaml:"project"`
	// timeouts block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#timeouts GoogleComputeStoragePool#timeouts}
	Timeouts *GoogleComputeStoragePoolTimeouts `field:"optional" json:"timeouts" yaml:"timeouts"`
	// A reference to the zone where the storage pool resides.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.32.0/docs/resources/google_compute_storage_pool#zone GoogleComputeStoragePool#zone}
	Zone *string `field:"optional" json:"zone" yaml:"zone"`
}

