package opensearchdomainsamloptions

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/aws/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/aws/opensearchdomainsamloptions/internal"
)

type OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference interface {
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
	EntityId() *string
	SetEntityId(val *string)
	EntityIdInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() *OpensearchDomainSamlOptionsSamlOptionsIdp
	SetInternalValue(val *OpensearchDomainSamlOptionsSamlOptionsIdp)
	MetadataContent() *string
	SetMetadataContent(val *string)
	MetadataContentInput() *string
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
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference
type jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) EntityId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) EntityIdInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityIdInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) InternalValue() *OpensearchDomainSamlOptionsSamlOptionsIdp {
	var returns *OpensearchDomainSamlOptionsSamlOptionsIdp
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) MetadataContent() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataContent",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) MetadataContentInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"metadataContentInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewOpensearchDomainSamlOptionsSamlOptionsIdpOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference {
	_init_.Initialize()

	if err := validateNewOpensearchDomainSamlOptionsSamlOptionsIdpOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomainSamlOptions.OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOpensearchDomainSamlOptionsSamlOptionsIdpOutputReference_Override(o OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-aws.opensearchDomainSamlOptions.OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference",
		[]any{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetEntityId(val *string) {
	if err := j.validateSetEntityIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityId",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetInternalValue(val *OpensearchDomainSamlOptionsSamlOptionsIdp) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetMetadataContent(val *string) {
	if err := j.validateSetMetadataContentParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"metadataContent",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		o,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OpensearchDomainSamlOptionsSamlOptionsIdpOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
