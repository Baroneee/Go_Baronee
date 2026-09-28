### LeetCode 746: Min Cost Climbing Stairs — Personal Journal

---

#### 🇻🇳 [VIETNAMESE VERSION]

##### Ý tưởng

Để bước qua bậc thứ $N$ (đỉnh cầu thang), ta có 2 lựa chọn:

1. Bước từ bậc $N-1$ sang (trả chi phí của bậc $N-1$).
2. Bước từ bậc $N-2$ sang (trả chi phí của bậc $N-2$).

Vì muốn tối thiểu hóa tổng chi phí, chi phí nhỏ nhất để bước qua bậc $N$ sẽ là:
`dp[N] = min(dp[N-1] + cost[N-1], dp[N-2] + cost[N-2])`

##### Tối ưu thuật toán

- **Version 1: Dynamic Programming dùng mảng (O(N) Space)**
  - **Ý tưởng:** Gọi `dp[i]` là chi phí tối thiểu để bước qua bậc thứ `i`.
  - **Thực thi:** Dùng một mảng `dp` kích thước `N + 1` để lưu lại trạng thái của từng bậc, tính toán lần lượt từ `2` đến `N`.

- **Version 2: Tối ưu bộ nhớ (O(1) Space)**
  - **Ý tưởng:** Nhận thấy tính kế thừa của bài toán — để tính giá trị tại bậc hiện tại (`curr`), ta chỉ cần đúng 2 trạng thái ngay trước nó (`curr` cũ và `prev`).
  - **Thực thi:** Thay vì lưu toàn bộ mảng `dp`, ta chỉ dùng 2 biến `prev` và `curr` để mô phỏng lại quá trình dịch chuyển DP, giúp tiết kiệm bộ nhớ tối đa mà không cần lưu lại các kết quả cũ không dùng tới.

---

#### 🇬🇧 [ENGLISH VERSION]

##### Key Insight & Approach

To reach past step N (the top of the staircase), there are two possible options:

1. Step from step N - 1 (paying its cost).
2. Step from step N - 2 (paying its cost).

Since we want to minimize the total effort, the minimum cost to reach step N is:
`dp[N] = min(dp[N-1] + cost[N-1], dp[N-2] + cost[N-2])`

##### Optimizing Solutions

- **Version 1: Full Dynamic Programming Array (O(N) Space)**
  - **Idea:** Define `dp[i]` as the minimum cost to reach step `i`.
  - **Implementation:** Use an array `dp` of size `N + 1` to store the state for every step iteratively from `2` to `N`.

- **Version 2: Space Optimization (O(1) Space)**
  - **Idea:** Notice the overlapping subproblem pattern — calculating the current state `curr` only requires the two immediately preceding states (`curr` and `prev`).
  - **Implementation:** Replace the full `dp` array with two tracking variables, `prev` and `curr`, to simulate the DP transitions in constant space without storing unnecessary past values.
