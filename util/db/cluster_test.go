package db

import (
	"context"
	"sync"
	"testing"
	"time"

	log "github.com/sirupsen/logrus"

	"github.com/stretchr/testify/assert"
	"k8s.io/client-go/kubernetes/fake"

	"github.com/argoproj/argo-cd/common"
	"github.com/argoproj/argo-cd/pkg/apis/application/v1alpha1"
	"github.com/argoproj/argo-cd/util/settings"
)

const (
	fakeNamespace = "fake-ns"
	syncMessage   = "Sync successful"
	debugFlag     = true
)

func Test_serverToSecretName(t *testing.T) {
	name, err := serverToSecretName("http://foo")
	assert.NoError(t, err)
	assert.Equal(t, "cluster-foo-752281925", name)
}

func TestWatchClustersNoClustersRegistered(t *testing.T) {
	kubeclientset := fake.NewSimpleClientset()
	settingsManager := settings.NewSettingsManager(context.Background(), kubeclientset, fakeNamespace)
	db := NewDB(fakeNamespace, settingsManager, kubeclientset)
	timeout := time.Second * 5

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	addedClusters := make(chan *v1alpha1.Cluster)

	go func() {
		assert.NoError(t, db.WatchClusters(ctx, func(cluster *v1alpha1.Cluster) {
			addedClusters <- cluster
		}, func(oldCluster *v1alpha1.Cluster, newCluster *v1alpha1.Cluster) {
			assert.Fail(t, "no cluster modifications expected")
		}, func(clusterServer string) {
			assert.Fail(t, "no cluster removals expected")
		}))
	}()

	select {
	case addedCluster := <-addedClusters:
		assert.Equal(t, addedCluster.Server, common.KubernetesInternalAPIServerAddr)
	case <-time.After(timeout):
		assert.Fail(t, "no cluster event within timeout")
	}
}

func TestWatchClusters(t *testing.T) {
	const cluserServerAddr = "http://minikube"

	kubeclientset := fake.NewSimpleClientset()
	settingsManager := settings.NewSettingsManager(context.Background(), kubeclientset, fakeNamespace)
	db := NewDB(fakeNamespace, settingsManager, kubeclientset)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	addedClusters := make([]string, 0)
	updatedClusters := make([]string, 0)
	deletedClusters := make([]string, 0)

	var wg sync.WaitGroup
	wg.Add(4)

	go func() {
		assert.NoError(t, db.WatchClusters(ctx, func(cluster *v1alpha1.Cluster) {
			addedClusters = append(addedClusters, cluster.Server)
			wg.Done()
		}, func(oldCluster *v1alpha1.Cluster, newCluster *v1alpha1.Cluster) {
			updatedClusters = append(updatedClusters, newCluster.Server)
			assert.Equal(t, syncMessage, newCluster.ConnectionState.Message)
			assert.Empty(t, oldCluster.ConnectionState.Message)
			wg.Done()
		}, func(clusterServer string) {
			deletedClusters = append(deletedClusters, clusterServer)
			wg.Done()
		}))
	}()

	err := crudCluster(ctx, db, cluserServerAddr, syncMessage)
	assert.NoError(t, err, "Test prepare test data crdCluster failed")

	wg.Wait()

	assert.ElementsMatch(t, []string{common.KubernetesInternalAPIServerAddr, cluserServerAddr}, addedClusters)
	assert.ElementsMatch(t, []string{"http://minikube"}, updatedClusters)
	assert.ElementsMatch(t, []string{"http://minikube"}, deletedClusters)
}

//Cluster with address common.KubernetesInternalAPIServerAddr is local cluster
//In this test we crud local cluster
func TestWatchClustersLocalCluster(t *testing.T) {
	kubeclientset := fake.NewSimpleClientset()
	settingsManager := settings.NewSettingsManager(context.Background(), kubeclientset, fakeNamespace)
	db := NewDB(fakeNamespace, settingsManager, kubeclientset)

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	addedClusters := make([]string, 0)
	updatedClusters := make([]string, 0)

	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		assert.NoError(t, db.WatchClusters(ctx, func(cluster *v1alpha1.Cluster) {
			if debugFlag {
				log.Info("called adding event for local cluster")
			}
			addedClusters = append(addedClusters, cluster.Server)
		}, func(oldCluster *v1alpha1.Cluster, newCluster *v1alpha1.Cluster) {
			if debugFlag {
				log.Info("called update event for local cluster")
			}
			updatedClusters = append(updatedClusters, newCluster.Server)
			wg.Done()
		}, func(clusterServer string) {
			assert.Fail(t, "Not expecting delete for local cluster")
		}))
	}()

	//crud local cluster
	err := crudCluster(ctx, db, common.KubernetesInternalAPIServerAddr, syncMessage)
	assert.NoError(t, err, "Test prepare test data crdCluster failed")

	wg.Wait()

	assert.ElementsMatch(t, []string{common.KubernetesInternalAPIServerAddr}, addedClusters)
	assert.ElementsMatch(t, []string{common.KubernetesInternalAPIServerAddr, common.KubernetesInternalAPIServerAddr}, updatedClusters)
}

func crudCluster(ctx context.Context, db ArgoDB, cluserServerAddr string, message string) error {
	if debugFlag {
		log.Info("starting create local cluster")
	}
	cluster, err := db.CreateCluster(ctx, &v1alpha1.Cluster{Server: cluserServerAddr})
	if err != nil {
		return err
	}
	if debugFlag {
		log.Info("starting create local cluster - done")
	}

	if debugFlag {
		log.Info("starting update local cluster")
	}
	cluster.ConnectionState.Message = message
	cluster, err = db.UpdateCluster(ctx, cluster)
	if err != nil {
		return err
	}
	if debugFlag {
		log.Info("starting update local cluster - done")
	}

	if debugFlag {
		log.Info("starting delete local cluster")
	}
	err = db.DeleteCluster(ctx, cluster.Server)
	if err != nil {
		return err
	}
	return err
	if debugFlag {
		log.Info("starting delete local cluster - done")
	}
}
