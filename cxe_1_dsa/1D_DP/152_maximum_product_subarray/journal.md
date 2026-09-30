### LeetCode 152: Maximum Product Subarray — Personal Journal

---

#### 🇻🇳 [VIETNAMESE VERSION]

##### Ý tưởng cốt lõi (Core Insight)

Tích của mảng con phụ thuộc lớn vào dấu của số hiện tại. Khi gặp một số âm, giá trị cực đại có thể đảo thành cực tiểu và ngược lại. Do đó, bài toán yêu cầu lưu trữ song song hai trạng thái tại mỗi vị trí $i$: tích lớn nhất (`currMax`) và tích nhỏ nhất (`currMin`) của mảng con kết thúc đúng tại $i$.

##### Tối ưu thuật toán

- **Version 1: State-Optimized Dynamic Programming (O(N) Time, O(1) Space)**
  - **Ý tưởng:** Trạng thái tại $i$ hoàn toàn phụ thuộc vào trạng thái $i-1$. Thay vì dùng mảng DP 1D $O(N)$ để lưu tất cả giá trị, ta tối ưu bộ nhớ xuống $O(1)$ bằng cách chỉ duy trì 2 biến `currMax` và `currMin`. Khi gặp số âm, hoán đổi hai biến này trước khi tính toán để đảm bảo tính đúng đắn.
  - **Thực thi:** Khởi tạo `currMin`, `currMax`, `ans` bằng `nums[0]`. Duyệt từ $i = 1$, nếu `nums[i] < 0`, hoán đổi `currMin` và `currMax`. Cập nhật `currMin = min(nums[i], currMin * nums[i])`, `currMax = max(nums[i], currMax * nums[i])` và liên tục cập nhật `ans = max(ans, currMax)`.

- **Version 2: Brute Force (O(N²) Time, O(1) Space)**
  - **Ý tưởng:** Duyệt qua tất cả các cặp chỉ số $[i, j]$ đại diện cho điểm bắt đầu và kết thúc của mảng con, tính tích của từng mảng con và tìm giá trị lớn nhất.
  - **Thực thi:** Dùng 2 vòng lặp lồng nhau. Vòng lặp ngoài cố định $i$, vòng lặp trong mở rộng $j$, tính tích dồn `currentProduct *= nums[j]` và cập nhật `ans`.

---

#### 🇬🇧 [ENGLISH VERSION]

##### Key Insight & Approach

The product of a subarray heavily depends on the sign of the current number. A negative number flips the maximum product into a minimum and vice versa. Thus, the problem requires tracking two concurrent states at each position $i$: the maximum product (`currMax`) and the minimum product (`currMin`) of the subarray ending strictly at $i$.

##### Optimizing Solutions

- **Version 1: State-Optimized Dynamic Programming (O(N) Time, O(1) Space)**
  - **Idea:** The state at index $i$ depends strictly on state $i-1$. Instead of maintaining $O(N)$ DP arrays, we optimize space to $O(1)$ by keeping only two state variables: `currMax` and `currMin`. When encountering a negative number, swap these variables prior to calculation to maintain state correctness.
  - **Implementation:** Initialize `currMin`, `currMax`, and `ans` with `nums[0]`. Iterate from $i = 1$. If `nums[i] < 0`, swap `currMin` and `currMax`. Update `currMin = min(nums[i], currMin * nums[i])`, `currMax = max(nums[i], currMax * nums[i])`, and track the global maximum with `ans = max(ans, currMax)`.

- **Version 2: Brute Force (O(N²) Time, O(1) Space)**
  - **Idea:** Evaluate all possible subarray index pairs $[i, j]$ (start and end bounds), compute their cumulative products, and record the overall maximum.
  - **Implementation:** Use two nested loops. The outer loop fixes the start bound $i$, while the inner loop expands the end bound $j$, updating `currentProduct *= nums[j]` and tracking the maximum value in `ans`.
