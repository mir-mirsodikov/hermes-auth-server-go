package gateway

import (
	"database/sql"
	"encoding/json"
	"log"

	"github.com/Hermes-chat-App/hermes-auth-server/internal/model"
)

func CreateUser(user *model.User) (*model.User, error) {
	userJson, err := json.Marshal(user)
	if err != nil {
		return nil, err
	}

	log.Print("userJson", string(userJson))

	userReturn, err := json.Marshal(&model.User{})
	if err != nil {
		return nil, err
	}

	db.Exec(`call create_user(@input, $2);`, sql.Named("input", `[{"username": "testing", "email": "testing", "name": "testing"}]`), userReturn)

	log.Print(string(userReturn))

	parsedUserReturn := &model.User{}

	if err := json.Unmarshal(userReturn, parsedUserReturn); err != nil {
		return nil, err
	}

	return parsedUserReturn, nil
}
