package actions

import (
	"gh_actions/util"
	"net/http"
)

func GetRunsByWorkflowIdResponse() http.Response {

	var respStr = `
	{
		"total_count": 1,
		"workflow_runs": [
			{
				"id": 4604017430,
				"name": "Build_Test_Deploy",
				"node_id": "WFR_kwLOHWWfL88AAAABEmu7Fg",
				"head_branch": "main",
				"head_sha": "9c7563b4dc6ff35518383b2ac175094cd754a705",
				"path": ".github/workflows/Build_Test_Deploy.yml",
				"display_title": "Build_Test_Deploy",
				"run_number": 90,
				"event": "workflow_dispatch",
				"status": "completed",
				"conclusion": "success",
				"workflow_id": 29051699,
				"check_suite_id": 12008855344,
				"check_suite_node_id": "CS_kwDOHWWfL88AAAACy8iXMA",
				"url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430",
				"html_url": "https://github.com/middesb/middetest001/actions/runs/4604017430",
				"pull_requests": [],
				"created_at": "2023-04-04T04:35:12Z",
				"updated_at": "2023-04-04T04:35:47Z",
				"actor": {
					"login": "middesb",
					"id": 105188650,
					"node_id": "U_kgDOBkUNKg",
					"avatar_url": "https://avatars.githubusercontent.com/u/105188650?v=4",
					"gravatar_id": "",
					"url": "https://api.github.com/users/middesb",
					"html_url": "https://github.com/middesb",
					"followers_url": "https://api.github.com/users/middesb/followers",
					"following_url": "https://api.github.com/users/middesb/following{/other_user}",
					"gists_url": "https://api.github.com/users/middesb/gists{/gist_id}",
					"starred_url": "https://api.github.com/users/middesb/starred{/owner}{/repo}",
					"subscriptions_url": "https://api.github.com/users/middesb/subscriptions",
					"organizations_url": "https://api.github.com/users/middesb/orgs",
					"repos_url": "https://api.github.com/users/middesb/repos",
					"events_url": "https://api.github.com/users/middesb/events{/privacy}",
					"received_events_url": "https://api.github.com/users/middesb/received_events",
					"type": "User",
					"site_admin": false
				},
				"run_attempt": 1,
				"referenced_workflows": [],
				"run_started_at": "2023-04-04T04:35:12Z",
				"triggering_actor": {
					"login": "middesb",
					"id": 105188650,
					"node_id": "U_kgDOBkUNKg",
					"avatar_url": "https://avatars.githubusercontent.com/u/105188650?v=4",
					"gravatar_id": "",
					"url": "https://api.github.com/users/middesb",
					"html_url": "https://github.com/middesb",
					"followers_url": "https://api.github.com/users/middesb/followers",
					"following_url": "https://api.github.com/users/middesb/following{/other_user}",
					"gists_url": "https://api.github.com/users/middesb/gists{/gist_id}",
					"starred_url": "https://api.github.com/users/middesb/starred{/owner}{/repo}",
					"subscriptions_url": "https://api.github.com/users/middesb/subscriptions",
					"organizations_url": "https://api.github.com/users/middesb/orgs",
					"repos_url": "https://api.github.com/users/middesb/repos",
					"events_url": "https://api.github.com/users/middesb/events{/privacy}",
					"received_events_url": "https://api.github.com/users/middesb/received_events",
					"type": "User",
					"site_admin": false
				},
				"jobs_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430/jobs",
				"logs_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430/logs",
				"check_suite_url": "https://api.github.com/repos/middesb/middetest001/check-suites/12008855344",
				"artifacts_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430/artifacts",
				"cancel_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430/cancel",
				"rerun_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430/rerun",
				"previous_attempt_url": null,
				"workflow_url": "https://api.github.com/repos/middesb/middetest001/actions/workflows/29051699",
				"head_commit": {
					"id": "9c7563b4dc6ff35518383b2ac175094cd754a705",
					"tree_id": "87d82997a1fc5b3a143c63576fa85d227bef2432",
					"message": "Update README.md",
					"timestamp": "2023-01-19T01:30:10Z",
					"author": {
						"name": "tamilds08",
						"email": "tamilmaran.s@releaseiq.io"
					},
					"committer": {
						"name": "GitHub",
						"email": "noreply@github.com"
					}
				},
				"repository": {
					"id": 493199151,
					"node_id": "R_kgDOHWWfLw",
					"name": "middetest001",
					"full_name": "middesb/middetest001",
					"private": false,
					"owner": {
						"login": "middesb",
						"id": 105188650,
						"node_id": "U_kgDOBkUNKg",
						"avatar_url": "https://avatars.githubusercontent.com/u/105188650?v=4",
						"gravatar_id": "",
						"url": "https://api.github.com/users/middesb",
						"html_url": "https://github.com/middesb",
						"followers_url": "https://api.github.com/users/middesb/followers",
						"following_url": "https://api.github.com/users/middesb/following{/other_user}",
						"gists_url": "https://api.github.com/users/middesb/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/middesb/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/middesb/subscriptions",
						"organizations_url": "https://api.github.com/users/middesb/orgs",
						"repos_url": "https://api.github.com/users/middesb/repos",
						"events_url": "https://api.github.com/users/middesb/events{/privacy}",
						"received_events_url": "https://api.github.com/users/middesb/received_events",
						"type": "User",
						"site_admin": false
					},
					"html_url": "https://github.com/middesb/middetest001",
					"description": null,
					"fork": false,
					"url": "https://api.github.com/repos/middesb/middetest001",
					"forks_url": "https://api.github.com/repos/middesb/middetest001/forks",
					"keys_url": "https://api.github.com/repos/middesb/middetest001/keys{/key_id}",
					"collaborators_url": "https://api.github.com/repos/middesb/middetest001/collaborators{/collaborator}",
					"teams_url": "https://api.github.com/repos/middesb/middetest001/teams",
					"hooks_url": "https://api.github.com/repos/middesb/middetest001/hooks",
					"issue_events_url": "https://api.github.com/repos/middesb/middetest001/issues/events{/number}",
					"events_url": "https://api.github.com/repos/middesb/middetest001/events",
					"assignees_url": "https://api.github.com/repos/middesb/middetest001/assignees{/user}",
					"branches_url": "https://api.github.com/repos/middesb/middetest001/branches{/branch}",
					"tags_url": "https://api.github.com/repos/middesb/middetest001/tags",
					"blobs_url": "https://api.github.com/repos/middesb/middetest001/git/blobs{/sha}",
					"git_tags_url": "https://api.github.com/repos/middesb/middetest001/git/tags{/sha}",
					"git_refs_url": "https://api.github.com/repos/middesb/middetest001/git/refs{/sha}",
					"trees_url": "https://api.github.com/repos/middesb/middetest001/git/trees{/sha}",
					"statuses_url": "https://api.github.com/repos/middesb/middetest001/statuses/{sha}",
					"languages_url": "https://api.github.com/repos/middesb/middetest001/languages",
					"stargazers_url": "https://api.github.com/repos/middesb/middetest001/stargazers",
					"contributors_url": "https://api.github.com/repos/middesb/middetest001/contributors",
					"subscribers_url": "https://api.github.com/repos/middesb/middetest001/subscribers",
					"subscription_url": "https://api.github.com/repos/middesb/middetest001/subscription",
					"commits_url": "https://api.github.com/repos/middesb/middetest001/commits{/sha}",
					"git_commits_url": "https://api.github.com/repos/middesb/middetest001/git/commits{/sha}",
					"comments_url": "https://api.github.com/repos/middesb/middetest001/comments{/number}",
					"issue_comment_url": "https://api.github.com/repos/middesb/middetest001/issues/comments{/number}",
					"contents_url": "https://api.github.com/repos/middesb/middetest001/contents/{+path}",
					"compare_url": "https://api.github.com/repos/middesb/middetest001/compare/{base}...{head}",
					"merges_url": "https://api.github.com/repos/middesb/middetest001/merges",
					"archive_url": "https://api.github.com/repos/middesb/middetest001/{archive_format}{/ref}",
					"downloads_url": "https://api.github.com/repos/middesb/middetest001/downloads",
					"issues_url": "https://api.github.com/repos/middesb/middetest001/issues{/number}",
					"pulls_url": "https://api.github.com/repos/middesb/middetest001/pulls{/number}",
					"milestones_url": "https://api.github.com/repos/middesb/middetest001/milestones{/number}",
					"notifications_url": "https://api.github.com/repos/middesb/middetest001/notifications{?since,all,participating}",
					"labels_url": "https://api.github.com/repos/middesb/middetest001/labels{/name}",
					"releases_url": "https://api.github.com/repos/middesb/middetest001/releases{/id}",
					"deployments_url": "https://api.github.com/repos/middesb/middetest001/deployments"
				},
				"head_repository": {
					"id": 493199151,
					"node_id": "R_kgDOHWWfLw",
					"name": "middetest001",
					"full_name": "middesb/middetest001",
					"private": false,
					"owner": {
						"login": "middesb",
						"id": 105188650,
						"node_id": "U_kgDOBkUNKg",
						"avatar_url": "https://avatars.githubusercontent.com/u/105188650?v=4",
						"gravatar_id": "",
						"url": "https://api.github.com/users/middesb",
						"html_url": "https://github.com/middesb",
						"followers_url": "https://api.github.com/users/middesb/followers",
						"following_url": "https://api.github.com/users/middesb/following{/other_user}",
						"gists_url": "https://api.github.com/users/middesb/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/middesb/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/middesb/subscriptions",
						"organizations_url": "https://api.github.com/users/middesb/orgs",
						"repos_url": "https://api.github.com/users/middesb/repos",
						"events_url": "https://api.github.com/users/middesb/events{/privacy}",
						"received_events_url": "https://api.github.com/users/middesb/received_events",
						"type": "User",
						"site_admin": false
					},
					"html_url": "https://github.com/middesb/middetest001",
					"description": null,
					"fork": false,
					"url": "https://api.github.com/repos/middesb/middetest001",
					"forks_url": "https://api.github.com/repos/middesb/middetest001/forks",
					"keys_url": "https://api.github.com/repos/middesb/middetest001/keys{/key_id}",
					"collaborators_url": "https://api.github.com/repos/middesb/middetest001/collaborators{/collaborator}",
					"teams_url": "https://api.github.com/repos/middesb/middetest001/teams",
					"hooks_url": "https://api.github.com/repos/middesb/middetest001/hooks",
					"issue_events_url": "https://api.github.com/repos/middesb/middetest001/issues/events{/number}",
					"events_url": "https://api.github.com/repos/middesb/middetest001/events",
					"assignees_url": "https://api.github.com/repos/middesb/middetest001/assignees{/user}",
					"branches_url": "https://api.github.com/repos/middesb/middetest001/branches{/branch}",
					"tags_url": "https://api.github.com/repos/middesb/middetest001/tags",
					"blobs_url": "https://api.github.com/repos/middesb/middetest001/git/blobs{/sha}",
					"git_tags_url": "https://api.github.com/repos/middesb/middetest001/git/tags{/sha}",
					"git_refs_url": "https://api.github.com/repos/middesb/middetest001/git/refs{/sha}",
					"trees_url": "https://api.github.com/repos/middesb/middetest001/git/trees{/sha}",
					"statuses_url": "https://api.github.com/repos/middesb/middetest001/statuses/{sha}",
					"languages_url": "https://api.github.com/repos/middesb/middetest001/languages",
					"stargazers_url": "https://api.github.com/repos/middesb/middetest001/stargazers",
					"contributors_url": "https://api.github.com/repos/middesb/middetest001/contributors",
					"subscribers_url": "https://api.github.com/repos/middesb/middetest001/subscribers",
					"subscription_url": "https://api.github.com/repos/middesb/middetest001/subscription",
					"commits_url": "https://api.github.com/repos/middesb/middetest001/commits{/sha}",
					"git_commits_url": "https://api.github.com/repos/middesb/middetest001/git/commits{/sha}",
					"comments_url": "https://api.github.com/repos/middesb/middetest001/comments{/number}",
					"issue_comment_url": "https://api.github.com/repos/middesb/middetest001/issues/comments{/number}",
					"contents_url": "https://api.github.com/repos/middesb/middetest001/contents/{+path}",
					"compare_url": "https://api.github.com/repos/middesb/middetest001/compare/{base}...{head}",
					"merges_url": "https://api.github.com/repos/middesb/middetest001/merges",
					"archive_url": "https://api.github.com/repos/middesb/middetest001/{archive_format}{/ref}",
					"downloads_url": "https://api.github.com/repos/middesb/middetest001/downloads",
					"issues_url": "https://api.github.com/repos/middesb/middetest001/issues{/number}",
					"pulls_url": "https://api.github.com/repos/middesb/middetest001/pulls{/number}",
					"milestones_url": "https://api.github.com/repos/middesb/middetest001/milestones{/number}",
					"notifications_url": "https://api.github.com/repos/middesb/middetest001/notifications{?since,all,participating}",
					"labels_url": "https://api.github.com/repos/middesb/middetest001/labels{/name}",
					"releases_url": "https://api.github.com/repos/middesb/middetest001/releases{/id}",
					"deployments_url": "https://api.github.com/repos/middesb/middetest001/deployments"
				}
			}
		]
	}
	`

	resp := util.FetchResponseFromString(respStr)
	resp.StatusCode = 200

	return resp

}

func GetWorkflowRunByRunIdResponse(runAttempt string) http.Response {

	var respStr = `
	{

				"id": 4604017430,
				"name": "Build_Test_DeployTesting..",
				"node_id": "WFR_kwLOHWWfL88AAAABEmu7Fg",
				"head_branch": "main",
				"head_sha": "9c7563b4dc6ff35518383b2ac175094cd754a705",
				"path": ".github/workflows/Build_Test_Deploy.yml",
				"display_title": "Build_Test_Deploy",
				"run_number": 90,
				"event": "workflow_dispatch",
				"status": "completed",
				"conclusion": "success",
				"workflow_id": 29051699,
				"check_suite_id": 12008855344,
				"check_suite_node_id": "CS_kwDOHWWfL88AAAACy8iXMA",
				"url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430",
				"html_url": "https://github.com/middesb/middetest001/actions/runs/4604017430",
				"pull_requests": [],
				"created_at": "2023-04-04T04:35:12Z",
				"updated_at": "2023-04-04T04:35:47Z",
				"actor": {
					"login": "middesb",
					"id": 105188650,
					"node_id": "U_kgDOBkUNKg",
					"avatar_url": "https://avatars.githubusercontent.com/u/105188650?v=4",
					"gravatar_id": "",
					"url": "https://api.github.com/users/middesb",
					"html_url": "https://github.com/middesb",
					"followers_url": "https://api.github.com/users/middesb/followers",
					"following_url": "https://api.github.com/users/middesb/following{/other_user}",
					"gists_url": "https://api.github.com/users/middesb/gists{/gist_id}",
					"starred_url": "https://api.github.com/users/middesb/starred{/owner}{/repo}",
					"subscriptions_url": "https://api.github.com/users/middesb/subscriptions",
					"organizations_url": "https://api.github.com/users/middesb/orgs",
					"repos_url": "https://api.github.com/users/middesb/repos",
					"events_url": "https://api.github.com/users/middesb/events{/privacy}",
					"received_events_url": "https://api.github.com/users/middesb/received_events",
					"type": "User",
					"site_admin": false
				},
				"run_attempt": ` + runAttempt +
		`,
				"referenced_workflows": [],
				"run_started_at": "2023-04-04T04:35:12Z",
				"triggering_actor": {
					"login": "middesb",
					"id": 105188650,
					"node_id": "U_kgDOBkUNKg",
					"avatar_url": "https://avatars.githubusercontent.com/u/105188650?v=4",
					"gravatar_id": "",
					"url": "https://api.github.com/users/middesb",
					"html_url": "https://github.com/middesb",
					"followers_url": "https://api.github.com/users/middesb/followers",
					"following_url": "https://api.github.com/users/middesb/following{/other_user}",
					"gists_url": "https://api.github.com/users/middesb/gists{/gist_id}",
					"starred_url": "https://api.github.com/users/middesb/starred{/owner}{/repo}",
					"subscriptions_url": "https://api.github.com/users/middesb/subscriptions",
					"organizations_url": "https://api.github.com/users/middesb/orgs",
					"repos_url": "https://api.github.com/users/middesb/repos",
					"events_url": "https://api.github.com/users/middesb/events{/privacy}",
					"received_events_url": "https://api.github.com/users/middesb/received_events",
					"type": "User",
					"site_admin": false
				},
				"jobs_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430/jobs",
				"logs_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430/logs",
				"check_suite_url": "https://api.github.com/repos/middesb/middetest001/check-suites/12008855344",
				"artifacts_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430/artifacts",
				"cancel_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430/cancel",
				"rerun_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4604017430/rerun",
				"previous_attempt_url": null,
				"workflow_url": "https://api.github.com/repos/middesb/middetest001/actions/workflows/29051699",
				"head_commit": {
					"id": "9c7563b4dc6ff35518383b2ac175094cd754a705",
					"tree_id": "87d82997a1fc5b3a143c63576fa85d227bef2432",
					"message": "Update README.md",
					"timestamp": "2023-01-19T01:30:10Z",
					"author": {
						"name": "tamilds08",
						"email": "tamilmaran.s@releaseiq.io"
					},
					"committer": {
						"name": "GitHub",
						"email": "noreply@github.com"
					}
				},
				"repository": {
					"id": 493199151,
					"node_id": "R_kgDOHWWfLw",
					"name": "middetest001",
					"full_name": "middesb/middetest001",
					"private": false,
					"owner": {
						"login": "middesb",
						"id": 105188650,
						"node_id": "U_kgDOBkUNKg",
						"avatar_url": "https://avatars.githubusercontent.com/u/105188650?v=4",
						"gravatar_id": "",
						"url": "https://api.github.com/users/middesb",
						"html_url": "https://github.com/middesb",
						"followers_url": "https://api.github.com/users/middesb/followers",
						"following_url": "https://api.github.com/users/middesb/following{/other_user}",
						"gists_url": "https://api.github.com/users/middesb/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/middesb/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/middesb/subscriptions",
						"organizations_url": "https://api.github.com/users/middesb/orgs",
						"repos_url": "https://api.github.com/users/middesb/repos",
						"events_url": "https://api.github.com/users/middesb/events{/privacy}",
						"received_events_url": "https://api.github.com/users/middesb/received_events",
						"type": "User",
						"site_admin": false
					},
					"html_url": "https://github.com/middesb/middetest001",
					"description": null,
					"fork": false,
					"url": "https://api.github.com/repos/middesb/middetest001",
					"forks_url": "https://api.github.com/repos/middesb/middetest001/forks",
					"keys_url": "https://api.github.com/repos/middesb/middetest001/keys{/key_id}",
					"collaborators_url": "https://api.github.com/repos/middesb/middetest001/collaborators{/collaborator}",
					"teams_url": "https://api.github.com/repos/middesb/middetest001/teams",
					"hooks_url": "https://api.github.com/repos/middesb/middetest001/hooks",
					"issue_events_url": "https://api.github.com/repos/middesb/middetest001/issues/events{/number}",
					"events_url": "https://api.github.com/repos/middesb/middetest001/events",
					"assignees_url": "https://api.github.com/repos/middesb/middetest001/assignees{/user}",
					"branches_url": "https://api.github.com/repos/middesb/middetest001/branches{/branch}",
					"tags_url": "https://api.github.com/repos/middesb/middetest001/tags",
					"blobs_url": "https://api.github.com/repos/middesb/middetest001/git/blobs{/sha}",
					"git_tags_url": "https://api.github.com/repos/middesb/middetest001/git/tags{/sha}",
					"git_refs_url": "https://api.github.com/repos/middesb/middetest001/git/refs{/sha}",
					"trees_url": "https://api.github.com/repos/middesb/middetest001/git/trees{/sha}",
					"statuses_url": "https://api.github.com/repos/middesb/middetest001/statuses/{sha}",
					"languages_url": "https://api.github.com/repos/middesb/middetest001/languages",
					"stargazers_url": "https://api.github.com/repos/middesb/middetest001/stargazers",
					"contributors_url": "https://api.github.com/repos/middesb/middetest001/contributors",
					"subscribers_url": "https://api.github.com/repos/middesb/middetest001/subscribers",
					"subscription_url": "https://api.github.com/repos/middesb/middetest001/subscription",
					"commits_url": "https://api.github.com/repos/middesb/middetest001/commits{/sha}",
					"git_commits_url": "https://api.github.com/repos/middesb/middetest001/git/commits{/sha}",
					"comments_url": "https://api.github.com/repos/middesb/middetest001/comments{/number}",
					"issue_comment_url": "https://api.github.com/repos/middesb/middetest001/issues/comments{/number}",
					"contents_url": "https://api.github.com/repos/middesb/middetest001/contents/{+path}",
					"compare_url": "https://api.github.com/repos/middesb/middetest001/compare/{base}...{head}",
					"merges_url": "https://api.github.com/repos/middesb/middetest001/merges",
					"archive_url": "https://api.github.com/repos/middesb/middetest001/{archive_format}{/ref}",
					"downloads_url": "https://api.github.com/repos/middesb/middetest001/downloads",
					"issues_url": "https://api.github.com/repos/middesb/middetest001/issues{/number}",
					"pulls_url": "https://api.github.com/repos/middesb/middetest001/pulls{/number}",
					"milestones_url": "https://api.github.com/repos/middesb/middetest001/milestones{/number}",
					"notifications_url": "https://api.github.com/repos/middesb/middetest001/notifications{?since,all,participating}",
					"labels_url": "https://api.github.com/repos/middesb/middetest001/labels{/name}",
					"releases_url": "https://api.github.com/repos/middesb/middetest001/releases{/id}",
					"deployments_url": "https://api.github.com/repos/middesb/middetest001/deployments"
				},
				"head_repository": {
					"id": 493199151,
					"node_id": "R_kgDOHWWfLw",
					"name": "middetest001",
					"full_name": "middesb/middetest001",
					"private": false,
					"owner": {
						"login": "middesb",
						"id": 105188650,
						"node_id": "U_kgDOBkUNKg",
						"avatar_url": "https://avatars.githubusercontent.com/u/105188650?v=4",
						"gravatar_id": "",
						"url": "https://api.github.com/users/middesb",
						"html_url": "https://github.com/middesb",
						"followers_url": "https://api.github.com/users/middesb/followers",
						"following_url": "https://api.github.com/users/middesb/following{/other_user}",
						"gists_url": "https://api.github.com/users/middesb/gists{/gist_id}",
						"starred_url": "https://api.github.com/users/middesb/starred{/owner}{/repo}",
						"subscriptions_url": "https://api.github.com/users/middesb/subscriptions",
						"organizations_url": "https://api.github.com/users/middesb/orgs",
						"repos_url": "https://api.github.com/users/middesb/repos",
						"events_url": "https://api.github.com/users/middesb/events{/privacy}",
						"received_events_url": "https://api.github.com/users/middesb/received_events",
						"type": "User",
						"site_admin": false
					},
					"html_url": "https://github.com/middesb/middetest001",
					"description": null,
					"fork": false,
					"url": "https://api.github.com/repos/middesb/middetest001",
					"forks_url": "https://api.github.com/repos/middesb/middetest001/forks",
					"keys_url": "https://api.github.com/repos/middesb/middetest001/keys{/key_id}",
					"collaborators_url": "https://api.github.com/repos/middesb/middetest001/collaborators{/collaborator}",
					"teams_url": "https://api.github.com/repos/middesb/middetest001/teams",
					"hooks_url": "https://api.github.com/repos/middesb/middetest001/hooks",
					"issue_events_url": "https://api.github.com/repos/middesb/middetest001/issues/events{/number}",
					"events_url": "https://api.github.com/repos/middesb/middetest001/events",
					"assignees_url": "https://api.github.com/repos/middesb/middetest001/assignees{/user}",
					"branches_url": "https://api.github.com/repos/middesb/middetest001/branches{/branch}",
					"tags_url": "https://api.github.com/repos/middesb/middetest001/tags",
					"blobs_url": "https://api.github.com/repos/middesb/middetest001/git/blobs{/sha}",
					"git_tags_url": "https://api.github.com/repos/middesb/middetest001/git/tags{/sha}",
					"git_refs_url": "https://api.github.com/repos/middesb/middetest001/git/refs{/sha}",
					"trees_url": "https://api.github.com/repos/middesb/middetest001/git/trees{/sha}",
					"statuses_url": "https://api.github.com/repos/middesb/middetest001/statuses/{sha}",
					"languages_url": "https://api.github.com/repos/middesb/middetest001/languages",
					"stargazers_url": "https://api.github.com/repos/middesb/middetest001/stargazers",
					"contributors_url": "https://api.github.com/repos/middesb/middetest001/contributors",
					"subscribers_url": "https://api.github.com/repos/middesb/middetest001/subscribers",
					"subscription_url": "https://api.github.com/repos/middesb/middetest001/subscription",
					"commits_url": "https://api.github.com/repos/middesb/middetest001/commits{/sha}",
					"git_commits_url": "https://api.github.com/repos/middesb/middetest001/git/commits{/sha}",
					"comments_url": "https://api.github.com/repos/middesb/middetest001/comments{/number}",
					"issue_comment_url": "https://api.github.com/repos/middesb/middetest001/issues/comments{/number}",
					"contents_url": "https://api.github.com/repos/middesb/middetest001/contents/{+path}",
					"compare_url": "https://api.github.com/repos/middesb/middetest001/compare/{base}...{head}",
					"merges_url": "https://api.github.com/repos/middesb/middetest001/merges",
					"archive_url": "https://api.github.com/repos/middesb/middetest001/{archive_format}{/ref}",
					"downloads_url": "https://api.github.com/repos/middesb/middetest001/downloads",
					"issues_url": "https://api.github.com/repos/middesb/middetest001/issues{/number}",
					"pulls_url": "https://api.github.com/repos/middesb/middetest001/pulls{/number}",
					"milestones_url": "https://api.github.com/repos/middesb/middetest001/milestones{/number}",
					"notifications_url": "https://api.github.com/repos/middesb/middetest001/notifications{?since,all,participating}",
					"labels_url": "https://api.github.com/repos/middesb/middetest001/labels{/name}",
					"releases_url": "https://api.github.com/repos/middesb/middetest001/releases{/id}",
					"deployments_url": "https://api.github.com/repos/middesb/middetest001/deployments"
				}

	}
	`

	resp := util.FetchResponseFromString(respStr)
	resp.StatusCode = 200

	return resp

}

func GetWorkflowRunsSinceResponse(sinceTimeStamp string) http.Response {
	ttref, _ := util.ParseGHADate("2023-04-04T04:35:12Z")
	tt, _ := util.ParseGHADate(sinceTimeStamp)
	if tt.Sub(ttref) > 0 {
		return GetRunsByWorkflowIdResponse()
	} else {
		var respStr = `
		{
			"total_count": 0,
			"workflow_runs": []
		}`
		resp := util.FetchResponseFromString(respStr)
		resp.StatusCode = 200

		return resp
	}

}

func GetJobsByRunIdResponse() http.Response {
	respStr := `{
		"total_count": 3,
		"jobs": [
			{
				"id": 12086187654,
				"run_id": 4449625918,
				"workflow_name": "Build_Test_Deploy",
				"head_branch": "main",
				"run_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4449625918",
				"run_attempt": 1,
				"node_id": "CR_kwDOHWWfL88AAAAC0GSWhg",
				"head_sha": "9c7563b4dc6ff35518383b2ac175094cd754a705",
				"url": "https://api.github.com/repos/middesb/middetest001/actions/jobs/12086187654",
				"html_url": "https://github.com/middesb/middetest001/actions/runs/4449625918/jobs/7814065732",
				"status": "completed",
				"conclusion": "success",
				"created_at": "2023-03-17T16:34:18Z",
				"started_at": "2023-03-17T16:34:28Z",
				"completed_at": "2023-03-17T16:34:33Z",
				"name": "build1",
				"steps": [
					{
						"name": "Set up job",
						"status": "completed",
						"conclusion": "success",
						"number": 1,
						"started_at": "2023-03-17T22:04:27.000+05:30",
						"completed_at": "2023-03-17T22:04:29.000+05:30"
					},
					{
						"name": "init",
						"status": "completed",
						"conclusion": "success",
						"number": 2,
						"started_at": "2023-03-17T22:04:29.000+05:30",
						"completed_at": "2023-03-17T22:04:29.000+05:30"
					},
					{
						"name": "Run actions/checkout@v3",
						"status": "completed",
						"conclusion": "success",
						"number": 3,
						"started_at": "2023-03-17T22:04:29.000+05:30",
						"completed_at": "2023-03-17T22:04:30.000+05:30"
					},
					{
						"name": "execute",
						"status": "completed",
						"conclusion": "success",
						"number": 4,
						"started_at": "2023-03-17T22:04:31.000+05:30",
						"completed_at": "2023-03-17T22:04:31.000+05:30"
					},
					{
						"name": "Post Run actions/checkout@v3",
						"status": "completed",
						"conclusion": "success",
						"number": 8,
						"started_at": "2023-03-17T22:04:31.000+05:30",
						"completed_at": "2023-03-17T22:04:31.000+05:30"
					},
					{
						"name": "Complete job",
						"status": "completed",
						"conclusion": "success",
						"number": 9,
						"started_at": "2023-03-17T22:04:31.000+05:30",
						"completed_at": "2023-03-17T22:04:31.000+05:30"
					}
				],
				"check_run_url": "https://api.github.com/repos/middesb/middetest001/check-runs/12086187654",
				"labels": [
					"ubuntu-latest"
				],
				"runner_id": 2,
				"runner_name": "GitHub Actions 2",
				"runner_group_id": 2,
				"runner_group_name": "GitHub Actions"
			},
			{
				"id": 12086193973,
				"run_id": 4449625918,
				"workflow_name": "Build_Test_Deploy",
				"head_branch": "main",
				"run_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4449625918",
				"run_attempt": 1,
				"node_id": "CR_kwDOHWWfL88AAAAC0GSvNQ",
				"head_sha": "9c7563b4dc6ff35518383b2ac175094cd754a705",
				"url": "https://api.github.com/repos/middesb/middetest001/actions/jobs/12086193973",
				"html_url": "https://github.com/middesb/middetest001/actions/runs/4449625918/jobs/7814070502",
				"status": "completed",
				"conclusion": "success",
				"created_at": "2023-03-17T16:34:34Z",
				"started_at": "2023-03-17T16:34:41Z",
				"completed_at": "2023-03-17T16:34:45Z",
				"name": "test1",
				"steps": [
					{
						"name": "Set up job",
						"status": "completed",
						"conclusion": "success",
						"number": 1,
						"started_at": "2023-03-17T22:04:40.000+05:30",
						"completed_at": "2023-03-17T22:04:41.000+05:30"
					},
					{
						"name": "init",
						"status": "completed",
						"conclusion": "success",
						"number": 2,
						"started_at": "2023-03-17T22:04:41.000+05:30",
						"completed_at": "2023-03-17T22:04:41.000+05:30"
					},
					{
						"name": "Run actions/checkout@v3",
						"status": "completed",
						"conclusion": "success",
						"number": 3,
						"started_at": "2023-03-17T22:04:41.000+05:30",
						"completed_at": "2023-03-17T22:04:43.000+05:30"
					},
					{
						"name": "execute",
						"status": "completed",
						"conclusion": "success",
						"number": 4,
						"started_at": "2023-03-17T22:04:43.000+05:30",
						"completed_at": "2023-03-17T22:04:43.000+05:30"
					},
					{
						"name": "Run actions/checkout@v3",
						"status": "completed",
						"conclusion": "success",
						"number": 5,
						"started_at": "2023-03-17T22:04:43.000+05:30",
						"completed_at": "2023-03-17T22:04:43.000+05:30"
					},
					{
						"name": "Post Run actions/checkout@v3",
						"status": "completed",
						"conclusion": "success",
						"number": 9,
						"started_at": "2023-03-17T22:04:45.000+05:30",
						"completed_at": "2023-03-17T22:04:45.000+05:30"
					},
					{
						"name": "Post Run actions/checkout@v3",
						"status": "completed",
						"conclusion": "success",
						"number": 10,
						"started_at": "2023-03-17T22:04:45.000+05:30",
						"completed_at": "2023-03-17T22:04:45.000+05:30"
					},
					{
						"name": "Complete job",
						"status": "completed",
						"conclusion": "success",
						"number": 11,
						"started_at": "2023-03-17T22:04:44.000+05:30",
						"completed_at": "2023-03-17T22:04:44.000+05:30"
					}
				],
				"check_run_url": "https://api.github.com/repos/middesb/middetest001/check-runs/12086193973",
				"labels": [
					"ubuntu-latest"
				],
				"runner_id": 2,
				"runner_name": "GitHub Actions 2",
				"runner_group_id": 2,
				"runner_group_name": "GitHub Actions"
			},
			{
				"id": 12086198836,
				"run_id": 4449625918,
				"workflow_name": "Build_Test_Deploy",
				"head_branch": "main",
				"run_url": "https://api.github.com/repos/middesb/middetest001/actions/runs/4449625918",
				"run_attempt": 1,
				"node_id": "CR_kwDOHWWfL88AAAAC0GTCNA",
				"head_sha": "9c7563b4dc6ff35518383b2ac175094cd754a705",
				"url": "https://api.github.com/repos/middesb/middetest001/actions/jobs/12086198836",
				"html_url": "https://github.com/middesb/middetest001/actions/runs/4449625918/jobs/7814074143",
				"status": "completed",
				"conclusion": "success",
				"created_at": "2023-03-17T16:34:47Z",
				"started_at": "2023-03-17T16:34:55Z",
				"completed_at": "2023-03-17T16:34:59Z",
				"name": "deploy1",
				"steps": [
					{
						"name": "Set up job",
						"status": "completed",
						"conclusion": "success",
						"number": 1,
						"started_at": "2023-03-17T22:04:54.000+05:30",
						"completed_at": "2023-03-17T22:04:55.000+05:30"
					},
					{
						"name": "init",
						"status": "completed",
						"conclusion": "success",
						"number": 2,
						"started_at": "2023-03-17T22:04:55.000+05:30",
						"completed_at": "2023-03-17T22:04:55.000+05:30"
					},
					{
						"name": "Run actions/checkout@v3",
						"status": "completed",
						"conclusion": "success",
						"number": 3,
						"started_at": "2023-03-17T22:04:56.000+05:30",
						"completed_at": "2023-03-17T22:04:56.000+05:30"
					},
					{
						"name": "execute",
						"status": "completed",
						"conclusion": "success",
						"number": 4,
						"started_at": "2023-03-17T22:04:56.000+05:30",
						"completed_at": "2023-03-17T22:04:56.000+05:30"
					},
					{
						"name": "Run actions/checkout@v3",
						"status": "completed",
						"conclusion": "success",
						"number": 5,
						"started_at": "2023-03-17T22:04:56.000+05:30",
						"completed_at": "2023-03-17T22:04:57.000+05:30"
					},
					{
						"name": "Post Run actions/checkout@v3",
						"status": "completed",
						"conclusion": "success",
						"number": 9,
						"started_at": "2023-03-17T22:04:57.000+05:30",
						"completed_at": "2023-03-17T22:04:57.000+05:30"
					},
					{
						"name": "Post Run actions/checkout@v3",
						"status": "completed",
						"conclusion": "success",
						"number": 10,
						"started_at": "2023-03-17T22:04:57.000+05:30",
						"completed_at": "2023-03-17T22:04:57.000+05:30"
					},
					{
						"name": "Complete job",
						"status": "completed",
						"conclusion": "success",
						"number": 11,
						"started_at": "2023-03-17T22:04:57.000+05:30",
						"completed_at": "2023-03-17T22:04:58.000+05:30"
					}
				],
				"check_run_url": "https://api.github.com/repos/middesb/middetest001/check-runs/12086198836",
				"labels": [
					"ubuntu-latest"
				],
				"runner_id": 2,
				"runner_name": "GitHub Actions 2",
				"runner_group_id": 2,
				"runner_group_name": "GitHub Actions"
			}
		]
	}`
	resp := util.FetchResponseFromString(respStr)
	resp.StatusCode = 200

	return resp
}

func GetJobLogsByJobIdResponse() http.Response {
	respStr := `2023-03-15T06:12:14.1563645Z Requested labels: ubuntu-latest
	2023-03-15T06:12:14.1563694Z Job defined at: middesb/test002/.github/workflows/test.yml@refs/heads/main
	2023-03-15T06:12:14.1563724Z Waiting for a runner to pick up this job...
	2023-03-15T06:12:14.9943885Z Job is waiting for a hosted runner to come online.
	2023-03-15T06:12:18.3892429Z Job is about to start running on the hosted runner: GitHub Actions 3 (hosted)
	2023-03-15T06:12:20.7672853Z Current runner version: '2.303.0'
	2023-03-15T06:12:20.7698180Z ##[group]Operating System
	2023-03-15T06:12:20.7698704Z Ubuntu
	2023-03-15T06:12:20.7699056Z 22.04.2
	2023-03-15T06:12:20.7699270Z LTS
	2023-03-15T06:12:20.7699546Z ##[endgroup]
	2023-03-15T06:12:20.7699859Z ##[group]Runner Image
	2023-03-15T06:12:20.7700192Z Image: ubuntu-22.04
	2023-03-15T06:12:20.7700481Z Version: 20230305.1
	2023-03-15T06:12:20.7700960Z Included Software: https://github.com/actions/runner-images/blob/ubuntu22/20230305.1/images/linux/Ubuntu2204-Readme.md
	2023-03-15T06:12:20.7701573Z Image Release: https://github.com/actions/runner-images/releases/tag/ubuntu22%2F20230305.1
	2023-03-15T06:12:20.7701939Z ##[endgroup]
	2023-03-15T06:12:20.7702263Z ##[group]Runner Image Provisioner
	2023-03-15T06:12:20.7702614Z 2.0.119.1
	2023-03-15T06:12:20.7702846Z ##[endgroup]
	2023-03-15T06:12:20.7703768Z ##[group]GITHUB_TOKEN Permissions
	2023-03-15T06:12:20.7704333Z Actions: write
	2023-03-15T06:12:20.7704656Z Checks: write
	2023-03-15T06:12:20.7705122Z Contents: write
	2023-03-15T06:12:20.7705513Z Deployments: write
	2023-03-15T06:12:20.7705779Z Discussions: write
	2023-03-15T06:12:20.7706077Z Issues: write
	2023-03-15T06:12:20.7706362Z Metadata: read
	2023-03-15T06:12:20.7706614Z Packages: write
	2023-03-15T06:12:20.7706932Z Pages: write
	2023-03-15T06:12:20.7707243Z PullRequests: write
	2023-03-15T06:12:20.7707533Z RepositoryProjects: write
	2023-03-15T06:12:20.7707876Z SecurityEvents: write
	2023-03-15T06:12:20.7708180Z Statuses: write
	2023-03-15T06:12:20.7708452Z ##[endgroup]
	2023-03-15T06:12:20.7712041Z Secret source: Actions
	2023-03-15T06:12:20.7712543Z Prepare workflow directory
	2023-03-15T06:12:20.8541108Z Prepare all required actions
	2023-03-15T06:12:20.8726353Z Getting action download info
	2023-03-15T06:12:21.1052839Z Download action repository 'actions/checkout@v3' (SHA:ac593985615ec2ede58e132d2e21d2b1cbd6127c)
	2023-03-15T06:12:21.9867704Z Complete job name: build
	2023-03-15T06:12:22.0903686Z ##[group]Run echo run identifier 
	2023-03-15T06:12:22.0904237Z [36;1mecho run identifier [0m
	2023-03-15T06:12:22.0961006Z shell: /usr/bin/bash -e {0}
	2023-03-15T06:12:22.0961369Z ##[endgroup]
	2023-03-15T06:12:22.1216013Z run identifier
	2023-03-15T06:12:22.1387896Z ##[group]Run sleep 30s
	2023-03-15T06:12:22.1388253Z [36;1msleep 30s[0m
	2023-03-15T06:12:22.1442314Z shell: /usr/bin/bash --noprofile --norc -e -o pipefail {0}
	2023-03-15T06:12:22.1442714Z ##[endgroup]
	2023-03-15T06:12:52.1828012Z ##[group]Run actions/checkout@v3
	2023-03-15T06:12:52.1828286Z with:
	2023-03-15T06:12:52.1828493Z   repository: middesb/test002
	2023-03-15T06:12:52.1828909Z   token: ***
	2023-03-15T06:12:52.1829102Z   ssh-strict: true
	2023-03-15T06:12:52.1829314Z   persist-credentials: true
	2023-03-15T06:12:52.1829524Z   clean: true
	2023-03-15T06:12:52.1829710Z   fetch-depth: 1
	2023-03-15T06:12:52.1829887Z   lfs: false
	2023-03-15T06:12:52.1830073Z   submodules: false
	2023-03-15T06:12:52.1830286Z   set-safe-directory: true
	2023-03-15T06:12:52.1830493Z ##[endgroup]
	2023-03-15T06:12:52.4151857Z Syncing repository: middesb/test002
	2023-03-15T06:12:52.4153521Z ##[group]Getting Git version info
	2023-03-15T06:12:52.4154077Z Working directory is '/home/runner/work/test002/test002'
	2023-03-15T06:12:52.4154557Z [command]/usr/bin/git version
	2023-03-15T06:12:52.4202145Z git version 2.39.2
	2023-03-15T06:12:52.4232658Z ##[endgroup]
	2023-03-15T06:12:52.4258567Z Temporarily overriding HOME='/home/runner/work/_temp/bb2d827d-caf4-48e2-b3ff-1b3cc283f39e' before making global git config changes
	2023-03-15T06:12:52.4259028Z Adding repository directory to the temporary git global config as a safe directory
	2023-03-15T06:12:52.4259516Z [command]/usr/bin/git config --global --add safe.directory /home/runner/work/test002/test002
	2023-03-15T06:12:52.4301344Z Deleting the contents of '/home/runner/work/test002/test002'
	2023-03-15T06:12:52.4306613Z ##[group]Initializing the repository
	2023-03-15T06:12:52.4310306Z [command]/usr/bin/git init /home/runner/work/test002/test002
	2023-03-15T06:12:52.4369019Z hint: Using 'master' as the name for the initial branch. This default branch name
	2023-03-15T06:12:52.4370266Z hint: is subject to change. To configure the initial branch name to use in all
	2023-03-15T06:12:52.4370706Z hint: of your new repositories, which will suppress this warning, call:
	2023-03-15T06:12:52.4370993Z hint: 
	2023-03-15T06:12:52.4371438Z hint: 	git config --global init.defaultBranch <name>
	2023-03-15T06:12:52.4371704Z hint: 
	2023-03-15T06:12:52.4372081Z hint: Names commonly chosen instead of 'master' are 'main', 'trunk' and
	2023-03-15T06:12:52.4372512Z hint: 'development'. The just-created branch can be renamed via this command:
	2023-03-15T06:12:52.4372777Z hint: 
	2023-03-15T06:12:52.4373012Z hint: 	git branch -m <name>
	2023-03-15T06:12:52.4381058Z Initialized empty Git repository in /home/runner/work/test002/test002/.git/
	2023-03-15T06:12:52.4395831Z [command]/usr/bin/git remote add origin https://github.com/middesb/test002
	2023-03-15T06:12:52.4432699Z ##[endgroup]
	2023-03-15T06:12:52.4433317Z ##[group]Disabling automatic garbage collection
	2023-03-15T06:12:52.4436059Z [command]/usr/bin/git config --local gc.auto 0
	2023-03-15T06:12:52.4466370Z ##[endgroup]
	2023-03-15T06:12:52.4466920Z ##[group]Setting up auth
	2023-03-15T06:12:52.4472906Z [command]/usr/bin/git config --local --name-only --get-regexp core\.sshCommand
	2023-03-15T06:12:52.4505369Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'core\.sshCommand' && git config --local --unset-all 'core.sshCommand' || :"
	2023-03-15T06:12:52.4809306Z [command]/usr/bin/git config --local --name-only --get-regexp http\.https\:\/\/github\.com\/\.extraheader
	2023-03-15T06:12:52.4839231Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'http\.https\:\/\/github\.com\/\.extraheader' && git config --local --unset-all 'http.https://github.com/.extraheader' || :"
	2023-03-15T06:12:52.5049549Z [command]/usr/bin/git config --local http.https://github.com/.extraheader AUTHORIZATION: basic ***
	2023-03-15T06:12:52.5088906Z ##[endgroup]
	2023-03-15T06:12:52.5089338Z ##[group]Fetching the repository
	2023-03-15T06:12:52.5095929Z [command]/usr/bin/git -c protocol.version=2 fetch --no-tags --prune --progress --no-recurse-submodules --depth=1 origin +2b3db99db9f2710ee49612a2fc296a301da3d31a:refs/remotes/origin/main
	2023-03-15T06:12:52.8291149Z remote: Enumerating objects: 6, done.        
	2023-03-15T06:12:52.8299614Z remote: Counting objects:  16% (1/6)        
	2023-03-15T06:12:52.8300428Z remote: Counting objects:  33% (2/6)        
	2023-03-15T06:12:52.8300850Z remote: Counting objects:  50% (3/6)        
	2023-03-15T06:12:52.8301436Z remote: Counting objects:  66% (4/6)        
	2023-03-15T06:12:52.8301835Z remote: Counting objects:  83% (5/6)        
	2023-03-15T06:12:52.8302413Z remote: Counting objects: 100% (6/6)        
	2023-03-15T06:12:52.8302804Z remote: Counting objects: 100% (6/6), done.        
	2023-03-15T06:12:52.8303443Z remote: Compressing objects:  33% (1/3)        
	2023-03-15T06:12:52.8303872Z remote: Compressing objects:  66% (2/3)        
	2023-03-15T06:12:52.8304463Z remote: Compressing objects: 100% (3/3)        
	2023-03-15T06:12:52.8304879Z remote: Compressing objects: 100% (3/3), done.        
	2023-03-15T06:12:52.8305790Z remote: Total 6 (delta 0), reused 1 (delta 0), pack-reused 0        
	2023-03-15T06:12:52.8365964Z From https://github.com/middesb/test002
	2023-03-15T06:12:52.8366715Z  * [new ref]         2b3db99db9f2710ee49612a2fc296a301da3d31a -> origin/main
	2023-03-15T06:12:52.8391549Z ##[endgroup]
	2023-03-15T06:12:52.8392085Z ##[group]Determining the checkout info
	2023-03-15T06:12:52.8394292Z ##[endgroup]
	2023-03-15T06:12:52.8394815Z ##[group]Checking out the ref
	2023-03-15T06:12:52.8399199Z [command]/usr/bin/git checkout --progress --force -B main refs/remotes/origin/main
	2023-03-15T06:12:52.8441523Z Switched to a new branch 'main'
	2023-03-15T06:12:52.8444828Z branch 'main' set up to track 'origin/main'.
	2023-03-15T06:12:52.8452986Z ##[endgroup]
	2023-03-15T06:12:52.8490875Z [command]/usr/bin/git log -1 --format='%H'
	2023-03-15T06:12:52.8541934Z '2b3db99db9f2710ee49612a2fc296a301da3d31a'
	2023-03-15T06:12:52.8660480Z ##[group]Run echo "Hello, world!"
	2023-03-15T06:12:52.8660741Z [36;1mecho "Hello, world!"[0m
	2023-03-15T06:12:52.8715076Z shell: /usr/bin/bash -e {0}
	2023-03-15T06:12:52.8715295Z ##[endgroup]
	2023-03-15T06:12:52.8789797Z Hello, world!
	2023-03-15T06:12:52.8810948Z ##[group]Run echo "Add other actions to build."
	2023-03-15T06:12:52.8811394Z [36;1mecho "Add other actions to build."[0m
	2023-03-15T06:12:52.8811700Z [36;1mecho "test, and deploy your project."[0m
	2023-03-15T06:12:52.8811953Z [36;1mecho run identifier - [0m
	2023-03-15T06:12:52.8860717Z shell: /usr/bin/bash -e {0}
	2023-03-15T06:12:52.8860933Z ##[endgroup]
	2023-03-15T06:12:52.8927413Z Add other actions to build.
	2023-03-15T06:12:52.8927792Z test, and deploy your project.
	2023-03-15T06:12:52.8928296Z run identifier -
	2023-03-15T06:12:52.8969889Z Post job cleanup.
	2023-03-15T06:12:53.0207414Z [command]/usr/bin/git version
	2023-03-15T06:12:53.0254044Z git version 2.39.2
	2023-03-15T06:12:53.0343523Z Temporarily overriding HOME='/home/runner/work/_temp/6eff454d-3994-41ae-817c-ad46ef30a307' before making global git config changes
	2023-03-15T06:12:53.0345634Z Adding repository directory to the temporary git global config as a safe directory
	2023-03-15T06:12:53.0351560Z [command]/usr/bin/git config --global --add safe.directory /home/runner/work/test002/test002
	2023-03-15T06:12:53.0402268Z [command]/usr/bin/git config --local --name-only --get-regexp core\.sshCommand
	2023-03-15T06:12:53.0442282Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'core\.sshCommand' && git config --local --unset-all 'core.sshCommand' || :"
	2023-03-15T06:12:53.0681424Z [command]/usr/bin/git config --local --name-only --get-regexp http\.https\:\/\/github\.com\/\.extraheader
	2023-03-15T06:12:53.0705065Z http.https://github.com/.extraheader
	2023-03-15T06:12:53.0723878Z [command]/usr/bin/git config --local --unset-all http.https://github.com/.extraheader
	2023-03-15T06:12:53.0763155Z [command]/usr/bin/git submodule foreach --recursive sh -c "git config --local --name-only --get-regexp 'http\.https\:\/\/github\.com\/\.extraheader' && git config --local --unset-all 'http.https://github.com/.extraheader' || :"
	2023-03-15T06:12:53.1197914Z Evaluate and set job outputs
	2023-03-15T06:12:53.1209424Z Set output 'output1'
	2023-03-15T06:12:53.1210311Z Set output 'output2'
	2023-03-15T06:12:53.1210909Z Cleaning up orphan processes
	`
	resp := util.FetchResponseFromString(respStr)
	resp.StatusCode = 200

	return resp
}

func RunWorkflowByIdResponse() http.Response {
	respStr := ""
	resp := util.FetchResponseFromString(respStr)
	util.Log(resp)
	resp.StatusCode = 200

	return resp
}
func RerunWorkflowByRunIdResponse() http.Response {
	respStr := ""
	resp := util.FetchResponseFromString(respStr)
	util.Log(resp)
	resp.StatusCode = 200

	return resp
}
