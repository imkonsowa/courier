package providers

import "testing"

func TestNewRouteProvider_Ignitable(t *testing.T) {
	var _ Provider = (*RouteProvider)(nil)
}
