package cloudfrontresponseheaderspolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/cloudfrontresponseheaderspolicy/internal"
)

type CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference interface {
	cdktf.ComplexObject
	AccessControlMaxAgeSec() *float64
	SetAccessControlMaxAgeSec(val *float64)
	AccessControlMaxAgeSecInput() *float64
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() any
	// Experimental.
	SetComplexObjectIndex(val any)
	// set to true if this item is from inside a set and needs tolist() for accessing it set to "0" for single list items.
	// Experimental.
	ComplexObjectIsFromSet() *bool
	// Experimental.
	SetComplexObjectIsFromSet(val *bool)
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	IncludeSubdomains() any
	SetIncludeSubdomains(val any)
	IncludeSubdomainsInput() any
	InternalValue() *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity
	SetInternalValue(val *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity)
	Override() any
	SetOverride(val any)
	OverrideInput() any
	Preload() any
	SetPreload(val any)
	PreloadInput() any
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]any
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable
	// Experimental.
	GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool
	// Experimental.
	GetListAttribute(terraformAttribute *string) *[]*string
	// Experimental.
	GetNumberAttribute(terraformAttribute *string) *float64
	// Experimental.
	GetNumberListAttribute(terraformAttribute *string) *[]*float64
	// Experimental.
	GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64
	// Experimental.
	GetStringAttribute(terraformAttribute *string) *string
	// Experimental.
	GetStringMapAttribute(terraformAttribute *string) *map[string]*string
	// Experimental.
	InterpolationAsList() cdktf.IResolvable
	// Experimental.
	InterpolationForAttribute(property *string) cdktf.IResolvable
	ResetIncludeSubdomains()
	ResetPreload()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference
type jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) AccessControlMaxAgeSec() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSec",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) AccessControlMaxAgeSecInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"accessControlMaxAgeSecInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) IncludeSubdomains() any {
	var returns any
	_jsii_.Get(
		j,
		"includeSubdomains",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) IncludeSubdomainsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeSubdomainsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) InternalValue() *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity {
	var returns *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) Override() any {
	var returns any
	_jsii_.Get(
		j,
		"override",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) OverrideInput() any {
	var returns any
	_jsii_.Get(
		j,
		"overrideInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) Preload() any {
	var returns any
	_jsii_.Get(
		j,
		"preload",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) PreloadInput() any {
	var returns any
	_jsii_.Get(
		j,
		"preloadInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference {
	_init_.Initialize()

	if err := validateNewCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference_Override(c CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudfrontResponseHeadersPolicy.CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetAccessControlMaxAgeSec(val *float64) {
	if err := j.validateSetAccessControlMaxAgeSecParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"accessControlMaxAgeSec",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetIncludeSubdomains(val any) {
	if err := j.validateSetIncludeSubdomainsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeSubdomains",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetInternalValue(val *CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetOverride(val any) {
	if err := j.validateSetOverrideParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"override",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetPreload(val any) {
	if err := j.validateSetPreloadParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"preload",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) ResetIncludeSubdomains() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeSubdomains",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) ResetPreload() {
	_jsii_.InvokeVoid(
		c,
		"resetPreload",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		c,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudfrontResponseHeadersPolicySecurityHeadersConfigStrictTransportSecurityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
