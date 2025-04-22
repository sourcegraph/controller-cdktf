package googlehealthcarepipelinejob

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehealthcarepipelinejob/internal"
)

type GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference interface {
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
	ImportUriPrefix() *string
	SetImportUriPrefix(val *string)
	ImportUriPrefixInput() *string
	InternalValue() *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource
	SetInternalValue(val *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Uri() *string
	SetUri(val *string)
	UriInput() *string
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

// The jsii proxy struct for GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference
type jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ImportUriPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUriPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ImportUriPrefixInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"importUriPrefixInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) InternalValue() *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource {
	var returns *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) Uri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) UriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"uriInput",
		&returns,
	)
	return returns
}


func NewGoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHealthcarePipelineJob.GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference_Override(g GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleHealthcarePipelineJob.GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetImportUriPrefix(val *string) {
	if err := j.validateSetImportUriPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"importUriPrefix",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetInternalValue(val *GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSource) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference)SetUri(val *string) {
	if err := j.validateSetUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"uri",
		val,
	)
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHealthcarePipelineJobReconciliationPipelineJobMergeConfigWhistleConfigSourceOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

