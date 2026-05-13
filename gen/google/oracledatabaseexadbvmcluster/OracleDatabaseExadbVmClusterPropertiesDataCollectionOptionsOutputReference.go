package oracledatabaseexadbvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google/jsii"

	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/google/oracledatabaseexadbvmcluster/internal"
)

type OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference interface {
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
	InternalValue() *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions
	SetInternalValue(val *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions)
	IsDiagnosticsEventsEnabled() interface{}
	SetIsDiagnosticsEventsEnabled(val interface{})
	IsDiagnosticsEventsEnabledInput() interface{}
	IsHealthMonitoringEnabled() interface{}
	SetIsHealthMonitoringEnabled(val interface{})
	IsHealthMonitoringEnabledInput() interface{}
	IsIncidentLogsEnabled() interface{}
	SetIsIncidentLogsEnabled(val interface{})
	IsIncidentLogsEnabledInput() interface{}
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
	ResetIsDiagnosticsEventsEnabled()
	ResetIsHealthMonitoringEnabled()
	ResetIsIncidentLogsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(_context cdktf.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference
type jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference struct {
	internal.Type__cdktfComplexObject
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) InternalValue() *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions {
	var returns *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsDiagnosticsEventsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDiagnosticsEventsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsDiagnosticsEventsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDiagnosticsEventsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsHealthMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHealthMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsHealthMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHealthMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsIncidentLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIncidentLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsIncidentLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIncidentLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) TerraformResource() cdktf.IInterpolatingParent {
	var returns cdktf.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference{}

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseExadbVmCluster.OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference_Override(o OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference, terraformResource cdktf.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktf/provider-google.oracleDatabaseExadbVmCluster.OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		o,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetInternalValue(val *OracleDatabaseExadbVmClusterPropertiesDataCollectionOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetIsDiagnosticsEventsEnabled(val interface{}) {
	if err := j.validateSetIsDiagnosticsEventsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDiagnosticsEventsEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetIsHealthMonitoringEnabled(val interface{}) {
	if err := j.validateSetIsHealthMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isHealthMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetIsIncidentLogsEnabled(val interface{}) {
	if err := j.validateSetIsIncidentLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isIncidentLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetTerraformResource(val cdktf.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
	if err := o.validateGetAnyMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]interface{}

	_jsii_.Invoke(
		o,
		"getAnyMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktf.IResolvable {
	if err := o.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
	if err := o.validateGetBooleanMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*bool

	_jsii_.Invoke(
		o,
		"getBooleanMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
	if err := o.validateGetListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*string

	_jsii_.Invoke(
		o,
		"getListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
	if err := o.validateGetNumberAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *float64

	_jsii_.Invoke(
		o,
		"getNumberAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
	if err := o.validateGetNumberListAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *[]*float64

	_jsii_.Invoke(
		o,
		"getNumberListAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
	if err := o.validateGetNumberMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*float64

	_jsii_.Invoke(
		o,
		"getNumberMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
	if err := o.validateGetStringAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		o,
		"getStringAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
	if err := o.validateGetStringMapAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns *map[string]*string

	_jsii_.Invoke(
		o,
		"getStringMapAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) InterpolationAsList() cdktf.IResolvable {
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) InterpolationForAttribute(property *string) cdktf.IResolvable {
	if err := o.validateInterpolationForAttributeParameters(property); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		o,
		"interpolationForAttribute",
		[]interface{}{property},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ResetIsDiagnosticsEventsEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsDiagnosticsEventsEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ResetIsHealthMonitoringEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsHealthMonitoringEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ResetIsIncidentLogsEnabled() {
	_jsii_.InvokeVoid(
		o,
		"resetIsIncidentLogsEnabled",
		nil, // no parameters
	)
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) Resolve(_context cdktf.IResolveContext) interface{} {
	if err := o.validateResolveParameters(_context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		o,
		"resolve",
		[]interface{}{_context},
		&returns,
	)

	return returns
}

func (o *jsiiProxy_OracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		o,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

