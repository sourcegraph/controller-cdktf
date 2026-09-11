package googleoracledatabaseexadbvmcluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googleoracledatabaseexadbvmcluster/internal"
)

type GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference interface {
	cdktn.ComplexObject
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
	InternalValue() *GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptions
	SetInternalValue(val *GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptions)
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
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	// Experimental.
	ComputeFqn() *string
	// Experimental.
	GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{}
	// Experimental.
	GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable
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
	InterpolationAsList() cdktn.IResolvable
	// Experimental.
	InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable
	ResetIsDiagnosticsEventsEnabled()
	ResetIsHealthMonitoringEnabled()
	ResetIsIncidentLogsEnabled()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference
type jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) InternalValue() *GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptions {
	var returns *GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptions
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsDiagnosticsEventsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDiagnosticsEventsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsDiagnosticsEventsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isDiagnosticsEventsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsHealthMonitoringEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHealthMonitoringEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsHealthMonitoringEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isHealthMonitoringEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsIncidentLogsEnabled() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIncidentLogsEnabled",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) IsIncidentLogsEnabledInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"isIncidentLogsEnabledInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}


func NewGoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseExadbVmCluster.GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference_Override(g GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleOracleDatabaseExadbVmCluster.GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetInternalValue(val *GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptions) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetIsDiagnosticsEventsEnabled(val interface{}) {
	if err := j.validateSetIsDiagnosticsEventsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isDiagnosticsEventsEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetIsHealthMonitoringEnabled(val interface{}) {
	if err := j.validateSetIsHealthMonitoringEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isHealthMonitoringEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetIsIncidentLogsEnabled(val interface{}) {
	if err := j.validateSetIsIncidentLogsEnabledParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"isIncidentLogsEnabled",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateGetBooleanAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"getBooleanAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
	if err := g.validateInterpolationForAttributeParameters(terraformAttribute); err != nil {
		panic(err)
	}
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationForAttribute",
		[]interface{}{terraformAttribute},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ResetIsDiagnosticsEventsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIsDiagnosticsEventsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ResetIsHealthMonitoringEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIsHealthMonitoringEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ResetIsIncidentLogsEnabled() {
	_jsii_.InvokeVoid(
		g,
		"resetIsIncidentLogsEnabled",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
	if err := g.validateResolveParameters(context); err != nil {
		panic(err)
	}
	var returns interface{}

	_jsii_.Invoke(
		g,
		"resolve",
		[]interface{}{context},
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleOracleDatabaseExadbVmClusterPropertiesDataCollectionOptionsOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

