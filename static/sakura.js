window.addEventListener('DOMContentLoaded', () => {
  const minSize = 15;
  const maxSize = 45;
  const section = document.querySelector('.cherry-blossom-container');

  const createPetal = () => {
    if (window.innerWidth < 1140) return;
    const petalEl = document.createElement('span');
    petalEl.className = 'petal';
    const size = Math.random() * (maxSize + 1 - minSize) + minSize;
    petalEl.style.width = `${size}px`;
    petalEl.style.height = `${size}px`;
    petalEl.style.left = Math.random() * innerWidth + 'px';
    section.appendChild(petalEl);

    // 一定時間が経てば花びらを消す
    setTimeout(() => {
      petalEl.remove();
    }, 10000);
  }

  setInterval(createPetal, 300);
});