package pagesproject


type PagesProjectSourceConfig struct {
	// Project production branch name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#production_branch PagesProject#production_branch}
	ProductionBranch *string `field:"required" json:"productionBranch" yaml:"productionBranch"`
	// Toggle deployments on this repo. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#deployments_enabled PagesProject#deployments_enabled}
	DeploymentsEnabled interface{} `field:"optional" json:"deploymentsEnabled" yaml:"deploymentsEnabled"`
	// Project owner username.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#owner PagesProject#owner}
	Owner *string `field:"optional" json:"owner" yaml:"owner"`
	// Enable Pages to comment on Pull Requests. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#pr_comments_enabled PagesProject#pr_comments_enabled}
	PrCommentsEnabled interface{} `field:"optional" json:"prCommentsEnabled" yaml:"prCommentsEnabled"`
	// Branches will be excluded from automatic deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#preview_branch_excludes PagesProject#preview_branch_excludes}
	PreviewBranchExcludes *[]*string `field:"optional" json:"previewBranchExcludes" yaml:"previewBranchExcludes"`
	// Branches will be included for automatic deployment.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#preview_branch_includes PagesProject#preview_branch_includes}
	PreviewBranchIncludes *[]*string `field:"optional" json:"previewBranchIncludes" yaml:"previewBranchIncludes"`
	// Preview Deployment Setting. Defaults to `all`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#preview_deployment_setting PagesProject#preview_deployment_setting}
	PreviewDeploymentSetting *string `field:"optional" json:"previewDeploymentSetting" yaml:"previewDeploymentSetting"`
	// Enable production deployments. Defaults to `true`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#production_deployment_enabled PagesProject#production_deployment_enabled}
	ProductionDeploymentEnabled interface{} `field:"optional" json:"productionDeploymentEnabled" yaml:"productionDeploymentEnabled"`
	// Project repository name.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/cloudflare/cloudflare/4.3.0/docs/resources/pages_project#repo_name PagesProject#repo_name}
	RepoName *string `field:"optional" json:"repoName" yaml:"repoName"`
}

