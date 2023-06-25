self.addEventListener("activate", async () => {
  self.showNotification("Vibration Sample", {
    body: "Buzz! Buzz!",
    icon: "../images/touch/chrome-touch-icon-192x192.png",
    vibrate: [200, 100, 200, 100, 200, 100, 200],
    tag: "vibration-sample",
  });
});
