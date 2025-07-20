from bot.db.activity import get_activities_by_id
from telebot.types import InlineKeyboardMarkup, InlineKeyboardButton
from dotenv import load_dotenv
import os

load_dotenv(override=True)

PRIVATE_CHANNEL_ID = os.getenv("PRIVATE_CHANNEL_ID")

if PRIVATE_CHANNEL_ID is None:
    raise ValueError("PRIVATE_CHANNEL_ID is not set in .env file.")
PRIVATE_CHANNEL_ID = int(PRIVATE_CHANNEL_ID)


def forward_handler(bot):
    @bot.callback_query_handler(func=lambda call: call.data.startswith("activity_"))
    def forward(call):
        try:
            activity_id = int(call.data.split("_")[1])
            data = get_activities_by_id(activity_id)

            if not data:
                bot.answer_callback_query(call.id, "Activity not found!")
                return

            message_id, title = data

            bot.forward_message(
                chat_id=call.message.chat.id,
                from_chat_id=PRIVATE_CHANNEL_ID,
                message_id=message_id
            )

            bot.answer_callback_query(call.id, f"Activity: {title}")

        except Exception as e:
            print(f"[Error] Unexpected error in forwarding activity: {e}")
            bot.answer_callback_query(call.id, "Failed to forward the activity.")
