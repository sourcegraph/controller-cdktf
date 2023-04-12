package googlecomputeregionurlmap

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/googlebeta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/googlebeta/googlecomputeregionurlmap/internal"
)

type GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference interface {
	cdktf.ComplexObject
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
	InternalValue() *GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch
	SetInternalValue(val *GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch)
	RangeEnd() *float64
	SetRangeEnd(val *float64)
	RangeEndInput() *float64
	RangeStart() *float64
	SetRangeStart(val *float64)
	RangeStartInput() *float64
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference
type jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) InternalValue() *GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch {
	var returns *GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) RangeEnd() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rangeEnd",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) RangeEndInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rangeEndInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) RangeStart() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rangeStart",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) RangeStartInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"rangeStartInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference_Override(g GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-googlebeta.googleComputeRegionUrlMap.GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference)SetInternalValue(val *GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatch) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference)SetRangeEnd(val *float64) {
	if err := j.validateSetRangeEndParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeEnd",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference)SetRangeStart(val *float64) {
	if err := j.validateSetRangeStartParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"rangeStart",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeRegionUrlMapPathMatcherRouteRulesMatchRulesHeaderMatchesRangeMatchOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

