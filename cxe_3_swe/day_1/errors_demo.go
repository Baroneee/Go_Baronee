package main

import (
	"errors"
	"fmt"
)

var ErrNotFound = errors.New("not found")

// ValidationError tuỳ chỉnh có khả năng bọc lỗi khác (ví dụ: ErrNotFound)
type ValidationError struct {
	Field string
	Msg   string
	Err   error // Lỗi gốc bị bọc bên trong
}

func (e *ValidationError) Error() string {
	if e.Err != nil {
		return fmt.Sprintf("%s: %s (%v)", e.Field, e.Msg, e.Err)
	}
	return fmt.Sprintf("%s: %s", e.Field, e.Msg)
}

// Unwrap cho phép errors.Is và errors.As "nhìn xuyên qua" ValidationError để tìm lỗi bên trong
func (e *ValidationError) Unwrap() error {
	return e.Err
}

func findUser(id int) error {
	if id <= 0 {
		return &ValidationError{
			Field: "id",
			Msg:   fmt.Sprintf("invalid id %d", id),
			Err:   ErrNotFound, // Bọc ErrNotFound vào trong ValidationError
		}
	}
	return nil
}

func main() {
	err := findUser(-1)

	// Kiểm tra errors.Is sẽ trả về true vì ValidationError có Unwrap() trả về ErrNotFound
	if errors.Is(err, ErrNotFound) {
		fmt.Println("Is NotFound: true")
	}

	// Kiểm tra errors.As sẽ trả về true và gán thành công vào biến ve
	var ve *ValidationError
	if errors.As(err, &ve) {
		fmt.Printf("As ValidationError -> Field: %s, Msg: %s\n", ve.Field, ve.Msg)
	}

	fmt.Println("Full error:", err)
}