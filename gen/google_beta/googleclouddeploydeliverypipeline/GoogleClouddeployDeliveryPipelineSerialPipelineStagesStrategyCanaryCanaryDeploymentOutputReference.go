package googleclouddeploydeliverypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleclouddeploydeliverypipeline/internal"
)

type GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference interface {
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
	InternalValue() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment
	SetInternalValue(val *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment)
	Percentages() *[]*float64
	SetPercentages(val *[]*float64)
	PercentagesInput() *[]*float64
	Postdeploy() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployOutputReference
	PostdeployInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy
	Predeploy() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployOutputReference
	PredeployInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Verify() interface{}
	SetVerify(val interface{})
	VerifyInput() interface{}
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
	PutPostdeploy(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy)
	PutPredeploy(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy)
	ResetPostdeploy()
	ResetPredeploy()
	ResetVerify()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference
type jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) InternalValue() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment {
	var returns *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) Percentages() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"percentages",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) PercentagesInput() *[]*float64 {
	var returns *[]*float64
	_jsii_.Get(
		j,
		"percentagesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) Postdeploy() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployOutputReference {
	var returns GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeployOutputReference
	_jsii_.Get(
		j,
		"postdeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) PostdeployInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy {
	var returns *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy
	_jsii_.Get(
		j,
		"postdeployInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) Predeploy() GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployOutputReference {
	var returns GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeployOutputReference
	_jsii_.Get(
		j,
		"predeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) PredeployInput() *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy {
	var returns *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy
	_jsii_.Get(
		j,
		"predeployInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) Verify() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) VerifyInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"verifyInput",
		&returns,
	)
	return returns
}


func NewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployDeliveryPipeline.GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference_Override(g GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleClouddeployDeliveryPipeline.GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference)SetInternalValue(val *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeployment) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference)SetPercentages(val *[]*float64) {
	if err := j.validateSetPercentagesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"percentages",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference)SetVerify(val interface{}) {
	if err := j.validateSetVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verify",
		val,
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) PutPostdeploy(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPostdeploy) {
	if err := g.validatePutPostdeployParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPostdeploy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) PutPredeploy(value *GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentPredeploy) {
	if err := g.validatePutPredeployParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putPredeploy",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) ResetPostdeploy() {
	_jsii_.InvokeVoid(
		g,
		"resetPostdeploy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) ResetPredeploy() {
	_jsii_.InvokeVoid(
		g,
		"resetPredeploy",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) ResetVerify() {
	_jsii_.InvokeVoid(
		g,
		"resetVerify",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleClouddeployDeliveryPipelineSerialPipelineStagesStrategyCanaryCanaryDeploymentOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

