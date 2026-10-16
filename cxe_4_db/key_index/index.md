# Index

- Là cấu trúc phụ giúp tìm dữ liệu nhanh mà không cần phải scan toàn bộ bảng
- Single-column index: Index trên 1 cột
  - Ví dụ: CREATE INDEX idx_users_email ON users(email);
    Phù hợp khi filter/join/sắp xếp theo đúng cột đó.
- Multi-column index: Index trên nhiều cột
  - Ví dụ: CREATE INDEX idx_orders_user_date ON orders(user_id, created_at);
  - Thứ tự cột rất quan trọng (leftmost prefix rule). Index này hỗ trợ tốt:
    - WHERE user_id = ?
    - WHERE user_id = ? AND created_at > ?
      Nhưng không hỗ trợ tốt WHERE created_at > ? (thiếu cột đầu).
