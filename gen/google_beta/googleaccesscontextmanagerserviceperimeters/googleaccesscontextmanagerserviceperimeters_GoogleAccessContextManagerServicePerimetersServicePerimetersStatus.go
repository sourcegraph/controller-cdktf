package googleaccesscontextmanagerserviceperimeters


type GoogleAccessContextManagerServicePerimetersServicePerimetersStatus struct {
	// A list of AccessLevel resource names that allow resources within the ServicePerimeter to be accessed from the internet.
	//
	// AccessLevels listed must be in the same policy as this
	// ServicePerimeter. Referencing a nonexistent AccessLevel is a
	// syntax error. If no AccessLevel names are listed, resources within
	// the perimeter can only be accessed via GCP calls with request
	// origins within the perimeter. For Service Perimeter Bridge, must
	// be empty.
	//
	// Format: accessPolicies/{policy_id}/accessLevels/{access_level_name}
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_service_perimeters#access_levels GoogleAccessContextManagerServicePerimeters#access_levels}
	AccessLevels *[]*string `field:"optional" json:"accessLevels" yaml:"accessLevels"`
	// egress_policies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_service_perimeters#egress_policies GoogleAccessContextManagerServicePerimeters#egress_policies}
	EgressPolicies interface{} `field:"optional" json:"egressPolicies" yaml:"egressPolicies"`
	// ingress_policies block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_service_perimeters#ingress_policies GoogleAccessContextManagerServicePerimeters#ingress_policies}
	IngressPolicies interface{} `field:"optional" json:"ingressPolicies" yaml:"ingressPolicies"`
	// A list of GCP resources that are inside of the service perimeter. Currently only projects are allowed. Format: projects/{project_number}.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_service_perimeters#resources GoogleAccessContextManagerServicePerimeters#resources}
	Resources *[]*string `field:"optional" json:"resources" yaml:"resources"`
	// GCP services that are subject to the Service Perimeter restrictions.
	//
	// Must contain a list of services. For example, if
	// 'storage.googleapis.com' is specified, access to the storage
	// buckets inside the perimeter must meet the perimeter's access
	// restrictions.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_service_perimeters#restricted_services GoogleAccessContextManagerServicePerimeters#restricted_services}
	RestrictedServices *[]*string `field:"optional" json:"restrictedServices" yaml:"restrictedServices"`
	// vpc_accessible_services block.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_access_context_manager_service_perimeters#vpc_accessible_services GoogleAccessContextManagerServicePerimeters#vpc_accessible_services}
	VpcAccessibleServices *GoogleAccessContextManagerServicePerimetersServicePerimetersStatusVpcAccessibleServices `field:"optional" json:"vpcAccessibleServices" yaml:"vpcAccessibleServices"`
}
