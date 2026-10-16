# Transaction là một đơn vị công việc logic: hoặc tất cả thành công, hoặc tất cả thất bại(all-or-nothing)

# ACID(4 thuộc tính bắt buộc của 1 transaction tốt):

- Atomicty: Tất cả hoặc không gì cả
- Consistency: Database luôn chuyển từ trạng thái hợp lệ sang trạng thái hợp lệ
- Isolation: Các transaction chạy đồng thời không được can thiệp lẫn nhau
- Durability: Khi commit rồi thì dữ liệu không mất dù hệ thống sập

## Atomicity được thực hiện như thế nào?

- Database thường sử dụng các cơ chế như:
  - Transaction log hoặc write-ahead log (WAL).
  - Thông tin undo hoặc cơ chế tương đương để hoàn tác thay đổi.
  - Recovery khi database khởi động lại sau sự cố.
- Cơ chế cụ thể khác nhau giữa các hệ quản trị. Ý tưởng chung là database có đủ thông tin để xác định thay đổi nào cần được giữ lại và thay đổi nào cần được hoàn tác.
- Lưu ý: Atomicity áp dụng cho các thao tác thuộc cùng transaction trong phạm vi database hỗ trợ. Nó không tự động hoàn tác email đã gửi, tiền đã chuyển qua một ngân hàng bên ngoài hay API bên thứ ba đã được gọi. Những trường hợp đó cần cơ chế phối hợp bổ sung.

## Các lỗi thường gặp khi transaction chạy đồng thời

- Dirty Read: đọc dữ liệu chưa commit
- Non-repeatable Read: đọc lại cùng một dòng như nhận giá trị khác
- Phantom Read: xuất hiện hoặc biến mất các dòng phù hợp điều kiện

## Các mức Isolation phổ biến

- Read uncommitted: Có thể cho phép dirty read
- Read committed: Chỉ đọc dữ liệu đã commit; có thể gặp non-repeatable read
- Repeatable read: Bảo vệ khả năng đọc lặp lại theo định nghĩa của mức này, xử lý phantom và ácc bất thường khác tuỳ hệ quản trị
- Serializable: Mức mạnh nhất, nhằm bảo đảm kết quả tương đương với một thứ tự thực thi tuần tự hợp lệ.
- Chú ý: PostgreSQL triển khai READ UNCOMMITTED như READ COMMITTED. REPEATABLE READ của PostgreSQL sử dụng snapshot isolation và ngăn phantom read theo cách thông thường, nhưng vẫn có thể xảy ra một số bất thường như write skew. SERIALIZABLE có thể yêu cầu ứng dụng thử lại transaction khi phát hiện xung đột.

- Isolation cao có thể giúp ngăn nhiều bất thường, nhưng cũng có thể dẫn đến:
  - Nhiều xung đột hơn.
  - Transaction phải chờ nhau.
  - Deadlock.
  - Transaction bị từ chối và cần retry.
  - Giảm thông lượng trong một số workload.
- Vì vậy, cần chọn mức isolation phù hợp với yêu cầu nghiệp vụ và hiểu cơ chế của database đang sử dụng.

## Database bảo đảm dữ liệu bằng cách nào?

- Một kỹ thuật phổ biến là Write-Ahead Logging (WAL) hoặc transaction log.
- Ý tưởng cơ bản:
  - Database ghi thông tin cần thiết để khôi phục thay đổi vào log.
  - Log được lưu bền vững theo cơ chế mà database yêu cầu trước khi xác nhận commit.
  - Nếu máy chủ bị sập, database đọc log khi khởi động lại để khôi phục những thay đổi đã commit và loại bỏ hoặc hoàn tác những thay đổi chưa commit khi cần thiết.
- Database không nhất thiết phải ghi ngay tất cả các trang dữ liệu đã thay đổi vào file dữ liệu chính trước khi trả về commit. Transaction log giúp việc phục hồi có thể diễn ra ngay cả khi các trang dữ liệu được ghi xuống sau đó.

## Durability có nghĩa là không bao giờ mất dữ liệu không?

- Không hoàn toàn. Durability là cam kết trong phạm vi cấu hình và những loại sự cố database hỗ trợ.
- Ví dụ, durability có thể bị ảnh hưởng nếu:
  - Ổ đĩa bị hỏng hoàn toàn và không có bản sao dự phòng.
  - Cấu hình ghi log không bảo đảm độ bền như mong muốn.
  - Replication bất đồng bộ chưa kịp sao chép dữ liệu trước khi máy chủ chính gặp sự cố.
  - Người quản trị vô tình xóa dữ liệu.
  - Một lỗi nghiêm trọng xảy ra ngoài phạm vi bảo vệ của hệ thống.
- Vì vậy, durability không thay thế cho backup, replication, kiểm tra khả năng phục hồi và kế hoạch khôi phục sau thảm họa. Ngoài ra, cũng cần phân biệt hai tình huống:
  - Database đã commit nhưng phản hồi bị mất: Ứng dụng có thể không biết giao dịch thành công hay chưa.
  - Database chưa commit: Transaction có thể chưa được lưu thành công.
- Đây là lý do các hệ thống thanh toán thường sử dụng mã giao dịch duy nhất hoặc cơ chế idempotency để tránh xử lý một yêu cầu hai lần khi ứng dụng retry.

## Write-Ahead Log (WAL)

- Để đảm bảo Durability của ACID, hệ thống không thể chỉ ghi vào data file rồi mới commit.
- WAL là cơ chế:

Ghi log trước → mới ghi dữ liệu thật.

- Mọi thay đổi đều được ghi vào WAL (sequential write – rất nhanh) trước khi được apply vào data page.
- Mục đích:
  - Crash recovery: Khi hệ thống sập, đọc lại WAL để redo những transaction đã commit.
  - Point-in-time recovery (PITR).
  - Replication (streaming replication dựa trên WAL).
- Trường hợp đặc biệt:
  - fsync trên WAL quyết định độ an toàn vs tốc độ.
  - Checkpoint: định kỳ đẩy dữ liệu từ bộ nhớ đệm xuống đĩa và cắt bỏ phần WAL cũ.
  - Trong PostgreSQL, WAL cũng dùng cho logical decoding, CDC…
