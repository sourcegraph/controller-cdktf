package securesourcemanagerhook


type SecureSourceManagerHookPushOption struct {
	// Trigger hook for matching branches only.
	//
	// Specified as glob pattern. If empty or *, events for all branches are
	// reported. Examples: main, {main,release*}.
	// See https://pkg.go.dev/github.com/gobwas/glob documentation.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/secure_source_manager_hook#branch_filter SecureSourceManagerHook#branch_filter}
	BranchFilter *string `field:"optional" json:"branchFilter" yaml:"branchFilter"`
}

