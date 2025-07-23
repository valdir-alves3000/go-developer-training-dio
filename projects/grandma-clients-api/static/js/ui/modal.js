export function initModal() {
    const modal = document.getElementById('user-modal');
    const closeBtn = document.getElementById('modal-close');
    const cancelBtn = document.getElementById('cancel-btn');

    function closeModal() {
        modal.classList.remove('active');
    }

    [closeBtn, cancelBtn].forEach(btn => btn.addEventListener('click', closeModal));

    modal.addEventListener('click', e => {
        if (e.target === modal) closeModal();
    });
}

export function openModal() {
    document.getElementById('user-modal').classList.add('active');
}
