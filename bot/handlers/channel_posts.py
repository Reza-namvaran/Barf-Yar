from bot.db.activity import insert_activity
from bot.utils.helpers import title_spliter


def post_handler(bot):
    @bot.channel_post_handler(
        content_types = ["text" , "photo" , "video" , "audio" , "voice" , "document"]
    )
    def update_db_content(msg):
        try:
            msg_title = title_spliter(msg.text) if msg.content_type == "text" else title_spliter(msg.caption or "")
            insert_activity(msg.message_id , msg_title)
        except Exception as e:
            print(f"[Error] Unexpected error happend : {e}")