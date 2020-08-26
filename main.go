package main

import (
	"context"
	"fmt"

	"github.com/argoproj/argo-cd/pkg/apiclient"
	"github.com/argoproj/argo-cd/pkg/apiclient/application"
	"github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func main() {
	clientset := apiclient.NewClientOrDie(&apiclient.ClientOptions{Insecure: true, PlainText: true, ServerAddr: "localhost:8080"})
	closer, appClient := clientset.NewApplicationClientOrDie()
	defer closer.Close()
	name := "guestbook"
	_, err := appClient.Get(context.Background(), &application.ApplicationQuery{Name: &name})
	if err == nil {
		println("app exists")
		return
	}

	_, err = appClient.Create(context.Background(), &application.ApplicationCreateRequest{Application: v1alpha1.Application{
		ObjectMeta: v1.ObjectMeta{
			Name: "guestbook",
		},
		Spec: v1alpha1.ApplicationSpec{
			Source: v1alpha1.ApplicationSource{
				RepoURL: "https://github.com/argoproj/argocd-example-apps",
				Path:    "guestbook",
			},
			Destination: v1alpha1.ApplicationDestination{
				Namespace: "default",
				Server:    "https://kubernetes.default.svc",
			},
		},
	}})
	if err != nil {
		println(fmt.Sprintf("cannot create app: %v", err))
	}
}
