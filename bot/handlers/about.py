
def handle_about(bot):
    @bot.callback_query_handler(func=lambda call : call.data == "data1")
    def about(call):
        bot.answer_callback_query(call.id,text="Hey")
