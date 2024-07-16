package googleosconfigguestpolicies


type GoogleOsConfigGuestPoliciesAssignment struct {
	// group_labels block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#group_labels GoogleOsConfigGuestPolicies#group_labels}
	GroupLabels interface{} `field:"optional" json:"groupLabels" yaml:"groupLabels"`
	// Targets VM instances whose name starts with one of these prefixes.
	//
	// Like labels, this is another way to group VM instances when targeting configs,
	// for example prefix="prod-".
	// Only supported for project-level policies.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#instance_name_prefixes GoogleOsConfigGuestPolicies#instance_name_prefixes}
	InstanceNamePrefixes *[]*string `field:"optional" json:"instanceNamePrefixes" yaml:"instanceNamePrefixes"`
	// Targets any of the instances specified.
	//
	// Instances are specified by their URI in the form
	// zones/[ZONE]/instances/[INSTANCE_NAME].
	// Instance targeting is uncommon and is supported to facilitate the management of changes
	// by the instance or to target specific VM instances for development and testing.
	// Only supported for project-level policies and must reference instances within this project.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#instances GoogleOsConfigGuestPolicies#instances}
	Instances *[]*string `field:"optional" json:"instances" yaml:"instances"`
	// os_types block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#os_types GoogleOsConfigGuestPolicies#os_types}
	OsTypes interface{} `field:"optional" json:"osTypes" yaml:"osTypes"`
	// Targets instances in any of these zones.
	//
	// Leave empty to target instances in any zone.
	// Zonal targeting is uncommon and is supported to facilitate the management of changes by zone.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_os_config_guest_policies#zones GoogleOsConfigGuestPolicies#zones}
	Zones *[]*string `field:"optional" json:"zones" yaml:"zones"`
}

