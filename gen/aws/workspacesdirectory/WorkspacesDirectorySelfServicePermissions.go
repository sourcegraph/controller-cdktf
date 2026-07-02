package workspacesdirectory

type WorkspacesDirectorySelfServicePermissions struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/workspaces_directory#change_compute_type WorkspacesDirectory#change_compute_type}.
	ChangeComputeType any `field:"optional" json:"changeComputeType" yaml:"changeComputeType"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/workspaces_directory#increase_volume_size WorkspacesDirectory#increase_volume_size}.
	IncreaseVolumeSize any `field:"optional" json:"increaseVolumeSize" yaml:"increaseVolumeSize"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/workspaces_directory#rebuild_workspace WorkspacesDirectory#rebuild_workspace}.
	RebuildWorkspace any `field:"optional" json:"rebuildWorkspace" yaml:"rebuildWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/workspaces_directory#restart_workspace WorkspacesDirectory#restart_workspace}.
	RestartWorkspace any `field:"optional" json:"restartWorkspace" yaml:"restartWorkspace"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/aws/4.54.0/docs/resources/workspaces_directory#switch_running_mode WorkspacesDirectory#switch_running_mode}.
	SwitchRunningMode any `field:"optional" json:"switchRunningMode" yaml:"switchRunningMode"`
}
