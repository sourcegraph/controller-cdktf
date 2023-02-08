package apitoken


type ApiTokenPolicy struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_token#permission_groups ApiToken#permission_groups}.
	PermissionGroups *[]*string `field:"required" json:"permissionGroups" yaml:"permissionGroups"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_token#resources ApiToken#resources}.
	Resources *map[string]*string `field:"required" json:"resources" yaml:"resources"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/cloudflare/r/api_token#effect ApiToken#effect}.
	Effect *string `field:"optional" json:"effect" yaml:"effect"`
}

