package googlememcacheinstance


type GoogleMemcacheInstanceMemcacheParameters struct {
	// User-defined set of parameters to use in the memcache process.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/google-beta/r/google_memcache_instance#params GoogleMemcacheInstance#params}
	Params *map[string]*string `field:"optional" json:"params" yaml:"params"`
}

