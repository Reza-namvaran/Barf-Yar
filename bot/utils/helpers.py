def title_spliter(text: str):
    if not text:
        return "بدون عنوان"
    title = text.strip().split("\n")[0]
    return title.strip() or "بدون عنوان"
