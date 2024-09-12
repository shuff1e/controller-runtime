package main

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/kubectl/pkg/scheme"
	api "sigs.k8s.io/controller-runtime/examples/crd/pkg"
	"sigs.k8s.io/controller-runtime/pkg/client/fake"
)

func TestFake(t *testing.T) {
	s := scheme.Scheme

	s.AddKnownTypes(api.SchemeGroupVersion, &api.ChaosPod{}, &api.ChaosPodList{})

	chaosPod := &api.ChaosPod{
		ObjectMeta: metav1.ObjectMeta{
			Namespace: "default",
			Name: "chaos-pod-1",
		},
		Spec: api.ChaosPodSpec{
			NextStop: metav1.Time{
				time.Unix(123,0),
			},
		},
	}

	chaosPodList := &api.ChaosPodList{
		Items: []api.ChaosPod{
			{
				ObjectMeta: metav1.ObjectMeta{
					Namespace: "default",
					Name: "chaos-pod-list-1",
				},
				Spec: api.ChaosPodSpec{
					NextStop: metav1.Time{
						time.Unix(456,0),
					},
				},
			},
		},
	}

	fakeClient := fake.NewClientBuilder().
		WithScheme(s).
		WithLists(chaosPodList).
		WithObjects(chaosPod).
		Build()

	ret := &api.ChaosPod{}
	err := fakeClient.Get(context.TODO(), types.NamespacedName{
		Namespace: "default",
		Name:      "chaos-pod-1",
	}, ret)
	assert.Nil(t, err)
	assert.Equal(t, 123, int(ret.Spec.NextStop.Unix()))

	ret = &api.ChaosPod{}
	err = fakeClient.Get(context.TODO(), types.NamespacedName{
		Namespace: "default",
		Name:      "chaos-pod-list-1",
	}, ret)
	assert.Nil(t, err)
	assert.Equal(t, 456, int(ret.Spec.NextStop.Unix()))
}