package cmd_test

import (
	"bytes"
	"code.cloudfoundry.org/cli/plugin"
	"github.com/AP-Hunt/cf-traverse/testfixtures"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/AP-Hunt/cf-traverse/cmd"
)

var _ = Describe("service_offering", func() {
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

	Describe("instances_of SERVICE_OFFERING_GUID|SERVICE_OFFERING_NAME", func() {
		It("gets all of the instances of all of the service plans belonging to the given service offering guid", func() {
			cmd := NewServiceOfferingsCommand(cliConnection)
			cmd.SetArgs([]string{"instances_of", testfixtures.V3ServiceOfferingGuid})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3ServiceInstancesByMultiplePlanListing))
		})

		It("gets all of the instances of all of the service plans belonging to the given service offering name", func() {
			cmd := NewServiceOfferingsCommand(cliConnection)
			cmd.SetArgs([]string{"instances_of", testfixtures.V3ServiceOfferingName})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3ServiceInstancesByMultiplePlanListing))
		})
	})
})
