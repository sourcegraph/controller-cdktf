package discoveryenginedataconnector

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/discoveryenginedataconnector/internal"
)

type DiscoveryEngineDataConnectorEntitiesOutputReference interface {
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
	DataStore() *string
	EntityName() *string
	SetEntityName(val *string)
	EntityNameInput() *string
	// Experimental.
	Fqn() *string
	InternalValue() any
	SetInternalValue(val any)
	KeyPropertyMappings() *map[string]*string
	SetKeyPropertyMappings(val *map[string]*string)
	KeyPropertyMappingsInput() *map[string]*string
	Params() *string
	SetParams(val *string)
	ParamsInput() *string
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
	ResetEntityName()
	ResetKeyPropertyMappings()
	ResetParams()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) any
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for DiscoveryEngineDataConnectorEntitiesOutputReference
type jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) ComplexObjectIndex() any {
	var returns any
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) DataStore() *string {
	var returns *string
	_jsii_.Get(
		j,
		"dataStore",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) EntityName() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityName",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) EntityNameInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"entityNameInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) InternalValue() any {
	var returns any
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) KeyPropertyMappings() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"keyPropertyMappings",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) KeyPropertyMappingsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"keyPropertyMappingsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) Params() *string {
	var returns *string
	_jsii_.Get(
		j,
		"params",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) ParamsInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"paramsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func NewDiscoveryEngineDataConnectorEntitiesOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) DiscoveryEngineDataConnectorEntitiesOutputReference {
	_init_.Initialize()

	if err := validateNewDiscoveryEngineDataConnectorEntitiesOutputReferenceParameters(terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet); err != nil {
		panic(err)
	}
	j := jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineDataConnector.DiscoveryEngineDataConnectorEntitiesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		&j,
	)

	return &j
}

func NewDiscoveryEngineDataConnectorEntitiesOutputReference_Override(d DiscoveryEngineDataConnectorEntitiesOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, complexObjectIndex *float64, complexObjectIsFromSet *bool) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.discoveryEngineDataConnector.DiscoveryEngineDataConnectorEntitiesOutputReference",
		[]any{terraformResource, terraformAttribute, complexObjectIndex, complexObjectIsFromSet},
		d,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) SetComplexObjectIndex(val any) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) SetEntityName(val *string) {
	if err := j.validateSetEntityNameParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"entityName",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) SetInternalValue(val any) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) SetKeyPropertyMappings(val *map[string]*string) {
	if err := j.validateSetKeyPropertyMappingsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"keyPropertyMappings",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) SetParams(val *string) {
	if err := j.validateSetParamsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"params",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]any {
	if err := d.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]any

	_jsii_.Invoke(
		d,
		"getAnyMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := d.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"getBooleanAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := d.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		d,
		"getBooleanMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := d.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		d,
		"getListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := d.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		d,
		"getNumberAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := d.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		d,
		"getNumberListAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := d.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		d,
		"getNumberMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := d.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		d,
		"getStringAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := d.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		d,
		"getStringMapAttribute",
		[]any{terraformAttribute},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := d.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		d,
		"interpolationForAttribute",
		[]any{property},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) ResetEntityName() {
	_jsii_.InvokeVoid(
		d,
		"resetEntityName",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) ResetKeyPropertyMappings() {
	_jsii_.InvokeVoid(
		d,
		"resetKeyPropertyMappings",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) ResetParams() {
	_jsii_.InvokeVoid(
		d,
		"resetParams",
		nil, // no parameters
	)
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) Resolve(_context cdktf.IResolveContext) any {
	if err := d.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns any

	_jsii_.Invoke(
		d,
		"resolve",
		[]any{_context},
		&returns,
	)

	return returns
}

func (d *jsiiProxy_DiscoveryEngineDataConnectorEntitiesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		d,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}
