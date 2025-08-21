

# Barf Yar ❄️  
Telegram bot for managing charity activities, contributors, and broadcast targeted reminders.  

## ✨ Features
- **Telegram Bot (pyTelegramBotAPI)**
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

## 🏫 About the Project
Barf Yar is a collaborative project developed by students of the **Computer Engineering Department at Bu-Ali Sina University (BASU)**, under the guidance of the **BASU Computer Engineering Student Association**.  

This system was designed to solve a real need for **Barf-No(برف نو)** charity: managing activities, collaborators, and targeted reminders in an efficient and scalable way.  

By combining a modern **admin panel** (Go + PostgreSQL) with a dynamic **Telegram bot**, the project empowers charity administrators to:  
- Coordinate events and activities more effectively  
- Engage volunteers through Telegram with minimal effort  
- Organize supporters using dynamic labels and broadcast reminders at the right time  

This project demonstrates the ability of BASU students to deliver impactful, community-driven solutions that blend **computer science education** with **real-world social responsibility**.  

---

## Contributors
This project is made possible thanks to the dedication and collaboration of our amazing team from BASU Computer Engineering Student Association:

* **Daniel Keshavarz Nejad** – Project Lead, guiding the vision and direction of the project
* **Reza Namvaran** – Technical Lead / Manager, coordinating development and ensuring technical excellence
* **Saeed Mazaheri** – Back-end Developer & Mentor, providing guidance and core backend solutions
* **Hossein Fazel** – Back-end / Bot Developer & Mentor, supporting development and mentoring team members
* **Pouya Tavakoli** – Back-end Developer, implementing key features and logic
* **Maryam Nokohan** – Intern, contributing with enthusiasm and fresh ideas
* **Taha Sadeghi** – Intern, assisting with development and learning through practice
* **Kasra Ali Rezaee** – Intern, supporting tasks and gaining hands-on experience

> Every contribution has made this project stronger—thank you all for your hard work and dedication!

<a href="https://github.com/Reza-namvaran/barf-yar/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=Reza-namvaran/barf-yar" />
</a>

Made with [contrib.rocks](https://contrib.rocks).
