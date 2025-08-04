from dotenv import load_dotenv
import os

load_dotenv(override=True)

PRIVATE_CHANNEL_ID = os.getenv("PRIVATE_CHANNEL_ID")
ABOUT_ID=os.getenv("ABOUT_ID")

def handle_about(bot):
    @bot.message_handler(func=lambda message: message.text == "درباره برف نو")
    def about(message):
        bot.copy_message(
                chat_id=message.chat.id,
                from_chat_id=PRIVATE_CHANNEL_ID,
                message_id=ABOUT_ID
            )
