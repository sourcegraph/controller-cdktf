package googlehypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehypercomputeclustercluster/internal"
)

type GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleHypercomputeclusterClusterComputeResourcesConfig
	SetInternalValue(val *GoogleHypercomputeclusterClusterComputeResourcesConfig)
	NewFlexStartInstances() GoogleHypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstancesOutputReference
	NewFlexStartInstancesInput() *GoogleHypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances
	NewOnDemandInstances() GoogleHypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstancesOutputReference
	NewOnDemandInstancesInput() *GoogleHypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances
	NewReservedInstances() GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstancesOutputReference
	NewReservedInstancesInput() *GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstances
	NewSpotInstances() GoogleHypercomputeclusterClusterComputeResourcesConfigNewSpotInstancesOutputReference
	NewSpotInstancesInput() *GoogleHypercomputeclusterClusterComputeResourcesConfigNewSpotInstances
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	PutNewFlexStartInstances(value *GoogleHypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances)
	PutNewOnDemandInstances(value *GoogleHypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances)
	PutNewReservedInstances(value *GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstances)
	PutNewSpotInstances(value *GoogleHypercomputeclusterClusterComputeResourcesConfigNewSpotInstances)
	ResetNewFlexStartInstances()
	ResetNewOnDemandInstances()
	ResetNewReservedInstances()
	ResetNewSpotInstances()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference
type jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) InternalValue() *GoogleHypercomputeclusterClusterComputeResourcesConfig {
	var returns *GoogleHypercomputeclusterClusterComputeResourcesConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) NewFlexStartInstances() GoogleHypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstancesOutputReference {
	var returns GoogleHypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstancesOutputReference
	_jsii_.Get(
		j,
		"newFlexStartInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) NewFlexStartInstancesInput() *GoogleHypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances {
	var returns *GoogleHypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances
	_jsii_.Get(
		j,
		"newFlexStartInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) NewOnDemandInstances() GoogleHypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstancesOutputReference {
	var returns GoogleHypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstancesOutputReference
	_jsii_.Get(
		j,
		"newOnDemandInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) NewOnDemandInstancesInput() *GoogleHypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances {
	var returns *GoogleHypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances
	_jsii_.Get(
		j,
		"newOnDemandInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) NewReservedInstances() GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstancesOutputReference {
	var returns GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstancesOutputReference
	_jsii_.Get(
		j,
		"newReservedInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) NewReservedInstancesInput() *GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstances {
	var returns *GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstances
	_jsii_.Get(
		j,
		"newReservedInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) NewSpotInstances() GoogleHypercomputeclusterClusterComputeResourcesConfigNewSpotInstancesOutputReference {
	var returns GoogleHypercomputeclusterClusterComputeResourcesConfigNewSpotInstancesOutputReference
	_jsii_.Get(
		j,
		"newSpotInstances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) NewSpotInstancesInput() *GoogleHypercomputeclusterClusterComputeResourcesConfigNewSpotInstances {
	var returns *GoogleHypercomputeclusterClusterComputeResourcesConfigNewSpotInstances
	_jsii_.Get(
		j,
		"newSpotInstancesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHypercomputeclusterClusterComputeResourcesConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference_Override(g GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference)SetInternalValue(val *GoogleHypercomputeclusterClusterComputeResourcesConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) PutNewFlexStartInstances(value *GoogleHypercomputeclusterClusterComputeResourcesConfigNewFlexStartInstances) {
	if err := g.validatePutNewFlexStartInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNewFlexStartInstances",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) PutNewOnDemandInstances(value *GoogleHypercomputeclusterClusterComputeResourcesConfigNewOnDemandInstances) {
	if err := g.validatePutNewOnDemandInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNewOnDemandInstances",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) PutNewReservedInstances(value *GoogleHypercomputeclusterClusterComputeResourcesConfigNewReservedInstances) {
	if err := g.validatePutNewReservedInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNewReservedInstances",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) PutNewSpotInstances(value *GoogleHypercomputeclusterClusterComputeResourcesConfigNewSpotInstances) {
	if err := g.validatePutNewSpotInstancesParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putNewSpotInstances",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) ResetNewFlexStartInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetNewFlexStartInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) ResetNewOnDemandInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetNewOnDemandInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) ResetNewReservedInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetNewReservedInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) ResetNewSpotInstances() {
	_jsii_.InvokeVoid(
		g,
		"resetNewSpotInstances",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterComputeResourcesConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

