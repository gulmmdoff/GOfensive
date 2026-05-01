![image alt](https://github.com/gulmmdoff/GOfensive/blob/2d6c7f408b1773e9c0c37be5a49f1c9092e5febc/GOfensive.png).
# ⚔️ GOfensive: Lightweight & Malleable C2 Framework

**GOfensive**, real dünya kiber-hücum infrastrukturunu simulyasiya etmək üçün hazırlanmış modern bir Command & Control (C2) sistemidir. Go dilinin gücündən istifadə edərək, hədəf sistemlərdə minimal iz buraxan və genişləndirilə bilən modulyar bir arxitektura təqdim edir.

---

## 🚀 Features

- **Agent-Based Architecture:** Yüngül və effektiv agent idarəetməsi.
- **Remote Shell Execution:** Hədəf sistemlərdə birbaşa əmr icrası.
- **Real-time Monitoring:** Aktiv agentlərin anlıq izlənilməsi və status hesabatı.
- **Minimalist Web UI:** Sadə, sürətli və effektiv idarəetmə paneli.
- **Zero-Dependency Agent:** Heç bir xarici kitabxanadan asılı olmayan Go implementation.
- **Lab & CTF Ready:** Xüsusi laboratoriya və CTF mühitləri üçün optimallaşdırılmış struktur.

---

## 🏗️ Project Structure

- `cmd/` — Giriş nöqtələri: Server və Agent-in ana kodları.
- `internal/` — Core məntiq: Handler-lər, modellər və təhlükəsizlik alqoritmləri.
- `ui/` — Frontend: Operatorun idarəetmə paneli.

---

## 🌐 Architecture

GOfensive üç tərəfli sinxron əlaqə üzərində qurulub:
`Agent <----HTTP----> C2 Server <----> Web UI`

1. **Agents:** Müəyyən intervallarla serverə "check-in" edir və tapşırıqları götürür.
2. **Server:** Tapşırıqların idarə edilməsi, cavabların emalı və bazada saxlanılması.
3. **UI:** Operatorun sistemləri vizual idarə etməsi üçün interfeys.

---

## ⚡ Quick Start

### 1. Repository-ni klonlayın
```bash
git clone [https://github.com/gulmmdoff/GOfensive.git](https://github.com/gulmmdoff/GOfensive.git)
cd GOfensive
