package plugin

import (
	"context"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/argoproj/gitops-engine/pkg/utils/kube"
	"github.com/argoproj/pkg/sync"
	"github.com/mattn/go-zglob"

	"github.com/argoproj/argo-cd/v2/cmpserver/apiclient"
	executil "github.com/argoproj/argo-cd/v2/util/exec"
	kubeutil "github.com/argoproj/argo-cd/v2/util/kube"
)

// Service implements ConfigManagementPluginService interface
type Service struct {
	initConstants CMPServerInitConstants
}

type CMPServerInitConstants struct {
	ConfigFilePath string
	PluginConfig   PluginConfig
}

// NewService returns a new instance of the ConfigManagementPluginService
func NewService(initConstants CMPServerInitConstants) *Service {
	return &Service{
		initConstants: initConstants,
	}
}

var manifestGenerateLock = sync.NewKeyLock()

func runCommand(command Command, path string, env []string) (string, error) {
	if len(command.Command) == 0 {
		return "", fmt.Errorf("Command is empty")
	}
	cmd := exec.Command(command.Command[0], append(command.Command[1:], command.Args...)...)
	cmd.Env = env
	cmd.Dir = path
	return executil.Run(cmd)
}

// Environ returns a list of environment variables in name=value format from a list of variables
func environ(envVars []*apiclient.EnvEntry) []string {
	var environ []string
	for _, item := range envVars {
		if item != nil && item.Name != "" && item.Value != "" {
			environ = append(environ, fmt.Sprintf("%s=%s", item.Name, item.Value))
		}
	}
	return environ
}

// GenerateManifest runs generate command from plugin config file and returns generated manifest files
func (s *Service) GenerateManifest(ctx context.Context, q *apiclient.ManifestRequest) (*apiclient.ManifestResponse, error) {
	config := s.initConstants.PluginConfig

	// Plugins can request to lock the complete repository when they need to
	// use git client operations.
	if config.Spec.LockRepo {
		manifestGenerateLock.Lock(q.RepoPath)
		defer manifestGenerateLock.Unlock(q.RepoPath)
	} else if !config.Spec.AllowConcurrency {
		manifestGenerateLock.Lock(q.AppPath)
		defer manifestGenerateLock.Unlock(q.AppPath)
	}

	if err := ValidatePluginConfig(config); err != nil {
		return &apiclient.ManifestResponse{}, fmt.Errorf("cmp server should have valid plugin config file, instead %v", err)
	}

	env := append(os.Environ(), environ(q.Env)...)
	if len(config.Spec.Init.Command) > 0 {
		_, err := runCommand(config.Spec.Init, q.AppPath, env)
		if err != nil {
			return &apiclient.ManifestResponse{}, err
		}
	}

	out, err := runCommand(config.Spec.Generate, q.AppPath, env)
	if err != nil {
		return &apiclient.ManifestResponse{}, err
	}

	targetObjs, err := kube.SplitYAML([]byte(out))
	if err != nil {
		return &apiclient.ManifestResponse{}, err
	}

	manifests, err := kubeutil.FromUnstructured(targetObjs)
	return &apiclient.ManifestResponse{
		Manifests: manifests,
	}, err
}

// MatchRepository checks whether the application repository type is supported by config management plugin server
func (s *Service) MatchRepository(ctx context.Context, q *apiclient.RepositoryRequest) (*apiclient.RepositoryResponse, error) {
	var repoResponse apiclient.RepositoryResponse
	config := s.initConstants.PluginConfig
	if config.Spec.Find.Glob != "" {
		pattern := strings.TrimSuffix(q.Path, "/") + "/" + strings.TrimPrefix(config.Spec.Find.Glob, "/")
		// filepath.Glob doesn't have '**' support hence selecting third-party lib
		// https://github.com/golang/go/issues/11862
		matches, err := zglob.Glob(pattern)
		if err != nil || len(matches) == 0 {
			return &repoResponse, err
		} else if len(matches) > 0 {
			repoResponse.IsSupported = true
			return &repoResponse, nil
		}
	}

	find, err := runCommand(config.Spec.Find, q.Path, os.Environ())
	if err != nil {
		return &repoResponse, err
	}

	var isSupported bool
	if find != "" {
		isSupported = true
	}
	return &apiclient.RepositoryResponse{
		IsSupported: isSupported,
	}, nil
}