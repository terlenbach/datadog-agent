// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build !linux_bpf

package ebpfcheck

import (
	"github.com/DataDog/datadog-agent/pkg/ebpf"
)

// EBPFProbe is not implemented on non-linux systems
type EBPFProbe struct{}

// NewEBPFProbe is not implemented on non-linux systems
func NewEBPFProbe(cfg *ebpf.Config) (*EBPFProbe, error) {
	return nil, ebpf.ErrNotImplemented
}

// Close is not implemented on non-linux systems
func (t *EBPFProbe) Close() {}

// GetAndFlush is not implemented on non-linux systems
func (t *EBPFProbe) GetAndFlush() EBPFStats {
	return EBPFStats{}
}
