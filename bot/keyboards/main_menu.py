from telebot.types import ReplyKeyboardMarkup, KeyboardButton

def get_main_menu():
    markup = ReplyKeyboardMarkup(resize_keyboard=True, row_width=2)
    info_button = KeyboardButton(text="درباره برف نو")
    activity_button = KeyboardButton(text="📋 فعالیت های برف نو")
    markup.add(activity_button, info_button)
    return markup
