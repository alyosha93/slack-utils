package utils

import (
	"log"

	"github.com/nlopes/slack"
)

// GetAll returns all users for a workspace
func (u *User) GetAll() []slack.User {
	users, err := u.Client.GetUsers()
	if err != nil {
		log.Printf("Error getting user profiles: %v", err)
		return nil
	}
	return users
}

// EmailsToSlackIDs takes in an array of email addresses and finds the IDs of any workplace members with those emails
func (u *User) EmailsToSlackIDs(emails []string) []string {
	users := u.GetAll()
	ids := toSlackIDs(users, emails)

	return ids
}

func toSlackIDs(users []slack.User, emails []string) []string {
	var ids []string

	for _, email := range emails {
		for _, user := range users {
			if user.Profile.Email == email {
				ids = append(ids, user.ID)
			}
		}
	}

	return ids
}
