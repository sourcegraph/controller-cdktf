package appgroupassignments


type AppGroupAssignmentsGroup struct {
	// A group to associate with the application.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignments#id AppGroupAssignments#id}
	//
	// Please be aware that the id field is automatically added to all resources in Terraform providers using a Terraform provider SDK version below 2.
	// If you experience problems setting this value it might not be settable. Please take a look at the provider documentation to ensure it should be settable.
	Id *string `field:"required" json:"id" yaml:"id"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignments#priority AppGroupAssignments#priority}.
	Priority *float64 `field:"optional" json:"priority" yaml:"priority"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/app_group_assignments#profile AppGroupAssignments#profile}.
	Profile *string `field:"optional" json:"profile" yaml:"profile"`
}

