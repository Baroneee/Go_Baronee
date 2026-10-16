# B-tree (Balanced Tree)

- Là cấu trúc dữ liệu mặc định của hầu hết index
  - Cây cân bằng -> mọi đường đi từ gốc đến là gần như bằng nhau -> Tìm kiếm O(log n)
  - Hỗ trợ tốt: equality, range, order by, group by
  - Không tối ưu cho full-text search
- Mục đích của indexing: Tăng tốc đọc, nhưng làm chậm mỗi lần ghi

# Trường hợp đặc biệt

- Over-indexing -> ghi chậm, tốn dung lượng
- Index không được dùng nếu selectivity thấp (cột có ít giá trị khác nhau, Ví dụ: giới tính)
- Covering index: index chưa đủ cột cần thiết -> không phải quay lại bảng (Index Only Scan)
