from bot.db.activity import get_activities
from telebot.types import InlineKeyboardMarkup, InlineKeyboardButton

ACTIVITY_IN_ROW = 3

def activity_handler(bot):
    @bot.callback_query_handler(func = lambda call : call.data == "data2")
    def show_activity_list(call):
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
            bot.send_message(call.message.chat.id , "حالا میتوانید لیست فعالیت های خیریه ما رو ببینید" , reply_markup = markup)
        except Exception as e:
            print(f"[Error] Unexpected error happend : {e}")