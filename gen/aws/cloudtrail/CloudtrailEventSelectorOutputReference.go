package cloudtrail

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/cloudtrail/internal"
)

type CloudtrailEventSelectorOutputReference interface {
	cdktf.ComplexObject
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
	DataResource() CloudtrailEventSelectorDataResourceList
	DataResourceInput() any
	ExcludeManagementEventSources() *[]*string
	SetExcludeManagementEventSources(val *[]*string)
	ExcludeManagementEventSourcesInput() *[]*string
	// Experimental.
	Fqn() *string
	IncludeManagementEvents() any
	SetIncludeManagementEvents(val any)
	IncludeManagementEventsInput() any
	InternalValue() any
	SetInternalValue(val any)
	ReadWriteType() *string
	SetReadWriteType(val *string)
	ReadWriteTypeInput() *string
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
	PutDataResource(value any)
	ResetDataResource()
	ResetExcludeManagementEventSources()
	ResetIncludeManagementEvents()
	ResetReadWriteType()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for CloudtrailEventSelectorOutputReference
type jsiiProxy_CloudtrailEventSelectorOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) DataResource() CloudtrailEventSelectorDataResourceList {
	var returns CloudtrailEventSelectorDataResourceList
	_jsii_.Get(
		j,
		"dataResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) DataResourceInput() any {
	var returns any
	_jsii_.Get(
		j,
		"dataResourceInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) ExcludeManagementEventSources() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeManagementEventSources",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) ExcludeManagementEventSourcesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"excludeManagementEventSourcesInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) IncludeManagementEvents() any {
	var returns any
	_jsii_.Get(
		j,
		"includeManagementEvents",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) IncludeManagementEventsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"includeManagementEventsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) ReadWriteType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readWriteType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) ReadWriteTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"readWriteTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewCloudtrailEventSelectorOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) CloudtrailEventSelectorOutputReference {
	_init_.Initialize()

	if err := validateNewCloudtrailEventSelectorOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_CloudtrailEventSelectorOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.cloudtrail.CloudtrailEventSelectorOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewCloudtrailEventSelectorOutputReference_Override(c CloudtrailEventSelectorOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.cloudtrail.CloudtrailEventSelectorOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		c,
	)
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) SetExcludeManagementEventSources(val *[]*string) {
	if err := j.validateSetExcludeManagementEventSourcesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"excludeManagementEventSources",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) SetIncludeManagementEvents(val any) {
	if err := j.validateSetIncludeManagementEventsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"includeManagementEvents",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) SetReadWriteType(val *string) {
	if err := j.validateSetReadWriteTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"readWriteType",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_CloudtrailEventSelectorOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		c,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) PutDataResource(value any) {
	if err := c.validatePutDataResourceParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		c,
		"putDataResource",
		[]any{value},
	)
}

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) ResetDataResource() {
	_jsii_.InvokeVoid(
		c,
		"resetDataResource",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) ResetExcludeManagementEventSources() {
	_jsii_.InvokeVoid(
		c,
		"resetExcludeManagementEventSources",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) ResetIncludeManagementEvents() {
	_jsii_.InvokeVoid(
		c,
		"resetIncludeManagementEvents",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) ResetReadWriteType() {
	_jsii_.InvokeVoid(
		c,
		"resetReadWriteType",
		nil, // no parameters
	)
}

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) Resolve(_context cdktf.IResolveContext) any {
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

func (c *jsiiProxy_CloudtrailEventSelectorOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		c,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
