package testfixtures

import (
	plugin "code.cloudfoundry.org/cli/plugin"
	plugin_models "code.cloudfoundry.org/cli/plugin/models"
)

type TestCLIConnection struct {
	apiAddr string
}

func NewTestCLIConnection(apiAddr string) plugin.CliConnection {
	return &TestCLIConnection{
		apiAddr: apiAddr,
	}
}

func (cli *TestCLIConnection) CliCommandWithoutTerminalOutput(args ...string) ([]string, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) CliCommand(args ...string) ([]string, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetCurrentOrg() (plugin_models.Organization, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetCurrentSpace() (plugin_models.Space, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) Username() (string, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) UserGuid() (string, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) UserEmail() (string, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) IsLoggedIn() (bool, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

// IsSSLDisabled returns true if and only if the user is connected to the Cloud Controller API with the
// `--skip-ssl-validation` flag set unless the CLI configuration file cannot be read, in which case it
// returns an error.
func (cli *TestCLIConnection) IsSSLDisabled() (bool, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) HasOrganization() (bool, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) HasSpace() (bool, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) ApiEndpoint() (string, error) {
	return cli.apiAddr, nil
}

func (cli *TestCLIConnection) ApiVersion() (string, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) HasAPIEndpoint() (bool, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) LoggregatorEndpoint() (string, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) DopplerEndpoint() (string, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) AccessToken() (string, error) {
	return "access_token", nil
}

func (cli *TestCLIConnection) GetApp(_ string) (plugin_models.GetAppModel, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetApps() ([]plugin_models.GetAppsModel, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetOrgs() ([]plugin_models.GetOrgs_Model, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetSpaces() ([]plugin_models.GetSpaces_Model, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetOrgUsers(_ string, _ ...string) ([]plugin_models.GetOrgUsers_Model, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetSpaceUsers(_ string, _ string) ([]plugin_models.GetSpaceUsers_Model, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetServices() ([]plugin_models.GetServices_Model, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetService(_ string) (plugin_models.GetService_Model, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetOrg(_ string) (plugin_models.GetOrg_Model, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}

func (cli *TestCLIConnection) GetSpace(_ string) (plugin_models.GetSpace_Model, error) {
	panic("not implemented") // Deliberately not implemented. Not needed for its purpose.
}
