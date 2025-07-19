from dotenv import load_dotenv
from bot.handlers import start, about , channel_posts , activity
import telebot
import os

load_dotenv(override = True)
BOT_TOKEN = os.getenv("BOT_TOKEN")

bot = telebot.TeleBot(BOT_TOKEN, parse_mode="HTML")

start.handle_start(bot)
about.handle_about(bot)
channel_posts.post_handler(bot)
activity.activity_handler(bot)

print("Bot is running")
bot.infinity_polling()