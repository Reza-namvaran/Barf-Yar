from dotenv import load_dotenv
from bot.handlers import start, about
import telebot
import os

load_dotenv()
BOT_TOKEN = os.getenv("BOT_TOKEN")

bot = telebot.TeleBot(BOT_TOKEN, parse_mode="HTML")

start.handle_start(bot)
about.handle_about(bot)

print("Bot is running")
bot.infinity_polling()