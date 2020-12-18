package cmd_test

import (
	"bytes"

	"code.cloudfoundry.org/cli/plugin"
	. "github.com/AP-Hunt/cf-traverse/cmd"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/AP-Hunt/cf-traverse/testfixtures"
)

var _ = Describe("service_instance", func() {
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
		apiServer.PathReturns(testfixtures.V3ServiceInstanceListingPath, []byte(testfixtures.V3ServiceInstanceListing))
		apiServer.PathReturns(testfixtures.V3ServicePlanPath, []byte(testfixtures.V3ServicePlan))
	})

	AfterEach(func() {
		apiServer.Stop()
	})

	Describe("space SERVICE_INSTANCE_GUID|SERVICE_INSTANCE_NAME", func() {
		It("gets the space the service instance belongs to when given a UUID", func() {
			cmd := NewServiceCommand(cliConnection)
			cmd.SetArgs([]string{"space", testfixtures.V3ServiceInstanceGuid})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3Space))
		})

		It("gets the space the service instance belongs to when given a service instance name", func(){
			cmd := NewServiceCommand(cliConnection)
			cmd.SetArgs([]string{"space", testfixtures.V3ServiceInstanceName})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3Space))
		})
	})

	Describe("org SERVICE_INSTANCE_GUID|SERVICE_INSTANCE_NAME", func() {
		It("gets the org the service instance belongs to when give a UUID", func() {

			cmd := NewServiceCommand(cliConnection)
			cmd.SetArgs([]string{"org", testfixtures.V3ServiceInstanceGuid})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3Org))
		})

		It("gets the org the service instance belongs to when give a service instance name", func() {

			cmd := NewServiceCommand(cliConnection)
			cmd.SetArgs([]string{"org", testfixtures.V3ServiceInstanceName})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3Org))
		})
	})

	Describe("plan SERVICE_INSTANCE_GUID|SERVICE_INSTANCE_NAME", func(){
		It("gets the service plan the service instance is an instance of when give a UUID", func() {

			cmd := NewServiceCommand(cliConnection)
			cmd.SetArgs([]string{"plan", testfixtures.V3ServiceInstanceGuid})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3ServicePlan))
		})

		It("gets the service plan the service instance is an instance of when give a service instance name", func() {

			cmd := NewServiceCommand(cliConnection)
			cmd.SetArgs([]string{"plan", testfixtures.V3ServiceInstanceName})
			cmd.SetOut(&out)
			err := cmd.Execute()

			Expect(err).ToNot(HaveOccurred())
			Expect(out.String()).To(Equal(testfixtures.V3ServicePlan))
		})
	})
})
