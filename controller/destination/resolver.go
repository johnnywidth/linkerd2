package destination

type streamingDestinationResolver interface {
	canResolve(host string, port int) (bool, error)
	streamResolution(host string, port int, listener endpointUpdateListener) error
	streamProfiles(host string, listener profileUpdateListener) error
	stop()
}
