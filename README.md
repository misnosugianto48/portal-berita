# Portal Berita

Portal berita dibangun dengan **hexagonal architecture** (ports & adapters), yang memisahkan bisnis murni dari framework dan teknik eksternal.

## Susunan folder

```
.
├── cmd/
├── configs/
│   └── config.go
├── databases/
├── internal/
│   ├── adapters/          # Implementasi ports (concrete adapters)
│   │   ├── cloudflare/    # Cloudflare external service
│   │   ├── handlers/      # HTTP request handlers
│   │   └── repositories/  # Database repositories
│   ├── apps/
│   │   └── app.go         # App wiring & initialization
│   └── core/              # Core business logic
│       ├── domains/       # Domain entities & business rules
│       └── services/      # Application services (orchestration)
├── libs/
│   ├── conv/
│   └── jwt/
├── main.go                # Application entry point
```

## Struktur layer

- **Core domain** (`internal/core/domains`)
  - Berisi entitas dan aturan bisnis murni.
  - Tidak bergantung pada framework, database, atau library eksternal.

- **Ports** (interface di `internal/core`)
  - Kontrak antara core dan lapisan luar, ditulis sebagai interface.
  - Core bergantung pada interface ini, bukan implementasi konkret.

- **Adapters** (`internal/adapters`)
  - Implementasi konkret dari ports.
  - `repositories`: akses database.
  - `handlers`: terima permintaan HTTP dan panggil service.
  - `cloudflare`: integrasi layanan eksternal Cloudflare.

- **Services** (`internal/core/services`)
  - Koordinasi logika aplikasi antara domain dan port.

- **Apps** (`internal/apps`)
  - Inisialisasi dan wiring semua komponen.

- **main.go**: titik masuk aplikasi, memanggil apps untuk menjalankan server.

## Manfaat

- Bisnis logic dapat diuji tanpa database atau HTTP.
- Database atau framework dapat diganti tanpa mengubah core.
- Tiap layer punya tanggung jawab yang jelas, mudah dipelihara.
