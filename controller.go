/*
Copyright 2018 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"

	"github.com/go-logr/logr"

	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

type reconcilePod struct {
	client client.Client
	log    logr.Logger
}

var _ reconcile.Reconciler = &reconcilePod{}

func (r *reconcilePod) Reconcile(request reconcile.Request) (reconcile.Result, error) {
	// set up a convinient log object so we don't have to type request over and over again
	log := r.log.WithValues("request", request)

	// Fetch the Pod from the cache
	rs := &corev1.Pod{}
	err := r.client.Get(context.TODO(), request.NamespacedName, rs)
	if errors.IsNotFound(err) {
		log.Error(nil, "Could not find Pod")
		return reconcile.Result{}, nil
	}

	if err != nil {
		log.Error(err, "Could not fetch Pod")
		return reconcile.Result{}, err
	}

	// Print the Pod
	log.Info("Reconciling Pod", "pod name", rs.ObjectMeta.Name)

	return reconcile.Result{}, nil
}
