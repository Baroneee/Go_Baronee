# Key

## Natural key

- dùng dữ liệu có ý nghĩa thực tế như email, CCCD

## Surrogate key

- Là khoá nhân tao không mang ý nghĩa nghiệp vụ (thường là id kiểu SERIAL, UUID hoặc auto-increment)
- Mục đích: Ổn định, ngắn, không phụ thuộc dữ liệu nghiệp vụ thay đổi
- Trường hợp đặc biệt:
  - Dùng UUID khi cần phân tán (distributed system) hoặc merge dữ liệu từ nhiều nguồn
  - Tránh dùng surrogate key nếu natural key ổn định và được truy vấn thường xuyên

#
