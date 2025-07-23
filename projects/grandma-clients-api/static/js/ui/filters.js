export function addFilterButtons() {
    const filterContainer = document.querySelector('.filter-container');
    const users = document.querySelectorAll('.user-list-container .user');

    filterContainer.querySelectorAll(".filter-btn:not([data-filter='all'])").forEach(btn => btn.remove());

    const uniqueStates = [...new Set(Array.from(users).map(user => user.getAttribute('data-state')).filter(Boolean))];

    uniqueStates.forEach(state => {
        const btn = document.createElement('button');
        btn.classList.add('filter-btn');
        btn.dataset.filter = state;
        btn.textContent = state;
        filterContainer.appendChild(btn);
    });

    initFilterEvents();
}

function initFilterEvents() {
    document.querySelectorAll('.filter-btn').forEach(btn => {
        btn.addEventListener('click', function () {
            const filter = this.dataset.filter;

            document.querySelectorAll('.filter-btn').forEach(b => b.classList.remove('active'));
            this.classList.add('active');

            document.querySelectorAll('.user').forEach(user => {
                const match = filter === 'all' || user.dataset.state === filter;
                if (match) {
                    user.style.display = 'block';
                    setTimeout(() => {
                        user.style.opacity = 1;
                        user.style.transform = 'translateY(0)';
                    }, 100);
                } else {
                    user.style.opacity = 0;
                    user.style.transform = 'translateY(20px)';
                    setTimeout(() => {
                        user.style.display = 'none';
                    }, 300);
                }
            });
        });
    });
}
