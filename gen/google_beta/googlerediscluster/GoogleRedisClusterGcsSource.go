package googlerediscluster


type GoogleRedisClusterGcsSource struct {
	// URIs of the GCS objects to import. Example: gs://bucket1/object1, gs://bucket2/folder2/object2.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/6.30.0/docs/resources/google_redis_cluster#uris GoogleRedisCluster#uris}
	Uris *[]*string `field:"required" json:"uris" yaml:"uris"`
}

