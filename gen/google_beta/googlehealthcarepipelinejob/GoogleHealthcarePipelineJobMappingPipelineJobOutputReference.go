package googlehealthcarepipelinejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehealthcarepipelinejob/internal"
)

type GoogleHealthcarePipelineJobMappingPipelineJobOutputReference interface {
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
	FhirStoreDestination() *string
	SetFhirStoreDestination(val *string)
	FhirStoreDestinationInput() *string
	FhirStreamingSource() GoogleHealthcarePipelineJobMappingPipelineJobFhirStreamingSourceOutputReference
	FhirStreamingSourceInput() *GoogleHealthcarePipelineJobMappingPipelineJobFhirStreamingSource
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleHealthcarePipelineJobMappingPipelineJob
	SetInternalValue(val *GoogleHealthcarePipelineJobMappingPipelineJob)
	MappingConfig() GoogleHealthcarePipelineJobMappingPipelineJobMappingConfigOutputReference
	MappingConfigInput() *GoogleHealthcarePipelineJobMappingPipelineJobMappingConfig
	ReconciliationDestination() interface{}
	SetReconciliationDestination(val interface{})
	ReconciliationDestinationInput() interface{}
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
	PutFhirStreamingSource(value *GoogleHealthcarePipelineJobMappingPipelineJobFhirStreamingSource)
	PutMappingConfig(value *GoogleHealthcarePipelineJobMappingPipelineJobMappingConfig)
	ResetFhirStoreDestination()
	ResetFhirStreamingSource()
	ResetReconciliationDestination()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHealthcarePipelineJobMappingPipelineJobOutputReference
type jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) FhirStoreDestination() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fhirStoreDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) FhirStoreDestinationInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fhirStoreDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) FhirStreamingSource() GoogleHealthcarePipelineJobMappingPipelineJobFhirStreamingSourceOutputReference {
	var returns GoogleHealthcarePipelineJobMappingPipelineJobFhirStreamingSourceOutputReference
	_jsii_.Get(
		j,
		"fhirStreamingSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) FhirStreamingSourceInput() *GoogleHealthcarePipelineJobMappingPipelineJobFhirStreamingSource {
	var returns *GoogleHealthcarePipelineJobMappingPipelineJobFhirStreamingSource
	_jsii_.Get(
		j,
		"fhirStreamingSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) InternalValue() *GoogleHealthcarePipelineJobMappingPipelineJob {
	var returns *GoogleHealthcarePipelineJobMappingPipelineJob
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) MappingConfig() GoogleHealthcarePipelineJobMappingPipelineJobMappingConfigOutputReference {
	var returns GoogleHealthcarePipelineJobMappingPipelineJobMappingConfigOutputReference
	_jsii_.Get(
		j,
		"mappingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) MappingConfigInput() *GoogleHealthcarePipelineJobMappingPipelineJobMappingConfig {
	var returns *GoogleHealthcarePipelineJobMappingPipelineJobMappingConfig
	_jsii_.Get(
		j,
		"mappingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) ReconciliationDestination() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reconciliationDestination",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) ReconciliationDestinationInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"reconciliationDestinationInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleHealthcarePipelineJobMappingPipelineJobOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleHealthcarePipelineJobMappingPipelineJobOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHealthcarePipelineJobMappingPipelineJobOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHealthcarePipelineJob.GoogleHealthcarePipelineJobMappingPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleHealthcarePipelineJobMappingPipelineJobOutputReference_Override(g GoogleHealthcarePipelineJobMappingPipelineJobOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHealthcarePipelineJob.GoogleHealthcarePipelineJobMappingPipelineJobOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference)SetFhirStoreDestination(val *string) {
	if err := j.validateSetFhirStoreDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"fhirStoreDestination",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference)SetInternalValue(val *GoogleHealthcarePipelineJobMappingPipelineJob) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference)SetReconciliationDestination(val interface{}) {
	if err := j.validateSetReconciliationDestinationParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"reconciliationDestination",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) PutFhirStreamingSource(value *GoogleHealthcarePipelineJobMappingPipelineJobFhirStreamingSource) {
	if err := g.validatePutFhirStreamingSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putFhirStreamingSource",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) PutMappingConfig(value *GoogleHealthcarePipelineJobMappingPipelineJobMappingConfig) {
	if err := g.validatePutMappingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putMappingConfig",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) ResetFhirStoreDestination() {
	_jsii_.InvokeVoid(
		g,
		"resetFhirStoreDestination",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) ResetFhirStreamingSource() {
	_jsii_.InvokeVoid(
		g,
		"resetFhirStreamingSource",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) ResetReconciliationDestination() {
	_jsii_.InvokeVoid(
		g,
		"resetReconciliationDestination",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobMappingPipelineJobOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

