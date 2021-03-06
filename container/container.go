// Copyright 2014 Google Inc. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package container

import "github.com/google/cadvisor/info"

// Listing types.
const (
	LIST_SELF = iota
	LIST_RECURSIVE
)

type ListType int

// SubcontainerEvent types.
const (
	SUBCONTAINER_ADD = iota
	SUBCONTAINER_DELETE
)

type SubcontainerEvent struct {
	// The type of event that occurred.
	EventType int

	// The full container name of the container where the event occurred.
	Name string
}

// Interface for container operation handlers.
type ContainerHandler interface {
	ContainerReference() (info.ContainerReference, error)
	GetSpec() (info.ContainerSpec, error)
	GetStats() (*info.ContainerStats, error)
	ListContainers(listType ListType) ([]info.ContainerReference, error)
	ListThreads(listType ListType) ([]int, error)
	ListProcesses(listType ListType) ([]int, error)
	WatchSubcontainers(events chan SubcontainerEvent) error
}
