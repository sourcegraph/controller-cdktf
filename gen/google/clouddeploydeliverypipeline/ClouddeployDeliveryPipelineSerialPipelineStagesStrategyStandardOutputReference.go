package clouddeploydeliverypipeline

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/clouddeploydeliverypipeline/internal"
)

type ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference interface {
	cdktf.ComplexObject
	Analysis() ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisOutputReference
	AnalysisInput() *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysis
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
	InternalValue() *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard
	SetInternalValue(val *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard)
	Postdeploy() ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployOutputReference
	PostdeployInput() *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy
	Predeploy() ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployOutputReference
	PredeployInput() *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Verify() any
	SetVerify(val any)
	VerifyConfig() ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardVerifyConfigOutputReference
	VerifyConfigInput() *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardVerifyConfig
	VerifyInput() any
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
	PutAnalysis(value *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysis)
	PutPostdeploy(value *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy)
	PutPredeploy(value *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy)
	PutVerifyConfig(value *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardVerifyConfig)
	ResetAnalysis()
	ResetPostdeploy()
	ResetPredeploy()
	ResetVerify()
	ResetVerifyConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference
type jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) Analysis() ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisOutputReference {
	var returns ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysisOutputReference
	_jsii_.Get(
		j,
		"analysis",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) AnalysisInput() *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysis {
	var returns *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysis
	_jsii_.Get(
		j,
		"analysisInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) InternalValue() *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard {
	var returns *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) Postdeploy() ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployOutputReference {
	var returns ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeployOutputReference
	_jsii_.Get(
		j,
		"postdeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) PostdeployInput() *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy {
	var returns *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy
	_jsii_.Get(
		j,
		"postdeployInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) Predeploy() ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployOutputReference {
	var returns ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeployOutputReference
	_jsii_.Get(
		j,
		"predeploy",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) PredeployInput() *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy {
	var returns *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy
	_jsii_.Get(
		j,
		"predeployInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) Verify() any {
	var returns any
	_jsii_.Get(
		j,
		"verify",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) VerifyConfig() ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardVerifyConfigOutputReference {
	var returns ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardVerifyConfigOutputReference
	_jsii_.Get(
		j,
		"verifyConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) VerifyConfigInput() *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardVerifyConfig {
	var returns *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardVerifyConfig
	_jsii_.Get(
		j,
		"verifyConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) VerifyInput() any {
	var returns any
	_jsii_.Get(
		j,
		"verifyInput",
		&returns,
	)
	return returns
}

func NewClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference {
	_init_.Initialize()

	if err := validateNewClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployDeliveryPipeline.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference_Override(c ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.clouddeployDeliveryPipeline.ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) SetInternalValue(val *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandard) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) SetVerify(val any) {
	if err := j.validateSetVerifyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"verify",
		val,
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) PutAnalysis(value *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardAnalysis) {
	if err := c.validatePutAnalysisParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putAnalysis",
		[]any{value},
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) PutPostdeploy(value *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPostdeploy) {
	if err := c.validatePutPostdeployParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPostdeploy",
		[]any{value},
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) PutPredeploy(value *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardPredeploy) {
	if err := c.validatePutPredeployParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putPredeploy",
		[]any{value},
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) PutVerifyConfig(value *ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardVerifyConfig) {
	if err := c.validatePutVerifyConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putVerifyConfig",
		[]any{value},
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) ResetAnalysis() {
	_jsii_.InvokeVoid(
		c,
		"resetAnalysis",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) ResetPostdeploy() {
	_jsii_.InvokeVoid(
		c,
		"resetPostdeploy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) ResetPredeploy() {
	_jsii_.InvokeVoid(
		c,
		"resetPredeploy",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) ResetVerify() {
	_jsii_.InvokeVoid(
		c,
		"resetVerify",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) ResetVerifyConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetVerifyConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_ClouddeployDeliveryPipelineSerialPipelineStagesStrategyStandardOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
