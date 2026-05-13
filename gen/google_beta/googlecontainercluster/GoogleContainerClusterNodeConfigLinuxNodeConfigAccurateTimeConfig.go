package googlecontainercluster


type GoogleContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig struct {
	// Whether to enable accurate time synchronization with PTP-KVM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/7.31.0/docs/resources/google_container_cluster#enable_ptp_kvm_time_sync GoogleContainerCluster#enable_ptp_kvm_time_sync}
	EnablePtpKvmTimeSync interface{} `field:"optional" json:"enablePtpKvmTimeSync" yaml:"enablePtpKvmTimeSync"`
}

