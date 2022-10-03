//go:build no_runtime_type_checking
// +build no_runtime_type_checking

package githubauthbackend

// Building without runtime type checking enabled, so all the below just return nil

func (g *jsiiProxy_GithubAuthBackendTuneList) validateGetParameters(index *float64) error {
	return nil
}

func (g *jsiiProxy_GithubAuthBackendTuneList) validateResolveParameters(_context cdktf.IResolveContext) error {
	return nil
}

func (j *jsiiProxy_GithubAuthBackendTuneList) validateSetInternalValueParameters(val interface{}) error {
	return nil
}

func (j *jsiiProxy_GithubAuthBackendTuneList) validateSetTerraformAttributeParameters(val *string) error {
	return nil
}

func (j *jsiiProxy_GithubAuthBackendTuneList) validateSetTerraformResourceParameters(val cdktf.IInterpolatingParent) error {
	return nil
}

func (j *jsiiProxy_GithubAuthBackendTuneList) validateSetWrapsSetParameters(val *bool) error {
	return nil
}

func validateNewGithubAuthBackendTuneListParameters(terraformResource cdktf.IInterpolatingParent, terraformAttribute *string, wrapsSet *bool) error {
	return nil
}

