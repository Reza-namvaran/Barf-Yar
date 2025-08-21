document.addEventListener('DOMContentLoaded', function() {
  const form = document.getElementById('addForm');
  if (form) {
    form.addEventListener('submit', function(e) {
      e.preventDefault();
      const formData = {
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
  }
});

// Delete
let deleteId = null;
const deleteModal = document.getElementById("deleteModal");
const cancelDelete = document.getElementById("cancelDelete");
const confirmDelete = document.getElementById("confirmDelete");

function openDeleteModal(id) {
  deleteId = id;
  deleteModal.style.display = "flex";
}

function closeDeleteModal() {
  deleteId = null;
  deleteModal.style.display = "none";
}

cancelDelete.addEventListener("click", closeDeleteModal);

confirmDelete.addEventListener("click", () => {
  if (!deleteId) return;

  fetch(`/dashboard/activities/${deleteId}`, { method: "DELETE"}).then(res => {
    if (res.ok) {
      document.getElementById(`row-${deleteId}`).remove();
    } else {
      res.text().then(alert);
    }
  }).finally(closeDeleteModal)
});

window.addEventListener("click", e => {
  if (e.target === deleteModal) closeDeleteModal();
});

// Edit
function toggleEdit(id) {
  const row = document.getElementById(`row-${id}`);
  const cells = row.querySelectorAll("td");
  const messageId = cells[1].innerText.trim();
  const title = cells[2].innerText.trim();
  const prompt = cells[3].innerText.trim() !== "—" ? cells[3].innerText.trim() : "";

  // Replace cells with input fields
  cells[1].innerHTML = `<input type="number" id="msg-${id}" value="${messageId}">`;
  cells[2].innerHTML = `<input type="text" id="title-${id}" value="${title}">`;
  cells[3].innerHTML = `<input type="number" id="prompt-${id}" value="${prompt}">`;

  // Replace actions with Save/Cancel
  cells[4].innerHTML = `
    <button class="btn btn-save" onclick="saveEdit(${id})">Save</button>
    <button class="btn btn-cancel" onclick="location.reload()">Cancel</button>
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
