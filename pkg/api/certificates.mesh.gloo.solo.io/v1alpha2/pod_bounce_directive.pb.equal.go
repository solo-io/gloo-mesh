// Code generated by protoc-gen-ext. DO NOT EDIT.
// source: github.com/solo-io/gloo-mesh/api/certificates/pod_bounce_directive.proto

package v1alpha2

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"strings"

	"github.com/golang/protobuf/proto"
	equality "github.com/solo-io/protoc-gen-ext/pkg/equality"
)

// ensure the imports are used
var (
	_ = errors.New("")
	_ = fmt.Print
	_ = binary.LittleEndian
	_ = bytes.Compare
	_ = strings.Compare
	_ = equality.Equalizer(nil)
	_ = proto.Message(nil)
)

// Equal function
func (m *PodBounceDirectiveSpec) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PodBounceDirectiveSpec)
	if !ok {
		that2, ok := that.(PodBounceDirectiveSpec)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetPodsToBounce()) != len(target.GetPodsToBounce()) {
		return false
	}
	for idx, v := range m.GetPodsToBounce() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPodsToBounce()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPodsToBounce()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *PodBounceDirectiveStatus) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PodBounceDirectiveStatus)
	if !ok {
		that2, ok := that.(PodBounceDirectiveStatus)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetPodsBounced()) != len(target.GetPodsBounced()) {
		return false
	}
	for idx, v := range m.GetPodsBounced() {

		if h, ok := interface{}(v).(equality.Equalizer); ok {
			if !h.Equal(target.GetPodsBounced()[idx]) {
				return false
			}
		} else {
			if !proto.Equal(v, target.GetPodsBounced()[idx]) {
				return false
			}
		}

	}

	return true
}

// Equal function
func (m *PodBounceDirectiveSpec_PodSelector) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PodBounceDirectiveSpec_PodSelector)
	if !ok {
		that2, ok := that.(PodBounceDirectiveSpec_PodSelector)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if strings.Compare(m.GetNamespace(), target.GetNamespace()) != 0 {
		return false
	}

	if len(m.GetLabels()) != len(target.GetLabels()) {
		return false
	}
	for k, v := range m.GetLabels() {

		if strings.Compare(v, target.GetLabels()[k]) != 0 {
			return false
		}

	}

	if m.GetWaitForReplicas() != target.GetWaitForReplicas() {
		return false
	}

	return true
}

// Equal function
func (m *PodBounceDirectiveStatus_BouncedPodSet) Equal(that interface{}) bool {
	if that == nil {
		return m == nil
	}

	target, ok := that.(*PodBounceDirectiveStatus_BouncedPodSet)
	if !ok {
		that2, ok := that.(PodBounceDirectiveStatus_BouncedPodSet)
		if ok {
			target = &that2
		} else {
			return false
		}
	}
	if target == nil {
		return m == nil
	} else if m == nil {
		return false
	}

	if len(m.GetBouncedPods()) != len(target.GetBouncedPods()) {
		return false
	}
	for idx, v := range m.GetBouncedPods() {

		if strings.Compare(v, target.GetBouncedPods()[idx]) != 0 {
			return false
		}

	}

	return true
}
