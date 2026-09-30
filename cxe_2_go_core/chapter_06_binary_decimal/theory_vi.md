## CÁC KHÁI NIỆM KỸ THUẬT ĐƯỢC ĐỀ CẬP

- Thông số kỹ thuật (Specifications)
- Biên dịch (Compilation)
- Tệp nhị phân/ Tệp thực thi (A binary/ an executable)

## THÔNG SỐ KỸ THUẬT CỦA ỨNG DỤNG (APPLICATION SPECIFICATION)

- Trước khi viết bất kỳ mã nguồn nào, chúng ta cần quyết định xem ứng dụng của mình sẽ làm gì. Hầu hết các dự án đều bắt đầu bằng giai đoạn này. Đây gọi là giai đoạn lập thông số kỹ thuật. Giai đoạn này nhằm đưa ra các yêu cầu chính xác mà ứng dụng cần phải đáp ứng. Những yêu cầu đó chính là các thông số kỹ thuật (specs).
- Thông số kỹ thuật cho ứng dụng của chúng ta rất đơn giản: khi được khởi chạy, ứng dụng sẽ hiển thị ngày giờ và sau đó thoát.

## THƯ MỤC DỰ ÁN (PROJECT DIRECTORY)

- Một ứng dụng Go được cấu thành từ các tệp tin. Trên các tệp đó, chúng ta sẽ viết mã Go. Chúng ta gọi những tệp này là “tệp mã nguồn” (source files). Một ứng dụng được lưu trữ trong một thư mục chính. Thư mục chính này có thể chỉ chứa một tệp mã nguồn duy nhất. Tuy nhiên, phần lớn thời gian, nó bao gồm một vài thư mục con.
- Chúng ta sẽ tạo thư mục chính này cho ứng dụng của mình. Bạn có thể thực hiện việc này bằng dòng lệnh:
  $ cd Documents/code
  $ mkdir dateAndTime

## MÔI TRƯỜNG PHÁT TRIỂN (IDE)

- Về lý thuyết, bạn có thể viết mã bằng một trình soạn thảo văn bản tiêu chuẩn. Tuy nhiên, trên thị trường có những phần mềm chuyên dụng được phát triển đặc biệt cho lập trình viên. Chúng được gọi là IDE (Integrated Development Environment - Môi trường phát triển tích hợp).
- IDE cung cấp các tính năng như:
  - Tô màu tự động cho các từ khóa (làm nổi bật cú pháp - syntax highlighting)
  - Tự động hoàn thành mã (autocompletion)
  - Khả năng tái cấu trúc mã (refactoring)

## Tệp mã nguồn (Source file)
