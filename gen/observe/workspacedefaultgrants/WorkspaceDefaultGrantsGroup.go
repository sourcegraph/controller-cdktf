package workspacedefaultgrants


type WorkspaceDefaultGrantsGroup struct {
	// The Observe ID for the group to grant access to.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/workspace_default_grants#oid WorkspaceDefaultGrants#oid}
	Oid *string `field:"required" json:"oid" yaml:"oid"`
	// The permission to grant. Must be one of `view` or `edit`.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/workspace_default_grants#permission WorkspaceDefaultGrants#permission}
	Permission *string `field:"required" json:"permission" yaml:"permission"`
	// Limits which object types this default grant applies to.
	//
	// Must be one of
	// `dashboard`, `datastream`, `monitor`, `referencetable`, or `worksheet`.
	// If not set, this default grant applies to all object types.
	// Note: Datasets are not represented here as they inherit default grants
	// based on their input datasets instead.
	//
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/observeinc/observe/0.14.47/docs/resources/workspace_default_grants#object_types WorkspaceDefaultGrants#object_types}
	ObjectTypes *[]*string `field:"optional" json:"objectTypes" yaml:"objectTypes"`
}

