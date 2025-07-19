from bot.db.activity import insert_activity
from bot.utils.helpers import title_spliter


def post_handler(bot):
    @bot.channel_post_handler(
        content_types = ["text" , "photo" , "video" , "audio" , "voice" , "document"]
    )
    def update_db_content(msg):
        msg_title = title_spliter(msg.text) if msg.content_type == "text" else title_spliter(msg.caption)
        insert_activity(msg.message_id , msg_title)