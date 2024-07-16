package googlecontainercluster


type GoogleContainerClusterProtectConfigWorkloadConfig struct {
	// Sets which mode of auditing should be used for the cluster's workloads. Accepted values are DISABLED, BASIC.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_container_cluster#audit_mode GoogleContainerCluster#audit_mode}
	AuditMode *string `field:"required" json:"auditMode" yaml:"auditMode"`
}

