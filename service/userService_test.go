package service

import (
	"golang_bootstrap_project/mocks"
	"golang_bootstrap_project/model"
	"golang_bootstrap_project/repository"
	"reflect"
	"testing"
)

func getUserRepoMock() repository.UserRepo {
	userRepoMock := new(mocks.UserRepo)
	userRepoMock.On("FindActiveUserWithUserId", "LatVAlY").Return(&model.UserDTO{
		Id:        1,
		FirstName: "Abdel",
		LastName:  "Latrache",
		UserId:    "LatVAlY",
		Email:     "test@test.com",
	}, nil)
	return userRepoMock
}

func TestUserServiceImpl_FindActiveUserWithUserId(t *testing.T) {
	type fields struct {
		UserRepository repository.UserRepo
	}
	type args struct {
		userId string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *model.UserDTO
		wantErr bool
	}{
		{
			name: "ShouldReturnUserObjectWithGivenUserId",
			fields: fields{
				UserRepository: getUserRepoMock(),
			},
			args: args{
				userId: "LatVAlY",
			},
			want: &model.UserDTO{
				Id:        1,
				FirstName: "Abdel",
				LastName:  "Latrache",
				UserId:    "LatVAlY",
				Email:     "test@test.com",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := UserServiceImpl{
				UserRepository: tt.fields.UserRepository,
			}
			got, err := u.FindActiveUserWithUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("FindActiveUserWithUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindActiveUserWithUserId() got = %v, want %v", got, tt.want)
			}
		})
	}
}
