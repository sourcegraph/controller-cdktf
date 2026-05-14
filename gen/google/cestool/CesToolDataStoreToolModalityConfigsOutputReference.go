package cestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestool/internal"
)

type CesToolDataStoreToolModalityConfigsOutputReference interface {
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
	GroundingConfig() CesToolDataStoreToolModalityConfigsGroundingConfigOutputReference
	GroundingConfigInput() *CesToolDataStoreToolModalityConfigsGroundingConfig
	InternalValue() interface{}
	SetInternalValue(val interface{})
	ModalityType() *string
	SetModalityType(val *string)
	ModalityTypeInput() *string
	RewriterConfig() CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference
	RewriterConfigInput() *CesToolDataStoreToolModalityConfigsRewriterConfig
	SummarizationConfig() CesToolDataStoreToolModalityConfigsSummarizationConfigOutputReference
	SummarizationConfigInput() *CesToolDataStoreToolModalityConfigsSummarizationConfig
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
	PutGroundingConfig(value *CesToolDataStoreToolModalityConfigsGroundingConfig)
	PutRewriterConfig(value *CesToolDataStoreToolModalityConfigsRewriterConfig)
	PutSummarizationConfig(value *CesToolDataStoreToolModalityConfigsSummarizationConfig)
	ResetGroundingConfig()
	ResetRewriterConfig()
	ResetSummarizationConfig()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolDataStoreToolModalityConfigsOutputReference
type jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GroundingConfig() CesToolDataStoreToolModalityConfigsGroundingConfigOutputReference {
	var returns CesToolDataStoreToolModalityConfigsGroundingConfigOutputReference
	_jsii_.Get(
		j,
		"groundingConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GroundingConfigInput() *CesToolDataStoreToolModalityConfigsGroundingConfig {
	var returns *CesToolDataStoreToolModalityConfigsGroundingConfig
	_jsii_.Get(
		j,
		"groundingConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) InternalValue() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) ModalityType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modalityType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) ModalityTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"modalityTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) RewriterConfig() CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference {
	var returns CesToolDataStoreToolModalityConfigsRewriterConfigOutputReference
	_jsii_.Get(
		j,
		"rewriterConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) RewriterConfigInput() *CesToolDataStoreToolModalityConfigsRewriterConfig {
	var returns *CesToolDataStoreToolModalityConfigsRewriterConfig
	_jsii_.Get(
		j,
		"rewriterConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) SummarizationConfig() CesToolDataStoreToolModalityConfigsSummarizationConfigOutputReference {
	var returns CesToolDataStoreToolModalityConfigsSummarizationConfigOutputReference
	_jsii_.Get(
		j,
		"summarizationConfig",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) SummarizationConfigInput() *CesToolDataStoreToolModalityConfigsSummarizationConfig {
	var returns *CesToolDataStoreToolModalityConfigsSummarizationConfig
	_jsii_.Get(
		j,
		"summarizationConfigInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewCesToolDataStoreToolModalityConfigsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CesToolDataStoreToolModalityConfigsOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolDataStoreToolModalityConfigsOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesTool.CesToolDataStoreToolModalityConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCesToolDataStoreToolModalityConfigsOutputReference_Override(c CesToolDataStoreToolModalityConfigsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesTool.CesToolDataStoreToolModalityConfigsOutputReference",
		[]interface{}{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference)SetInternalValue(val interface{}) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference)SetModalityType(val *string) {
	if err := j.validateSetModalityTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"modalityType",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := c.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		c,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := c.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := c.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		c,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := c.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		c,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := c.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		c,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := c.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		c,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := c.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		c,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := c.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		c,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := c.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		c,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := c.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) PutGroundingConfig(value *CesToolDataStoreToolModalityConfigsGroundingConfig) {
	if err := c.validatePutGroundingConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putGroundingConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) PutRewriterConfig(value *CesToolDataStoreToolModalityConfigsRewriterConfig) {
	if err := c.validatePutRewriterConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putRewriterConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) PutSummarizationConfig(value *CesToolDataStoreToolModalityConfigsSummarizationConfig) {
	if err := c.validatePutSummarizationConfigParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putSummarizationConfig",
		[]interface{}{value},
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) ResetGroundingConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetGroundingConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) ResetRewriterConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetRewriterConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) ResetSummarizationConfig() {
	_jsii_.InvokeVoid(
		c,
		"resetSummarizationConfig",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := c.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		c,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolModalityConfigsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

