package googleeventarctrigger

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleeventarctrigger/internal"
)

type GoogleEventarcTriggerDestinationOutputReference interface {
	cdktf.ComplexObject
	CloudFunction() *string
	SetCloudFunction(val *string)
	CloudFunctionInput() *string
	CloudRunService() GoogleEventarcTriggerDestinationCloudRunServiceOutputReference
	CloudRunServiceInput() *GoogleEventarcTriggerDestinationCloudRunService
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
	Gke() GoogleEventarcTriggerDestinationGkeOutputReference
	GkeInput() *GoogleEventarcTriggerDestinationGke
	InternalValue() *GoogleEventarcTriggerDestination
	SetInternalValue(val *GoogleEventarcTriggerDestination)
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Workflow() *string
	SetWorkflow(val *string)
	WorkflowInput() *string
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
	PutCloudRunService(value *GoogleEventarcTriggerDestinationCloudRunService)
	PutGke(value *GoogleEventarcTriggerDestinationGke)
	ResetCloudFunction()
	ResetCloudRunService()
	ResetGke()
	ResetWorkflow()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleEventarcTriggerDestinationOutputReference
type jsiiProxy_GoogleEventarcTriggerDestinationOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) CloudFunction() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFunction",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) CloudFunctionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"cloudFunctionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) CloudRunService() GoogleEventarcTriggerDestinationCloudRunServiceOutputReference {
	var returns GoogleEventarcTriggerDestinationCloudRunServiceOutputReference
	_jsii_.Get(
		j,
		"cloudRunService",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) CloudRunServiceInput() *GoogleEventarcTriggerDestinationCloudRunService {
	var returns *GoogleEventarcTriggerDestinationCloudRunService
	_jsii_.Get(
		j,
		"cloudRunServiceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) Gke() GoogleEventarcTriggerDestinationGkeOutputReference {
	var returns GoogleEventarcTriggerDestinationGkeOutputReference
	_jsii_.Get(
		j,
		"gke",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) GkeInput() *GoogleEventarcTriggerDestinationGke {
	var returns *GoogleEventarcTriggerDestinationGke
	_jsii_.Get(
		j,
		"gkeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) InternalValue() *GoogleEventarcTriggerDestination {
	var returns *GoogleEventarcTriggerDestination
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) Workflow() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflow",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) WorkflowInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"workflowInput",
		&returns,
	)
	return returns
}


func NewGoogleEventarcTriggerDestinationOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleEventarcTriggerDestinationOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleEventarcTriggerDestinationOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleEventarcTriggerDestinationOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEventarcTrigger.GoogleEventarcTriggerDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleEventarcTriggerDestinationOutputReference_Override(g GoogleEventarcTriggerDestinationOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleEventarcTrigger.GoogleEventarcTriggerDestinationOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference)SetCloudFunction(val *string) {
	if err := j.validateSetCloudFunctionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"cloudFunction",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference)SetInternalValue(val *GoogleEventarcTriggerDestination) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference)SetWorkflow(val *string) {
	if err := j.validateSetWorkflowParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"workflow",
		val,
	)
}

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) PutCloudRunService(value *GoogleEventarcTriggerDestinationCloudRunService) {
	if err := g.validatePutCloudRunServiceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putCloudRunService",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) PutGke(value *GoogleEventarcTriggerDestinationGke) {
	if err := g.validatePutGkeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putGke",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) ResetCloudFunction() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudFunction",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) ResetCloudRunService() {
	_jsii_.InvokeVoid(
		g,
		"resetCloudRunService",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) ResetGke() {
	_jsii_.InvokeVoid(
		g,
		"resetGke",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) ResetWorkflow() {
	_jsii_.InvokeVoid(
		g,
		"resetWorkflow",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleEventarcTriggerDestinationOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

