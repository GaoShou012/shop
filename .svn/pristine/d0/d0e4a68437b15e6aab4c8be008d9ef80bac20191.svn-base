package GORMDemo

import (
	"errors"
	"fmt"
)

func makeQuery(query string, key string, value interface{}, operation string) (string, error) {
	var err error

	if query != "" {
		query = query + " " + operation + " " + fmt.Sprintf("%s = '%v'", key, value)
	} else {
		query = fmt.Sprintf("%s = '%v'", key, value)
	}

	return query, err
}

func makeQueryByFields(fields map[string]interface{}) (string, error) {
	var err error
	var query string

	for k, v := range fields {
		query, err = makeQuery(query, k, v, "AND")
		if err != nil {
			panic(err)
		}
	}

	return query, nil
}

func convertFieldsToUser(fields map[string]interface{}) (*User, error) {
	user := new(User)

	for k, v := range fields {
		switch k {
		case "email":
			if email, ok := v.(string); !ok {
				return nil, errors.New("convertFieldsToUser fail, the type of email is illegal")
			} else {
				user.Email = email
			}
		}
	}

	return user, nil
}