package ingress

import (
	"fmt"

	runtimev1alpha1 "github.com/crossplane/crossplane-runtime/apis/core/v1alpha1"
	"k8s.io/api/networking/v1beta1"
	"sigs.k8s.io/controller-runtime/pkg/client"

	standardv1alpha1 "github.com/oam-dev/kubevela/api/v1alpha1"
)

const TypeNginx = "nginx"

const (
	StatusReady  = "Ready"
	StatusSynced = "Synced"
)

type RouteIngress interface {
	Construct(routeTrait *standardv1alpha1.Route) []*v1beta1.Ingress
	CheckStatus(routeTrait *standardv1alpha1.Route) (string, []runtimev1alpha1.Condition)
}

func GetRouteIngress(provider string, client client.Client) (RouteIngress, error) {
	var routeIngress RouteIngress
	switch provider {
	case TypeNginx, "":
		routeIngress = &Nginx{Client: client}
	default:
		return nil, fmt.Errorf("unknow route ingress provider '%v', only '%s' is supported now", provider, TypeNginx)
	}
	return routeIngress, nil
}