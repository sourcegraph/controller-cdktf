package googlelustreinstance


type GoogleLustreInstanceAccessRulesOptions struct {
	// The squash mode for the default access rule. Possible values: NO_SQUASH ROOT_SQUASH.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#default_squash_mode GoogleLustreInstance#default_squash_mode}
	DefaultSquashMode *string `field:"required" json:"defaultSquashMode" yaml:"defaultSquashMode"`
	// access_rules block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#access_rules GoogleLustreInstance#access_rules}
	AccessRules interface{} `field:"optional" json:"accessRules" yaml:"accessRules"`
	// The user squash GID for the default access rule.
	//
	// This user squash GID applies to all root users connecting from clients
	// that are not matched by any of the access rules. If not set, the default
	// is 0 (no GID squash).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#default_squash_gid GoogleLustreInstance#default_squash_gid}
	DefaultSquashGid *float64 `field:"optional" json:"defaultSquashGid" yaml:"defaultSquashGid"`
	// The user squash UID for the default access rule.
	//
	// This user squash UID applies to all root users connecting from clients
	// that are not matched by any of the access rules. If not set, the default
	// is 0 (no UID squash).
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_lustre_instance#default_squash_uid GoogleLustreInstance#default_squash_uid}
	DefaultSquashUid *float64 `field:"optional" json:"defaultSquashUid" yaml:"defaultSquashUid"`
}

