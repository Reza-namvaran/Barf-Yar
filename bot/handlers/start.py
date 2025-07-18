from bot.keyboards.main_menu import get_main_menu

def handle_start(bot):
    @bot.message_handler(commands=['start'])
    def start(message):
        bot.send_message(message.chat.id,text="Welcome",reply_markup = get_main_menu())