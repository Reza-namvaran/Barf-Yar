from telebot.types import InlineKeyboardMarkup, InlineKeyboardButton
from bot.db.activity import add_collaborator, get_collaborators

def collaboration_handler(bot):
    @bot.callback_query_handler(func=lambda call: call.data.startswith("collab_join_"))
    def join_collaboration(call):
        try:
            activity_id = int(call.data.split("_")[2])
            user_id = call.from_user.id
            add_collaborator(activity_id, user_id)
            bot.answer_callback_query(call.id, "You have joined the collaboration!")
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

def get_collaboration_menu(activity_id):
    keyboard = InlineKeyboardMarkup()
    keyboard.row(
        InlineKeyboardButton("Join Collaboration", callback_data=f"collab_join_{activity_id}")
    )
    return keyboard
