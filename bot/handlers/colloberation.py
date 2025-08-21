from telebot.types import InlineKeyboardMarkup, InlineKeyboardButton
from bot.db.activity import add_collaborator, get_collaborators, get_support_prompt
from dotenv import load_dotenv
import os

load_dotenv(override=True)

PRIVATE_CHANNEL_ID = os.getenv("PRIVATE_CHANNEL_ID")

if PRIVATE_CHANNEL_ID is None:
    raise ValueError("PRIVATE_CHANNEL_ID is not set in .env file.")
PRIVATE_CHANNEL_ID = int(PRIVATE_CHANNEL_ID)

def collaboration_handler(bot):
    @bot.callback_query_handler(func=lambda call: call.data.startswith("collab_join_"))
    def join_collaboration(call):
        try:
            activity_id = int(call.data.split("_")[2])
            user_id = call.from_user.id
            add_collaborator(activity_id, user_id)

            bot.answer_callback_query(call.id, "✅ You have joined the collaboration!")

            bot.send_message(
                call.message.chat.id,
                f"{call.from_user.first_name} عزیز\nf"{call.from_user.first_name} عزیز\nاز مشارکت شما در این فعالیت ممنونیم"
از مشارکت شما در این فعالیت ممنونیم"
            )

        except Exception as e:
            print(f"[Error] Unexpected error in joining collaboration: {e}")
            bot.answer_callback_query(call.id, "Failed to join the collaboration.")
            
    @bot.callback_query_handler(func=lambda call: call.data.startswith("collab_view_"))
    def view_collaborators(call):
        try:
            activity_id = int(call.data.split("_")[2])
            collaborators = get_collaborators(activity_id)
            if collaborators:
                message = f"Collaborators for activity {activity_id}:\n" + "\n".join(str(uid) for uid in collaborators)
            else:
                message = f"No collaborators for activity {activity_id} yet."
            bot.answer_callback_query(call.id, message, show_alert=True)
        except Exception as e:
            print(f"[Error] Unexpected error in viewing collaborators: {e}")
            bot.answer_callback_query(call.id, "Failed to view collaborators.")

    @bot.callback_query_handler(func=lambda call: call.data.startswith("collab_info_"))
    def show_info(call):
        try:
            activity_id = int(call.data.split("_")[2])
            prompt = get_support_prompt(activity_id)

            if not prompt:
                bot.answer_callback_query(call.id, "No support prompt for this activity.")
                return

            prompt_message_id = prompt[0]

            print(prompt_message_id)
            # Attach join button for this prompt
            join_keyboard = InlineKeyboardMarkup()
            join_keyboard.row(
                InlineKeyboardButton("میخواهم مشارکت کنم", callback_data=f"collab_join_{activity_id}")
            )

            bot.copy_message(
                chat_id=call.message.chat.id,
                from_chat_id=PRIVATE_CHANNEL_ID,
                message_id=prompt_message_id,
                reply_markup=join_keyboard
            )

            bot.answer_callback_query(call.id)
        except Exception as e:
            print(f"[Error] show_info failed: {e}")
            bot.answer_callback_query(call.id, "Error fetching info")


def get_collaboration_menu(activity_id):
    keyboard = InlineKeyboardMarkup()
    keyboard.row(
        InlineKeyboardButton("اطلاعات بیشتر", callback_data=f"collab_info_{activity_id}")
    )
    return keyboard
