import { openModal } from './modal.js';

export function initEditUserModal() {
    const modalTitle = document.getElementById('modal-title');

    document.querySelectorAll('.action-button.view, .action-button.edit').forEach(btn => {
        btn.addEventListener('click', async function () {
            const userId = this.getAttribute('data-id');
            const isEdit = this.classList.contains('edit');
            const name = this.parentElement.parentElement.getAttribute('data-name');
            const age = this.parentElement.parentElement.getAttribute('data-age');
            const city = this.parentElement.parentElement.getAttribute('data-city');
            const state = this.parentElement.parentElement.getAttribute('data-state');

            document.getElementById('user-id').value = userId;
            document.getElementById('user-name').value = name;
            document.getElementById('user-age').value = age;
            document.getElementById('user-city').value = city;
            document.getElementById('user-state').value = state;

            const isReadOnly = !isEdit;
            ['user-name', 'user-age', 'user-city', 'user-state'].forEach(id => {
                document.getElementById(id).readOnly = isReadOnly;
            });

            modalTitle.textContent = isEdit ? 'Edit User' : 'View User Details';
            document.querySelector('.btn-primary').style.display = isEdit ? 'block' : 'none';

            openModal();

        });
    });
}

export function initAddUserModal() {
    const addUserBtn = document.getElementById('add-user-btn');
    const userForm = document.getElementById('user-form');
    const modal = document.getElementById('user-modal');
    const modalTitle = document.getElementById('modal-title');

    addUserBtn.addEventListener('click', function () {
        userForm.reset();
        document.getElementById('user-id').value = '';

        ['user-name', 'user-age', 'user-city', 'user-state'].forEach(id => {
            document.getElementById(id).readOnly = false;
        });

        modalTitle.textContent = 'Add New User';
        document.querySelector('.btn-primary').style.display = 'block';

        modal.classList.add('active');
    });
}