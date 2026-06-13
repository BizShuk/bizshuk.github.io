/* resume.html - Resume page logic
 * - Updates the copyright year in the footer
 * - Adds drag-to-scroll for the horizontal photo gallery on desktop
 */
document.getElementById('year').textContent = new Date().getFullYear();

// Mouse drag scrolling for the photo gallery
const slider = document.querySelector('.photo-gallery');
let isDown = false;
let startX;
let scrollLeft;

slider.addEventListener('mousedown', (e) => {
    isDown = true;
    slider.style.cursor = 'grabbing';
    startX = e.pageX - slider.offsetLeft;
    scrollLeft = slider.scrollLeft;
});

slider.addEventListener('mouseleave', () => {
    isDown = false;
    slider.style.cursor = 'default';
});

slider.addEventListener('mouseup', () => {
    isDown = false;
    slider.style.cursor = 'default';
});

slider.addEventListener('mousemove', (e) => {
    if (!isDown) return;
    e.preventDefault();
    const x = e.pageX - slider.offsetLeft;
    const walk = (x - startX) * 2; // scroll-fast
    slider.scrollLeft = scrollLeft - walk;
});
