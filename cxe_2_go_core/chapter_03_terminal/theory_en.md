## CÁC KHÁI NIỆM KỸ THUẬT ĐƯỢC ĐỀ CẬP

- GUI (Graphical User Interface): Giao diện người dùng đồ họa.
- CLI (Command Line Interface): Giao diện dòng lệnh.
- Shell: Trình thông dịch dòng lệnh.
- Bash: Một loại shell phổ biến (Bourne Again Shell).
- WSL (Windows Subsystem For Linux): Hệ thống con Windows dành cho Linux.

## GIỚI THIỆU

- Cuộc sống hằng ngày của một lập trình viên Go đòi hỏi việc sử dụng terminal. Chương trình này sẽ giải thích terminal là gì và cách sử dụng nó. Nếu bạn đã quen thuộc với nó, bạn có thể bỏ qua chương này.

## GIAO DIỆN NGƯỜI DÙNG ĐỒ HOẠ (GUI)

- Các hệ điều hành như macOS, Windows hoặc Linux cung cấp giao diện người dùng đồ họa (GUI) rất phong phú. Để khởi chạy một chương trình đã được cài đặt trên máy tính, bạn thường nhấp đúp chuột vào một biểu tượng nằm trên màn hình nền (desktop). Các chương trình sẽ được hiển thị trong các cửa sổ có giao diện tương tác: menu, thanh bên, nút bấm...
- Chúng ta nói rằng các chương trình đó cung cấp giao diện người dùng đồ họa (GUI). Phần lớn người dùng sẽ sử dụng các chương trình cung cấp các giao diện này. Giao diện đồ họa rất dễ sử dụng và trực quan.

## GIAO DIỆN DÒNG LỆNH (CLI)

- Giao diện người dùng đồ họa không phải lúc nào cũng tồn tại. Những chiếc máy tính đầu tiên không có khả năng đó. Nhưng làm thế nào để người dùng của những chiếc máy tính đó khởi chạy và sử dụng các chương trình? Máy tính được trang bị giao diện dòng lệnh (Command-line interface). Giao diện này còn được gọi là “shell”.
- Shell là một chương trình có thể truyền các câu lệnh đến hệ điều hành. Shell là một thuật ngữ chung chỉ các chương trình như vậy. Shell nổi tiếng nhất là bash (Bourne Again Shell). Bash được tích hợp sẵn theo mặc định trên macOS và phần lớn các bản phân phối Linux. Windows cũng được tích hợp sẵn một shell theo mặc định (nhưng không phải là bash).

## CÁCH TƯƠNG TÁC VỚI SHELL: TERMINAL

- Shell từng được hiển thị trực tiếp cho người dùng ngay sau khi khởi động trên các máy tính cũ. Trên máy tính hiện đại, chúng ta phải khởi chạy một chương trình để tương tác với shell. Chương trình này thường được gọi là terminal. Chúng ta sẽ xem cách mở một terminal trên macOS, Linux (GNOME) và Windows.

### macOS

- Mở ứng dụng Finder (biểu tượng Finder).
- Trên thanh menu, nhấp vào “Go” và sau đó chọn “Utilities”.
- Một cửa sổ sẽ mở ra. Nhấp vào “Terminal” (Ứng dụng Terminal).
- Một cửa sổ terminal mới sẽ mở ra.

### Linux (Ubuntu)

- Trên Ubuntu, bạn có thể sử dụng phím tắt Ctrl + Alt + T.
- Bạn cũng có thể khởi chạy terminal thông qua Ubuntu Dash. Nhập từ khóa “Terminal” và ứng dụng sẽ xuất hiện.

### Windows

- Nhấp vào nút Start, sau đó trong hộp văn bản, nhập “cmd” (dành cho command prompt).
- Nhấp vào ứng dụng cmd.
- Một cửa sổ màu đen sẽ xuất hiện; đó chính là terminal của bạn!

#### Cmder

- Terminal và Windows shell mặc định không thực sự thiết thực trong việc sử dụng hằng ngày. Tôi khuyên bạn nên cài đặt Cmder để giúp cuộc sống lập trình viên trên Windows của bạn trở nên dễ dàng hơn. Cmder là một trình giả lập (emulator) cho phép bạn sử dụng các câu lệnh có sẵn trên Linux/macOS. Quá trình cài đặt rất dễ dàng (tải xuống phiên bản mới nhất trên GitHub), sau đó chạy trình hướng dẫn cài đặt.
- Sau khi cài đặt Cmder, hãy khởi chạy chương trình “Cmder” để mở terminal hoàn toàn mới của bạn.

#### Bash trên Windows

- Theo mặc định, bạn không thể sử dụng bash trên máy tính Windows. Đây không phải là vấn đề lớn, nhưng điều đó có nghĩa là bạn sẽ phải tìm các câu lệnh tương đương trên Windows cho mỗi câu lệnh macOS/Linux. Việc này có thể trở nên rườm rà vào một thời điểm nào đó vì rất nhiều ví dụ và hướng dẫn trên web không phải lúc nào cũng cung cấp các câu lệnh tương đương cho Windows.
- Microsoft đã thông báo rằng giờ đây bạn có thể cài đặt “Windows Subsystem for Linux” (WSL) trên máy tính Windows của mình. Đây là một tin vui vì khi đó bạn có thể sử dụng bash. Bạn có thể tìm thấy hướng dẫn cài đặt trên trang web của Microsoft.
- Tôi thực sự khuyên bạn nên cài đặt tính năng này, bởi vì nó sẽ làm cho cuộc sống của bạn dễ dàng hơn ngay cả khi tôi cố gắng cung cấp các lệnh tương đương trên Windows cho các câu lệnh cơ bản trong các phần tiếp theo.
