package googlebackupdrmanagementserver


type GoogleBackupDrManagementServerNetworks struct {
	// Network with format 'projects/{{project_id}}/global/networks/{{network_id}}'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_backup_dr_management_server#network GoogleBackupDrManagementServer#network}
	Network *string `field:"required" json:"network" yaml:"network"`
	// Type of Network peeringMode Default value: "PRIVATE_SERVICE_ACCESS" Possible values: ["PRIVATE_SERVICE_ACCESS"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_backup_dr_management_server#peering_mode GoogleBackupDrManagementServer#peering_mode}
	PeeringMode *string `field:"optional" json:"peeringMode" yaml:"peeringMode"`
}

