package subscription

import (
	"fmt"
	"net/url"
	"time"

	kymaeventingv1alpha1 "github.com/kyma-project/kyma/components/eventing-controller/api/v1alpha1"
	"github.com/kyma-project/kyma/tests/function-controller/pkg/helpers"
	"github.com/kyma-project/kyma/tests/function-controller/pkg/resource"
	"github.com/kyma-project/kyma/tests/function-controller/pkg/shared"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"

	k8sruntime "k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"k8s.io/client-go/util/retry"
)

type Subscription struct {
	resCli      *resource.Resource
	name        string
	namespace   string
	waitTimeout time.Duration
	log         *logrus.Entry
	verbose     bool
}

func New(name string, c shared.Container) *Subscription {
	gvr := schema.GroupVersionResource{
		Group:    "eventing.kyma-project.io",
		Version:  "v1alpha1",
		Resource: "subscriptions",
	}

	return &Subscription{
		resCli:      resource.New(c.DynamicCli, gvr, c.Namespace, c.Log, c.Verbose),
		name:        name,
		namespace:   c.Namespace,
		waitTimeout: c.WaitTimeout,
		log:         c.Log,
		verbose:     c.Verbose,
	}
}

func (s *Subscription) Create(sinkUrl *url.URL) (subscription *kymaeventingv1alpha1.Subscription, err error) {
	subscription = &kymaeventingv1alpha1.Subscription{
		ObjectMeta: metav1.ObjectMeta{
			Name:      s.name,
			Namespace: s.namespace,
		},
		Spec: kymaeventingv1alpha1.SubscriptionSpec{
			Protocol:         "",
			ProtocolSettings: &kymaeventingv1alpha1.ProtocolSettings{},
			Sink:             sinkUrl.String(),
		},
	}

	unstructuredSub, err := toUnstructured(subscription)
	if err != nil {
		return subscription, errors.Wrapf(err, "while creating Subscription %s", s.name)
	}

	_, err = s.resCli.Create(&unstructuredSub)
	if err != nil {
		return subscription, errors.Wrapf(err, "while creating Subscription %s in namespace %s", s.name, s.namespace)
	}
	return subscription, err
}

func (s *Subscription) Delete() error {
	err := s.resCli.Delete(s.name)
	if err != nil {
		return errors.Wrapf(err, "while deleting Subscription %s in namespace %s", s.name, s.namespace)
	}

	return nil
}

func (s *Subscription) Get() (*unstructured.Unstructured, error) {
	subscription, err := s.resCli.Get(s.name)
	if err != nil {
		return &unstructured.Unstructured{}, errors.Wrapf(err, "while getting Subscription %s in namespace %s", s.name, s.namespace)
	}

	return subscription, nil
}

func (s *Subscription) WaitForStatusRunning(subscription *kymaeventingv1alpha1.Subscription) error {
	return retry.OnError(retry.DefaultBackoff, func(err error) bool {
		if err != nil {
			return true
		}
		return false
	}, func() error {
		if subscription.Status.Ready {
			return nil
		}
		return fmt.Errorf("Subscription: %s is not ready yet, retrying", subscription.GetName())
	})
}

func (s *Subscription) LogResource() error {
	subscription, err := s.Get()
	if err != nil {
		return err
	}

	out, err := helpers.PrettyMarshall(subscription)
	if err != nil {
		return err
	}

	s.log.Infof("%s", out)
	return nil
}

// not used, can be removed
func toSubscription(unstructuredSub *unstructured.Unstructured) (*kymaeventingv1alpha1.Subscription, error) {
	subscription := new(kymaeventingv1alpha1.Subscription)
	err := k8sruntime.DefaultUnstructuredConverter.FromUnstructured(unstructuredSub.Object, subscription)
	if err != nil {
		return nil, err
	}
	return subscription, nil
}

func toUnstructured(sub *kymaeventingv1alpha1.Subscription) (*unstructured.Unstructured, error) {
	object, err := k8sruntime.DefaultUnstructuredConverter.ToUnstructured(&sub)
	if err != nil {
		return nil, err
	}
	return &unstructured.Unstructured{Object: object}, nil
}
