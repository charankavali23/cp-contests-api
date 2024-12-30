package models

type CodeChefContestDetails struct {
	ContestCode string `json:"contest_code"`
	ContestName string `json:"contest_name"`
	ContestStartDate string `json:"contest_start_date"`
	ContestEndDate string `json:"contest_end_date"`
	ContestStartDateIso string `json:"contest_start_date_iso"`
	ContestEndDateIso string `json:"contest_end_date_iso"`
	ContestDuration string `json:"contest_duration"`
	DistinctUsers int `json:"distinct_users"`
}

type CodeChefSkillTestDetails struct {
	CodeChefContestDetails
	ProblemsCount int `json:"problems_count"`
}

type CodeChefContests struct {
	Status string `json:"status"`
	Message string `json:"message"`
	PresentContests []CodeChefContestDetails `json:"present_contests"`
	FutureContests []CodeChefContestDetails `json:"future_contests"`
	PractiseContests []CodeChefContestDetails `json:"practise_contests"`
	PastContests []CodeChefContestDetails `json:"past_contests"`
	SkillTests []CodeChefSkillTestDetails `json:"skill_tests"`
	Banners [] struct {
		Image string `json:"image"`
		Link string `json:"link"`
	} `json:"banners"`
}
