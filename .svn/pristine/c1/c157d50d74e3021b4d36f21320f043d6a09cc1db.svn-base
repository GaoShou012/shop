package GORMDemo

type UserController struct {}

func (u *UserController) IsExists(fields map[string]interface{}) (int, error) {

	query, err := makeQueryByFields(fields)
	if err != nil {
		return 0, err
	}

	return (&User{}).IsExist(query)
}

func (u *UserController) Create(fields map[string]interface{}) (*User, error) {

	user, err := convertFieldsToUser(fields)
	if err != nil {
		return nil, err
	}

	err = user.Create()
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserController) Update(user *User, fields map[string]interface{}) (*User, error) {

	err := user.Update(fields)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (u *UserController) RetrieveOne(fields map[string]interface{}) (*User, error) {

	conditions, err := makeQueryByFields(fields)
	if err != nil {
		return nil, err
	}

	users, err := (&User{}).Retrieve(conditions)
	if err != nil {
		return nil, err
	}

	if len(users) > 0 {
		return &users[0], nil
	}
	return nil, nil
}

func (u *UserController) RetrieveAll(fields map[string]interface{}) ([]User, error) {
	conditions, err := makeQueryByFields(fields)
	if err != nil {
		return nil, err
	}

	users, err := (&User{}).Retrieve(conditions)
	if err != nil {
		return nil, err
	}

	if len(users) > 0 {
		return users, nil
	}
	return nil, nil
}

func (u *UserController) DeleteAll(fields map[string]interface{}) error {
	conditions, err := makeQueryByFields(fields)
	if err != nil {
		return err
	}

	return (&User{}).Delete(conditions)
}
