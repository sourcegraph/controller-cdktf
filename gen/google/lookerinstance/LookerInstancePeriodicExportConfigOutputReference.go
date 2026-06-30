package lookerinstance

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/lookerinstance/internal"
)

type LookerInstancePeriodicExportConfigOutputReference interface {
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
	// Experimental.
	Fqn() *string
	GcsUri() *string
	SetGcsUri(val *string)
	GcsUriInput() *string
	InternalValue() *LookerInstancePeriodicExportConfig
	SetInternalValue(val *LookerInstancePeriodicExportConfig)
	KmsKey() *string
	SetKmsKey(val *string)
	KmsKeyInput() *string
	StartTime() LookerInstancePeriodicExportConfigStartTimeOutputReference
	StartTimeInput() *LookerInstancePeriodicExportConfigStartTime
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
	PutStartTime(value *LookerInstancePeriodicExportConfigStartTime)
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for LookerInstancePeriodicExportConfigOutputReference
type jsiiProxy_LookerInstancePeriodicExportConfigOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GcsUri() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsUri",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GcsUriInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"gcsUriInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) InternalValue() *LookerInstancePeriodicExportConfig {
	var returns *LookerInstancePeriodicExportConfig
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) KmsKey() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKey",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) KmsKeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"kmsKeyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) StartTime() LookerInstancePeriodicExportConfigStartTimeOutputReference {
	var returns LookerInstancePeriodicExportConfigStartTimeOutputReference
	_jsii_.Get(
		j,
		"startTime",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) StartTimeInput() *LookerInstancePeriodicExportConfigStartTime {
	var returns *LookerInstancePeriodicExportConfigStartTime
	_jsii_.Get(
		j,
		"startTimeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewLookerInstancePeriodicExportConfigOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) LookerInstancePeriodicExportConfigOutputReference {
	_init_.Initialize()

	if err := validateNewLookerInstancePeriodicExportConfigOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_LookerInstancePeriodicExportConfigOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.lookerInstance.LookerInstancePeriodicExportConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewLookerInstancePeriodicExportConfigOutputReference_Override(l LookerInstancePeriodicExportConfigOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.lookerInstance.LookerInstancePeriodicExportConfigOutputReference",
		[]any{terraformResource, terraformAttribute},
		l,
	)
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) SetGcsUri(val *string) {
	if err := j.validateSetGcsUriParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"gcsUri",
		val,
	)
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) SetInternalValue(val *LookerInstancePeriodicExportConfig) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) SetKmsKey(val *string) {
	if err := j.validateSetKmsKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"kmsKey",
		val,
	)
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := l.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		l,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := l.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := l.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		l,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := l.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		l,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := l.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		l,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := l.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		l,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := l.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		l,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := l.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		l,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := l.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		l,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := l.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		l,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) PutStartTime(value *LookerInstancePeriodicExportConfigStartTime) {
	if err := l.validatePutStartTimeParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		l,
		"putStartTime",
		[]any{value},
	)
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := l.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		l,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (l *jsiiProxy_LookerInstancePeriodicExportConfigOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		l,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
