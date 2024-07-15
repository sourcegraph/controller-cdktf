package googlegkehubfeaturemembership


type GoogleGkeHubFeatureMembershipConfigmanagementPolicyController struct {
	// Sets the interval for Policy Controller Audit Scans (in seconds). When set to 0, this disables audit functionality altogether.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#audit_interval_seconds GoogleGkeHubFeatureMembership#audit_interval_seconds}
	AuditIntervalSeconds *string `field:"optional" json:"auditIntervalSeconds" yaml:"auditIntervalSeconds"`
	// Enables the installation of Policy Controller. If false, the rest of PolicyController fields take no effect.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#enabled GoogleGkeHubFeatureMembership#enabled}
	Enabled interface{} `field:"optional" json:"enabled" yaml:"enabled"`
	// The set of namespaces that are excluded from Policy Controller checks.
	//
	// Namespaces do not need to currently exist on the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#exemptable_namespaces GoogleGkeHubFeatureMembership#exemptable_namespaces}
	ExemptableNamespaces *[]*string `field:"optional" json:"exemptableNamespaces" yaml:"exemptableNamespaces"`
	// Logs all denies and dry run failures.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#log_denies_enabled GoogleGkeHubFeatureMembership#log_denies_enabled}
	LogDeniesEnabled interface{} `field:"optional" json:"logDeniesEnabled" yaml:"logDeniesEnabled"`
	// monitoring block.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#monitoring GoogleGkeHubFeatureMembership#monitoring}
	Monitoring *GoogleGkeHubFeatureMembershipConfigmanagementPolicyControllerMonitoring `field:"optional" json:"monitoring" yaml:"monitoring"`
	// Enable or disable mutation in policy controller.
	//
	// If true, mutation CRDs, webhook and controller deployment will be deployed to the cluster.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#mutation_enabled GoogleGkeHubFeatureMembership#mutation_enabled}
	MutationEnabled interface{} `field:"optional" json:"mutationEnabled" yaml:"mutationEnabled"`
	// Enables the ability to use Constraint Templates that reference to objects other than the object currently being evaluated.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#referential_rules_enabled GoogleGkeHubFeatureMembership#referential_rules_enabled}
	ReferentialRulesEnabled interface{} `field:"optional" json:"referentialRulesEnabled" yaml:"referentialRulesEnabled"`
	// Installs the default template library along with Policy Controller.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gke_hub_feature_membership#template_library_installed GoogleGkeHubFeatureMembership#template_library_installed}
	TemplateLibraryInstalled interface{} `field:"optional" json:"templateLibraryInstalled" yaml:"templateLibraryInstalled"`
}

