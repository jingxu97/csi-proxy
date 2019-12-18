// Code generated by csi-proxy-api-gen. DO NOT EDIT.

package smb

import (
	"github.com/kubernetes-csi/csi-proxy/client/apiversion"
	srvtypes "github.com/kubernetes-csi/csi-proxy/internal/server/types"
	"github.com/kubernetes-csi/csi-proxy/internal/server/smb/internal"
	"github.com/kubernetes-csi/csi-proxy/internal/server/smb/internal/v1alpha1"
)

const name = "smb"

// ensure the server defines all the required methods
var _ internal.ServerInterface = &Server{}

func (s *Server) VersionedAPIs() []*srvtypes.VersionedAPI {
	v1alpha1Server := v1alpha1.NewVersionedServer(s)

	return []*srvtypes.VersionedAPI{
		{
			Group:      name,
			Version:    apiversion.NewVersionOrPanic("v1alpha1"),
			Registrant: v1alpha1Server.Register,
		},
	}
}
