package types

type User struct {
	Name string `json:"name"`
	Age  int64  `json:"age"`
}

type CreateUserResponse struct {
	*ApiResponse
}

type CreateRequest struct {
	Name string `json:"name" binding:"required"`
	Age  int64  `json:"age" binding:"required"`
}

type GetUserResponse struct {
	*ApiResponse
	Users []*User `json:"result"`
}

type UpdateUserResponse struct {
	*ApiResponse
}

type UpdateRequest struct {
	Name string `json:"name" binding:"required"`

	UpdatedAge int64 `json:"updatedAge" binding:"required"`
}

type DeleteUserResponse struct {
	*ApiResponse
}

type DeleteRequest struct {
	Name string `json:"name" binding:"required"`
}
