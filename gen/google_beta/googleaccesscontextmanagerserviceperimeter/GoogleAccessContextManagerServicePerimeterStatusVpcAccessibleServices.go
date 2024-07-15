package googleaccesscontextmanagerserviceperimeter


type GoogleAccessContextManagerServicePerimeterStatusVpcAccessibleServices struct {
	// The list of APIs usable within the Service Perimeter. Must be empty unless 'enableRestriction' is True.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter#allowed_services GoogleAccessContextManagerServicePerimeter#allowed_services}
	AllowedServices *[]*string `field:"optional" json:"allowedServices" yaml:"allowedServices"`
	// Whether to restrict API calls within the Service Perimeter to the list of APIs specified in 'allowedServices'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_access_context_manager_service_perimeter#enable_restriction GoogleAccessContextManagerServicePerimeter#enable_restriction}
	EnableRestriction interface{} `field:"optional" json:"enableRestriction" yaml:"enableRestriction"`
}

