

# برف‌یار (Barf-Yar) ❄️  
Telegram bot for managing charity activities, contributors, and broadcast targeted reminders.  

## ✨ Features
- **Telegram Bot**
  - List activities dynamically
  - Uses an archive channel to reduce media load
  - Forward activity & support prompts
  - Join to contributors
  - Track supporters for each activity

- **Admin Panel (Go + PostgreSQL)**
  - Secure login
  - Manage activities (CRUD)
  - Assign **support prompts** to activities
  - View supporters of each activity
  - Export supporters (JSON, CSV, PDF)

---

## How to test
Make sure you have a `.env` file with variables in `.env.example`

### 🐳 Docker Setup (Admin Panel)
1. **Build and start**
   ```bash
   sudo docker-compose up --build
   ``` 

Panel will run on → `http://localhost:8080`

2. **Create an admin user**
   ```bash
   sudo docker-compose exec panel /app/create_admin
   ```
3. **Working with DB**
   ```bash
   sudo docker-compose exec db psql -U <user> -d <database>
   ```
---

## 🤖 Telegram Bot Setup

1. Make sure your `.env` file contains:

   ```env
   BOT_TOKEN=your-telegram-bot-token
   PRIVATE_CHANNEL_ID=CHANNEL_ID_TO_ARCHIVE_MESSAGES
   DATABASE_URL=postgres://<username>:<password>@<host>:<port>/<DB_NAME>?sslmode=disable
   ```

2. Install dependencies (Python 3.10+):

   ```bash
   pip install -r requirements.txt
   ```

3. Run the bot:

   ```bash
   python -m bot.main
   ```

---
