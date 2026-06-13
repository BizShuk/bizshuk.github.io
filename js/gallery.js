/* index.html - Gallery page logic
 * Loads gallery data from data/gallery.json and renders it into the grid.
 */
document.addEventListener('DOMContentLoaded', () => {
    const galleryContainer = document.getElementById('gallery');

    fetch('data/gallery.json')
        .then(response => response.json())
        .then(data => {
            data.forEach((item, index) => {
                // Items with a link become anchors; otherwise render as a decorative block.
                const isLink = item.link && item.link.trim() !== '';
                const elem = document.createElement(isLink ? 'a' : 'div');

                elem.className = 'gallery-item';

                if (isLink) {
                    elem.href = item.link;
                    elem.target = '_blank';
                    elem.rel = 'noopener noreferrer';
                }

                // Extract background image url from style string
                // e.g. "url('https://i.imgur.com/SpbbJ3z.jpg?1')"
                if (item.style && item.style.backgroundImage) {
                    elem.style.backgroundImage = item.style.backgroundImage;
                }

                // Staggered animation
                elem.style.animationDelay = `${index * 0.05}s`;

                if (item.img_title) {
                    const overlay = document.createElement('div');
                    overlay.className = 'gallery-item-overlay';

                    const title = document.createElement('h3');
                    title.className = 'gallery-item-title';
                    title.textContent = item.img_title;

                    overlay.appendChild(title);
                    elem.appendChild(overlay);
                }

                galleryContainer.appendChild(elem);
            });
        })
        .catch(error => console.error('Error loading gallery data:', error));
});
