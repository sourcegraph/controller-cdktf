package googlenetworkconnectivitymulticlouddatatransferconfig


type GoogleNetworkConnectivityMulticloudDataTransferConfigServices struct {
	// The name of the service, like "big-query" or "cloud-storage". This corresponds to the map key in the API.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_network_connectivity_multicloud_data_transfer_config#service_name GoogleNetworkConnectivityMulticloudDataTransferConfig#service_name}
	ServiceName *string `field:"required" json:"serviceName" yaml:"serviceName"`
}

