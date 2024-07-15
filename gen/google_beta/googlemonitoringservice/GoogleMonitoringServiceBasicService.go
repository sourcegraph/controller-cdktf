package googlemonitoringservice


type GoogleMonitoringServiceBasicService struct {
	// Labels that specify the resource that emits the monitoring data which is used for SLO reporting of this 'Service'.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_service#service_labels GoogleMonitoringService#service_labels}
	ServiceLabels *map[string]*string `field:"optional" json:"serviceLabels" yaml:"serviceLabels"`
	// The type of service that this basic service defines, e.g. APP_ENGINE service type.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_monitoring_service#service_type GoogleMonitoringService#service_type}
	ServiceType *string `field:"optional" json:"serviceType" yaml:"serviceType"`
}

