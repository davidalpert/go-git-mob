package msg

type Source int

const (
	// comments from: https://git-scm.com/docs/githooks#_prepare_commit_msg
	UnknownSource  Source = iota
	EmptySource           // commit: source not provided
	MessageSource         // commit-with-message: if a -m or -F option was given
	TemplateSource        // commit-with-template: if a -t option was given or the configuration option commit.template is set
	MergeSource           // merge-commit: if the commit is a merge or a .git/MERGE_MSG file exists
	SquashSource          // squash-commit: if a .git/SQUASH_MSG file exists
	CommitSource          // amending: followed by a commit object name (if a -c, -C or --amend option was given)
)

func CommitMsgSourceFromString(s string) Source {
	switch s {
	case "":
		return EmptySource
	case "message":
		return MessageSource
	case "template":
		return TemplateSource
	case "merge":
		return MergeSource
	case "squash":
		return SquashSource
	case "commit":
		return CommitSource
	}
	return UnknownSource
}

func (s Source) String() string {
	switch s {
	case EmptySource:
		return ""
	case MessageSource:
		return "message"
	case TemplateSource:
		return "template"
	case MergeSource:
		return "merge"
	case SquashSource:
		return "squash"
	case CommitSource:
		return "commit"
	}
	return "unknown"
}
