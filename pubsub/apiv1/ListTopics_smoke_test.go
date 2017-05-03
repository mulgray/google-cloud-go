// Copyright 2017, Google Inc. All rights reserved.
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

// AUTO-GENERATED CODE. DO NOT EDIT.

package pubsub

import (
	pubsubpb "google.golang.org/genproto/googleapis/pubsub/v1"
)

import (
	"testing"

	"cloud.google.com/go/internal/testutil"
	"golang.org/x/net/context"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

var _ = iterator.Done

func TestPublisherSmoke(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping smoke test in short mode")
	}
	ctx := context.Background()
	ts := testutil.TokenSource(ctx, DefaultAuthScopes()...)
	if ts == nil {
		t.Skip("Integration tests skipped. See CONTRIBUTING.md for details")
	}

	projectId := testutil.ProjID()
	uidSpace := testutil.NewUIDSpace("TestPublisherSmoke")
	_, _ = projectId, uidSpace

	c, err := NewPublisherClient(ctx, option.WithTokenSource(ts))
	if err != nil {
		t.Fatal(err)
	}

	var formattedProject string = PublisherProjectPath(projectId)
	var request = &pubsubpb.ListTopicsRequest{
		Project: formattedProject,
	}

	iter := c.ListTopics(ctx, request)
	resp, err := iter.Next()
	if err == iterator.Done {
		t.Log("no items")
	} else if err != nil {
		t.Error(err)
	} else {
		t.Log(resp)
	}
}
