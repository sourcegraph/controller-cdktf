package googlerediscluster


type GoogleRedisClusterTimeouts struct {
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_redis_cluster#create GoogleRedisCluster#create}.
	Create *string `field:"optional" json:"create" yaml:"create"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_redis_cluster#delete GoogleRedisCluster#delete}.
	Delete *string `field:"optional" json:"delete" yaml:"delete"`
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/5.38.0/docs/resources/google_redis_cluster#update GoogleRedisCluster#update}.
	Update *string `field:"optional" json:"update" yaml:"update"`
}

