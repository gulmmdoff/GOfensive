Quick Start
1. Build & Run Server

git clone https://github.com/gulmmdoff/GOfensive.git
cd GOfensive
go run cmd/server/main.go
2. Compile Agent
Agenti hədəfə uyğun olaraq aşağıdakı kimi compile edin:
 
go build -o agent.exe ./cmd/agent/main.go

3. Command Execution
Server qalxdıqdan sonra http://0.0.0.0:8080 ünvanına daxil olun, agenti seçin və əmrləri icra edin:

Plaintext
whoami
hostname
dir



Gələcək Planlar:
Bu layihənin v1.0 (Beta) versiyasıdır. Gələcək yenilənmələrdə upload/download modulları, in-memory execution və daha mürəkkəb və yeni modullar əlavə edilərək tam upgradeable olacaqdır.
