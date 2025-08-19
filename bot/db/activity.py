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
    cur.execute("SELECT id , title FROM activities")
    rows = cur.fetchall()
    cur.close()
    conn.close()
    return rows

def get_activities_by_id(activity_id: int):
    conn = get_conn()
    cur = conn.cursor()
    cur.execute("SELECT message_id, title FROM activities WHERE id = %s", (activity_id,))
    row = cur.fetchone()
    cur.close()
    conn.close()
    return row

def update_activity_title(message_id: int, new_title: str):
    conn = get_conn()
    cur = conn.cursor()
    cur.execute(
        "UPDATE activities SET title = %s WHERE message_id = %s",
        (new_title , message_id)
    )
    conn.commit()
    cur.close()
    conn.close()
    
def add_collaborator(activity_id: int, user_id: int):
    conn = get_conn()
    cur = conn.cursor()
    try:
        cur.execute(
            "INSERT INTO activity_supporters (activity_id, user_id) VALUES (%s, %s)",
            (activity_id, user_id)
        )
        conn.commit()
    except psycopg2.IntegrityError:
        # User already a collaborator (UNIQUE constraint violation)
        pass
    finally:
        cur.close()
        conn.close()

def get_collaborators(activity_id: int):
    conn = get_conn()
    cur = conn.cursor()
    cur.execute(
        "SELECT user_id FROM activity_supporters WHERE activity_id = %s",
        (activity_id,)
    )
    rows = cur.fetchall()
    cur.close()
    conn.close()
    return [row[0] for row in rows]