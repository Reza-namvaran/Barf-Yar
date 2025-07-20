
def handle_about(bot):
    @bot.message_handler(func=lambda message: message.text == "درباره برف نو")
    def about(message):
        bot.send_message(message.chat.id, "توضیحات مربوط به انجمن")
