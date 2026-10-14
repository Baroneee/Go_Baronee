# Method là gì

- Phương thức là hàm đi kèm với bộ thu (receiver)
- Receiver là 1 tham số đặc biệt
- Receiver không được liệt kê trong danh sách tham số thông thường mà được đặt trước tên phương thức
- Receiver có kiểu T hoặc \*T:
  - Khi có kiểu T thì gọi là value receiver
  - Khi có kiểu \*T thì gọi là pointer receiver
- T là base type

## Cấu trúc phương thức

// methods/first-example/main.go
package main

import (
"os/user"
"time"

    "github.com/Rhymond/go-money"

)

type Item struct {
ID string
}

type Cart struct {
ID string
CreatedAt time.Time
UpdatedAt time.Time
lockedAt time.Time
user.User
Items []Item
CurrencyCode string
isLocked bool
}

func (c *Cart) TotalPrice() (*money.Money, error) {
// ...
return nil, nil
}

func (c \*Cart) Lock() error {
// ...
return nil
}

- Chúng ta đã định nghĩa kiểu Cart. Kiểu này có 2 phương thức là TotalPrice và Lock. Hai phương thức này thực chất là các hàm:
  - Được gắn với kiểu Cart
  - Receiver của TotalPrice đặt tên là c và có kiểu \*Cart
  - Receiver của Lock đặt tên là c và có kiểu \*Cart

## Phương thức là các khả năng

- Thông qua các phương thức, chúng ta trao thêm các khả năng cho kiểu Kart, trong ví dụ trước ta thêm khả năng cho bất kì ai thao tác với Cart có thể Lock hoặc tính TotalPrice

## Tên method

- Tên method phải là duy nhất bên trong 1 method set
- Method set của một kiểu T là nhóm tất cả method có receiver là T
- Method set của một kiểu *T là nhóm tất cả method có receiver là T và *T
- Điều này có nghĩa là KHÔNG THỂ có 2 method trùng tên, ngay cả khi method đầu tiên dùng T còn method thứ hai dùng \*T.

## Cách gọi method

- Các method được gọi bằng dot nation

## Nên dùng T hay \*T

- Khi dùng T dữ liệu sẽ được sao chép nội bộ trước khi method được thực thi -> Thao tác trên 1 bản sao của biến. Vì vậy nên khi dùng T thì không thể sửa đổi dữ liệu gốc được truyền vào
- Với \*T dữ liệu được truyền vào có thể bị thay đổi trực tiếp bởi method
