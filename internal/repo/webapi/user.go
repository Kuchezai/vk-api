package webapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"vk-api/internal/entity"
)

type UserWebAPI struct {
	AccessToken string
}

func NewUserWebAPI(token string) *UserWebAPI {
	return &UserWebAPI{token}
}

// duplicates user from entity. but json tags are kept only here -> user from entity package don't know about implementation -> no abstraction leak
type userDTO struct {
	ID         int    `json:"id"`
	FirstName  string `json:"first_name"`
	SecondName string `json:"last_name"`
}

type userResponse struct {
	Response []userDTO `json:"response"`
	Error    `json:"error"`
}

type friendsResponse struct {
	Response struct {
		Items []userDTO `json:"items"`
	} `json:"response"`

	Error `json:"error"`
}

type Error struct {
	Code int    `json:"error_code"`
	Msg  string `json:"error_msg"`
}

func (r *UserWebAPI) User(userID string) (entity.User, error) {
	apiUrl := fmt.Sprintf("https://api.vk.com/method/users.get?user_ids=%s&access_token=%s&v=5.131", userID, r.AccessToken)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return entity.User{}, err
	}
	defer resp.Body.Close()

	ur := userResponse{}
	err = json.NewDecoder(resp.Body).Decode(&ur)
	if err != nil {
		return entity.User{}, err
	}

	if len(ur.Response) == 0 {
		return entity.User{}, fmt.Errorf("err: %s", "user not found")
	}

	if ur.Error.Code != 0 {
		return entity.User{}, fmt.Errorf("err: %s", ur.Error.Msg)
	}

	return entity.User(ur.Response[0]), err

}

func (r *UserWebAPI) FriendsByID(userID string) ([]entity.User, error) {
	apiUrl := fmt.Sprintf("https://api.vk.com/method/friends.get?user_id=%s&fields=counters&access_token=%s&v=5.131", userID, r.AccessToken)

	resp, err := http.Get(apiUrl)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	fr := friendsResponse{}
	err = json.NewDecoder(resp.Body).Decode(&fr)
	if err != nil {
		return nil, err
	}

	if fr.Error.Code != 0 {
		if fr.Error.Code == 100 {
			return nil, fmt.Errorf("err: %s", "invalid user_id")
		} else if fr.Error.Code == 30 {
			return nil, fmt.Errorf("err: %s", "user is private")
		} else {
			return nil, fmt.Errorf("err: %s", fr.Error.Msg)
		}
	}

	users := make([]entity.User, 0, len(fr.Response.Items))

	for _, us := range fr.Response.Items {
		users = append(users, entity.User(us))
	}

	return users, nil
}
