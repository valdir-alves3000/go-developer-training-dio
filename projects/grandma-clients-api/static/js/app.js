import { initModal } from './ui/modal.js';
import { initUserForm } from './ui/userForm.js';
import { initAddUserModal,initEditUserModal } from './ui/userActions.js';
import { attachDeleteEvents } from './ui/deleteButtons.js';
import { addFilterButtons } from './ui/filters.js';

document.addEventListener('DOMContentLoaded', () => {
    addFilterButtons();
    attachDeleteEvents();
    initModal();
    initUserForm();
    initAddUserModal();
    initEditUserModal();
});
