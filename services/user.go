package services

import "strings"

var AuthorizedUserIds []string

func GetUserName(userId string) string {
	userIdParts := strings.Split(userId, "|")
	userName := userIdParts[0]

	return strings.Replace(userName, "_", " ", -1)
}

func CheckUserid(userId string) bool {
	for _, auhorizedUserid := range AuthorizedUserIds {
		if userId == auhorizedUserid {
			return true
		}
	}
	return false
}
