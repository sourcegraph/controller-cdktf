package googlestoragebucket

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestoragebucket/internal"
)

type GoogleStorageBucketIpFilterOutputReference interface {
	cdktf.ComplexObject
	AllowAllServiceAgentAccess() interface{}
	SetAllowAllServiceAgentAccess(val interface{})
	AllowAllServiceAgentAccessInput() interface{}
	AllowCrossOrgVpcs() interface{}
	SetAllowCrossOrgVpcs(val interface{})
	AllowCrossOrgVpcsInput() interface{}
	// the index of the complex object in a list.
	// Experimental.
	ComplexObjectIndex() interface{}
	// Experimental.
	SetComplexObjectIndex(val interface{})
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
	InternalValue() *GoogleStorageBucketIpFilter
	SetInternalValue(val *GoogleStorageBucketIpFilter)
	Mode() *string
	SetMode(val *string)
	ModeInput() *string
	PublicNetworkSource() GoogleStorageBucketIpFilterPublicNetworkSourceOutputReference
	PublicNetworkSourceInput() *GoogleStorageBucketIpFilterPublicNetworkSource
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	VpcNetworkSources() GoogleStorageBucketIpFilterVpcNetworkSourcesList
	VpcNetworkSourcesInput() interface{}
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
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
	PutPublicNetworkSource(value *GoogleStorageBucketIpFilterPublicNetworkSource)
	PutVpcNetworkSources(value interface{})
	ResetAllowAllServiceAgentAccess()
	ResetAllowCrossOrgVpcs()
	ResetPublicNetworkSource()
	ResetVpcNetworkSources()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageBucketIpFilterOutputReference
type jsiiProxy_GoogleStorageBucketIpFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) AllowAllServiceAgentAccess() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllServiceAgentAccess",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) AllowAllServiceAgentAccessInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowAllServiceAgentAccessInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) AllowCrossOrgVpcs() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCrossOrgVpcs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) AllowCrossOrgVpcsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"allowCrossOrgVpcsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) InternalValue() *GoogleStorageBucketIpFilter {
	var returns *GoogleStorageBucketIpFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) Mode() *string {
	var returns *string
	_jsii_.Get(
		j,
		"mode",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) ModeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) PublicNetworkSource() GoogleStorageBucketIpFilterPublicNetworkSourceOutputReference {
	var returns GoogleStorageBucketIpFilterPublicNetworkSourceOutputReference
	_jsii_.Get(
		j,
		"publicNetworkSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) PublicNetworkSourceInput() *GoogleStorageBucketIpFilterPublicNetworkSource {
	var returns *GoogleStorageBucketIpFilterPublicNetworkSource
	_jsii_.Get(
		j,
		"publicNetworkSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) VpcNetworkSources() GoogleStorageBucketIpFilterVpcNetworkSourcesList {
	var returns GoogleStorageBucketIpFilterVpcNetworkSourcesList
	_jsii_.Get(
		j,
		"vpcNetworkSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) VpcNetworkSourcesInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"vpcNetworkSourcesInput",
		&returns,
	)
	return returns
}


func NewGoogleStorageBucketIpFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageBucketIpFilterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageBucketIpFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageBucketIpFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageBucket.GoogleStorageBucketIpFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageBucketIpFilterOutputReference_Override(g GoogleStorageBucketIpFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageBucket.GoogleStorageBucketIpFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference)SetAllowAllServiceAgentAccess(val interface{}) {
	if err := j.validateSetAllowAllServiceAgentAccessParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowAllServiceAgentAccess",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference)SetAllowCrossOrgVpcs(val interface{}) {
	if err := j.validateSetAllowCrossOrgVpcsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"allowCrossOrgVpcs",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference)SetInternalValue(val *GoogleStorageBucketIpFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference)SetMode(val *string) {
	if err := j.validateSetModeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"mode",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageBucketIpFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := g.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		g,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := g.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		g,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := g.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		g,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := g.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		g,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := g.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		g,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := g.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		g,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := g.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		g,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := g.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		g,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) PutPublicNetworkSource(value *GoogleStorageBucketIpFilterPublicNetworkSource) {
	if err := g.validatePutPublicNetworkSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPublicNetworkSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) PutVpcNetworkSources(value interface{}) {
	if err := g.validatePutVpcNetworkSourcesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putVpcNetworkSources",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) ResetAllowAllServiceAgentAccess() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowAllServiceAgentAccess",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) ResetAllowCrossOrgVpcs() {
	_jsii_.InvokeVoid(
		g,
		"resetAllowCrossOrgVpcs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) ResetPublicNetworkSource() {
	_jsii_.InvokeVoid(
		g,
		"resetPublicNetworkSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) ResetVpcNetworkSources() {
	_jsii_.InvokeVoid(
		g,
		"resetVpcNetworkSources",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := g.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageBucketIpFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

