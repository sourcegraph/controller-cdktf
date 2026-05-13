package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigKubeletConfigEvictionMinimumReclaim struct {
	// Defines percentage of minimum reclaim for imagefs.available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#imagefs_available GoogleContainerNodePool#imagefs_available}
	ImagefsAvailable *string `field:"optional" json:"imagefsAvailable" yaml:"imagefsAvailable"`
	// Defines percentage of minimum reclaim for imagefs.inodesFree.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#imagefs_inodes_free GoogleContainerNodePool#imagefs_inodes_free}
	ImagefsInodesFree *string `field:"optional" json:"imagefsInodesFree" yaml:"imagefsInodesFree"`
	// Defines percentage of minimum reclaim for memory.available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#memory_available GoogleContainerNodePool#memory_available}
	MemoryAvailable *string `field:"optional" json:"memoryAvailable" yaml:"memoryAvailable"`
	// Defines percentage of minimum reclaim for nodefs.available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#nodefs_available GoogleContainerNodePool#nodefs_available}
	NodefsAvailable *string `field:"optional" json:"nodefsAvailable" yaml:"nodefsAvailable"`
	// Defines percentage of minimum reclaim for nodefs.inodesFree.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#nodefs_inodes_free GoogleContainerNodePool#nodefs_inodes_free}
	NodefsInodesFree *string `field:"optional" json:"nodefsInodesFree" yaml:"nodefsInodesFree"`
	// Defines percentage of minimum reclaim for pid.available.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#pid_available GoogleContainerNodePool#pid_available}
	PidAvailable *string `field:"optional" json:"pidAvailable" yaml:"pidAvailable"`
}

