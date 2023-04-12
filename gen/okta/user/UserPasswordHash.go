package user


type UserPasswordHash struct {
	// The algorithm used to generate the hash using the password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user#algorithm User#algorithm}
	Algorithm *string `field:"required" json:"algorithm" yaml:"algorithm"`
	// For SHA-512, SHA-256, SHA-1, MD5, This is the actual base64-encoded hash of the password (and salt, if used).
	//
	// This is the Base64 encoded value of the SHA-512/SHA-256/SHA-1/MD5 digest that was computed by either pre-fixing or post-fixing the salt to the password, depending on the saltOrder. If a salt was not used in the source system, then this should just be the the Base64 encoded value of the password's SHA-512/SHA-256/SHA-1/MD5 digest. For BCRYPT, This is the actual radix64-encoded hashed password.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user#value User#value}
	Value *string `field:"required" json:"value" yaml:"value"`
	// Only required for salted hashes.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user#salt User#salt}
	Salt *string `field:"optional" json:"salt" yaml:"salt"`
	// Specifies whether salt was pre- or postfixed to the password before hashing.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user#salt_order User#salt_order}
	SaltOrder *string `field:"optional" json:"saltOrder" yaml:"saltOrder"`
	// Governs the strength of the hash and the time required to compute it. Only required for BCRYPT algorithm.
	//
	// Docs at Terraform Registry: {@link https://www.terraform.io/docs/providers/okta/r/user#work_factor User#work_factor}
	WorkFactor *float64 `field:"optional" json:"workFactor" yaml:"workFactor"`
}

