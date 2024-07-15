package datagoogleiampolicy


type DataGoogleIamPolicyAuditConfig struct {
	// audit_log_configs block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/data-sources/google_iam_policy#audit_log_configs DataGoogleIamPolicy#audit_log_configs}
	AuditLogConfigs interface{} `field:"required" json:"auditLogConfigs" yaml:"auditLogConfigs"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/data-sources/google_iam_policy#service DataGoogleIamPolicy#service}.
	Service *string `field:"required" json:"service" yaml:"service"`
}

