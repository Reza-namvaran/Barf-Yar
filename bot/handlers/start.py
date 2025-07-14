from bot.keyboards.main_menu import get_main_menu

def handle_start(bot):
    @bot.message_handler(commands=['start'])
    def start(message):
        pass