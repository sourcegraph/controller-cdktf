package googlegkeonprembaremetalcluster


type GoogleGkeonpremBareMetalClusterBinaryAuthorization struct {
	// Mode of operation for binauthz policy evaluation. If unspecified, defaults to DISABLED. Possible values: ["DISABLED", "PROJECT_SINGLETON_POLICY_ENFORCE"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_gkeonprem_bare_metal_cluster#evaluation_mode GoogleGkeonpremBareMetalCluster#evaluation_mode}
	EvaluationMode *string `field:"optional" json:"evaluationMode" yaml:"evaluationMode"`
}

