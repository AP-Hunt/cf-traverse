package cmd_test

import (
	"bytes"

	"code.cloudfoundry.org/cli/plugin"
	. "github.com/AP-Hunt/cf-traverse/cmd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/AP-Hunt/cf-traverse/testfixtures"
)

var _ = Describe("service", func() {
	var apiServer *testfixtures.APIServer
	var cliConnection plugin.CliConnection
	var out bytes.Buffer

	BeforeEach(func() {
		apiServer = testfixtures.NewAPIServer()
		cliConnection = testfixtures.NewTestCLIConnection("http://" + apiServer.ListenerAddr())
		out = bytes.Buffer{}

		apiServer.PathReturns(testfixtures.V3ServiceInstancePath, []byte(testfixtures.V3ServiceInstance))
		apiServer.PathReturns(testfixtures.V3SpacePath, []byte(testfixtures.V3Space))
		apiServer.PathReturns(testfixtures.V3OrgPath, []byte(testfixtures.V3Org))
	})

	AfterEach(func() {
		apiServer.Stop()
	})

	Describe("space SERVICE_GUID", func() {
		It("gets the space the service belongs to", func() {
			cmd := NewServiceCommand(cliConnection)
			cmd.SetArgs([]string{"space", testfixtures.V3ServiceInstanceGuid})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3Space))
		})
	})

	Describe("org SERVICE_GUID", func() {
		It("gets the org the service belongs to", func() {

			cmd := NewServiceCommand(cliConnection)
			cmd.SetArgs([]string{"org", testfixtures.V3ServiceInstanceGuid})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3Org))
		})
	})
})
