package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var (
	RootURL = "https://www.fifa.com"
)

var (
	ChromeDriverPath = "D:\\1_System\\1_Package\\chromedriver_win32\\chromedriver"
)

var (
	// URL之小组赛
	MatchesURLGroupPhase = "https://www.fifa.com/worldcup/matches/?#groupphase"
	// URL之淘汰赛
	MatchesURLKnockOutPhase = "https://www.fifa.com/worldcup/matches/?#knockoutphase"
	// URL之参赛球队
	TeamsURL = "https://www.fifa.com/worldcup/teams/"
	// URL之参赛球队分组情况
	GroupsURL = "https://www.fifa.com/worldcup/groups/"
	// URL之球员角色(守门员、后卫、中场、前锋...)
	PlayersURL = "https://www.fifa.com/worldcup/players/browser/"
	// URL之球员列表
	PlayersURLList = "https://www.fifa.com/worldcup/players/_libraries/byposition/all/_players-list"
	// URL之教练
	CoachesURL = "https://www.fifa.com/worldcup/players/coaches/"
	// URL之统计分析
	StatisticsURL = "https://www.fifa.com/worldcup/statistics/"
	// URL之统计分析球队的进球得分
	StatisticsTeamGoalURL = "https://www.fifa.com/worldcup/statistics/teams/goal-scored"
	// URL之统计分析球队的射手
	StatisticsTeamShots = "https://www.fifa.com/worldcup/statistics/teams/shots"
	// URL之统计分析球队的吃牌等纪律情况
	StatisticsTeamDisciplinary = "https://www.fifa.com/worldcup/statistics/teams/disciplinary"
	// URL之统计分析球员进球得分
	StatisticsPlayerGoalScoredURL = "https://www.fifa.com/worldcup/statistics/players/goal-scored"
	// URL之统计分析球员的救援数
	StatisticsPlayerSaves = "https://www.fifa.com/worldcup/statistics/players/saves"
	// URL之统计分析球员的射手排行榜
	StatisticsPlayerShots = "https://www.fifa.com/worldcup/statistics/players/shots"
	// URL之统计分析球员的吃牌等纪律情况
	StatisticsPlayerDisciplinary = "https://www.fifa.com/worldcup/statistics/players/disciplinary"
	// URL之奖项情况
	AwardsURL = "https://www.fifa.com/worldcup/awards/"
	// URL之历届世界杯经典瞬间
	HistoryURL = "https://www.fifa.com/worldcup/classic/"
)

func init() {
	viper.SetConfigName("settings")
	viper.AddConfigPath("./infra/config")
	viper.SetConfigType("yml")
	//viper.SetConfigFile()
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Errorf("Fatal error config file: %s \n", err)
		return
	}
}

// 获取Postgres数据库配置信息
func GetPostgreConfig() string {

	var (
		dbName   string
		host     string
		port     string
		user     string
		password string
		sslMode  string
	)

	dbName = viper.GetString("db.dbName")
	host = viper.GetString("db.host")
	port = viper.GetString("db.port")
	user = viper.GetString("db.user")
	password = viper.GetString("db.password")
	sslMode = viper.GetString("db.sslMode")

	return fmt.Sprintf("dbname=%s host=%s port=%s user=%s password=%s sslmode=%s", dbName, host, port, user, password, sslMode)
}
