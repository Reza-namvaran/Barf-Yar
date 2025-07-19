from bot.handlers.start import handle_start

def handle_about(bot):
    @bot.callback_query_handler(func=lambda call:"data1")
    def about(call):
        if call.data == "data1":
            bot.answer_callback_query(call.id,text="Hey")
        elif call.data == "data2":
            pass
