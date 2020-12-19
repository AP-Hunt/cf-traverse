package testfixtures

// Configures the API server to serve the different
// responses at different paths
func ConfigureAPIServer(apiServer *APIServer) {
	// Orgs
	apiServer.PathReturns(V3OrgPath, []byte(V3Org))

	// Service instances
	apiServer.PathReturns(V3ServiceInstancePath, []byte(V3ServiceInstance))
	apiServer.PathReturns(V3ServiceInstanceByNameListingPath, []byte(V3ServiceInstancesByNameListing))
	apiServer.PathReturns(V3ServiceInstancesBySinglePlanListingPath, []byte(V3ServiceInstancesBySinglePlanListing))

	// Service offerings
	apiServer.PathReturns(V3ServiceOfferingPath, []byte(V3ServiceOffering))

	// Service plans
	apiServer.PathReturns(V3ServicePlanPath, []byte(V3ServicePlan))

	// Spaces
	apiServer.PathReturns(V3SpacePath, []byte(V3Space))
}
