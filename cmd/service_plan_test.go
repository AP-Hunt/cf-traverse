package cmd_test

import (
	"bytes"

	"code.cloudfoundry.org/cli/plugin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/AP-Hunt/cf-traverse/cmd"
	"github.com/AP-Hunt/cf-traverse/testfixtures"
)

var _ = Describe("service_plan", func() {
	var apiServer *testfixtures.APIServer
	var cliConnection plugin.CliConnection
	var out bytes.Buffer

	BeforeEach(func() {
		apiServer = testfixtures.NewAPIServer()
		cliConnection = testfixtures.NewTestCLIConnection("http://" + apiServer.ListenerAddr())
		out = bytes.Buffer{}
		testfixtures.ConfigureAPIServer(apiServer)
	})

	AfterEach(func() {
		apiServer.Stop()
	})

	Describe("instances_of SERVICE_PLAN_GUID", func() {
		It("gets all of the instances of the gievn service plan guid", func() {
			cmd := NewServicePlansCommand(cliConnection)
			cmd.SetArgs([]string{"instances_of", testfixtures.V3ServicePlanGuid})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3ServiceInstancesBySinglePlanListing))
		})
	})
})
