## [v1.0.0] - 2025-08-21
### Added
- Initial **Admin Panel** (Go + PostgreSQL + Docker):
  - Manage activities (create, edit, deletes).
  - View supporters linked to activities.
  - Export supporters of an activity (PDF, CSV, JSON).
  - Simple dashboard with activity count and system status.
- Initial **Telegram Bot** (Python + Telebot):  
  - Forward activities from private channel.  
  - Join collaboration button.  
  - Supporter tracking in PostgreSQL.  

### Infrastructure
- Docker setup for Go-based panel and psql db.  
- `.env` support for secrets.  