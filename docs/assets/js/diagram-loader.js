// Diagram loading functionality for DUKPT documentation
document.addEventListener('DOMContentLoaded', async function() {
    // Set theme based on system preference
    const prefersDark = window.matchMedia('(prefers-color-scheme: dark)').matches;
    document.documentElement.setAttribute('data-bs-theme', prefersDark ? 'dark' : 'light');
    
    // Check if Mermaid is loaded first
    if (typeof mermaid !== 'undefined') {
        const mermaidTheme = prefersDark ? 'dark' : 'neutral';
        mermaid.initialize({ startOnLoad: false, theme: mermaidTheme });
        
        // Load diagrams dynamically from external .mmd files
        const diagrams = document.querySelectorAll('.mermaid[data-source]');
        await Promise.all(Array.from(diagrams).map(async (diagram) => {
            try {
                const response = await fetch(`assets/diagrams/${diagram.dataset.source}.mmd`);
                const content = await response.text();
                diagram.textContent = content;
            } catch (error) {
                diagram.innerHTML = `<div class="alert alert-danger">Failed to load diagram: ${diagram.dataset.source}<br><small>${error.message}</small></div>`;
            }
        }));
        
        // Render all loaded diagrams
        mermaid.run();
    } else {
        console.warn('Mermaid.js failed to load from CDN. Diagrams will show as text.');
        // Add a notice to the page
        const notice = document.createElement('div');
        notice.className = 'alert alert-warning mx-3 mt-3';
        notice.innerHTML = '<strong>Note:</strong> This page uses Mermaid.js from an external CDN for diagram rendering. If diagrams appear as text, please ensure external resources can be loaded.';
        document.body.insertBefore(notice, document.body.firstChild);
    }
    
    // Listen for system theme changes
    window.matchMedia('(prefers-color-scheme: dark)').addEventListener('change', (e) => {
        document.documentElement.setAttribute('data-bs-theme', e.matches ? 'dark' : 'light');
        // Note: Mermaid diagrams will require page refresh for theme change
    });
});