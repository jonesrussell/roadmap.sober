console.log('main.js loaded');

// Get the close panel button and slide over elements
const closePanelButton = document.getElementById('close-panel');
const slideOver = document.getElementById('slide-over');

// Function to open the drawer
function openDrawer() {
  console.log('openDrawer function called');
  slideOver.classList.remove('hidden');
}

// Function to dispatch the openDrawer event
function dispatchOpenDrawerEvent() {
  console.log('dispatchOpenDrawerEvent function called');
  document.dispatchEvent(new CustomEvent('openDrawerEvent'));
}

// Add event listener to the close panel button
closePanelButton.addEventListener('click', () => {
  slideOver.classList.add('hidden');
});

// Add event listener to the document
document.addEventListener('openDrawerEvent', openDrawer);
