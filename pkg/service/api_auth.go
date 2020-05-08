package service

import (
	"context"
	"log"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	"github.com/influenzanet/user-management-service/pkg/api"
	"github.com/influenzanet/user-management-service/pkg/models"
	"github.com/influenzanet/user-management-service/pkg/pwhash"
	"github.com/influenzanet/user-management-service/pkg/tokens"
	"github.com/influenzanet/user-management-service/pkg/utils"
)

func (s *userManagementServer) Status(ctx context.Context, _ *empty.Empty) (*api.ServiceStatus, error) {
	return &api.ServiceStatus{
		Status:  api.ServiceStatus_NORMAL,
		Msg:     "service running",
		Version: apiVersion,
	}, nil
}

func (s *userManagementServer) LoginWithEmail(ctx context.Context, req *api.LoginWithEmailMsg) (*api.TokenResponse, error) {
	if req == nil || req.Email == "" || req.Password == "" {
		return nil, status.Error(codes.InvalidArgument, "invalid username and/or password")
	}

	instanceID := req.InstanceId
	if instanceID == "" {
		instanceID = "default"
	}
	user, err := s.userDBservice.GetUserByEmail(instanceID, req.Email)

	if err != nil {
		return nil, status.Error(codes.InvalidArgument, "invalid username and/or password")
	}

	match, err := pwhash.ComparePasswordWithHash(user.Account.Password, req.Password)
	if err != nil || !match {
		return nil, status.Error(codes.InvalidArgument, "invalid username and/or password")
	}

	var username string
	if len(user.Roles) > 1 || len(user.Roles) == 1 && user.Roles[0] != "PARTICIPANT" {
		username = user.Account.AccountID
	}
	apiUser := user.ToAPI()

	// Access Token
	token, err := tokens.GenerateNewToken(
		apiUser.Id,
		apiUser.Profiles[0].Id,
		user.Roles,
		instanceID,
		s.JWT.TokenExpiryInterval,
		username,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Refresh Token
	rt, err := tokens.GenerateUniqueTokenString()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user.AddRefreshToken(rt)

	user, err = s.userDBservice.UpdateUser(req.InstanceId, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.userDBservice.UpdateLoginTime(instanceID, user.ID.Hex()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response := &api.TokenResponse{
		AccessToken:       token,
		RefreshToken:      rt,
		ExpiresIn:         int32(s.JWT.TokenExpiryInterval / time.Minute),
		Profiles:          apiUser.Profiles,
		SelectedProfileId: apiUser.Profiles[0].Id,
		PreferredLanguage: apiUser.Account.PreferredLanguage,
	}
	return response, nil

}

func (s *userManagementServer) SignupWithEmail(ctx context.Context, req *api.SignupWithEmailMsg) (*api.TokenResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}
	if !utils.CheckEmailFormat(req.Email) {
		return nil, status.Error(codes.InvalidArgument, "email not valid")
	}
	if !utils.CheckPasswordFormat(req.Password) {
		return nil, status.Error(codes.InvalidArgument, "password too weak")
	}

	password, err := pwhash.HashPassword(req.Password)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Create user DB object from request:
	newUser := models.User{
		Account: models.Account{
			Type:               "email",
			AccountID:          req.Email,
			AccountConfirmedAt: 0, // not confirmed yet
			Password:           password,
			PreferredLanguage:  req.PreferredLanguage,
		},
		Roles: []string{"PARTICIPANT"},
		Profiles: []models.Profile{
			{
				ID:                 primitive.NewObjectID(),
				Nickname:           req.Email,
				ConsentConfirmedAt: time.Now().Unix(),
				AvatarID:           "default",
			},
		},
	}
	newUser.AddNewEmail(req.Email, false)

	if req.WantsNewsletter {
		newUser.ContactPreferences.SubscribedToNewsletter = true
		newUser.ContactPreferences.SendNewsletterTo = []string{newUser.ContactInfos[0].ID.Hex()}
	}

	instanceID := req.InstanceId
	if instanceID == "" {
		instanceID = "default"
	}

	id, err := s.userDBservice.AddUser(instanceID, newUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	newUser.ID, _ = primitive.ObjectIDFromHex(id)

	log.Println("TODO: generate account confirmation token for newly created user")
	log.Println("TODO: send email for newly created user")
	// TODO: generate email confirmation token
	// TODO: send email with confirmation request

	var username string
	if len(newUser.Roles) > 1 || len(newUser.Roles) == 1 && newUser.Roles[0] != "PARTICIPANT" {
		username = newUser.Account.AccountID
	}
	apiUser := newUser.ToAPI()

	// Access Token
	token, err := tokens.GenerateNewToken(
		apiUser.Id,
		apiUser.Profiles[0].Id,
		newUser.Roles,
		instanceID,
		s.JWT.TokenExpiryInterval,
		username,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Refresh Token
	rt, err := tokens.GenerateUniqueTokenString()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	newUser.AddRefreshToken(rt)

	newUser, err = s.userDBservice.UpdateUser(req.InstanceId, newUser)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	if err := s.userDBservice.UpdateLoginTime(instanceID, newUser.ID.Hex()); err != nil {
		return nil, status.Error(codes.InvalidArgument, err.Error())
	}

	response := &api.TokenResponse{
		AccessToken:       token,
		RefreshToken:      rt,
		ExpiresIn:         int32(s.JWT.TokenExpiryInterval / time.Minute),
		Profiles:          apiUser.Profiles,
		SelectedProfileId: apiUser.Profiles[0].Id,
		PreferredLanguage: apiUser.Account.PreferredLanguage,
	}
	return response, nil
}

func (s *userManagementServer) SwitchProfile(ctx context.Context, req *api.SwitchProfileRequest) (*api.TokenResponse, error) {
	if req == nil || utils.IsTokenEmpty(req.Token) || req.ProfileId == "" || req.RefreshToken == "" {
		return nil, status.Error(codes.InvalidArgument, "missing argument")
	}
	user, err := s.userDBservice.GetUserByID(req.Token.InstanceId, req.Token.Id)
	if err != nil {
		return nil, status.Error(codes.Internal, "user not found")
	}

	profile, err := user.FindProfile(req.ProfileId)
	if err != nil {
		return nil, status.Error(codes.Internal, "profile not found")
	}

	if err := user.RemoveRefreshToken(req.RefreshToken); err != nil {
		return nil, status.Error(codes.Internal, "wrong refresh token")
	}

	var username string
	if len(user.Roles) > 1 || len(user.Roles) == 1 && user.Roles[0] != "PARTICIPANT" {
		username = user.Account.AccountID
	}
	apiUser := user.ToAPI()

	// Access Token
	token, err := tokens.GenerateNewToken(
		apiUser.Id,
		profile.ID.Hex(),
		user.Roles,
		req.Token.InstanceId,
		s.JWT.TokenExpiryInterval,
		username,
	)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	// Refresh Token
	rt, err := tokens.GenerateUniqueTokenString()
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}
	user.AddRefreshToken(rt)

	user, err = s.userDBservice.UpdateUser(req.Token.InstanceId, user)
	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	response := &api.TokenResponse{
		AccessToken:       token,
		RefreshToken:      rt,
		ExpiresIn:         int32(s.JWT.TokenExpiryInterval / time.Minute),
		Profiles:          apiUser.Profiles,
		SelectedProfileId: profile.ID.Hex(),
		PreferredLanguage: apiUser.Account.PreferredLanguage,
	}
	return response, nil
}