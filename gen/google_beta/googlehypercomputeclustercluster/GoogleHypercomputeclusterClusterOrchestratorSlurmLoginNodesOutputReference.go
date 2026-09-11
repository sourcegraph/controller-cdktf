package googlehypercomputeclustercluster

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/google_beta/jsii"

	"github.com/open-constructs/cdk-terrain-go/cdktn"
	"github.com/sourcegraph/controller-cdktf/gen/google_beta/googlehypercomputeclustercluster/internal"
)

type GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference interface {
	cdktn.ComplexObject
	BootDisk() GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDiskOutputReference
	BootDiskInput() *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk
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
	Count() *string
	SetCount(val *string)
	CountInput() *string
	// The creation stack of this resolvable which will be appended to errors thrown during resolution.
	//
	// If this returns an empty array the stack will not be attached.
	// Experimental.
	CreationStack() *[]*string
	EnableOsLogin() interface{}
	SetEnableOsLogin(val interface{})
	EnableOsLoginInput() interface{}
	EnablePublicIps() interface{}
	SetEnablePublicIps(val interface{})
	EnablePublicIpsInput() interface{}
	// Experimental.
	Fqn() *string
	Instances() GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesInstancesList
	InternalValue() *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes
	SetInternalValue(val *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes)
	Labels() *map[string]*string
	SetLabels(val *map[string]*string)
	LabelsInput() *map[string]*string
	MachineType() *string
	SetMachineType(val *string)
	MachineTypeInput() *string
	StartupScript() *string
	SetStartupScript(val *string)
	StartupScriptInput() *string
	StorageConfigs() GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesStorageConfigsList
	StorageConfigsInput() interface{}
	// Experimental.
	TerraformAttribute() *string
	// Experimental.
	SetTerraformAttribute(val *string)
	// Experimental.
	TerraformResource() cdktn.IInterpolatingParent
	// Experimental.
	SetTerraformResource(val cdktn.IInterpolatingParent)
	Zone() *string
	SetZone(val *string)
	ZoneInput() *string
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
	PutBootDisk(value *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk)
	PutStorageConfigs(value interface{})
	ResetBootDisk()
	ResetEnableOsLogin()
	ResetEnablePublicIps()
	ResetLabels()
	ResetStartupScript()
	ResetStorageConfigs()
	// Produce the Token's value at resolution time.
	// Experimental.
	Resolve(context cdktn.IResolveContext) interface{}
	// Return a string representation of this resolvable object.
	//
	// Returns a reversible string representation.
	// Experimental.
	ToString() *string
}

// The jsii proxy struct for GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference
type jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference struct {
	internal.Type__cdktnComplexObject
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) BootDisk() GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDiskOutputReference {
	var returns GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDiskOutputReference
	_jsii_.Get(
		j,
		"bootDisk",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) BootDiskInput() *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk {
	var returns *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk
	_jsii_.Get(
		j,
		"bootDiskInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ComplexObjectIndex() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"complexObjectIndex",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ComplexObjectIsFromSet() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"complexObjectIsFromSet",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Count() *string {
	var returns *string
	_jsii_.Get(
		j,
		"count",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) CountInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"countInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) CreationStack() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"creationStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) EnableOsLogin() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOsLogin",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) EnableOsLoginInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enableOsLoginInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) EnablePublicIps() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicIps",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) EnablePublicIpsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"enablePublicIpsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Instances() GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesInstancesList {
	var returns GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesInstancesList
	_jsii_.Get(
		j,
		"instances",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) InternalValue() *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes {
	var returns *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes
	_jsii_.Get(
		j,
		"internalValue",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Labels() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labels",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) LabelsInput() *map[string]*string {
	var returns *map[string]*string
	_jsii_.Get(
		j,
		"labelsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) MachineTypeInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineTypeInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) StartupScript() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScript",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) StartupScriptInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"startupScriptInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) StorageConfigs() GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesStorageConfigsList {
	var returns GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesStorageConfigsList
	_jsii_.Get(
		j,
		"storageConfigs",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) StorageConfigsInput() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"storageConfigsInput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) TerraformAttribute() *string {
	var returns *string
	_jsii_.Get(
		j,
		"terraformAttribute",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) TerraformResource() cdktn.IInterpolatingParent {
	var returns cdktn.IInterpolatingParent
	_jsii_.Get(
		j,
		"terraformResource",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ZoneInput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zoneInput",
		&returns,
	)
	return returns
}


func NewGoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference(terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference {
	_init_.Initialize()

	if err := validateNewGoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReferenceParameters(terraformResource, terraformAttribute); err != nil {
		panic(err)
	}
	j := jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference{}

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		&j,
	)

	return &j
}

func NewGoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference_Override(g GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference, terraformResource cdktn.IInterpolatingParent, terraformAttribute *string) {
	_init_.Initialize()

	_jsii_.Create(
		"@cdktn/provider-google-beta.googleHypercomputeclusterCluster.GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference",
		[]interface{}{terraformResource, terraformAttribute},
		g,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetComplexObjectIndex(val interface{}) {
	if err := j.validateSetComplexObjectIndexParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIndex",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetComplexObjectIsFromSet(val *bool) {
	if err := j.validateSetComplexObjectIsFromSetParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"complexObjectIsFromSet",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetCount(val *string) {
	if err := j.validateSetCountParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"count",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetEnableOsLogin(val interface{}) {
	if err := j.validateSetEnableOsLoginParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enableOsLogin",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetEnablePublicIps(val interface{}) {
	if err := j.validateSetEnablePublicIpsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"enablePublicIps",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetInternalValue(val *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodes) {
	if err := j.validateSetInternalValueParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"internalValue",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetLabels(val *map[string]*string) {
	if err := j.validateSetLabelsParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"labels",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetMachineType(val *string) {
	if err := j.validateSetMachineTypeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetStartupScript(val *string) {
	if err := j.validateSetStartupScriptParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"startupScript",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetTerraformAttribute(val *string) {
	if err := j.validateSetTerraformAttributeParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformAttribute",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetTerraformResource(val cdktn.IInterpolatingParent) {
	if err := j.validateSetTerraformResourceParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"terraformResource",
		val,
	)
}

func (j *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ComputeFqn() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"computeFqn",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetAnyMapAttribute(terraformAttribute *string) *map[string]interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetBooleanAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetBooleanMapAttribute(terraformAttribute *string) *map[string]*bool {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetListAttribute(terraformAttribute *string) *[]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetNumberAttribute(terraformAttribute *string) *float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetNumberListAttribute(terraformAttribute *string) *[]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetNumberMapAttribute(terraformAttribute *string) *map[string]*float64 {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetStringAttribute(terraformAttribute *string) *string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) GetStringMapAttribute(terraformAttribute *string) *map[string]*string {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) InterpolationAsList() cdktn.IResolvable {
	var returns cdktn.IResolvable

	_jsii_.Invoke(
		g,
		"interpolationAsList",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) InterpolationForAttribute(terraformAttribute *string) cdktn.IResolvable {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) PutBootDisk(value *GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesBootDisk) {
	if err := g.validatePutBootDiskParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putBootDisk",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) PutStorageConfigs(value interface{}) {
	if err := g.validatePutStorageConfigsParameters(value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		g,
		"putStorageConfigs",
		[]interface{}{value},
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetBootDisk() {
	_jsii_.InvokeVoid(
		g,
		"resetBootDisk",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetEnableOsLogin() {
	_jsii_.InvokeVoid(
		g,
		"resetEnableOsLogin",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetEnablePublicIps() {
	_jsii_.InvokeVoid(
		g,
		"resetEnablePublicIps",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetLabels() {
	_jsii_.InvokeVoid(
		g,
		"resetLabels",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetStartupScript() {
	_jsii_.InvokeVoid(
		g,
		"resetStartupScript",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ResetStorageConfigs() {
	_jsii_.InvokeVoid(
		g,
		"resetStorageConfigs",
		nil, // no parameters
	)
}

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) Resolve(context cdktn.IResolveContext) interface{} {
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

func (g *jsiiProxy_GoogleHypercomputeclusterClusterOrchestratorSlurmLoginNodesOutputReference) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		g,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

