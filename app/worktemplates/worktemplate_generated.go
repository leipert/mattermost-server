// Copyright (c) 2015-present Mattermost, Inc. All Rights Reserved.
// See LICENSE.txt for license information.

// Code generated by "make generate-worktemplates"
// DO NOT EDIT

package worktemplates

func init() {
	registerWorkTemplateCategory("product_teams", wtc846b565cd80043537945134a54812e07)
	registerWorkTemplateCategory("devops", wtca21c218df41f6d7fd032535fe20394e2)
	registerWorkTemplateCategory("companywide", wtca6def90c2edac0c33650ac8ebee1e094)
	registerWorkTemplateCategory("leadership", wtce9b74766edff1096ba7c67999ca259b6)
	registerWorkTemplate("product_teams/feature_release:v1", wt00a1b44a5831c0a3acb14787b3fdd352)
	registerWorkTemplate("product_teams/goals_and_okrs:v1", wt5baa68055bf9ea423273662e01ccc575)
	registerWorkTemplate("product_teams/bug_bash:v1", wtfeb56bc6a8f277c47b503bd1c92d830e)
	registerWorkTemplate("product_teams/sprint_planning:v1", wt8d2ef53deac5517eb349dc5de6150196)
	registerWorkTemplate("product_teams/product_roadmap:v1", wt00ab91a945627f4a624957dd80490bb2)
	registerWorkTemplate("devops/incident_resolution:v1", wtce19b9352a59d6a5d26f292d83e84377)
	registerWorkTemplate("devops/product_release:v1", wt37406285a41c18bcdeb881189f7acde0)
	registerWorkTemplate("companywide/goals_and_okrs:v1", wtf7b846d35810f8272eeb9a1a562025b5)
	registerWorkTemplate("companywide/create_project:v1", wtb9ab412890c2410c7b49eec8f12e7edc)
	registerWorkTemplate("leadership/goals_and_okrs:v1", wt32ab773bfe021e3d4913931041552559)

	// Register categories strings
	_ = T("worktemplate.category.product_teams")
	_ = T("worktemplate.category.devops")
	_ = T("worktemplate.category.companywide")
	_ = T("worktemplate.category.leadership")

	// Register translation strings
	_ = T("worktemplate.product_teams.feature_release.description.channel")
	_ = T("worktemplate.product_teams.feature_release.description.board")
	_ = T("worktemplate.product_teams.feature_release.description.playbook")
	_ = T("worktemplate.product_teams.feature_release.description.integration")
	_ = T("worktemplate.product_teams.goals_and_okrs.channel")
	_ = T("worktemplate.product_teams.goals_and_okrs.board")
	_ = T("worktemplate.product_teams.goals_and_okrs.integration")
	_ = T("worktemplate.product_teams.bug_bash.channel")
	_ = T("worktemplate.product_teams.bug_bash.board")
	_ = T("worktemplate.product_teams.bug_bash.playbook")
	_ = T("worktemplate.product_teams.bug_bash.integration")
	_ = T("worktemplate.product_teams.sprint_planning.channel")
	_ = T("worktemplate.product_teams.sprint_planning.board")
	_ = T("worktemplate.product_teams.sprint_planning.integration")
	_ = T("worktemplate.product_teams.product_roadmap.channel")
	_ = T("worktemplate.product_teams.product_roadmap.board")
	_ = T("worktemplate.devops.incident_resolution.description.channel")
	_ = T("worktemplate.devops.incident_resolution.description.board")
	_ = T("worktemplate.devops.incident_resolution.description.playbook")
	_ = T("worktemplate.devops.product_release.channel")
	_ = T("worktemplate.devops.product_release.board")
	_ = T("worktemplate.devops.product_release.playbook")
	_ = T("worktemplate.companywide.goals_and_okrs.channel")
	_ = T("worktemplate.companywide.goals_and_okrs.board")
	_ = T("worktemplate.companywide.goals_and_okrs.integration")
	_ = T("worktemplate.companywide.create_project.channel")
	_ = T("worktemplate.companywide.create_project.board")
	_ = T("worktemplate.companywide.create_project.integration")
	_ = T("worktemplate.leadership.goals_and_okrs.channel")
	_ = T("worktemplate.leadership.goals_and_okrs.board")
	_ = T("worktemplate.leadership.goals_and_okrs.integration")
}

var wtc846b565cd80043537945134a54812e07 = &WorkTemplateCategory{
	ID:   "product_teams",
	Name: "worktemplate.category.product_teams",
}

var wtca21c218df41f6d7fd032535fe20394e2 = &WorkTemplateCategory{
	ID:   "devops",
	Name: "worktemplate.category.devops",
}

var wtca6def90c2edac0c33650ac8ebee1e094 = &WorkTemplateCategory{
	ID:   "companywide",
	Name: "worktemplate.category.companywide",
}

var wtce9b74766edff1096ba7c67999ca259b6 = &WorkTemplateCategory{
	ID:   "leadership",
	Name: "worktemplate.category.leadership",
}

var wt00a1b44a5831c0a3acb14787b3fdd352 = &WorkTemplate{
	ID:           "product_teams/feature_release:v1",
	Category:     "product_teams",
	UseCase:      "Manage feature release",
	Illustration: "/static/worktemplates/product_teams/feature_release/feature_release.svg",
	Visibility:   "public",

	Description: Description{
		Channel: &TranslatableString{
			ID:             "worktemplate.product_teams.feature_release.description.channel",
			DefaultMessage: "Chat with your team in a Feature Release channel that connects easily with your boards, playbooks and app bots.",
			Illustration:   "",
		},
		Board: &TranslatableString{
			ID:             "worktemplate.product_teams.feature_release.description.board",
			DefaultMessage: "Use our Meeting Agenda board template for recurring meetings like standup and our Project Tasks board to manage the progress of tasks along the way.",
			Illustration:   "",
		},
		Playbook: &TranslatableString{
			ID:             "worktemplate.product_teams.feature_release.description.playbook",
			DefaultMessage: "Create transparent workflows across development teams to ensure your feature development process is seamless.",
			Illustration:   "",
		},
		Integration: &TranslatableString{
			ID:             "worktemplate.product_teams.feature_release.description.integration",
			DefaultMessage: "Increase productivity in your channel by integrating a Jira bot and Github bot. These will be downloaded for you.",
			Illustration:   "/static/worktemplates/product_teams/feature_release/integrations.png",
		},
	},
	Content: []Content{
		{
			Channel: &Channel{
				ID:           "feature-release",
				Name:         "Feature Release",
				Purpose:      "",
				Playbook:     "product-release-playbook",
				Illustration: "/static/worktemplates/product_teams/feature_release/channel.png",
			},
		},
		{
			Board: &Board{
				ID:           "board-meeting-agenda",
				Template:     "54fcf9c610f0ac5e4c522c0657c90602",
				Name:         "Meeting Agenda",
				Channel:      "feature-release",
				Illustration: "/static/worktemplates/boards/meeting_agenda.svg",
			},
		},
		{
			Board: &Board{
				ID:           "board-project-task",
				Template:     "a4ec399ab4f2088b1051c3cdf1dde4c3",
				Name:         "Project Task",
				Channel:      "feature-release",
				Illustration: "/static/worktemplates/boards/project_tasks.svg",
			},
		},
		{
			Playbook: &Playbook{
				Template:     "Product Release",
				Name:         "Feature release",
				ID:           "product-release-playbook",
				Illustration: "/static/worktemplates/playbooks/product_release.svg",
			},
		},
		{
			Integration: &Integration{
				ID: "jira",
			},
		},
		{
			Integration: &Integration{
				ID: "github",
			},
		},
	},
}

var wt5baa68055bf9ea423273662e01ccc575 = &WorkTemplate{
	ID:           "product_teams/goals_and_okrs:v1",
	Category:     "product_teams",
	UseCase:      "Set goals and OKR's",
	Illustration: "/static/worktemplates/product_teams/goals_and_okrs/goals_and_okrs.svg",
	Visibility:   "public",

	Description: Description{
		Channel: &TranslatableString{
			ID:             "worktemplate.product_teams.goals_and_okrs.channel",
			DefaultMessage: "Clear focus is essential to team success and with this Project you can document the team’s goals and OKR’s as well as post updates in the dedicated channel.",
			Illustration:   "",
		},
		Board: &TranslatableString{
			ID:             "worktemplate.product_teams.goals_and_okrs.board",
			DefaultMessage: "Clear focus is essential to team success and with this Project you can document the team’s goals and OKR’s as well as post updates in the dedicated channel.",
			Illustration:   "",
		},

		Integration: &TranslatableString{
			ID:             "worktemplate.product_teams.goals_and_okrs.integration",
			DefaultMessage: "Clear focus is essential to team success and with this Project you can document the team’s goals and OKR’s as well as post updates in the dedicated channel.",
			Illustration:   "/static/worktemplates/integrations.svg",
		},
	},
	Content: []Content{
		{
			Channel: &Channel{
				ID:           "channel-1674845108569",
				Name:         "Goals and OKR",
				Purpose:      "",
				Playbook:     "",
				Illustration: "/static/worktemplates/product_teams/goals_and_okrs/channel.png",
			},
		},
		{
			Board: &Board{
				ID:           "board-1674845139258",
				Template:     "7ba22ccfdfac391d63dea5c4b8cde0de",
				Name:         "Goals and OKR",
				Channel:      "channel-1674845108569",
				Illustration: "/static/worktemplates/boards/company_goal_and_okrs.svg",
			},
		},
		{
			Board: &Board{
				ID:           "board-1674845175528",
				Template:     "54fcf9c610f0ac5e4c522c0657c90602",
				Name:         "Meeting Agenda",
				Channel:      "channel-1674845108569",
				Illustration: "/static/worktemplates/boards/meeting_agenda.svg",
			},
		},
		{
			Integration: &Integration{
				ID: "zoom",
			},
		},
	},
}

var wtfeb56bc6a8f277c47b503bd1c92d830e = &WorkTemplate{
	ID:           "product_teams/bug_bash:v1",
	Category:     "product_teams",
	UseCase:      "Run a bug bash",
	Illustration: "/static/worktemplates/product_teams/bug_bash/bug_bash.svg",
	Visibility:   "public",

	Description: Description{
		Channel: &TranslatableString{
			ID:             "worktemplate.product_teams.bug_bash.channel",
			DefaultMessage: "Get organized and bash all the bugs with  this project! Build momentum and measure progress using included Playbook, Board, and Channel.",
			Illustration:   "",
		},
		Board: &TranslatableString{
			ID:             "worktemplate.product_teams.bug_bash.board",
			DefaultMessage: "Get organized and bash all the bugs with  this project! Build momentum and measure progress using included Playbook, Board, and Channel.",
			Illustration:   "",
		},
		Playbook: &TranslatableString{
			ID:             "worktemplate.product_teams.bug_bash.playbook",
			DefaultMessage: "Get organized and bash all the bugs with  this project! Build momentum and measure progress using included Playbook, Board, and Channel.",
			Illustration:   "",
		},
		Integration: &TranslatableString{
			ID:             "worktemplate.product_teams.bug_bash.integration",
			DefaultMessage: "Get organized and bash all the bugs with  this project! Build momentum and measure progress using included Playbook, Board, and Channel.",
			Illustration:   "/static/worktemplates/integrations.svg",
		},
	},
	Content: []Content{
		{
			Playbook: &Playbook{
				Template:     "Bug Bash",
				Name:         "Bug Bash",
				ID:           "playbook-1674844017943",
				Illustration: "/static/worktemplates/playbooks/bug_bash.svg",
			},
		},
		{
			Channel: &Channel{
				ID:           "channel-1674844017943",
				Name:         "Bug Bash",
				Purpose:      "",
				Playbook:     "playbook-1674844017943",
				Illustration: "/static/worktemplates/product_teams/bug_bash/channel.svg",
			},
		},
		{
			Integration: &Integration{
				ID: "jira",
			},
		},
	},
}

var wt8d2ef53deac5517eb349dc5de6150196 = &WorkTemplate{
	ID:           "product_teams/sprint_planning:v1",
	Category:     "product_teams",
	UseCase:      "Plan sprints",
	Illustration: "/static/worktemplates/product_teams/sprint_planning/sprint_planning.svg",
	Visibility:   "public",

	Description: Description{
		Channel: &TranslatableString{
			ID:             "worktemplate.product_teams.sprint_planning.channel",
			DefaultMessage: "Use a Project to make sprint planning a breeze. The channel keeps the conversation and questions focused. The sprint plan keeps everyone on task for the week and the Retrospective board brings the team together to continuously improve.",
			Illustration:   "",
		},
		Board: &TranslatableString{
			ID:             "worktemplate.product_teams.sprint_planning.board",
			DefaultMessage: "Use a Project to make sprint planning a breeze. The channel keeps the conversation and questions focused. The sprint plan keeps everyone on task for the week and the Retrospective board brings the team together to continuously improve.",
			Illustration:   "",
		},

		Integration: &TranslatableString{
			ID:             "worktemplate.product_teams.sprint_planning.integration",
			DefaultMessage: "Use a Project to make sprint planning a breeze. The channel keeps the conversation and questions focused. The sprint plan keeps everyone on task for the week and the Retrospective board brings the team together to continuously improve.",
			Illustration:   "/static/worktemplates/integrations.svg",
		},
	},
	Content: []Content{
		{
			Channel: &Channel{
				ID:           "channel-1674850783500",
				Name:         "Sprint planning",
				Purpose:      "",
				Playbook:     "",
				Illustration: "/static/worktemplates/product_teams/sprint_planning/channel.png",
			},
		},
		{
			Board: &Board{
				ID:           "board-1674850783973",
				Template:     "99b74e26d2f5d0a9b346d43c0a7bfb09",
				Name:         "Sprint planning",
				Channel:      "channel-1674850783500",
				Illustration: "/static/worktemplates/boards/sprint_planner.svg",
			},
		},
		{
			Integration: &Integration{
				ID: "zoom",
			},
		},
	},
}

var wt00ab91a945627f4a624957dd80490bb2 = &WorkTemplate{
	ID:           "product_teams/product_roadmap:v1",
	Category:     "product_teams",
	UseCase:      "Create a product roadmap",
	Illustration: "/static/worktemplates/product_teams/product_roadmap/product_roadmap.svg",
	Visibility:   "public",

	Description: Description{
		Channel: &TranslatableString{
			ID:             "worktemplate.product_teams.product_roadmap.channel",
			DefaultMessage: "Description of why the channel(s) are needed",
			Illustration:   "",
		},
		Board: &TranslatableString{
			ID:             "worktemplate.product_teams.product_roadmap.board",
			DefaultMessage: "Description of why the board(s) are needed",
			Illustration:   "",
		},
	},
	Content: []Content{
		{
			Channel: &Channel{
				ID:           "channel-1674851139450",
				Name:         "Product Roadmap",
				Purpose:      "",
				Playbook:     "",
				Illustration: "/static/worktemplates/product_teams/product_roadmap/channel.png",
			},
		},
		{
			Board: &Board{
				ID:           "board-1674851139759",
				Template:     "b728c6ca730e2cfc229741c5a4712b65",
				Name:         "Product Roadmap",
				Channel:      "channel-1674851139450",
				Illustration: "/static/worktemplates/boards/roadmap.svg",
			},
		},
	},
}

var wtce19b9352a59d6a5d26f292d83e84377 = &WorkTemplate{
	ID:           "devops/incident_resolution:v1",
	Category:     "devops",
	UseCase:      "Resolve incidents",
	Illustration: "/static/worktemplates/devops/incident_resolution/incident_resolution.png",
	Visibility:   "public",

	Description: Description{
		Channel: &TranslatableString{
			ID:             "worktemplate.devops.incident_resolution.description.channel",
			DefaultMessage: "When everything is going wrong, having a repeatable process is the key to making sure everything is made right as quickly as possible. This Project combines everything Mattermost offers to ensure the fires are put out and stakeholders informed along the way.",
			Illustration:   "",
		},
		Board: &TranslatableString{
			ID:             "worktemplate.devops.incident_resolution.description.board",
			DefaultMessage: "When everything is going wrong, having a repeatable process is the key to making sure everything is made right as quickly as possible. This Project combines everything Mattermost offers to ensure the fires are put out and stakeholders informed along the way.",
			Illustration:   "",
		},
		Playbook: &TranslatableString{
			ID:             "worktemplate.devops.incident_resolution.description.playbook",
			DefaultMessage: "When everything is going wrong, having a repeatable process is the key to making sure everything is made right as quickly as possible. This Project combines everything Mattermost offers to ensure the fires are put out and stakeholders informed along the way.",
			Illustration:   "",
		},
	},
	Content: []Content{
		{
			Playbook: &Playbook{
				Template:     "Incident Resolution",
				Name:         "Incident Resolution",
				ID:           "irpb",
				Illustration: "/static/worktemplates/playbooks/incident_resolution.svg",
			},
		},
		{
			Channel: &Channel{
				ID:           "irc",
				Name:         "Incident Resolution",
				Purpose:      "",
				Playbook:     "irpb",
				Illustration: "/static/worktemplates/devops/incident_resolution/channel.png",
			},
		},
		{
			Board: &Board{
				ID:           "irb",
				Template:     "a4ec399ab4f2088b1051c3cdf1dde4c3",
				Name:         "Incident Resolution",
				Channel:      "irc",
				Illustration: "/static/worktemplates/boards/project_tasks.svg",
			},
		},
	},
}

var wt37406285a41c18bcdeb881189f7acde0 = &WorkTemplate{
	ID:           "devops/product_release:v1",
	Category:     "devops",
	UseCase:      "Prepare a product release",
	Illustration: "/static/worktemplates/devops/product_release/product_release.svg",
	Visibility:   "public",

	Description: Description{
		Channel: &TranslatableString{
			ID:             "worktemplate.devops.product_release.channel",
			DefaultMessage: "Don’t miss a step during a product release with this Project. Assign tasks from the Playbook checklist and hit milestones with the Board. Use Channels to keep everyone on the same page.",
			Illustration:   "",
		},
		Board: &TranslatableString{
			ID:             "worktemplate.devops.product_release.board",
			DefaultMessage: "Don’t miss a step during a product release with this Project. Assign tasks from the Playbook checklist and hit milestones with the Board. Use Channels to keep everyone on the same page.",
			Illustration:   "",
		},
		Playbook: &TranslatableString{
			ID:             "worktemplate.devops.product_release.playbook",
			DefaultMessage: "Don’t miss a step during a product release with this Project. Assign tasks from the Playbook checklist and hit milestones with the Board. Use Channels to keep everyone on the same page.",
			Illustration:   "",
		},
	},
	Content: []Content{
		{
			Playbook: &Playbook{
				Template:     "Product Release",
				Name:         "Product Release",
				ID:           "playbook-1674851385983",
				Illustration: "/static/worktemplates/playbooks/product_release.svg",
			},
		},
		{
			Channel: &Channel{
				ID:           "channel-1674851385983",
				Name:         "Product Release",
				Purpose:      "",
				Playbook:     "playbook-1674851385983",
				Illustration: "/static/worktemplates/devops/product_release/channel.png",
			},
		},
		{
			Board: &Board{
				ID:           "board-1674851386432",
				Template:     "a4ec399ab4f2088b1051c3cdf1dde4c3",
				Name:         "Product Release",
				Channel:      "channel-1674851385983",
				Illustration: "/static/worktemplates/boards/project_tasks.svg",
			},
		},
	},
}

var wtf7b846d35810f8272eeb9a1a562025b5 = &WorkTemplate{
	ID:           "companywide/goals_and_okrs:v1",
	Category:     "companywide",
	UseCase:      "Set goals and OKR's",
	Illustration: "/static/worktemplates/companywide/goals_and_okrs/goals_and_okrs.svg",
	Visibility:   "public",

	Description: Description{
		Channel: &TranslatableString{
			ID:             "worktemplate.companywide.goals_and_okrs.channel",
			DefaultMessage: "Clear focus is essential to team success and with this Project you can document the team’s goals and OKR’s as well as post updates in the dedicated channel.",
			Illustration:   "",
		},
		Board: &TranslatableString{
			ID:             "worktemplate.companywide.goals_and_okrs.board",
			DefaultMessage: "Clear focus is essential to team success and with this Project you can document the team’s goals and OKR’s as well as post updates in the dedicated channel.",
			Illustration:   "",
		},

		Integration: &TranslatableString{
			ID:             "worktemplate.companywide.goals_and_okrs.integration",
			DefaultMessage: "Clear focus is essential to team success and with this Project you can document the team’s goals and OKR’s as well as post updates in the dedicated channel.",
			Illustration:   "/static/worktemplates/integrations.svg",
		},
	},
	Content: []Content{
		{
			Channel: &Channel{
				ID:           "channel-1674845108569",
				Name:         "Goals and OKR",
				Purpose:      "",
				Playbook:     "",
				Illustration: "/static/worktemplates/companywide/goals_and_okrs/channel.png",
			},
		},
		{
			Board: &Board{
				ID:           "board-1674845139258",
				Template:     "7ba22ccfdfac391d63dea5c4b8cde0de",
				Name:         "Goals and OKR",
				Channel:      "channel-1674845108569",
				Illustration: "/static/worktemplates/boards/company_goal_and_okrs.svg",
			},
		},
		{
			Integration: &Integration{
				ID: "zoom",
			},
		},
	},
}

var wtb9ab412890c2410c7b49eec8f12e7edc = &WorkTemplate{
	ID:           "companywide/create_project:v1",
	Category:     "companywide",
	UseCase:      "Create a project",
	Illustration: "/static/worktemplates/companywide/create_project/create_project.svg",
	Visibility:   "public",

	Description: Description{
		Channel: &TranslatableString{
			ID:             "worktemplate.companywide.create_project.channel",
			DefaultMessage: "Plan a Roadmap using this Project Board and collaborate on topic in the channel created with this template.",
			Illustration:   "",
		},
		Board: &TranslatableString{
			ID:             "worktemplate.companywide.create_project.board",
			DefaultMessage: "Plan a Roadmap using this Project Board and collaborate on topic in the channel created with this template.",
			Illustration:   "",
		},

		Integration: &TranslatableString{
			ID:             "worktemplate.companywide.create_project.integration",
			DefaultMessage: "Plan a Roadmap using this Project Board and collaborate on topic in the channel created with this template.",
			Illustration:   "/static/worktemplates/integrations.svg",
		},
	},
	Content: []Content{
		{
			Channel: &Channel{
				ID:           "channel-1674851940114",
				Name:         "Create Project",
				Purpose:      "",
				Playbook:     "",
				Illustration: "/static/worktemplates/companywide/create_project/channel.png",
			},
		},
		{
			Board: &Board{
				ID:           "board-1674851940548",
				Template:     "a4ec399ab4f2088b1051c3cdf1dde4c3",
				Name:         "Create Project",
				Channel:      "channel-1674851940114",
				Illustration: "/static/worktemplates/boards/project_tasks.svg",
			},
		},
		{
			Integration: &Integration{
				ID: "jira",
			},
		},
		{
			Integration: &Integration{
				ID: "github",
			},
		},
		{
			Integration: &Integration{
				ID: "zoom",
			},
		},
	},
}

var wt32ab773bfe021e3d4913931041552559 = &WorkTemplate{
	ID:           "leadership/goals_and_okrs:v1",
	Category:     "leadership",
	UseCase:      "Set goals and OKR's",
	Illustration: "/static/worktemplates/leadership/goals_and_okrs/goals_and_okrs.svg",
	Visibility:   "public",

	Description: Description{
		Channel: &TranslatableString{
			ID:             "worktemplate.leadership.goals_and_okrs.channel",
			DefaultMessage: "Clear focus is essential to team success and with this Project you can document the team’s goals and OKR’s as well as post updates in the dedicated channel.",
			Illustration:   "",
		},
		Board: &TranslatableString{
			ID:             "worktemplate.leadership.goals_and_okrs.board",
			DefaultMessage: "Clear focus is essential to team success and with this Project you can document the team’s goals and OKR’s as well as post updates in the dedicated channel.",
			Illustration:   "",
		},

		Integration: &TranslatableString{
			ID:             "worktemplate.leadership.goals_and_okrs.integration",
			DefaultMessage: "Clear focus is essential to team success and with this Project you can document the team’s goals and OKR’s as well as post updates in the dedicated channel.",
			Illustration:   "/static/worktemplates/integrations.svg",
		},
	},
	Content: []Content{
		{
			Channel: &Channel{
				ID:           "channel-1674845108569",
				Name:         "Goals and OKR",
				Purpose:      "",
				Playbook:     "",
				Illustration: "/static/worktemplates/leadership/goals_and_okrs/channel.png",
			},
		},
		{
			Board: &Board{
				ID:           "board-1674845139258",
				Template:     "7ba22ccfdfac391d63dea5c4b8cde0de",
				Name:         "Goals and OKR",
				Channel:      "channel-1674845108569",
				Illustration: "/static/worktemplates/boards/company_goal_and_okrs.svg",
			},
		},
		{
			Integration: &Integration{
				ID: "zoom",
			},
		},
	},
}
