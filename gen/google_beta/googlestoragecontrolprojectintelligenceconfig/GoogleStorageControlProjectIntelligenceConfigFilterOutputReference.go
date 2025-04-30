package googlestoragecontrolprojectintelligenceconfig

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlestoragecontrolprojectintelligenceconfig/internal"
)

type GoogleStorageControlProjectIntelligenceConfigFilterOutputReference interface {
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
	ExcludedCloudStorageBuckets() GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBucketsOutputReference
	ExcludedCloudStorageBucketsInput() *GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets
	ExcludedCloudStorageLocations() GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocationsOutputReference
	ExcludedCloudStorageLocationsInput() *GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations
	// Experimental.
	Fqn() *string
	IncludedCloudStorageBuckets() GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBucketsOutputReference
	IncludedCloudStorageBucketsInput() *GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets
	IncludedCloudStorageLocations() GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocationsOutputReference
	IncludedCloudStorageLocationsInput() *GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations
	InternalValue() *GoogleStorageControlProjectIntelligenceConfigFilter
	SetInternalValue(val *GoogleStorageControlProjectIntelligenceConfigFilter)
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
	PutExcludedCloudStorageBuckets(value *GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets)
	PutExcludedCloudStorageLocations(value *GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations)
	PutIncludedCloudStorageBuckets(value *GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets)
	PutIncludedCloudStorageLocations(value *GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations)
	ResetExcludedCloudStorageBuckets()
	ResetExcludedCloudStorageLocations()
	ResetIncludedCloudStorageBuckets()
	ResetIncludedCloudStorageLocations()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleStorageControlProjectIntelligenceConfigFilterOutputReference
type jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ExcludedCloudStorageBuckets() GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBucketsOutputReference {
	var returns GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBucketsOutputReference
	_jsii_.Get(
		j,
		"excludedCloudStorageBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ExcludedCloudStorageBucketsInput() *GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets {
	var returns *GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets
	_jsii_.Get(
		j,
		"excludedCloudStorageBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ExcludedCloudStorageLocations() GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocationsOutputReference {
	var returns GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocationsOutputReference
	_jsii_.Get(
		j,
		"excludedCloudStorageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ExcludedCloudStorageLocationsInput() *GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations {
	var returns *GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations
	_jsii_.Get(
		j,
		"excludedCloudStorageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) IncludedCloudStorageBuckets() GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBucketsOutputReference {
	var returns GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBucketsOutputReference
	_jsii_.Get(
		j,
		"includedCloudStorageBuckets",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) IncludedCloudStorageBucketsInput() *GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets {
	var returns *GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets
	_jsii_.Get(
		j,
		"includedCloudStorageBucketsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) IncludedCloudStorageLocations() GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocationsOutputReference {
	var returns GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocationsOutputReference
	_jsii_.Get(
		j,
		"includedCloudStorageLocations",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) IncludedCloudStorageLocationsInput() *GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations {
	var returns *GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations
	_jsii_.Get(
		j,
		"includedCloudStorageLocationsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) InternalValue() *GoogleStorageControlProjectIntelligenceConfigFilter {
	var returns *GoogleStorageControlProjectIntelligenceConfigFilter
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleStorageControlProjectIntelligenceConfigFilterOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleStorageControlProjectIntelligenceConfigFilterOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleStorageControlProjectIntelligenceConfigFilterOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageControlProjectIntelligenceConfig.GoogleStorageControlProjectIntelligenceConfigFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleStorageControlProjectIntelligenceConfigFilterOutputReference_Override(g GoogleStorageControlProjectIntelligenceConfigFilterOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleStorageControlProjectIntelligenceConfig.GoogleStorageControlProjectIntelligenceConfigFilterOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference)SetInternalValue(val *GoogleStorageControlProjectIntelligenceConfigFilter) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) PutExcludedCloudStorageBuckets(value *GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageBuckets) {
	if err := g.validatePutExcludedCloudStorageBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludedCloudStorageBuckets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) PutExcludedCloudStorageLocations(value *GoogleStorageControlProjectIntelligenceConfigFilterExcludedCloudStorageLocations) {
	if err := g.validatePutExcludedCloudStorageLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putExcludedCloudStorageLocations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) PutIncludedCloudStorageBuckets(value *GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageBuckets) {
	if err := g.validatePutIncludedCloudStorageBucketsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIncludedCloudStorageBuckets",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) PutIncludedCloudStorageLocations(value *GoogleStorageControlProjectIntelligenceConfigFilterIncludedCloudStorageLocations) {
	if err := g.validatePutIncludedCloudStorageLocationsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putIncludedCloudStorageLocations",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ResetExcludedCloudStorageBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedCloudStorageBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ResetExcludedCloudStorageLocations() {
	_jsii_.InvokeVoid(
		g,
		"resetExcludedCloudStorageLocations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ResetIncludedCloudStorageBuckets() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludedCloudStorageBuckets",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ResetIncludedCloudStorageLocations() {
	_jsii_.InvokeVoid(
		g,
		"resetIncludedCloudStorageLocations",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleStorageControlProjectIntelligenceConfigFilterOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

