package tilerender

import (
	"fmt"
	"reflect"

	"app"
	"gopnik"
)

type zoomRender struct {
	Render  *RenderPool
	MinZoom uint
	MaxZoom uint
	Tags    []string
}

type BadCoordError struct {
	coord gopnik.TileCoord
}

func NewBadCoordError(coord gopnik.TileCoord) *BadCoordError {
	return &BadCoordError{
		coord: coord,
	}
}

func (self *BadCoordError) Error() string {
	return fmt.Sprintf("No situable render for %v", self.coord)
}

func IsBadCoordError(err error) bool {
	return reflect.TypeOf(err) == reflect.TypeOf(&BadCoordError{})
}

type MultiRenderPool struct {
	renders []zoomRender
}

func NewMultiRenderPool(poolsCfg app.RenderPoolsConfig) (*MultiRenderPool, error) {
	self := &MultiRenderPool{
		renders: make([]zoomRender, len(poolsCfg.RenderPools)),
	}

	for i := 0; i < len(self.renders); i++ {
		var err error
		self.renders[i].MinZoom = poolsCfg.RenderPools[i].MinZoom
		self.renders[i].MaxZoom = poolsCfg.RenderPools[i].MaxZoom
		self.renders[i].Tags = poolsCfg.RenderPools[i].Tags
		self.renders[i].Render, err = NewRenderPool(
			poolsCfg.RenderPools[i].Cmd, poolsCfg.RenderPools[i].PoolSize,
			poolsCfg.RenderPools[i].QueueSize, poolsCfg.RenderPools[i].RenderTTL)
		if err != nil {
			return nil, err
		}
	}

	return self, nil
}

func (self *MultiRenderPool) EnqueueRequest(coord gopnik.TileCoord, resCh chan<- *RenderPoolResponse) error {
RL:
	for _, renderCfg := range self.renders {
		if renderCfg.Tags != nil {
		TL:
			for _, cfgT := range renderCfg.Tags {
				for _, inT := range coord.Tags {
					if inT == cfgT {
						continue TL
					}
				}
				continue RL
			}
		}
		if coord.Zoom < uint64(renderCfg.MinZoom) || coord.Zoom > uint64(renderCfg.MaxZoom) {
			continue
		}
		return renderCfg.Render.EnqueueRequest(coord, resCh)
	}
	return NewBadCoordError(coord)
}

func (self *MultiRenderPool) Reload() {
	for _, r := range self.renders {
		r.Render.Reload()
	}
}

func (self *MultiRenderPool) Stop() {
	for _, rend := range self.renders {
		rend.Render.Stop()
	}
}
