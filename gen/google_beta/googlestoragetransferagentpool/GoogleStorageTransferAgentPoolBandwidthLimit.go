package googlestoragetransferagentpool


type GoogleStorageTransferAgentPoolBandwidthLimit struct {
	// Bandwidth rate in megabytes per second, distributed across all the agents in the pool.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_storage_transfer_agent_pool#limit_mbps GoogleStorageTransferAgentPool#limit_mbps}
	LimitMbps *string `field:"required" json:"limitMbps" yaml:"limitMbps"`
}

