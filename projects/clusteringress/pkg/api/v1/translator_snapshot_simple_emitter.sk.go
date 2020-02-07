// Code generated by solo-kit. DO NOT EDIT.

package v1

import (
	"context"
	fmt "fmt"
	"time"

	github_com_solo_io_gloo_projects_clusteringress_pkg_api_external_knative "github.com/solo-io/gloo/projects/clusteringress/pkg/api/external/knative"
	gloo_solo_io "github.com/solo-io/gloo/projects/gloo/pkg/api/v1"

	"go.opencensus.io/stats"
	"go.uber.org/zap"

	"github.com/solo-io/go-utils/contextutils"
	"github.com/solo-io/go-utils/errutils"
	"github.com/solo-io/solo-kit/pkg/api/v1/clients"
)

type TranslatorSimpleEmitter interface {
	Snapshots(ctx context.Context) (<-chan *TranslatorSnapshot, <-chan error, error)
}

func NewTranslatorSimpleEmitter(aggregatedWatch clients.ResourceWatch) TranslatorSimpleEmitter {
	return NewTranslatorSimpleEmitterWithEmit(aggregatedWatch, make(chan struct{}))
}

func NewTranslatorSimpleEmitterWithEmit(aggregatedWatch clients.ResourceWatch, emit <-chan struct{}) TranslatorSimpleEmitter {
	return &translatorSimpleEmitter{
		aggregatedWatch: aggregatedWatch,
		forceEmit:       emit,
	}
}

type translatorSimpleEmitter struct {
	forceEmit       <-chan struct{}
	aggregatedWatch clients.ResourceWatch
}

func (c *translatorSimpleEmitter) Snapshots(ctx context.Context) (<-chan *TranslatorSnapshot, <-chan error, error) {
	snapshots := make(chan *TranslatorSnapshot)
	errs := make(chan error)

	untyped, watchErrs, err := c.aggregatedWatch(ctx)
	if err != nil {
		return nil, nil, err
	}

	go errutils.AggregateErrs(ctx, errs, watchErrs, "translator-emitter")

	go func() {
		currentSnapshot := TranslatorSnapshot{}
		timer := time.NewTicker(time.Second * 1)
		var previousHash uint64
		sync := func() {
			currentHash, err := currentSnapshot.Hash(nil)
			if err != nil {
				contextutils.LoggerFrom(ctx).Panicw("error while hashing, this should never happen", zap.Error(err))
			}
			if previousHash == currentHash {
				return
			}

			previousHash = currentHash

			stats.Record(ctx, mTranslatorSnapshotOut.M(1))
			sentSnapshot := currentSnapshot.Clone()
			snapshots <- &sentSnapshot
		}

		defer func() {
			close(snapshots)
			close(errs)
		}()

		for {
			record := func() { stats.Record(ctx, mTranslatorSnapshotIn.M(1)) }

			select {
			case <-timer.C:
				sync()
			case <-ctx.Done():
				return
			case <-c.forceEmit:
				sentSnapshot := currentSnapshot.Clone()
				snapshots <- &sentSnapshot
			case untypedList := <-untyped:
				record()

				currentSnapshot = TranslatorSnapshot{}
				for _, res := range untypedList {
					switch typed := res.(type) {
					case *gloo_solo_io.Secret:
						currentSnapshot.Secrets = append(currentSnapshot.Secrets, typed)
					case *github_com_solo_io_gloo_projects_clusteringress_pkg_api_external_knative.ClusterIngress:
						currentSnapshot.Clusteringresses = append(currentSnapshot.Clusteringresses, typed)
					default:
						select {
						case errs <- fmt.Errorf("TranslatorSnapshotEmitter "+
							"cannot process resource %v of type %T", res.GetMetadata().Ref(), res):
						case <-ctx.Done():
							return
						}
					}
				}

			}
		}
	}()
	return snapshots, errs, nil
}
