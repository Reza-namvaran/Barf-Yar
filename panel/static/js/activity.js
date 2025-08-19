document.getElementById('addForm').addEventListener('submit', function(e) {
    e.preventDefault();
    const formData = {
        id: parseInt(this.id.value),
        message_id: parseInt(this.message_id.value),
        title: this.title.value,
        prompt_message_id: this.prompt_message_id.value ? parseInt(this.prompt_message_id.value) : null
    };
    fetch(this.action, {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        body: JSON.stringify(formData)
    }).then(res => {
        if (res.ok) location.reload();
        else res.text().then(alert);
    });
});

function deleteActivity(id) {
    fetch(`/dashboard/activities/${id}`, { method: 'DELETE' }).then(res => {
        if (res.ok) location.reload();
        else res.text().then(alert);
    });
}

function toggleEdit(id) {
    const row = document.getElementById(`row-${id}`);
    const cells = row.querySelectorAll("td");
    const messageId = cells[1].innerText.trim();
    const title = cells[2].innerText.trim();
    const prompt = cells[3].innerText.trim() !== "-" ? cells[3].innerText.trim() : "";
  
    // Replace cells with input fields
    cells[1].innerHTML = `<input type="number" id="msg-${id}" value="${messageId}">`;
    cells[2].innerHTML = `<input type="text" id="title-${id}" value="${title}">`;
    cells[3].innerHTML = `<input type="number" id="prompt-${id}" value="${prompt}">`;
  
    // Replace actions with Save/Cancel
    cells[4].innerHTML = `
      <button onclick="saveEdit(${id})">Save</button>
      <button onclick="location.reload()">Cancel</button>
    `;
}
  
function saveEdit(id) {
    const messageId = parseInt(document.getElementById(`msg-${id}`).value);
    const title = document.getElementById(`title-${id}`).value;
    let promptVal = document.getElementById(`prompt-${id}`).value;
    const promptId = promptVal ? parseInt(promptVal) : null;
  
    fetch(`/dashboard/activities/${id}`, {
      method: 'PUT',
      headers: { 'Content-Type': 'application/json' },
      body: JSON.stringify({
        id: id,
        message_id: messageId,
        title: title,
        prompt_message_id: promptId
      })
    }).then(res => {
      if (res.ok) location.reload();
      else res.text().then(alert);
    });
  }