package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigEndpointConfig struct {
	// Optional. If true, enable http access to specific ports on the cluster from external sources. Defaults to false.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_workflow_template#enable_http_port_access GoogleDataprocWorkflowTemplate#enable_http_port_access}
	EnableHttpPortAccess interface{} `field:"optional" json:"enableHttpPortAccess" yaml:"enableHttpPortAccess"`
}

