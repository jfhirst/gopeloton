package main

type MyProfileStruct struct {
	Username               string      `json:"username"`
	LastName               string      `json:"last_name"`
	IsDemo                 bool        `json:"is_demo"`
	Weight                 float64     `json:"weight"`
	IsProfilePrivate       bool        `json:"is_profile_private"`
	CyclingFtpWorkoutID    interface{} `json:"cycling_ftp_workout_id"`
	CreatedCountry         string      `json:"created_country"`
	Height                 float64     `json:"height"`
	IsProvisional          bool        `json:"is_provisional"`
	CyclingFtp             int         `json:"cycling_ftp"`
	ReferralCode           interface{} `json:"referral_code"`
	DefaultHeartRateZones  []float64   `json:"default_heart_rate_zones"`
	ID                     string      `json:"id"`
	TotalPendingFollowers  int         `json:"total_pending_followers"`
	BlockExplicit          bool        `json:"block_explicit"`
	FacebookAccessToken    interface{} `json:"facebook_access_token"`
	CustomizedMaxHeartRate int         `json:"customized_max_heart_rate"`
	IsStravaAuthenticated  bool        `json:"is_strava_authenticated"`
	HardwareSettings       interface{} `json:"hardware_settings"`
	IsCompleteProfile      bool        `json:"is_complete_profile"`
	InstructorID           interface{} `json:"instructor_id"`
	PairedDevices          []struct {
		SerialNumber     string `json:"serial_number"`
		Name             string `json:"name"`
		PairedDeviceType string `json:"paired_device_type"`
	} `json:"paired_devices"`
	V1ReferralsMade      int         `json:"v1_referrals_made"`
	LastWorkoutAt        int         `json:"last_workout_at"`
	Location             string      `json:"location"`
	IsInternalBetaTester bool        `json:"is_internal_beta_tester"`
	FacebookID           interface{} `json:"facebook_id"`
	WorkoutCounts        []struct {
		Count   int    `json:"count"`
		IconURL string `json:"icon_url"`
		Name    string `json:"name"`
		Slug    string `json:"slug"`
	} `json:"workout_counts"`
	HasActiveDigitalSubscription   bool   `json:"has_active_digital_subscription"`
	TotalNonPedalingMetricWorkouts int    `json:"total_non_pedaling_metric_workouts"`
	PhoneNumber                    string `json:"phone_number"`
	ContractAgreements             []struct {
		ContractID        string `json:"contract_id"`
		AgreedAt          int    `json:"agreed_at"`
		TreadContractURL  string `json:"tread_contract_url"`
		ContractCreatedAt int    `json:"contract_created_at"`
		ContractType      string `json:"contract_type"`
		BikeContractURL   string `json:"bike_contract_url"`
	} `json:"contract_agreements"`
	QuickHits struct {
		SpeedShortcuts   interface{} `json:"speed_shortcuts"`
		InclineShortcuts interface{} `json:"incline_shortcuts"`
		QuickHitsEnabled bool        `json:"quick_hits_enabled"`
	} `json:"quick_hits"`
	ExternalMusicAuthList []struct {
		Status   string      `json:"status"`
		Email    interface{} `json:"email"`
		Provider string      `json:"provider"`
	} `json:"external_music_auth_list"`
	FirstName                   string        `json:"first_name"`
	CardExpiresAt               int           `json:"card_expires_at"`
	CyclingFtpSource            interface{}   `json:"cycling_ftp_source"`
	MiddleInitial               string        `json:"middle_initial"`
	SubscriptionCreditsUsed     int           `json:"subscription_credits_used"`
	HasSignedWaiver             bool          `json:"has_signed_waiver"`
	CustomizedHeartRateZones    []interface{} `json:"customized_heart_rate_zones"`
	DefaultMaxHeartRate         int           `json:"default_max_heart_rate"`
	IsExternalBetaTester        bool          `json:"is_external_beta_tester"`
	ObfuscatedEmail             string        `json:"obfuscated_email"`
	TotalPedalingMetricWorkouts int           `json:"total_pedaling_metric_workouts"`
	CyclingWorkoutFtp           int           `json:"cycling_workout_ftp"`
	ReferralsMade               int           `json:"referrals_made"`
	Name                        string        `json:"name"`
	IsFitbitAuthenticated       bool          `json:"is_fitbit_authenticated"`
	HasActiveDeviceSubscription bool          `json:"has_active_device_subscription"`
	Gender                      string        `json:"gender"`
	CreatedAt                   int           `json:"created_at"`
	Email                       string        `json:"email"`
	MemberGroups                interface{}   `json:"member_groups"`
	Birthday                    int           `json:"birthday"`
	ImageURL                    string        `json:"image_url"`
	TotalFollowing              int           `json:"total_following"`
	EstimatedCyclingFtp         int           `json:"estimated_cycling_ftp"`
	CanCharge                   bool          `json:"can_charge"`
	SubscriptionCredits         int           `json:"subscription_credits"`
	TotalFollowers              int           `json:"total_followers"`
	TotalWorkouts               int           `json:"total_workouts"`
}

type WorkoutSummary struct {
	Count   int `json:"count"`
	Summary struct {
		Two01909 int `json:"2019-09"`
	} `json:"summary"`
	PageCount      int    `json:"page_count"`
	ShowNext       bool   `json:"show_next"`
	SortBy         string `json:"sort_by"`
	ShowPrevious   bool   `json:"show_previous"`
	Limit          int    `json:"limit"`
	AggregateStats []interface {
	} `json:"aggregate_stats"`
	Total int `json:"total"`
	Data  []struct {
		WorkoutType               string  `json:"workout_type"`
		TotalWork                 float64 `json:"total_work"`
		IsTotalWorkPersonalRecord bool    `json:"is_total_work_personal_record"`
		DeviceType                string  `json:"device_type"`
		Timezone                  string  `json:"timezone"`
		DeviceTimeCreatedAt       int     `json:"device_time_created_at"`
		ID                        string  `json:"id"`
		FitbitID                  string  `json:"fitbit_id"`
		PelotonID                 string  `json:"peloton_id"`
		UserID                    string  `json:"user_id"`
		Title                     string  `json:"title"`
		HasLeaderboardMetrics     bool    `json:"has_leaderboard_metrics"`
		HasPedalingMetrics        bool    `json:"has_pedaling_metrics"`
		Platform                  string  `json:"platform"`
		MetricsType               string  `json:"metrics_type"`
		FitnessDiscipline         string  `json:"fitness_discipline"`
		Status                    string  `json:"status"`
		StartTime                 int     `json:"start_time"`
		Name                      string  `json:"name"`
		StravaID                  string  `json:"strava_id"`
		Created                   int     `json:"created"`
		CreatedAt                 int     `json:"created_at"`
		EndTime                   int     `json:"end_time"`
	} `json:"data"`
	Page int `json:"page"`
}
type PerformanceStruct struct {
	IsClassPlanShown         bool          `json:"is_class_plan_shown"`
	SplitsData               []interface{} `json:"splits_data"`
	TargetPerformanceMetrics struct {
		CadenceTimeInRange int `json:"cadence_time_in_range"`
		TargetGraphMetrics []struct {
			GraphData struct {
				Upper   []int `json:"upper"`
				Lower   []int `json:"lower"`
				Average []int `json:"average"`
			} `json:"graph_data"`
			Average int    `json:"average"`
			Min     int    `json:"min"`
			Type    string `json:"type"`
			Max     int    `json:"max"`
		} `json:"target_graph_metrics"`
		ResistanceTimeInRange int `json:"resistance_time_in_range"`
	} `json:"target_performance_metrics"`
	LocationData     []interface{} `json:"location_data"`
	AverageSummaries []struct {
		DisplayName string  `json:"display_name"`
		Slug        string  `json:"slug"`
		Value       float64 `json:"value"`
		DisplayUnit string  `json:"display_unit"`
	} `json:"average_summaries"`
	Metrics []struct {
		DisplayName         string    `json:"display_name"`
		MaxValue            float64   `json:"max_value"`
		AverageValue        float64   `json:"average_value"`
		DisplayUnit         string    `json:"display_unit"`
		Values              []float64 `json:"values"`
		Slug                string    `json:"slug"`
		MissingDataDuration int       `json:"missing_data_duration,omitempty"`
		Zones               []struct {
			DisplayName string `json:"display_name"`
			MaxValue    int    `json:"max_value"`
			MinValue    int    `json:"min_value"`
			Range       string `json:"range"`
			Duration    int    `json:"duration"`
			Slug        string `json:"slug"`
		} `json:"zones,omitempty"`
	} `json:"metrics"`
	SegmentList []struct {
		IntensityInMets float64 `json:"intensity_in_mets"`
		Name            string  `json:"name"`
		StartTimeOffset int     `json:"start_time_offset"`
		IconURL         string  `json:"icon_url"`
		IconName        string  `json:"icon_name"`
		IconSlug        string  `json:"icon_slug"`
		Length          int     `json:"length"`
		MetricsType     string  `json:"metrics_type"`
		ID              string  `json:"id"`
	} `json:"segment_list"`
	Duration               int         `json:"duration"`
	IsLocationDataAccurate interface{} `json:"is_location_data_accurate"`
	HasAppleWatchMetrics   bool        `json:"has_apple_watch_metrics"`
	Summaries              []struct {
		DisplayName string  `json:"display_name"`
		Slug        string  `json:"slug"`
		Value       float64 `json:"value"`
		DisplayUnit string  `json:"display_unit"`
	} `json:"summaries"`
	SecondsSincePedalingStart []int `json:"seconds_since_pedaling_start"`
}
