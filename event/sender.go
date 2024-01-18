package event

import (
	"fmt"
	"time"

	"code.gitea.io/gitea/modules/context"

	"code.gitea.io/gitea/event/infrastructure/kafka"
)

const (
	topicGitClone = "gitea_git_clone"
	topicGitLFS   = "gitea_git_lfs"

	typGitClone = "git_clone"

	fieldRepoName  = "repo_name"
	fieldRepoOwner = "repo_owner"
	fieldDoerEmail = "doer_email"
)

func GitCloneSender(ctx *context.Context) error {
	fmt.Println("GitCloneSender start...")

	msg := MsgNormal{
		Type:    typGitClone,
		User:    ctx.Doer.Name,
		Desc:    "this message means someone clone the repository",
		Details: map[string]string{
			// fieldRepoName:  ctx.Repo.Repository.Name,
			// fieldRepoOwner: ctx.Repo.Repository.OwnerName,
			// fieldDoerEmail: ctx.Doer.Email,
		},
		CreatedAt: time.Now().Unix(),
	}

	return kafka.Publish(topicGitClone, &msg, nil)
}

func GitLFSDownloadSender(msg *MsgNormal) error {
	return kafka.Publish(topicGitLFS, msg, nil)
}
