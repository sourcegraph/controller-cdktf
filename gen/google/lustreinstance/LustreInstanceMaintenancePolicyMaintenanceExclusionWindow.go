package lustreinstance


type LustreInstanceMaintenancePolicyMaintenanceExclusionWindow struct {
	// end_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#end_date LustreInstance#end_date}
	EndDate *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowEndDate `field:"required" json:"endDate" yaml:"endDate"`
	// start_date block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#start_date LustreInstance#start_date}
	StartDate *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowStartDate `field:"required" json:"startDate" yaml:"startDate"`
	// time block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/lustre_instance#time LustreInstance#time}
	Time *LustreInstanceMaintenancePolicyMaintenanceExclusionWindowTime `field:"required" json:"time" yaml:"time"`
}

