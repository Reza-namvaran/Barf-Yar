from bot.db.activity import get_activities
from telebot.types import InlineKeyboardMarkup, InlineKeyboardButton

ACTIVITY_IN_ROW = 3

def activity_handler(bot):
    @bot.message_handler(func=lambda message: message.text == "📋 فعالیت های برف نو")
    def show_activity_list(message):
        try:
            markup = InlineKeyboardMarkup(row_width = ACTIVITY_IN_ROW)
            all_activities = get_activities()
            buttons = [
                InlineKeyboardButton (
                    text = f"{title}",
                    callback_data = f"activity_{id}"
                )
                for id , title in all_activities
            ]
            markup.add(*buttons)
            bot.send_message(message.chat.id , "حالا میتوانید لیست فعالیت های خیریه ما رو ببینید" , reply_markup = markup)
        except Exception as e:
            print(f"[Error] Unexpected error happend : {e}")