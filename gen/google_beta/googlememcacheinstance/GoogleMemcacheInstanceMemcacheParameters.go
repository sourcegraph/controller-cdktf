package googlememcacheinstance


type GoogleMemcacheInstanceMemcacheParameters struct {
	// User-defined set of parameters to use in the memcache process.
	//
	// Docs at Terraform Registry: {@link https://registry.terraform.io/providers/hashicorp/google-beta/4.65.2/docs/resources/google_memcache_instance#params GoogleMemcacheInstance#params}
	Params *map[string]*string `field:"optional" json:"params" yaml:"params"`
}

