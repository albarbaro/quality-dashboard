package server

import (
	"context"
	"time"

	gh "github.com/google/go-github/v44/github"
	"github.com/redhat-appstudio/quality-studio/api/apis/codecov"
	"github.com/redhat-appstudio/quality-studio/pkg/storage"
	"go.uber.org/zap"
)

// rotationStrategy describes a strategy for generating server configuration from a file.
type rotationStrategy struct {
	// Time between rotations.
	rotationFrequency time.Duration
}

// startUpdateCache begins repo information rotation in a new goroutine, closing once the context is canceled.
func (s *Server) startUpdateStorage(ctx context.Context, strategy rotationStrategy, now func() time.Time) {
	go func() {
		// Try to rotate immediately so properly configured repositories.
		if err := s.rotate(); err != nil {
			s.cfg.Logger.Sugar().Infof("Update failed: %v", err)
		}
		for {
			select {
			case <-ctx.Done():
				return
			case <-time.After(strategy.rotationFrequency):
				if err := s.rotate(); err != nil {
					s.cfg.Logger.Sugar().Infof("failed to update cache: %v", err)
				}
			}
		}
	}()
}

func (s *Server) rotate() error {
	err := s.CacheRepositoriesInformation()
	if err != nil {
		s.cfg.Logger.Sugar().Errorf("Failed to update cache", zap.Error(err))
		return err
	}

	return nil
}

func staticRotationStrategy() rotationStrategy {
	return rotationStrategy{
		// Setting these values to 30 Minutes is easier than having a flag indicating no rotation.
		rotationFrequency: time.Minute * 30,
	}
}

func (s *Server) CacheRepositoriesInformation() error {
	storageRepos, err := s.cfg.Storage.ListRepositories()
	if err != nil {
		return err
	}

	for _, repo := range storageRepos {
		workf, err := s.getGithubWorkflows(repo.GitOrganization, repo.RepositoryName)
		if err != nil {
			return err
		}

		for _, w := range workf.Workflows {
			err = s.cfg.Storage.ReCreateWorkflow(storage.GithubWorkflows{
				WorkflowName: w.GetName(),
				BadgeURL:     w.GetBadgeURL(),
				HTMLURL:      w.GetHTMLURL(),
				JobURL:       w.GetURL(),
				State:        w.GetState(),
			}, repo.RepositoryName)
			if err != nil {
				return err
			}
		}

		coverage, err := s.getCodeCoverage(repo.GitOrganization, repo.RepositoryName)
		if err != nil {
			return err
		}
		totalCoverageConverted, _ := coverage.Commit.Totals.TotalCoverage.Float64()
		err = s.cfg.Storage.UpdateCoverage(storage.Coverage{
			GitOrganization:    repo.GitOrganization,
			RepositoryName:     repo.RepositoryName,
			CoveragePercentage: totalCoverageConverted,
		}, repo.RepositoryName)
		if err != nil {
			return err
		}
	}
	s.cfg.Logger.Info("Successfully updated the storage data")

	return nil
}

func (s *Server) getGithubWorkflows(gitOrganization string, repoName string) (*gh.Workflows, error) {
	return s.cfg.Github.GetRepositoryWorkflows(gitOrganization, repoName)
}

func (s *Server) getCodeCoverage(gitOrganization string, repoName string) (codecov.GitHubTagResponse, error) {
	return s.cfg.CodeCov.GetCodeCovInfo(gitOrganization, repoName)
}
