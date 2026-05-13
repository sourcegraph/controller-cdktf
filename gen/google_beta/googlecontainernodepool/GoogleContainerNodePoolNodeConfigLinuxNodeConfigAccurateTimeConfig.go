package googlecontainernodepool


type GoogleContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfig struct {
	// Whether to enable accurate time synchronization with PTP-KVM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_node_pool#enable_ptp_kvm_time_sync GoogleContainerNodePool#enable_ptp_kvm_time_sync}
	EnablePtpKvmTimeSync interface{} `field:"optional" json:"enablePtpKvmTimeSync" yaml:"enablePtpKvmTimeSync"`
}

