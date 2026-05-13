package googlebackupdrrestoreworkload

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlebackupdrrestoreworkload/internal"
)

type GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference interface {
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
	ConsumeAllocationType() *string
	SetConsumeAllocationType(val *string)
	ConsumeAllocationTypeInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	// Experimental.
	Fqn() *string
	InternalValue() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity
	SetInternalValue(val *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity)
	Key() *string
	SetKey(val *string)
	KeyInput() *string
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktf.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktf.IInterpolatingParent)
	Values() *[]*string
	SetValues(val *[]*string)
	ValuesInput() *[]*string
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
	ResetConsumeAllocationType()
	ResetKey()
	ResetValues()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference
type jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) ConsumeAllocationType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumeAllocationType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) ConsumeAllocationTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"consumeAllocationTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) InternalValue() *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity {
	var returns *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) Key() *string {
	var returns *string
	_jsii_.Get(
		j,
		"key",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) KeyInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"keyInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) Values() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"values",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) ValuesInput() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"valuesInput",
		&returns,
	)
	return returns
}


func NewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference_Override(g GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google-beta.googleBackupDrRestoreWorkload.GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference)SetConsumeAllocationType(val *string) {
	if err := j.validateSetConsumeAllocationTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"consumeAllocationType",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference)SetInternalValue(val *GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinity) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference)SetKey(val *string) {
	if err := j.validateSetKeyParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"key",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference)SetValues(val *[]*string) {
	if err := j.validateSetValuesParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"values",
		val,
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) ResetConsumeAllocationType() {
	_jsii_.InvokeVoid(
		g,
		"resetConsumeAllocationType",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) ResetKey() {
	_jsii_.InvokeVoid(
		g,
		"resetKey",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) ResetValues() {
	_jsii_.InvokeVoid(
		g,
		"resetValues",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleBackupDrRestoreWorkloadComputeInstanceRestorePropertiesAllocationAffinityOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

