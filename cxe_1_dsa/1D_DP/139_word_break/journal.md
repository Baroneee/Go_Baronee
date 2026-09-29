### LeetCode 139: Word Break — Personal Journal

---

#### 🇻🇳 [VIETNAMESE VERSION]

##### Ý tưởng cốt lõi (Core Insight)

Một chuỗi `s` có thể ghép thành công nếu tồn tại ít nhất một điểm cắt `i` sao cho tiền tố `s[:i]` ghép được thành công và phần còn lại `s[i:]` bắt đầu bằng một từ hợp lệ trong `wordDict`.

##### Tối ưu thuật toán

- **Version 1: Dynamic Programming (Bottom-Up - O(N \* M) Time, O(N) Space)**
  - **Ý tưởng:** Gọi `dp[i]` là trạng thái cho biết chuỗi con `s[:i]` có thể ghép được từ `wordDict` hay không.
  - **Thực thi:** Khởi tạo `dp[0] = true` (chuỗi rỗng). Duyệt qua từng chỉ số `i` từ `0` đến `len(s)-1`. Nếu `dp[i] == true`, kiểm tra xem đoạn `s[i:]` có bắt đầu bằng từ `w` nào trong `wordDict` hay không (`strings.HasPrefix`). Nếu có, đánh dấu `dp[i + len(w)] = true`. Kết quả cuối cùng nằm ở `dp[len(s)]`.

- **Version 2: Brute Force Backtracking (Top-Down DFS)**
  - **Ý tưởng:** Xây dựng cây quyết định (Decision Tree), tại mỗi vị trí `start`, thử ghép từng từ trong `wordDict`. Nếu hợp lệ, đệ quy đi tiếp tới `start + len(w)`.
  - **Thực thi:** Dùng hàm đệ quy `backtrack(start)`. Nếu `start == len(s)`, trả về `true`. Ngược lại, duyệt qua từng từ `w`. Nếu `s[start:]` có tiền tố `w`, gọi đệ quy `backtrack(start + len(w))`. Nếu nhánh đệ quy trả về `false`, thuật toán tự động lùi lại (backtrack) để thử từ tiếp theo.

---

#### 🇬🇧 [ENGLISH VERSION]

##### Key Insight & Approach

A string `s` can be successfully segmented if there exists at least one partition point `i` such that the prefix `s[:i]` is segmentable and the remaining suffix `s[i:]` starts with a valid word from `wordDict`.

##### Optimizing Solutions

- **Version 1: Dynamic Programming (Bottom-Up - O(N \* M) Time, O(N) Space)**
  - **Idea:** Define `dp[i]` as a boolean indicating whether the substring prefix `s[:i]` can be formed using words from `wordDict`.
  - **Implementation:** Initialize `dp[0] = true` (empty string). Iterate through index `i` from `0` to `len(s)-1`. If `dp[i]` is `true`, check if `s[i:]` starts with any word `w` in `wordDict` (`strings.HasPrefix`). If matched, set `dp[i + len(w)] = true`. The answer is stored in `dp[len(s)]`.

- **Version 2: Brute Force Backtracking (Top-Down DFS)**
  - **Idea:** Build a decision tree where at each `start` position, try matching every word in `wordDict`. If valid, recursively proceed to `start + len(w)`.
  - **Implementation:** Use a recursive `backtrack(start)` function. Base case: if `start == len(s)`, return `true`. Otherwise, iterate through each word `w`. If `s[start:]` starts with `w`, call `backtrack(start + len(w))`. If a branch returns `false`, it backtracks to test the next word in `wordDict`.
