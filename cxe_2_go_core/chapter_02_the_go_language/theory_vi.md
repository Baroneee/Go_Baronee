## CÁC KHÁI NIỆM KỸ THUẬT ĐƯỢC ĐỀ CẬP

- Thời gian biên dịch (Build time)
- Ngôn ngữ kiểu tĩnh (Statically typed language)
- Tính đồng thời (Concurrency)
- Bộ thu gom rác (Garbage collector)
- Phần phụ thuộc phần mềm (Software dependency)

## NGUỒN GỐC VỀ SỰ RA ĐỜI

- Có một truyền thuyết xung quanh sự ra đời của Go. Ngôn ngữ này ra đời bên trong một văn phòng của Google, và nó xuất hiện trong một quá trình biên dịch rất dài mất tới 45 phút.
- Câu chuyện này được kể bởi Rob Pike trong [@go-at-google]. Nó cung cấp cho chúng ta một thông tin quý giá về những động lực đằng sau sự ra đời của Go. Thời gian biên dịch quá lâu và gây đau đầu... họ buộc phải tìm cách tránh khỏi điều đó; đó chính là điểm khởi nguồn của Go.
- Robert Griesemer, Ken Thompson và Rob Pike là những lập trình viên đã bắt đầu làm việc với Go từ năm 2007. Rob Pike tuyên bố rằng đến giữa năm 2008, ngôn ngữ này đã “được thiết kế phần lớn và việc triển khai (trình biên dịch, môi trường thời gian chạy) bắt đầu hoạt động”. Sau đó, Ian Lance Taylor và Russ Cox đã gia nhập đội ngũ vào năm 2008 [@pike2009go].
- Go là một ngôn ngữ lập trình mã nguồn mở được duy trì bởi cộng đồng và một đội ngũ cốt lõi các lập trình viên làm việc tại Google. Ngày 16 tháng 3 năm 2011 là ngày phát hành phiên bản Go đầu tiên (Nó được đặt tên là “r56”). Phiên bản Go 1 được phát hành vào ngày 28 tháng 3 năm 2012.

## ĐỘNG LỰC

- Go (hoặc Golang) được Google xây dựng để giải quyết các vấn đề của công ty lớn.
- Những thách thức đối với phần mềm tại các công ty lớn trên toàn cầu là gì?
  - Cơ sở mã (codebase) của các dịch vụ Google rất khổng lồ. Google có hàng triệu dòng mã.
  - Những dòng mã đó được viết bằng các ngôn ngữ khác nhau: C, C++, Java và các ngôn ngữ khác.
  - Thời gian biên dịch của các ứng dụng đó “đã kéo dài tới nhiều phút, thậm chí hàng giờ”.
  - Việc cập nhật một số phần của ứng dụng có thể rất tốn kém.
- Mục tiêu của những Gophers đầu tiên là làm cho cuộc sống của các lập trình viên trở nên dễ dàng hơn bằng cách:
  - Giảm đáng kể thời gian biên dịch của các chương trình.
  - Thiết kế một ngôn ngữ dễ học, dễ đọc và dễ gỡ lỗi (debug) cho các lập trình viên trẻ đã tiếp xúc với C, C++, hoặc Java.
  - Thiết kế một hệ thống quản lý phần phụ thuộc (dependency management) hiệu quả.
  - Xây dựng một ngôn ngữ có thể tạo ra phần mềm có khả năng mở rộng (scale) tốt trên phần cứng.

### ĐỊNH NGHĨA MỘT SỐ KHÁI NIỆM

- Thời gian biên dịch (Build time): lượng thời gian cần thiết để trình biên dịch tạo ra một tệp thực thi máy có thể đọc được.
- Ngôn ngữ kiểu tĩnh (Statically Typed Language): Đưa ra định nghĩa chính xác về khái niệm này lúc này còn quá sớm.
- Phần phụ thuộc (Dependency): một phần mềm được sử dụng bởi một phần mềm khác.
- Khả năng mở rộng (Scalability): khả năng của một chương trình trong việc xử lý lượng tác vụ ngày càng tăng cần thực hiện. Ví dụ, một trang mạng được coi là có khả năng mở rộng nếu nó có thể chấp nhận số lượng yêu cầu ngày càng tăng mà không bị sập (downtime) hoặc tăng độ trễ tải.

## CÁC TÍNH NĂNG CỐT LÕI CỦA GO

- Những người sáng tạo ra Go đã tập trung nỗ lực vào một số lựa chọn thiết kế quan trọng:
  - Là một ngôn ngữ biên dịch.
  - Có ngữ nghĩa dễ hiểu và dễ học.
  - Kiểu tĩnh (Statically typed).
  - Có tích hợp sẵn tính đồng thời (concurrency), một hệ thống giúp các lập trình viên làm việc dễ dàng hơn.
  - Quản lý phần phụ thuộc mạnh mẽ.
  - Có bộ thu gom rác (garbage collector).
- Mục tiêu chính, như Rob Pike đã trình bày, là cung cấp cho các lập trình viên một ngôn ngữ dễ học dùng cho việc “kỹ thuật hóa các dự án phần mềm lớn”.

### MỘT SỐ KHÁI NIỆM

- Tính đồng thời (Concurrency): Một chương trình có tính đồng thời khi các tác vụ có thể được thực thi không theo thứ tự hoặc theo một thứ tự một phần
- Bộ thu gom rác (Garbage collector - thường gọi là GC): Khi xây dựng các chương trình, chúng ta cần lưu trữ dữ liệu và lấy dữ liệu từ bộ nhớ. Bộ nhớ không phải là nguồn tài nguyên vô hạn. Do đó, lập trình viên phải đảm bảo rằng các phần tử không sử dụng được lưu trữ trong bộ nhớ sẽ bị xóa bỏ theo thời gian. Việc đưa dữ liệu vào bộ nhớ gọi là cấp phát (allocation); hành động ngược lại, bao gồm việc xóa dữ liệu khỏi bộ nhớ, gọi là thu hồi (deallocation). Vai trò của bộ thu gom rác là thu hồi bộ nhớ khi nó không còn được sử dụng nữa. Khi ngôn ngữ không có bộ thu gom rác nào, lập trình viên phải tự quản lý rác của mình và giải phóng bộ nhớ không còn dùng tới... Thật may mắn, Go có một bộ thu gom rác.

## TỰ KIỂM TRA

- Tính đồng thời (concurrency) có nghĩa là gì?
  - Là khi các tác vụ của 1 chương trình có thể thực thi cùng một thời điểm mà không có thứ tự hoặc theo thứ tự 1 phần
- Trung bình, Go có thời gian biên dịch rất dài phải không? Đúng hay Sai?
  - Sai, Go được tạo ra để giải quyết chính xác vấn đề này!

## CÁC ĐIỂM CỐT LÕI CẦN GHI NHỚ

- Go ra đời năm 2007.
- Phiên bản Go 1 được phát hành năm 2012.
- Ngôn ngữ này dễ hiểu. Ngữ nghĩa của nó vẫn giữ được sự đơn giản.
- Nó có kiểu tĩnh (statically typed).
- Nó là ngôn ngữ biên dịch (compiled).
- Bạn có thể viết các chương trình đồng thời với Go.
