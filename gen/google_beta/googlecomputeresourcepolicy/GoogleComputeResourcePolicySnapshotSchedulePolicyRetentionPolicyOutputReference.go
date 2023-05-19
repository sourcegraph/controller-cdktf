package googlecomputeresourcepolicy

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlecomputeresourcepolicy/internal"
)

type GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference interface {
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
	InternalValue() *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy
	SetInternalValue(val *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy)
	MaxRetentionDays() *float64
	SetMaxRetentionDays(val *float64)
	MaxRetentionDaysInput() *float64
	OnSourceDiskDelete() *string
	SetOnSourceDiskDelete(val *string)
	OnSourceDiskDeleteInput() *string
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
	ResetOnSourceDiskDelete()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference
type jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) InternalValue() *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy {
	var returns *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) MaxRetentionDays() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetentionDays",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) MaxRetentionDaysInput() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"maxRetentionDaysInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) OnSourceDiskDelete() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onSourceDiskDelete",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) OnSourceDiskDeleteInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"onSourceDiskDeleteInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference{}

	_jsii_.Create(
		"google-beta.googleComputeResourcePolicy.GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference_Override(g GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"google-beta.googleComputeResourcePolicy.GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference)SetInternalValue(val *GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicy) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference)SetMaxRetentionDays(val *float64) {
	if err := j.validateSetMaxRetentionDaysParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"maxRetentionDays",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference)SetOnSourceDiskDelete(val *string) {
	if err := j.validateSetOnSourceDiskDeleteParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"onSourceDiskDelete",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) ResetOnSourceDiskDelete() {
	_jsii_.InvokeVoid(
		g,
		"resetOnSourceDiskDelete",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleComputeResourcePolicySnapshotSchedulePolicyRetentionPolicyOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

