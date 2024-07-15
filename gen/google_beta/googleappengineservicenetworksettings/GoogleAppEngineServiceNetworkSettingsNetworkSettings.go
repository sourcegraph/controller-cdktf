package googleappengineservicenetworksettings


type GoogleAppEngineServiceNetworkSettingsNetworkSettings struct {
	// The ingress settings for version or service. Default value: "INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED" Possible values: ["INGRESS_TRAFFIC_ALLOWED_UNSPECIFIED", "INGRESS_TRAFFIC_ALLOWED_ALL", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_ONLY", "INGRESS_TRAFFIC_ALLOWED_INTERNAL_AND_LB"].
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_app_engine_service_network_settings#ingress_traffic_allowed GoogleAppEngineServiceNetworkSettings#ingress_traffic_allowed}
	IngressTrafficAllowed *string `field:"optional" json:"ingressTrafficAllowed" yaml:"ingressTrafficAllowed"`
}

