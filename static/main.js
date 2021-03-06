// Constants for time calculations
const SECOND = 1000;
const MINUTE = 60 * SECOND;
const HOUR = 60 * MINUTE;

// Configure Sentry
Sentry.init({
  dsn:
    "https://dcb228402ca04737a8efc0ef1486eb3d@o538041.ingest.sentry.io/5656062",
  maxBreadcrumbs: 50,
  debug: true,
});

// Function to handle fetching status
const fetchStatus = async () => {
  const resp = await fetch("/status");
  const data = await resp.json();

  return data.IsOpen ? "Open" : "Closed";
};

// Update status in DOM
const updateStatus = (status) => {
  document.getElementById("currentStatus").textContent = status;
};

// Send request to toggle door
const toggleDoor = async () => {
  await fetch("/toggle");
  // Update status after 30 seconds
  await new Promise((_) => setTimeout(_, 30 * SECOND));
  const status = await fetchStatus();
  updateStatus(status);
};

// Fetch status on page load
fetchStatus().then(updateStatus);

// Toggle button handler
document.getElementById("toggleButton").addEventListener("click", toggleDoor);

// Refetch status every 10 minutes
setInterval(async () => {
  const status = await fetchStatus();
  updateStatus(status);
}, 5 * MINUTE);
