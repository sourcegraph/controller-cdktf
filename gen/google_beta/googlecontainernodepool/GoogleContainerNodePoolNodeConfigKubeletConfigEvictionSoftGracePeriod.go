package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigKubeletConfigEvictionSoftGracePeriod struct {
	// Defines grace period for the imagefs.available soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#imagefs_available GoogleContainerNodePool#imagefs_available}
	ImagefsAvailable *string `field:"optional" json:"imagefsAvailable" yaml:"imagefsAvailable"`
	// Defines grace period for the imagefs.inodesFree soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#imagefs_inodes_free GoogleContainerNodePool#imagefs_inodes_free}
	ImagefsInodesFree *string `field:"optional" json:"imagefsInodesFree" yaml:"imagefsInodesFree"`
	// Defines grace period for the memory.available soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#memory_available GoogleContainerNodePool#memory_available}
	MemoryAvailable *string `field:"optional" json:"memoryAvailable" yaml:"memoryAvailable"`
	// Defines grace period for the nodefs.available soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#nodefs_available GoogleContainerNodePool#nodefs_available}
	NodefsAvailable *string `field:"optional" json:"nodefsAvailable" yaml:"nodefsAvailable"`
	// Defines grace period for the nodefs.inodesFree soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#nodefs_inodes_free GoogleContainerNodePool#nodefs_inodes_free}
	NodefsInodesFree *string `field:"optional" json:"nodefsInodesFree" yaml:"nodefsInodesFree"`
	// Defines grace period for the pid.available soft eviction threshold.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#pid_available GoogleContainerNodePool#pid_available}
	PidAvailable *string `field:"optional" json:"pidAvailable" yaml:"pidAvailable"`
}

