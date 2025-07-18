import psycopg2
from bot.db import get_conn

def insert_activity(message_id: int, title: str):
    conn = get_conn()
    cur = conn.cursor()
    cur.execute(
        "INSERT INTO activities (message_id, title) VALUES (%s, %s)",
        (message_id, title)
    )
    conn.commit()
    cur.close()
    conn.close()

def get_activities():
    conn = get_conn()
    cur = conn.cursor()
    cur.execute("SELECT id, title FROM activities")
    rows = cur.fetchall()
    cur.close()
    conn.close()
    return rows
