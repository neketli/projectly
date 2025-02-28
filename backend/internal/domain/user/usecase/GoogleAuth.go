package usecase

import (
	"context"
	"encoding/json"
	"fmt"
	"math/rand"
	"projectly-server/internal/domain/user/entity"
	"time"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var googleConfig = &oauth2.Config{}

func (uc *userUseCase) GoogleLogin(ctx context.Context, redirectURL string) string {
	googleConfig = &oauth2.Config{
		ClientID:     uc.config.GoogleAuthProvider.ClientID,
		ClientSecret: uc.config.GoogleAuthProvider.ClientSecret,
		RedirectURL:  redirectURL,
		Scopes: []string{
			"https://www.googleapis.com/auth/userinfo.email",
			"https://www.googleapis.com/auth/userinfo.profile",
		},
		Endpoint: google.Endpoint,
	}

	return googleConfig.AuthCodeURL("state", oauth2.AccessTypeOffline)
}

func (u *userUseCase) GoogleCallback(ctx context.Context, code string) (*entity.User, error) {
	token, err := googleConfig.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange token: %v", err)
	}

	client := googleConfig.Client(ctx, token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	defer resp.Body.Close()

	var userInfo struct {
		Email      string `json:"email"`
		ID         string `json:"id"`
		GivenName  string `json:"given_name"`
		FamilyName string `json:"family_name"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&userInfo); err != nil {
		return nil, fmt.Errorf("failed to decode user info: %v", err)
	}

	existingUser, err := u.GetUserByEmail(ctx, userInfo.Email)
	if err != nil && err != entity.ErrNoUserFound {
		return nil, fmt.Errorf("failed to get user info: %v", err)
	}
	if err == entity.ErrNoUserFound {
		user := &entity.User{
			Email:    userInfo.Email,
			Name:     userInfo.GivenName,
			Surname:  userInfo.FamilyName,
			Password: generateRandomPassword(),
			Meta: &entity.UserMeta{
				Provider:   "google",
				ProviderID: userInfo.ID,
			},
		}

		err = u.CreateUser(ctx, user)
		if err != nil {
			return nil, err
		}
		return user, nil
	}

	if existingUser.Meta.Provider == "google" {
		return &existingUser, nil
	}

	existingUser.Meta = &entity.UserMeta{
		Provider:   "google",
		ProviderID: userInfo.ID,
	}
	err = u.UpdateUser(ctx, &existingUser)
	if err != nil {
		return nil, fmt.Errorf("failed to update user info: %v", err)
	}
	return &existingUser, nil

}

func generateRandomPassword() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	const length = 12
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))

	password := make([]byte, length)
	for i := range password {
		password[i] = charset[seededRand.Intn(len(charset))]
	}
	return string(password)
}
