# Predeclared/ Basic types

- Kiểu Boolean: bool
- Kiểu chuỗi (String): string
- Các kiểu số (Numeric types):
  - Số nguyên không dấu: uint, uint8, uint32, uint64
  - Số nguyên có dấu: int, int8, int32, int64
  - Số thực dấu phẩy động: float, float32, float64
  - Số phức: complex64, complex128

# Composite types

- arrays
- pointers
- functions
- slices
- maps và channels
- struct
- interfaces

## Struct type

- Example:
  struct {
  Name string
  Capacity uint8
  Rooms uint8
  Smoking bool
  }

## Type definition

- Example:
  type Firstname string

  type Country struct {
  Name string
  CapitalCity string
  }

## Type struct variable creation

### With fields names

- Example:
  france := Country{
  Name: "France",
  CapitalCity: "Paris",
  }

empty := Country{}

### Without fields names

- Example:
  belgium := Country{
  "Belgium",
  "Bruxelles",
  }
- Giá trị các trường phải được chỉ định theo đúng thứ tự như khi khai báo cấu trúc
- Không được bỏ sót trường nào
- Không được trộn lẫn cả hai cú pháp

# Selector expression

- Example:
  usa := Country{
  Name: "United States of America",
  }
  usa.CapitalCity = "Washington DC"

# Embedded fields

- Trong một kiểu cấu trúc, chúng ta có thể thêm các trường nhúng. Các trường nhúng được định nghĩa một cách ẩn (implicitly).
- Example:
  type Hotel struct {
  Name string
  Capacity uint8
  Rooms uint8
  Smoking bool
  Country
  }

type Country struct {
Name string
CapitalCity string
}

- Trong kiểu cấu trúc Hotel, chúng ta có một trường nhúng là Country. Bản thân Country cũng là một kiểu cấu trúc khác.
- Các trường nhúng không có tên tường minh. Tên trường chính là tên của kiểu dữ liệu đó.

## Pointer type

- Example:
  type Hotel struct {
  Name string
  \*Country
  }

type Country struct {
Name string
CapitalCity string
}

- Ở đây chúng ta nhúng kiểu \*Country (con trỏ trỏ tới một phần tử kiểu Country). Tên trường chính là tên kiểu: Country.
- Example:
  hotel := Hotel{
  Name: "Hotel super luxe",
  Country: &Country{Name: "France"},
  }
  fmt.Println(hotel.Country.Name)
