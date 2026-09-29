## CÁC THÀNH PHẦN PHẦN CỨNG

- Một máy tính gồm 4 thành phần chính: + Đơn vị bộ nhớ (MU): nơi lưu trữ dữ liệu và các chương trình. + Đơn vị số học và logic (ALU): Thực hiện các phép toán số học và login trên dữ liệu được lưu trữ trong MU.
  - Đơn vị nhập và xuất (I/OU): Chịu trách nhiệm nạp dữ liệu vào MU từ 1 thiết bị đầu vào. Và nó cũng gửi dữ liệu từ đơn vị bộ nhớ đến một thiết bị đầu ra + Đơn vị điều khiển (CU): Nhận các chỉ thị từ chương trình và điều khiển hoạt động của các đơn vị khác.
    -> 4 thành phần này biểu diễn một sơ đồ cấu trúc của các thành phần trong máy tính.

## BỘ NHỚ

- Máy tính bao gồm 2 bộ nhớ:
  - Bộ nhớ trung tâm
  - Bộ nhớ phụ
- Có 2 phân loại bộ nhớ:
  - Dễ bay hơi (Volatile)
  - Không dễ bay hơi (Non-volatile)

### BỘ NHỚ TRUNG TÂM

- Chứa 2 loại lưu trữ:
  - Bộ nhớ truy cập ngẫu nhiên (RAM): Cần nguồn điện để duy trì dữ liệu. Khi tắt máy tính, dữ liệu trong bộ nhớ này sẽ bị xoá. Hệ điều hành và các chương tình được sử dụng sẽ nạp vào bộ nhớ này. Loại bộ nhớ này có tính chất là Volatile
  - Bộ nhớ chỉ đọc (ROM): Đây là bộ nhớ chưa dữ liệu cần thiết để máy tính chạy chính xác. Loại bộ nhớ này có tính chất Non-volatile (Khi tắt máy tính dữ liệu sẽ không bị xoá). Nó được thiết kế để chỉ đọc và hệ thống không thể cập nhật

### BỘ NHỚ PHỤ

- Loại bộ nhớ này có tính chất Non-volatile. Khi mất nguồn điện, dữ liệu được lưu trữ sẽ không bị xoá. Ví dụ: USB, Ổ cứng, CD-ROM, DVD,...
- Việc đọc và ghi đối với loại bộ nhớ này chậm hơn so với RAM.
- Một số ổ đĩa cứng truy cập bộ nhớ tuần tự -> Hệ thống phải tuân theo một trình tự cụ thể. Việc này mất nhiều thời gian hơn so với chế độ truy cập ngẫu nhiên. Lưu ý rằng vẫn có 1 số ổ đĩa cứng cho phép truy cập ngẫu nhiên.

#### Ổ ĐĨA CỨNG

- Ổ đĩa cứng hay còn gọi là ổ đĩa cứng cơ (HDD), bao gồm các đĩa từ tính quay tròn. Dữ liệu được đọc và ghi nhờ một đầu đọc ghi từ tính di chuyển. Các thao tác đọc và ghi sẽ tạo ra chuyển động quay và di chuyển của đầu từ -> từ đó tiêu tốn thời gian.
- Ổ đĩa thể rắn (SSD) không được cấu tạo như vậy. Không có đầu từ tính hay đĩa từ tính nào cả. Thay vào đó, dữ liệu được lưu trữ trong các ô nhớ flash. Việc truy cập dữ liệu sẽ nhanh hơn trên loại đĩa này.

## CPU

- CPU là viết tắt của Central Processing Unit (Bộ xử lý trung tâm). CPU còn được gọi là bộ vi xử lý (processor). CPU bao gồm:
  - ALU
  - CU
- CPU chịu trách nhiệm thực thi các chỉ thị do một chương trình đưa ra. Ví dụ, chương trình có thể yêu cầu thực hiện phép cộng giữa hai số. Các số đó sẽ được truy xuất từ đơn vị bộ nhớ và chuyển đến ALU. Chương trình cũng có thể yêu cầu thực hiện thao tác nhập/xuất như đọc dữ liệu từ ổ đĩa cứng và nạp nó vào RAM để xử lý thêm. CPU sẽ thực thi các chỉ thị đó.
- CPU là thành phần trung tâm của một máy tính.

## CHƯƠNG TRÌNH LÀ GÌ?

- Để làm cho máy tính thực hiện điều gì đó, chúng ta phải cung cấp cho chúng các chỉ thị chính xác. Tập hợp các chỉ thị này được gọi là "chương trình" (program).
- Theo một định nghĩa chính thống hơn, chương trình là "sự kết hợp giữa các chỉ thị máy tính và định nghĩa dữ liệu cho phép phần cứng máy tính tính toán".
- Chúng ta sẽ đưa ra các chỉ thị theo ngôn ngữ con người thông thường NHƯNG máy tính không hiểu các câu tiếng người. Những câu này cần được dịch sang ngôn ngữ mà máy hiểu được -> Ngôn ngữ đó là gì?

## LÀM THỂ NÀO ĐỂ GIAO TIẾP ĐƯỢC VỚI MÁY TÍNH?

### NGÔN NGỮ LẬP TRÌNH LÀ NGÔN NGỮ HÌNH THỨC

- Các chỉ thị được đưa ra cho máy tính được viết bằng các ngôn ngữ lập trình. Ngôn ngữ lập trình là các ngôn ngữ hình thức (formal languages). Chúng bao gồm các từ được cấu tạo từ một bảng chữ cái (một tập hợp các ký tự riêng biệt). Các từ đó được tổ chức tuân theo các quy tắc cụ thể. Go là một ngôn ngữ lập trình.
- Có 2 loại ngôn ngữ lập trình:
  - Cấp thấp (Low level)
  - Cấp cao (High level)
- Ngôn ngữ lập trình cấp thấp gần gũi hơn với các chỉ thị của bộ xử lý. Ngôn ngữ cấp cao cung cấp các cấu trúc giúp chúng dễ học và dễ sử dụng hơn trong công việc hàng ngày.
- Một số ngôn ngữ cấp cao được biên dịch, số khác được thông dịch, và một số nằm ở giữa. Chúng ta sẽ xem hai thuật ngữ này có ý nghĩa gì trong các phần tiếp theo.

### NGÔN NGỮ MÁY

- Để giao tiếp với đơn vị xử lý của máy tính, chúng ta có thể sử dụng ngôn ngữ máy (machine language). Ngôn ngữ máy bao gồm hoàn toàn các số 0 và 1. Một chỉ thị được viết bằng ngôn ngữ máy là một chuỗi các số 0 và 1. Mỗi bộ xử lý (hoặc họ bộ xử lý) sẽ định nghĩa một danh sách các chỉ thị được gọi là tập lệnh (instruction set). Có chỉ thị để cộng thêm một số, tăng lên một đơn vị, giảm đi một đơn vị, sao chép dữ liệu từ vị trí này trong bộ nhớ sang nơi khác...
- Việc viết chương trình máy tính trực tiếp bằng ngôn ngữ máy là khả thi. Tuy nhiên, điều này không dễ dàng.

### NGÔN NGỮ HỢP NGỮ (ASSEMBLY)

- Ngôn ngữ hợp ngữ là một ngôn ngữ lập trình cấp thấp. Các chỉ thị của một chương trình viết bằng hợp ngữ tương ứng với các chỉ thị máy. Ngôn ngữ hợp ngữ sử dụng các từ ký hiệu viết tắt (mnemonics) tương ứng với một chỉ thị máy. Ví dụ, MOV sẽ hướng dẫn máy tính di chuyển dữ liệu từ vị trí này sang vị trí khác. Các nhà phát triển cũng có thể chú thích (comment) vào mã nguồn (điều không thể làm được với ngôn ngữ máy).
- Để tạo một chương trình bằng ngôn ngữ hợp ngữ, lập trình viên sẽ viết các chỉ thị vào một hoặc nhiều tệp. Các tệp này được gọi là tệp nguồn (source files).
- Dưới đây là một ví dụ về một chỉ thị được viết bằng hợp ngữ x86 Linux:
  mov eax,1
  int 0x80
- Hai dòng này sẽ thực hiện một lệnh gọi hệ thống (system call) nhằm đóng chương trình (số "1" đại diện cho số hiệu lệnh gọi hệ thống có nghĩa là "thoát chương trình"). Lưu ý rằng ngôn ngữ hợp ngữ khác nhau giữa các máy khác nhau. Chúng ta gọi nó là phụ thuộc vào máy (machine-specific).
- Một chương trình dịch hợp ngữ (assembler) được sử dụng để chuyển đổi các tệp nguồn được viết bằng ngôn ngữ hợp ngữ thành các tệp mã đối tượng (object code files). Chúng ta gọi quá trình này là hợp dịch chương trình. Bộ liên kết (linker) sau đó sẽ biến đổi các tệp mã đối tượng này thành một tệp thực thi (executable file). Một tệp thực thi chứa tất cả các chỉ thị cần thiết của máy tính để khởi chạy chương trình.
  ![Quy trình dịch mã Assembly](https://www.practical-go-lessons.com/img/assembly_to_executable.926437cf.png)

### CÁC NGÔN NGỮ CẤP CAO

- Có rất nhiều ngôn ngữ cấp cao trên thị trường, như Go. Những ngôn ngữ này không bị ràng buộc chặt chẽ với kiến trúc máy. Chúng cung cấp một cách thức thuận tiện để viết các chỉ thị. Ví dụ, nếu chúng ta muốn thực hiện một lệnh gọi hệ thống để thoát chương trình, bằng Go chúng ta có thể viết:
  os.Exit(1)
- Với ngôn ngữ C, chúng ta có thể viết:
  exit(1)
- Với Java, chúng ta có thể viết:
  System.exit(1);
- Trong ví dụ này, chúng ta không phải di chuyển một số vào một thanh ghi (register); chúng ta sử dụng các cấu trúc của ngôn ngữ (hàm, gói, phương thức, biến, kiểu dữ liệu...). Mục tiêu của cuốn sách này là cung cấp cho bạn các định nghĩa chính xác và ngắn gọn về các công cụ này để xây dựng các ứng dụng Go.
- Các chương trình cấp cao cũng được viết thành các tệp. Các tệp này được gọi là "tệp nguồn" (source files). Nhìn chung, các ngôn ngữ lập trình yêu cầu thêm một phần mở rộng cụ thể vào tên tệp. Đối với các chương trình Go, chúng ta sẽ thêm .go vào cuối mỗi tệp mà chúng ta viết. Trong PHP, phần mở rộng là .php.
- Khi các tệp nguồn được viết, chương trình do chúng định nghĩa không thể được thực thi ngay lập tức. Tệp nguồn cần được biên dịch bằng cách sử dụng một trình biên dịch (compiler). Trình biên dịch sẽ biến đổi các tệp nguồn thành một tệp thực thi. Trình biên dịch cũng là một chương trình. Go là một phần của gia đình ngôn ngữ biên dịch.
- Go là một ngôn ngữ biên dịch

![Quy trình Trình dịch (Compiler)](https://www.practical-go-lessons.com/img/compiler.adc1a3b1.png)

#### BIÊN DỊCH VS THÔNG DỊCH

- Lưu ý rằng một số ngôn ngữ lập trình là ngôn ngữ thông dịch (interpreted). Khi các tệp nguồn đã được viết, lập trình viên không cần phải biên dịch mã nguồn. Với các tệp nguồn đã sẵn sàng, hệ thống có thể thực thi chương trình nhờ một trình thông dịch (interpreter). Mỗi chỉ thị được viết trong tệp nguồn được dịch và thực thi bởi trình thông dịch. Trong một số trường hợp, các trình thông dịch lưu trữ một phiên bản đã biên dịch của chương trình trong bộ nhớ đệm (cache) để tăng hiệu suất (các tệp nguồn không bị dịch lại mỗi lần). PHP, Python, Ruby, Perl là các ngôn ngữ thông dịch.

## TỰ KIỂM TRA

### CÁC CÂU HỎI

##### Các chương trình được lưu trữ ở đâu ?

- Ở các đơn vị bộ nhớ (MU)

##### Đọc dữ liệu từ ổ đĩa cứng chậm hơn đọc dữ liệu từ RAM. Đúng hay sai?

- Đúng vì ổ đĩa cứng cần thời gian để các đầu đọc và đĩa từ xoay để và đọc còn RAM thì chỉ cần lấy từ các ô nhớ có sẵn.

##### Bạn có thể ghi dữ liệu vào ROM không? Đúng hay sai?

- Sai bởi vì ROM là Read-Only Memory nghĩa là chúng ta chỉ có thể đọc dữ liệu từ ROM chứ không thể ghi.

##### Hai loại bộ nhớ là gì?

- Bộ nhớ trung tâm () và bộ nhớ phụ

##### Định nghĩa của "bộ nhớ dễ bay hơi" (volatile memory) là gì?

- Là khi bị ngắt nguồn điện thì dữ liệu sẽ bị mất đi.

##### Chương trình nào chuyển đổi mã viết bằng ngôn ngữ hợp ngữ thành mã đối tượng?

- Chương trình dịch hợp ngữ (Assembler).

##### Chương trình nào chuyển đổi mã đối tượng thành tệp thực thi?

- Bộ liên kết (Linker)

##### Nêu hai ưu điểm của ngôn ngữ cấp cao so với ngôn ngữ cấp thấp?

- Chúng cung cấp các cấu trúc cấp cao dễ sử dụng hơn.
- Mã nguồn sẽ không phụ thuộc vào kiến trúc kỹ thuật của một máy cụ thể. Chúng ta gọi mã đó là tính di động (portable).

##### Go là một ngôn ngữ thông dịch? Đúng hay sai?

- Sai vì Go là một ngôn ngữ biên dịch (Complied language)

## Các điểm cốt lõi cần ghi nhớ

- Ở cấp độ vĩ mô, máy tính bao gồm:
  - Đơn vị bộ nhớ (MU): để lưu trữ dữ liệu và chương trình.
  - Đơn vị số học và logic (ALU): để thực hiện tính toán.
  - Đơn vị nhập và xuất (IOU): để quản lý các thiết bị đầu vào và thiết bị đầu ra.
  - Đơn vị điều khiển (CU): quản lý MU, ALU và IOU theo các chỉ thị được đưa ra bởi chương trình đang thực thi.
- CPU có nghĩa là Central Processing Unit (còn được gọi là bộ vi xử lý hoặc chip xử lý), bao gồm ALU và CU.
- Một chương trình là một tập hợp các chỉ thị.
- Lập trình viên viết chương trình bằng ngôn ngữ lập trình.
- Ngôn ngữ lập trình bao gồm các từ và ký tự phải được sắp xếp tuân theo các quy tắc chỉ định.
- Có các ngôn ngữ lập trình cấp cao và cấp thấp.
- Ngôn ngữ máy và ngôn ngữ hợp ngữ là cấp thấp. Các chỉ thị được viết bằng các ngôn ngữ này gắn kết chặt chẽ với tổ chức và năng lực của phần cứng. Chúng cung cấp rất ít tính trừu tượng.
- Go là một ngôn ngữ lập trình cấp cao và được biên dịch.
