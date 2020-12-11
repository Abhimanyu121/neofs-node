package searchsvc

import (
	"context"
	"crypto/ecdsa"

	"github.com/nspcc-dev/neofs-api-go/pkg/client"
	"github.com/nspcc-dev/neofs-api-go/pkg/container"
	objectSDK "github.com/nspcc-dev/neofs-api-go/pkg/object"
	"github.com/nspcc-dev/neofs-node/pkg/network"
	"github.com/nspcc-dev/neofs-node/pkg/services/object_manager/placement"
	"github.com/nspcc-dev/neofs-node/pkg/util/logger"
	"go.uber.org/zap"
)

type statusError struct {
	status int
	err    error
}

type execCtx struct {
	svc *Service

	ctx context.Context

	prm Prm

	statusError

	log *logger.Logger
}

const (
	statusUndefined int = iota
	statusOK
)

func (exec *execCtx) prepare() {
	if _, ok := exec.prm.writer.(*uniqueIDWriter); !ok {
		exec.prm.writer = newUniqueAddressWriter(exec.prm.writer)
	}
}

func (exec *execCtx) setLogger(l *logger.Logger) {
	exec.log = l.With(
		zap.String("request", "SEARCH"),
		zap.Stringer("container", exec.containerID()),
		zap.Bool("local", exec.isLocal()),
		zap.Bool("with session", exec.prm.common.SessionToken() != nil),
		zap.Bool("with bearer", exec.prm.common.BearerToken() != nil),
	)
}

func (exec execCtx) context() context.Context {
	return exec.ctx
}

func (exec execCtx) isLocal() bool {
	return exec.prm.common.LocalOnly()
}

func (exec execCtx) key() *ecdsa.PrivateKey {
	return exec.prm.key
}

func (exec execCtx) callOptions() []client.CallOption {
	return exec.prm.callOpts
}

func (exec execCtx) remotePrm() *client.SearchObjectParams {
	return &exec.prm.SearchObjectParams
}

func (exec *execCtx) containerID() *container.ID {
	return exec.prm.ContainerID()
}

func (exec *execCtx) searchFilters() objectSDK.SearchFilters {
	return exec.prm.SearchFilters()
}

func (exec *execCtx) generateTraverser(cid *container.ID) (*placement.Traverser, bool) {
	t, err := exec.svc.traverserGenerator.generateTraverser(cid)

	switch {
	default:
		exec.status = statusUndefined
		exec.err = err

		exec.log.Debug("could not generate container traverser",
			zap.String("error", err.Error()),
		)

		return nil, false
	case err == nil:
		return t, true
	}
}

func (exec execCtx) remoteClient(node *network.Address) (searchClient, bool) {
	ipAddr, err := node.IPAddrString()

	log := exec.log.With(zap.Stringer("node", node))

	switch {
	default:
		exec.status = statusUndefined
		exec.err = err

		log.Debug("could not calculate node IP address")
	case err == nil:
		c, err := exec.svc.clientCache.get(exec.key(), ipAddr)

		switch {
		default:
			exec.status = statusUndefined
			exec.err = err

			log.Debug("could not construct remote node client")
		case err == nil:
			return c, true
		}
	}

	return nil, false
}

func (exec *execCtx) writeIDList(ids []*objectSDK.ID) {
	err := exec.prm.writer.WriteIDs(ids)

	switch {
	default:
		exec.status = statusUndefined
		exec.err = err

		exec.log.Debug("could not write object identifiers",
			zap.String("error", err.Error()),
		)
	case err == nil:
		exec.status = statusOK
		exec.err = nil
	}
}