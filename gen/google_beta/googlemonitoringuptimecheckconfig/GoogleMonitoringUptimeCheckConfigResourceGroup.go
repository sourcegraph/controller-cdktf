package googlemonitoringuptimecheckconfig


type GoogleMonitoringUptimeCheckConfigResourceGroup struct {
	// The group of resources being monitored. Should be the 'name' of a group.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_uptime_check_config#group_id GoogleMonitoringUptimeCheckConfig#group_id}
	GroupId *string `field:"optional" json:"groupId" yaml:"groupId"`
	// The resource type of the group members. Possible values: ["RESOURCE_TYPE_UNSPECIFIED", "INSTANCE", "AWS_ELB_LOAD_BALANCER"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_uptime_check_config#resource_type GoogleMonitoringUptimeCheckConfig#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

