package users

import "testing"

func TestAddedUserCanBeRetrieved(t *testing.T) {
    email := "john.doe@gmail.com"
    AddUser(email, "some-token")
    user, err := GetUserByEmail(email)
    if err != nil || user.Email != email {
        t.Fail()
    }
}
