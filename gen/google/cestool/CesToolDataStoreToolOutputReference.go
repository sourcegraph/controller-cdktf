package cestool

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/cestool/internal"
)

type CesToolDataStoreToolOutputReference interface {
	cdktf.ComplexObject
	BoostSpecs() CesToolDataStoreToolBoostSpecsList
	BoostSpecsInput() any
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
	Description() *string
	SetDescription(val *string)
	DescriptionInput() *string
	EngineSource() CesToolDataStoreToolEngineSourceOutputReference
	EngineSourceInput() *CesToolDataStoreToolEngineSource
	// Experimental.
	Fqn() *string
	InternalValue() *CesToolDataStoreTool
	SetInternalValue(val *CesToolDataStoreTool)
	MaxResults() *float64
	SetMaxResults(val *float64)
	MaxResultsInput() *float64
	ModalityConfigs() CesToolDataStoreToolModalityConfigsList
	ModalityConfigsInput() any
	Name() *string
	SetName(val *string)
	NameInput() *string
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
	PutBoostSpecs(value any)
	PutEngineSource(value *CesToolDataStoreToolEngineSource)
	PutModalityConfigs(value any)
	ResetBoostSpecs()
	ResetDescription()
	ResetEngineSource()
	ResetMaxResults()
	ResetModalityConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CesToolDataStoreToolOutputReference
type jsiiProxy_CesToolDataStoreToolOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) BoostSpecs() CesToolDataStoreToolBoostSpecsList {
	var returns CesToolDataStoreToolBoostSpecsList
	_jsii_.Get(
		j,
		"boostSpecs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) BoostSpecsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"boostSpecsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) Description() *string {
	var returns *string
	_jsii_.Get(
		j,
		"description",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) DescriptionInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"descriptionInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) EngineSource() CesToolDataStoreToolEngineSourceOutputReference {
	var returns CesToolDataStoreToolEngineSourceOutputReference
	_jsii_.Get(
		j,
		"engineSource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) EngineSourceInput() *CesToolDataStoreToolEngineSource {
	var returns *CesToolDataStoreToolEngineSource
	_jsii_.Get(
		j,
		"engineSourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) InternalValue() *CesToolDataStoreTool {
	var returns *CesToolDataStoreTool
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) MaxResults() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResults",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) MaxResultsInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxResultsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) ModalityConfigs() CesToolDataStoreToolModalityConfigsList {
	var returns CesToolDataStoreToolModalityConfigsList
	_jsii_.Get(
		j,
		"modalityConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) ModalityConfigsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"modalityConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) Name() *string {
	var returns *string
	_jsii_.Get(
		j,
		"name",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) NameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"nameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCesToolDataStoreToolOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) CesToolDataStoreToolOutputReference {
	_init_.Initialize()

	if err := validateNewCesToolDataStoreToolOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_CesToolDataStoreToolOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.cesTool.CesToolDataStoreToolOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewCesToolDataStoreToolOutputReference_Override(c CesToolDataStoreToolOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.cesTool.CesToolDataStoreToolOutputReference",
		[]any{terraformResource, terraformAttribute},
		c,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) SetDescription(val *string) {
	if err := j.validateSetDescriptionParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"description",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) SetInternalValue(val *CesToolDataStoreTool) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) SetMaxResults(val *float64) {
	if err := j.validateSetMaxResultsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxResults",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) SetName(val *string) {
	if err := j.validateSetNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"name",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CesToolDataStoreToolOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) PutBoostSpecs(value any) {
	if err := c.validatePutBoostSpecsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putBoostSpecs",
		[]any{value},
	)
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) PutEngineSource(value *CesToolDataStoreToolEngineSource) {
	if err := c.validatePutEngineSourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putEngineSource",
		[]any{value},
	)
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) PutModalityConfigs(value any) {
	if err := c.validatePutModalityConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putModalityConfigs",
		[]any{value},
	)
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) ResetBoostSpecs() {
	_jsii_.InvokeVoid(
		c,
		"resetBoostSpecs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) ResetDescription() {
	_jsii_.InvokeVoid(
		c,
		"resetDescription",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) ResetEngineSource() {
	_jsii_.InvokeVoid(
		c,
		"resetEngineSource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) ResetMaxResults() {
	_jsii_.InvokeVoid(
		c,
		"resetMaxResults",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) ResetModalityConfigs() {
	_jsii_.InvokeVoid(
		c,
		"resetModalityConfigs",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CesToolDataStoreToolOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
