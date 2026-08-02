// Function to toggle between dark and light mode
const toggleDarkMode = () => {
    const body = document.body;
    body.classList.toggle('dark');
    document.getElementById('dark-mode-toggle').setAttribute('aria-pressed', body.classList.contains('dark'));
    
    // Optionally, store the preference in localStorage
    if (body.classList.contains('dark')) {
        localStorage.setItem('dark-mode', 'enabled');
    } else {
        localStorage.setItem('dark-mode', 'disabled');
    }
};

// Check localStorage for dark mode preference on load
if (localStorage.getItem('dark-mode') === 'enabled') {
    document.body.classList.add('dark');
}

document.getElementById('dark-mode-toggle').setAttribute('aria-pressed', document.body.classList.contains('dark'));
