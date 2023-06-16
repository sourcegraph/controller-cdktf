package googleassuredworkloadsworkload


type GoogleAssuredWorkloadsWorkloadResourceSettings struct {
	// Resource identifier. For a project this represents project_number. If the project is already taken, the workload creation will fail.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_assured_workloads_workload#resource_id GoogleAssuredWorkloadsWorkload#resource_id}
	ResourceId *string `field:"optional" json:"resourceId" yaml:"resourceId"`
	// Indicates the type of resource.
	//
	// This field should be specified to correspond the id to the right project type (CONSUMER_PROJECT or ENCRYPTION_KEYS_PROJECT) Possible values: RESOURCE_TYPE_UNSPECIFIED, CONSUMER_PROJECT, ENCRYPTION_KEYS_PROJECT, KEYRING, CONSUMER_FOLDER
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.69.1/docs/resources/google_assured_workloads_workload#resource_type GoogleAssuredWorkloadsWorkload#resource_type}
	ResourceType *string `field:"optional" json:"resourceType" yaml:"resourceType"`
}

