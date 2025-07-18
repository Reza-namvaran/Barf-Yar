from telebot.types import InlineKeyboardMarkup, InlineKeyboardButton, ReplyKeyboardMarkup, KeyboardButton

def get_main_menu():
    first_markup = InlineKeyboardMarkup(row_width=2)
    info_button = InlineKeyboardButton(text="Info",callback_data="data1")
    activity_button = InlineKeyboardButton(text="Activity",callback_data="data2")
    first_markup.add(activity_button,info_button)
    return first_markup
