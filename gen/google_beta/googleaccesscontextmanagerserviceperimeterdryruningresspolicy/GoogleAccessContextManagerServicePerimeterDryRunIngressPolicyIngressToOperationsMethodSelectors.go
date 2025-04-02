package googleaccesscontextmanagerserviceperimeterdryruningresspolicy


type GoogleAccessContextManagerServicePerimeterDryRunIngressPolicyIngressToOperationsMethodSelectors struct {
	// Value for method should be a valid method name for the corresponding serviceName in 'ApiOperation'.
	//
	// If '*' used as value for 'method', then
	// ALL methods and permissions are allowed.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_access_context_manager_service_perimeter_dry_run_ingress_policy#method GoogleAccessContextManagerServicePerimeterDryRunIngressPolicy#method}
	Method *string `field:"optional" json:"method" yaml:"method"`
	// Value for permission should be a valid Cloud IAM permission for the corresponding 'serviceName' in 'ApiOperation'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.45.0/docs/resources/google_access_context_manager_service_perimeter_dry_run_ingress_policy#permission GoogleAccessContextManagerServicePerimeterDryRunIngressPolicy#permission}
	Permission *string `field:"optional" json:"permission" yaml:"permission"`
}

