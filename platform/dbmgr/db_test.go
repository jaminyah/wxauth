package dbmgr

import (
	"log"
	"testing"
	"wxauth/datatype"

	"github.com/stretchr/testify/require"
)

func TestInsertUser(t *testing.T) {

	arg := datatype.UserDataModel{
		Email:    "jan@mail.com",
		PassHash: "2468",
		Salt:     "69127",
		UserRole: "Client",
		Services: "Notify, WebRTC",
	}

	err := handle.InsertUser(arg)
	require.NoError(t, err)
}

func TestInsertUser2(t *testing.T) {

	arg := datatype.UserDataModel{
		Email:    "der@mail.com",
		PassHash: "72019",
		Salt:     "204159",
		UserRole: "Client",
		Services: "Notify",
	}

	err := handle.InsertUser(arg)
	require.NoError(t, err)
}

func TestGetUser(t *testing.T) {

	arg := datatype.UserDataModel{
		Email:    "der@mail.com",
		PassHash: "72019",
		Salt:     "76157",
		UserRole: "Client",
		Services: "Notify",
	}
	user, err := handle.GetUser(arg.Email)

	require.NoError(t, err)
	require.Equal(t, arg.Email, user.Email)
	require.Equal(t, arg.PassHash, user.PassHash)
	require.Equal(t, arg.UserRole, user.UserRole)
	require.Equal(t, arg.Services, user.Services)

}

func TestReadUsers(t *testing.T) {

	users, err := handle.ReadUsers()
	size := len(users)

	require.NoError(t, err)
	require.Equal(t, 2, size)

}

func TestUpdateMail(t *testing.T) {

	current := "der@mail.com"
	update := "derek@mail.com"

	user, err := handle.GetUser(current)
	if err != nil {
		log.Fatal("testUpdateMail - get user error")
	}

	err = handle.UpdateMail(user, update)
	require.NoError(t, err)
}

func TestUpdatePassRSA(t *testing.T) {

	current := "jank@mail.com"
	update := "jYi123"

	user, err := handle.GetUser(current)
	if err != nil {
		log.Fatal("testUpdateToken - get user error")
	}

	err = handle.UpdatePassRSA(user, update)
	require.NoError(t, err)
}

func TestUpdateServices(t *testing.T) {

	current := "jan@mail.com"
	update := "Nofity, WebRTC"

	user, err := handle.GetUser(current)
	if err != nil {
		log.Fatal("testUpdateServices - get user error")
	}

	err = handle.UpdateServices(user, update)
	require.NoError(t, err)
}

func TestDeleteUser(t *testing.T) {

	arg := datatype.UserDataModel{
		Email:    "der@mail.com",
		PassHash: "72019",
		Salt:     "76157",
		UserRole: "Client",
		Services: "Notify",
	}
	err := handle.DeleteUser(arg.Email)
	require.NoError(t, err)
}
