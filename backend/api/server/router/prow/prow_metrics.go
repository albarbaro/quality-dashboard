package prow

import (
	"context"
	"net/http"

	"github.com/redhat-appstudio/quality-studio/api/types"
	"github.com/redhat-appstudio/quality-studio/pkg/utils/httputils"
)

// version godoc
// @Summary Prow Jobs info
// @Description returns all prow jobs related to git_organization and repository_name
// @Tags Prow Jobs info
// @Accept json
// @Produce json
// @Param repository body GitRepositoryRequest true "repository name"
// @Param organization body GitRepositoryRequest true "git_organization"
// @Router /prow/results/get [get]
// @Success 200 {Object} []db.Prow
// @Failure 400 {object} types.ErrorResponse
func (s *jobRouter) getProwMetrics(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	repositoryName := r.URL.Query()["repository_name"]
	gitOrgazanitation := r.URL.Query()["git_organization"]
	jobType := r.URL.Query()["job_type"]

	if len(repositoryName) == 0 {
		return httputils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{
			Message:    "repository_name value not present in query",
			StatusCode: 400,
		})
	} else if len(gitOrgazanitation) == 0 {
		return httputils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{
			Message:    "git_organization value not present in query",
			StatusCode: 400,
		})
	} else if len(jobType) == 0 {
		return httputils.WriteJSON(w, http.StatusBadRequest, types.ErrorResponse{
			Message:    "job_type value not present in query",
			StatusCode: 400,
		})
	}
	return httputils.WriteJSON(w, http.StatusOK, s.Storage.GetMetrics(gitOrgazanitation[0], repositoryName[0], jobType[0]))
}
