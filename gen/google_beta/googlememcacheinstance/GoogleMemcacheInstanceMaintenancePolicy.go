package googlememcacheinstance


type GoogleMemcacheInstanceMaintenancePolicy struct {
	// weekly_maintenance_window block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_memcache_instance#weekly_maintenance_window GoogleMemcacheInstance#weekly_maintenance_window}
	WeeklyMaintenanceWindow interface{} `field:"required" json:"weeklyMaintenanceWindow" yaml:"weeklyMaintenanceWindow"`
	// Optional. Description of what this policy is for. Create/Update methods return INVALID_ARGUMENT if the length is greater than 512.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_memcache_instance#description GoogleMemcacheInstance#description}
	Description *string `field:"optional" json:"description" yaml:"description"`
}

