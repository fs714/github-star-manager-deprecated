package public

import (
	"net/http"

	"github.com/fs714/github-star-manager/db/sqlite/sqlite_tables"
	"github.com/fs714/github-star-manager/pkg/utils/code"
	"github.com/fs714/github-star-manager/pkg/utils/log"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func GetRepos(c *gin.Context) {
	repos, msg, err := doGetRepos(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"status": code.RespCommonError,
			"msg":    msg,
			"data":   "",
		})

		log.Errorf("failed to get repos:\n%+v", err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": code.RespOk,
		"msg":    "",
		"data":   repos,
	})
}

func doGetRepos(c *gin.Context) (repos []sqlite_tables.Repo, msg string, err error) {
	user := c.Query("user")
	if user == "" {
		msg = "user is required"
		err = errors.New(msg)
		return
	}

	repos, err = sqlite_tables.GetReposByUser(user)
	if err != nil {
		msg = "failed to get repos from database"
		err = errors.Wrap(err, msg)
		return
	}

	return
}
