// Code generated by protoc-gen-solo-kit. DO NOT EDIT.

package v1

import (
	
	"go.uber.org/zap"
	"github.com/mitchellh/hashstructure"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources"
	"github.com/solo-io/solo-kit/pkg/api/v1/resources/core"
)

type ObservabilitySnapshot struct {
	Prometheusconfigs PrometheusconfigsByNamespace
}

func (s ObservabilitySnapshot) Clone() ObservabilitySnapshot {
	return ObservabilitySnapshot{
		Prometheusconfigs: s.Prometheusconfigs.Clone(),
	}
}

func (s ObservabilitySnapshot) snapshotToHash() ObservabilitySnapshot {
	snapshotForHashing := s.Clone()
	for _, config := range snapshotForHashing.Prometheusconfigs.List() {
		resources.UpdateMetadata(config, func(meta *core.Metadata) {
			meta.ResourceVersion = ""
		})
	}

	return snapshotForHashing
}

func (s ObservabilitySnapshot) Hash() uint64 {
	return s.hashStruct(s.snapshotToHash())
 }

 func (s ObservabilitySnapshot) HashFields() []zap.Field {
	snapshotForHashing := s.snapshotToHash()
	var fields []zap.Field
	prometheusconfigs := s.hashStruct(snapshotForHashing.Prometheusconfigs.List())
	fields = append(fields, zap.Uint64("prometheusconfigs", prometheusconfigs ))

	return append(fields, zap.Uint64("snapshotHash",  s.hashStruct(snapshotForHashing)))
 }
 
func (s ObservabilitySnapshot) hashStruct(v interface{}) uint64 {
	h, err := hashstructure.Hash(v, nil)
	 if err != nil {
		 panic(err)
	 }
	 return h
 }


