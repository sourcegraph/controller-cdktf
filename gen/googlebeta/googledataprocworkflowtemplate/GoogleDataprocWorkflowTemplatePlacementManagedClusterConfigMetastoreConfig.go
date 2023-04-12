package googledataprocworkflowtemplate


type GoogleDataprocWorkflowTemplatePlacementManagedClusterConfigMetastoreConfig struct {
	// Required. Resource name of an existing Dataproc Metastore service. Example: * `projects/[project_id]/locations/[dataproc_region]/services/[service-name]`.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_dataproc_workflow_template#dataproc_metastore_service GoogleDataprocWorkflowTemplate#dataproc_metastore_service}
	DataprocMetastoreService *string `field:"required" json:"dataprocMetastoreService" yaml:"dataprocMetastoreService"`
}

