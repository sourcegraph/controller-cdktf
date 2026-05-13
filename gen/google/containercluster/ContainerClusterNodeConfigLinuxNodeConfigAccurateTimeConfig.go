package containercluster


type ContainerClusterNodeConfigLinuxNodeConfigAccurateTimeConfig struct {
	// Whether to enable accurate time synchronization with PTP-KVM.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google/7.31.0/docs/resources/container_cluster#enable_ptp_kvm_time_sync ContainerCluster#enable_ptp_kvm_time_sync}
	EnablePtpKvmTimeSync interface{} `field:"optional" json:"enablePtpKvmTimeSync" yaml:"enablePtpKvmTimeSync"`
}

