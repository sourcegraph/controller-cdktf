package googlecolabnotebookexecution


type GoogleColabNotebookExecutionCustomEnvironmentSpec struct {
	// machine_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#machine_spec GoogleColabNotebookExecution#machine_spec}
	MachineSpec *GoogleColabNotebookExecutionCustomEnvironmentSpecMachineSpec `field:"optional" json:"machineSpec" yaml:"machineSpec"`
	// network_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#network_spec GoogleColabNotebookExecution#network_spec}
	NetworkSpec *GoogleColabNotebookExecutionCustomEnvironmentSpecNetworkSpec `field:"optional" json:"networkSpec" yaml:"networkSpec"`
	// persistent_disk_spec block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_colab_notebook_execution#persistent_disk_spec GoogleColabNotebookExecution#persistent_disk_spec}
	PersistentDiskSpec *GoogleColabNotebookExecutionCustomEnvironmentSpecPersistentDiskSpec `field:"optional" json:"persistentDiskSpec" yaml:"persistentDiskSpec"`
}

