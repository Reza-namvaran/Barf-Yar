// function initiateExport(button) {
//   const format = document.getElementById("export-format").value;
//   const activityId = button.dataset.activityId;
//   const activityTitle = button.dataset.activityTitle;

//   if (!format) {
//     alert("Please select an export format.");
//     return;
//   }

//   button.disabled = true;
//   button.innerText = "Downloading...";

//   fetch(`/dashboard/activities/${activityId}/supporters/export/${format}`)
//     .then((res) => {
//       if (!res.ok) return res.text().then((msg) => { throw new Error(msg); });
//       return res.blob();
//     })
//     .then((blob) => {
//       const url = URL.createObjectURL(blob);
//       const today = new Date().toISOString().split("T")[0];
//       const safeTitle = activityTitle.replace(/[^a-z0-9_-]/gi, "_");
//       const fileName = `${safeTitle}_supporters_${today}.${format}`;

//       const link = document.createElement("a");
//       link.href = url;
//       link.download = fileName;
//       link.click();
//       URL.revokeObjectURL(url);
//     })
//     .catch((err) => alert("Error: " + err.message))
//     .finally(() => {
//       button.disabled = false;
//       button.innerText = "Export";
//     });
// }
