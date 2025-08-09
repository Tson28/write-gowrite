# Hướng dẫn Upload lên GitHub

## Cách 1: Sử dụng GitHub Desktop
1. Tải GitHub Desktop từ: https://desktop.github.com/
2. Clone repository: git@github.com:Tson28/write-gowrite.git
3. Copy tất cả files trong thư mục này vào repository
4. Commit và Push

## Cách 2: Sử dụng GitHub Web Interface
1. Tạo repository mới trên GitHub: https://github.com/Tson28/write-gowrite
2. Upload files thủ công qua web interface
3. Hoặc sử dụng GitHub CLI nếu có

## Cách 3: Sử dụng GitHub API (cần token)
1. Tạo Personal Access Token trên GitHub
2. Chạy script với token:
   ```bash
   export GITHUB_TOKEN='your_token_here'
   ./upload_to_github.sh
   ```

## Files đã được chuẩn bị:
- Tất cả files .go với hashtag #author: Nguyễn Thái Sơn
- go.mod và go.sum
- README.md và LICENSE
- Example program trong thư mục example/
- Cấu hình Git cơ bản

## Tác giả: Nguyễn Thái Sơn
