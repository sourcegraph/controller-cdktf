package pubsubsubscription

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/pubsubsubscription/internal"
)

type PubsubSubscriptionBigqueryConfigOutputReference interface {
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
	DropUnknownFields() any
	SetDropUnknownFields(val any)
	DropUnknownFieldsInput() any
	// Experimental.
	Fqn() *string
	InternalValue() *PubsubSubscriptionBigqueryConfig
	SetInternalValue(val *PubsubSubscriptionBigqueryConfig)
	ServiceAccountEmail() *string
	SetServiceAccountEmail(val *string)
	ServiceAccountEmailInput() *string
	Table() *string
	SetTable(val *string)
	TableInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	UseTableSchema() any
	SetUseTableSchema(val any)
	UseTableSchemaInput() any
	UseTopicSchema() any
	SetUseTopicSchema(val any)
	UseTopicSchemaInput() any
	WriteMetadata() any
	SetWriteMetadata(val any)
	WriteMetadataInput() any
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
	ResetDropUnknownFields()
	ResetServiceAccountEmail()
	ResetUseTableSchema()
	ResetUseTopicSchema()
	ResetWriteMetadata()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for PubsubSubscriptionBigqueryConfigOutputReference
type jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) DropUnknownFields() any {
	var returns any
	_jsii_.Get(
		j,
		"dropUnknownFields",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) DropUnknownFieldsInput() any {
	var returns any
	_jsii_.Get(
		j,
		"dropUnknownFieldsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) InternalValue() *PubsubSubscriptionBigqueryConfig {
	var returns *PubsubSubscriptionBigqueryConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ServiceAccountEmail() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmail",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ServiceAccountEmailInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"serviceAccountEmailInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) Table() *string {
	var returns *string
	_jsii_.Get(
		j,
		"table",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) TableInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"tableInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) UseTableSchema() any {
	var returns any
	_jsii_.Get(
		j,
		"useTableSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) UseTableSchemaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useTableSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) UseTopicSchema() any {
	var returns any
	_jsii_.Get(
		j,
		"useTopicSchema",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) UseTopicSchemaInput() any {
	var returns any
	_jsii_.Get(
		j,
		"useTopicSchemaInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) WriteMetadata() any {
	var returns any
	_jsii_.Get(
		j,
		"writeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) WriteMetadataInput() any {
	var returns any
	_jsii_.Get(
		j,
		"writeMetadataInput",
		&returns,
	)
	return returns
}

func NewPubsubSubscriptionBigqueryConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) PubsubSubscriptionBigqueryConfigOutputReference {
	_init_.Initialize()

	if err := validateNewPubsubSubscriptionBigqueryConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.pubsubSubscription.PubsubSubscriptionBigqueryConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewPubsubSubscriptionBigqueryConfigOutputReference_Override(p PubsubSubscriptionBigqueryConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.pubsubSubscription.PubsubSubscriptionBigqueryConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		p,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetDropUnknownFields(val any) {
	if err := j.validateSetDropUnknownFieldsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"dropUnknownFields",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetInternalValue(val *PubsubSubscriptionBigqueryConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetServiceAccountEmail(val *string) {
	if err := j.validateSetServiceAccountEmailParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"serviceAccountEmail",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetTable(val *string) {
	if err := j.validateSetTableParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"table",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetUseTableSchema(val any) {
	if err := j.validateSetUseTableSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTableSchema",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetUseTopicSchema(val any) {
	if err := j.validateSetUseTopicSchemaParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"useTopicSchema",
		val,
	)
}

func (j *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) SetWriteMetadata(val any) {
	if err := j.validateSetWriteMetadataParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"writeMetadata",
		val,
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := p.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		p,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := p.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := p.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		p,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := p.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		p,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := p.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		p,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := p.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		p,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := p.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		p,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := p.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		p,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := p.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		p,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := p.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		p,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ResetDropUnknownFields() {
	_jsii_.InvokeVoid(
		p,
		"resetDropUnknownFields",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ResetServiceAccountEmail() {
	_jsii_.InvokeVoid(
		p,
		"resetServiceAccountEmail",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ResetUseTableSchema() {
	_jsii_.InvokeVoid(
		p,
		"resetUseTableSchema",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ResetUseTopicSchema() {
	_jsii_.InvokeVoid(
		p,
		"resetUseTopicSchema",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ResetWriteMetadata() {
	_jsii_.InvokeVoid(
		p,
		"resetWriteMetadata",
		nil, // no parameters
	)
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := p.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		p,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (p *jsiiProxy_PubsubSubscriptionBigqueryConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		p,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
