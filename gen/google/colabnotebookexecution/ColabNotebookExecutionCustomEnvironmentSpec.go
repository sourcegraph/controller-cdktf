package colabnotebookexecution


type ColabNotebookExecutionCustomEnvironmentSpec struct {
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/colab_notebook_execution#machine_spec ColabNotebookExecution#machine_spec}
	MachineSpec *ColabNotebookExecutionCustomEnvironmentSpecMachineSpec `field:"optional" json:"machineSpec" yaml:"machineSpec"`
	// network_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/colab_notebook_execution#network_spec ColabNotebookExecution#network_spec}
	NetworkSpec *ColabNotebookExecutionCustomEnvironmentSpecNetworkSpec `field:"optional" json:"networkSpec" yaml:"networkSpec"`
	// persistent_disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/colab_notebook_execution#persistent_disk_spec ColabNotebookExecution#persistent_disk_spec}
	PersistentDiskSpec *ColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec `field:"optional" json:"persistentDiskSpec" yaml:"persistentDiskSpec"`
}

