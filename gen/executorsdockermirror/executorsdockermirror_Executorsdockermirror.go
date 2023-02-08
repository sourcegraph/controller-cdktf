// executorsdockermirror
package executorsdockermirror

import (
	_jsii_ "github.com/aws/jsii-runtime-go/runtime"
	_init_ "github.com/sourcegraph/controller-cdktf/gen/executorsdockermirror/jsii"

	"github.com/aws/constructs-go/constructs/v10"
	"github.com/hashicorp/terraform-cdk-go/cdktf"
	"github.com/sourcegraph/controller-cdktf/gen/executorsdockermirror/internal"
)

type Executorsdockermirror interface {
	cdktf.TerraformModule
	AssignPublicIp() *bool
	SetAssignPublicIp(val *bool)
	BootDiskSize() *float64
	SetBootDiskSize(val *float64)
	// Experimental.
	CdktfStack() cdktf.TerraformStack
	// Experimental.
	ConstructNodeMetadata() *map[string]interface{}
	// Experimental.
	DependsOn() *[]*string
	// Experimental.
	SetDependsOn(val *[]*string)
	DiskSize() *float64
	SetDiskSize(val *float64)
	// Experimental.
	ForEach() cdktf.ITerraformIterator
	// Experimental.
	SetForEach(val cdktf.ITerraformIterator)
	// Experimental.
	Fqn() *string
	// Experimental.
	FriendlyUniqueId() *string
	HttpAccessCidrRanges() *[]*string
	SetHttpAccessCidrRanges(val *[]*string)
	HttpMetricsAccessCidrRanges() *[]*string
	SetHttpMetricsAccessCidrRanges(val *[]*string)
	InstanceTagPrefix() *string
	SetInstanceTagPrefix(val *string)
	IpAddressOutput() *string
	MachineImage() *string
	SetMachineImage(val *string)
	MachineType() *string
	SetMachineType(val *string)
	NetworkId() *string
	SetNetworkId(val *string)
	// The tree node.
	Node() constructs.Node
	// Experimental.
	Providers() *[]interface{}
	// Experimental.
	RawOverrides() interface{}
	// Experimental.
	SkipAssetCreationFromLocalModules() *bool
	// Experimental.
	Source() *string
	SubnetId() *string
	SetSubnetId(val *string)
	// Experimental.
	Version() *string
	Zone() *string
	SetZone(val *string)
	// Experimental.
	AddOverride(path *string, value interface{})
	// Experimental.
	AddProvider(provider interface{})
	// Experimental.
	GetString(output *string) *string
	// Experimental.
	InterpolationForOutput(moduleOutput *string) cdktf.IResolvable
	// Overrides the auto-generated logical ID with a specific ID.
	// Experimental.
	OverrideLogicalId(newLogicalId *string)
	// Resets a previously passed logical Id to use the auto-generated logical id again.
	// Experimental.
	ResetOverrideLogicalId()
	SynthesizeAttributes() *map[string]interface{}
	// Experimental.
	ToMetadata() interface{}
	// Returns a string representation of this construct.
	ToString() *string
	// Experimental.
	ToTerraform() interface{}
}

// The jsii proxy struct for Executorsdockermirror
type jsiiProxy_Executorsdockermirror struct {
	internal.Type__cdktfTerraformModule
}

func (j *jsiiProxy_Executorsdockermirror) AssignPublicIp() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"assignPublicIp",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) BootDiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"bootDiskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) CdktfStack() cdktf.TerraformStack {
	var returns cdktf.TerraformStack
	_jsii_.Get(
		j,
		"cdktfStack",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) ConstructNodeMetadata() *map[string]interface{} {
	var returns *map[string]interface{}
	_jsii_.Get(
		j,
		"constructNodeMetadata",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) DependsOn() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"dependsOn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) DiskSize() *float64 {
	var returns *float64
	_jsii_.Get(
		j,
		"diskSize",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) ForEach() cdktf.ITerraformIterator {
	var returns cdktf.ITerraformIterator
	_jsii_.Get(
		j,
		"forEach",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) Fqn() *string {
	var returns *string
	_jsii_.Get(
		j,
		"fqn",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) FriendlyUniqueId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"friendlyUniqueId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) HttpAccessCidrRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpAccessCidrRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) HttpMetricsAccessCidrRanges() *[]*string {
	var returns *[]*string
	_jsii_.Get(
		j,
		"httpMetricsAccessCidrRanges",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) InstanceTagPrefix() *string {
	var returns *string
	_jsii_.Get(
		j,
		"instanceTagPrefix",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) IpAddressOutput() *string {
	var returns *string
	_jsii_.Get(
		j,
		"ipAddressOutput",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) MachineImage() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineImage",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) MachineType() *string {
	var returns *string
	_jsii_.Get(
		j,
		"machineType",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) NetworkId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"networkId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) Node() constructs.Node {
	var returns constructs.Node
	_jsii_.Get(
		j,
		"node",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) Providers() *[]interface{} {
	var returns *[]interface{}
	_jsii_.Get(
		j,
		"providers",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) RawOverrides() interface{} {
	var returns interface{}
	_jsii_.Get(
		j,
		"rawOverrides",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) SkipAssetCreationFromLocalModules() *bool {
	var returns *bool
	_jsii_.Get(
		j,
		"skipAssetCreationFromLocalModules",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) Source() *string {
	var returns *string
	_jsii_.Get(
		j,
		"source",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) SubnetId() *string {
	var returns *string
	_jsii_.Get(
		j,
		"subnetId",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) Version() *string {
	var returns *string
	_jsii_.Get(
		j,
		"version",
		&returns,
	)
	return returns
}

func (j *jsiiProxy_Executorsdockermirror) Zone() *string {
	var returns *string
	_jsii_.Get(
		j,
		"zone",
		&returns,
	)
	return returns
}


func NewExecutorsdockermirror(scope constructs.Construct, id *string, options *ExecutorsdockermirrorOptions) Executorsdockermirror {
	_init_.Initialize()

	if err := validateNewExecutorsdockermirrorParameters(scope, id, options); err != nil {
		panic(err)
	}
	j := jsiiProxy_Executorsdockermirror{}

	_jsii_.Create(
		"executorsdockermirror.Executorsdockermirror",
		[]interface{}{scope, id, options},
		&j,
	)

	return &j
}

func NewExecutorsdockermirror_Override(e Executorsdockermirror, scope constructs.Construct, id *string, options *ExecutorsdockermirrorOptions) {
	_init_.Initialize()

	_jsii_.Create(
		"executorsdockermirror.Executorsdockermirror",
		[]interface{}{scope, id, options},
		e,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetAssignPublicIp(val *bool) {
	_jsii_.Set(
		j,
		"assignPublicIp",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetBootDiskSize(val *float64) {
	_jsii_.Set(
		j,
		"bootDiskSize",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetDependsOn(val *[]*string) {
	_jsii_.Set(
		j,
		"dependsOn",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetDiskSize(val *float64) {
	_jsii_.Set(
		j,
		"diskSize",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetForEach(val cdktf.ITerraformIterator) {
	_jsii_.Set(
		j,
		"forEach",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetHttpAccessCidrRanges(val *[]*string) {
	_jsii_.Set(
		j,
		"httpAccessCidrRanges",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetHttpMetricsAccessCidrRanges(val *[]*string) {
	_jsii_.Set(
		j,
		"httpMetricsAccessCidrRanges",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetInstanceTagPrefix(val *string) {
	if err := j.validateSetInstanceTagPrefixParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"instanceTagPrefix",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetMachineImage(val *string) {
	_jsii_.Set(
		j,
		"machineImage",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetMachineType(val *string) {
	_jsii_.Set(
		j,
		"machineType",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetNetworkId(val *string) {
	if err := j.validateSetNetworkIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"networkId",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetSubnetId(val *string) {
	if err := j.validateSetSubnetIdParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"subnetId",
		val,
	)
}

func (j *jsiiProxy_Executorsdockermirror)SetZone(val *string) {
	if err := j.validateSetZoneParameters(val); err != nil {
		panic(err)
	}
	_jsii_.Set(
		j,
		"zone",
		val,
	)
}

// Checks if `x` is a construct.
//
// Use this method instead of `instanceof` to properly detect `Construct`
// instances, even when the construct library is symlinked.
//
// Explanation: in JavaScript, multiple copies of the `constructs` library on
// disk are seen as independent, completely different libraries. As a
// consequence, the class `Construct` in each copy of the `constructs` library
// is seen as a different class, and an instance of one class will not test as
// `instanceof` the other class. `npm install` will not create installations
// like this, but users may manually symlink construct libraries together or
// use a monorepo tool: in those cases, multiple copies of the `constructs`
// library can be accidentally installed, and `instanceof` will behave
// unpredictably. It is safest to avoid using `instanceof`, and using
// this type-testing method instead.
//
// Returns: true if `x` is an object created from a class which extends `Construct`.
func Executorsdockermirror_IsConstruct(x interface{}) *bool {
	_init_.Initialize()

	if err := validateExecutorsdockermirror_IsConstructParameters(x); err != nil {
		panic(err)
	}
	var returns *bool

	_jsii_.StaticInvoke(
		"executorsdockermirror.Executorsdockermirror",
		"isConstruct",
		[]interface{}{x},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsdockermirror) AddOverride(path *string, value interface{}) {
	if err := e.validateAddOverrideParameters(path, value); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addOverride",
		[]interface{}{path, value},
	)
}

func (e *jsiiProxy_Executorsdockermirror) AddProvider(provider interface{}) {
	if err := e.validateAddProviderParameters(provider); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"addProvider",
		[]interface{}{provider},
	)
}

func (e *jsiiProxy_Executorsdockermirror) GetString(output *string) *string {
	if err := e.validateGetStringParameters(output); err != nil {
		panic(err)
	}
	var returns *string

	_jsii_.Invoke(
		e,
		"getString",
		[]interface{}{output},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsdockermirror) InterpolationForOutput(moduleOutput *string) cdktf.IResolvable {
	if err := e.validateInterpolationForOutputParameters(moduleOutput); err != nil {
		panic(err)
	}
	var returns cdktf.IResolvable

	_jsii_.Invoke(
		e,
		"interpolationForOutput",
		[]interface{}{moduleOutput},
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsdockermirror) OverrideLogicalId(newLogicalId *string) {
	if err := e.validateOverrideLogicalIdParameters(newLogicalId); err != nil {
		panic(err)
	}
	_jsii_.InvokeVoid(
		e,
		"overrideLogicalId",
		[]interface{}{newLogicalId},
	)
}

func (e *jsiiProxy_Executorsdockermirror) ResetOverrideLogicalId() {
	_jsii_.InvokeVoid(
		e,
		"resetOverrideLogicalId",
		nil, // no parameters
	)
}

func (e *jsiiProxy_Executorsdockermirror) SynthesizeAttributes() *map[string]interface{} {
	var returns *map[string]interface{}

	_jsii_.Invoke(
		e,
		"synthesizeAttributes",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsdockermirror) ToMetadata() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toMetadata",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsdockermirror) ToString() *string {
	var returns *string

	_jsii_.Invoke(
		e,
		"toString",
		nil, // no parameters
		&returns,
	)

	return returns
}

func (e *jsiiProxy_Executorsdockermirror) ToTerraform() interface{} {
	var returns interface{}

	_jsii_.Invoke(
		e,
		"toTerraform",
		nil, // no parameters
		&returns,
	)

	return returns
}

