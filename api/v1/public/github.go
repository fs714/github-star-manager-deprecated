package public

import (
	"database/sql"
	"net/http"

	"github.com/fs714/github-star-manager/db/sqlite/sqlite_tables"
	"github.com/fs714/github-star-manager/pkg/github_api"
	"github.com/fs714/github-star-manager/pkg/utils/code"
	"github.com/fs714/github-star-manager/pkg/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func SyncFromGithub(c *gin.Context) {
	msg, err := doSyncFromGithub(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code.RespCommonError,
			"msg":    msg,
			"data":   "",
		})

		log.Errorf("failed to sync from github:\n%+v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code.RespOk,
		"msg":    "",
		"data":   "",
	})
}

func doSyncFromGithub(c *gin.Context) (msg string, err error) {
	var postData = struct {
		User string
	}{}
	err = c.ShouldBindJSON(&postData)
	if err != nil {
		msg = "failed to bind post json to struct"
		err = errors.Wrap(err, msg)
		return
	}

	repos, err := github_api.GetStarredRepos(postData.User)
	if err != nil {
		msg = "failed to get starred repos from github"
		err = errors.Wrap(err, msg)
		return
	}

	for _, repo := range repos {
		_, err = sqlite_tables.GetRepoByNameAndUser(*repo.Repository.FullName, postData.User)
		if err != nil {
			if errors.Cause(err) == sql.ErrNoRows {
				err = nil
			} else {
				msg = "failed to get repo from db"
				err = errors.Wrap(err, msg)
				return
			}
		} else {
			continue
		}

		repoDB := sqlite_tables.Repo{
			User:      postData.User,
			CreatedAt: repo.Repository.CreatedAt.Unix(),
			UpdatedAt: repo.Repository.UpdatedAt.Unix(),
			PushedAt:  repo.Repository.PushedAt.Unix(),
		}

		if repo.Repository.FullName != nil {
			repoDB.Name = *repo.Repository.FullName
		}

		if repo.Repository.HTMLURL != nil {
			repoDB.Url = *repo.Repository.HTMLURL
		}

		if repo.Repository.Language != nil {
			repoDB.Language = *repo.Repository.Language
		}

		if repo.Repository.StargazersCount != nil {
			repoDB.StarsCount = *repo.Repository.StargazersCount
		}

		if repo.Repository.ForksCount != nil {
			repoDB.ForksCount = *repo.Repository.ForksCount
		}

		if repo.Repository.Description != nil {
			repoDB.Description = *repo.Repository.Description
		}

		err = sqlite_tables.InsertRepo(repoDB)
		if err != nil {
			msg = "failed to insert repos to db"
			err = errors.Wrap(err, msg)
			return
		}
	}

	return
}
