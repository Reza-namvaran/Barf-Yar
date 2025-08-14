from dotenv import load_dotenv
from bot.handlers import start, about , activity , forward_handler
import telebot
import os

load_dotenv(override = True)
BOT_TOKEN = os.getenv("BOT_TOKEN")

bot = telebot.TeleBot(BOT_TOKEN, parse_mode="HTML")

start.handle_start(bot)
about.handle_about(bot)
activity.activity_handler(bot)
forward_handler.forward_handler(bot)

print("Bot is running")
bot.infinity_polling()