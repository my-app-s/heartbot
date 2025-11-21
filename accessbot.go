package heartbot

import (
    "strconv"
	"os"
    "log"
	"github.com/joho/godotenv" // For work getenv
)

func init() {
	err:=godotenv.Load() // Initialization env file
	if err!=nil { log.Fatal("[FATAL Initialization package heartbot/accessbot]Not .env file.") }
}

// AllowedUser set status debug from env.
func AllowedUser() int64{
	// Realization
	allowedUserID:=os.Getenv("ALLOWED_USER_ID") // Get ALLOWED_USER_ID
	parseID,err:=strconv.Atoi(allowedUserID)
	if err!=nil { log.Panicf("[WARNING AllowedUser]Invalid ALLOWED_USER_ID value error: %s.",err) }
	return int64(parseID)
}

// SetDebugBot set status debug from env.
func AccessCode() int64{
	// Realization
	code:=os.Getenv("ACCESS_CODE") // Get ACCESS_CODE
	parseCode,err:=strconv.Atoi(code)
	if err!=nil { log.Panicf("[PANIC AccessCode]%s.",err) }
	return int64(parseCode)
}

// GetStatusCheck get status check access from env.
func GetStatusCheck() bool{
	// Realization
	statusCheckAccess:=os.Getenv("STATUS_CHECK_ACCESS")
	parse,err:=strconv.ParseBool(statusCheckAccess)
	if err!=nil { log.Panicf("[PANIC AccessCode]%s.",err) }
	return parse
}