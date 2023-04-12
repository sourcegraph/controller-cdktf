package teamaccess


type TeamAccessPermissions struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/team_access#runs TeamAccess#runs}.
	Runs *string `field:"required" json:"runs" yaml:"runs"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/team_access#run_tasks TeamAccess#run_tasks}.
	RunTasks interface{} `field:"required" json:"runTasks" yaml:"runTasks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/team_access#sentinel_mocks TeamAccess#sentinel_mocks}.
	SentinelMocks *string `field:"required" json:"sentinelMocks" yaml:"sentinelMocks"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/team_access#state_versions TeamAccess#state_versions}.
	StateVersions *string `field:"required" json:"stateVersions" yaml:"stateVersions"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/team_access#variables TeamAccess#variables}.
	Variables *string `field:"required" json:"variables" yaml:"variables"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/tfe/r/team_access#workspace_locking TeamAccess#workspace_locking}.
	WorkspaceLocking interface{} `field:"required" json:"workspaceLocking" yaml:"workspaceLocking"`
}

