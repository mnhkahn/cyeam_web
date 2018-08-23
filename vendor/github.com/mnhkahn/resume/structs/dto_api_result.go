package structs

type Params struct {
	Output        string
	GitHub        string
	RepoLimit     int
	TouTiao       string
	TouTiaoLimit  int
	StackOverflow string
	Weixin        string
}

type ResumeResult struct {
	Profile              *Profile
	TouTiao              *TouTiao
	Weixin               string
	HotArticles          []*Article
	Repos                *Repos
	StackOverflowProfile *StackOverflowProfile
	StackOverflowTags    *StackOverflowTags
	Answers              []*AnswerResult
}

type TouTiao struct {
	Name string
	ID   string
}

type Article struct {
	Title string
	Url   string
}

type Profile struct {
	AvatarURL         string `json:"avatar_url"`
	Bio               string `json:"bio"`
	Blog              string `json:"blog"`
	Company           string `json:"company"`
	CreatedAt         string `json:"created_at"`
	Email             string `json:"email"`
	EventsURL         string `json:"events_url"`
	Followers         int    `json:"followers"`
	FollowersURL      string `json:"followers_url"`
	Following         int    `json:"following"`
	FollowingURL      string `json:"following_url"`
	GistsURL          string `json:"gists_url"`
	GravatarID        string `json:"gravatar_id"`
	Hireable          bool   `json:"hireable"`
	HTMLURL           string `json:"html_url"`
	ID                int    `json:"id"`
	Location          string `json:"location"`
	Login             string `json:"login"`
	Name              string `json:"name"`
	OrganizationsURL  string `json:"organizations_url"`
	PublicGists       int    `json:"public_gists"`
	PublicRepos       int    `json:"public_repos"`
	ReceivedEventsURL string `json:"received_events_url"`
	ReposURL          string `json:"repos_url"`
	SiteAdmin         bool   `json:"site_admin"`
	StarredURL        string `json:"starred_url"`
	SubscriptionsURL  string `json:"subscriptions_url"`
	Type              string `json:"type"`
	UpdatedAt         string `json:"updated_at"`
	URL               string `json:"url"`
}

type Repos struct {
	IncompleteResults bool    `json:"incomplete_results"`
	Items             []*Item `json:"items"`
	TotalCount        int     `json:"total_count"`
}

type Item struct {
	ArchiveURL       string      `json:"archive_url"`
	AssigneesURL     string      `json:"assignees_url"`
	BlobsURL         string      `json:"blobs_url"`
	BranchesURL      string      `json:"branches_url"`
	CloneURL         string      `json:"clone_url"`
	CollaboratorsURL string      `json:"collaborators_url"`
	CommentsURL      string      `json:"comments_url"`
	CommitsURL       string      `json:"commits_url"`
	CompareURL       string      `json:"compare_url"`
	ContentsURL      string      `json:"contents_url"`
	ContributorsURL  string      `json:"contributors_url"`
	CreatedAt        string      `json:"created_at"`
	DefaultBranch    string      `json:"default_branch"`
	DeploymentsURL   string      `json:"deployments_url"`
	Description      interface{} `json:"description"`
	DownloadsURL     string      `json:"downloads_url"`
	EventsURL        string      `json:"events_url"`
	Fork             bool        `json:"fork"`
	Forks            int         `json:"forks"`
	ForksCount       int         `json:"forks_count"`
	ForksURL         string      `json:"forks_url"`
	FullName         string      `json:"full_name"`
	GitCommitsURL    string      `json:"git_commits_url"`
	GitRefsURL       string      `json:"git_refs_url"`
	GitTagsURL       string      `json:"git_tags_url"`
	GitURL           string      `json:"git_url"`
	HasDownloads     bool        `json:"has_downloads"`
	HasIssues        bool        `json:"has_issues"`
	HasPages         bool        `json:"has_pages"`
	HasWiki          bool        `json:"has_wiki"`
	Homepage         interface{} `json:"homepage"`
	HooksURL         string      `json:"hooks_url"`
	HTMLURL          string      `json:"html_url"`
	ID               int         `json:"id"`
	IssueCommentURL  string      `json:"issue_comment_url"`
	IssueEventsURL   string      `json:"issue_events_url"`
	IssuesURL        string      `json:"issues_url"`
	KeysURL          string      `json:"keys_url"`
	LabelsURL        string      `json:"labels_url"`
	Language         string      `json:"language"`
	LanguagesURL     string      `json:"languages_url"`
	MergesURL        string      `json:"merges_url"`
	MilestonesURL    string      `json:"milestones_url"`
	MirrorURL        interface{} `json:"mirror_url"`
	Name             string      `json:"name"`
	NotificationsURL string      `json:"notifications_url"`
	// OpenIssues       int         `json:"open_issues"`
	// OpenIssuesCount  int         `json:"open_issues_count"`
	Owner struct {
		AvatarURL    string `json:"avatar_url"`
		EventsURL    string `json:"events_url"`
		FollowersURL string `json:"followers_url"`
		FollowingURL string `json:"following_url"`
		GistsURL     string `json:"gists_url"`
		GravatarID   string `json:"gravatar_id"`
		HTMLURL      string `json:"html_url"`
		// ID                int    `json:"id"`
		Login             string `json:"login"`
		OrganizationsURL  string `json:"organizations_url"`
		ReceivedEventsURL string `json:"received_events_url"`
		ReposURL          string `json:"repos_url"`
		SiteAdmin         bool   `json:"site_admin"`
		StarredURL        string `json:"starred_url"`
		SubscriptionsURL  string `json:"subscriptions_url"`
		Type              string `json:"type"`
		URL               string `json:"url"`
	} `json:"owner"`
	Private     bool   `json:"private"`
	PullsURL    string `json:"pulls_url"`
	PushedAt    string `json:"pushed_at"`
	ReleasesURL string `json:"releases_url"`
	// Score           int    `json:"score"`
	// Size            int    `json:"size"`
	SSHURL string `json:"ssh_url"`
	// StargazersCount int    `json:"stargazers_count"`
	StargazersURL   string `json:"stargazers_url"`
	StatusesURL     string `json:"statuses_url"`
	SubscribersURL  string `json:"subscribers_url"`
	SubscriptionURL string `json:"subscription_url"`
	SvnURL          string `json:"svn_url"`
	TagsURL         string `json:"tags_url"`
	TeamsURL        string `json:"teams_url"`
	TreesURL        string `json:"trees_url"`
	UpdatedAt       string `json:"updated_at"`
	URL             string `json:"url"`
	// Watchers        int    `json:"watchers"`
	// WatchersCount   int    `json:"watchers_count"`
}

type StackOverflowProfile struct {
	HasMore bool `json:"has_more"`
	Items   []struct {
		AcceptRate  int `json:"accept_rate"`
		AccountID   int `json:"account_id"`
		Age         int `json:"age"`
		BadgeCounts struct {
			Bronze int `json:"bronze"`
			Gold   int `json:"gold"`
			Silver int `json:"silver"`
		} `json:"badge_counts"`
		CreationDate            int    `json:"creation_date"`
		DisplayName             string `json:"display_name"`
		IsEmployee              bool   `json:"is_employee"`
		LastAccessDate          int    `json:"last_access_date"`
		LastModifiedDate        int    `json:"last_modified_date"`
		Link                    string `json:"link"`
		Location                string `json:"location"`
		ProfileImage            string `json:"profile_image"`
		Reputation              int    `json:"reputation"`
		ReputationChangeDay     int    `json:"reputation_change_day"`
		ReputationChangeMonth   int    `json:"reputation_change_month"`
		ReputationChangeQuarter int    `json:"reputation_change_quarter"`
		ReputationChangeWeek    int    `json:"reputation_change_week"`
		ReputationChangeYear    int    `json:"reputation_change_year"`
		UserID                  int    `json:"user_id"`
		UserType                string `json:"user_type"`
		WebsiteURL              string `json:"website_url"`
	} `json:"items"`
	QuotaMax       int `json:"quota_max"`
	QuotaRemaining int `json:"quota_remaining"`
}

type StackOverflowTags struct {
	HasMore bool `json:"has_more"`
	Items   []struct {
		Count           int    `json:"count"`
		HasSynonyms     bool   `json:"has_synonyms"`
		IsModeratorOnly bool   `json:"is_moderator_only"`
		IsRequired      bool   `json:"is_required"`
		Name            string `json:"name"`
		UserID          int    `json:"user_id"`
	} `json:"items"`
	QuotaMax       int `json:"quota_max"`
	QuotaRemaining int `json:"quota_remaining"`
}

type StackOverflowAnswers struct {
	HasMore bool `json:"has_more"`
	Items   []struct {
		AnswerID         int  `json:"answer_id"`
		CreationDate     int  `json:"creation_date"`
		IsAccepted       bool `json:"is_accepted"`
		LastActivityDate int  `json:"last_activity_date"`
		Owner            struct {
			AcceptRate   int    `json:"accept_rate"`
			DisplayName  string `json:"display_name"`
			Link         string `json:"link"`
			ProfileImage string `json:"profile_image"`
			Reputation   int    `json:"reputation"`
			UserID       int    `json:"user_id"`
			UserType     string `json:"user_type"`
		} `json:"owner"`
		QuestionID int `json:"question_id"`
		Score      int `json:"score"`
	} `json:"items"`
	QuotaMax       int `json:"quota_max"`
	QuotaRemaining int `json:"quota_remaining"`
}

type StackOverflowQuestions struct {
	HasMore bool `json:"has_more"`
	Items   []struct {
		AcceptedAnswerID int    `json:"accepted_answer_id"`
		AnswerCount      int    `json:"answer_count"`
		CreationDate     int    `json:"creation_date"`
		IsAnswered       bool   `json:"is_answered"`
		LastActivityDate int    `json:"last_activity_date"`
		LastEditDate     int    `json:"last_edit_date"`
		Link             string `json:"link"`
		Owner            struct {
			AcceptRate   int    `json:"accept_rate"`
			DisplayName  string `json:"display_name"`
			Link         string `json:"link"`
			ProfileImage string `json:"profile_image"`
			Reputation   int    `json:"reputation"`
			UserID       int    `json:"user_id"`
			UserType     string `json:"user_type"`
		} `json:"owner"`
		QuestionID int      `json:"question_id"`
		Score      int      `json:"score"`
		Tags       []string `json:"tags"`
		Title      string   `json:"title"`
		ViewCount  int      `json:"view_count"`
	} `json:"items"`
	QuotaMax       int `json:"quota_max"`
	QuotaRemaining int `json:"quota_remaining"`
}

type AnswerResult struct {
	Title string
	AID   int
	QID   int
}
