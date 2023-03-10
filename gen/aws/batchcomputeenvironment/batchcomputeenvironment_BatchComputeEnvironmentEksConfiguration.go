package batchcomputeenvironment


type BatchComputeEnvironmentEksConfiguration struct {
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#eks_cluster_arn BatchComputeEnvironment#eks_cluster_arn}.
	EksClusterArn *string `field:"required" json:"eksClusterArn" yaml:"eksClusterArn"`
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/aws/r/batch_compute_environment#kubernetes_namespace BatchComputeEnvironment#kubernetes_namespace}.
	KubernetesNamespace *string `field:"required" json:"kubernetesNamespace" yaml:"kubernetesNamespace"`
}

