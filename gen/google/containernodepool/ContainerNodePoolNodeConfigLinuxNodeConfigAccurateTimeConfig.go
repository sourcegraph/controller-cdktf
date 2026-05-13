package containernodepool


type ContainerNodePoolNodeConfigLinuxNodeConfigAccurateTimeConfig struct {
	// Whether to enable accurate time synchronization with PTP-KVM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_node_pool#enable_ptp_kvm_time_sync ContainerNodePool#enable_ptp_kvm_time_sync}
	EnablePtpKvmTimeSync interface{} `field:"optional" json:"enablePtpKvmTimeSync" yaml:"enablePtpKvmTimeSync"`
}

